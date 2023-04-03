// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssmassociation

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type OutputLocation struct {
	// S3BucketName: string, required
	S3BucketName terra.StringValue `hcl:"s3_bucket_name,attr" validate:"required"`
	// S3KeyPrefix: string, optional
	S3KeyPrefix terra.StringValue `hcl:"s3_key_prefix,attr"`
	// S3Region: string, optional
	S3Region terra.StringValue `hcl:"s3_region,attr"`
}

type Targets struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type OutputLocationAttributes struct {
	ref terra.Reference
}

func (ol OutputLocationAttributes) InternalRef() (terra.Reference, error) {
	return ol.ref, nil
}

func (ol OutputLocationAttributes) InternalWithRef(ref terra.Reference) OutputLocationAttributes {
	return OutputLocationAttributes{ref: ref}
}

func (ol OutputLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ol.ref.InternalTokens()
}

func (ol OutputLocationAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("s3_bucket_name"))
}

func (ol OutputLocationAttributes) S3KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("s3_key_prefix"))
}

func (ol OutputLocationAttributes) S3Region() terra.StringValue {
	return terra.ReferenceAsString(ol.ref.Append("s3_region"))
}

type TargetsAttributes struct {
	ref terra.Reference
}

func (t TargetsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TargetsAttributes) InternalWithRef(ref terra.Reference) TargetsAttributes {
	return TargetsAttributes{ref: ref}
}

func (t TargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TargetsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TargetsAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("values"))
}

type OutputLocationState struct {
	S3BucketName string `json:"s3_bucket_name"`
	S3KeyPrefix  string `json:"s3_key_prefix"`
	S3Region     string `json:"s3_region"`
}

type TargetsState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
