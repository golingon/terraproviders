// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package accesscontextmanagerserviceperimeteringresspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IngressFrom struct {
	// Identities: list of string, optional
	Identities terra.ListValue[terra.StringValue] `hcl:"identities,attr"`
	// IdentityType: string, optional
	IdentityType terra.StringValue `hcl:"identity_type,attr"`
	// Sources: min=0
	Sources []Sources `hcl:"sources,block" validate:"min=0"`
}

type Sources struct {
	// AccessLevel: string, optional
	AccessLevel terra.StringValue `hcl:"access_level,attr"`
	// Resource: string, optional
	Resource terra.StringValue `hcl:"resource,attr"`
}

type IngressTo struct {
	// Resources: list of string, optional
	Resources terra.ListValue[terra.StringValue] `hcl:"resources,attr"`
	// Operations: min=0
	Operations []Operations `hcl:"operations,block" validate:"min=0"`
}

type Operations struct {
	// ServiceName: string, optional
	ServiceName terra.StringValue `hcl:"service_name,attr"`
	// MethodSelectors: min=0
	MethodSelectors []MethodSelectors `hcl:"method_selectors,block" validate:"min=0"`
}

type MethodSelectors struct {
	// Method: string, optional
	Method terra.StringValue `hcl:"method,attr"`
	// Permission: string, optional
	Permission terra.StringValue `hcl:"permission,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type IngressFromAttributes struct {
	ref terra.Reference
}

func (_if IngressFromAttributes) InternalRef() (terra.Reference, error) {
	return _if.ref, nil
}

func (_if IngressFromAttributes) InternalWithRef(ref terra.Reference) IngressFromAttributes {
	return IngressFromAttributes{ref: ref}
}

func (_if IngressFromAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return _if.ref.InternalTokens()
}

func (_if IngressFromAttributes) Identities() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](_if.ref.Append("identities"))
}

func (_if IngressFromAttributes) IdentityType() terra.StringValue {
	return terra.ReferenceAsString(_if.ref.Append("identity_type"))
}

func (_if IngressFromAttributes) Sources() terra.ListValue[SourcesAttributes] {
	return terra.ReferenceAsList[SourcesAttributes](_if.ref.Append("sources"))
}

type SourcesAttributes struct {
	ref terra.Reference
}

func (s SourcesAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SourcesAttributes) InternalWithRef(ref terra.Reference) SourcesAttributes {
	return SourcesAttributes{ref: ref}
}

func (s SourcesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SourcesAttributes) AccessLevel() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("access_level"))
}

func (s SourcesAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("resource"))
}

type IngressToAttributes struct {
	ref terra.Reference
}

func (it IngressToAttributes) InternalRef() (terra.Reference, error) {
	return it.ref, nil
}

func (it IngressToAttributes) InternalWithRef(ref terra.Reference) IngressToAttributes {
	return IngressToAttributes{ref: ref}
}

func (it IngressToAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return it.ref.InternalTokens()
}

func (it IngressToAttributes) Resources() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](it.ref.Append("resources"))
}

func (it IngressToAttributes) Operations() terra.ListValue[OperationsAttributes] {
	return terra.ReferenceAsList[OperationsAttributes](it.ref.Append("operations"))
}

type OperationsAttributes struct {
	ref terra.Reference
}

func (o OperationsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OperationsAttributes) InternalWithRef(ref terra.Reference) OperationsAttributes {
	return OperationsAttributes{ref: ref}
}

func (o OperationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OperationsAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("service_name"))
}

func (o OperationsAttributes) MethodSelectors() terra.ListValue[MethodSelectorsAttributes] {
	return terra.ReferenceAsList[MethodSelectorsAttributes](o.ref.Append("method_selectors"))
}

type MethodSelectorsAttributes struct {
	ref terra.Reference
}

func (ms MethodSelectorsAttributes) InternalRef() (terra.Reference, error) {
	return ms.ref, nil
}

func (ms MethodSelectorsAttributes) InternalWithRef(ref terra.Reference) MethodSelectorsAttributes {
	return MethodSelectorsAttributes{ref: ref}
}

func (ms MethodSelectorsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ms.ref.InternalTokens()
}

func (ms MethodSelectorsAttributes) Method() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("method"))
}

func (ms MethodSelectorsAttributes) Permission() terra.StringValue {
	return terra.ReferenceAsString(ms.ref.Append("permission"))
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

type IngressFromState struct {
	Identities   []string       `json:"identities"`
	IdentityType string         `json:"identity_type"`
	Sources      []SourcesState `json:"sources"`
}

type SourcesState struct {
	AccessLevel string `json:"access_level"`
	Resource    string `json:"resource"`
}

type IngressToState struct {
	Resources  []string          `json:"resources"`
	Operations []OperationsState `json:"operations"`
}

type OperationsState struct {
	ServiceName     string                 `json:"service_name"`
	MethodSelectors []MethodSelectorsState `json:"method_selectors"`
}

type MethodSelectorsState struct {
	Method     string `json:"method"`
	Permission string `json:"permission"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}