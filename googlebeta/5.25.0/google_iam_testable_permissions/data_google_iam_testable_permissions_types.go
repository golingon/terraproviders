// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_iam_testable_permissions

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataPermissionsAttributes struct {
	ref terra.Reference
}

func (p DataPermissionsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p DataPermissionsAttributes) InternalWithRef(ref terra.Reference) DataPermissionsAttributes {
	return DataPermissionsAttributes{ref: ref}
}

func (p DataPermissionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p DataPermissionsAttributes) ApiDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("api_disabled"))
}

func (p DataPermissionsAttributes) CustomSupportLevel() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("custom_support_level"))
}

func (p DataPermissionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p DataPermissionsAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("stage"))
}

func (p DataPermissionsAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("title"))
}

type DataPermissionsState struct {
	ApiDisabled        bool   `json:"api_disabled"`
	CustomSupportLevel string `json:"custom_support_level"`
	Name               string `json:"name"`
	Stage              string `json:"stage"`
	Title              string `json:"title"`
}
