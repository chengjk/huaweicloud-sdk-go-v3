package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRouteTableRequest struct {
	// 路由表ID

	RoutetableId string `json:"routetable_id"`

	Body *UpdateRoutetableReqBody `json:"body,omitempty"`
}

func (o UpdateRouteTableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRouteTableRequest struct{}"
	}

	return strings.Join([]string{"UpdateRouteTableRequest", string(data)}, " ")
}