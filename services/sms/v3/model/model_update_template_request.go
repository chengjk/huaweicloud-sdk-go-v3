package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateTemplateRequest struct {
	// 需要修改信息的模板的id

	Id string `json:"id"`

	Body *UpdateTemplateReq `json:"body,omitempty"`
}

func (o UpdateTemplateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateTemplateRequest", string(data)}, " ")
}