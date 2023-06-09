// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package networkmanagercorenetwork

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Edges struct{}

type Segments struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type EdgesAttributes struct {
	ref terra.Reference
}

func (e EdgesAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EdgesAttributes) InternalWithRef(ref terra.Reference) EdgesAttributes {
	return EdgesAttributes{ref: ref}
}

func (e EdgesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EdgesAttributes) Asn() terra.NumberValue {
	return terra.ReferenceAsNumber(e.ref.Append("asn"))
}

func (e EdgesAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("edge_location"))
}

func (e EdgesAttributes) InsideCidrBlocks() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](e.ref.Append("inside_cidr_blocks"))
}

type SegmentsAttributes struct {
	ref terra.Reference
}

func (s SegmentsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SegmentsAttributes) InternalWithRef(ref terra.Reference) SegmentsAttributes {
	return SegmentsAttributes{ref: ref}
}

func (s SegmentsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SegmentsAttributes) EdgeLocations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("edge_locations"))
}

func (s SegmentsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SegmentsAttributes) SharedSegments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("shared_segments"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type EdgesState struct {
	Asn              float64  `json:"asn"`
	EdgeLocation     string   `json:"edge_location"`
	InsideCidrBlocks []string `json:"inside_cidr_blocks"`
}

type SegmentsState struct {
	EdgeLocations  []string `json:"edge_locations"`
	Name           string   `json:"name"`
	SharedSegments []string `json:"shared_segments"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
