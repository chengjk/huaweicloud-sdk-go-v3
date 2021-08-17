package model

import (
	"encoding/json"

	"strings"
)

// 批量删除测试用例请求体
type BatchDeleteTestCaseRequestBody struct {
	// 注册结果返回的服务id

	ServiceId int32 `json:"service_id"`
	// 测试用例唯一标识，数组长度小于50个

	TestcaseIdList []string `json:"testcase_id_list"`
}

func (o BatchDeleteTestCaseRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchDeleteTestCaseRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteTestCaseRequestBody", string(data)}, " ")
}