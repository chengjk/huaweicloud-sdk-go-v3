package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIndicatorDetailResponse struct {

	// Id value
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *IndicatorDetail `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowIndicatorDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIndicatorDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowIndicatorDetailResponse", string(data)}, " ")
}
