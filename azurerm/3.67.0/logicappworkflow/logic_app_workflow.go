// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package logicappworkflow

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AccessControl struct {
	// Action: optional
	Action *Action `hcl:"action,block"`
	// Content: optional
	Content *Content `hcl:"content,block"`
	// Trigger: optional
	Trigger *Trigger `hcl:"trigger,block"`
	// WorkflowManagement: optional
	WorkflowManagement *WorkflowManagement `hcl:"workflow_management,block"`
}

type Action struct {
	// AllowedCallerIpAddressRange: set of string, required
	AllowedCallerIpAddressRange terra.SetValue[terra.StringValue] `hcl:"allowed_caller_ip_address_range,attr" validate:"required"`
}

type Content struct {
	// AllowedCallerIpAddressRange: set of string, required
	AllowedCallerIpAddressRange terra.SetValue[terra.StringValue] `hcl:"allowed_caller_ip_address_range,attr" validate:"required"`
}

type Trigger struct {
	// AllowedCallerIpAddressRange: set of string, required
	AllowedCallerIpAddressRange terra.SetValue[terra.StringValue] `hcl:"allowed_caller_ip_address_range,attr" validate:"required"`
	// OpenAuthenticationPolicy: min=0
	OpenAuthenticationPolicy []OpenAuthenticationPolicy `hcl:"open_authentication_policy,block" validate:"min=0"`
}

type OpenAuthenticationPolicy struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Claim: min=1
	Claim []Claim `hcl:"claim,block" validate:"min=1"`
}

type Claim struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type WorkflowManagement struct {
	// AllowedCallerIpAddressRange: set of string, required
	AllowedCallerIpAddressRange terra.SetValue[terra.StringValue] `hcl:"allowed_caller_ip_address_range,attr" validate:"required"`
}

type Identity struct {
	// IdentityIds: set of string, optional
	IdentityIds terra.SetValue[terra.StringValue] `hcl:"identity_ids,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
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

type AccessControlAttributes struct {
	ref terra.Reference
}

func (ac AccessControlAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AccessControlAttributes) InternalWithRef(ref terra.Reference) AccessControlAttributes {
	return AccessControlAttributes{ref: ref}
}

func (ac AccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AccessControlAttributes) Action() terra.ListValue[ActionAttributes] {
	return terra.ReferenceAsList[ActionAttributes](ac.ref.Append("action"))
}

func (ac AccessControlAttributes) Content() terra.ListValue[ContentAttributes] {
	return terra.ReferenceAsList[ContentAttributes](ac.ref.Append("content"))
}

func (ac AccessControlAttributes) Trigger() terra.ListValue[TriggerAttributes] {
	return terra.ReferenceAsList[TriggerAttributes](ac.ref.Append("trigger"))
}

func (ac AccessControlAttributes) WorkflowManagement() terra.ListValue[WorkflowManagementAttributes] {
	return terra.ReferenceAsList[WorkflowManagementAttributes](ac.ref.Append("workflow_management"))
}

type ActionAttributes struct {
	ref terra.Reference
}

func (a ActionAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a ActionAttributes) InternalWithRef(ref terra.Reference) ActionAttributes {
	return ActionAttributes{ref: ref}
}

func (a ActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a ActionAttributes) AllowedCallerIpAddressRange() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("allowed_caller_ip_address_range"))
}

type ContentAttributes struct {
	ref terra.Reference
}

func (c ContentAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ContentAttributes) InternalWithRef(ref terra.Reference) ContentAttributes {
	return ContentAttributes{ref: ref}
}

func (c ContentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ContentAttributes) AllowedCallerIpAddressRange() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](c.ref.Append("allowed_caller_ip_address_range"))
}

type TriggerAttributes struct {
	ref terra.Reference
}

func (t TriggerAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TriggerAttributes) InternalWithRef(ref terra.Reference) TriggerAttributes {
	return TriggerAttributes{ref: ref}
}

func (t TriggerAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TriggerAttributes) AllowedCallerIpAddressRange() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](t.ref.Append("allowed_caller_ip_address_range"))
}

func (t TriggerAttributes) OpenAuthenticationPolicy() terra.SetValue[OpenAuthenticationPolicyAttributes] {
	return terra.ReferenceAsSet[OpenAuthenticationPolicyAttributes](t.ref.Append("open_authentication_policy"))
}

type OpenAuthenticationPolicyAttributes struct {
	ref terra.Reference
}

func (oap OpenAuthenticationPolicyAttributes) InternalRef() (terra.Reference, error) {
	return oap.ref, nil
}

func (oap OpenAuthenticationPolicyAttributes) InternalWithRef(ref terra.Reference) OpenAuthenticationPolicyAttributes {
	return OpenAuthenticationPolicyAttributes{ref: ref}
}

func (oap OpenAuthenticationPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oap.ref.InternalTokens()
}

func (oap OpenAuthenticationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(oap.ref.Append("name"))
}

func (oap OpenAuthenticationPolicyAttributes) Claim() terra.SetValue[ClaimAttributes] {
	return terra.ReferenceAsSet[ClaimAttributes](oap.ref.Append("claim"))
}

type ClaimAttributes struct {
	ref terra.Reference
}

func (c ClaimAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ClaimAttributes) InternalWithRef(ref terra.Reference) ClaimAttributes {
	return ClaimAttributes{ref: ref}
}

func (c ClaimAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ClaimAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("name"))
}

func (c ClaimAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type WorkflowManagementAttributes struct {
	ref terra.Reference
}

func (wm WorkflowManagementAttributes) InternalRef() (terra.Reference, error) {
	return wm.ref, nil
}

func (wm WorkflowManagementAttributes) InternalWithRef(ref terra.Reference) WorkflowManagementAttributes {
	return WorkflowManagementAttributes{ref: ref}
}

func (wm WorkflowManagementAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return wm.ref.InternalTokens()
}

func (wm WorkflowManagementAttributes) AllowedCallerIpAddressRange() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](wm.ref.Append("allowed_caller_ip_address_range"))
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

type AccessControlState struct {
	Action             []ActionState             `json:"action"`
	Content            []ContentState            `json:"content"`
	Trigger            []TriggerState            `json:"trigger"`
	WorkflowManagement []WorkflowManagementState `json:"workflow_management"`
}

type ActionState struct {
	AllowedCallerIpAddressRange []string `json:"allowed_caller_ip_address_range"`
}

type ContentState struct {
	AllowedCallerIpAddressRange []string `json:"allowed_caller_ip_address_range"`
}

type TriggerState struct {
	AllowedCallerIpAddressRange []string                        `json:"allowed_caller_ip_address_range"`
	OpenAuthenticationPolicy    []OpenAuthenticationPolicyState `json:"open_authentication_policy"`
}

type OpenAuthenticationPolicyState struct {
	Name  string       `json:"name"`
	Claim []ClaimState `json:"claim"`
}

type ClaimState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type WorkflowManagementState struct {
	AllowedCallerIpAddressRange []string `json:"allowed_caller_ip_address_range"`
}

type IdentityState struct {
	IdentityIds []string `json:"identity_ids"`
	PrincipalId string   `json:"principal_id"`
	TenantId    string   `json:"tenant_id"`
	Type        string   `json:"type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
