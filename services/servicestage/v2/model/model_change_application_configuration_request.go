/*
 * ServiceStage
 *
 * ServiceStage的API,包括应用管理和仓库授权管理
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeApplicationConfigurationRequest struct {
	ApplicationId string                   `json:"application_id"`
	Body          *ApplicationConfigModify `json:"body,omitempty"`
}

func (o ChangeApplicationConfigurationRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ChangeApplicationConfigurationRequest", string(data)}, " ")
}
