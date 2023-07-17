// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package athenadatabase

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AclConfiguration struct {
	// S3AclOption: string, required
	S3AclOption terra.StringValue `hcl:"s3_acl_option,attr" validate:"required"`
}

type EncryptionConfiguration struct {
	// EncryptionOption: string, required
	EncryptionOption terra.StringValue `hcl:"encryption_option,attr" validate:"required"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
}

type AclConfigurationAttributes struct {
	ref terra.Reference
}

func (ac AclConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac AclConfigurationAttributes) InternalWithRef(ref terra.Reference) AclConfigurationAttributes {
	return AclConfigurationAttributes{ref: ref}
}

func (ac AclConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac AclConfigurationAttributes) S3AclOption() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("s3_acl_option"))
}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) EncryptionOption() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_option"))
}

func (ec EncryptionConfigurationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type AclConfigurationState struct {
	S3AclOption string `json:"s3_acl_option"`
}

type EncryptionConfigurationState struct {
	EncryptionOption string `json:"encryption_option"`
	KmsKey           string `json:"kms_key"`
}