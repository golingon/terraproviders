// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datadirectoryroletemplates

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RoleTemplates struct{}

type Timeouts struct {
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type RoleTemplatesAttributes struct {
	ref terra.Reference
}

func (rt RoleTemplatesAttributes) InternalRef() (terra.Reference, error) {
	return rt.ref, nil
}

func (rt RoleTemplatesAttributes) InternalWithRef(ref terra.Reference) RoleTemplatesAttributes {
	return RoleTemplatesAttributes{ref: ref}
}

func (rt RoleTemplatesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rt.ref.InternalTokens()
}

func (rt RoleTemplatesAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("description"))
}

func (rt RoleTemplatesAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("display_name"))
}

func (rt RoleTemplatesAttributes) ObjectId() terra.StringValue {
	return terra.ReferenceAsString(rt.ref.Append("object_id"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type RoleTemplatesState struct {
	Description string `json:"description"`
	DisplayName string `json:"display_name"`
	ObjectId    string `json:"object_id"`
}

type TimeoutsState struct {
	Read string `json:"read"`
}
