// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpclattice_listener

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DefaultAction struct {
	// DefaultActionFixedResponse: optional
	FixedResponse *DefaultActionFixedResponse `hcl:"fixed_response,block"`
	// DefaultActionForward: min=0
	Forward []DefaultActionForward `hcl:"forward,block" validate:"min=0"`
}

type DefaultActionFixedResponse struct {
	// StatusCode: number, required
	StatusCode terra.NumberValue `hcl:"status_code,attr" validate:"required"`
}

type DefaultActionForward struct {
	// DefaultActionForwardTargetGroups: min=0
	TargetGroups []DefaultActionForwardTargetGroups `hcl:"target_groups,block" validate:"min=0"`
}

type DefaultActionForwardTargetGroups struct {
	// TargetGroupIdentifier: string, optional
	TargetGroupIdentifier terra.StringValue `hcl:"target_group_identifier,attr"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type DefaultActionAttributes struct {
	ref terra.Reference
}

func (da DefaultActionAttributes) InternalRef() (terra.Reference, error) {
	return da.ref, nil
}

func (da DefaultActionAttributes) InternalWithRef(ref terra.Reference) DefaultActionAttributes {
	return DefaultActionAttributes{ref: ref}
}

func (da DefaultActionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return da.ref.InternalTokens()
}

func (da DefaultActionAttributes) FixedResponse() terra.ListValue[DefaultActionFixedResponseAttributes] {
	return terra.ReferenceAsList[DefaultActionFixedResponseAttributes](da.ref.Append("fixed_response"))
}

func (da DefaultActionAttributes) Forward() terra.ListValue[DefaultActionForwardAttributes] {
	return terra.ReferenceAsList[DefaultActionForwardAttributes](da.ref.Append("forward"))
}

type DefaultActionFixedResponseAttributes struct {
	ref terra.Reference
}

func (fr DefaultActionFixedResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr DefaultActionFixedResponseAttributes) InternalWithRef(ref terra.Reference) DefaultActionFixedResponseAttributes {
	return DefaultActionFixedResponseAttributes{ref: ref}
}

func (fr DefaultActionFixedResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr DefaultActionFixedResponseAttributes) StatusCode() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("status_code"))
}

type DefaultActionForwardAttributes struct {
	ref terra.Reference
}

func (f DefaultActionForwardAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DefaultActionForwardAttributes) InternalWithRef(ref terra.Reference) DefaultActionForwardAttributes {
	return DefaultActionForwardAttributes{ref: ref}
}

func (f DefaultActionForwardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DefaultActionForwardAttributes) TargetGroups() terra.ListValue[DefaultActionForwardTargetGroupsAttributes] {
	return terra.ReferenceAsList[DefaultActionForwardTargetGroupsAttributes](f.ref.Append("target_groups"))
}

type DefaultActionForwardTargetGroupsAttributes struct {
	ref terra.Reference
}

func (tg DefaultActionForwardTargetGroupsAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg DefaultActionForwardTargetGroupsAttributes) InternalWithRef(ref terra.Reference) DefaultActionForwardTargetGroupsAttributes {
	return DefaultActionForwardTargetGroupsAttributes{ref: ref}
}

func (tg DefaultActionForwardTargetGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg DefaultActionForwardTargetGroupsAttributes) TargetGroupIdentifier() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("target_group_identifier"))
}

func (tg DefaultActionForwardTargetGroupsAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tg.ref.Append("weight"))
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

type DefaultActionState struct {
	FixedResponse []DefaultActionFixedResponseState `json:"fixed_response"`
	Forward       []DefaultActionForwardState       `json:"forward"`
}

type DefaultActionFixedResponseState struct {
	StatusCode float64 `json:"status_code"`
}

type DefaultActionForwardState struct {
	TargetGroups []DefaultActionForwardTargetGroupsState `json:"target_groups"`
}

type DefaultActionForwardTargetGroupsState struct {
	TargetGroupIdentifier string  `json:"target_group_identifier"`
	Weight                float64 `json:"weight"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
