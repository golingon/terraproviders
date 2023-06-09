// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataami

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type BlockDeviceMappings struct{}

type ProductCodes struct{}

type Filter struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type BlockDeviceMappingsAttributes struct {
	ref terra.Reference
}

func (bdm BlockDeviceMappingsAttributes) InternalRef() (terra.Reference, error) {
	return bdm.ref, nil
}

func (bdm BlockDeviceMappingsAttributes) InternalWithRef(ref terra.Reference) BlockDeviceMappingsAttributes {
	return BlockDeviceMappingsAttributes{ref: ref}
}

func (bdm BlockDeviceMappingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdm.ref.InternalTokens()
}

func (bdm BlockDeviceMappingsAttributes) DeviceName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("device_name"))
}

func (bdm BlockDeviceMappingsAttributes) Ebs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bdm.ref.Append("ebs"))
}

func (bdm BlockDeviceMappingsAttributes) NoDevice() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("no_device"))
}

func (bdm BlockDeviceMappingsAttributes) VirtualName() terra.StringValue {
	return terra.ReferenceAsString(bdm.ref.Append("virtual_name"))
}

type ProductCodesAttributes struct {
	ref terra.Reference
}

func (pc ProductCodesAttributes) InternalRef() (terra.Reference, error) {
	return pc.ref, nil
}

func (pc ProductCodesAttributes) InternalWithRef(ref terra.Reference) ProductCodesAttributes {
	return ProductCodesAttributes{ref: ref}
}

func (pc ProductCodesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pc.ref.InternalTokens()
}

func (pc ProductCodesAttributes) ProductCodeId() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("product_code_id"))
}

func (pc ProductCodesAttributes) ProductCodeType() terra.StringValue {
	return terra.ReferenceAsString(pc.ref.Append("product_code_type"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FilterAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("values"))
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

type BlockDeviceMappingsState struct {
	DeviceName  string            `json:"device_name"`
	Ebs         map[string]string `json:"ebs"`
	NoDevice    string            `json:"no_device"`
	VirtualName string            `json:"virtual_name"`
}

type ProductCodesState struct {
	ProductCodeId   string `json:"product_code_id"`
	ProductCodeType string `json:"product_code_type"`
}

type FilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
