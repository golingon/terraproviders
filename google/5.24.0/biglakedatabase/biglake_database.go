// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package biglakedatabase

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type HiveOptions struct {
	// LocationUri: string, optional
	LocationUri terra.StringValue `hcl:"location_uri,attr"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type HiveOptionsAttributes struct {
	ref terra.Reference
}

func (ho HiveOptionsAttributes) InternalRef() (terra.Reference, error) {
	return ho.ref, nil
}

func (ho HiveOptionsAttributes) InternalWithRef(ref terra.Reference) HiveOptionsAttributes {
	return HiveOptionsAttributes{ref: ref}
}

func (ho HiveOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ho.ref.InternalTokens()
}

func (ho HiveOptionsAttributes) LocationUri() terra.StringValue {
	return terra.ReferenceAsString(ho.ref.Append("location_uri"))
}

func (ho HiveOptionsAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ho.ref.Append("parameters"))
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

type HiveOptionsState struct {
	LocationUri string            `json:"location_uri"`
	Parameters  map[string]string `json:"parameters"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
