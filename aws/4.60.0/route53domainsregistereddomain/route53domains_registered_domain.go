// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package route53domainsregistereddomain

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AdminContact struct {
	// AddressLine1: string, optional
	AddressLine1 terra.StringValue `hcl:"address_line_1,attr"`
	// AddressLine2: string, optional
	AddressLine2 terra.StringValue `hcl:"address_line_2,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// ContactType: string, optional
	ContactType terra.StringValue `hcl:"contact_type,attr"`
	// CountryCode: string, optional
	CountryCode terra.StringValue `hcl:"country_code,attr"`
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// ExtraParams: map of string, optional
	ExtraParams terra.MapValue[terra.StringValue] `hcl:"extra_params,attr"`
	// Fax: string, optional
	Fax terra.StringValue `hcl:"fax,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// OrganizationName: string, optional
	OrganizationName terra.StringValue `hcl:"organization_name,attr"`
	// PhoneNumber: string, optional
	PhoneNumber terra.StringValue `hcl:"phone_number,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// ZipCode: string, optional
	ZipCode terra.StringValue `hcl:"zip_code,attr"`
}

type NameServer struct {
	// GlueIps: set of string, optional
	GlueIps terra.SetValue[terra.StringValue] `hcl:"glue_ips,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type RegistrantContact struct {
	// AddressLine1: string, optional
	AddressLine1 terra.StringValue `hcl:"address_line_1,attr"`
	// AddressLine2: string, optional
	AddressLine2 terra.StringValue `hcl:"address_line_2,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// ContactType: string, optional
	ContactType terra.StringValue `hcl:"contact_type,attr"`
	// CountryCode: string, optional
	CountryCode terra.StringValue `hcl:"country_code,attr"`
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// ExtraParams: map of string, optional
	ExtraParams terra.MapValue[terra.StringValue] `hcl:"extra_params,attr"`
	// Fax: string, optional
	Fax terra.StringValue `hcl:"fax,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// OrganizationName: string, optional
	OrganizationName terra.StringValue `hcl:"organization_name,attr"`
	// PhoneNumber: string, optional
	PhoneNumber terra.StringValue `hcl:"phone_number,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// ZipCode: string, optional
	ZipCode terra.StringValue `hcl:"zip_code,attr"`
}

type TechContact struct {
	// AddressLine1: string, optional
	AddressLine1 terra.StringValue `hcl:"address_line_1,attr"`
	// AddressLine2: string, optional
	AddressLine2 terra.StringValue `hcl:"address_line_2,attr"`
	// City: string, optional
	City terra.StringValue `hcl:"city,attr"`
	// ContactType: string, optional
	ContactType terra.StringValue `hcl:"contact_type,attr"`
	// CountryCode: string, optional
	CountryCode terra.StringValue `hcl:"country_code,attr"`
	// Email: string, optional
	Email terra.StringValue `hcl:"email,attr"`
	// ExtraParams: map of string, optional
	ExtraParams terra.MapValue[terra.StringValue] `hcl:"extra_params,attr"`
	// Fax: string, optional
	Fax terra.StringValue `hcl:"fax,attr"`
	// FirstName: string, optional
	FirstName terra.StringValue `hcl:"first_name,attr"`
	// LastName: string, optional
	LastName terra.StringValue `hcl:"last_name,attr"`
	// OrganizationName: string, optional
	OrganizationName terra.StringValue `hcl:"organization_name,attr"`
	// PhoneNumber: string, optional
	PhoneNumber terra.StringValue `hcl:"phone_number,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// ZipCode: string, optional
	ZipCode terra.StringValue `hcl:"zip_code,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AdminContactAttributes struct {
	ref terra.Reference
}

func (ac AdminContactAttributes) InternalRef() terra.Reference {
	return ac.ref
}

func (ac AdminContactAttributes) InternalWithRef(ref terra.Reference) AdminContactAttributes {
	return AdminContactAttributes{ref: ref}
}

func (ac AdminContactAttributes) InternalTokens() hclwrite.Tokens {
	return ac.ref.InternalTokens()
}

func (ac AdminContactAttributes) AddressLine1() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("address_line_1"))
}

func (ac AdminContactAttributes) AddressLine2() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("address_line_2"))
}

func (ac AdminContactAttributes) City() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("city"))
}

func (ac AdminContactAttributes) ContactType() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("contact_type"))
}

func (ac AdminContactAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("country_code"))
}

func (ac AdminContactAttributes) Email() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("email"))
}

func (ac AdminContactAttributes) ExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ac.ref.Append("extra_params"))
}

func (ac AdminContactAttributes) Fax() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("fax"))
}

func (ac AdminContactAttributes) FirstName() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("first_name"))
}

func (ac AdminContactAttributes) LastName() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("last_name"))
}

func (ac AdminContactAttributes) OrganizationName() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("organization_name"))
}

func (ac AdminContactAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("phone_number"))
}

func (ac AdminContactAttributes) State() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("state"))
}

func (ac AdminContactAttributes) ZipCode() terra.StringValue {
	return terra.ReferenceString(ac.ref.Append("zip_code"))
}

type NameServerAttributes struct {
	ref terra.Reference
}

func (ns NameServerAttributes) InternalRef() terra.Reference {
	return ns.ref
}

func (ns NameServerAttributes) InternalWithRef(ref terra.Reference) NameServerAttributes {
	return NameServerAttributes{ref: ref}
}

func (ns NameServerAttributes) InternalTokens() hclwrite.Tokens {
	return ns.ref.InternalTokens()
}

func (ns NameServerAttributes) GlueIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](ns.ref.Append("glue_ips"))
}

func (ns NameServerAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ns.ref.Append("name"))
}

type RegistrantContactAttributes struct {
	ref terra.Reference
}

func (rc RegistrantContactAttributes) InternalRef() terra.Reference {
	return rc.ref
}

func (rc RegistrantContactAttributes) InternalWithRef(ref terra.Reference) RegistrantContactAttributes {
	return RegistrantContactAttributes{ref: ref}
}

func (rc RegistrantContactAttributes) InternalTokens() hclwrite.Tokens {
	return rc.ref.InternalTokens()
}

func (rc RegistrantContactAttributes) AddressLine1() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("address_line_1"))
}

func (rc RegistrantContactAttributes) AddressLine2() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("address_line_2"))
}

func (rc RegistrantContactAttributes) City() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("city"))
}

func (rc RegistrantContactAttributes) ContactType() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("contact_type"))
}

func (rc RegistrantContactAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("country_code"))
}

func (rc RegistrantContactAttributes) Email() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("email"))
}

func (rc RegistrantContactAttributes) ExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](rc.ref.Append("extra_params"))
}

func (rc RegistrantContactAttributes) Fax() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("fax"))
}

func (rc RegistrantContactAttributes) FirstName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("first_name"))
}

func (rc RegistrantContactAttributes) LastName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("last_name"))
}

func (rc RegistrantContactAttributes) OrganizationName() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("organization_name"))
}

func (rc RegistrantContactAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("phone_number"))
}

func (rc RegistrantContactAttributes) State() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("state"))
}

func (rc RegistrantContactAttributes) ZipCode() terra.StringValue {
	return terra.ReferenceString(rc.ref.Append("zip_code"))
}

type TechContactAttributes struct {
	ref terra.Reference
}

func (tc TechContactAttributes) InternalRef() terra.Reference {
	return tc.ref
}

func (tc TechContactAttributes) InternalWithRef(ref terra.Reference) TechContactAttributes {
	return TechContactAttributes{ref: ref}
}

func (tc TechContactAttributes) InternalTokens() hclwrite.Tokens {
	return tc.ref.InternalTokens()
}

func (tc TechContactAttributes) AddressLine1() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("address_line_1"))
}

func (tc TechContactAttributes) AddressLine2() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("address_line_2"))
}

func (tc TechContactAttributes) City() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("city"))
}

func (tc TechContactAttributes) ContactType() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("contact_type"))
}

func (tc TechContactAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("country_code"))
}

func (tc TechContactAttributes) Email() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("email"))
}

func (tc TechContactAttributes) ExtraParams() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](tc.ref.Append("extra_params"))
}

func (tc TechContactAttributes) Fax() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("fax"))
}

func (tc TechContactAttributes) FirstName() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("first_name"))
}

func (tc TechContactAttributes) LastName() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("last_name"))
}

func (tc TechContactAttributes) OrganizationName() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("organization_name"))
}

func (tc TechContactAttributes) PhoneNumber() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("phone_number"))
}

func (tc TechContactAttributes) State() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("state"))
}

func (tc TechContactAttributes) ZipCode() terra.StringValue {
	return terra.ReferenceString(tc.ref.Append("zip_code"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type AdminContactState struct {
	AddressLine1     string            `json:"address_line_1"`
	AddressLine2     string            `json:"address_line_2"`
	City             string            `json:"city"`
	ContactType      string            `json:"contact_type"`
	CountryCode      string            `json:"country_code"`
	Email            string            `json:"email"`
	ExtraParams      map[string]string `json:"extra_params"`
	Fax              string            `json:"fax"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	OrganizationName string            `json:"organization_name"`
	PhoneNumber      string            `json:"phone_number"`
	State            string            `json:"state"`
	ZipCode          string            `json:"zip_code"`
}

