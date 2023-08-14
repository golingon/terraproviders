// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ssmresourcedatasync

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type S3Destination struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
	// SyncFormat: string, optional
	SyncFormat terra.StringValue `hcl:"sync_format,attr"`
}

type S3DestinationAttributes struct {
	ref terra.Reference
}

func (sd S3DestinationAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd S3DestinationAttributes) InternalWithRef(ref terra.Reference) S3DestinationAttributes {
	return S3DestinationAttributes{ref: ref}
}

func (sd S3DestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd S3DestinationAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("bucket_name"))
}

func (sd S3DestinationAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("kms_key_arn"))
}

func (sd S3DestinationAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("prefix"))
}

func (sd S3DestinationAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("region"))
}

func (sd S3DestinationAttributes) SyncFormat() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("sync_format"))
}

type S3DestinationState struct {
	BucketName string `json:"bucket_name"`
	KmsKeyArn  string `json:"kms_key_arn"`
	Prefix     string `json:"prefix"`
	Region     string `json:"region"`
	SyncFormat string `json:"sync_format"`
}