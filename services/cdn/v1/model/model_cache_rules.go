package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CacheRules struct {

	// 匹配所有文件 all 按文件后缀匹配 file_extension 按目录匹配 catalog 全路径匹配 full_path URL匹配正则表达式 regex 按首页匹配 home_page
	MatchType string `json:"match_type"`

	// 缓存匹配设置。 当match_type为off时，为空。当match_type为on时，为文件后缀，输入首字符为“.”，以“;”进行分隔，如.jpg;.zip;.exe，并且输入的文件名后缀总数不超过20个。 当rule_type为2时，为目录，输入要求以“/”作为首字符，以“;”进行分隔，如/test/folder01;/test/folder02，并且输入的目录路径总数不超过20个。    当rule_type为3时，为全路径，输入要求以“/”作为首字符，支持匹配指定目录下的具体文件，或者带通配符“*”的文件，如/test/index.html或/test/_*.jpg
	MatchValue *string `json:"match_value,omitempty"`

	// 缓存时间。最大支持365天。
	Ttl int32 `json:"ttl"`

	// 缓存时间单位。1：秒；2：分；3：小时；4：天
	TtlUnit string `json:"ttl_unit"`

	// 此条配置的权重值, 默认值1，数值越大，优先级越高。取值范围为1-100，权重值不能相同
	Priority int32 `json:"priority"`

	// 缓存遵循源站开关（on：打开，off：关闭）
	FollowOrigin string `json:"follow_origin"`

	// 忽略指定URL参数： del_params 保留指定URL参数： reserve_params 忽略全部URL参数： ignore_url_params 使用完整： URL full_url
	UrlParameterType string `json:"url_parameter_type"`

	// 最多设置10条，以\",\"分隔
	UrlParameterValue *string `json:"url_parameter_value,omitempty"`
}

func (o CacheRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CacheRules struct{}"
	}

	return strings.Join([]string{"CacheRules", string(data)}, " ")
}