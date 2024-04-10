// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package networkmanager

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CrossTenantScopes struct{}

type Scope struct {
	// ManagementGroupIds: list of string, optional
	ManagementGroupIds terra.ListValue[terra.StringValue] `hcl:"management_group_ids,attr"`
	// SubscriptionIds: list of string, optional
	SubscriptionIds terra.ListValue[terra.StringValue] `hcl:"subscription_ids,attr"`
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

type CrossTenantScopesAttributes struct {
	ref terra.Reference
}

func (cts CrossTenantScopesAttributes) InternalRef() (terra.Reference, error) {
	return cts.ref, nil
}

func (cts CrossTenantScopesAttributes) InternalWithRef(ref terra.Reference) CrossTenantScopesAttributes {
	return CrossTenantScopesAttributes{ref: ref}
}

func (cts CrossTenantScopesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cts.ref.InternalTokens()
}

func (cts CrossTenantScopesAttributes) ManagementGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cts.ref.Append("management_groups"))
}

func (cts CrossTenantScopesAttributes) Subscriptions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cts.ref.Append("subscriptions"))
}

func (cts CrossTenantScopesAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(cts.ref.Append("tenant_id"))
}

type ScopeAttributes struct {
	ref terra.Reference
}

func (s ScopeAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ScopeAttributes) InternalWithRef(ref terra.Reference) ScopeAttributes {
	return ScopeAttributes{ref: ref}
}

func (s ScopeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ScopeAttributes) ManagementGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("management_group_ids"))
}

func (s ScopeAttributes) SubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("subscription_ids"))
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

type CrossTenantScopesState struct {
	ManagementGroups []string `json:"management_groups"`
	Subscriptions    []string `json:"subscriptions"`
	TenantId         string   `json:"tenant_id"`
}

type ScopeState struct {
	ManagementGroupIds []string `json:"management_group_ids"`
	SubscriptionIds    []string `json:"subscription_ids"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
