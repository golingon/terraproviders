// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package gameliftscript

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type StorageLocation struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ObjectVersion: string, optional
	ObjectVersion terra.StringValue `hcl:"object_version,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}

type StorageLocationAttributes struct {
	ref terra.Reference
}

func (sl StorageLocationAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl StorageLocationAttributes) InternalWithRef(ref terra.Reference) StorageLocationAttributes {
	return StorageLocationAttributes{ref: ref}
}

func (sl StorageLocationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sl.ref.InternalTokens()
}

func (sl StorageLocationAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("bucket"))
}

func (sl StorageLocationAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("key"))
}

func (sl StorageLocationAttributes) ObjectVersion() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("object_version"))
}

func (sl StorageLocationAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("role_arn"))
}

type StorageLocationState struct {
	Bucket        string `json:"bucket"`
	Key           string `json:"key"`
	ObjectVersion string `json:"object_version"`
	RoleArn       string `json:"role_arn"`
}
