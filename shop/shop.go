package shop

type Result struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Owner      int    `json:"owner"`
	NickName   string `json:"nickName"`
	CreateTime string `json:"createTime"`
}
