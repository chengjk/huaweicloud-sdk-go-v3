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
type KeystoneUpdateIdentityProviderRequestBody struct {
	IdentityProvider *IdentityproviderOption `json:"identity_provider"`
}

func (o KeystoneUpdateIdentityProviderRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneUpdateIdentityProviderRequestBody", string(data)}, " ")
}
