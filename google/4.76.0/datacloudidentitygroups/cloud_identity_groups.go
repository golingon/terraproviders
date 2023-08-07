// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudidentitygroups

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Groups struct {
	// GroupKey: min=0
	GroupKey []GroupKey `hcl:"group_key,block" validate:"min=0"`
}

type GroupKey struct{}

type GroupsAttributes struct {
	ref terra.Reference
}

func (g GroupsAttributes) InternalRef() (terra.Reference, error) {
	return g.ref, nil
}

func (g GroupsAttributes) InternalWithRef(ref terra.Reference) GroupsAttributes {
	return GroupsAttributes{ref: ref}
}

func (g GroupsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return g.ref.InternalTokens()
}

func (g GroupsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("create_time"))
}

func (g GroupsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("description"))
}

func (g GroupsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("display_name"))
}

func (g GroupsAttributes) InitialGroupConfig() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("initial_group_config"))
}

func (g GroupsAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](g.ref.Append("labels"))
}

func (g GroupsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("name"))
}

func (g GroupsAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("parent"))
}

func (g GroupsAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("update_time"))
}

func (g GroupsAttributes) GroupKey() terra.ListValue[GroupKeyAttributes] {
	return terra.ReferenceAsList[GroupKeyAttributes](g.ref.Append("group_key"))
}

type GroupKeyAttributes struct {
	ref terra.Reference
}

func (gk GroupKeyAttributes) InternalRef() (terra.Reference, error) {
	return gk.ref, nil
}

func (gk GroupKeyAttributes) InternalWithRef(ref terra.Reference) GroupKeyAttributes {
	return GroupKeyAttributes{ref: ref}
}

func (gk GroupKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gk.ref.InternalTokens()
}

func (gk GroupKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("id"))
}

func (gk GroupKeyAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(gk.ref.Append("namespace"))
}

type GroupsState struct {
	CreateTime         string            `json:"create_time"`
	Description        string            `json:"description"`
	DisplayName        string            `json:"display_name"`
	InitialGroupConfig string            `json:"initial_group_config"`
	Labels             map[string]string `json:"labels"`
	Name               string            `json:"name"`
	Parent             string            `json:"parent"`
	UpdateTime         string            `json:"update_time"`
	GroupKey           []GroupKeyState   `json:"group_key"`
}

type GroupKeyState struct {
	Id        string `json:"id"`
	Namespace string `json:"namespace"`
}
