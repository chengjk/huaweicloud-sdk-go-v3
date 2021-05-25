package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateIpGroupRequest struct {
	// 待更新的IP地址组的id

	IpgroupId string `json:"ipgroup_id"`

	Body *UpdateIpGroupRequestBody `json:"body,omitempty"`
}

func (o UpdateIpGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateIpGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupRequest", string(data)}, " ")
}
