// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataOpensearchserverlessVpcEndpoint creates a new instance of [DataOpensearchserverlessVpcEndpoint].
func NewDataOpensearchserverlessVpcEndpoint(name string, args DataOpensearchserverlessVpcEndpointArgs) *DataOpensearchserverlessVpcEndpoint {
	return &DataOpensearchserverlessVpcEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOpensearchserverlessVpcEndpoint)(nil)

// DataOpensearchserverlessVpcEndpoint represents the Terraform data resource aws_opensearchserverless_vpc_endpoint.
type DataOpensearchserverlessVpcEndpoint struct {
	Name string
	Args DataOpensearchserverlessVpcEndpointArgs
}

// DataSource returns the Terraform object type for [DataOpensearchserverlessVpcEndpoint].
func (ove *DataOpensearchserverlessVpcEndpoint) DataSource() string {
	return "aws_opensearchserverless_vpc_endpoint"
}

// LocalName returns the local name for [DataOpensearchserverlessVpcEndpoint].
func (ove *DataOpensearchserverlessVpcEndpoint) LocalName() string {
	return ove.Name
}

// Configuration returns the configuration (args) for [DataOpensearchserverlessVpcEndpoint].
func (ove *DataOpensearchserverlessVpcEndpoint) Configuration() interface{} {
	return ove.Args
}

// Attributes returns the attributes for [DataOpensearchserverlessVpcEndpoint].
func (ove *DataOpensearchserverlessVpcEndpoint) Attributes() dataOpensearchserverlessVpcEndpointAttributes {
	return dataOpensearchserverlessVpcEndpointAttributes{ref: terra.ReferenceDataResource(ove)}
}

// DataOpensearchserverlessVpcEndpointArgs contains the configurations for aws_opensearchserverless_vpc_endpoint.
type DataOpensearchserverlessVpcEndpointArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcEndpointId: string, required
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr" validate:"required"`
}
type dataOpensearchserverlessVpcEndpointAttributes struct {
	ref terra.Reference
}

// CreatedDate returns a reference to field created_date of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(ove.ref.Append("created_date"))
}

// Id returns a reference to field id of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ove.ref.Append("id"))
}

// Name returns a reference to field name of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ove.ref.Append("name"))
}

// SecurityGroupIds returns a reference to field security_group_ids of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) SecurityGroupIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ove.ref.Append("security_group_ids"))
}

// SubnetIds returns a reference to field subnet_ids of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) SubnetIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ove.ref.Append("subnet_ids"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(ove.ref.Append("vpc_endpoint_id"))
}

// VpcId returns a reference to field vpc_id of aws_opensearchserverless_vpc_endpoint.
func (ove dataOpensearchserverlessVpcEndpointAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ove.ref.Append("vpc_id"))
}