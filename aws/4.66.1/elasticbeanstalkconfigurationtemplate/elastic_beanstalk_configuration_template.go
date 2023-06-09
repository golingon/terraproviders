// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elasticbeanstalkconfigurationtemplate

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Setting struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Namespace: string, required
	Namespace terra.StringValue `hcl:"namespace,attr" validate:"required"`
	// Resource: string, optional
	Resource terra.StringValue `hcl:"resource,attr"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
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

func (s SettingAttributes) Namespace() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("namespace"))
}

func (s SettingAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("resource"))
}

func (s SettingAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("value"))
}

type SettingState struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Resource  string `json:"resource"`
	Value     string `json:"value"`
}
