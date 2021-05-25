package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCustomPolicyRequest struct {
	// 待删除的自定义策略ID，获取方式请参见：[自定义策略ID](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=ListCustomPolicies)。

	RoleId string `json:"role_id"`
}

func (o DeleteCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteCustomPolicyRequest", string(data)}, " ")
}
