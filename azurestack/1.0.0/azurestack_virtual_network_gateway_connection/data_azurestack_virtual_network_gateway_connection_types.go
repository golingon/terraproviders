// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_virtual_network_gateway_connection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataIpsecPolicyAttributes struct {
	ref terra.Reference
}

func (ip DataIpsecPolicyAttributes) InternalRef() (terra.Reference, error) {
	return ip.ref, nil
}

func (ip DataIpsecPolicyAttributes) InternalWithRef(ref terra.Reference) DataIpsecPolicyAttributes {
	return DataIpsecPolicyAttributes{ref: ref}
}

func (ip DataIpsecPolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ip.ref.InternalTokens()
}

func (ip DataIpsecPolicyAttributes) DhGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("dh_group"))
}

func (ip DataIpsecPolicyAttributes) IkeEncryption() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_encryption"))
}

func (ip DataIpsecPolicyAttributes) IkeIntegrity() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ike_integrity"))
}

func (ip DataIpsecPolicyAttributes) IpsecEncryption() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ipsec_encryption"))
}

func (ip DataIpsecPolicyAttributes) IpsecIntegrity() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("ipsec_integrity"))
}

func (ip DataIpsecPolicyAttributes) PfsGroup() terra.StringValue {
	return terra.ReferenceAsString(ip.ref.Append("pfs_group"))
}

func (ip DataIpsecPolicyAttributes) SaDatasize() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_datasize"))
}

func (ip DataIpsecPolicyAttributes) SaLifetime() terra.NumberValue {
	return terra.ReferenceAsNumber(ip.ref.Append("sa_lifetime"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataIpsecPolicyState struct {
	DhGroup         string  `json:"dh_group"`
	IkeEncryption   string  `json:"ike_encryption"`
	IkeIntegrity    string  `json:"ike_integrity"`
	IpsecEncryption string  `json:"ipsec_encryption"`
	IpsecIntegrity  string  `json:"ipsec_integrity"`
	PfsGroup        string  `json:"pfs_group"`
	SaDatasize      float64 `json:"sa_datasize"`
	SaLifetime      float64 `json:"sa_lifetime"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
