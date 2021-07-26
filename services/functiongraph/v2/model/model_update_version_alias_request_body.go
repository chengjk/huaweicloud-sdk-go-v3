package model

import (
	"encoding/json"

	"strings"
)

type UpdateVersionAliasRequestBody struct {
	// 别名对应的版本名称。

	Version string `json:"version"`
	// 别名描述信息。

	Description *string `json:"description,omitempty"`
	// 灰度版本信息

	AdditionalVersionWeights map[string]int32 `json:"additional_version_weights,omitempty"`
}

func (o UpdateVersionAliasRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVersionAliasRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateVersionAliasRequestBody", string(data)}, " ")
}
