package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateFunctionCodeRequestBody struct {
	// 函数执行入口 规则：xx.xx，必须包含“. ” 举例：对于node.js函数：myfunction.handler，则表示函数的文件名为myfunction.js，执行的入口函数名为handler。

	Handler string `json:"handler"`
	// 函数代码类型，取值有4种。 inline: UI在线编辑代码。 zip: 函数代码为zip包。 obs: 函数代码来源于obs存储。 jar: 函数代码为jar包，主要针对Java函数。

	CodeType UpdateFunctionCodeRequestBodyCodeType `json:"code_type"`
	// 当CodeType为obs时，该值为函数代码包在OBS上的地址，CodeType为其他值时，该字段为空。

	CodeUrl *string `json:"code_url,omitempty"`
	// 函数的文件名，当CodeType为jar/zip时必须提供该字段，inline和obs不需要提供。

	CodeFilename *string `json:"code_filename,omitempty"`

	FuncCode *FuncCode `json:"func_code"`
	// 依赖id列表

	DependList *[]string `json:"depend_list,omitempty"`
	// 函数依赖代码包列表。

	Dependencies *[]Dependency `json:"dependencies,omitempty"`
}

func (o UpdateFunctionCodeRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateFunctionCodeRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFunctionCodeRequestBody", string(data)}, " ")
}

type UpdateFunctionCodeRequestBodyCodeType struct {
	value string
}

type UpdateFunctionCodeRequestBodyCodeTypeEnum struct {
	INLINE UpdateFunctionCodeRequestBodyCodeType
	ZIP    UpdateFunctionCodeRequestBodyCodeType
	OBS    UpdateFunctionCodeRequestBodyCodeType
	JAR    UpdateFunctionCodeRequestBodyCodeType
}

func GetUpdateFunctionCodeRequestBodyCodeTypeEnum() UpdateFunctionCodeRequestBodyCodeTypeEnum {
	return UpdateFunctionCodeRequestBodyCodeTypeEnum{
		INLINE: UpdateFunctionCodeRequestBodyCodeType{
			value: "inline",
		},
		ZIP: UpdateFunctionCodeRequestBodyCodeType{
			value: "zip",
		},
		OBS: UpdateFunctionCodeRequestBodyCodeType{
			value: "obs",
		},
		JAR: UpdateFunctionCodeRequestBodyCodeType{
			value: "jar",
		},
	}
}

func (c UpdateFunctionCodeRequestBodyCodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateFunctionCodeRequestBodyCodeType) UnmarshalJSON(b []byte) error {
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
