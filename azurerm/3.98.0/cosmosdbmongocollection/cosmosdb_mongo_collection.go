// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cosmosdbmongocollection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type SystemIndexes struct{}

type AutoscaleSettings struct {
	// MaxThroughput: number, optional
	MaxThroughput terra.NumberValue `hcl:"max_throughput,attr"`
}

type Index struct {
	// Keys: list of string, required
	Keys terra.ListValue[terra.StringValue] `hcl:"keys,attr" validate:"required"`
	// Unique: bool, optional
	Unique terra.BoolValue `hcl:"unique,attr"`
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

type SystemIndexesAttributes struct {
	ref terra.Reference
}

func (si SystemIndexesAttributes) InternalRef() (terra.Reference, error) {
	return si.ref, nil
}

func (si SystemIndexesAttributes) InternalWithRef(ref terra.Reference) SystemIndexesAttributes {
	return SystemIndexesAttributes{ref: ref}
}

func (si SystemIndexesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return si.ref.InternalTokens()
}

func (si SystemIndexesAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](si.ref.Append("keys"))
}

func (si SystemIndexesAttributes) Unique() terra.BoolValue {
	return terra.ReferenceAsBool(si.ref.Append("unique"))
}

type AutoscaleSettingsAttributes struct {
	ref terra.Reference
}

func (as AutoscaleSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AutoscaleSettingsAttributes) InternalWithRef(ref terra.Reference) AutoscaleSettingsAttributes {
	return AutoscaleSettingsAttributes{ref: ref}
}

func (as AutoscaleSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AutoscaleSettingsAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceAsNumber(as.ref.Append("max_throughput"))
}

type IndexAttributes struct {
	ref terra.Reference
}

func (i IndexAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i IndexAttributes) InternalWithRef(ref terra.Reference) IndexAttributes {
	return IndexAttributes{ref: ref}
}

func (i IndexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i IndexAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("keys"))
}

func (i IndexAttributes) Unique() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("unique"))
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

type SystemIndexesState struct {
	Keys   []string `json:"keys"`
	Unique bool     `json:"unique"`
}

type AutoscaleSettingsState struct {
	MaxThroughput float64 `json:"max_throughput"`
}

type IndexState struct {
	Keys   []string `json:"keys"`
	Unique bool     `json:"unique"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
