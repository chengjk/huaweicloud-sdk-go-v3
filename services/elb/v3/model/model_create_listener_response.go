/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateListenerResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId *string   `json:"request_id,omitempty"`
	Listener  *Listener `json:"listener,omitempty"`
}

func (o CreateListenerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateListenerResponse", string(data)}, " ")
}
