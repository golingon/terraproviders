// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package computeinterconnectattachment

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PrivateInterconnectInfo struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type PrivateInterconnectInfoAttributes struct {
	ref terra.Reference
}

func (pii PrivateInterconnectInfoAttributes) InternalRef() (terra.Reference, error) {
	return pii.ref, nil
}

func (pii PrivateInterconnectInfoAttributes) InternalWithRef(ref terra.Reference) PrivateInterconnectInfoAttributes {
	return PrivateInterconnectInfoAttributes{ref: ref}
}

func (pii PrivateInterconnectInfoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pii.ref.InternalTokens()
}

func (pii PrivateInterconnectInfoAttributes) Tag8021Q() terra.NumberValue {
	return terra.ReferenceAsNumber(pii.ref.Append("tag8021q"))
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

type PrivateInterconnectInfoState struct {
	Tag8021Q float64 `json:"tag8021q"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
