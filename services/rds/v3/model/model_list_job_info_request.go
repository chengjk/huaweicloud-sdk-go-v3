package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListJobInfoRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 任务ID。

	Id string `json:"id"`
}

func (o ListJobInfoRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListJobInfoRequest struct{}"
	}

	return strings.Join([]string{"ListJobInfoRequest", string(data)}, " ")
}
