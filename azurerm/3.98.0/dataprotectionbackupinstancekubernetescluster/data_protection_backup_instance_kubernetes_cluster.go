// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package dataprotectionbackupinstancekubernetescluster

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type BackupDatasourceParameters struct {
	// ClusterScopedResourcesEnabled: bool, optional
	ClusterScopedResourcesEnabled terra.BoolValue `hcl:"cluster_scoped_resources_enabled,attr"`
	// ExcludedNamespaces: list of string, optional
	ExcludedNamespaces terra.ListValue[terra.StringValue] `hcl:"excluded_namespaces,attr"`
	// ExcludedResourceTypes: list of string, optional
	ExcludedResourceTypes terra.ListValue[terra.StringValue] `hcl:"excluded_resource_types,attr"`
	// IncludedNamespaces: list of string, optional
	IncludedNamespaces terra.ListValue[terra.StringValue] `hcl:"included_namespaces,attr"`
	// IncludedResourceTypes: list of string, optional
	IncludedResourceTypes terra.ListValue[terra.StringValue] `hcl:"included_resource_types,attr"`
	// LabelSelectors: list of string, optional
	LabelSelectors terra.ListValue[terra.StringValue] `hcl:"label_selectors,attr"`
	// VolumeSnapshotEnabled: bool, optional
	VolumeSnapshotEnabled terra.BoolValue `hcl:"volume_snapshot_enabled,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type BackupDatasourceParametersAttributes struct {
	ref terra.Reference
}

func (bdp BackupDatasourceParametersAttributes) InternalRef() (terra.Reference, error) {
	return bdp.ref, nil
}

func (bdp BackupDatasourceParametersAttributes) InternalWithRef(ref terra.Reference) BackupDatasourceParametersAttributes {
	return BackupDatasourceParametersAttributes{ref: ref}
}

func (bdp BackupDatasourceParametersAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return bdp.ref.InternalTokens()
}

func (bdp BackupDatasourceParametersAttributes) ClusterScopedResourcesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bdp.ref.Append("cluster_scoped_resources_enabled"))
}

func (bdp BackupDatasourceParametersAttributes) ExcludedNamespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bdp.ref.Append("excluded_namespaces"))
}

func (bdp BackupDatasourceParametersAttributes) ExcludedResourceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bdp.ref.Append("excluded_resource_types"))
}

func (bdp BackupDatasourceParametersAttributes) IncludedNamespaces() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bdp.ref.Append("included_namespaces"))
}

func (bdp BackupDatasourceParametersAttributes) IncludedResourceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bdp.ref.Append("included_resource_types"))
}

func (bdp BackupDatasourceParametersAttributes) LabelSelectors() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](bdp.ref.Append("label_selectors"))
}

func (bdp BackupDatasourceParametersAttributes) VolumeSnapshotEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(bdp.ref.Append("volume_snapshot_enabled"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type BackupDatasourceParametersState struct {
	ClusterScopedResourcesEnabled bool     `json:"cluster_scoped_resources_enabled"`
	ExcludedNamespaces            []string `json:"excluded_namespaces"`
	ExcludedResourceTypes         []string `json:"excluded_resource_types"`
	IncludedNamespaces            []string `json:"included_namespaces"`
	IncludedResourceTypes         []string `json:"included_resource_types"`
	LabelSelectors                []string `json:"label_selectors"`
	VolumeSnapshotEnabled         bool     `json:"volume_snapshot_enabled"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
