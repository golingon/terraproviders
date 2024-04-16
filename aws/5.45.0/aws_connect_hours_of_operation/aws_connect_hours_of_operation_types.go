// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_connect_hours_of_operation

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	// Day: string, required
	Day terra.StringValue `hcl:"day,attr" validate:"required"`
	// ConfigEndTime: required
	EndTime *ConfigEndTime `hcl:"end_time,block" validate:"required"`
	// ConfigStartTime: required
	StartTime *ConfigStartTime `hcl:"start_time,block" validate:"required"`
}

type ConfigEndTime struct {
	// Hours: number, required
	Hours terra.NumberValue `hcl:"hours,attr" validate:"required"`
	// Minutes: number, required
	Minutes terra.NumberValue `hcl:"minutes,attr" validate:"required"`
}

type ConfigStartTime struct {
	// Hours: number, required
	Hours terra.NumberValue `hcl:"hours,attr" validate:"required"`
	// Minutes: number, required
	Minutes terra.NumberValue `hcl:"minutes,attr" validate:"required"`
}

type ConfigAttributes struct {
	ref terra.Reference
}

func (c ConfigAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigAttributes) InternalWithRef(ref terra.Reference) ConfigAttributes {
	return ConfigAttributes{ref: ref}
}

func (c ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigAttributes) Day() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("day"))
}

func (c ConfigAttributes) EndTime() terra.ListValue[ConfigEndTimeAttributes] {
	return terra.ReferenceAsList[ConfigEndTimeAttributes](c.ref.Append("end_time"))
}

func (c ConfigAttributes) StartTime() terra.ListValue[ConfigStartTimeAttributes] {
	return terra.ReferenceAsList[ConfigStartTimeAttributes](c.ref.Append("start_time"))
}

type ConfigEndTimeAttributes struct {
	ref terra.Reference
}

func (et ConfigEndTimeAttributes) InternalRef() (terra.Reference, error) {
	return et.ref, nil
}

func (et ConfigEndTimeAttributes) InternalWithRef(ref terra.Reference) ConfigEndTimeAttributes {
	return ConfigEndTimeAttributes{ref: ref}
}

func (et ConfigEndTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return et.ref.InternalTokens()
}

func (et ConfigEndTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(et.ref.Append("hours"))
}

func (et ConfigEndTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(et.ref.Append("minutes"))
}

type ConfigStartTimeAttributes struct {
	ref terra.Reference
}

func (st ConfigStartTimeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st ConfigStartTimeAttributes) InternalWithRef(ref terra.Reference) ConfigStartTimeAttributes {
	return ConfigStartTimeAttributes{ref: ref}
}

func (st ConfigStartTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st ConfigStartTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("hours"))
}

func (st ConfigStartTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("minutes"))
}

type ConfigState struct {
	Day       string                 `json:"day"`
	EndTime   []ConfigEndTimeState   `json:"end_time"`
	StartTime []ConfigStartTimeState `json:"start_time"`
}

type ConfigEndTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
}

type ConfigStartTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
}
