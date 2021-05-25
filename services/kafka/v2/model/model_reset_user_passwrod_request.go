package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetUserPasswrodRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 用户名称。

	UserName string `json:"user_name"`

	Body *ResetUserPasswrodReq `json:"body,omitempty"`
}

func (o ResetUserPasswrodRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetUserPasswrodRequest struct{}"
	}

	return strings.Join([]string{"ResetUserPasswrodRequest", string(data)}, " ")
}
