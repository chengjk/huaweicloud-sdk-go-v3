package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTagRequest struct {
	// 密钥ID

	KeyId string `json:"key_id"`
	// 标签键的值

	Key string `json:"key"`
	// API版本号

	VersionId string `json:"version_id"`
}

func (o DeleteTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteTagRequest", string(data)}, " ")
}
