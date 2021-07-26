package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPremiumHostRequest struct {
	// 独享模式域名ID

	HostId string `json:"host_id"`
}

func (o ShowPremiumHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPremiumHostRequest struct{}"
	}

	return strings.Join([]string{"ShowPremiumHostRequest", string(data)}, " ")
}
