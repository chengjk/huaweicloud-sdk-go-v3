package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreatePublicipTagRequest struct {
	// 资源ID

	PublicipId string `json:"publicip_id"`

	Body *CreatePublicipTagRequestBody `json:"body,omitempty"`
}

func (o CreatePublicipTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreatePublicipTagRequest struct{}"
	}

	return strings.Join([]string{"CreatePublicipTagRequest", string(data)}, " ")
}
