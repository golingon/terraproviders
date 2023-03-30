// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataDocdbEngineVersion(name string, args DataDocdbEngineVersionArgs) *DataDocdbEngineVersion {
	return &DataDocdbEngineVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDocdbEngineVersion)(nil)

type DataDocdbEngineVersion struct {
	Name string
	Args DataDocdbEngineVersionArgs
}

func (dev *DataDocdbEngineVersion) DataSource() string {
	return "aws_docdb_engine_version"
}

func (dev *DataDocdbEngineVersion) LocalName() string {
	return dev.Name
}

func (dev *DataDocdbEngineVersion) Configuration() interface{} {
	return dev.Args
}

func (dev *DataDocdbEngineVersion) Attributes() dataDocdbEngineVersionAttributes {
	return dataDocdbEngineVersionAttributes{ref: terra.ReferenceDataResource(dev)}
}

type DataDocdbEngineVersionArgs struct {
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
type dataDocdbEngineVersionAttributes struct {
	ref terra.Reference
}

func (dev dataDocdbEngineVersionAttributes) Engine() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("engine"))
}

func (dev dataDocdbEngineVersionAttributes) EngineDescription() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("engine_description"))
}

func (dev dataDocdbEngineVersionAttributes) ExportableLogTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dev.ref.Append("exportable_log_types"))
}

func (dev dataDocdbEngineVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("id"))
}

func (dev dataDocdbEngineVersionAttributes) ParameterGroupFamily() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("parameter_group_family"))
}

func (dev dataDocdbEngineVersionAttributes) PreferredVersions() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dev.ref.Append("preferred_versions"))
}

func (dev dataDocdbEngineVersionAttributes) SupportsLogExportsToCloudwatch() terra.BoolValue {
	return terra.ReferenceBool(dev.ref.Append("supports_log_exports_to_cloudwatch"))
}

func (dev dataDocdbEngineVersionAttributes) ValidUpgradeTargets() terra.SetValue[terra.StringValue] {
	return terra.ReferenceSet[terra.StringValue](dev.ref.Append("valid_upgrade_targets"))
}

func (dev dataDocdbEngineVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("version"))
}

func (dev dataDocdbEngineVersionAttributes) VersionDescription() terra.StringValue {
	return terra.ReferenceString(dev.ref.Append("version_description"))
}
