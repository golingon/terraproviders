// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datavpclatticelistener

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DefaultAction struct {
	// FixedResponse: min=0
	FixedResponse []FixedResponse `hcl:"fixed_response,block" validate:"min=0"`
	// Forward: min=0
	Forward []Forward `hcl:"forward,block" validate:"min=0"`
}

type FixedResponse struct{}

type Forward struct {
	// TargetGroups: min=0
	TargetGroups []TargetGroups `hcl:"target_groups,block" validate:"min=0"`
}

type TargetGroups struct{}

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

func (da DefaultActionAttributes) FixedResponse() terra.ListValue[FixedResponseAttributes] {
	return terra.ReferenceAsList[FixedResponseAttributes](da.ref.Append("fixed_response"))
}

func (da DefaultActionAttributes) Forward() terra.ListValue[ForwardAttributes] {
	return terra.ReferenceAsList[ForwardAttributes](da.ref.Append("forward"))
}

type FixedResponseAttributes struct {
	ref terra.Reference
}

func (fr FixedResponseAttributes) InternalRef() (terra.Reference, error) {
	return fr.ref, nil
}

func (fr FixedResponseAttributes) InternalWithRef(ref terra.Reference) FixedResponseAttributes {
	return FixedResponseAttributes{ref: ref}
}

func (fr FixedResponseAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return fr.ref.InternalTokens()
}

func (fr FixedResponseAttributes) StatusCode() terra.NumberValue {
	return terra.ReferenceAsNumber(fr.ref.Append("status_code"))
}

type ForwardAttributes struct {
	ref terra.Reference
}

func (f ForwardAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f ForwardAttributes) InternalWithRef(ref terra.Reference) ForwardAttributes {
	return ForwardAttributes{ref: ref}
}

func (f ForwardAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f ForwardAttributes) TargetGroups() terra.ListValue[TargetGroupsAttributes] {
	return terra.ReferenceAsList[TargetGroupsAttributes](f.ref.Append("target_groups"))
}

type TargetGroupsAttributes struct {
	ref terra.Reference
}

func (tg TargetGroupsAttributes) InternalRef() (terra.Reference, error) {
	return tg.ref, nil
}

func (tg TargetGroupsAttributes) InternalWithRef(ref terra.Reference) TargetGroupsAttributes {
	return TargetGroupsAttributes{ref: ref}
}

func (tg TargetGroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tg.ref.InternalTokens()
}

func (tg TargetGroupsAttributes) TargetGroupIdentifier() terra.StringValue {
	return terra.ReferenceAsString(tg.ref.Append("target_group_identifier"))
}

func (tg TargetGroupsAttributes) Weight() terra.NumberValue {
	return terra.ReferenceAsNumber(tg.ref.Append("weight"))
}

type DefaultActionState struct {
	FixedResponse []FixedResponseState `json:"fixed_response"`
	Forward       []ForwardState       `json:"forward"`
}

type FixedResponseState struct {
	StatusCode float64 `json:"status_code"`
}

type ForwardState struct {
	TargetGroups []TargetGroupsState `json:"target_groups"`
}

type TargetGroupsState struct {
	TargetGroupIdentifier string  `json:"target_group_identifier"`
	Weight                float64 `json:"weight"`
}
