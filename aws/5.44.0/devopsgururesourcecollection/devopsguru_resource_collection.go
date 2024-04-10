// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package devopsgururesourcecollection

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Cloudformation struct {
	// StackNames: list of string, required
	StackNames terra.ListValue[terra.StringValue] `hcl:"stack_names,attr" validate:"required"`
}

type Tags struct {
	// AppBoundaryKey: string, required
	AppBoundaryKey terra.StringValue `hcl:"app_boundary_key,attr" validate:"required"`
	// TagValues: list of string, required
	TagValues terra.ListValue[terra.StringValue] `hcl:"tag_values,attr" validate:"required"`
}

type CloudformationAttributes struct {
	ref terra.Reference
}

func (c CloudformationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CloudformationAttributes) InternalWithRef(ref terra.Reference) CloudformationAttributes {
	return CloudformationAttributes{ref: ref}
}

func (c CloudformationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CloudformationAttributes) StackNames() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](c.ref.Append("stack_names"))
}

type TagsAttributes struct {
	ref terra.Reference
}

func (t TagsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TagsAttributes) InternalWithRef(ref terra.Reference) TagsAttributes {
	return TagsAttributes{ref: ref}
}

func (t TagsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TagsAttributes) AppBoundaryKey() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("app_boundary_key"))
}

func (t TagsAttributes) TagValues() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("tag_values"))
}

type CloudformationState struct {
	StackNames []string `json:"stack_names"`
}

type TagsState struct {
	AppBoundaryKey string   `json:"app_boundary_key"`
	TagValues      []string `json:"tag_values"`
}
