// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssm_maintenance_window_target

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type Targets struct {
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Values: list of string, required
	Values terra.ListValue[terra.StringValue] `hcl:"values,attr" validate:"required"`
}

type TargetsAttributes struct {
	ref terra.Reference
}

func (t TargetsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TargetsAttributes) InternalWithRef(ref terra.Reference) TargetsAttributes {
	return TargetsAttributes{ref: ref}
}

func (t TargetsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TargetsAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("key"))
}

func (t TargetsAttributes) Values() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](t.ref.Append("values"))
}

type TargetsState struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}