type NameServerState struct {
	GlueIps []string `json:"glue_ips"`
	Name    string   `json:"name"`
}

type RegistrantContactState struct {
	AddressLine1     string            `json:"address_line_1"`
	AddressLine2     string            `json:"address_line_2"`
	City             string            `json:"city"`
	ContactType      string            `json:"contact_type"`
	CountryCode      string            `json:"country_code"`
	Email            string            `json:"email"`
	ExtraParams      map[string]string `json:"extra_params"`
	Fax              string            `json:"fax"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	OrganizationName string            `json:"organization_name"`
	PhoneNumber      string            `json:"phone_number"`
	State            string            `json:"state"`
	ZipCode          string            `json:"zip_code"`
}

type TechContactState struct {
	AddressLine1     string            `json:"address_line_1"`
	AddressLine2     string            `json:"address_line_2"`
	City             string            `json:"city"`
	ContactType      string            `json:"contact_type"`
	CountryCode      string            `json:"country_code"`
	Email            string            `json:"email"`
	ExtraParams      map[string]string `json:"extra_params"`
	Fax              string            `json:"fax"`
	FirstName        string            `json:"first_name"`
	LastName         string            `json:"last_name"`
	OrganizationName string            `json:"organization_name"`
	PhoneNumber      string            `json:"phone_number"`
	State            string            `json:"state"`
	ZipCode          string            `json:"zip_code"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Update string `json:"update"`
}
