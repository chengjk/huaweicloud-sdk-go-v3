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
type KeystoneCreateProtocolResponse struct {
	Protocol *ProtocolResult `json:"protocol,omitempty"`
}

func (o KeystoneCreateProtocolResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"KeystoneCreateProtocolResponse", string(data)}, " ")
}
