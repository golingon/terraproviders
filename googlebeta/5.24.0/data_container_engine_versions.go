// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataContainerEngineVersions creates a new instance of [DataContainerEngineVersions].
func NewDataContainerEngineVersions(name string, args DataContainerEngineVersionsArgs) *DataContainerEngineVersions {
	return &DataContainerEngineVersions{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerEngineVersions)(nil)

// DataContainerEngineVersions represents the Terraform data resource google_container_engine_versions.
type DataContainerEngineVersions struct {
	Name string
	Args DataContainerEngineVersionsArgs
}

// DataSource returns the Terraform object type for [DataContainerEngineVersions].
func (cev *DataContainerEngineVersions) DataSource() string {
	return "google_container_engine_versions"
}

// LocalName returns the local name for [DataContainerEngineVersions].
func (cev *DataContainerEngineVersions) LocalName() string {
	return cev.Name
}

// Configuration returns the configuration (args) for [DataContainerEngineVersions].
func (cev *DataContainerEngineVersions) Configuration() interface{} {
	return cev.Args
}

// Attributes returns the attributes for [DataContainerEngineVersions].
func (cev *DataContainerEngineVersions) Attributes() dataContainerEngineVersionsAttributes {
	return dataContainerEngineVersionsAttributes{ref: terra.ReferenceDataResource(cev)}
}

// DataContainerEngineVersionsArgs contains the configurations for google_container_engine_versions.
type DataContainerEngineVersionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// VersionPrefix: string, optional
	VersionPrefix terra.StringValue `hcl:"version_prefix,attr"`
}
type dataContainerEngineVersionsAttributes struct {
	ref terra.Reference
}

// DefaultClusterVersion returns a reference to field default_cluster_version of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) DefaultClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("default_cluster_version"))
}

// Id returns a reference to field id of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("id"))
}

// LatestMasterVersion returns a reference to field latest_master_version of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) LatestMasterVersion() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("latest_master_version"))
}

// LatestNodeVersion returns a reference to field latest_node_version of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) LatestNodeVersion() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("latest_node_version"))
}

// Location returns a reference to field location of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("location"))
}

// Project returns a reference to field project of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("project"))
}

// ReleaseChannelDefaultVersion returns a reference to field release_channel_default_version of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) ReleaseChannelDefaultVersion() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cev.ref.Append("release_channel_default_version"))
}

// ReleaseChannelLatestVersion returns a reference to field release_channel_latest_version of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) ReleaseChannelLatestVersion() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cev.ref.Append("release_channel_latest_version"))
}

// ValidMasterVersions returns a reference to field valid_master_versions of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) ValidMasterVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cev.ref.Append("valid_master_versions"))
}

// ValidNodeVersions returns a reference to field valid_node_versions of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) ValidNodeVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cev.ref.Append("valid_node_versions"))
}

// VersionPrefix returns a reference to field version_prefix of google_container_engine_versions.
func (cev dataContainerEngineVersionsAttributes) VersionPrefix() terra.StringValue {
	return terra.ReferenceAsString(cev.ref.Append("version_prefix"))
}
