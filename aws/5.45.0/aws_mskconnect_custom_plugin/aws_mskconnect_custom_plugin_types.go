// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_mskconnect_custom_plugin

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Location struct {
	// LocationS3: required
	S3 *LocationS3 `hcl:"s3,block" validate:"required"`
}

type LocationS3 struct {
	// BucketArn: string, required
	BucketArn terra.StringValue `hcl:"bucket_arn,attr" validate:"required"`
	// FileKey: string, required
	FileKey terra.StringValue `hcl:"file_key,attr" validate:"required"`
	// ObjectVersion: string, optional
	ObjectVersion terra.StringValue `hcl:"object_version,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type LocationAttributes struct {
	ref terra.Reference
}

func (l LocationAttributes) InternalRef() (terra.Reference, error) {
	return l.ref, nil
}

func (l LocationAttributes) InternalWithRef(ref terra.Reference) LocationAttributes {
	return LocationAttributes{ref: ref}
}

func (l LocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return l.ref.InternalTokens()
}

func (l LocationAttributes) S3() terra.ListValue[LocationS3Attributes] {
	return terra.ReferenceAsList[LocationS3Attributes](l.ref.Append("s3"))
}

type LocationS3Attributes struct {
	ref terra.Reference
}

func (s LocationS3Attributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s LocationS3Attributes) InternalWithRef(ref terra.Reference) LocationS3Attributes {
	return LocationS3Attributes{ref: ref}
}

func (s LocationS3Attributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s LocationS3Attributes) BucketArn() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("bucket_arn"))
}

func (s LocationS3Attributes) FileKey() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("file_key"))
}

func (s LocationS3Attributes) ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("object_version"))
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

type LocationState struct {
	S3 []LocationS3State `json:"s3"`
}

type LocationS3State struct {
	BucketArn     string `json:"bucket_arn"`
	FileKey       string `json:"file_key"`
	ObjectVersion string `json:"object_version"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
