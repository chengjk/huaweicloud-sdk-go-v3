package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListImageAutoSyncReposDetailsRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType ListImageAutoSyncReposDetailsRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`
}

func (o ListImageAutoSyncReposDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListImageAutoSyncReposDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListImageAutoSyncReposDetailsRequest", string(data)}, " ")
}

type ListImageAutoSyncReposDetailsRequestContentType struct {
	value string
}

type ListImageAutoSyncReposDetailsRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ListImageAutoSyncReposDetailsRequestContentType
	APPLICATION_JSON             ListImageAutoSyncReposDetailsRequestContentType
}

func GetListImageAutoSyncReposDetailsRequestContentTypeEnum() ListImageAutoSyncReposDetailsRequestContentTypeEnum {
	return ListImageAutoSyncReposDetailsRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ListImageAutoSyncReposDetailsRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ListImageAutoSyncReposDetailsRequestContentType{
			value: "application/json",
		},
	}
}

func (c ListImageAutoSyncReposDetailsRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListImageAutoSyncReposDetailsRequestContentType) UnmarshalJSON(b []byte) error {
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
