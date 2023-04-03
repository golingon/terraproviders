// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package siterecoveryprotectioncontainermapping

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutomaticUpdate struct {
	// AutomationAccountId: string, optional
	AutomationAccountId terra.StringValue `hcl:"automation_account_id,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
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

type AutomaticUpdateAttributes struct {
	ref terra.Reference
}

func (au AutomaticUpdateAttributes) InternalRef() (terra.Reference, error) {
	return au.ref, nil
}

func (au AutomaticUpdateAttributes) InternalWithRef(ref terra.Reference) AutomaticUpdateAttributes {
	return AutomaticUpdateAttributes{ref: ref}
}

func (au AutomaticUpdateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return au.ref.InternalTokens()
}

func (au AutomaticUpdateAttributes) AutomationAccountId() terra.StringValue {
	return terra.ReferenceAsString(au.ref.Append("automation_account_id"))
}

func (au AutomaticUpdateAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(au.ref.Append("enabled"))
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

type AutomaticUpdateState struct {
	AutomationAccountId string `json:"automation_account_id"`
	Enabled             bool   `json:"enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}