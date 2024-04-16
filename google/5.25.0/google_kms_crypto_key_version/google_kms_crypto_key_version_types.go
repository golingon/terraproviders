// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_kms_crypto_key_version

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type ExternalProtectionLevelOptions struct {
	// EkmConnectionKeyPath: string, optional
	EkmConnectionKeyPath terra.StringValue `hcl:"ekm_connection_key_path,attr"`
	// ExternalKeyUri: string, optional
	ExternalKeyUri terra.StringValue `hcl:"external_key_uri,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AttestationAttributes struct {
	ref terra.Reference
}

func (a AttestationAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AttestationAttributes) InternalWithRef(ref terra.Reference) AttestationAttributes {
	return AttestationAttributes{ref: ref}
}

func (a AttestationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AttestationAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("content"))
}

func (a AttestationAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("format"))
}

func (a AttestationAttributes) CertChains() terra.ListValue[AttestationCertChainsAttributes] {
	return terra.ReferenceAsList[AttestationCertChainsAttributes](a.ref.Append("cert_chains"))
}

func (a AttestationAttributes) ExternalProtectionLevelOptions() terra.ListValue[AttestationExternalProtectionLevelOptionsAttributes] {
	return terra.ReferenceAsList[AttestationExternalProtectionLevelOptionsAttributes](a.ref.Append("external_protection_level_options"))
}

type AttestationCertChainsAttributes struct {
	ref terra.Reference
}

func (cc AttestationCertChainsAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc AttestationCertChainsAttributes) InternalWithRef(ref terra.Reference) AttestationCertChainsAttributes {
	return AttestationCertChainsAttributes{ref: ref}
}

func (cc AttestationCertChainsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc AttestationCertChainsAttributes) CaviumCerts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("cavium_certs"))
}

func (cc AttestationCertChainsAttributes) GoogleCardCerts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("google_card_certs"))
}

func (cc AttestationCertChainsAttributes) GooglePartitionCerts() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cc.ref.Append("google_partition_certs"))
}

type AttestationExternalProtectionLevelOptionsAttributes struct {
	ref terra.Reference
}

func (eplo AttestationExternalProtectionLevelOptionsAttributes) InternalRef() (terra.Reference, error) {
	return eplo.ref, nil
}

func (eplo AttestationExternalProtectionLevelOptionsAttributes) InternalWithRef(ref terra.Reference) AttestationExternalProtectionLevelOptionsAttributes {
	return AttestationExternalProtectionLevelOptionsAttributes{ref: ref}
}

func (eplo AttestationExternalProtectionLevelOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eplo.ref.InternalTokens()
}

func (eplo AttestationExternalProtectionLevelOptionsAttributes) EkmConnectionKeyPath() terra.StringValue {
	return terra.ReferenceAsString(eplo.ref.Append("ekm_connection_key_path"))
}

func (eplo AttestationExternalProtectionLevelOptionsAttributes) ExternalKeyUri() terra.StringValue {
	return terra.ReferenceAsString(eplo.ref.Append("external_key_uri"))
}

type ExternalProtectionLevelOptionsAttributes struct {
	ref terra.Reference
}

func (eplo ExternalProtectionLevelOptionsAttributes) InternalRef() (terra.Reference, error) {
	return eplo.ref, nil
}

func (eplo ExternalProtectionLevelOptionsAttributes) InternalWithRef(ref terra.Reference) ExternalProtectionLevelOptionsAttributes {
	return ExternalProtectionLevelOptionsAttributes{ref: ref}
}

func (eplo ExternalProtectionLevelOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return eplo.ref.InternalTokens()
}

func (eplo ExternalProtectionLevelOptionsAttributes) EkmConnectionKeyPath() terra.StringValue {
	return terra.ReferenceAsString(eplo.ref.Append("ekm_connection_key_path"))
}

func (eplo ExternalProtectionLevelOptionsAttributes) ExternalKeyUri() terra.StringValue {
	return terra.ReferenceAsString(eplo.ref.Append("external_key_uri"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AttestationState struct {
	Content                        string                                           `json:"content"`
	Format                         string                                           `json:"format"`
	CertChains                     []AttestationCertChainsState                     `json:"cert_chains"`
	ExternalProtectionLevelOptions []AttestationExternalProtectionLevelOptionsState `json:"external_protection_level_options"`
}

type AttestationCertChainsState struct {
	CaviumCerts          []string `json:"cavium_certs"`
	GoogleCardCerts      []string `json:"google_card_certs"`
	GooglePartitionCerts []string `json:"google_partition_certs"`
}

type AttestationExternalProtectionLevelOptionsState struct {
	EkmConnectionKeyPath string `json:"ekm_connection_key_path"`
	ExternalKeyUri       string `json:"external_key_uri"`
}

type ExternalProtectionLevelOptionsState struct {
	EkmConnectionKeyPath string `json:"ekm_connection_key_path"`
	ExternalKeyUri       string `json:"external_key_uri"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
