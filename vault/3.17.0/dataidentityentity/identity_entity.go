// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataidentityentity

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Aliases struct{}

type AliasesAttributes struct {
	ref terra.Reference
}

func (a AliasesAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AliasesAttributes) InternalWithRef(ref terra.Reference) AliasesAttributes {
	return AliasesAttributes{ref: ref}
}

func (a AliasesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return a.ref.InternalTokens()
}

func (a AliasesAttributes) CanonicalId() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("canonical_id"))
}

func (a AliasesAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("creation_time"))
}

func (a AliasesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("id"))
}

func (a AliasesAttributes) LastUpdateTime() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("last_update_time"))
}

func (a AliasesAttributes) MergedFromCanonicalIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](a.ref.Append("merged_from_canonical_ids"))
}

func (a AliasesAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](a.ref.Append("metadata"))
}

func (a AliasesAttributes) MountAccessor() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("mount_accessor"))
}

func (a AliasesAttributes) MountPath() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("mount_path"))
}

func (a AliasesAttributes) MountType() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("mount_type"))
}

func (a AliasesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("name"))
}

type AliasesState struct {
	CanonicalId            string            `json:"canonical_id"`
	CreationTime           string            `json:"creation_time"`
	Id                     string            `json:"id"`
	LastUpdateTime         string            `json:"last_update_time"`
	MergedFromCanonicalIds []string          `json:"merged_from_canonical_ids"`
	Metadata               map[string]string `json:"metadata"`
	MountAccessor          string            `json:"mount_accessor"`
	MountPath              string            `json:"mount_path"`
	MountType              string            `json:"mount_type"`
	Name                   string            `json:"name"`
}
