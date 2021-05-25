package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreJobInstanceRequest struct {
	// 作业名称.

	JobName string `json:"job_name"`
	// 作业实例id.

	InstanceId string `json:"instance_id"`
}

func (o RestoreJobInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreJobInstanceRequest struct{}"
	}

	return strings.Join([]string{"RestoreJobInstanceRequest", string(data)}, " ")
}
