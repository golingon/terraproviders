// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapublicmaintenanceconfigurations

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Configs struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ConfigsAttributes struct {
	ref terra.Reference
}

func (c ConfigsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigsAttributes) InternalWithRef(ref terra.Reference) ConfigsAttributes {
	return ConfigsAttributes{ref: ref}
}

func (c ConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("description"))
}

func (c ConfigsAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("duration"))
}

func (c ConfigsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("id"))
}

func (c ConfigsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("location"))
}

func (c ConfigsAttributes) MaintenanceScope() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("maintenance_scope"))
}

func (c ConfigsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ConfigsAttributes) RecurEvery() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("recur_every"))
}

func (c ConfigsAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("time_zone"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type ConfigsState struct {
	Description      string `json:"description"`
	Duration         string `json:"duration"`
	Id               string `json:"id"`
	Location         string `json:"location"`
	MaintenanceScope string `json:"maintenance_scope"`
	Name             string `json:"name"`
	RecurEvery       string `json:"recur_every"`
	TimeZone         string `json:"time_zone"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}