package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBranchesRequest struct {
	// 授权名称。

	XRepoAuth string `json:"X-Repo-Auth"`
	// 组织ID。

	Namespace string `json:"namespace"`
	// 仓库项目ID，如果含有“/”，需要将“/”替换为“:”。

	Project string `json:"project"`
}

func (o ListBranchesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBranchesRequest struct{}"
	}

	return strings.Join([]string{"ListBranchesRequest", string(data)}, " ")
}
