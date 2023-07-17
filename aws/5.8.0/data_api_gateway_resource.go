// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataApiGatewayResource creates a new instance of [DataApiGatewayResource].
func NewDataApiGatewayResource(name string, args DataApiGatewayResourceArgs) *DataApiGatewayResource {
	return &DataApiGatewayResource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiGatewayResource)(nil)

// DataApiGatewayResource represents the Terraform data resource aws_api_gateway_resource.
type DataApiGatewayResource struct {
	Name string
	Args DataApiGatewayResourceArgs
}

// DataSource returns the Terraform object type for [DataApiGatewayResource].
func (agr *DataApiGatewayResource) DataSource() string {
	return "aws_api_gateway_resource"
}

// LocalName returns the local name for [DataApiGatewayResource].
func (agr *DataApiGatewayResource) LocalName() string {
	return agr.Name
}

// Configuration returns the configuration (args) for [DataApiGatewayResource].
func (agr *DataApiGatewayResource) Configuration() interface{} {
	return agr.Args
}

// Attributes returns the attributes for [DataApiGatewayResource].
func (agr *DataApiGatewayResource) Attributes() dataApiGatewayResourceAttributes {
	return dataApiGatewayResourceAttributes{ref: terra.ReferenceDataResource(agr)}
}

// DataApiGatewayResourceArgs contains the configurations for aws_api_gateway_resource.
type DataApiGatewayResourceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
}
type dataApiGatewayResourceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_api_gateway_resource.
func (agr dataApiGatewayResourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("id"))
}

// ParentId returns a reference to field parent_id of aws_api_gateway_resource.
func (agr dataApiGatewayResourceAttributes) ParentId() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("parent_id"))
}

// Path returns a reference to field path of aws_api_gateway_resource.
func (agr dataApiGatewayResourceAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("path"))
}

// PathPart returns a reference to field path_part of aws_api_gateway_resource.
func (agr dataApiGatewayResourceAttributes) PathPart() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("path_part"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_resource.
func (agr dataApiGatewayResourceAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agr.ref.Append("rest_api_id"))
}