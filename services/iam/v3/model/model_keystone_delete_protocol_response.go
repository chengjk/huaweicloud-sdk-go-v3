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
type KeystoneDeleteProtocolResponse struct {
}

func (o KeystoneDeleteProtocolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneDeleteProtocolResponse", string(data)}, " ")
}
