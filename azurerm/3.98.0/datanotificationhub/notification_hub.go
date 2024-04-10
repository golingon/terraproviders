// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datanotificationhub

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ApnsCredential struct{}

type GcmCredential struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type ApnsCredentialAttributes struct {
	ref terra.Reference
}

func (ac ApnsCredentialAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac ApnsCredentialAttributes) InternalWithRef(ref terra.Reference) ApnsCredentialAttributes {
	return ApnsCredentialAttributes{ref: ref}
}

func (ac ApnsCredentialAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac ApnsCredentialAttributes) ApplicationMode() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("application_mode"))
}

func (ac ApnsCredentialAttributes) BundleId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("bundle_id"))
}

func (ac ApnsCredentialAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("key_id"))
}

func (ac ApnsCredentialAttributes) TeamId() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("team_id"))
}

func (ac ApnsCredentialAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("token"))
}

type GcmCredentialAttributes struct {
	ref terra.Reference
}

func (gc GcmCredentialAttributes) InternalRef() (terra.Reference, error) {
	return gc.ref, nil
}

func (gc GcmCredentialAttributes) InternalWithRef(ref terra.Reference) GcmCredentialAttributes {
	return GcmCredentialAttributes{ref: ref}
}

func (gc GcmCredentialAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gc.ref.InternalTokens()
}

func (gc GcmCredentialAttributes) ApiKey() terra.StringValue {
	return terra.ReferenceAsString(gc.ref.Append("api_key"))
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

type ApnsCredentialState struct {
	ApplicationMode string `json:"application_mode"`
	BundleId        string `json:"bundle_id"`
	KeyId           string `json:"key_id"`
	TeamId          string `json:"team_id"`
	Token           string `json:"token"`
}

type GcmCredentialState struct {
	ApiKey string `json:"api_key"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
