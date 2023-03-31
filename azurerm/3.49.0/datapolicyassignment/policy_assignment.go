// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapolicyassignment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct{}

type NonComplianceMessage struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() terra.Reference {
	return i.ref
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() hclwrite.Tokens {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceString(i.ref.Append("type"))
}

type NonComplianceMessageAttributes struct {
	ref terra.Reference
}

func (ncm NonComplianceMessageAttributes) InternalRef() terra.Reference {
	return ncm.ref
}

func (ncm NonComplianceMessageAttributes) InternalWithRef(ref terra.Reference) NonComplianceMessageAttributes {
	return NonComplianceMessageAttributes{ref: ref}
}

func (ncm NonComplianceMessageAttributes) InternalTokens() hclwrite.Tokens {
	return ncm.ref.InternalTokens()
}

func (ncm NonComplianceMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceString(ncm.ref.Append("content"))
}

func (ncm NonComplianceMessageAttributes) PolicyDefinitionReferenceId() terra.StringValue {
	return terra.ReferenceString(ncm.ref.Append("policy_definition_reference_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type NonComplianceMessageState struct {
	Content                     string `json:"content"`
	PolicyDefinitionReferenceId string `json:"policy_definition_reference_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
