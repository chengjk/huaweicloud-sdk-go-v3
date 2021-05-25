package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateAlarmRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 告警规则的ID。

	AlarmId string `json:"alarm_id"`

	Body *UpdateAlarmRequestBody `json:"body,omitempty"`
}

func (o UpdateAlarmRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAlarmRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmRequest", string(data)}, " ")
}
