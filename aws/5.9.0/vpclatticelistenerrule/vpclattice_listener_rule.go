// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package vpclatticelistenerrule

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Action struct {
	// FixedResponse: optional
	FixedResponse *FixedResponse `hcl:"fixed_response,block"`
	// Forward: optional
	Forward *Forward `hcl:"forward,block"`
}

type FixedResponse struct {
	// StatusCode: number, required
	StatusCode terra.NumberValue `hcl:"status_code,attr" validate:"required"`
}

type Forward struct {
	// TargetGroups: min=1,max=2
	TargetGroups []TargetGroups `hcl:"target_groups,block" validate:"min=1,max=2"`
}

type TargetGroups struct {
	// TargetGroupIdentifier: string, required
	TargetGroupIdentifier terra.StringValue `hcl:"target_group_identifier,attr" validate:"required"`
	// Weight: number, optional
	Weight terra.NumberValue `hcl:"weight,attr"`
}

type Match struct {
	// HttpMatch: optional
	HttpMatch *HttpMatch `hcl:"http_match,block"`
}

type HttpMatch struct {
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// HeaderMatches: min=0,max=5
	HeaderMatches []HeaderMatches `hcl:"header_matches,block" validate:"min=0,max=5"`
	// PathMatch: optional
	PathMatch *PathMatch `hcl:"path_match,block"`
}

type HeaderMatches struct {
	// CaseSensitive: bool, optional
	CaseSensitive terra.BoolValue `hcl:"case_sensitive,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// HeaderMatchesMatch: required
	Match *HeaderMatchesMatch `hcl:"match,block" validate:"required"`
}

type HeaderMatchesMatch struct {
	// Contains: string, optional
	Contains terra.StringValue `hcl:"contains,attr"`
	// Exact: string, optional
	Exact terra.StringValue `hcl:"exact,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}

type PathMatch struct {
	// CaseSensitive: bool, optional
	CaseSensitive terra.BoolValue `hcl:"case_sensitive,attr"`
	// PathMatchMatch: required
	Match *PathMatchMatch `hcl:"match,block" validate:"required"`
}

type PathMatchMatch struct {
	// Exact: string, optional
	Exact terra.StringValue `hcl:"exact,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
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

func (a ActionAttributes) FixedResponse() terra.ListValue[FixedResponseAttributes] {
	return terra.ReferenceAsList[FixedResponseAttributes](a.ref.Append("fixed_response"))
}

func (a ActionAttributes) Forward() terra.ListValue[ForwardAttributes] {
	return terra.ReferenceAsList[ForwardAttributes](a.ref.Append("forward"))
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

type MatchAttributes struct {
	ref terra.Reference
}

func (m MatchAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MatchAttributes) InternalWithRef(ref terra.Reference) MatchAttributes {
	return MatchAttributes{ref: ref}
}

func (m MatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MatchAttributes) HttpMatch() terra.ListValue[HttpMatchAttributes] {
	return terra.ReferenceAsList[HttpMatchAttributes](m.ref.Append("http_match"))
}

type HttpMatchAttributes struct {
	ref terra.Reference
}

func (hm HttpMatchAttributes) InternalRef() (terra.Reference, error) {
	return hm.ref, nil
}

func (hm HttpMatchAttributes) InternalWithRef(ref terra.Reference) HttpMatchAttributes {
	return HttpMatchAttributes{ref: ref}
}

func (hm HttpMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hm.ref.InternalTokens()
}

func (hm HttpMatchAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("method"))
}

func (hm HttpMatchAttributes) HeaderMatches() terra.ListValue[HeaderMatchesAttributes] {
	return terra.ReferenceAsList[HeaderMatchesAttributes](hm.ref.Append("header_matches"))
}

func (hm HttpMatchAttributes) PathMatch() terra.ListValue[PathMatchAttributes] {
	return terra.ReferenceAsList[PathMatchAttributes](hm.ref.Append("path_match"))
}

type HeaderMatchesAttributes struct {
	ref terra.Reference
}

func (hm HeaderMatchesAttributes) InternalRef() (terra.Reference, error) {
	return hm.ref, nil
}

func (hm HeaderMatchesAttributes) InternalWithRef(ref terra.Reference) HeaderMatchesAttributes {
	return HeaderMatchesAttributes{ref: ref}
}

func (hm HeaderMatchesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return hm.ref.InternalTokens()
}

func (hm HeaderMatchesAttributes) CaseSensitive() terra.BoolValue {
	return terra.ReferenceAsBool(hm.ref.Append("case_sensitive"))
}

func (hm HeaderMatchesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hm.ref.Append("name"))
}

func (hm HeaderMatchesAttributes) Match() terra.ListValue[HeaderMatchesMatchAttributes] {
	return terra.ReferenceAsList[HeaderMatchesMatchAttributes](hm.ref.Append("match"))
}

type HeaderMatchesMatchAttributes struct {
	ref terra.Reference
}

func (m HeaderMatchesMatchAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m HeaderMatchesMatchAttributes) InternalWithRef(ref terra.Reference) HeaderMatchesMatchAttributes {
	return HeaderMatchesMatchAttributes{ref: ref}
}

func (m HeaderMatchesMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m HeaderMatchesMatchAttributes) Contains() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("contains"))
}

func (m HeaderMatchesMatchAttributes) Exact() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("exact"))
}

func (m HeaderMatchesMatchAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("prefix"))
}

type PathMatchAttributes struct {
	ref terra.Reference
}

func (pm PathMatchAttributes) InternalRef() (terra.Reference, error) {
	return pm.ref, nil
}

func (pm PathMatchAttributes) InternalWithRef(ref terra.Reference) PathMatchAttributes {
	return PathMatchAttributes{ref: ref}
}

func (pm PathMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pm.ref.InternalTokens()
}

func (pm PathMatchAttributes) CaseSensitive() terra.BoolValue {
	return terra.ReferenceAsBool(pm.ref.Append("case_sensitive"))
}

func (pm PathMatchAttributes) Match() terra.ListValue[PathMatchMatchAttributes] {
	return terra.ReferenceAsList[PathMatchMatchAttributes](pm.ref.Append("match"))
}

type PathMatchMatchAttributes struct {
	ref terra.Reference
}

func (m PathMatchMatchAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m PathMatchMatchAttributes) InternalWithRef(ref terra.Reference) PathMatchMatchAttributes {
	return PathMatchMatchAttributes{ref: ref}
}

func (m PathMatchMatchAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m PathMatchMatchAttributes) Exact() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("exact"))
}

func (m PathMatchMatchAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("prefix"))
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

type ActionState struct {
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

type MatchState struct {
	HttpMatch []HttpMatchState `json:"http_match"`
}

type HttpMatchState struct {
	Method        string               `json:"method"`
	HeaderMatches []HeaderMatchesState `json:"header_matches"`
	PathMatch     []PathMatchState     `json:"path_match"`
}

type HeaderMatchesState struct {
	CaseSensitive bool                      `json:"case_sensitive"`
	Name          string                    `json:"name"`
	Match         []HeaderMatchesMatchState `json:"match"`
}

type HeaderMatchesMatchState struct {
	Contains string `json:"contains"`
	Exact    string `json:"exact"`
	Prefix   string `json:"prefix"`
}

type PathMatchState struct {
	CaseSensitive bool                  `json:"case_sensitive"`
	Match         []PathMatchMatchState `json:"match"`
}

type PathMatchMatchState struct {
	Exact  string `json:"exact"`
	Prefix string `json:"prefix"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
