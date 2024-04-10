// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datauser

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Addresses struct{}

type CustomSchemas struct{}

type Emails struct{}

type ExternalIds struct{}

type Ims struct{}

type Keywords struct{}

type Languages struct{}

type Locations struct{}

type Name struct{}

type Organizations struct{}

type Phones struct{}

type PosixAccounts struct{}

type Relations struct{}

type SshPublicKeys struct{}

type Websites struct{}

type AddressesAttributes struct {
	ref terra.Reference
}

func (a AddressesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AddressesAttributes) InternalWithRef(ref terra.Reference) AddressesAttributes {
	return AddressesAttributes{ref: ref}
}

func (a AddressesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AddressesAttributes) Country() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("country"))
}

func (a AddressesAttributes) CountryCode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("country_code"))
}

func (a AddressesAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("custom_type"))
}

func (a AddressesAttributes) ExtendedAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("extended_address"))
}

func (a AddressesAttributes) Formatted() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("formatted"))
}

func (a AddressesAttributes) Locality() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("locality"))
}

func (a AddressesAttributes) PoBox() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("po_box"))
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

func (a AddressesAttributes) SourceIsStructured() terra.BoolValue {
	return terra.ReferenceAsBool(a.ref.Append("source_is_structured"))
}

func (a AddressesAttributes) StreetAddress() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("street_address"))
}

func (a AddressesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("type"))
}

type CustomSchemasAttributes struct {
	ref terra.Reference
}

func (cs CustomSchemasAttributes) InternalRef() (terra.Reference, error) {
	return cs.ref, nil
}

func (cs CustomSchemasAttributes) InternalWithRef(ref terra.Reference) CustomSchemasAttributes {
	return CustomSchemasAttributes{ref: ref}
}

func (cs CustomSchemasAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cs.ref.InternalTokens()
}

func (cs CustomSchemasAttributes) SchemaName() terra.StringValue {
	return terra.ReferenceAsString(cs.ref.Append("schema_name"))
}

func (cs CustomSchemasAttributes) SchemaValues() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cs.ref.Append("schema_values"))
}

type EmailsAttributes struct {
	ref terra.Reference
}

func (e EmailsAttributes) InternalRef() (terra.Reference, error) {
	return e.ref, nil
}

func (e EmailsAttributes) InternalWithRef(ref terra.Reference) EmailsAttributes {
	return EmailsAttributes{ref: ref}
}

func (e EmailsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return e.ref.InternalTokens()
}

func (e EmailsAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("address"))
}

func (e EmailsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("custom_type"))
}

func (e EmailsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(e.ref.Append("primary"))
}

func (e EmailsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(e.ref.Append("type"))
}

type ExternalIdsAttributes struct {
	ref terra.Reference
}

func (ei ExternalIdsAttributes) InternalRef() (terra.Reference, error) {
	return ei.ref, nil
}

func (ei ExternalIdsAttributes) InternalWithRef(ref terra.Reference) ExternalIdsAttributes {
	return ExternalIdsAttributes{ref: ref}
}

func (ei ExternalIdsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ei.ref.InternalTokens()
}

func (ei ExternalIdsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("custom_type"))
}

func (ei ExternalIdsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("type"))
}

func (ei ExternalIdsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ei.ref.Append("value"))
}

type ImsAttributes struct {
	ref terra.Reference
}

func (i ImsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i ImsAttributes) InternalWithRef(ref terra.Reference) ImsAttributes {
	return ImsAttributes{ref: ref}
}

func (i ImsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i ImsAttributes) CustomProtocol() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("custom_protocol"))
}

func (i ImsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("custom_type"))
}

func (i ImsAttributes) Im() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("im"))
}

func (i ImsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("primary"))
}

func (i ImsAttributes) Protocol() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("protocol"))
}

func (i ImsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("type"))
}

type KeywordsAttributes struct {
	ref terra.Reference
}

func (k KeywordsAttributes) InternalRef() (terra.Reference, error) {
	return k.ref, nil
}

func (k KeywordsAttributes) InternalWithRef(ref terra.Reference) KeywordsAttributes {
	return KeywordsAttributes{ref: ref}
}

func (k KeywordsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return k.ref.InternalTokens()
}

func (k KeywordsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("custom_type"))
}

func (k KeywordsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("type"))
}

func (k KeywordsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(k.ref.Append("value"))
}

type LanguagesAttributes struct {
	ref terra.Reference
}

func (l LanguagesAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LanguagesAttributes) InternalWithRef(ref terra.Reference) LanguagesAttributes {
	return LanguagesAttributes{ref: ref}
}

func (l LanguagesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LanguagesAttributes) CustomLanguage() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("custom_language"))
}

func (l LanguagesAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("language_code"))
}

func (l LanguagesAttributes) Preference() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("preference"))
}

