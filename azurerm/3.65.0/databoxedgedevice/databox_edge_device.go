// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package databoxedgedevice

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DeviceProperties struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DevicePropertiesAttributes struct {
	ref terra.Reference
}

func (dp DevicePropertiesAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DevicePropertiesAttributes) InternalWithRef(ref terra.Reference) DevicePropertiesAttributes {
	return DevicePropertiesAttributes{ref: ref}
}

func (dp DevicePropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DevicePropertiesAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("capacity"))
}

func (dp DevicePropertiesAttributes) ConfiguredRoleTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dp.ref.Append("configured_role_types"))
}

func (dp DevicePropertiesAttributes) Culture() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("culture"))
}

func (dp DevicePropertiesAttributes) HcsVersion() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("hcs_version"))
}

func (dp DevicePropertiesAttributes) Model() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("model"))
}

func (dp DevicePropertiesAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("node_count"))
}

func (dp DevicePropertiesAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("serial_number"))
}

func (dp DevicePropertiesAttributes) SoftwareVersion() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("software_version"))
}

func (dp DevicePropertiesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("status"))
}

func (dp DevicePropertiesAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("time_zone"))
}

func (dp DevicePropertiesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("type"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type DevicePropertiesState struct {
	Capacity            float64  `json:"capacity"`
	ConfiguredRoleTypes []string `json:"configured_role_types"`
	Culture             string   `json:"culture"`
	HcsVersion          string   `json:"hcs_version"`
	Model               string   `json:"model"`
	NodeCount           float64  `json:"node_count"`
	SerialNumber        string   `json:"serial_number"`
	SoftwareVersion     string   `json:"software_version"`
	Status              string   `json:"status"`
	TimeZone            string   `json:"time_zone"`
	Type                string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
