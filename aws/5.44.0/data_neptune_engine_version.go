// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataNeptuneEngineVersion creates a new instance of [DataNeptuneEngineVersion].
func NewDataNeptuneEngineVersion(name string, args DataNeptuneEngineVersionArgs) *DataNeptuneEngineVersion {
	return &DataNeptuneEngineVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNeptuneEngineVersion)(nil)

// DataNeptuneEngineVersion represents the Terraform data resource aws_neptune_engine_version.
type DataNeptuneEngineVersion struct {
	Name string
	Args DataNeptuneEngineVersionArgs
}

// DataSource returns the Terraform object type for [DataNeptuneEngineVersion].
func (nev *DataNeptuneEngineVersion) DataSource() string {
	return "aws_neptune_engine_version"
}

// LocalName returns the local name for [DataNeptuneEngineVersion].
func (nev *DataNeptuneEngineVersion) LocalName() string {
	return nev.Name
}

// Configuration returns the configuration (args) for [DataNeptuneEngineVersion].
func (nev *DataNeptuneEngineVersion) Configuration() interface{} {
	return nev.Args
}

// Attributes returns the attributes for [DataNeptuneEngineVersion].
func (nev *DataNeptuneEngineVersion) Attributes() dataNeptuneEngineVersionAttributes {
	return dataNeptuneEngineVersionAttributes{ref: terra.ReferenceDataResource(nev)}
}

// DataNeptuneEngineVersionArgs contains the configurations for aws_neptune_engine_version.
type DataNeptuneEngineVersionArgs struct {
	// Engine: string, optional
	Engine terra.StringValue `hcl:"engine,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ParameterGroupFamily: string, optional
	ParameterGroupFamily terra.StringValue `hcl:"parameter_group_family,attr"`
	// PreferredVersions: list of string, optional
	PreferredVersions terra.ListValue[terra.StringValue] `hcl:"preferred_versions,attr"`
	// Version: string, optional
	Version terra.StringValue `hcl:"version,attr"`
}
type dataNeptuneEngineVersionAttributes struct {
	ref terra.Reference
}

// Engine returns a reference to field engine of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) Engine() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("engine"))
}

// EngineDescription returns a reference to field engine_description of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) EngineDescription() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("engine_description"))
}

// ExportableLogTypes returns a reference to field exportable_log_types of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) ExportableLogTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nev.ref.Append("exportable_log_types"))
}

// Id returns a reference to field id of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("id"))
}

// ParameterGroupFamily returns a reference to field parameter_group_family of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) ParameterGroupFamily() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("parameter_group_family"))
}

// PreferredVersions returns a reference to field preferred_versions of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) PreferredVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](nev.ref.Append("preferred_versions"))
}

// SupportedTimezones returns a reference to field supported_timezones of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) SupportedTimezones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nev.ref.Append("supported_timezones"))
}

// SupportsLogExportsToCloudwatch returns a reference to field supports_log_exports_to_cloudwatch of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) SupportsLogExportsToCloudwatch() terra.BoolValue {
	return terra.ReferenceAsBool(nev.ref.Append("supports_log_exports_to_cloudwatch"))
}

// SupportsReadReplica returns a reference to field supports_read_replica of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) SupportsReadReplica() terra.BoolValue {
	return terra.ReferenceAsBool(nev.ref.Append("supports_read_replica"))
}

// ValidUpgradeTargets returns a reference to field valid_upgrade_targets of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) ValidUpgradeTargets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](nev.ref.Append("valid_upgrade_targets"))
}

// Version returns a reference to field version of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("version"))
}

// VersionDescription returns a reference to field version_description of aws_neptune_engine_version.
func (nev dataNeptuneEngineVersionAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceAsString(nev.ref.Append("version_description"))
}
