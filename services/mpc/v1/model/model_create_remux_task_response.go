package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateRemuxTaskResponse struct {
	// 任务ID

	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRemuxTaskResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRemuxTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateRemuxTaskResponse", string(data)}, " ")
}
