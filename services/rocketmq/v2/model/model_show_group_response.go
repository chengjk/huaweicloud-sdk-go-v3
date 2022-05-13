package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowGroupResponse struct {

	// 是否启用。
	Enabled *bool `json:"enabled,omitempty"`

	// 是否广播。
	Broadcast *bool `json:"broadcast,omitempty"`

	// 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// 消费组名称。
	Name *string `json:"name,omitempty"`

	// 最大重试次数。
	RetryMaxTime float32 `json:"retry_max_time,omitempty"`

	// 是否重头消费。
	FromBeginning  *bool `json:"from_beginning,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupResponse", string(data)}, " ")
}