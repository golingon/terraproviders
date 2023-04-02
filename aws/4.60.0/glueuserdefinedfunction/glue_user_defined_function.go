// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package glueuserdefinedfunction

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ResourceUris struct {
	// ResourceType: string, required
	ResourceType terra.StringValue `hcl:"resource_type,attr" validate:"required"`
	// Uri: string, required
	Uri terra.StringValue `hcl:"uri,attr" validate:"required"`
}

type ResourceUrisAttributes struct {
	ref terra.Reference
}

func (ru ResourceUrisAttributes) InternalRef() (terra.Reference, error) {
	return ru.ref, nil
}

func (ru ResourceUrisAttributes) InternalWithRef(ref terra.Reference) ResourceUrisAttributes {
	return ResourceUrisAttributes{ref: ref}
}

func (ru ResourceUrisAttributes) InternalTokens() hclwrite.Tokens {
	return ru.ref.InternalTokens()
}

func (ru ResourceUrisAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(ru.ref.Append("resource_type"))
}

func (ru ResourceUrisAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(ru.ref.Append("uri"))
}

type ResourceUrisState struct {
	ResourceType string `json:"resource_type"`
	Uri          string `json:"uri"`
}
