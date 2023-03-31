// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datavirtualnetworkgatewayconnection

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IpsecPolicy struct{}

type TrafficSelectorPolicy struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type IpsecPolicyAttributes struct {
	ref terra.Reference
}

func (ip IpsecPolicyAttributes) InternalRef() terra.Reference {
	return ip.ref
}

func (ip IpsecPolicyAttributes) InternalWithRef(ref terra.Reference) IpsecPolicyAttributes {
	return IpsecPolicyAttributes{ref: ref}
}

func (ip IpsecPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return ip.ref.InternalTokens()
}

func (ip IpsecPolicyAttributes) DhGroup() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("dh_group"))
}

func (ip IpsecPolicyAttributes) IkeEncryption() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("ike_encryption"))
}

func (ip IpsecPolicyAttributes) IkeIntegrity() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("ike_integrity"))
}

func (ip IpsecPolicyAttributes) IpsecEncryption() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("ipsec_encryption"))
}

func (ip IpsecPolicyAttributes) IpsecIntegrity() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("ipsec_integrity"))
}

func (ip IpsecPolicyAttributes) PfsGroup() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("pfs_group"))
}

func (ip IpsecPolicyAttributes) SaDatasize() terra.NumberValue {
	return terra.ReferenceNumber(ip.ref.Append("sa_datasize"))
}

func (ip IpsecPolicyAttributes) SaLifetime() terra.NumberValue {
	return terra.ReferenceNumber(ip.ref.Append("sa_lifetime"))
}

type TrafficSelectorPolicyAttributes struct {
	ref terra.Reference
}

func (tsp TrafficSelectorPolicyAttributes) InternalRef() terra.Reference {
	return tsp.ref
}

func (tsp TrafficSelectorPolicyAttributes) InternalWithRef(ref terra.Reference) TrafficSelectorPolicyAttributes {
	return TrafficSelectorPolicyAttributes{ref: ref}
}

func (tsp TrafficSelectorPolicyAttributes) InternalTokens() hclwrite.Tokens {
	return tsp.ref.InternalTokens()
}

func (tsp TrafficSelectorPolicyAttributes) LocalAddressCidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](tsp.ref.Append("local_address_cidrs"))
}

func (tsp TrafficSelectorPolicyAttributes) RemoteAddressCidrs() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](tsp.ref.Append("remote_address_cidrs"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
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

type TrafficSelectorPolicyState struct {
	LocalAddressCidrs  []string `json:"local_address_cidrs"`
	RemoteAddressCidrs []string `json:"remote_address_cidrs"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
