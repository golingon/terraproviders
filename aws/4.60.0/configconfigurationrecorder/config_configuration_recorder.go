// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package configconfigurationrecorder

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RecordingGroup struct {
	// AllSupported: bool, optional
	AllSupported terra.BoolValue `hcl:"all_supported,attr"`
	// IncludeGlobalResourceTypes: bool, optional
	IncludeGlobalResourceTypes terra.BoolValue `hcl:"include_global_resource_types,attr"`
	// ResourceTypes: set of string, optional
	ResourceTypes terra.SetValue[terra.StringValue] `hcl:"resource_types,attr"`
}

type RecordingGroupAttributes struct {
	ref terra.Reference
}

func (rg RecordingGroupAttributes) InternalRef() (terra.Reference, error) {
	return rg.ref, nil
}

func (rg RecordingGroupAttributes) InternalWithRef(ref terra.Reference) RecordingGroupAttributes {
	return RecordingGroupAttributes{ref: ref}
}

func (rg RecordingGroupAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return rg.ref.InternalTokens()
}

func (rg RecordingGroupAttributes) AllSupported() terra.BoolValue {
	return terra.ReferenceAsBool(rg.ref.Append("all_supported"))
}

func (rg RecordingGroupAttributes) IncludeGlobalResourceTypes() terra.BoolValue {
	return terra.ReferenceAsBool(rg.ref.Append("include_global_resource_types"))
}

func (rg RecordingGroupAttributes) ResourceTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rg.ref.Append("resource_types"))
}

type RecordingGroupState struct {
	AllSupported               bool     `json:"all_supported"`
	IncludeGlobalResourceTypes bool     `json:"include_global_resource_types"`
	ResourceTypes              []string `json:"resource_types"`
}
