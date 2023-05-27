// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcpeeringconnections "github.com/golingon/terraproviders/aws/5.0.1/datavpcpeeringconnections"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcPeeringConnections creates a new instance of [DataVpcPeeringConnections].
func NewDataVpcPeeringConnections(name string, args DataVpcPeeringConnectionsArgs) *DataVpcPeeringConnections {
	return &DataVpcPeeringConnections{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcPeeringConnections)(nil)

// DataVpcPeeringConnections represents the Terraform data resource aws_vpc_peering_connections.
type DataVpcPeeringConnections struct {
	Name string
	Args DataVpcPeeringConnectionsArgs
}

// DataSource returns the Terraform object type for [DataVpcPeeringConnections].
func (vpc *DataVpcPeeringConnections) DataSource() string {
	return "aws_vpc_peering_connections"
}

// LocalName returns the local name for [DataVpcPeeringConnections].
func (vpc *DataVpcPeeringConnections) LocalName() string {
	return vpc.Name
}

// Configuration returns the configuration (args) for [DataVpcPeeringConnections].
func (vpc *DataVpcPeeringConnections) Configuration() interface{} {
	return vpc.Args
}

// Attributes returns the attributes for [DataVpcPeeringConnections].
func (vpc *DataVpcPeeringConnections) Attributes() dataVpcPeeringConnectionsAttributes {
	return dataVpcPeeringConnectionsAttributes{ref: terra.ReferenceDataResource(vpc)}
}

// DataVpcPeeringConnectionsArgs contains the configurations for aws_vpc_peering_connections.
type DataVpcPeeringConnectionsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []datavpcpeeringconnections.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcpeeringconnections.Timeouts `hcl:"timeouts,block"`
}
type dataVpcPeeringConnectionsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpc_peering_connections.
func (vpc dataVpcPeeringConnectionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_vpc_peering_connections.
func (vpc dataVpcPeeringConnectionsAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](vpc.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_vpc_peering_connections.
func (vpc dataVpcPeeringConnectionsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpc.ref.Append("tags"))
}

func (vpc dataVpcPeeringConnectionsAttributes) Filter() terra.SetValue[datavpcpeeringconnections.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcpeeringconnections.FilterAttributes](vpc.ref.Append("filter"))
}

func (vpc dataVpcPeeringConnectionsAttributes) Timeouts() datavpcpeeringconnections.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcpeeringconnections.TimeoutsAttributes](vpc.ref.Append("timeouts"))
}
