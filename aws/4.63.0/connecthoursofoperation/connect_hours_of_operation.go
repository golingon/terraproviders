// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package connecthoursofoperation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Config struct {
	// Day: string, required
	Day terra.StringValue `hcl:"day,attr" validate:"required"`
	// EndTime: required
	EndTime *EndTime `hcl:"end_time,block" validate:"required"`
	// StartTime: required
	StartTime *StartTime `hcl:"start_time,block" validate:"required"`
}

type EndTime struct {
	// Hours: number, required
	Hours terra.NumberValue `hcl:"hours,attr" validate:"required"`
	// Minutes: number, required
	Minutes terra.NumberValue `hcl:"minutes,attr" validate:"required"`
}

type StartTime struct {
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

func (c ConfigAttributes) EndTime() terra.ListValue[EndTimeAttributes] {
	return terra.ReferenceAsList[EndTimeAttributes](c.ref.Append("end_time"))
}

func (c ConfigAttributes) StartTime() terra.ListValue[StartTimeAttributes] {
	return terra.ReferenceAsList[StartTimeAttributes](c.ref.Append("start_time"))
}

type EndTimeAttributes struct {
	ref terra.Reference
}

func (et EndTimeAttributes) InternalRef() (terra.Reference, error) {
	return et.ref, nil
}

func (et EndTimeAttributes) InternalWithRef(ref terra.Reference) EndTimeAttributes {
	return EndTimeAttributes{ref: ref}
}

func (et EndTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return et.ref.InternalTokens()
}

func (et EndTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(et.ref.Append("hours"))
}

func (et EndTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(et.ref.Append("minutes"))
}

type StartTimeAttributes struct {
	ref terra.Reference
}

func (st StartTimeAttributes) InternalRef() (terra.Reference, error) {
	return st.ref, nil
}

func (st StartTimeAttributes) InternalWithRef(ref terra.Reference) StartTimeAttributes {
	return StartTimeAttributes{ref: ref}
}

func (st StartTimeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return st.ref.InternalTokens()
}

func (st StartTimeAttributes) Hours() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("hours"))
}

func (st StartTimeAttributes) Minutes() terra.NumberValue {
	return terra.ReferenceAsNumber(st.ref.Append("minutes"))
}

type ConfigState struct {
	Day       string           `json:"day"`
	EndTime   []EndTimeState   `json:"end_time"`
	StartTime []StartTimeState `json:"start_time"`
}

type EndTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
}

type StartTimeState struct {
	Hours   float64 `json:"hours"`
	Minutes float64 `json:"minutes"`
}
