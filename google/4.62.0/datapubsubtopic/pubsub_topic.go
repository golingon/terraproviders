// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datapubsubtopic

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MessageStoragePolicy struct{}

type SchemaSettings struct{}

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

type MessageStoragePolicyState struct {
	AllowedPersistenceRegions []string `json:"allowed_persistence_regions"`
}

type SchemaSettingsState struct {
	Encoding string `json:"encoding"`
	Schema   string `json:"schema"`
}
