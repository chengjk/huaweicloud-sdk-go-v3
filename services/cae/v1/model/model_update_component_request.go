package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateComponentRequest struct {

	// 组件id。
	ComponentId string `json:"component_id"`

	// 环境id。
	XEnvironmentID string `json:"X-Environment-ID"`

	// 租户的企业项目id。
	XEnterpriseProjectID *string `json:"X-Enterprise-Project-ID,omitempty"`

	// 应用id。
	ApplicationId string `json:"application_id"`

	Body *UpdateComponentRequestBody `json:"body,omitempty"`
}

func (o UpdateComponentRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComponentRequest struct{}"
	}

	return strings.Join([]string{"UpdateComponentRequest", string(data)}, " ")
}