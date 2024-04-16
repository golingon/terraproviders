// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_appmesh_gateway_route

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_appmesh_gateway_route.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aagr *DataSource) DataSource() string {
	return "aws_appmesh_gateway_route"
}

// LocalName returns the local name for [DataSource].
func (aagr *DataSource) LocalName() string {
	return aagr.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aagr *DataSource) Configuration() interface{} {
	return aagr.Args
}

// Attributes returns the attributes for [DataSource].
func (aagr *DataSource) Attributes() dataAwsAppmeshGatewayRouteAttributes {
	return dataAwsAppmeshGatewayRouteAttributes{ref: terra.ReferenceDataSource(aagr)}
}

// DataArgs contains the configurations for aws_appmesh_gateway_route.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MeshName: string, required
	MeshName terra.StringValue `hcl:"mesh_name,attr" validate:"required"`
	// MeshOwner: string, optional
	MeshOwner terra.StringValue `hcl:"mesh_owner,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VirtualGatewayName: string, required
	VirtualGatewayName terra.StringValue `hcl:"virtual_gateway_name,attr" validate:"required"`
}

type dataAwsAppmeshGatewayRouteAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("arn"))
}

// CreatedDate returns a reference to field created_date of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("id"))
}

// LastUpdatedDate returns a reference to field last_updated_date of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) LastUpdatedDate() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("last_updated_date"))
}

// MeshName returns a reference to field mesh_name of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) MeshName() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("mesh_name"))
}

// MeshOwner returns a reference to field mesh_owner of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) MeshOwner() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("mesh_owner"))
}

// Name returns a reference to field name of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("name"))
}

// ResourceOwner returns a reference to field resource_owner of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) ResourceOwner() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("resource_owner"))
}

// Tags returns a reference to field tags of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aagr.ref.Append("tags"))
}

// VirtualGatewayName returns a reference to field virtual_gateway_name of aws_appmesh_gateway_route.
func (aagr dataAwsAppmeshGatewayRouteAttributes) VirtualGatewayName() terra.StringValue {
	return terra.ReferenceAsString(aagr.ref.Append("virtual_gateway_name"))
}

func (aagr dataAwsAppmeshGatewayRouteAttributes) Spec() terra.ListValue[DataSpecAttributes] {
	return terra.ReferenceAsList[DataSpecAttributes](aagr.ref.Append("spec"))
}
