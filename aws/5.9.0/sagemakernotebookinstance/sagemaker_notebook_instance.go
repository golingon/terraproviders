// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package sagemakernotebookinstance

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type InstanceMetadataServiceConfiguration struct {
	// MinimumInstanceMetadataServiceVersion: string, optional
	MinimumInstanceMetadataServiceVersion terra.StringValue `hcl:"minimum_instance_metadata_service_version,attr"`
}

type InstanceMetadataServiceConfigurationAttributes struct {
	ref terra.Reference
}

func (imsc InstanceMetadataServiceConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return imsc.ref, nil
}

func (imsc InstanceMetadataServiceConfigurationAttributes) InternalWithRef(ref terra.Reference) InstanceMetadataServiceConfigurationAttributes {
	return InstanceMetadataServiceConfigurationAttributes{ref: ref}
}

func (imsc InstanceMetadataServiceConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return imsc.ref.InternalTokens()
}

func (imsc InstanceMetadataServiceConfigurationAttributes) MinimumInstanceMetadataServiceVersion() terra.StringValue {
	return terra.ReferenceAsString(imsc.ref.Append("minimum_instance_metadata_service_version"))
}

type InstanceMetadataServiceConfigurationState struct {
	MinimumInstanceMetadataServiceVersion string `json:"minimum_instance_metadata_service_version"`
}