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
type KeystoneCreateAgencyTokenRequestBody struct {
	Auth *AgencyTokenAuth `json:"auth"`
}

func (o KeystoneCreateAgencyTokenRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateAgencyTokenRequestBody", string(data)}, " ")
}
