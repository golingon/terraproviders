// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataidentitystoreuser

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Addresses struct{}

type Emails struct{}

type ExternalIds struct{}

type Name struct{}

type PhoneNumbers struct{}

type AlternateIdentifier struct {
	// ExternalId: optional
	ExternalId *ExternalId `hcl:"external_id,block"`
	// UniqueAttribute: optional
	UniqueAttribute *UniqueAttribute `hcl:"unique_attribute,block"`
}

type ExternalId struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
}

type UniqueAttribute struct {
	// AttributePath: string, required
	AttributePath terra.StringValue `hcl:"attribute_path,attr" validate:"required"`
	// AttributeValue: string, required
	AttributeValue terra.StringValue `hcl:"attribute_value,attr" validate:"required"`
}

type Filter struct {
	// AttributePath: string, required
	AttributePath terra.StringValue `hcl:"attribute_path,attr" validate:"required"`
	// AttributeValue: string, required
	AttributeValue terra.StringValue `hcl:"attribute_value,attr" validate:"required"`
}

type AddressesAttributes struct {
	ref terra.Reference
}

func (a AddressesAttributes) InternalRef() terra.Reference {
	return a.ref
}

func (a AddressesAttributes) InternalWithRef(ref terra.Reference) AddressesAttributes {
	return AddressesAttributes{ref: ref}
}

func (a AddressesAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AddressesAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("country"))
}

func (a AddressesAttributes) Formatted() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("formatted"))
}

func (a AddressesAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("locality"))
}

func (a AddressesAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("postal_code"))
}

func (a AddressesAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("primary"))
}

func (a AddressesAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("region"))
}

func (a AddressesAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("street_address"))
}

func (a AddressesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

type EmailsAttributes struct {
	ref terra.Reference
}

func (e EmailsAttributes) InternalRef() terra.Reference {
	return e.ref
}

func (e EmailsAttributes) InternalWithRef(ref terra.Reference) EmailsAttributes {
	return EmailsAttributes{ref: ref}
}

func (e EmailsAttributes) InternalTokens() hclwrite.Tokens {
	return e.ref.InternalTokens()
}

func (e EmailsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("primary"))
}

func (e EmailsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

func (e EmailsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("value"))
}

type ExternalIdsAttributes struct {
	ref terra.Reference
}

func (ei ExternalIdsAttributes) InternalRef() terra.Reference {
	return ei.ref
}

func (ei ExternalIdsAttributes) InternalWithRef(ref terra.Reference) ExternalIdsAttributes {
	return ExternalIdsAttributes{ref: ref}
}

func (ei ExternalIdsAttributes) InternalTokens() hclwrite.Tokens {
	return ei.ref.InternalTokens()
}

func (ei ExternalIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei ExternalIdsAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type NameAttributes struct {
	ref terra.Reference
}

func (n NameAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NameAttributes) InternalWithRef(ref terra.Reference) NameAttributes {
	return NameAttributes{ref: ref}
}

func (n NameAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NameAttributes) FamilyName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("family_name"))
}

func (n NameAttributes) Formatted() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("formatted"))
}

func (n NameAttributes) GivenName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("given_name"))
}

func (n NameAttributes) HonorificPrefix() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("honorific_prefix"))
}

func (n NameAttributes) HonorificSuffix() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("honorific_suffix"))
}

func (n NameAttributes) MiddleName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("middle_name"))
}

type PhoneNumbersAttributes struct {
	ref terra.Reference
}

func (pn PhoneNumbersAttributes) InternalRef() terra.Reference {
	return pn.ref
}

func (pn PhoneNumbersAttributes) InternalWithRef(ref terra.Reference) PhoneNumbersAttributes {
	return PhoneNumbersAttributes{ref: ref}
}

func (pn PhoneNumbersAttributes) InternalTokens() hclwrite.Tokens {
	return pn.ref.InternalTokens()
}

func (pn PhoneNumbersAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(pn.ref.Append("primary"))
}

func (pn PhoneNumbersAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pn.ref.Append("type"))
}

func (pn PhoneNumbersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pn.ref.Append("value"))
}

type AlternateIdentifierAttributes struct {
	ref terra.Reference
}

func (ai AlternateIdentifierAttributes) InternalRef() terra.Reference {
	return ai.ref
}

func (ai AlternateIdentifierAttributes) InternalWithRef(ref terra.Reference) AlternateIdentifierAttributes {
	return AlternateIdentifierAttributes{ref: ref}
}

func (ai AlternateIdentifierAttributes) InternalTokens() hclwrite.Tokens {
	return ai.ref.InternalTokens()
}

func (ai AlternateIdentifierAttributes) ExternalId() terra.ListValue[ExternalIdAttributes] {
	return terra.ReferenceAsList[ExternalIdAttributes](ai.ref.Append("external_id"))
}

func (ai AlternateIdentifierAttributes) UniqueAttribute() terra.ListValue[UniqueAttributeAttributes] {
	return terra.ReferenceAsList[UniqueAttributeAttributes](ai.ref.Append("unique_attribute"))
}

type ExternalIdAttributes struct {
	ref terra.Reference
}

func (ei ExternalIdAttributes) InternalRef() terra.Reference {
	return ei.ref
}

func (ei ExternalIdAttributes) InternalWithRef(ref terra.Reference) ExternalIdAttributes {
	return ExternalIdAttributes{ref: ref}
}

func (ei ExternalIdAttributes) InternalTokens() hclwrite.Tokens {
	return ei.ref.InternalTokens()
}

func (ei ExternalIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei ExternalIdAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type UniqueAttributeAttributes struct {
	ref terra.Reference
}

func (ua UniqueAttributeAttributes) InternalRef() terra.Reference {
	return ua.ref
}

func (ua UniqueAttributeAttributes) InternalWithRef(ref terra.Reference) UniqueAttributeAttributes {
	return UniqueAttributeAttributes{ref: ref}
}

func (ua UniqueAttributeAttributes) InternalTokens() hclwrite.Tokens {
	return ua.ref.InternalTokens()
}

func (ua UniqueAttributeAttributes) AttributePath() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_path"))
}

func (ua UniqueAttributeAttributes) AttributeValue() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_value"))
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) AttributePath() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("attribute_path"))
}

func (f FilterAttributes) AttributeValue() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("attribute_value"))
}

type AddressesState struct {
	Country       string `json:"country"`
	Formatted     string `json:"formatted"`
	Locality      string `json:"locality"`
	PostalCode    string `json:"postal_code"`
	Primary       bool   `json:"primary"`
	Region        string `json:"region"`
	StreetAddress string `json:"street_address"`
	Type          string `json:"type"`
}

type EmailsState struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type ExternalIdsState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type NameState struct {
	FamilyName      string `json:"family_name"`
	Formatted       string `json:"formatted"`
	GivenName       string `json:"given_name"`
	HonorificPrefix string `json:"honorific_prefix"`
	HonorificSuffix string `json:"honorific_suffix"`
	MiddleName      string `json:"middle_name"`
}

type PhoneNumbersState struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type AlternateIdentifierState struct {
	ExternalId      []ExternalIdState      `json:"external_id"`
	UniqueAttribute []UniqueAttributeState `json:"unique_attribute"`
}

type ExternalIdState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type UniqueAttributeState struct {
	AttributePath  string `json:"attribute_path"`
	AttributeValue string `json:"attribute_value"`
}

type FilterState struct {
	AttributePath  string `json:"attribute_path"`
	AttributeValue string `json:"attribute_value"`
}
