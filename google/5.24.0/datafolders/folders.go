// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package datafolders

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Folders struct{}

type FoldersAttributes struct {
	ref terra.Reference
}

func (f FoldersAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f FoldersAttributes) InternalWithRef(ref terra.Reference) FoldersAttributes {
	return FoldersAttributes{ref: ref}
}

func (f FoldersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f FoldersAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("create_time"))
}

func (f FoldersAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("delete_time"))
}

func (f FoldersAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("display_name"))
}

func (f FoldersAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("etag"))
}

func (f FoldersAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f FoldersAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("parent"))
}

func (f FoldersAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("state"))
}

func (f FoldersAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("update_time"))
}

type FoldersState struct {
	CreateTime  string `json:"create_time"`
	DeleteTime  string `json:"delete_time"`
	DisplayName string `json:"display_name"`
	Etag        string `json:"etag"`
	Name        string `json:"name"`
	Parent      string `json:"parent"`
	State       string `json:"state"`
	UpdateTime  string `json:"update_time"`
}
