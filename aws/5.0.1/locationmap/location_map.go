// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package locationmap

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Configuration struct {
	// Style: string, required
	Style terra.StringValue `hcl:"style,attr" validate:"required"`
}

type ConfigurationAttributes struct {
	ref terra.Reference
}

func (c ConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c ConfigurationAttributes) InternalWithRef(ref terra.Reference) ConfigurationAttributes {
	return ConfigurationAttributes{ref: ref}
}

func (c ConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c ConfigurationAttributes) Style() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("style"))
}

type ConfigurationState struct {
	Style string `json:"style"`
}
