// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datamonitoringnotificationchannel

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type SensitiveLabels struct{}

type SensitiveLabelsAttributes struct {
	ref terra.Reference
}

func (sl SensitiveLabelsAttributes) InternalRef() (terra.Reference, error) {
	return sl.ref, nil
}

func (sl SensitiveLabelsAttributes) InternalWithRef(ref terra.Reference) SensitiveLabelsAttributes {
	return SensitiveLabelsAttributes{ref: ref}
}

func (sl SensitiveLabelsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sl.ref.InternalTokens()
}

func (sl SensitiveLabelsAttributes) AuthToken() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("auth_token"))
}

func (sl SensitiveLabelsAttributes) Password() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("password"))
}

func (sl SensitiveLabelsAttributes) ServiceKey() terra.StringValue {
	return terra.ReferenceAsString(sl.ref.Append("service_key"))
}

type SensitiveLabelsState struct {
	AuthToken  string `json:"auth_token"`
	Password   string `json:"password"`
	ServiceKey string `json:"service_key"`
}