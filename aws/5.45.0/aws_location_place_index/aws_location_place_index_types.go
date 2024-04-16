// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_location_place_index

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataSourceConfiguration struct {
	// IntendedUse: string, optional
	IntendedUse terra.StringValue `hcl:"intended_use,attr"`
}

type DataSourceConfigurationAttributes struct {
	ref terra.Reference
}

func (dsc DataSourceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return dsc.ref, nil
}

func (dsc DataSourceConfigurationAttributes) InternalWithRef(ref terra.Reference) DataSourceConfigurationAttributes {
	return DataSourceConfigurationAttributes{ref: ref}
}

func (dsc DataSourceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dsc.ref.InternalTokens()
}

func (dsc DataSourceConfigurationAttributes) IntendedUse() terra.StringValue {
	return terra.ReferenceAsString(dsc.ref.Append("intended_use"))
}

type DataSourceConfigurationState struct {
	IntendedUse string `json:"intended_use"`
}
