// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package subnet

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Delegation struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServiceDelegation: required
	ServiceDelegation *ServiceDelegation `hcl:"service_delegation,block" validate:"required"`
}

type ServiceDelegation struct {
	// Actions: list of string, optional
	Actions terra.ListValue[terra.StringValue] `hcl:"actions,attr"`
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

type DelegationAttributes struct {
	ref terra.Reference
}

func (d DelegationAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DelegationAttributes) InternalWithRef(ref terra.Reference) DelegationAttributes {
	return DelegationAttributes{ref: ref}
}

func (d DelegationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DelegationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d DelegationAttributes) ServiceDelegation() terra.ListValue[ServiceDelegationAttributes] {
	return terra.ReferenceAsList[ServiceDelegationAttributes](d.ref.Append("service_delegation"))
}

type ServiceDelegationAttributes struct {
	ref terra.Reference
}

func (sd ServiceDelegationAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd ServiceDelegationAttributes) InternalWithRef(ref terra.Reference) ServiceDelegationAttributes {
	return ServiceDelegationAttributes{ref: ref}
}

func (sd ServiceDelegationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd ServiceDelegationAttributes) Actions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sd.ref.Append("actions"))
}

func (sd ServiceDelegationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("name"))
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

type DelegationState struct {
	Name              string                   `json:"name"`
	ServiceDelegation []ServiceDelegationState `json:"service_delegation"`
}

type ServiceDelegationState struct {
	Actions []string `json:"actions"`
	Name    string   `json:"name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
