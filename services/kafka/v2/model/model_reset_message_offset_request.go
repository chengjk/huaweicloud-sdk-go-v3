package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetMessageOffsetRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 消费组名称。

	Group string `json:"group"`

	Body *ResetMessageOffsetReq `json:"body,omitempty"`
}

func (o ResetMessageOffsetRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetMessageOffsetRequest struct{}"
	}

	return strings.Join([]string{"ResetMessageOffsetRequest", string(data)}, " ")
}
