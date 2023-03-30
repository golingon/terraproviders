// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package worklinkfleet

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type IdentityProvider struct {
	// SamlMetadata: string, required
	SamlMetadata terra.StringValue `hcl:"saml_metadata,attr" validate:"required"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
}

type Network struct {
	// SecurityGroupIds: set of string, required
	SecurityGroupIds terra.SetValue[terra.StringValue] `hcl:"security_group_ids,attr" validate:"required"`
	// SubnetIds: set of string, required
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}

type IdentityProviderAttributes struct {
	ref terra.Reference
}

func (ip IdentityProviderAttributes) InternalRef() terra.Reference {
	return ip.ref
}

func (ip IdentityProviderAttributes) InternalWithRef(ref terra.Reference) IdentityProviderAttributes {
	return IdentityProviderAttributes{ref: ref}
}

func (ip IdentityProviderAttributes) InternalTokens() hclwrite.Tokens {
	return ip.ref.InternalTokens()
}

func (ip IdentityProviderAttributes) SamlMetadata() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("saml_metadata"))
}

func (ip IdentityProviderAttributes) Type() terra.StringValue {
	return terra.ReferenceString(ip.ref.Append("type"))
}

type NetworkAttributes struct {
	ref terra.Reference
}

func (n NetworkAttributes) InternalRef() terra.Reference {
	return n.ref
}

func (n NetworkAttributes) InternalWithRef(ref terra.Reference) NetworkAttributes {
	return NetworkAttributes{ref: ref}
}

func (n NetworkAttributes) InternalTokens() hclwrite.Tokens {
	return n.ref.InternalTokens()
}

func (n NetworkAttributes) SecurityGroupIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](n.ref.Append("security_group_ids"))
}

func (n NetworkAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](n.ref.Append("subnet_ids"))
}

func (n NetworkAttributes) VpcId() terra.StringValue {
	return terra.ReferenceString(n.ref.Append("vpc_id"))
}

type IdentityProviderState struct {
	SamlMetadata string `json:"saml_metadata"`
	Type         string `json:"type"`
}

type NetworkState struct {
	SecurityGroupIds []string `json:"security_group_ids"`
	SubnetIds        []string `json:"subnet_ids"`
	VpcId            string   `json:"vpc_id"`
}
