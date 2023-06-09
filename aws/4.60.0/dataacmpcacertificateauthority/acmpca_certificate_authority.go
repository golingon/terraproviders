// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataacmpcacertificateauthority

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RevocationConfiguration struct {
	// CrlConfiguration: min=0
	CrlConfiguration []CrlConfiguration `hcl:"crl_configuration,block" validate:"min=0"`
	// OcspConfiguration: min=0
	OcspConfiguration []OcspConfiguration `hcl:"ocsp_configuration,block" validate:"min=0"`
}

type CrlConfiguration struct{}

type OcspConfiguration struct{}

type RevocationConfigurationAttributes struct {
	ref terra.Reference
}

func (rc RevocationConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return rc.ref, nil
}

func (rc RevocationConfigurationAttributes) InternalWithRef(ref terra.Reference) RevocationConfigurationAttributes {
	return RevocationConfigurationAttributes{ref: ref}
}

func (rc RevocationConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rc.ref.InternalTokens()
}

func (rc RevocationConfigurationAttributes) CrlConfiguration() terra.ListValue[CrlConfigurationAttributes] {
	return terra.ReferenceAsList[CrlConfigurationAttributes](rc.ref.Append("crl_configuration"))
}

func (rc RevocationConfigurationAttributes) OcspConfiguration() terra.ListValue[OcspConfigurationAttributes] {
	return terra.ReferenceAsList[OcspConfigurationAttributes](rc.ref.Append("ocsp_configuration"))
}

type CrlConfigurationAttributes struct {
	ref terra.Reference
}

func (cc CrlConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CrlConfigurationAttributes) InternalWithRef(ref terra.Reference) CrlConfigurationAttributes {
	return CrlConfigurationAttributes{ref: ref}
}

func (cc CrlConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CrlConfigurationAttributes) CustomCname() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("custom_cname"))
}

func (cc CrlConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enabled"))
}

func (cc CrlConfigurationAttributes) ExpirationInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(cc.ref.Append("expiration_in_days"))
}

func (cc CrlConfigurationAttributes) S3BucketName() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("s3_bucket_name"))
}

func (cc CrlConfigurationAttributes) S3ObjectAcl() terra.StringValue {
	return terra.ReferenceAsString(cc.ref.Append("s3_object_acl"))
}

type OcspConfigurationAttributes struct {
	ref terra.Reference
}

func (oc OcspConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return oc.ref, nil
}

func (oc OcspConfigurationAttributes) InternalWithRef(ref terra.Reference) OcspConfigurationAttributes {
	return OcspConfigurationAttributes{ref: ref}
}

func (oc OcspConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return oc.ref.InternalTokens()
}

func (oc OcspConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(oc.ref.Append("enabled"))
}

func (oc OcspConfigurationAttributes) OcspCustomCname() terra.StringValue {
	return terra.ReferenceAsString(oc.ref.Append("ocsp_custom_cname"))
}

type RevocationConfigurationState struct {
	CrlConfiguration  []CrlConfigurationState  `json:"crl_configuration"`
	OcspConfiguration []OcspConfigurationState `json:"ocsp_configuration"`
}

type CrlConfigurationState struct {
	CustomCname      string  `json:"custom_cname"`
	Enabled          bool    `json:"enabled"`
	ExpirationInDays float64 `json:"expiration_in_days"`
	S3BucketName     string  `json:"s3_bucket_name"`
	S3ObjectAcl      string  `json:"s3_object_acl"`
}

type OcspConfigurationState struct {
	Enabled         bool   `json:"enabled"`
	OcspCustomCname string `json:"ocsp_custom_cname"`
}
