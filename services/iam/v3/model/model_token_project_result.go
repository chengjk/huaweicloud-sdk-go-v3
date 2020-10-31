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
type TokenProjectResult struct {
	// 项目名。
	Name string `json:"name"`
	// 项目ID。
	Id     string                    `json:"id"`
	Domain *TokenProjectDomainResult `json:"domain"`
}

func (o TokenProjectResult) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TokenProjectResult", string(data)}, " ")
}
