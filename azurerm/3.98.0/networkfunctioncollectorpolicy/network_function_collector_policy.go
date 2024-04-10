// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package networkfunctioncollectorpolicy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type IpfxEmission struct {
	// DestinationTypes: list of string, required
	DestinationTypes terra.ListValue[terra.StringValue] `hcl:"destination_types,attr" validate:"required"`
}

type IpfxIngestion struct {
	// SourceResourceIds: list of string, required
	SourceResourceIds terra.ListValue[terra.StringValue] `hcl:"source_resource_ids,attr" validate:"required"`
}

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

type IpfxEmissionAttributes struct {
	ref terra.Reference
}

func (ie IpfxEmissionAttributes) InternalRef() (terra.Reference, error) {
	return ie.ref, nil
}

func (ie IpfxEmissionAttributes) InternalWithRef(ref terra.Reference) IpfxEmissionAttributes {
	return IpfxEmissionAttributes{ref: ref}
}

func (ie IpfxEmissionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ie.ref.InternalTokens()
}

func (ie IpfxEmissionAttributes) DestinationTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ie.ref.Append("destination_types"))
}

type IpfxIngestionAttributes struct {
	ref terra.Reference
}

func (ii IpfxIngestionAttributes) InternalRef() (terra.Reference, error) {
	return ii.ref, nil
}

func (ii IpfxIngestionAttributes) InternalWithRef(ref terra.Reference) IpfxIngestionAttributes {
	return IpfxIngestionAttributes{ref: ref}
}

func (ii IpfxIngestionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ii.ref.InternalTokens()
}

func (ii IpfxIngestionAttributes) SourceResourceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ii.ref.Append("source_resource_ids"))
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

type IpfxEmissionState struct {
	DestinationTypes []string `json:"destination_types"`
}

type IpfxIngestionState struct {
	SourceResourceIds []string `json:"source_resource_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
