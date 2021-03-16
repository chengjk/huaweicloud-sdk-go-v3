package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateOpenIdConnectConfigRequest struct {
	IdpId string                                `json:"idp_id"`
	Body  *UpdateOpenIdConnectConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateOpenIdConnectConfigRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateOpenIdConnectConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateOpenIdConnectConfigRequest", string(data)}, " ")
}