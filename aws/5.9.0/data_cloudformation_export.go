// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudformationExport creates a new instance of [DataCloudformationExport].
func NewDataCloudformationExport(name string, args DataCloudformationExportArgs) *DataCloudformationExport {
	return &DataCloudformationExport{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudformationExport)(nil)

// DataCloudformationExport represents the Terraform data resource aws_cloudformation_export.
type DataCloudformationExport struct {
	Name string
	Args DataCloudformationExportArgs
}

// DataSource returns the Terraform object type for [DataCloudformationExport].
func (ce *DataCloudformationExport) DataSource() string {
	return "aws_cloudformation_export"
}

// LocalName returns the local name for [DataCloudformationExport].
func (ce *DataCloudformationExport) LocalName() string {
	return ce.Name
}

// Configuration returns the configuration (args) for [DataCloudformationExport].
func (ce *DataCloudformationExport) Configuration() interface{} {
	return ce.Args
}

// Attributes returns the attributes for [DataCloudformationExport].
func (ce *DataCloudformationExport) Attributes() dataCloudformationExportAttributes {
	return dataCloudformationExportAttributes{ref: terra.ReferenceDataResource(ce)}
}

// DataCloudformationExportArgs contains the configurations for aws_cloudformation_export.
type DataCloudformationExportArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataCloudformationExportAttributes struct {
	ref terra.Reference
}

// ExportingStackId returns a reference to field exporting_stack_id of aws_cloudformation_export.
func (ce dataCloudformationExportAttributes) ExportingStackId() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("exporting_stack_id"))
}

// Id returns a reference to field id of aws_cloudformation_export.
func (ce dataCloudformationExportAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudformation_export.
func (ce dataCloudformationExportAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("name"))
}

// Value returns a reference to field value of aws_cloudformation_export.
func (ce dataCloudformationExportAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("value"))
}
