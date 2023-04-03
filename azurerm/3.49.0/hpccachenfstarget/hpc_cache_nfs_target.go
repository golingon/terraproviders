// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package hpccachenfstarget

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type NamespaceJunction struct {
	// AccessPolicyName: string, optional
	AccessPolicyName terra.StringValue `hcl:"access_policy_name,attr"`
	// NamespacePath: string, required
	NamespacePath terra.StringValue `hcl:"namespace_path,attr" validate:"required"`
	// NfsExport: string, required
	NfsExport terra.StringValue `hcl:"nfs_export,attr" validate:"required"`
	// TargetPath: string, optional
	TargetPath terra.StringValue `hcl:"target_path,attr"`
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

type NamespaceJunctionAttributes struct {
	ref terra.Reference
}

func (nj NamespaceJunctionAttributes) InternalRef() (terra.Reference, error) {
	return nj.ref, nil
}

func (nj NamespaceJunctionAttributes) InternalWithRef(ref terra.Reference) NamespaceJunctionAttributes {
	return NamespaceJunctionAttributes{ref: ref}
}

func (nj NamespaceJunctionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return nj.ref.InternalTokens()
}

func (nj NamespaceJunctionAttributes) AccessPolicyName() terra.StringValue {
	return terra.ReferenceAsString(nj.ref.Append("access_policy_name"))
}

func (nj NamespaceJunctionAttributes) NamespacePath() terra.StringValue {
	return terra.ReferenceAsString(nj.ref.Append("namespace_path"))
}

func (nj NamespaceJunctionAttributes) NfsExport() terra.StringValue {
	return terra.ReferenceAsString(nj.ref.Append("nfs_export"))
}

func (nj NamespaceJunctionAttributes) TargetPath() terra.StringValue {
	return terra.ReferenceAsString(nj.ref.Append("target_path"))
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

type NamespaceJunctionState struct {
	AccessPolicyName string `json:"access_policy_name"`
	NamespacePath    string `json:"namespace_path"`
	NfsExport        string `json:"nfs_export"`
	TargetPath       string `json:"target_path"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
