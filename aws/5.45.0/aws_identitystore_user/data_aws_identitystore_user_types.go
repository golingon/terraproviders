// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_identitystore_user

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataAlternateIdentifier struct {
	// AlternateIdentifierExternalId: optional
	ExternalId *DataAlternateIdentifierExternalId `hcl:"external_id,block"`
	// AlternateIdentifierUniqueAttribute: optional
	UniqueAttribute *DataAlternateIdentifierUniqueAttribute `hcl:"unique_attribute,block"`
}

type DataAlternateIdentifierExternalId struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Issuer: string, required
	Issuer terra.StringValue `hcl:"issuer,attr" validate:"required"`
}

type DataAlternateIdentifierUniqueAttribute struct {
	// AttributePath: string, required
	AttributePath terra.StringValue `hcl:"attribute_path,attr" validate:"required"`
	// AttributeValue: string, required
	AttributeValue terra.StringValue `hcl:"attribute_value,attr" validate:"required"`
}

type DataFilter struct {
	// AttributePath: string, required
	AttributePath terra.StringValue `hcl:"attribute_path,attr" validate:"required"`
	// AttributeValue: string, required
	AttributeValue terra.StringValue `hcl:"attribute_value,attr" validate:"required"`
}

type DataAddressesAttributes struct {
	ref terra.Reference
}

func (a DataAddressesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a DataAddressesAttributes) InternalWithRef(ref terra.Reference) DataAddressesAttributes {
	return DataAddressesAttributes{ref: ref}
}

func (a DataAddressesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a DataAddressesAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("country"))
}

func (a DataAddressesAttributes) Formatted() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("formatted"))
}

func (a DataAddressesAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("locality"))
}

func (a DataAddressesAttributes) PostalCode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("postal_code"))
}

func (a DataAddressesAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("primary"))
}

func (a DataAddressesAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("region"))
}

func (a DataAddressesAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("street_address"))
}

func (a DataAddressesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

type DataEmailsAttributes struct {
	ref terra.Reference
}

func (e DataEmailsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e DataEmailsAttributes) InternalWithRef(ref terra.Reference) DataEmailsAttributes {
	return DataEmailsAttributes{ref: ref}
}

func (e DataEmailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e DataEmailsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("primary"))
}

func (e DataEmailsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

func (e DataEmailsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("value"))
}

type DataExternalIdsAttributes struct {
	ref terra.Reference
}

func (ei DataExternalIdsAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei DataExternalIdsAttributes) InternalWithRef(ref terra.Reference) DataExternalIdsAttributes {
	return DataExternalIdsAttributes{ref: ref}
}

func (ei DataExternalIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei DataExternalIdsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei DataExternalIdsAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type DataNameAttributes struct {
	ref terra.Reference
}

func (n DataNameAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n DataNameAttributes) InternalWithRef(ref terra.Reference) DataNameAttributes {
	return DataNameAttributes{ref: ref}
}

func (n DataNameAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n DataNameAttributes) FamilyName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("family_name"))
}

func (n DataNameAttributes) Formatted() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("formatted"))
}

func (n DataNameAttributes) GivenName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("given_name"))
}

func (n DataNameAttributes) HonorificPrefix() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("honorific_prefix"))
}

func (n DataNameAttributes) HonorificSuffix() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("honorific_suffix"))
}

func (n DataNameAttributes) MiddleName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("middle_name"))
}

type DataPhoneNumbersAttributes struct {
	ref terra.Reference
}

func (pn DataPhoneNumbersAttributes) InternalRef() (terra.Reference, error) {
	return pn.ref, nil
}

func (pn DataPhoneNumbersAttributes) InternalWithRef(ref terra.Reference) DataPhoneNumbersAttributes {
	return DataPhoneNumbersAttributes{ref: ref}
}

func (pn DataPhoneNumbersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pn.ref.InternalTokens()
}

func (pn DataPhoneNumbersAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(pn.ref.Append("primary"))
}

func (pn DataPhoneNumbersAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(pn.ref.Append("type"))
}

func (pn DataPhoneNumbersAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(pn.ref.Append("value"))
}

type DataAlternateIdentifierAttributes struct {
	ref terra.Reference
}

func (ai DataAlternateIdentifierAttributes) InternalRef() (terra.Reference, error) {
	return ai.ref, nil
}

