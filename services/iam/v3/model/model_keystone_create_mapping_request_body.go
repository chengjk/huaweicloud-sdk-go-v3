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
type KeystoneCreateMappingRequestBody struct {
	Mapping *MappingOption `json:"mapping"`
}

func (o KeystoneCreateMappingRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateMappingRequestBody", string(data)}, " ")
}
