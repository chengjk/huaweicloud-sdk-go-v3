package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NeutronRemoveFirewallRuleRequest struct {
	// 网络ACL防火墙策略ID

	FirewallPolicyId string `json:"firewall_policy_id"`

	Body *NeutronRemoveFirewallRuleRequestBody `json:"body,omitempty"`
}

func (o NeutronRemoveFirewallRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NeutronRemoveFirewallRuleRequest struct{}"
	}

	return strings.Join([]string{"NeutronRemoveFirewallRuleRequest", string(data)}, " ")
}
