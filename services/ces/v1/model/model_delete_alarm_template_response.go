/*
 * CES
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
type DeleteAlarmTemplateResponse struct {
}

func (o DeleteAlarmTemplateResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeleteAlarmTemplateResponse", string(data)}, " ")
}
