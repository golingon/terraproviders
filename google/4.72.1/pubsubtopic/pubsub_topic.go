// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package pubsubtopic

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MessageStoragePolicy struct {
	// AllowedPersistenceRegions: list of string, required
	AllowedPersistenceRegions terra.ListValue[terra.StringValue] `hcl:"allowed_persistence_regions,attr" validate:"required"`
}

type SchemaSettings struct {
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// Schema: string, required
	Schema terra.StringValue `hcl:"schema,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type MessageStoragePolicyAttributes struct {
	ref terra.Reference
}

func (msp MessageStoragePolicyAttributes) InternalRef() (terra.Reference, error) {
	return msp.ref, nil
}

func (msp MessageStoragePolicyAttributes) InternalWithRef(ref terra.Reference) MessageStoragePolicyAttributes {
	return MessageStoragePolicyAttributes{ref: ref}
}

func (msp MessageStoragePolicyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return msp.ref.InternalTokens()
}

func (msp MessageStoragePolicyAttributes) AllowedPersistenceRegions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](msp.ref.Append("allowed_persistence_regions"))
}

type SchemaSettingsAttributes struct {
	ref terra.Reference
}

func (ss SchemaSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss SchemaSettingsAttributes) InternalWithRef(ref terra.Reference) SchemaSettingsAttributes {
	return SchemaSettingsAttributes{ref: ref}
}

func (ss SchemaSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss SchemaSettingsAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("encoding"))
}

func (ss SchemaSettingsAttributes) Schema() terra.StringValue {
	return terra.ReferenceAsString(ss.ref.Append("schema"))
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

type MessageStoragePolicyState struct {
	AllowedPersistenceRegions []string `json:"allowed_persistence_regions"`
}

type SchemaSettingsState struct {
	Encoding string `json:"encoding"`
	Schema   string `json:"schema"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}