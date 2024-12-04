package request

import (
	"main.go/model/common/request"
	"main.go/model/manage"
)

type TestDataSearch struct {
	manage.TestData
	request.PageInfo
}
