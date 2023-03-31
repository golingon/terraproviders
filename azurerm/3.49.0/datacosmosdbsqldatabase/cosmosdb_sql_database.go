// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacosmosdbsqldatabase

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type AutoscaleSettings struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type AutoscaleSettingsAttributes struct {
	ref terra.Reference
}

func (as AutoscaleSettingsAttributes) InternalRef() terra.Reference {
	return as.ref
}

func (as AutoscaleSettingsAttributes) InternalWithRef(ref terra.Reference) AutoscaleSettingsAttributes {
	return AutoscaleSettingsAttributes{ref: ref}
}

func (as AutoscaleSettingsAttributes) InternalTokens() hclwrite.Tokens {
	return as.ref.InternalTokens()
}

func (as AutoscaleSettingsAttributes) MaxThroughput() terra.NumberValue {
	return terra.ReferenceNumber(as.ref.Append("max_throughput"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

type AutoscaleSettingsState struct {
	MaxThroughput float64 `json:"max_throughput"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
