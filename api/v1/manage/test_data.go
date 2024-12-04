package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/response"
	manageReq "main.go/model/manage/request"
)

type TestDataApi struct {
}

// TestDataList get test data
func (m *TestDataApi) TestDataList(c *gin.Context) {
	global.GVA_LOG.Info("get test data")
	var pageInfo manageReq.TestDataSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := testDataService.GetTestDataList(pageInfo); err != nil {
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
