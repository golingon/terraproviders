// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_app_engine_service_split_traffic

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Split struct {
	// Allocations: map of string, required
	Allocations terra.MapValue[terra.StringValue] `hcl:"allocations,attr" validate:"required"`
	// ShardBy: string, optional
	ShardBy terra.StringValue `hcl:"shard_by,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type SplitAttributes struct {
	ref terra.Reference
}

func (s SplitAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SplitAttributes) InternalWithRef(ref terra.Reference) SplitAttributes {
	return SplitAttributes{ref: ref}
}

func (s SplitAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SplitAttributes) Allocations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("allocations"))
}

func (s SplitAttributes) ShardBy() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("shard_by"))
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

type SplitState struct {
	Allocations map[string]string `json:"allocations"`
	ShardBy     string            `json:"shard_by"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
