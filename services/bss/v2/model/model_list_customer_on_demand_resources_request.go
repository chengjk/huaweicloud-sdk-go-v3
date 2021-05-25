package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCustomerOnDemandResourcesRequest struct {
	// 语言。中文：zh_CN英文：en_US缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`

	Body *QueryCustomerOnDemandResourcesReq `json:"body,omitempty"`
}

func (o ListCustomerOnDemandResourcesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomerOnDemandResourcesRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerOnDemandResourcesRequest", string(data)}, " ")
}
