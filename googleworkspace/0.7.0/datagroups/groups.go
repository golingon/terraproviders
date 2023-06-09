// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datagroups

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Groups struct{}

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

func (g GroupsAttributes) AdminCreated() terra.BoolValue {
	return terra.ReferenceAsBool(g.ref.Append("admin_created"))
}

func (g GroupsAttributes) Aliases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("aliases"))
}

func (g GroupsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("description"))
}

func (g GroupsAttributes) DirectMembersCount() terra.NumberValue {
	return terra.ReferenceAsNumber(g.ref.Append("direct_members_count"))
}

func (g GroupsAttributes) Email() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("email"))
}

func (g GroupsAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("etag"))
}

func (g GroupsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("id"))
}

func (g GroupsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(g.ref.Append("name"))
}

func (g GroupsAttributes) NonEditableAliases() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](g.ref.Append("non_editable_aliases"))
}

type GroupsState struct {
	AdminCreated       bool     `json:"admin_created"`
	Aliases            []string `json:"aliases"`
	Description        string   `json:"description"`
	DirectMembersCount float64  `json:"direct_members_count"`
	Email              string   `json:"email"`
	Etag               string   `json:"etag"`
	Id                 string   `json:"id"`
	Name               string   `json:"name"`
	NonEditableAliases []string `json:"non_editable_aliases"`
}
