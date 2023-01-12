package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MonitorMethod struct {
	value string
}

type MonitorMethodEnum struct {
	MAX     MonitorMethod
	MIN     MonitorMethod
	AVERAGE MonitorMethod
}

func GetMonitorMethodEnum() MonitorMethodEnum {
	return MonitorMethodEnum{
		MAX: MonitorMethod{
			value: "max",
		},
		MIN: MonitorMethod{
			value: "min",
		},
		AVERAGE: MonitorMethod{
			value: "average",
		},
	}
}

func (c MonitorMethod) Value() string {
	return c.value
}

func (c MonitorMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorMethod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
