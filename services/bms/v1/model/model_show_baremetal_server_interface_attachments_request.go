package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBaremetalServerInterfaceAttachmentsRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`
}

func (o ShowBaremetalServerInterfaceAttachmentsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerInterfaceAttachmentsRequest struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerInterfaceAttachmentsRequest", string(data)}, " ")
}
