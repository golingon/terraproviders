// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_cdn_frontdoor_custom_domain

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataTimeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type DataTlsAttributes struct {
	ref terra.Reference
}

func (t DataTlsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTlsAttributes) InternalWithRef(ref terra.Reference) DataTlsAttributes {
	return DataTlsAttributes{ref: ref}
}

func (t DataTlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTlsAttributes) CdnFrontdoorSecretId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("cdn_frontdoor_secret_id"))
}

func (t DataTlsAttributes) CertificateType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("certificate_type"))
}

func (t DataTlsAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("minimum_tls_version"))
}

type DataTimeoutsAttributes struct {
	ref terra.Reference
}

func (t DataTimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t DataTimeoutsAttributes) InternalWithRef(ref terra.Reference) DataTimeoutsAttributes {
	return DataTimeoutsAttributes{ref: ref}
}

func (t DataTimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t DataTimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type DataTlsState struct {
	CdnFrontdoorSecretId string `json:"cdn_frontdoor_secret_id"`
	CertificateType      string `json:"certificate_type"`
	MinimumTlsVersion    string `json:"minimum_tls_version"`
}

type DataTimeoutsState struct {
	Read string `json:"read"`
}
