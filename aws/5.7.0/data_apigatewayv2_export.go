// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataApigatewayv2Export creates a new instance of [DataApigatewayv2Export].
func NewDataApigatewayv2Export(name string, args DataApigatewayv2ExportArgs) *DataApigatewayv2Export {
	return &DataApigatewayv2Export{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApigatewayv2Export)(nil)

// DataApigatewayv2Export represents the Terraform data resource aws_apigatewayv2_export.
type DataApigatewayv2Export struct {
	Name string
	Args DataApigatewayv2ExportArgs
}

// DataSource returns the Terraform object type for [DataApigatewayv2Export].
func (ae *DataApigatewayv2Export) DataSource() string {
	return "aws_apigatewayv2_export"
}

// LocalName returns the local name for [DataApigatewayv2Export].
func (ae *DataApigatewayv2Export) LocalName() string {
	return ae.Name
}

// Configuration returns the configuration (args) for [DataApigatewayv2Export].
func (ae *DataApigatewayv2Export) Configuration() interface{} {
	return ae.Args
}

// Attributes returns the attributes for [DataApigatewayv2Export].
func (ae *DataApigatewayv2Export) Attributes() dataApigatewayv2ExportAttributes {
	return dataApigatewayv2ExportAttributes{ref: terra.ReferenceDataResource(ae)}
}

// DataApigatewayv2ExportArgs contains the configurations for aws_apigatewayv2_export.
type DataApigatewayv2ExportArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// ExportVersion: string, optional
	ExportVersion terra.StringValue `hcl:"export_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IncludeExtensions: bool, optional
	IncludeExtensions terra.BoolValue `hcl:"include_extensions,attr"`
	// OutputType: string, required
	OutputType terra.StringValue `hcl:"output_type,attr" validate:"required"`
	// Specification: string, required
	Specification terra.StringValue `hcl:"specification,attr" validate:"required"`
	// StageName: string, optional
	StageName terra.StringValue `hcl:"stage_name,attr"`
}
type dataApigatewayv2ExportAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("api_id"))
}

// Body returns a reference to field body of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("body"))
}

// ExportVersion returns a reference to field export_version of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) ExportVersion() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("export_version"))
}

// Id returns a reference to field id of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("id"))
}

// IncludeExtensions returns a reference to field include_extensions of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) IncludeExtensions() terra.BoolValue {
	return terra.ReferenceAsBool(ae.ref.Append("include_extensions"))
}

// OutputType returns a reference to field output_type of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) OutputType() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("output_type"))
}

// Specification returns a reference to field specification of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) Specification() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("specification"))
}

// StageName returns a reference to field stage_name of aws_apigatewayv2_export.
func (ae dataApigatewayv2ExportAttributes) StageName() terra.StringValue {
	return terra.ReferenceAsString(ae.ref.Append("stage_name"))
}
