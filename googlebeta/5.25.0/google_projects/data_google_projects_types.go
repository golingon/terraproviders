// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_projects

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataProjectsAttributes struct {
	ref terra.Reference
}

func (p DataProjectsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataProjectsAttributes) InternalWithRef(ref terra.Reference) DataProjectsAttributes {
	return DataProjectsAttributes{ref: ref}
}

func (p DataProjectsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataProjectsAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("create_time"))
}

func (p DataProjectsAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("labels"))
}

func (p DataProjectsAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("lifecycle_state"))
}

func (p DataProjectsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p DataProjectsAttributes) Number() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("number"))
}

func (p DataProjectsAttributes) Parent() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](p.ref.Append("parent"))
}

func (p DataProjectsAttributes) ProjectId() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("project_id"))
}

type DataProjectsState struct {
	CreateTime     string            `json:"create_time"`
	Labels         map[string]string `json:"labels"`
	LifecycleState string            `json:"lifecycle_state"`
	Name           string            `json:"name"`
	Number         string            `json:"number"`
	Parent         map[string]string `json:"parent"`
	ProjectId      string            `json:"project_id"`
}
