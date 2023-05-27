// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package gluedatacatalogencryptionsettings

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type DataCatalogEncryptionSettings struct {
	// ConnectionPasswordEncryption: required
	ConnectionPasswordEncryption *ConnectionPasswordEncryption `hcl:"connection_password_encryption,block" validate:"required"`
	// EncryptionAtRest: required
	EncryptionAtRest *EncryptionAtRest `hcl:"encryption_at_rest,block" validate:"required"`
}

type ConnectionPasswordEncryption struct {
	// AwsKmsKeyId: string, optional
	AwsKmsKeyId terra.StringValue `hcl:"aws_kms_key_id,attr"`
	// ReturnConnectionPasswordEncrypted: bool, required
	ReturnConnectionPasswordEncrypted terra.BoolValue `hcl:"return_connection_password_encrypted,attr" validate:"required"`
}

type EncryptionAtRest struct {
	// CatalogEncryptionMode: string, required
	CatalogEncryptionMode terra.StringValue `hcl:"catalog_encryption_mode,attr" validate:"required"`
	// SseAwsKmsKeyId: string, optional
	SseAwsKmsKeyId terra.StringValue `hcl:"sse_aws_kms_key_id,attr"`
}

type DataCatalogEncryptionSettingsAttributes struct {
	ref terra.Reference
}

func (dces DataCatalogEncryptionSettingsAttributes) InternalRef() (terra.Reference, error) {
	return dces.ref, nil
}

func (dces DataCatalogEncryptionSettingsAttributes) InternalWithRef(ref terra.Reference) DataCatalogEncryptionSettingsAttributes {
	return DataCatalogEncryptionSettingsAttributes{ref: ref}
}

func (dces DataCatalogEncryptionSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dces.ref.InternalTokens()
}

func (dces DataCatalogEncryptionSettingsAttributes) ConnectionPasswordEncryption() terra.ListValue[ConnectionPasswordEncryptionAttributes] {
	return terra.ReferenceAsList[ConnectionPasswordEncryptionAttributes](dces.ref.Append("connection_password_encryption"))
}

func (dces DataCatalogEncryptionSettingsAttributes) EncryptionAtRest() terra.ListValue[EncryptionAtRestAttributes] {
	return terra.ReferenceAsList[EncryptionAtRestAttributes](dces.ref.Append("encryption_at_rest"))
}

type ConnectionPasswordEncryptionAttributes struct {
	ref terra.Reference
}

func (cpe ConnectionPasswordEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return cpe.ref, nil
}

func (cpe ConnectionPasswordEncryptionAttributes) InternalWithRef(ref terra.Reference) ConnectionPasswordEncryptionAttributes {
	return ConnectionPasswordEncryptionAttributes{ref: ref}
}

func (cpe ConnectionPasswordEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cpe.ref.InternalTokens()
}

func (cpe ConnectionPasswordEncryptionAttributes) AwsKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(cpe.ref.Append("aws_kms_key_id"))
}

func (cpe ConnectionPasswordEncryptionAttributes) ReturnConnectionPasswordEncrypted() terra.BoolValue {
	return terra.ReferenceAsBool(cpe.ref.Append("return_connection_password_encrypted"))
}

type EncryptionAtRestAttributes struct {
	ref terra.Reference
}

func (ear EncryptionAtRestAttributes) InternalRef() (terra.Reference, error) {
	return ear.ref, nil
}

func (ear EncryptionAtRestAttributes) InternalWithRef(ref terra.Reference) EncryptionAtRestAttributes {
	return EncryptionAtRestAttributes{ref: ref}
}

func (ear EncryptionAtRestAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ear.ref.InternalTokens()
}

func (ear EncryptionAtRestAttributes) CatalogEncryptionMode() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("catalog_encryption_mode"))
}

func (ear EncryptionAtRestAttributes) SseAwsKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(ear.ref.Append("sse_aws_kms_key_id"))
}

type DataCatalogEncryptionSettingsState struct {
	ConnectionPasswordEncryption []ConnectionPasswordEncryptionState `json:"connection_password_encryption"`
	EncryptionAtRest             []EncryptionAtRestState             `json:"encryption_at_rest"`
}

type ConnectionPasswordEncryptionState struct {
	AwsKmsKeyId                       string `json:"aws_kms_key_id"`
	ReturnConnectionPasswordEncrypted bool   `json:"return_connection_password_encrypted"`
}

type EncryptionAtRestState struct {
	CatalogEncryptionMode string `json:"catalog_encryption_mode"`
	SseAwsKmsKeyId        string `json:"sse_aws_kms_key_id"`
}
