// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package gluepartitionindex

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type PartitionIndex struct {
	// IndexName: string, optional
	IndexName terra.StringValue `hcl:"index_name,attr"`
	// Keys: list of string, optional
	Keys terra.ListValue[terra.StringValue] `hcl:"keys,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type PartitionIndexAttributes struct {
	ref terra.Reference
}

func (pi PartitionIndexAttributes) InternalRef() (terra.Reference, error) {
	return pi.ref, nil
}

func (pi PartitionIndexAttributes) InternalWithRef(ref terra.Reference) PartitionIndexAttributes {
	return PartitionIndexAttributes{ref: ref}
}

func (pi PartitionIndexAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return pi.ref.InternalTokens()
}

func (pi PartitionIndexAttributes) IndexName() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("index_name"))
}

func (pi PartitionIndexAttributes) IndexStatus() terra.StringValue {
	return terra.ReferenceAsString(pi.ref.Append("index_status"))
}

func (pi PartitionIndexAttributes) Keys() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pi.ref.Append("keys"))
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

type PartitionIndexState struct {
	IndexName   string   `json:"index_name"`
	IndexStatus string   `json:"index_status"`
	Keys        []string `json:"keys"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
}
