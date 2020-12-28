/*
 * BCS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 被邀请列表
type InvitedDomain struct {
	// 被邀请方租户名
	InvitedUsername *string `json:"invited_username,omitempty"`
}

func (o InvitedDomain) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"InvitedDomain", string(data)}, " ")
}