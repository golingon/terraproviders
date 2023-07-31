// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataecscluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ServiceConnectDefaults struct{}

type Setting struct{}

type ServiceConnectDefaultsAttributes struct {
	ref terra.Reference
}

func (scd ServiceConnectDefaultsAttributes) InternalRef() (terra.Reference, error) {
	return scd.ref, nil
}

func (scd ServiceConnectDefaultsAttributes) InternalWithRef(ref terra.Reference) ServiceConnectDefaultsAttributes {
	return ServiceConnectDefaultsAttributes{ref: ref}
}

func (scd ServiceConnectDefaultsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return scd.ref.InternalTokens()
}

func (scd ServiceConnectDefaultsAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(scd.ref.Append("namespace"))
}

type SettingAttributes struct {
	ref terra.Reference
}

func (s SettingAttributes) InternalRef() (terra.Reference, error) {
	return s.ref, nil
}

func (s SettingAttributes) InternalWithRef(ref terra.Reference) SettingAttributes {
	return SettingAttributes{ref: ref}
}

func (s SettingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return s.ref.InternalTokens()
}

func (s SettingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

func (s SettingAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
}

type ServiceConnectDefaultsState struct {
	Namespace string `json:"namespace"`
}

type SettingState struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
