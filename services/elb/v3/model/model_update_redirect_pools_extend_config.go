package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRedirectPoolsExtendConfig 转发到的后端主机组的配置。当action为REDIRECT_TO_POOL时生效。
type UpdateRedirectPoolsExtendConfig struct {

	// 是否开启url重定向
	RewriteUrlEnable *bool `json:"rewrite_url_enable,omitempty"`

	RewriteUrlConfig *CreateRewriteUrlConfig `json:"rewrite_url_config,omitempty"`
}

func (o UpdateRedirectPoolsExtendConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRedirectPoolsExtendConfig struct{}"
	}

	return strings.Join([]string{"UpdateRedirectPoolsExtendConfig", string(data)}, " ")
}
