// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cloudidentitygroup

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GroupKey struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Namespace: string, optional
	Namespace terra.StringValue `hcl:"namespace,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type GroupKeyAttributes struct {
	ref terra.Reference
}

func (gk GroupKeyAttributes) InternalRef() (terra.Reference, error) {
	return gk.ref, nil
}

func (gk GroupKeyAttributes) InternalWithRef(ref terra.Reference) GroupKeyAttributes {
	return GroupKeyAttributes{ref: ref}
}

func (gk GroupKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gk.ref.InternalTokens()
}

func (gk GroupKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("id"))
}

func (gk GroupKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("namespace"))
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

type GroupKeyState struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
