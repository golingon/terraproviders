// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_opensearchserverless_security_config

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataSamlOptions struct{}

type DataSamlOptionsAttributes struct {
	ref terra.Reference
}

func (so DataSamlOptionsAttributes) InternalRef() (terra.Reference, error) {
	return so.ref, nil
}

func (so DataSamlOptionsAttributes) InternalWithRef(ref terra.Reference) DataSamlOptionsAttributes {
	return DataSamlOptionsAttributes{ref: ref}
}

func (so DataSamlOptionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return so.ref.InternalTokens()
}

func (so DataSamlOptionsAttributes) GroupAttribute() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("group_attribute"))
}

func (so DataSamlOptionsAttributes) Metadata() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("metadata"))
}

func (so DataSamlOptionsAttributes) SessionTimeout() terra.NumberValue {
	return terra.ReferenceAsNumber(so.ref.Append("session_timeout"))
}

func (so DataSamlOptionsAttributes) UserAttribute() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("user_attribute"))
}

type DataSamlOptionsState struct {
	GroupAttribute string  `json:"group_attribute"`
	Metadata       string  `json:"metadata"`
	SessionTimeout float64 `json:"session_timeout"`
	UserAttribute  string  `json:"user_attribute"`
}
