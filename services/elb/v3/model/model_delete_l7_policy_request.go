package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteL7PolicyRequest struct {
	// 转发策略ID。

	L7policyId string `json:"l7policy_id"`
}

func (o DeleteL7PolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7PolicyRequest", string(data)}, " ")
}
