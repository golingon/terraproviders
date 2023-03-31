// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package macie2classificationexportconfiguration

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type S3Destination struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// KeyPrefix: string, optional
	KeyPrefix terra.StringValue `hcl:"key_prefix,attr"`
	// KmsKeyArn: string, required
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr" validate:"required"`
}

type S3DestinationAttributes struct {
	ref terra.Reference
}

func (sd S3DestinationAttributes) InternalRef() terra.Reference {
	return sd.ref
}

func (sd S3DestinationAttributes) InternalWithRef(ref terra.Reference) S3DestinationAttributes {
	return S3DestinationAttributes{ref: ref}
}

func (sd S3DestinationAttributes) InternalTokens() hclwrite.Tokens {
	return sd.ref.InternalTokens()
}

func (sd S3DestinationAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("bucket_name"))
}

func (sd S3DestinationAttributes) KeyPrefix() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("key_prefix"))
}

func (sd S3DestinationAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("kms_key_arn"))
}

type S3DestinationState struct {
	BucketName string `json:"bucket_name"`
	KeyPrefix  string `json:"key_prefix"`
	KmsKeyArn  string `json:"kms_key_arn"`
}
