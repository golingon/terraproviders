// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_directory_service_shared_directory

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Target struct {
	// Id: string, required
	Id terra.StringValue `hcl:"id,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type Timeouts struct {
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
}

type TargetAttributes struct {
	ref terra.Reference
}

func (t TargetAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TargetAttributes) InternalWithRef(ref terra.Reference) TargetAttributes {
	return TargetAttributes{ref: ref}
}

func (t TargetAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TargetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("id"))
}

func (t TargetAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("type"))
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

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

type TargetState struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type TimeoutsState struct {
	Delete string `json:"delete"`
}
