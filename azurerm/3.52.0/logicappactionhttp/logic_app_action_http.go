// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package logicappactionhttp

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RunAfter struct {
	// ActionName: string, required
	ActionName terra.StringValue `hcl:"action_name,attr" validate:"required"`
	// ActionResult: string, required
	ActionResult terra.StringValue `hcl:"action_result,attr" validate:"required"`
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

type RunAfterAttributes struct {
	ref terra.Reference
}

func (ra RunAfterAttributes) InternalRef() (terra.Reference, error) {
	return ra.ref, nil
}

func (ra RunAfterAttributes) InternalWithRef(ref terra.Reference) RunAfterAttributes {
	return RunAfterAttributes{ref: ref}
}

func (ra RunAfterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ra.ref.InternalTokens()
}

func (ra RunAfterAttributes) ActionName() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("action_name"))
}

func (ra RunAfterAttributes) ActionResult() terra.StringValue {
	return terra.ReferenceAsString(ra.ref.Append("action_result"))
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

type RunAfterState struct {
	ActionName   string `json:"action_name"`
	ActionResult string `json:"action_result"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
