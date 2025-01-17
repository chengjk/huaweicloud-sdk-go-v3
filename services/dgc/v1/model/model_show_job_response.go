package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowJobResponse Response Object
type ShowJobResponse struct {

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 节点定义
	Nodes *[]Node `json:"nodes,omitempty"`

	Schedule *Schedule `json:"schedule,omitempty"`

	// 作业参数定义
	Params *[]JobParam `json:"params,omitempty"`

	// 作业在目录树上的路径。创建作业时如果路径目录不存在，会自动创建目录，如/dir/a/，默认在根目录/。
	Directory *string `json:"directory,omitempty"`

	// 作业类型，REAL_TIME： 实时处理，BATCH：批处理
	ProcessType *ShowJobResponseProcessType `json:"processType,omitempty"`

	// 作业最后修改人
	LastUpdateUser *string `json:"lastUpdateUser,omitempty"`

	// 作业运行日志存放的OBS路径。
	LogPath *string `json:"logPath,omitempty"`

	BasicConfig *BasicConfig `json:"basicConfig,omitempty"`

	// 在开启审批开关后，需要填写该字段。表示创建作业的目标状态，有三种状态：SAVED、SUBMITTED和PRODUCTION，分别表示作业创建后是保存态，提交态，生产态。
	TargetStatus *ShowJobResponseTargetStatus `json:"targetStatus,omitempty"`

	// 在开启审批开关后，需要填写该字段，表示作业审批人。
	Approvers      *[]JobApprover `json:"approvers,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

type ShowJobResponseProcessType struct {
	value string
}

type ShowJobResponseProcessTypeEnum struct {
	BATCH     ShowJobResponseProcessType
	REAL_TIME ShowJobResponseProcessType
}

func GetShowJobResponseProcessTypeEnum() ShowJobResponseProcessTypeEnum {
	return ShowJobResponseProcessTypeEnum{
		BATCH: ShowJobResponseProcessType{
			value: "BATCH",
		},
		REAL_TIME: ShowJobResponseProcessType{
			value: "REAL_TIME",
		},
	}
}

func (c ShowJobResponseProcessType) Value() string {
	return c.value
}

func (c ShowJobResponseProcessType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseProcessType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ShowJobResponseTargetStatus struct {
	value string
}

type ShowJobResponseTargetStatusEnum struct {
	SAVED      ShowJobResponseTargetStatus
	SUBMITTED  ShowJobResponseTargetStatus
	PRODUCTION ShowJobResponseTargetStatus
}

func GetShowJobResponseTargetStatusEnum() ShowJobResponseTargetStatusEnum {
	return ShowJobResponseTargetStatusEnum{
		SAVED: ShowJobResponseTargetStatus{
			value: "SAVED",
		},
		SUBMITTED: ShowJobResponseTargetStatus{
			value: "SUBMITTED",
		},
		PRODUCTION: ShowJobResponseTargetStatus{
			value: "PRODUCTION",
		},
	}
}

func (c ShowJobResponseTargetStatus) Value() string {
	return c.value
}

func (c ShowJobResponseTargetStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowJobResponseTargetStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
