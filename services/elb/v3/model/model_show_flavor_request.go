package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowFlavorRequest struct {
	// 规格ID。

	FlavorId string `json:"flavor_id"`
}

func (o ShowFlavorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowFlavorRequest struct{}"
	}

	return strings.Join([]string{"ShowFlavorRequest", string(data)}, " ")
}
