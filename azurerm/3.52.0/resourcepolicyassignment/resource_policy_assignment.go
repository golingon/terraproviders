// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package resourcepolicyassignment

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type NonComplianceMessage struct {
	// Content: string, required
	Content terra.StringValue `hcl:"content,attr" validate:"required"`
	// PolicyDefinitionReferenceId: string, optional
	PolicyDefinitionReferenceId terra.StringValue `hcl:"policy_definition_reference_id,attr"`
}

type Overrides struct {
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
	// OverridesSelectors: min=0
	Selectors []OverridesSelectors `hcl:"selectors,block" validate:"min=0"`
}

type OverridesSelectors struct {
	// In: list of string, optional
	In terra.ListValue[terra.StringValue] `hcl:"in,attr"`
	// NotIn: list of string, optional
	NotIn terra.ListValue[terra.StringValue] `hcl:"not_in,attr"`
}

type ResourceSelectors struct {
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResourceSelectorsSelectors: min=1
	Selectors []ResourceSelectorsSelectors `hcl:"selectors,block" validate:"min=1"`
}

type ResourceSelectorsSelectors struct {
	// In: list of string, optional
	In terra.ListValue[terra.StringValue] `hcl:"in,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// NotIn: list of string, optional
	NotIn terra.ListValue[terra.StringValue] `hcl:"not_in,attr"`
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

type IdentityAttributes struct {
	ref terra.Reference
}

func (i IdentityAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IdentityAttributes) InternalWithRef(ref terra.Reference) IdentityAttributes {
	return IdentityAttributes{ref: ref}
}

func (i IdentityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IdentityAttributes) IdentityIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("identity_ids"))
}

func (i IdentityAttributes) PrincipalId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("principal_id"))
}

func (i IdentityAttributes) TenantId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("tenant_id"))
}

func (i IdentityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type NonComplianceMessageAttributes struct {
	ref terra.Reference
}

func (ncm NonComplianceMessageAttributes) InternalRef() (terra.Reference, error) {
	return ncm.ref, nil
}

func (ncm NonComplianceMessageAttributes) InternalWithRef(ref terra.Reference) NonComplianceMessageAttributes {
	return NonComplianceMessageAttributes{ref: ref}
}

func (ncm NonComplianceMessageAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ncm.ref.InternalTokens()
}

func (ncm NonComplianceMessageAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("content"))
}

func (ncm NonComplianceMessageAttributes) PolicyDefinitionReferenceId() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("policy_definition_reference_id"))
}

type OverridesAttributes struct {
	ref terra.Reference
}

func (o OverridesAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OverridesAttributes) InternalWithRef(ref terra.Reference) OverridesAttributes {
	return OverridesAttributes{ref: ref}
}

func (o OverridesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OverridesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("value"))
}

func (o OverridesAttributes) Selectors() terra.ListValue[OverridesSelectorsAttributes] {
	return terra.ReferenceAsList[OverridesSelectorsAttributes](o.ref.Append("selectors"))
}

type OverridesSelectorsAttributes struct {
	ref terra.Reference
}

func (s OverridesSelectorsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s OverridesSelectorsAttributes) InternalWithRef(ref terra.Reference) OverridesSelectorsAttributes {
	return OverridesSelectorsAttributes{ref: ref}
}

func (s OverridesSelectorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s OverridesSelectorsAttributes) In() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("in"))
}

func (s OverridesSelectorsAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("kind"))
}

func (s OverridesSelectorsAttributes) NotIn() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("not_in"))
}

type ResourceSelectorsAttributes struct {
	ref terra.Reference
}

func (rs ResourceSelectorsAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceSelectorsAttributes) InternalWithRef(ref terra.Reference) ResourceSelectorsAttributes {
	return ResourceSelectorsAttributes{ref: ref}
}

func (rs ResourceSelectorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceSelectorsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("name"))
}

func (rs ResourceSelectorsAttributes) Selectors() terra.ListValue[ResourceSelectorsSelectorsAttributes] {
	return terra.ReferenceAsList[ResourceSelectorsSelectorsAttributes](rs.ref.Append("selectors"))
}

type ResourceSelectorsSelectorsAttributes struct {
	ref terra.Reference
}

func (s ResourceSelectorsSelectorsAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s ResourceSelectorsSelectorsAttributes) InternalWithRef(ref terra.Reference) ResourceSelectorsSelectorsAttributes {
	return ResourceSelectorsSelectorsAttributes{ref: ref}
}

func (s ResourceSelectorsSelectorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s ResourceSelectorsSelectorsAttributes) In() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("in"))
}

func (s ResourceSelectorsSelectorsAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("kind"))
}

func (s ResourceSelectorsSelectorsAttributes) NotIn() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](s.ref.Append("not_in"))
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

type OverridesState struct {
	Value     string                    `json:"value"`
	Selectors []OverridesSelectorsState `json:"selectors"`
}

type OverridesSelectorsState struct {
	In    []string `json:"in"`
	Kind  string   `json:"kind"`
	NotIn []string `json:"not_in"`
}

type ResourceSelectorsState struct {
	Name      string                            `json:"name"`
	Selectors []ResourceSelectorsSelectorsState `json:"selectors"`
}

type ResourceSelectorsSelectorsState struct {
	In    []string `json:"in"`
	Kind  string   `json:"kind"`
	NotIn []string `json:"not_in"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}