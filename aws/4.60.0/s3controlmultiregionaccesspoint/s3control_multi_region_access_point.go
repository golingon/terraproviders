// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3controlmultiregionaccesspoint

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Details struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicAccessBlock: optional
	PublicAccessBlock *PublicAccessBlock `hcl:"public_access_block,block"`
	// Region: min=1,max=20
	Region []Region `hcl:"region,block" validate:"min=1,max=20"`
}

type PublicAccessBlock struct {
	// BlockPublicAcls: bool, optional
	BlockPublicAcls terra.BoolValue `hcl:"block_public_acls,attr"`
	// BlockPublicPolicy: bool, optional
	BlockPublicPolicy terra.BoolValue `hcl:"block_public_policy,attr"`
	// IgnorePublicAcls: bool, optional
	IgnorePublicAcls terra.BoolValue `hcl:"ignore_public_acls,attr"`
	// RestrictPublicBuckets: bool, optional
	RestrictPublicBuckets terra.BoolValue `hcl:"restrict_public_buckets,attr"`
}

type Region struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type DetailsAttributes struct {
	ref terra.Reference
}

func (d DetailsAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DetailsAttributes) InternalWithRef(ref terra.Reference) DetailsAttributes {
	return DetailsAttributes{ref: ref}
}

func (d DetailsAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d DetailsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(d.ref.Append("name"))
}

func (d DetailsAttributes) PublicAccessBlock() terra.ListValue[PublicAccessBlockAttributes] {
	return terra.ReferenceAsList[PublicAccessBlockAttributes](d.ref.Append("public_access_block"))
}

func (d DetailsAttributes) Region() terra.SetValue[RegionAttributes] {
	return terra.ReferenceAsSet[RegionAttributes](d.ref.Append("region"))
}

type PublicAccessBlockAttributes struct {
	ref terra.Reference
}

func (pab PublicAccessBlockAttributes) InternalRef() (terra.Reference, error) {
	return pab.ref, nil
}

func (pab PublicAccessBlockAttributes) InternalWithRef(ref terra.Reference) PublicAccessBlockAttributes {
	return PublicAccessBlockAttributes{ref: ref}
}

func (pab PublicAccessBlockAttributes) InternalTokens() hclwrite.Tokens {
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

type RegionAttributes struct {
	ref terra.Reference
}

func (r RegionAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r RegionAttributes) InternalWithRef(ref terra.Reference) RegionAttributes {
	return RegionAttributes{ref: ref}
}

func (r RegionAttributes) InternalTokens() hclwrite.Tokens {
	return r.ref.InternalTokens()
}

func (r RegionAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("bucket"))
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

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type DetailsState struct {
	Name              string                   `json:"name"`
	PublicAccessBlock []PublicAccessBlockState `json:"public_access_block"`
	Region            []RegionState            `json:"region"`
}

type PublicAccessBlockState struct {
	BlockPublicAcls       bool `json:"block_public_acls"`
	BlockPublicPolicy     bool `json:"block_public_policy"`
	IgnorePublicAcls      bool `json:"ignore_public_acls"`
	RestrictPublicBuckets bool `json:"restrict_public_buckets"`
}

type RegionState struct {
	Bucket string `json:"bucket"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
