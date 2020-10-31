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
type IdentityprovidersLinks struct {
	// 身份提供商的资源链接地址。
	Self string `json:"self"`
	// 协议的资源链接地址。
	Protocols string `json:"protocols"`
}

func (o IdentityprovidersLinks) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"IdentityprovidersLinks", string(data)}, " ")
}
