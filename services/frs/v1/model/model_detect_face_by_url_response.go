package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DetectFaceByUrlResponse struct {
	// 检测到的人脸。 调用失败时无此字段。

	Faces          *[]DetectFace `json:"faces,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o DetectFaceByUrlResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DetectFaceByUrlResponse struct{}"
	}

	return strings.Join([]string{"DetectFaceByUrlResponse", string(data)}, " ")
}