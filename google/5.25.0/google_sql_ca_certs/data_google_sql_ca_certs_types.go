// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_sql_ca_certs

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataCertsAttributes struct {
	ref terra.Reference
}

func (c DataCertsAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataCertsAttributes) InternalWithRef(ref terra.Reference) DataCertsAttributes {
	return DataCertsAttributes{ref: ref}
}

func (c DataCertsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataCertsAttributes) Cert() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("cert"))
}

func (c DataCertsAttributes) CommonName() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("common_name"))
}

func (c DataCertsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("create_time"))
}

func (c DataCertsAttributes) ExpirationTime() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("expiration_time"))
}

func (c DataCertsAttributes) Sha1Fingerprint() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("sha1_fingerprint"))
}

type DataCertsState struct {
	Cert            string `json:"cert"`
	CommonName      string `json:"common_name"`
	CreateTime      string `json:"create_time"`
	ExpirationTime  string `json:"expiration_time"`
	Sha1Fingerprint string `json:"sha1_fingerprint"`
}
