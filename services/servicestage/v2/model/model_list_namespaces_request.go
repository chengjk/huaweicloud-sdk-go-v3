package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListNamespacesRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
}

func (o ListNamespacesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListNamespacesRequest", string(data)}, " ")
}
