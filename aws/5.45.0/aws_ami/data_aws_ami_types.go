// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ami

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFilter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: set of string, required
	Values terra.SetValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataBlockDeviceMappingsAttributes struct {
	ref terra.Reference
}

func (bdm DataBlockDeviceMappingsAttributes) InternalRef() (terra.Reference, error) {
	return bdm.ref, nil
}

func (bdm DataBlockDeviceMappingsAttributes) InternalWithRef(ref terra.Reference) DataBlockDeviceMappingsAttributes {
	return DataBlockDeviceMappingsAttributes{ref: ref}
}

func (bdm DataBlockDeviceMappingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdm.ref.InternalTokens()
}

func (bdm DataBlockDeviceMappingsAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("device_name"))
}

func (bdm DataBlockDeviceMappingsAttributes) Ebs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bdm.ref.Append("ebs"))
}

func (bdm DataBlockDeviceMappingsAttributes) NoDevice() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("no_device"))
}

func (bdm DataBlockDeviceMappingsAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("virtual_name"))
}

type DataProductCodesAttributes struct {
	ref terra.Reference
}

func (pc DataProductCodesAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc DataProductCodesAttributes) InternalWithRef(ref terra.Reference) DataProductCodesAttributes {
	return DataProductCodesAttributes{ref: ref}
}

func (pc DataProductCodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc DataProductCodesAttributes) ProductCodeId() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("product_code_id"))
}

func (pc DataProductCodesAttributes) ProductCodeType() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("product_code_type"))
}

type DataFilterAttributes struct {
	ref terra.Reference
}

func (f DataFilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFilterAttributes) InternalWithRef(ref terra.Reference) DataFilterAttributes {
	return DataFilterAttributes{ref: ref}
}

func (f DataFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f DataFilterAttributes) Values() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("values"))
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

type DataBlockDeviceMappingsState struct {
	DeviceName  string            `json:"device_name"`
	Ebs         map[string]string `json:"ebs"`
	NoDevice    string            `json:"no_device"`
	VirtualName string            `json:"virtual_name"`
}

type DataProductCodesState struct {
	ProductCodeId   string `json:"product_code_id"`
	ProductCodeType string `json:"product_code_type"`
}

type DataFilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
