// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataiamtestablepermissions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Permissions struct{}

type PermissionsAttributes struct {
	ref terra.Reference
}

func (p PermissionsAttributes) InternalRef() (terra.Reference, error) {
	return p.ref, nil
}

func (p PermissionsAttributes) InternalWithRef(ref terra.Reference) PermissionsAttributes {
	return PermissionsAttributes{ref: ref}
}

func (p PermissionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return p.ref.InternalTokens()
}

func (p PermissionsAttributes) ApiDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(p.ref.Append("api_disabled"))
}

func (p PermissionsAttributes) CustomSupportLevel() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("custom_support_level"))
}

func (p PermissionsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("name"))
}

func (p PermissionsAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("stage"))
}

func (p PermissionsAttributes) Title() terra.StringValue {
	return terra.ReferenceAsString(p.ref.Append("title"))
}

type PermissionsState struct {
	ApiDisabled        bool   `json:"api_disabled"`
	CustomSupportLevel string `json:"custom_support_level"`
	Name               string `json:"name"`
	Stage              string `json:"stage"`
	Title              string `json:"title"`
}
