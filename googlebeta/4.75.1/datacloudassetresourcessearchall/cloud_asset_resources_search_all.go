// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package datacloudassetresourcessearchall

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Results struct{}

type ResultsAttributes struct {
	ref terra.Reference
}

func (r ResultsAttributes) InternalRef() (terra.Reference, error) {
	return r.ref, nil
}

func (r ResultsAttributes) InternalWithRef(ref terra.Reference) ResultsAttributes {
	return ResultsAttributes{ref: ref}
}

func (r ResultsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return r.ref.InternalTokens()
}

func (r ResultsAttributes) AdditionalAttributes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("additional_attributes"))
}

func (r ResultsAttributes) AssetType() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("asset_type"))
}

func (r ResultsAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("description"))
}

func (r ResultsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("display_name"))
}

func (r ResultsAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](r.ref.Append("labels"))
}

func (r ResultsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("location"))
}

func (r ResultsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("name"))
}

func (r ResultsAttributes) NetworkTags() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](r.ref.Append("network_tags"))
}

func (r ResultsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(r.ref.Append("project"))
}

type ResultsState struct {
	AdditionalAttributes []string          `json:"additional_attributes"`
	AssetType            string            `json:"asset_type"`
	Description          string            `json:"description"`
	DisplayName          string            `json:"display_name"`
	Labels               map[string]string `json:"labels"`
	Location             string            `json:"location"`
	Name                 string            `json:"name"`
	NetworkTags          []string          `json:"network_tags"`
	Project              string            `json:"project"`
}
