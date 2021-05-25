package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationConfigurationRequest struct {
	// 应用ID。

	ApplicationId string `json:"application_id"`
	// 环境ID。

	EnvironmentId string `json:"environment_id"`
}

func (o DeleteApplicationConfigurationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationConfigurationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationConfigurationRequest", string(data)}, " ")
}
