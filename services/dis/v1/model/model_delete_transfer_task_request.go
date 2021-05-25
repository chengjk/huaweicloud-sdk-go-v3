package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTransferTaskRequest struct {
	// 已创建的通道的名称。

	StreamName string `json:"stream_name"`
	// 待删除的转储任务名称。

	TaskName string `json:"task_name"`
}

func (o DeleteTransferTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTransferTaskRequest struct{}"
	}

	return strings.Join([]string{"DeleteTransferTaskRequest", string(data)}, " ")
}
