// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataInspectorRulesPackages creates a new instance of [DataInspectorRulesPackages].
func NewDataInspectorRulesPackages(name string, args DataInspectorRulesPackagesArgs) *DataInspectorRulesPackages {
	return &DataInspectorRulesPackages{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataInspectorRulesPackages)(nil)

// DataInspectorRulesPackages represents the Terraform data resource aws_inspector_rules_packages.
type DataInspectorRulesPackages struct {
	Name string
	Args DataInspectorRulesPackagesArgs
}

// DataSource returns the Terraform object type for [DataInspectorRulesPackages].
func (irp *DataInspectorRulesPackages) DataSource() string {
	return "aws_inspector_rules_packages"
}

// LocalName returns the local name for [DataInspectorRulesPackages].
func (irp *DataInspectorRulesPackages) LocalName() string {
	return irp.Name
}

// Configuration returns the configuration (args) for [DataInspectorRulesPackages].
func (irp *DataInspectorRulesPackages) Configuration() interface{} {
	return irp.Args
}

// Attributes returns the attributes for [DataInspectorRulesPackages].
func (irp *DataInspectorRulesPackages) Attributes() dataInspectorRulesPackagesAttributes {
	return dataInspectorRulesPackagesAttributes{ref: terra.ReferenceDataResource(irp)}
}

// DataInspectorRulesPackagesArgs contains the configurations for aws_inspector_rules_packages.
type DataInspectorRulesPackagesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataInspectorRulesPackagesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_inspector_rules_packages.
func (irp dataInspectorRulesPackagesAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](irp.ref.Append("arns"))
}

// Id returns a reference to field id of aws_inspector_rules_packages.
func (irp dataInspectorRulesPackagesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(irp.ref.Append("id"))
}