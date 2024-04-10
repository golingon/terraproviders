// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datardscluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MasterUserSecret struct{}

type MasterUserSecretAttributes struct {
	ref terra.Reference
}

func (mus MasterUserSecretAttributes) InternalRef() (terra.Reference, error) {
	return mus.ref, nil
}

func (mus MasterUserSecretAttributes) InternalWithRef(ref terra.Reference) MasterUserSecretAttributes {
	return MasterUserSecretAttributes{ref: ref}
}

func (mus MasterUserSecretAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mus.ref.InternalTokens()
}

func (mus MasterUserSecretAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(mus.ref.Append("kms_key_id"))
}

func (mus MasterUserSecretAttributes) SecretArn() terra.StringValue {
	return terra.ReferenceAsString(mus.ref.Append("secret_arn"))
}

func (mus MasterUserSecretAttributes) SecretStatus() terra.StringValue {
	return terra.ReferenceAsString(mus.ref.Append("secret_status"))
}

type MasterUserSecretState struct {
	KmsKeyId     string `json:"kms_key_id"`
	SecretArn    string `json:"secret_arn"`
	SecretStatus string `json:"secret_status"`
}
