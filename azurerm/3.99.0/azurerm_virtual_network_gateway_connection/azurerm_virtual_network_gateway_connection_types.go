// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_virtual_network_gateway_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CustomBgpAddresses struct {
	// Primary: string, required
	Primary terra.StringValue `hcl:"primary,attr" validate:"required"`
	// Secondary: string, optional
	Secondary terra.StringValue `hcl:"secondary,attr"`
}

type IpsecPolicy struct {
	// DhGroup: string, required
	DhGroup terra.StringValue `hcl:"dh_group,attr" validate:"required"`
	// IkeEncryption: string, required
	IkeEncryption terra.StringValue `hcl:"ike_encryption,attr" validate:"required"`
	// IkeIntegrity: string, required
	IkeIntegrity terra.StringValue `hcl:"ike_integrity,attr" validate:"required"`
	// IpsecEncryption: string, required
	IpsecEncryption terra.StringValue `hcl:"ipsec_encryption,attr" validate:"required"`
	// IpsecIntegrity: string, required
	IpsecIntegrity terra.StringValue `hcl:"ipsec_integrity,attr" validate:"required"`
	// PfsGroup: string, required
	PfsGroup terra.StringValue `hcl:"pfs_group,attr" validate:"required"`
	// SaDatasize: number, optional
	SaDatasize terra.NumberValue `hcl:"sa_datasize,attr"`
	// SaLifetime: number, optional
	SaLifetime terra.NumberValue `hcl:"sa_lifetime,attr"`
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

type TrafficSelectorPolicy struct {
	// LocalAddressCidrs: list of string, required
	LocalAddressCidrs terra.ListValue[terra.StringValue] `hcl:"local_address_cidrs,attr" validate:"required"`
	// RemoteAddressCidrs: list of string, required
	RemoteAddressCidrs terra.ListValue[terra.StringValue] `hcl:"remote_address_cidrs,attr" validate:"required"`
}

type CustomBgpAddressesAttributes struct {
	ref terra.Reference
}

func (cba CustomBgpAddressesAttributes) InternalRef() (terra.Reference, error) {
	return cba.ref, nil
}

func (cba CustomBgpAddressesAttributes) InternalWithRef(ref terra.Reference) CustomBgpAddressesAttributes {
	return CustomBgpAddressesAttributes{ref: ref}
}

func (cba CustomBgpAddressesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cba.ref.InternalTokens()
}

func (cba CustomBgpAddressesAttributes) Primary() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("primary"))
}

func (cba CustomBgpAddressesAttributes) Secondary() terra.StringValue {
	return terra.ReferenceAsString(cba.ref.Append("secondary"))
}

type IpsecPolicyAttributes struct {
	ref terra.Reference
}

func (ip IpsecPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip IpsecPolicyAttributes) InternalWithRef(ref terra.Reference) IpsecPolicyAttributes {
	return IpsecPolicyAttributes{ref: ref}
}

func (ip IpsecPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip IpsecPolicyAttributes) DhGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("dh_group"))
}

func (ip IpsecPolicyAttributes) IkeEncryption() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_encryption"))
}

func (ip IpsecPolicyAttributes) IkeIntegrity() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_integrity"))
}

func (ip IpsecPolicyAttributes) IpsecEncryption() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ipsec_encryption"))
}

func (ip IpsecPolicyAttributes) IpsecIntegrity() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ipsec_integrity"))
}

func (ip IpsecPolicyAttributes) PfsGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("pfs_group"))
}

func (ip IpsecPolicyAttributes) SaDatasize() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_datasize"))
}

func (ip IpsecPolicyAttributes) SaLifetime() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_lifetime"))
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

type TrafficSelectorPolicyAttributes struct {
	ref terra.Reference
}

func (tsp TrafficSelectorPolicyAttributes) InternalRef() (terra.Reference, error) {
	return tsp.ref, nil
}

func (tsp TrafficSelectorPolicyAttributes) InternalWithRef(ref terra.Reference) TrafficSelectorPolicyAttributes {
	return TrafficSelectorPolicyAttributes{ref: ref}
}

func (tsp TrafficSelectorPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return tsp.ref.InternalTokens()
}

func (tsp TrafficSelectorPolicyAttributes) LocalAddressCidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tsp.ref.Append("local_address_cidrs"))
}

func (tsp TrafficSelectorPolicyAttributes) RemoteAddressCidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](tsp.ref.Append("remote_address_cidrs"))
}

type CustomBgpAddressesState struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

type IpsecPolicyState struct {
	DhGroup         string  `json:"dh_group"`
	IkeEncryption   string  `json:"ike_encryption"`
	IkeIntegrity    string  `json:"ike_integrity"`
	IpsecEncryption string  `json:"ipsec_encryption"`
	IpsecIntegrity  string  `json:"ipsec_integrity"`
	PfsGroup        string  `json:"pfs_group"`
	SaDatasize      float64 `json:"sa_datasize"`
	SaLifetime      float64 `json:"sa_lifetime"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TrafficSelectorPolicyState struct {
	LocalAddressCidrs  []string `json:"local_address_cidrs"`
	RemoteAddressCidrs []string `json:"remote_address_cidrs"`
}
