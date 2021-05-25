package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteL7policyRequest struct {
	// 转发策略id

	L7policyId string `json:"l7policy_id"`
}

func (o DeleteL7policyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7policyRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7policyRequest", string(data)}, " ")
}
