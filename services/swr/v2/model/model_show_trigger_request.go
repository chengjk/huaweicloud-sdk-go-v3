package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowTriggerRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ShowTriggerRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
	// 触发器名称

	Trigger string `json:"trigger"`
}

func (o ShowTriggerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTriggerRequest struct{}"
	}

	return strings.Join([]string{"ShowTriggerRequest", string(data)}, " ")
}

type ShowTriggerRequestContentType struct {
	value string
}

type ShowTriggerRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowTriggerRequestContentType
	APPLICATION_JSON             ShowTriggerRequestContentType
}

func GetShowTriggerRequestContentTypeEnum() ShowTriggerRequestContentTypeEnum {
	return ShowTriggerRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowTriggerRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowTriggerRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowTriggerRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ShowTriggerRequestContentType) UnmarshalJSON(b []byte) error {
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
