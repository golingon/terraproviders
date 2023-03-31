// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package rdsglobalcluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type GlobalClusterMembers struct{}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type GlobalClusterMembersAttributes struct {
	ref terra.Reference
}

func (gcm GlobalClusterMembersAttributes) InternalRef() terra.Reference {
	return gcm.ref
}

func (gcm GlobalClusterMembersAttributes) InternalWithRef(ref terra.Reference) GlobalClusterMembersAttributes {
	return GlobalClusterMembersAttributes{ref: ref}
}

func (gcm GlobalClusterMembersAttributes) InternalTokens() hclwrite.Tokens {
	return gcm.ref.InternalTokens()
}

func (gcm GlobalClusterMembersAttributes) DbClusterArn() terra.StringValue {
	return terra.ReferenceAsString(gcm.ref.Append("db_cluster_arn"))
}

func (gcm GlobalClusterMembersAttributes) IsWriter() terra.BoolValue {
	return terra.ReferenceAsBool(gcm.ref.Append("is_writer"))
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

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type GlobalClusterMembersState struct {
	DbClusterArn string `json:"db_cluster_arn"`
	IsWriter     bool   `json:"is_writer"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
