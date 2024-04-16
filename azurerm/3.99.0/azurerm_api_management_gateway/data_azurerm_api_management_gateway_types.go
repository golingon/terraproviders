// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_api_management_gateway

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataLocationDataAttributes struct {
	ref terra.Reference
}

func (ld DataLocationDataAttributes) InternalRef() (terra.Reference, error) {
	return ld.ref, nil
}

func (ld DataLocationDataAttributes) InternalWithRef(ref terra.Reference) DataLocationDataAttributes {
	return DataLocationDataAttributes{ref: ref}
}

func (ld DataLocationDataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ld.ref.InternalTokens()
}

func (ld DataLocationDataAttributes) City() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("city"))
}

func (ld DataLocationDataAttributes) District() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("district"))
}

func (ld DataLocationDataAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("name"))
}

func (ld DataLocationDataAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ld.ref.Append("region"))
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

type DataLocationDataState struct {
	City     string `json:"city"`
	District string `json:"district"`
	Name     string `json:"name"`
	Region   string `json:"region"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
