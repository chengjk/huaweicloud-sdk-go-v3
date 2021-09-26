package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteValueListRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 引用表id

	Valuelistid string `json:"valuelistid"`
}

func (o DeleteValueListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteValueListRequest struct{}"
	}

	return strings.Join([]string{"DeleteValueListRequest", string(data)}, " ")
}
