package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type PublishImageResponse struct {

	// 资产id
	AssetId *string `json:"asset_id,omitempty"`

	// 资产版本
	Version        *string `json:"version,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishImageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishImageResponse struct{}"
	}

	return strings.Join([]string{"PublishImageResponse", string(data)}, " ")
}
