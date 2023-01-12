package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type HandwritingRequestBody struct {

	// 与url二选一  图像数据，base64编码，要求base64编码后大小不超过10MB。图片最小边不小于8px，最长边不超过8192px，支持JPEG、JPG、PNG、BMP、TIFF格式。  图片文件Base64编码字符串，点击[这里](https://support.huaweicloud.com/ocr_faq/ocr_01_0032.html)查看详细获取方式。
	Image *string `json:"image,omitempty"`

	// 与image二选一 图片的URL路径，目前支持： - 公网http/https url - OBS提供的url，使用OBS数据需要进行授权。包括对服务授权、临时授权、匿名公开授权，详情参见[配置OBS访问权限](https://support.huaweicloud.com/api-ocr/ocr_03_0132.html)。 > 说明： - 接口响应时间依赖于图片的下载时间，如果图片下载时间过长，会返回接口调用失败。 - 请保证被检测图片所在的存储服务稳定可靠，推荐使用OBS服务存储图片数据。
	Url *string `json:"url,omitempty"`

	// 快速模式开关，针对单行文字图片（要求图片只包含一行文字，且文字区域占比超过50%），打开时可以更快返回识别内容。可选值包括： - true：打开快速模式； - false：关闭快速模式。  > 说明： - 未传入该参数时默认为false，即关闭快速模式
	QuickMode *bool `json:"quick_mode,omitempty"`

	// 字符集设置，用户可以根据实际需要限定输出字符集范围。可选值如下所示。  - \"digit\": 数字模式；  - \"letter\": 大小写字母模式；  - \"digit_letter\": 数字+字母模式；  - \"general\": 数字+字母+中文模式；  > 说明： - 未传入该参数时，默认为“general”模式。
	CharSet *string `json:"char_set,omitempty"`

	// 校正图片的倾斜角度开关，可选值如下所示。 - true：校正图片的倾斜角度； - false：不校正图片的倾斜角度。  > 说明： - 支持任意角度的校正，未传入该参数时默认为“false”。
	DetectDirection *bool `json:"detect_direction,omitempty"`
}

func (o HandwritingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HandwritingRequestBody struct{}"
	}

	return strings.Join([]string{"HandwritingRequestBody", string(data)}, " ")
}
