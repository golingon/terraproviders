// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gkehubfeature

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ResourceState struct{}

type State struct {
	// StateState: min=0
	State []StateState `hcl:"state,block" validate:"min=0"`
}

type StateState struct{}

type Spec struct {
	// Multiclusteringress: optional
	Multiclusteringress *Multiclusteringress `hcl:"multiclusteringress,block"`
}

type Multiclusteringress struct {
	// ConfigMembership: string, required
	ConfigMembership terra.StringValue `hcl:"config_membership,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ResourceStateAttributes struct {
	ref terra.Reference
}

func (rs ResourceStateAttributes) InternalRef() (terra.Reference, error) {
	return rs.ref, nil
}

func (rs ResourceStateAttributes) InternalWithRef(ref terra.Reference) ResourceStateAttributes {
	return ResourceStateAttributes{ref: ref}
}

func (rs ResourceStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rs.ref.InternalTokens()
}

func (rs ResourceStateAttributes) HasResources() terra.BoolValue {
	return terra.ReferenceAsBool(rs.ref.Append("has_resources"))
}

func (rs ResourceStateAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(rs.ref.Append("state"))
}

type StateAttributes struct {
	ref terra.Reference
}

func (s StateAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StateAttributes) InternalWithRef(ref terra.Reference) StateAttributes {
	return StateAttributes{ref: ref}
}

func (s StateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StateAttributes) State() terra.ListValue[StateStateAttributes] {
	return terra.ReferenceAsList[StateStateAttributes](s.ref.Append("state"))
}

type StateStateAttributes struct {
	ref terra.Reference
}

func (s StateStateAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s StateStateAttributes) InternalWithRef(ref terra.Reference) StateStateAttributes {
	return StateStateAttributes{ref: ref}
}

func (s StateStateAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s StateStateAttributes) Code() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("code"))
}

func (s StateStateAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("description"))
}

func (s StateStateAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("update_time"))
}

type SpecAttributes struct {
	ref terra.Reference
}

func (s SpecAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SpecAttributes) InternalWithRef(ref terra.Reference) SpecAttributes {
	return SpecAttributes{ref: ref}
}

func (s SpecAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SpecAttributes) Multiclusteringress() terra.ListValue[MulticlusteringressAttributes] {
	return terra.ReferenceAsList[MulticlusteringressAttributes](s.ref.Append("multiclusteringress"))
}

type MulticlusteringressAttributes struct {
	ref terra.Reference
}

func (m MulticlusteringressAttributes) InternalRef() (terra.Reference, error) {
	return m.ref, nil
}

func (m MulticlusteringressAttributes) InternalWithRef(ref terra.Reference) MulticlusteringressAttributes {
	return MulticlusteringressAttributes{ref: ref}
}

func (m MulticlusteringressAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return m.ref.InternalTokens()
}

func (m MulticlusteringressAttributes) ConfigMembership() terra.StringValue {
	return terra.ReferenceAsString(m.ref.Append("config_membership"))
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

type ResourceStateState struct {
	HasResources bool   `json:"has_resources"`
	State        string `json:"state"`
}

type StateState struct {
	State []StateStateState `json:"state"`
}

type StateStateState struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	UpdateTime  string `json:"update_time"`
}

type SpecState struct {
	Multiclusteringress []MulticlusteringressState `json:"multiclusteringress"`
}

type MulticlusteringressState struct {
	ConfigMembership string `json:"config_membership"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
