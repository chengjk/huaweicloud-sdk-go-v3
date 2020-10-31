/*
 * ELB
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowLoadBalancerResponse struct {
	// 请求ID。  注：自动生成 。
	RequestId    *string       `json:"request_id,omitempty"`
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`
}

func (o ShowLoadBalancerResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowLoadBalancerResponse", string(data)}, " ")
}
