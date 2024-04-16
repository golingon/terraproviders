// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_guardduty_detector

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type DataFeaturesAttributes struct {
	ref terra.Reference
}

func (f DataFeaturesAttributes) InternalRef() (terra.Reference, error) {
	return f.ref, nil
}

func (f DataFeaturesAttributes) InternalWithRef(ref terra.Reference) DataFeaturesAttributes {
	return DataFeaturesAttributes{ref: ref}
}

func (f DataFeaturesAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return f.ref.InternalTokens()
}

func (f DataFeaturesAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

func (f DataFeaturesAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("status"))
}

func (f DataFeaturesAttributes) AdditionalConfiguration() terra.ListValue[DataFeaturesAdditionalConfigurationAttributes] {
	return terra.ReferenceAsList[DataFeaturesAdditionalConfigurationAttributes](f.ref.Append("additional_configuration"))
}

type DataFeaturesAdditionalConfigurationAttributes struct {
	ref terra.Reference
}

func (ac DataFeaturesAdditionalConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return ac.ref, nil
}

func (ac DataFeaturesAdditionalConfigurationAttributes) InternalWithRef(ref terra.Reference) DataFeaturesAdditionalConfigurationAttributes {
	return DataFeaturesAdditionalConfigurationAttributes{ref: ref}
}

func (ac DataFeaturesAdditionalConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ac.ref.InternalTokens()
}

func (ac DataFeaturesAdditionalConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("name"))
}

func (ac DataFeaturesAdditionalConfigurationAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(ac.ref.Append("status"))
}

type DataFeaturesState struct {
	Name                    string                                     `json:"name"`
	Status                  string                                     `json:"status"`
	AdditionalConfiguration []DataFeaturesAdditionalConfigurationState `json:"additional_configuration"`
}

type DataFeaturesAdditionalConfigurationState struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}
