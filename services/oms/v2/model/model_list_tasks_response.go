package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListTasksResponse struct {
	// 查询的任务详情

	Tasks *[]TaskResp `json:"tasks,omitempty"`
	// 满足查询条件的任务总数

	Count          *int64 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTasksResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTasksResponse struct{}"
	}

	return strings.Join([]string{"ListTasksResponse", string(data)}, " ")
}
