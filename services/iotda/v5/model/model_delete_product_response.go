package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteProductResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteProductResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteProductResponse struct{}"
	}

	return strings.Join([]string{"DeleteProductResponse", string(data)}, " ")
}