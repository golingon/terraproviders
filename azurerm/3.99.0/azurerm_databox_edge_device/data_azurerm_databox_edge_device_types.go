// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_databox_edge_device

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataDevicePropertiesAttributes struct {
	ref terra.Reference
}

func (dp DataDevicePropertiesAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp DataDevicePropertiesAttributes) InternalWithRef(ref terra.Reference) DataDevicePropertiesAttributes {
	return DataDevicePropertiesAttributes{ref: ref}
}

func (dp DataDevicePropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp DataDevicePropertiesAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("capacity"))
}

func (dp DataDevicePropertiesAttributes) ConfiguredRoleTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dp.ref.Append("configured_role_types"))
}

func (dp DataDevicePropertiesAttributes) Culture() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("culture"))
}

func (dp DataDevicePropertiesAttributes) HcsVersion() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("hcs_version"))
}

func (dp DataDevicePropertiesAttributes) Model() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("model"))
}

func (dp DataDevicePropertiesAttributes) NodeCount() terra.NumberValue {
	return terra.ReferenceAsNumber(dp.ref.Append("node_count"))
}

func (dp DataDevicePropertiesAttributes) SerialNumber() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("serial_number"))
}

func (dp DataDevicePropertiesAttributes) SoftwareVersion() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("software_version"))
}

func (dp DataDevicePropertiesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("status"))
}

func (dp DataDevicePropertiesAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("time_zone"))
}

func (dp DataDevicePropertiesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("type"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataDevicePropertiesState struct {
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

type DataTimeoutsState struct {
	Read string `json:"read"`
}
