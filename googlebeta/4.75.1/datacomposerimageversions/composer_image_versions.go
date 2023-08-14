// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacomposerimageversions

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type ImageVersions struct{}

type ImageVersionsAttributes struct {
	ref terra.Reference
}

func (iv ImageVersionsAttributes) InternalRef() (terra.Reference, error) {
	return iv.ref, nil
}

func (iv ImageVersionsAttributes) InternalWithRef(ref terra.Reference) ImageVersionsAttributes {
	return ImageVersionsAttributes{ref: ref}
}

func (iv ImageVersionsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return iv.ref.InternalTokens()
}

func (iv ImageVersionsAttributes) ImageVersionId() terra.StringValue {
	return terra.ReferenceAsString(iv.ref.Append("image_version_id"))
}

func (iv ImageVersionsAttributes) SupportedPythonVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](iv.ref.Append("supported_python_versions"))
}

type ImageVersionsState struct {
	ImageVersionId          string   `json:"image_version_id"`
	SupportedPythonVersions []string `json:"supported_python_versions"`
}