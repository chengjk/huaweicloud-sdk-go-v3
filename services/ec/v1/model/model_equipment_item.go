package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type EquipmentItem struct {

	// 智能企业网关设备ID
	Id string `json:"id"`

	// 智能企业网关ID
	IegId *string `json:"ieg_id,omitempty"`

	// esn
	Esn *string `json:"esn,omitempty"`

	// 设备名字
	Name *string `json:"name,omitempty"`

	// 设备类型
	Type *EquipmentItemType `json:"type,omitempty"`

	// 高可用类型
	HaType *EquipmentItemHaType `json:"ha_type,omitempty"`

	// 设备软件版本
	Version *string `json:"version,omitempty"`

	// 激活时间
	ActiveAt *sdktime.SdkTime `json:"active_at,omitempty"`

	// 上线时间
	GoLiveAt *sdktime.SdkTime `json:"go_live_at,omitempty"`

	// 设备启动时间
	StartUpAt *sdktime.SdkTime `json:"start_up_at,omitempty"`

	// VPN状态
	CloudAccessStatus *string `json:"cloud_access_status,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`
}

func (o EquipmentItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EquipmentItem struct{}"
	}

	return strings.Join([]string{"EquipmentItem", string(data)}, " ")
}

type EquipmentItemType struct {
	value string
}

type EquipmentItemTypeEnum struct {
	STANDARD EquipmentItemType
}

func GetEquipmentItemTypeEnum() EquipmentItemTypeEnum {
	return EquipmentItemTypeEnum{
		STANDARD: EquipmentItemType{
			value: "standard",
		},
	}
}

func (c EquipmentItemType) Value() string {
	return c.value
}

func (c EquipmentItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentItemType) UnmarshalJSON(b []byte) error {
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

type EquipmentItemHaType struct {
	value string
}

type EquipmentItemHaTypeEnum struct {
	ACTIVE  EquipmentItemHaType
	STANDBY EquipmentItemHaType
}

func GetEquipmentItemHaTypeEnum() EquipmentItemHaTypeEnum {
	return EquipmentItemHaTypeEnum{
		ACTIVE: EquipmentItemHaType{
			value: "Active",
		},
		STANDBY: EquipmentItemHaType{
			value: "Standby",
		},
	}
}

func (c EquipmentItemHaType) Value() string {
	return c.value
}

func (c EquipmentItemHaType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EquipmentItemHaType) UnmarshalJSON(b []byte) error {
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
