package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteDbUserRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 需要删除的帐号名。

	UserName string `json:"user_name"`
}

func (o DeleteDbUserRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteDbUserRequest struct{}"
	}

	return strings.Join([]string{"DeleteDbUserRequest", string(data)}, " ")
}