type LocationsAttributes struct {
	ref terra.Reference
}

func (l LocationsAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocationsAttributes) InternalWithRef(ref terra.Reference) LocationsAttributes {
	return LocationsAttributes{ref: ref}
}

func (l LocationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocationsAttributes) Area() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("area"))
}

func (l LocationsAttributes) BuildingId() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("building_id"))
}

func (l LocationsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("custom_type"))
}

func (l LocationsAttributes) DeskCode() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("desk_code"))
}

func (l LocationsAttributes) FloorName() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("floor_name"))
}

func (l LocationsAttributes) FloorSection() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("floor_section"))
}

func (l LocationsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(l.ref.Append("type"))
}

type NameAttributes struct {
	ref terra.Reference
}

func (n NameAttributes) InternalRef() (terra.Reference, error) {
	return n.ref, nil
}

func (n NameAttributes) InternalWithRef(ref terra.Reference) NameAttributes {
	return NameAttributes{ref: ref}
}

func (n NameAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return n.ref.InternalTokens()
}

func (n NameAttributes) FamilyName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("family_name"))
}

func (n NameAttributes) FullName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("full_name"))
}

func (n NameAttributes) GivenName() terra.StringValue {
	return terra.ReferenceAsString(n.ref.Append("given_name"))
}

type OrganizationsAttributes struct {
	ref terra.Reference
}

func (o OrganizationsAttributes) InternalRef() (terra.Reference, error) {
	return o.ref, nil
}

func (o OrganizationsAttributes) InternalWithRef(ref terra.Reference) OrganizationsAttributes {
	return OrganizationsAttributes{ref: ref}
}

func (o OrganizationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return o.ref.InternalTokens()
}

func (o OrganizationsAttributes) CostCenter() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("cost_center"))
}

func (o OrganizationsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("custom_type"))
}

func (o OrganizationsAttributes) Department() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("department"))
}

func (o OrganizationsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("description"))
}

func (o OrganizationsAttributes) Domain() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("domain"))
}

func (o OrganizationsAttributes) FullTimeEquivalent() terra.NumberValue {
	return terra.ReferenceAsNumber(o.ref.Append("full_time_equivalent"))
}

func (o OrganizationsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("location"))
}

func (o OrganizationsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("name"))
}

func (o OrganizationsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(o.ref.Append("primary"))
}

func (o OrganizationsAttributes) Symbol() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("symbol"))
}

func (o OrganizationsAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("title"))
}

func (o OrganizationsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(o.ref.Append("type"))
}

type PhonesAttributes struct {
	ref terra.Reference
}

func (p PhonesAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PhonesAttributes) InternalWithRef(ref terra.Reference) PhonesAttributes {
	return PhonesAttributes{ref: ref}
}

func (p PhonesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PhonesAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("custom_type"))
}

func (p PhonesAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("primary"))
}

func (p PhonesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("type"))
}

func (p PhonesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("value"))
}

type PosixAccountsAttributes struct {
	ref terra.Reference
}

func (pa PosixAccountsAttributes) InternalRef() (terra.Reference, error) {
	return pa.ref, nil
}

func (pa PosixAccountsAttributes) InternalWithRef(ref terra.Reference) PosixAccountsAttributes {
	return PosixAccountsAttributes{ref: ref}
}

func (pa PosixAccountsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pa.ref.InternalTokens()
}

func (pa PosixAccountsAttributes) AccountId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("account_id"))
}

func (pa PosixAccountsAttributes) Gecos() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("gecos"))
}

func (pa PosixAccountsAttributes) Gid() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("gid"))
}

func (pa PosixAccountsAttributes) HomeDirectory() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("home_directory"))
}

func (pa PosixAccountsAttributes) OperatingSystemType() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("operating_system_type"))
}

func (pa PosixAccountsAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(pa.ref.Append("primary"))
}

func (pa PosixAccountsAttributes) Shell() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("shell"))
}

func (pa PosixAccountsAttributes) SystemId() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("system_id"))
}

func (pa PosixAccountsAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("uid"))
}

func (pa PosixAccountsAttributes) Username() terra.StringValue {
	return terra.ReferenceAsString(pa.ref.Append("username"))
}

type RelationsAttributes struct {
	ref terra.Reference
}

func (r RelationsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RelationsAttributes) InternalWithRef(ref terra.Reference) RelationsAttributes {
	return RelationsAttributes{ref: ref}
}

func (r RelationsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RelationsAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("custom_type"))
}

func (r RelationsAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("type"))
}

func (r RelationsAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("value"))
}

type SshPublicKeysAttributes struct {
	ref terra.Reference
}

func (spk SshPublicKeysAttributes) InternalRef() (terra.Reference, error) {
	return spk.ref, nil
}

