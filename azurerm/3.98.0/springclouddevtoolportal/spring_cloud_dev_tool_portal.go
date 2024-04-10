// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package springclouddevtoolportal

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Sso struct {
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// ClientSecret: string, optional
	ClientSecret terra.StringValue `hcl:"client_secret,attr"`
	// MetadataUrl: string, optional
	MetadataUrl terra.StringValue `hcl:"metadata_url,attr"`
	// Scope: set of string, optional
	Scope terra.SetValue[terra.StringValue] `hcl:"scope,attr"`
}

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

type SsoAttributes struct {
	ref terra.Reference
}

func (s SsoAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SsoAttributes) InternalWithRef(ref terra.Reference) SsoAttributes {
	return SsoAttributes{ref: ref}
}

func (s SsoAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SsoAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("client_id"))
}

func (s SsoAttributes) ClientSecret() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("client_secret"))
}

func (s SsoAttributes) MetadataUrl() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("metadata_url"))
}

func (s SsoAttributes) Scope() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](s.ref.Append("scope"))
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

type SsoState struct {
	ClientId     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	MetadataUrl  string   `json:"metadata_url"`
	Scope        []string `json:"scope"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
