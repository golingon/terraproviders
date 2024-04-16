// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_secretsmanager_secret

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Replica struct {
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Region: string, required
	Region terra.StringValue `hcl:"region,attr" validate:"required"`
}

type ReplicaAttributes struct {
	ref terra.Reference
}

func (r ReplicaAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ReplicaAttributes) InternalWithRef(ref terra.Reference) ReplicaAttributes {
	return ReplicaAttributes{ref: ref}
}

func (r ReplicaAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ReplicaAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("kms_key_id"))
}

func (r ReplicaAttributes) LastAccessedDate() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("last_accessed_date"))
}

func (r ReplicaAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("region"))
}

func (r ReplicaAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status"))
}

func (r ReplicaAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("status_message"))
}

type ReplicaState struct {
	KmsKeyId         string `json:"kms_key_id"`
	LastAccessedDate string `json:"last_accessed_date"`
	Region           string `json:"region"`
	Status           string `json:"status"`
	StatusMessage    string `json:"status_message"`
}
