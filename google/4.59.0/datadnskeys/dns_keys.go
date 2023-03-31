// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadnskeys

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type KeySigningKeys struct {
	// KeySigningKeysDigests: min=0
	Digests []KeySigningKeysDigests `hcl:"digests,block" validate:"min=0"`
}

type KeySigningKeysDigests struct{}

type ZoneSigningKeys struct {
	// ZoneSigningKeysDigests: min=0
	Digests []ZoneSigningKeysDigests `hcl:"digests,block" validate:"min=0"`
}

type ZoneSigningKeysDigests struct{}

type KeySigningKeysAttributes struct {
	ref terra.Reference
}

func (ksk KeySigningKeysAttributes) InternalRef() terra.Reference {
	return ksk.ref
}

func (ksk KeySigningKeysAttributes) InternalWithRef(ref terra.Reference) KeySigningKeysAttributes {
	return KeySigningKeysAttributes{ref: ref}
}

func (ksk KeySigningKeysAttributes) InternalTokens() hclwrite.Tokens {
	return ksk.ref.InternalTokens()
}

func (ksk KeySigningKeysAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("algorithm"))
}

func (ksk KeySigningKeysAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("creation_time"))
}

func (ksk KeySigningKeysAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("description"))
}

func (ksk KeySigningKeysAttributes) DsRecord() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("ds_record"))
}

func (ksk KeySigningKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("id"))
}

func (ksk KeySigningKeysAttributes) IsActive() terra.BoolValue {
	return terra.ReferenceBool(ksk.ref.Append("is_active"))
}

func (ksk KeySigningKeysAttributes) KeyLength() terra.NumberValue {
	return terra.ReferenceNumber(ksk.ref.Append("key_length"))
}

func (ksk KeySigningKeysAttributes) KeyTag() terra.NumberValue {
	return terra.ReferenceNumber(ksk.ref.Append("key_tag"))
}

func (ksk KeySigningKeysAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceString(ksk.ref.Append("public_key"))
}

func (ksk KeySigningKeysAttributes) Digests() terra.ListValue[KeySigningKeysDigestsAttributes] {
	return terra.ReferenceList[KeySigningKeysDigestsAttributes](ksk.ref.Append("digests"))
}

type KeySigningKeysDigestsAttributes struct {
	ref terra.Reference
}

func (d KeySigningKeysDigestsAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d KeySigningKeysDigestsAttributes) InternalWithRef(ref terra.Reference) KeySigningKeysDigestsAttributes {
	return KeySigningKeysDigestsAttributes{ref: ref}
}

func (d KeySigningKeysDigestsAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d KeySigningKeysDigestsAttributes) Digest() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("digest"))
}

func (d KeySigningKeysDigestsAttributes) Type() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("type"))
}

type ZoneSigningKeysAttributes struct {
	ref terra.Reference
}

func (zsk ZoneSigningKeysAttributes) InternalRef() terra.Reference {
	return zsk.ref
}

func (zsk ZoneSigningKeysAttributes) InternalWithRef(ref terra.Reference) ZoneSigningKeysAttributes {
	return ZoneSigningKeysAttributes{ref: ref}
}

func (zsk ZoneSigningKeysAttributes) InternalTokens() hclwrite.Tokens {
	return zsk.ref.InternalTokens()
}

func (zsk ZoneSigningKeysAttributes) Algorithm() terra.StringValue {
	return terra.ReferenceString(zsk.ref.Append("algorithm"))
}

func (zsk ZoneSigningKeysAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceString(zsk.ref.Append("creation_time"))
}

func (zsk ZoneSigningKeysAttributes) Description() terra.StringValue {
	return terra.ReferenceString(zsk.ref.Append("description"))
}

func (zsk ZoneSigningKeysAttributes) Id() terra.StringValue {
	return terra.ReferenceString(zsk.ref.Append("id"))
}

func (zsk ZoneSigningKeysAttributes) IsActive() terra.BoolValue {
	return terra.ReferenceBool(zsk.ref.Append("is_active"))
}

func (zsk ZoneSigningKeysAttributes) KeyLength() terra.NumberValue {
	return terra.ReferenceNumber(zsk.ref.Append("key_length"))
}

func (zsk ZoneSigningKeysAttributes) KeyTag() terra.NumberValue {
	return terra.ReferenceNumber(zsk.ref.Append("key_tag"))
}

func (zsk ZoneSigningKeysAttributes) PublicKey() terra.StringValue {
	return terra.ReferenceString(zsk.ref.Append("public_key"))
}

func (zsk ZoneSigningKeysAttributes) Digests() terra.ListValue[ZoneSigningKeysDigestsAttributes] {
	return terra.ReferenceList[ZoneSigningKeysDigestsAttributes](zsk.ref.Append("digests"))
}

type ZoneSigningKeysDigestsAttributes struct {
	ref terra.Reference
}

func (d ZoneSigningKeysDigestsAttributes) InternalRef() terra.Reference {
	return d.ref
}

func (d ZoneSigningKeysDigestsAttributes) InternalWithRef(ref terra.Reference) ZoneSigningKeysDigestsAttributes {
	return ZoneSigningKeysDigestsAttributes{ref: ref}
}

func (d ZoneSigningKeysDigestsAttributes) InternalTokens() hclwrite.Tokens {
	return d.ref.InternalTokens()
}

func (d ZoneSigningKeysDigestsAttributes) Digest() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("digest"))
}

func (d ZoneSigningKeysDigestsAttributes) Type() terra.StringValue {
	return terra.ReferenceString(d.ref.Append("type"))
}

type KeySigningKeysState struct {
	Algorithm    string                       `json:"algorithm"`
	CreationTime string                       `json:"creation_time"`
	Description  string                       `json:"description"`
	DsRecord     string                       `json:"ds_record"`
	Id           string                       `json:"id"`
	IsActive     bool                         `json:"is_active"`
	KeyLength    float64                      `json:"key_length"`
	KeyTag       float64                      `json:"key_tag"`
	PublicKey    string                       `json:"public_key"`
	Digests      []KeySigningKeysDigestsState `json:"digests"`
}

type KeySigningKeysDigestsState struct {
	Digest string `json:"digest"`
	Type   string `json:"type"`
}

type ZoneSigningKeysState struct {
	Algorithm    string                        `json:"algorithm"`
	CreationTime string                        `json:"creation_time"`
	Description  string                        `json:"description"`
	Id           string                        `json:"id"`
	IsActive     bool                          `json:"is_active"`
	KeyLength    float64                       `json:"key_length"`
	KeyTag       float64                       `json:"key_tag"`
	PublicKey    string                        `json:"public_key"`
	Digests      []ZoneSigningKeysDigestsState `json:"digests"`
}

type ZoneSigningKeysDigestsState struct {
	Digest string `json:"digest"`
	Type   string `json:"type"`
}