func (spk SshPublicKeysAttributes) InternalWithRef(ref terra.Reference) SshPublicKeysAttributes {
	return SshPublicKeysAttributes{ref: ref}
}

func (spk SshPublicKeysAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return spk.ref.InternalTokens()
}

func (spk SshPublicKeysAttributes) ExpirationTimeUsec() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("expiration_time_usec"))
}

func (spk SshPublicKeysAttributes) Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("fingerprint"))
}

func (spk SshPublicKeysAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(spk.ref.Append("key"))
}

type WebsitesAttributes struct {
	ref terra.Reference
}

func (w WebsitesAttributes) InternalRef() (terra.Reference, error) {
	return w.ref, nil
}

func (w WebsitesAttributes) InternalWithRef(ref terra.Reference) WebsitesAttributes {
	return WebsitesAttributes{ref: ref}
}

func (w WebsitesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return w.ref.InternalTokens()
}

func (w WebsitesAttributes) CustomType() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("custom_type"))
}

func (w WebsitesAttributes) Primary() terra.BoolValue {
	return terra.ReferenceAsBool(w.ref.Append("primary"))
}

func (w WebsitesAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("type"))
}

func (w WebsitesAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(w.ref.Append("value"))
}

type AddressesState struct {
	Country            string `json:"country"`
	CountryCode        string `json:"country_code"`
	CustomType         string `json:"custom_type"`
	ExtendedAddress    string `json:"extended_address"`
	Formatted          string `json:"formatted"`
	Locality           string `json:"locality"`
	PoBox              string `json:"po_box"`
	PostalCode         string `json:"postal_code"`
	Primary            bool   `json:"primary"`
	Region             string `json:"region"`
	SourceIsStructured bool   `json:"source_is_structured"`
	StreetAddress      string `json:"street_address"`
	Type               string `json:"type"`
}

type CustomSchemasState struct {
	SchemaName   string            `json:"schema_name"`
	SchemaValues map[string]string `json:"schema_values"`
}

type EmailsState struct {
	Address    string `json:"address"`
	CustomType string `json:"custom_type"`
	Primary    bool   `json:"primary"`
	Type       string `json:"type"`
}

type ExternalIdsState struct {
	CustomType string `json:"custom_type"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type ImsState struct {
	CustomProtocol string `json:"custom_protocol"`
	CustomType     string `json:"custom_type"`
	Im             string `json:"im"`
	Primary        bool   `json:"primary"`
	Protocol       string `json:"protocol"`
	Type           string `json:"type"`
}

type KeywordsState struct {
	CustomType string `json:"custom_type"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type LanguagesState struct {
	CustomLanguage string `json:"custom_language"`
	LanguageCode   string `json:"language_code"`
	Preference     string `json:"preference"`
}

type LocationsState struct {
	Area         string `json:"area"`
	BuildingId   string `json:"building_id"`
	CustomType   string `json:"custom_type"`
	DeskCode     string `json:"desk_code"`
	FloorName    string `json:"floor_name"`
	FloorSection string `json:"floor_section"`
	Type         string `json:"type"`
}

type NameState struct {
	FamilyName string `json:"family_name"`
	FullName   string `json:"full_name"`
	GivenName  string `json:"given_name"`
}

type OrganizationsState struct {
	CostCenter         string  `json:"cost_center"`
	CustomType         string  `json:"custom_type"`
	Department         string  `json:"department"`
	Description        string  `json:"description"`
	Domain             string  `json:"domain"`
	FullTimeEquivalent float64 `json:"full_time_equivalent"`
	Location           string  `json:"location"`
	Name               string  `json:"name"`
	Primary            bool    `json:"primary"`
	Symbol             string  `json:"symbol"`
	Title              string  `json:"title"`
	Type               string  `json:"type"`
}

type PhonesState struct {
	CustomType string `json:"custom_type"`
	Primary    bool   `json:"primary"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type PosixAccountsState struct {
	AccountId           string `json:"account_id"`
	Gecos               string `json:"gecos"`
	Gid                 string `json:"gid"`
	HomeDirectory       string `json:"home_directory"`
	OperatingSystemType string `json:"operating_system_type"`
	Primary             bool   `json:"primary"`
	Shell               string `json:"shell"`
	SystemId            string `json:"system_id"`
	Uid                 string `json:"uid"`
	Username            string `json:"username"`
}

type RelationsState struct {
	CustomType string `json:"custom_type"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}

type SshPublicKeysState struct {
	ExpirationTimeUsec string `json:"expiration_time_usec"`
	Fingerprint        string `json:"fingerprint"`
	Key                string `json:"key"`
}

type WebsitesState struct {
	CustomType string `json:"custom_type"`
	Primary    bool   `json:"primary"`
	Type       string `json:"type"`
	Value      string `json:"value"`
}
