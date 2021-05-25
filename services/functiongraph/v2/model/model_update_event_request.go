package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEventRequest struct {
	// 事件ID。

	EventId string `json:"event_id"`
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`

	Body *UpdateEventRequestBody `json:"body,omitempty"`
}

func (o UpdateEventRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEventRequest struct{}"
	}

	return strings.Join([]string{"UpdateEventRequest", string(data)}, " ")
}
