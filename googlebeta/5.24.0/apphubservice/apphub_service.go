// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package apphubservice

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ServiceProperties struct{}

type ServiceReference struct{}

type Attributes struct {
	// BusinessOwners: min=0
	BusinessOwners []BusinessOwners `hcl:"business_owners,block" validate:"min=0"`
	// Criticality: optional
	Criticality *Criticality `hcl:"criticality,block"`
	// DeveloperOwners: min=0
	DeveloperOwners []DeveloperOwners `hcl:"developer_owners,block" validate:"min=0"`
	// Environment: optional
	Environment *Environment `hcl:"environment,block"`
	// OperatorOwners: min=0
	OperatorOwners []OperatorOwners `hcl:"operator_owners,block" validate:"min=0"`
}

type BusinessOwners struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
}

type Criticality struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type DeveloperOwners struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
}

type Environment struct {
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type OperatorOwners struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Email: string, required
	Email terra.StringValue `hcl:"email,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type ServicePropertiesAttributes struct {
	ref terra.Reference
}

func (sp ServicePropertiesAttributes) InternalRef() (terra.Reference, error) {
	return sp.ref, nil
}

func (sp ServicePropertiesAttributes) InternalWithRef(ref terra.Reference) ServicePropertiesAttributes {
	return ServicePropertiesAttributes{ref: ref}
}

func (sp ServicePropertiesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sp.ref.InternalTokens()
}

func (sp ServicePropertiesAttributes) GcpProject() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("gcp_project"))
}

func (sp ServicePropertiesAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("location"))
}

func (sp ServicePropertiesAttributes) Zone() terra.StringValue {
	return terra.ReferenceAsString(sp.ref.Append("zone"))
}

type ServiceReferenceAttributes struct {
	ref terra.Reference
}

func (sr ServiceReferenceAttributes) InternalRef() (terra.Reference, error) {
	return sr.ref, nil
}

func (sr ServiceReferenceAttributes) InternalWithRef(ref terra.Reference) ServiceReferenceAttributes {
	return ServiceReferenceAttributes{ref: ref}
}

func (sr ServiceReferenceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sr.ref.InternalTokens()
}

func (sr ServiceReferenceAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(sr.ref.Append("uri"))
}

type AttributesAttributes struct {
	ref terra.Reference
}

func (a AttributesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AttributesAttributes) InternalWithRef(ref terra.Reference) AttributesAttributes {
	return AttributesAttributes{ref: ref}
}

func (a AttributesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AttributesAttributes) BusinessOwners() terra.ListValue[BusinessOwnersAttributes] {
	return terra.ReferenceAsList[BusinessOwnersAttributes](a.ref.Append("business_owners"))
}

func (a AttributesAttributes) Criticality() terra.ListValue[CriticalityAttributes] {
	return terra.ReferenceAsList[CriticalityAttributes](a.ref.Append("criticality"))
}

func (a AttributesAttributes) DeveloperOwners() terra.ListValue[DeveloperOwnersAttributes] {
	return terra.ReferenceAsList[DeveloperOwnersAttributes](a.ref.Append("developer_owners"))
}

func (a AttributesAttributes) Environment() terra.ListValue[EnvironmentAttributes] {
	return terra.ReferenceAsList[EnvironmentAttributes](a.ref.Append("environment"))
}

func (a AttributesAttributes) OperatorOwners() terra.ListValue[OperatorOwnersAttributes] {
	return terra.ReferenceAsList[OperatorOwnersAttributes](a.ref.Append("operator_owners"))
}

type BusinessOwnersAttributes struct {
	ref terra.Reference
}

func (bo BusinessOwnersAttributes) InternalRef() (terra.Reference, error) {
	return bo.ref, nil
}

func (bo BusinessOwnersAttributes) InternalWithRef(ref terra.Reference) BusinessOwnersAttributes {
	return BusinessOwnersAttributes{ref: ref}
}

func (bo BusinessOwnersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bo.ref.InternalTokens()
}

func (bo BusinessOwnersAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(bo.ref.Append("display_name"))
}

func (bo BusinessOwnersAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(bo.ref.Append("email"))
}

type CriticalityAttributes struct {
	ref terra.Reference
}

func (c CriticalityAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CriticalityAttributes) InternalWithRef(ref terra.Reference) CriticalityAttributes {
	return CriticalityAttributes{ref: ref}
}

func (c CriticalityAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CriticalityAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("type"))
}

type DeveloperOwnersAttributes struct {
	ref terra.Reference
}

func (do DeveloperOwnersAttributes) InternalRef() (terra.Reference, error) {
	return do.ref, nil
}

func (do DeveloperOwnersAttributes) InternalWithRef(ref terra.Reference) DeveloperOwnersAttributes {
	return DeveloperOwnersAttributes{ref: ref}
}

func (do DeveloperOwnersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return do.ref.InternalTokens()
}

func (do DeveloperOwnersAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("display_name"))
}

func (do DeveloperOwnersAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(do.ref.Append("email"))
}

type EnvironmentAttributes struct {
	ref terra.Reference
}

func (e EnvironmentAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EnvironmentAttributes) InternalWithRef(ref terra.Reference) EnvironmentAttributes {
	return EnvironmentAttributes{ref: ref}
}

func (e EnvironmentAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EnvironmentAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

type OperatorOwnersAttributes struct {
	ref terra.Reference
}

func (oo OperatorOwnersAttributes) InternalRef() (terra.Reference, error) {
	return oo.ref, nil
}

func (oo OperatorOwnersAttributes) InternalWithRef(ref terra.Reference) OperatorOwnersAttributes {
	return OperatorOwnersAttributes{ref: ref}
}

func (oo OperatorOwnersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oo.ref.InternalTokens()
}

func (oo OperatorOwnersAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("display_name"))
}

func (oo OperatorOwnersAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(oo.ref.Append("email"))
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

type ServicePropertiesState struct {
	GcpProject string `json:"gcp_project"`
	Location   string `json:"location"`
	Zone       string `json:"zone"`
}

type ServiceReferenceState struct {
	Uri string `json:"uri"`
}

type AttributesState struct {
	BusinessOwners  []BusinessOwnersState  `json:"business_owners"`
	Criticality     []CriticalityState     `json:"criticality"`
	DeveloperOwners []DeveloperOwnersState `json:"developer_owners"`
	Environment     []EnvironmentState     `json:"environment"`
	OperatorOwners  []OperatorOwnersState  `json:"operator_owners"`
}

type BusinessOwnersState struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type CriticalityState struct {
	Type string `json:"type"`
}

type DeveloperOwnersState struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type EnvironmentState struct {
	Type string `json:"type"`
}

type OperatorOwnersState struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
