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

// Response Object
type KeystoneCreateUserTokenByPasswordResponse struct {
	Token         *TokenResult `json:"token,omitempty"`
	XSubjectToken *string      `json:"X-Subject-Token,omitempty"`
}

func (o KeystoneCreateUserTokenByPasswordResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateUserTokenByPasswordResponse", string(data)}, " ")
}
