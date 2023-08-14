// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datas3controlmultiregionaccesspoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type PublicAccessBlock struct{}

type Regions struct{}

type PublicAccessBlockAttributes struct {
	ref terra.Reference
}

func (pab PublicAccessBlockAttributes) InternalRef() (terra.Reference, error) {
	return pab.ref, nil
}

func (pab PublicAccessBlockAttributes) InternalWithRef(ref terra.Reference) PublicAccessBlockAttributes {
	return PublicAccessBlockAttributes{ref: ref}
}

func (pab PublicAccessBlockAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pab.ref.InternalTokens()
}

func (pab PublicAccessBlockAttributes) BlockPublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(pab.ref.Append("block_public_acls"))
}

func (pab PublicAccessBlockAttributes) BlockPublicPolicy() terra.BoolValue {
	return terra.ReferenceAsBool(pab.ref.Append("block_public_policy"))
}

func (pab PublicAccessBlockAttributes) IgnorePublicAcls() terra.BoolValue {
	return terra.ReferenceAsBool(pab.ref.Append("ignore_public_acls"))
}

func (pab PublicAccessBlockAttributes) RestrictPublicBuckets() terra.BoolValue {
	return terra.ReferenceAsBool(pab.ref.Append("restrict_public_buckets"))
}

type RegionsAttributes struct {
	ref terra.Reference
}

func (r RegionsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RegionsAttributes) InternalWithRef(ref terra.Reference) RegionsAttributes {
	return RegionsAttributes{ref: ref}
}

func (r RegionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r RegionsAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("bucket"))
}

func (r RegionsAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("region"))
}

type PublicAccessBlockState struct {
	BlockPublicAcls       bool `json:"block_public_acls"`
	BlockPublicPolicy     bool `json:"block_public_policy"`
	IgnorePublicAcls      bool `json:"ignore_public_acls"`
	RestrictPublicBuckets bool `json:"restrict_public_buckets"`
}

type RegionsState struct {
	Bucket string `json:"bucket"`
	Region string `json:"region"`
}
