package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/common/response"
	"main.go/model/example"
	"main.go/model/mall"
	mallReq "main.go/model/mall/request"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
	"time"
)

type ManageAdminUserApi struct {
}

// 创建AdminUser
func (m *ManageAdminUserApi) CreateAdminUser(c *gin.Context) {
	var params manageReq.MallAdminParam
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.AdminUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mallAdminUser := manage.MallAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: utils.MD5V([]byte(params.LoginPassword)),
	}
	if err := mallAdminUserService.CreateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败:", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// 修改密码
func (m *ManageAdminUserApi) UpdateAdminUserPassword(c *gin.Context) {
	var req manageReq.MallUpdatePasswordParam
	_ = c.ShouldBindJSON(&req)
	//mallAdminUserName := manage.MallAdminUser{
	//	LoginPassword: utils.MD5V([]byte(req.LoginPassword)),
	//}
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminPassWord(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

// 更新用户名
func (m *ManageAdminUserApi) UpdateAdminUserName(c *gin.Context) {
	var req manageReq.MallUpdateNameParam
	_ = c.ShouldBindJSON(&req)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminName(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AdminUserProfile 用id查询AdminUser
func (m *ManageAdminUserApi) AdminUserProfile(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, mallAdminUser := mallAdminUserService.GetMallAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		response.OkWithData(mallAdminUser, c)
	}
}
func (m *ManageAdminUserApi) AllUser(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, allUser := mallAdminUserService.GetAllUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		response.OkWithData(allUser, c)
	}
}
func (m *ManageAdminUserApi) AllShops(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, shops := mallAdminUserService.GetAllShops(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       shops,
			TotalCount: 200,
			CurrPage:   0,
			PageSize:   200,
		}, "获取成功", c)
	}
}
func (m *ManageAdminUserApi) AdminSaveMember(c *gin.Context) {
	var memberParams mallReq.AdminSaveMemberParam
	_ = c.ShouldBindJSON(&memberParams)

	adminToken := c.GetHeader("token")
	_ = adminToken
	//todo admin token check

	// check if member exist
	var memberList []mall.MallUser
	global.GVA_DB.Where("nick_name=? or login_name=?", memberParams.NickName, memberParams.Username).Find(&memberList)
	if len(memberList) > 0 {
		response.FailWithMessage("会员已存在", c)
		return
	}
	var member mall.MallUser
	member.NickName = memberParams.NickName
	member.LoginName = memberParams.Username
	member.PasswordMd5 = memberParams.Password
	member.IsDeleted = 0
	member.LockedFlag = 0
	member.CreateTime = common.JSONTime{Time: time.Now()}

	global.GVA_LOG.Info(fmt.Sprintf("admin save member, params: %v", memberParams))

	if err := global.GVA_DB.Save(&member).Error; err != nil {
		response.FailWithMessage("failed to create member: "+err.Error(), c)
	}

	var defaultAddress mall.MallUserAddress
	defaultAddress.CreateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UpdateTime = common.JSONTime{Time: time.Now()}
	defaultAddress.UserId = member.UserId
	defaultAddress.UserName = memberParams.RecipientName
	defaultAddress.UserPhone = memberParams.RecipientMobile
	defaultAddress.DetailAddress = memberParams.RecipientAddress
	defaultAddress.DefaultFlag = 1
	err := global.GVA_DB.Save(&defaultAddress).Error
	if err != nil {
		response.FailWithMessage("failed to create member: "+err.Error(), c)
	}
	response.OkWithData("ok", c)
}

// AdminLogin 管理员登陆
func (m *ManageAdminUserApi) AdminLogin(c *gin.Context) {
	var adminLoginParams manageReq.MallAdminLoginParam
	_ = c.ShouldBindJSON(&adminLoginParams)
	if err, _, adminToken := mallAdminUserService.AdminLogin(adminLoginParams); err != nil {
		response.FailWithMessage("登陆失败", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}

// AdminLogout 登出
func (m *ManageAdminUserApi) AdminLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallAdminUserTokenService.DeleteMallAdminUserToken(token); err != nil {
		response.FailWithMessage("登出失败", c)
	} else {
		response.OkWithMessage("登出成功", c)
	}

}

// UserList 商城注册用户列表
func (m *ManageAdminUserApi) UserList(c *gin.Context) {
	var pageInfo manageReq.MallUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallUserService.GetMallUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// LockUser 用户禁用与解除禁用(0-未锁定 1-已锁定)
func (m *ManageAdminUserApi) LockUser(c *gin.Context) {
	lockStatus, _ := strconv.Atoi(c.Param("lockStatus"))
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallUserService.LockUser(IDS, lockStatus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UploadFile 上传单图
// 此处上传图片的功能可用，但是原前端项目的图片链接为服务器地址，如需要显示图片，需要修改前端指向的图片链接
func (m *ManageAdminUserApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}

	// 这里直接使用本地的url
	urlPrefix := global.GVA_CONFIG.Local.Imageurlprefix
	response.OkWithData(urlPrefix+file.Url, c)
}
