package model

import (
	"encoding/json"

	"strings"
)

// 标签类别的多种语言输出。
type ImageTaggingItemBodyI18nType struct {
	// 中文标签类别

	Zh *string `json:"zh,omitempty"`
	// 英文标签类别

	En *string `json:"en,omitempty"`
}

func (o ImageTaggingItemBodyI18nType) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImageTaggingItemBodyI18nType struct{}"
	}

	return strings.Join([]string{"ImageTaggingItemBodyI18nType", string(data)}, " ")
}