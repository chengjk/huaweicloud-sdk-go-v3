package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSecretForScheduleRequest struct {
	// 凭据的资源标识符。

	SecretId string `json:"secret_id"`

	Body *DeleteSecretForScheduleRequestBody `json:"body,omitempty"`
}

func (o DeleteSecretForScheduleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSecretForScheduleRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecretForScheduleRequest", string(data)}, " ")
}
