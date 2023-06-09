// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cdnfrontdoorcustomdomain

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type Tls struct {
	// CdnFrontdoorSecretId: string, optional
	CdnFrontdoorSecretId terra.StringValue `hcl:"cdn_frontdoor_secret_id,attr"`
	// CertificateType: string, optional
	CertificateType terra.StringValue `hcl:"certificate_type,attr"`
	// MinimumTlsVersion: string, optional
	MinimumTlsVersion terra.StringValue `hcl:"minimum_tls_version,attr"`
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
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

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}

type TlsState struct {
	CdnFrontdoorSecretId string `json:"cdn_frontdoor_secret_id"`
	CertificateType      string `json:"certificate_type"`
	MinimumTlsVersion    string `json:"minimum_tls_version"`
}
