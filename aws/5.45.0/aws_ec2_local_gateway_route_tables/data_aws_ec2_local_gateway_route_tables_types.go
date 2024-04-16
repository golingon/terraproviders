// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ec2_local_gateway_route_tables

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

type DataFilterState struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
