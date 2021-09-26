package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListIgnoreRuleRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// policyid

	PolicyId string `json:"policy_id"`
	// 页码

	Page *int32 `json:"page,omitempty"`
	// 每页的条数

	Pagesize *int32 `json:"pagesize,omitempty"`
}

func (o ListIgnoreRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"ListIgnoreRuleRequest", string(data)}, " ")
}
