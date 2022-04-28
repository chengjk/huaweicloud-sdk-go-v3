package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOngoingWebinarsResponse struct {

	// 偏移量。
	Offset int32 `json:"offset"`

	// 查询个数。
	Limit int32 `json:"limit"`

	// 总记录数
	Count int64 `json:"count"`

	Data           *[]OpenWebinarOngoingInfo `json:"data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListOngoingWebinarsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOngoingWebinarsResponse struct{}"
	}

	return strings.Join([]string{"ListOngoingWebinarsResponse", string(data)}, " ")
}