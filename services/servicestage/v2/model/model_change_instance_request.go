package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeInstanceRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`

	Body *InstanceModify `json:"body,omitempty"`
}

func (o ChangeInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeInstanceRequest struct{}"
	}

	return strings.Join([]string{"ChangeInstanceRequest", string(data)}, " ")
}
