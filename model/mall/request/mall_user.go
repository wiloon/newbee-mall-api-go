package request

// RegisterUserParam 用户注册
type RegisterUserParam struct {
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
}

// UpdateUserInfoParam 更新用户信息
type UpdateUserInfoParam struct {
	NickName      string `json:"nickName"`
	PasswordMd5   string `json:"passwordMd5"`
	IntroduceSign string `json:"introduceSign"`
}

type UserLoginParam struct {
	LoginName   string `json:"loginName"`
	PasswordMd5 string `json:"passwordMd5"`
}
type AdminSaveMemberParam struct {
	NickName         string `json:"nickName"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	RecipientName    string `json:"recipientName"`
	RecipientMobile  string `json:"recipientMobile"`
	RecipientAddress string `json:"recipientAddress"`
}
