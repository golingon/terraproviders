// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package fmspolicy

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ExcludeMap struct {
	// Account: set of string, optional
	Account terra.SetValue[terra.StringValue] `hcl:"account,attr"`
	// Orgunit: set of string, optional
	Orgunit terra.SetValue[terra.StringValue] `hcl:"orgunit,attr"`
}

type IncludeMap struct {
	// Account: set of string, optional
	Account terra.SetValue[terra.StringValue] `hcl:"account,attr"`
	// Orgunit: set of string, optional
	Orgunit terra.SetValue[terra.StringValue] `hcl:"orgunit,attr"`
}

type SecurityServicePolicyData struct {
	// ManagedServiceData: string, optional
	ManagedServiceData terra.StringValue `hcl:"managed_service_data,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type ExcludeMapAttributes struct {
	ref terra.Reference
}

func (em ExcludeMapAttributes) InternalRef() (terra.Reference, error) {
	return em.ref, nil
}

func (em ExcludeMapAttributes) InternalWithRef(ref terra.Reference) ExcludeMapAttributes {
	return ExcludeMapAttributes{ref: ref}
}

func (em ExcludeMapAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return em.ref.InternalTokens()
}

func (em ExcludeMapAttributes) Account() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](em.ref.Append("account"))
}

func (em ExcludeMapAttributes) Orgunit() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](em.ref.Append("orgunit"))
}

type IncludeMapAttributes struct {
	ref terra.Reference
}

func (im IncludeMapAttributes) InternalRef() (terra.Reference, error) {
	return im.ref, nil
}

func (im IncludeMapAttributes) InternalWithRef(ref terra.Reference) IncludeMapAttributes {
	return IncludeMapAttributes{ref: ref}
}

func (im IncludeMapAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return im.ref.InternalTokens()
}

func (im IncludeMapAttributes) Account() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](im.ref.Append("account"))
}

func (im IncludeMapAttributes) Orgunit() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](im.ref.Append("orgunit"))
}

type SecurityServicePolicyDataAttributes struct {
	ref terra.Reference
}

func (sspd SecurityServicePolicyDataAttributes) InternalRef() (terra.Reference, error) {
	return sspd.ref, nil
}

func (sspd SecurityServicePolicyDataAttributes) InternalWithRef(ref terra.Reference) SecurityServicePolicyDataAttributes {
	return SecurityServicePolicyDataAttributes{ref: ref}
}

func (sspd SecurityServicePolicyDataAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sspd.ref.InternalTokens()
}

func (sspd SecurityServicePolicyDataAttributes) ManagedServiceData() terra.StringValue {
	return terra.ReferenceAsString(sspd.ref.Append("managed_service_data"))
}

func (sspd SecurityServicePolicyDataAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(sspd.ref.Append("type"))
}

type ExcludeMapState struct {
	Account []string `json:"account"`
	Orgunit []string `json:"orgunit"`
}

type IncludeMapState struct {
	Account []string `json:"account"`
	Orgunit []string `json:"orgunit"`
}

type SecurityServicePolicyDataState struct {
	ManagedServiceData string `json:"managed_service_data"`
	Type               string `json:"type"`
}
