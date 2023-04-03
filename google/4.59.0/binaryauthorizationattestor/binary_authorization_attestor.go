// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package binaryauthorizationattestor

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AttestationAuthorityNote struct {
	// NoteReference: string, required
	NoteReference terra.StringValue `hcl:"note_reference,attr" validate:"required"`
	// PublicKeys: min=0
	PublicKeys []PublicKeys `hcl:"public_keys,block" validate:"min=0"`
}

type PublicKeys struct {
	// AsciiArmoredPgpPublicKey: string, optional
	AsciiArmoredPgpPublicKey terra.StringValue `hcl:"ascii_armored_pgp_public_key,attr"`
	// Comment: string, optional
	Comment terra.StringValue `hcl:"comment,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PkixPublicKey: optional
	PkixPublicKey *PkixPublicKey `hcl:"pkix_public_key,block"`
}

type PkixPublicKey struct {
	// PublicKeyPem: string, optional
	PublicKeyPem terra.StringValue `hcl:"public_key_pem,attr"`
	// SignatureAlgorithm: string, optional
	SignatureAlgorithm terra.StringValue `hcl:"signature_algorithm,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AttestationAuthorityNoteAttributes struct {
	ref terra.Reference
}

func (aan AttestationAuthorityNoteAttributes) InternalRef() (terra.Reference, error) {
	return aan.ref, nil
}

func (aan AttestationAuthorityNoteAttributes) InternalWithRef(ref terra.Reference) AttestationAuthorityNoteAttributes {
	return AttestationAuthorityNoteAttributes{ref: ref}
}

func (aan AttestationAuthorityNoteAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aan.ref.InternalTokens()
}

func (aan AttestationAuthorityNoteAttributes) DelegationServiceAccountEmail() terra.StringValue {
	return terra.ReferenceAsString(aan.ref.Append("delegation_service_account_email"))
}

func (aan AttestationAuthorityNoteAttributes) NoteReference() terra.StringValue {
	return terra.ReferenceAsString(aan.ref.Append("note_reference"))
}

func (aan AttestationAuthorityNoteAttributes) PublicKeys() terra.ListValue[PublicKeysAttributes] {
	return terra.ReferenceAsList[PublicKeysAttributes](aan.ref.Append("public_keys"))
}

type PublicKeysAttributes struct {
	ref terra.Reference
}

func (pk PublicKeysAttributes) InternalRef() (terra.Reference, error) {
	return pk.ref, nil
}

func (pk PublicKeysAttributes) InternalWithRef(ref terra.Reference) PublicKeysAttributes {
	return PublicKeysAttributes{ref: ref}
}

func (pk PublicKeysAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pk.ref.InternalTokens()
}

func (pk PublicKeysAttributes) AsciiArmoredPgpPublicKey() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("ascii_armored_pgp_public_key"))
}

func (pk PublicKeysAttributes) Comment() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("comment"))
}

func (pk PublicKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("id"))
}

func (pk PublicKeysAttributes) PkixPublicKey() terra.ListValue[PkixPublicKeyAttributes] {
	return terra.ReferenceAsList[PkixPublicKeyAttributes](pk.ref.Append("pkix_public_key"))
}

type PkixPublicKeyAttributes struct {
	ref terra.Reference
}

func (ppk PkixPublicKeyAttributes) InternalRef() (terra.Reference, error) {
	return ppk.ref, nil
}

func (ppk PkixPublicKeyAttributes) InternalWithRef(ref terra.Reference) PkixPublicKeyAttributes {
	return PkixPublicKeyAttributes{ref: ref}
}

func (ppk PkixPublicKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ppk.ref.InternalTokens()
}

func (ppk PkixPublicKeyAttributes) PublicKeyPem() terra.StringValue {
	return terra.ReferenceAsString(ppk.ref.Append("public_key_pem"))
}

func (ppk PkixPublicKeyAttributes) SignatureAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(ppk.ref.Append("signature_algorithm"))
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

type AttestationAuthorityNoteState struct {
	DelegationServiceAccountEmail string            `json:"delegation_service_account_email"`
	NoteReference                 string            `json:"note_reference"`
	PublicKeys                    []PublicKeysState `json:"public_keys"`
}

type PublicKeysState struct {
	AsciiArmoredPgpPublicKey string               `json:"ascii_armored_pgp_public_key"`
	Comment                  string               `json:"comment"`
	Id                       string               `json:"id"`
	PkixPublicKey            []PkixPublicKeyState `json:"pkix_public_key"`
}

type PkixPublicKeyState struct {
	PublicKeyPem       string `json:"public_key_pem"`
	SignatureAlgorithm string `json:"signature_algorithm"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}