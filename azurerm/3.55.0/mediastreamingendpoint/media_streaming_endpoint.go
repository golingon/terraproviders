// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediastreamingendpoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Sku struct{}

type AccessControl struct {
	// AkamaiSignatureHeaderAuthenticationKey: min=0
	AkamaiSignatureHeaderAuthenticationKey []AkamaiSignatureHeaderAuthenticationKey `hcl:"akamai_signature_header_authentication_key,block" validate:"min=0"`
	// IpAllow: min=0
	IpAllow []IpAllow `hcl:"ip_allow,block" validate:"min=0"`
}

type AkamaiSignatureHeaderAuthenticationKey struct {
	// Base64Key: string, optional
	Base64Key terra.StringValue `hcl:"base64_key,attr"`
	// Expiration: string, optional
	Expiration terra.StringValue `hcl:"expiration,attr"`
	// Identifier: string, optional
	Identifier terra.StringValue `hcl:"identifier,attr"`
}

type IpAllow struct {
	// Address: string, optional
	Address terra.StringValue `hcl:"address,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// SubnetPrefixLength: number, optional
	SubnetPrefixLength terra.NumberValue `hcl:"subnet_prefix_length,attr"`
}

type CrossSiteAccessPolicy struct {
	// ClientAccessPolicy: string, optional
	ClientAccessPolicy terra.StringValue `hcl:"client_access_policy,attr"`
	// CrossDomainPolicy: string, optional
	CrossDomainPolicy terra.StringValue `hcl:"cross_domain_policy,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SkuAttributes struct {
	ref terra.Reference
}

func (s SkuAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SkuAttributes) InternalWithRef(ref terra.Reference) SkuAttributes {
	return SkuAttributes{ref: ref}
}

func (s SkuAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SkuAttributes) Capacity() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("capacity"))
}

func (s SkuAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

type AccessControlAttributes struct {
	ref terra.Reference
}

func (ac AccessControlAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AccessControlAttributes) InternalWithRef(ref terra.Reference) AccessControlAttributes {
	return AccessControlAttributes{ref: ref}
}

func (ac AccessControlAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AccessControlAttributes) AkamaiSignatureHeaderAuthenticationKey() terra.ListValue[AkamaiSignatureHeaderAuthenticationKeyAttributes] {
	return terra.ReferenceAsList[AkamaiSignatureHeaderAuthenticationKeyAttributes](ac.ref.Append("akamai_signature_header_authentication_key"))
}

func (ac AccessControlAttributes) IpAllow() terra.ListValue[IpAllowAttributes] {
	return terra.ReferenceAsList[IpAllowAttributes](ac.ref.Append("ip_allow"))
}

type AkamaiSignatureHeaderAuthenticationKeyAttributes struct {
	ref terra.Reference
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) InternalRef() (terra.Reference, error) {
	return ashak.ref, nil
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) InternalWithRef(ref terra.Reference) AkamaiSignatureHeaderAuthenticationKeyAttributes {
	return AkamaiSignatureHeaderAuthenticationKeyAttributes{ref: ref}
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ashak.ref.InternalTokens()
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) Base64Key() terra.StringValue {
	return terra.ReferenceAsString(ashak.ref.Append("base64_key"))
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(ashak.ref.Append("expiration"))
}

func (ashak AkamaiSignatureHeaderAuthenticationKeyAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(ashak.ref.Append("identifier"))
}

type IpAllowAttributes struct {
	ref terra.Reference
}

func (ia IpAllowAttributes) InternalRef() (terra.Reference, error) {
	return ia.ref, nil
}

func (ia IpAllowAttributes) InternalWithRef(ref terra.Reference) IpAllowAttributes {
	return IpAllowAttributes{ref: ref}
}

func (ia IpAllowAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ia.ref.InternalTokens()
}

func (ia IpAllowAttributes) Address() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("address"))
}

func (ia IpAllowAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ia.ref.Append("name"))
}

func (ia IpAllowAttributes) SubnetPrefixLength() terra.NumberValue {
	return terra.ReferenceAsNumber(ia.ref.Append("subnet_prefix_length"))
}

type CrossSiteAccessPolicyAttributes struct {
	ref terra.Reference
}

func (csap CrossSiteAccessPolicyAttributes) InternalRef() (terra.Reference, error) {
	return csap.ref, nil
}

func (csap CrossSiteAccessPolicyAttributes) InternalWithRef(ref terra.Reference) CrossSiteAccessPolicyAttributes {
	return CrossSiteAccessPolicyAttributes{ref: ref}
}

func (csap CrossSiteAccessPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return csap.ref.InternalTokens()
}

func (csap CrossSiteAccessPolicyAttributes) ClientAccessPolicy() terra.StringValue {
	return terra.ReferenceAsString(csap.ref.Append("client_access_policy"))
}

func (csap CrossSiteAccessPolicyAttributes) CrossDomainPolicy() terra.StringValue {
	return terra.ReferenceAsString(csap.ref.Append("cross_domain_policy"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type SkuState struct {
	Capacity float64 `json:"capacity"`
	Name     string  `json:"name"`
}

type AccessControlState struct {
	AkamaiSignatureHeaderAuthenticationKey []AkamaiSignatureHeaderAuthenticationKeyState `json:"akamai_signature_header_authentication_key"`
	IpAllow                                []IpAllowState                                `json:"ip_allow"`
}

type AkamaiSignatureHeaderAuthenticationKeyState struct {
	Base64Key  string `json:"base64_key"`
	Expiration string `json:"expiration"`
	Identifier string `json:"identifier"`
}

type IpAllowState struct {
	Address            string  `json:"address"`
	Name               string  `json:"name"`
	SubnetPrefixLength float64 `json:"subnet_prefix_length"`
}

type CrossSiteAccessPolicyState struct {
	ClientAccessPolicy string `json:"client_access_policy"`
	CrossDomainPolicy  string `json:"cross_domain_policy"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
