package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectTagsRequest struct {
	// 资源类型 目前有: smn_topic，主题 smn_sms，短信 smn_application，移动推送

	ResourceType string `json:"resource_type"`
}

func (o ListProjectTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectTagsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectTagsRequest", string(data)}, " ")
}
