// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package cosmosdbtable

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AutoscaleSettings struct {
	// MaxThroughput: number, optional
	MaxThroughput terra.NumberValue `hcl:"max_throughput,attr"`
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

type AutoscaleSettingsState struct {
	MaxThroughput float64 `json:"max_throughput"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
