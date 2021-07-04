package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VideoTemplateInfo struct {
	// 画质<br/>

	Quality VideoTemplateInfoQuality `json:"quality"`
	// 视频宽度<br/>

	Width *int32 `json:"width,omitempty"`
	// 视频高度<br/>

	Height *int32 `json:"height,omitempty"`
	// 码率<br/>

	Bitrate *int32 `json:"bitrate,omitempty"`
	// 帧率（默认为1，1代表自适应，单位是帧每秒）<br/>

	FrameRate *int32 `json:"frame_rate,omitempty"`
}

func (o VideoTemplateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoTemplateInfo struct{}"
	}

	return strings.Join([]string{"VideoTemplateInfo", string(data)}, " ")
}

type VideoTemplateInfoQuality struct {
	value string
}

type VideoTemplateInfoQualityEnum struct {
	FULL_HD VideoTemplateInfoQuality
	HD      VideoTemplateInfoQuality
	SD      VideoTemplateInfoQuality
	FLUENT  VideoTemplateInfoQuality
	AD      VideoTemplateInfoQuality
	E_2_K   VideoTemplateInfoQuality
	E_4_K   VideoTemplateInfoQuality
	UNKNOW  VideoTemplateInfoQuality
}

func GetVideoTemplateInfoQualityEnum() VideoTemplateInfoQualityEnum {
	return VideoTemplateInfoQualityEnum{
		FULL_HD: VideoTemplateInfoQuality{
			value: "FULL_HD",
		},
		HD: VideoTemplateInfoQuality{
			value: "HD",
		},
		SD: VideoTemplateInfoQuality{
			value: "SD",
		},
		FLUENT: VideoTemplateInfoQuality{
			value: "FLUENT",
		},
		AD: VideoTemplateInfoQuality{
			value: "AD",
		},
		E_2_K: VideoTemplateInfoQuality{
			value: "2K",
		},
		E_4_K: VideoTemplateInfoQuality{
			value: "4K",
		},
		UNKNOW: VideoTemplateInfoQuality{
			value: "UNKNOW",
		},
	}
}

func (c VideoTemplateInfoQuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoTemplateInfoQuality) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}