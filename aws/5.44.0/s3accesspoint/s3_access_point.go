// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package s3accesspoint

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PublicAccessBlockConfiguration struct {
	// BlockPublicAcls: bool, optional
	BlockPublicAcls terra.BoolValue `hcl:"block_public_acls,attr"`
	// BlockPublicPolicy: bool, optional
	BlockPublicPolicy terra.BoolValue `hcl:"block_public_policy,attr"`
	// IgnorePublicAcls: bool, optional
	IgnorePublicAcls terra.BoolValue `hcl:"ignore_public_acls,attr"`
	// RestrictPublicBuckets: bool, optional
	RestrictPublicBuckets terra.BoolValue `hcl:"restrict_public_buckets,attr"`
}

type VpcConfiguration struct {
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
}

type PublicAccessBlockConfigurationAttributes struct {
	ref terra.Reference
}

func (pabc PublicAccessBlockConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return pabc.ref, nil
}

func (pabc PublicAccessBlockConfigurationAttributes) InternalWithRef(ref terra.Reference) PublicAccessBlockConfigurationAttributes {
	return PublicAccessBlockConfigurationAttributes{ref: ref}
}

func (pabc PublicAccessBlockConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pabc.ref.InternalTokens()
}

func (pabc PublicAccessBlockConfigurationAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(pabc.ref.Append("block_public_acls"))
}

func (pabc PublicAccessBlockConfigurationAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(pabc.ref.Append("block_public_policy"))
}

func (pabc PublicAccessBlockConfigurationAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(pabc.ref.Append("ignore_public_acls"))
}

func (pabc PublicAccessBlockConfigurationAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(pabc.ref.Append("restrict_public_buckets"))
}

type VpcConfigurationAttributes struct {
	ref terra.Reference
}

func (vc VpcConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return vc.ref, nil
}

func (vc VpcConfigurationAttributes) InternalWithRef(ref terra.Reference) VpcConfigurationAttributes {
	return VpcConfigurationAttributes{ref: ref}
}

func (vc VpcConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vc.ref.InternalTokens()
}

func (vc VpcConfigurationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vc.ref.Append("vpc_id"))
}

type PublicAccessBlockConfigurationState struct {
	BlockPublicAcls       bool `json:"block_public_acls"`
	BlockPublicPolicy     bool `json:"block_public_policy"`
	IgnorePublicAcls      bool `json:"ignore_public_acls"`
	RestrictPublicBuckets bool `json:"restrict_public_buckets"`
}

type VpcConfigurationState struct {
	VpcId string `json:"vpc_id"`
}
