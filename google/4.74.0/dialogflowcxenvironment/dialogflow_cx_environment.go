// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dialogflowcxenvironment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type VersionConfigs struct {
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
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

type VersionConfigsAttributes struct {
	ref terra.Reference
}

func (vc VersionConfigsAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VersionConfigsAttributes) InternalWithRef(ref terra.Reference) VersionConfigsAttributes {
	return VersionConfigsAttributes{ref: ref}
}

func (vc VersionConfigsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VersionConfigsAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("version"))
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}

type VersionConfigsState struct {
	Version string `json:"version"`
}