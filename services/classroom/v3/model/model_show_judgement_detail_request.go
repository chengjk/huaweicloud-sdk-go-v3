package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowJudgementDetailRequest struct {
	// 判题任务ID

	JudgementId string `json:"judgement_id"`
}

func (o ShowJudgementDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowJudgementDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowJudgementDetailRequest", string(data)}, " ")
}