// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googleworkspace_privileges

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataItemsAttributes struct {
	ref terra.Reference
}

func (i DataItemsAttributes) InternalRef() (terra.Reference, error) {
	return i.ref, nil
}

func (i DataItemsAttributes) InternalWithRef(ref terra.Reference) DataItemsAttributes {
	return DataItemsAttributes{ref: ref}
}

func (i DataItemsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return i.ref.InternalTokens()
}

func (i DataItemsAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("etag"))
}

func (i DataItemsAttributes) IsOrgUnitScopable() terra.BoolValue {
	return terra.ReferenceAsBool(i.ref.Append("is_org_unit_scopable"))
}

func (i DataItemsAttributes) PrivilegeName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("privilege_name"))
}

func (i DataItemsAttributes) ServiceId() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("service_id"))
}

func (i DataItemsAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("service_name"))
}

type DataItemsState struct {
	Etag              string `json:"etag"`
	IsOrgUnitScopable bool   `json:"is_org_unit_scopable"`
	PrivilegeName     string `json:"privilege_name"`
	ServiceId         string `json:"service_id"`
	ServiceName       string `json:"service_name"`
}
