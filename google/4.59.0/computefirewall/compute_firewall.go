// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package computefirewall

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Allow struct {
	// Ports: list of string, optional
	Ports terra.ListValue[terra.StringValue] `hcl:"ports,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
}

type Deny struct {
	// Ports: list of string, optional
	Ports terra.ListValue[terra.StringValue] `hcl:"ports,attr"`
	// Protocol: string, required
	Protocol terra.StringValue `hcl:"protocol,attr" validate:"required"`
}

type LogConfig struct {
	// Metadata: string, required
	Metadata terra.StringValue `hcl:"metadata,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AllowAttributes struct {
	ref terra.Reference
}

func (a AllowAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AllowAttributes) InternalWithRef(ref terra.Reference) AllowAttributes {
	return AllowAttributes{ref: ref}
}

func (a AllowAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AllowAttributes) Ports() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](a.ref.Append("ports"))
}

func (a AllowAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(a.ref.Append("protocol"))
}

type DenyAttributes struct {
	ref terra.Reference
}

func (d DenyAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d DenyAttributes) InternalWithRef(ref terra.Reference) DenyAttributes {
	return DenyAttributes{ref: ref}
}

func (d DenyAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DenyAttributes) Ports() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](d.ref.Append("ports"))
}

func (d DenyAttributes) Protocol() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("protocol"))
}

type LogConfigAttributes struct {
	ref terra.Reference
}

func (lc LogConfigAttributes) InternalRef() terra.Reference {
	return lc.ref
}

func (lc LogConfigAttributes) InternalWithRef(ref terra.Reference) LogConfigAttributes {
	return LogConfigAttributes{ref: ref}
}

func (lc LogConfigAttributes) InternalTokens() hclwrite.Tokens {
	return lc.ref.InternalTokens()
}

func (lc LogConfigAttributes) Metadata() terra.StringValue {
	return terra.ReferenceString(lc.ref.Append("metadata"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type AllowState struct {
	Ports    []string `json:"ports"`
	Protocol string   `json:"protocol"`
}

type DenyState struct {
	Ports    []string `json:"ports"`
	Protocol string   `json:"protocol"`
}

type LogConfigState struct {
	Metadata string `json:"metadata"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
