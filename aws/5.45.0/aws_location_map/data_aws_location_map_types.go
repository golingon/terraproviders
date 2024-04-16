// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_location_map

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataConfigurationAttributes struct {
	ref terra.Reference
}

func (c DataConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c DataConfigurationAttributes) InternalWithRef(ref terra.Reference) DataConfigurationAttributes {
	return DataConfigurationAttributes{ref: ref}
}

func (c DataConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c DataConfigurationAttributes) Style() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("style"))
}

type DataConfigurationState struct {
	Style string `json:"style"`
}
