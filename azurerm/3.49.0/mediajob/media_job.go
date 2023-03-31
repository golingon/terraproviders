// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediajob

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InputAsset struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type OutputAsset struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
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

type InputAssetAttributes struct {
	ref terra.Reference
}

func (ia InputAssetAttributes) InternalRef() terra.Reference {
	return ia.ref
}

func (ia InputAssetAttributes) InternalWithRef(ref terra.Reference) InputAssetAttributes {
	return InputAssetAttributes{ref: ref}
}

func (ia InputAssetAttributes) InternalTokens() hclwrite.Tokens {
	return ia.ref.InternalTokens()
}

func (ia InputAssetAttributes) Label() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("label"))
}

func (ia InputAssetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ia.ref.Append("name"))
}

type OutputAssetAttributes struct {
	ref terra.Reference
}

func (oa OutputAssetAttributes) InternalRef() terra.Reference {
	return oa.ref
}

func (oa OutputAssetAttributes) InternalWithRef(ref terra.Reference) OutputAssetAttributes {
	return OutputAssetAttributes{ref: ref}
}

func (oa OutputAssetAttributes) InternalTokens() hclwrite.Tokens {
	return oa.ref.InternalTokens()
}

func (oa OutputAssetAttributes) Label() terra.StringValue {
	return terra.ReferenceString(oa.ref.Append("label"))
}

func (oa OutputAssetAttributes) Name() terra.StringValue {
	return terra.ReferenceString(oa.ref.Append("name"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type InputAssetState struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type OutputAssetState struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
