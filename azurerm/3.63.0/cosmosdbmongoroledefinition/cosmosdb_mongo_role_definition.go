// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package cosmosdbmongoroledefinition

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Privilege struct {
	// Actions: list of string, required
	Actions terra.ListValue[terra.StringValue] `hcl:"actions,attr" validate:"required"`
	// Resource: required
	Resource *Resource `hcl:"resource,block" validate:"required"`
}

type Resource struct {
	// CollectionName: string, optional
	CollectionName terra.StringValue `hcl:"collection_name,attr"`
	// DbName: string, optional
	DbName terra.StringValue `hcl:"db_name,attr"`
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

type PrivilegeAttributes struct {
	ref terra.Reference
}

func (p PrivilegeAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PrivilegeAttributes) InternalWithRef(ref terra.Reference) PrivilegeAttributes {
	return PrivilegeAttributes{ref: ref}
}

func (p PrivilegeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PrivilegeAttributes) Actions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](p.ref.Append("actions"))
}

func (p PrivilegeAttributes) Resource() terra.ListValue[ResourceAttributes] {
	return terra.ReferenceAsList[ResourceAttributes](p.ref.Append("resource"))
}

type ResourceAttributes struct {
	ref terra.Reference
}

func (r ResourceAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResourceAttributes) InternalWithRef(ref terra.Reference) ResourceAttributes {
	return ResourceAttributes{ref: ref}
}

func (r ResourceAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResourceAttributes) CollectionName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("collection_name"))
}

func (r ResourceAttributes) DbName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("db_name"))
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

type PrivilegeState struct {
	Actions  []string        `json:"actions"`
	Resource []ResourceState `json:"resource"`
}

type ResourceState struct {
	CollectionName string `json:"collection_name"`
	DbName         string `json:"db_name"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
