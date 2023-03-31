// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package ecrrepository

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type EncryptionConfiguration struct {
	// EncryptionType: string, optional
	EncryptionType terra.StringValue `hcl:"encryption_type,attr"`
	// KmsKey: string, optional
	KmsKey terra.StringValue `hcl:"kms_key,attr"`
}

type ImageScanningConfiguration struct {
	// ScanOnPush: bool, required
	ScanOnPush terra.BoolValue `hcl:"scan_on_push,attr" validate:"required"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type EncryptionConfigurationAttributes struct {
	ref terra.Reference
}

func (ec EncryptionConfigurationAttributes) InternalRef() terra.Reference {
	return ec.ref
}

func (ec EncryptionConfigurationAttributes) InternalWithRef(ref terra.Reference) EncryptionConfigurationAttributes {
	return EncryptionConfigurationAttributes{ref: ref}
}

func (ec EncryptionConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ec.ref.InternalTokens()
}

func (ec EncryptionConfigurationAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_type"))
}

func (ec EncryptionConfigurationAttributes) KmsKey() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("kms_key"))
}

type ImageScanningConfigurationAttributes struct {
	ref terra.Reference
}

func (isc ImageScanningConfigurationAttributes) InternalRef() terra.Reference {
	return isc.ref
}

func (isc ImageScanningConfigurationAttributes) InternalWithRef(ref terra.Reference) ImageScanningConfigurationAttributes {
	return ImageScanningConfigurationAttributes{ref: ref}
}

func (isc ImageScanningConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return isc.ref.InternalTokens()
}

func (isc ImageScanningConfigurationAttributes) ScanOnPush() terra.BoolValue {
	return terra.ReferenceAsBool(isc.ref.Append("scan_on_push"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type EncryptionConfigurationState struct {
	EncryptionType string `json:"encryption_type"`
	KmsKey         string `json:"kms_key"`
}

type ImageScanningConfigurationState struct {
	ScanOnPush bool `json:"scan_on_push"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
}