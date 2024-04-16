// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ecr_repository

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataEncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec DataEncryptionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec DataEncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) DataEncryptionConfigurationAttributes {
	return DataEncryptionConfigurationAttributes{ref: ref}
}

func (ec DataEncryptionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec DataEncryptionConfigurationAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_type"))
}

func (ec DataEncryptionConfigurationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type DataImageScanningConfigurationAttributes struct {
	ref terra.Reference
}

func (isc DataImageScanningConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return isc.ref, nil
}

func (isc DataImageScanningConfigurationAttributes) InternalWithRef(ref terra.Reference) DataImageScanningConfigurationAttributes {
	return DataImageScanningConfigurationAttributes{ref: ref}
}

func (isc DataImageScanningConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return isc.ref.InternalTokens()
}

func (isc DataImageScanningConfigurationAttributes) ScanOnPush() terra.BoolValue {
	return terra.ReferenceAsBool(isc.ref.Append("scan_on_push"))
}

type DataEncryptionConfigurationState struct {
	EncryptionType string `json:"encryption_type"`
	KmsKey         string `json:"kms_key"`
}

type DataImageScanningConfigurationState struct {
	ScanOnPush bool `json:"scan_on_push"`
}
