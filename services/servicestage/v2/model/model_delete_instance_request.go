package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstanceRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 组件ID。

	ComponentId string `json:"component_id"`
	// 组件实例ID。

	InstanceId string `json:"instance_id"`
	// 是否强制删除。

	Force *bool `json:"force,omitempty"`
}

func (o DeleteInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceRequest", string(data)}, " ")
}
