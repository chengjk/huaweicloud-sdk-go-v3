package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateSinkTaskRequest struct {
	// 实例转储ID。 请参考[实例生命周期][查询实例]接口返回的数据。

	ConnectorId string `json:"connector_id"`

	Body *CreateSinkTaskReq `json:"body,omitempty"`
}

func (o CreateSinkTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateSinkTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateSinkTaskRequest", string(data)}, " ")
}
