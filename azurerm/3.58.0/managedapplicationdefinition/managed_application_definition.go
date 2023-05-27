// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package managedapplicationdefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Authorization struct {
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// ServicePrincipalId: string, required
	ServicePrincipalId terra.StringValue `hcl:"service_principal_id,attr" validate:"required"`
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

type AuthorizationAttributes struct {
	ref terra.Reference
}

func (a AuthorizationAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AuthorizationAttributes) InternalWithRef(ref terra.Reference) AuthorizationAttributes {
	return AuthorizationAttributes{ref: ref}
}

func (a AuthorizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AuthorizationAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("role_definition_id"))
}

func (a AuthorizationAttributes) ServicePrincipalId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("service_principal_id"))
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

type AuthorizationState struct {
	RoleDefinitionId   string `json:"role_definition_id"`
	ServicePrincipalId string `json:"service_principal_id"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
