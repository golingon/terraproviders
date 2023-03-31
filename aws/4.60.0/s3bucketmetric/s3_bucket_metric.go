// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package s3bucketmetric

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Filter struct {
	// Prefix: string, optional
	Prefix terra.StringValue `hcl:"prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
}

type FilterAttributes struct {
	ref terra.Reference
}

func (f FilterAttributes) InternalRef() terra.Reference {
	return f.ref
}

func (f FilterAttributes) InternalWithRef(ref terra.Reference) FilterAttributes {
	return FilterAttributes{ref: ref}
}

func (f FilterAttributes) InternalTokens() hclwrite.Tokens {
	return f.ref.InternalTokens()
}

func (f FilterAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("prefix"))
}

func (f FilterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("tags"))
}

type FilterState struct {
	Prefix string            `json:"prefix"`
	Tags   map[string]string `json:"tags"`
}