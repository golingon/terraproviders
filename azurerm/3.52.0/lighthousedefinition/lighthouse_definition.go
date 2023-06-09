// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package lighthousedefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Authorization struct {
	// DelegatedRoleDefinitionIds: set of string, optional
	DelegatedRoleDefinitionIds terra.SetValue[terra.StringValue] `hcl:"delegated_role_definition_ids,attr"`
	// PrincipalDisplayName: string, optional
	PrincipalDisplayName terra.StringValue `hcl:"principal_display_name,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
}

type EligibleAuthorization struct {
	// PrincipalDisplayName: string, optional
	PrincipalDisplayName terra.StringValue `hcl:"principal_display_name,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
	// RoleDefinitionId: string, required
	RoleDefinitionId terra.StringValue `hcl:"role_definition_id,attr" validate:"required"`
	// JustInTimeAccessPolicy: optional
	JustInTimeAccessPolicy *JustInTimeAccessPolicy `hcl:"just_in_time_access_policy,block"`
}

type JustInTimeAccessPolicy struct {
	// MaximumActivationDuration: string, optional
	MaximumActivationDuration terra.StringValue `hcl:"maximum_activation_duration,attr"`
	// MultiFactorAuthProvider: string, optional
	MultiFactorAuthProvider terra.StringValue `hcl:"multi_factor_auth_provider,attr"`
	// Approver: min=0
	Approver []Approver `hcl:"approver,block" validate:"min=0"`
}

type Approver struct {
	// PrincipalDisplayName: string, optional
	PrincipalDisplayName terra.StringValue `hcl:"principal_display_name,attr"`
	// PrincipalId: string, required
	PrincipalId terra.StringValue `hcl:"principal_id,attr" validate:"required"`
}

type Plan struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Product: string, required
	Product terra.StringValue `hcl:"product,attr" validate:"required"`
	// Publisher: string, required
	Publisher terra.StringValue `hcl:"publisher,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
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

func (a AuthorizationAttributes) DelegatedRoleDefinitionIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("delegated_role_definition_ids"))
}

func (a AuthorizationAttributes) PrincipalDisplayName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("principal_display_name"))
}

func (a AuthorizationAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("principal_id"))
}

func (a AuthorizationAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("role_definition_id"))
}

type EligibleAuthorizationAttributes struct {
	ref terra.Reference
}

func (ea EligibleAuthorizationAttributes) InternalRef() (terra.Reference, error) {
	return ea.ref, nil
}

func (ea EligibleAuthorizationAttributes) InternalWithRef(ref terra.Reference) EligibleAuthorizationAttributes {
	return EligibleAuthorizationAttributes{ref: ref}
}

func (ea EligibleAuthorizationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ea.ref.InternalTokens()
}

func (ea EligibleAuthorizationAttributes) PrincipalDisplayName() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("principal_display_name"))
}

func (ea EligibleAuthorizationAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("principal_id"))
}

func (ea EligibleAuthorizationAttributes) RoleDefinitionId() terra.StringValue {
	return terra.ReferenceAsString(ea.ref.Append("role_definition_id"))
}

func (ea EligibleAuthorizationAttributes) JustInTimeAccessPolicy() terra.ListValue[JustInTimeAccessPolicyAttributes] {
	return terra.ReferenceAsList[JustInTimeAccessPolicyAttributes](ea.ref.Append("just_in_time_access_policy"))
}

type JustInTimeAccessPolicyAttributes struct {
	ref terra.Reference
}

func (jitap JustInTimeAccessPolicyAttributes) InternalRef() (terra.Reference, error) {
	return jitap.ref, nil
}

func (jitap JustInTimeAccessPolicyAttributes) InternalWithRef(ref terra.Reference) JustInTimeAccessPolicyAttributes {
	return JustInTimeAccessPolicyAttributes{ref: ref}
}

func (jitap JustInTimeAccessPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return jitap.ref.InternalTokens()
}

func (jitap JustInTimeAccessPolicyAttributes) MaximumActivationDuration() terra.StringValue {
	return terra.ReferenceAsString(jitap.ref.Append("maximum_activation_duration"))
}

func (jitap JustInTimeAccessPolicyAttributes) MultiFactorAuthProvider() terra.StringValue {
	return terra.ReferenceAsString(jitap.ref.Append("multi_factor_auth_provider"))
}

func (jitap JustInTimeAccessPolicyAttributes) Approver() terra.SetValue[ApproverAttributes] {
	return terra.ReferenceAsSet[ApproverAttributes](jitap.ref.Append("approver"))
}

type ApproverAttributes struct {
	ref terra.Reference
}

func (a ApproverAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ApproverAttributes) InternalWithRef(ref terra.Reference) ApproverAttributes {
	return ApproverAttributes{ref: ref}
}

func (a ApproverAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ApproverAttributes) PrincipalDisplayName() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("principal_display_name"))
}

func (a ApproverAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("principal_id"))
}

type PlanAttributes struct {
	ref terra.Reference
}

func (p PlanAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PlanAttributes) InternalWithRef(ref terra.Reference) PlanAttributes {
	return PlanAttributes{ref: ref}
}

func (p PlanAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PlanAttributes) Product() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("product"))
}

func (p PlanAttributes) Publisher() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("publisher"))
}

func (p PlanAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("version"))
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
	DelegatedRoleDefinitionIds []string `json:"delegated_role_definition_ids"`
	PrincipalDisplayName       string   `json:"principal_display_name"`
	PrincipalId                string   `json:"principal_id"`
	RoleDefinitionId           string   `json:"role_definition_id"`
}

type EligibleAuthorizationState struct {
	PrincipalDisplayName   string                        `json:"principal_display_name"`
	PrincipalId            string                        `json:"principal_id"`
	RoleDefinitionId       string                        `json:"role_definition_id"`
	JustInTimeAccessPolicy []JustInTimeAccessPolicyState `json:"just_in_time_access_policy"`
}

type JustInTimeAccessPolicyState struct {
	MaximumActivationDuration string          `json:"maximum_activation_duration"`
	MultiFactorAuthProvider   string          `json:"multi_factor_auth_provider"`
	Approver                  []ApproverState `json:"approver"`
}

type ApproverState struct {
	PrincipalDisplayName string `json:"principal_display_name"`
	PrincipalId          string `json:"principal_id"`
}

type PlanState struct {
	Name      string `json:"name"`
	Product   string `json:"product"`
	Publisher string `json:"publisher"`
	Version   string `json:"version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