func (ai DataAlternateIdentifierAttributes) InternalWithRef(ref terra.Reference) DataAlternateIdentifierAttributes {
	return DataAlternateIdentifierAttributes{ref: ref}
}

func (ai DataAlternateIdentifierAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ai.ref.InternalTokens()
}

func (ai DataAlternateIdentifierAttributes) ExternalId() terra.ListValue[DataAlternateIdentifierExternalIdAttributes] {
	return terra.ReferenceAsList[DataAlternateIdentifierExternalIdAttributes](ai.ref.Append("external_id"))
}

func (ai DataAlternateIdentifierAttributes) UniqueAttribute() terra.ListValue[DataAlternateIdentifierUniqueAttributeAttributes] {
	return terra.ReferenceAsList[DataAlternateIdentifierUniqueAttributeAttributes](ai.ref.Append("unique_attribute"))
}

type DataAlternateIdentifierExternalIdAttributes struct {
	ref terra.Reference
}

func (ei DataAlternateIdentifierExternalIdAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei DataAlternateIdentifierExternalIdAttributes) InternalWithRef(ref terra.Reference) DataAlternateIdentifierExternalIdAttributes {
	return DataAlternateIdentifierExternalIdAttributes{ref: ref}
}

func (ei DataAlternateIdentifierExternalIdAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei DataAlternateIdentifierExternalIdAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("id"))
}

func (ei DataAlternateIdentifierExternalIdAttributes) Issuer() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("issuer"))
}

type DataAlternateIdentifierUniqueAttributeAttributes struct {
	ref terra.Reference
}

func (ua DataAlternateIdentifierUniqueAttributeAttributes) InternalRef() (terra.Reference, error) {
	return ua.ref, nil
}

func (ua DataAlternateIdentifierUniqueAttributeAttributes) InternalWithRef(ref terra.Reference) DataAlternateIdentifierUniqueAttributeAttributes {
	return DataAlternateIdentifierUniqueAttributeAttributes{ref: ref}
}

func (ua DataAlternateIdentifierUniqueAttributeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ua.ref.InternalTokens()
}

func (ua DataAlternateIdentifierUniqueAttributeAttributes) AttributePath() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_path"))
}

func (ua DataAlternateIdentifierUniqueAttributeAttributes) AttributeValue() terra.StringValue {
	return terra.ReferenceAsString(ua.ref.Append("attribute_value"))
}

type DataFilterAttributes struct {
	ref terra.Reference
}

func (f DataFilterAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFilterAttributes) InternalWithRef(ref terra.Reference) DataFilterAttributes {
	return DataFilterAttributes{ref: ref}
}

func (f DataFilterAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFilterAttributes) AttributePath() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("attribute_path"))
}

func (f DataFilterAttributes) AttributeValue() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("attribute_value"))
}

type DataAddressesState struct {
	Country       string `json:"country"`
	Formatted     string `json:"formatted"`
	Locality      string `json:"locality"`
	PostalCode    string `json:"postal_code"`
	Primary       bool   `json:"primary"`
	Region        string `json:"region"`
	StreetAddress string `json:"street_address"`
	Type          string `json:"type"`
}

type DataEmailsState struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type DataExternalIdsState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type DataNameState struct {
	FamilyName      string `json:"family_name"`
	Formatted       string `json:"formatted"`
	GivenName       string `json:"given_name"`
	HonorificPrefix string `json:"honorific_prefix"`
	HonorificSuffix string `json:"honorific_suffix"`
	MiddleName      string `json:"middle_name"`
}

type DataPhoneNumbersState struct {
	Primary bool   `json:"primary"`
	Type    string `json:"type"`
	Value   string `json:"value"`
}

type DataAlternateIdentifierState struct {
	ExternalId      []DataAlternateIdentifierExternalIdState      `json:"external_id"`
	UniqueAttribute []DataAlternateIdentifierUniqueAttributeState `json:"unique_attribute"`
}

type DataAlternateIdentifierExternalIdState struct {
	Id     string `json:"id"`
	Issuer string `json:"issuer"`
}

type DataAlternateIdentifierUniqueAttributeState struct {
	AttributePath  string `json:"attribute_path"`
	AttributeValue string `json:"attribute_value"`
}

type DataFilterState struct {
	AttributePath  string `json:"attribute_path"`
	AttributeValue string `json:"attribute_value"`
}
