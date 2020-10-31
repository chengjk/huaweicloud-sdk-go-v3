/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

//
type ScopedTokenUser struct {
	Domain       *TokenDomainResult     `json:"domain"`
	OsFederation *TokenUserOsfederation `json:"OS-FEDERATION"`
	// 用户ID。
	Id string `json:"id"`
	// 用户名。
	Name string `json:"name"`
	// 密码过期时间（UTC时间），“”表示密码不过期。
	PasswordExpiresAt string `json:"password_expires_at"`
}

func (o ScopedTokenUser) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ScopedTokenUser", string(data)}, " ")
}
