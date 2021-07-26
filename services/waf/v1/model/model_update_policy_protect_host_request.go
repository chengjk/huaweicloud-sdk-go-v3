package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePolicyProtectHostRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// 域名id9从获取防护网站列表获取域名id）

	Hosts string `json:"hosts"`
}

func (o UpdatePolicyProtectHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePolicyProtectHostRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyProtectHostRequest", string(data)}, " ")
}
