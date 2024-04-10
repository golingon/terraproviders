// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datacdnfrontdoorcustomdomain

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Tls struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type TlsAttributes struct {
	ref terra.Reference
}

func (t TlsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TlsAttributes) InternalWithRef(ref terra.Reference) TlsAttributes {
	return TlsAttributes{ref: ref}
}

func (t TlsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TlsAttributes) CdnFrontdoorSecretId() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("cdn_frontdoor_secret_id"))
}

func (t TlsAttributes) CertificateType() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("certificate_type"))
}

func (t TlsAttributes) MinimumTlsVersion() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("minimum_tls_version"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type TlsState struct {
	CdnFrontdoorSecretId string `json:"cdn_frontdoor_secret_id"`
	CertificateType      string `json:"certificate_type"`
	MinimumTlsVersion    string `json:"minimum_tls_version"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
