// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datakmskey

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type MultiRegionConfiguration struct {
	// PrimaryKey: min=0
	PrimaryKey []PrimaryKey `hcl:"primary_key,block" validate:"min=0"`
	// ReplicaKeys: min=0
	ReplicaKeys []ReplicaKeys `hcl:"replica_keys,block" validate:"min=0"`
}

type PrimaryKey struct{}

type ReplicaKeys struct{}

type XksKeyConfiguration struct{}

type MultiRegionConfigurationAttributes struct {
	ref terra.Reference
}

func (mrc MultiRegionConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return mrc.ref, nil
}

func (mrc MultiRegionConfigurationAttributes) InternalWithRef(ref terra.Reference) MultiRegionConfigurationAttributes {
	return MultiRegionConfigurationAttributes{ref: ref}
}

func (mrc MultiRegionConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mrc.ref.InternalTokens()
}

func (mrc MultiRegionConfigurationAttributes) MultiRegionKeyType() terra.StringValue {
	return terra.ReferenceAsString(mrc.ref.Append("multi_region_key_type"))
}

func (mrc MultiRegionConfigurationAttributes) PrimaryKey() terra.ListValue[PrimaryKeyAttributes] {
	return terra.ReferenceAsList[PrimaryKeyAttributes](mrc.ref.Append("primary_key"))
}

func (mrc MultiRegionConfigurationAttributes) ReplicaKeys() terra.ListValue[ReplicaKeysAttributes] {
	return terra.ReferenceAsList[ReplicaKeysAttributes](mrc.ref.Append("replica_keys"))
}

type PrimaryKeyAttributes struct {
	ref terra.Reference
}

func (pk PrimaryKeyAttributes) InternalRef() (terra.Reference, error) {
	return pk.ref, nil
}

func (pk PrimaryKeyAttributes) InternalWithRef(ref terra.Reference) PrimaryKeyAttributes {
	return PrimaryKeyAttributes{ref: ref}
}

func (pk PrimaryKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pk.ref.InternalTokens()
}

func (pk PrimaryKeyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("arn"))
}

func (pk PrimaryKeyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(pk.ref.Append("region"))
}

type ReplicaKeysAttributes struct {
	ref terra.Reference
}

func (rk ReplicaKeysAttributes) InternalRef() (terra.Reference, error) {
	return rk.ref, nil
}

func (rk ReplicaKeysAttributes) InternalWithRef(ref terra.Reference) ReplicaKeysAttributes {
	return ReplicaKeysAttributes{ref: ref}
}

func (rk ReplicaKeysAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rk.ref.InternalTokens()
}

func (rk ReplicaKeysAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rk.ref.Append("arn"))
}

func (rk ReplicaKeysAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(rk.ref.Append("region"))
}

type XksKeyConfigurationAttributes struct {
	ref terra.Reference
}

func (xkc XksKeyConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return xkc.ref, nil
}

func (xkc XksKeyConfigurationAttributes) InternalWithRef(ref terra.Reference) XksKeyConfigurationAttributes {
	return XksKeyConfigurationAttributes{ref: ref}
}

func (xkc XksKeyConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return xkc.ref.InternalTokens()
}

func (xkc XksKeyConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(xkc.ref.Append("id"))
}

type MultiRegionConfigurationState struct {
	MultiRegionKeyType string             `json:"multi_region_key_type"`
	PrimaryKey         []PrimaryKeyState  `json:"primary_key"`
	ReplicaKeys        []ReplicaKeysState `json:"replica_keys"`
}

type PrimaryKeyState struct {
	Arn    string `json:"arn"`
	Region string `json:"region"`
}

type ReplicaKeysState struct {
	Arn    string `json:"arn"`
	Region string `json:"region"`
}

type XksKeyConfigurationState struct {
	Id string `json:"id"`
}
