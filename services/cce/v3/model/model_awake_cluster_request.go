package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AwakeClusterRequest struct {
	// 集群 ID，获取方式请参见[[如何获取接口URI中参数](https://support.huaweicloud.com/api-cce/cce_02_0271.html)](tag:hws)[[如何获取接口URI中参数](https://support.huaweicloud.com/intl/zh-cn/api-cce/cce_02_0271.html)](tag:hws_hk)

	ClusterId string `json:"cluster_id"`
}

func (o AwakeClusterRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AwakeClusterRequest struct{}"
	}

	return strings.Join([]string{"AwakeClusterRequest", string(data)}, " ")
}
