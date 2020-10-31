/*
 * DevStar
 *
 * DevStar API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type TagInfo struct {
	// 自定义标签id
	Id *string `json:"id,omitempty"`
	// 自定义标签名称
	Name *string `json:"name,omitempty"`
}

func (o TagInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TagInfo", string(data)}, " ")
}
