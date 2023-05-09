// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datavpcpeeringconnection "github.com/golingon/terraproviders/aws/4.66.1/datavpcpeeringconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataVpcPeeringConnection creates a new instance of [DataVpcPeeringConnection].
func NewDataVpcPeeringConnection(name string, args DataVpcPeeringConnectionArgs) *DataVpcPeeringConnection {
	return &DataVpcPeeringConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVpcPeeringConnection)(nil)

// DataVpcPeeringConnection represents the Terraform data resource aws_vpc_peering_connection.
type DataVpcPeeringConnection struct {
	Name string
	Args DataVpcPeeringConnectionArgs
}

// DataSource returns the Terraform object type for [DataVpcPeeringConnection].
func (vpc *DataVpcPeeringConnection) DataSource() string {
	return "aws_vpc_peering_connection"
}

// LocalName returns the local name for [DataVpcPeeringConnection].
func (vpc *DataVpcPeeringConnection) LocalName() string {
	return vpc.Name
}

// Configuration returns the configuration (args) for [DataVpcPeeringConnection].
func (vpc *DataVpcPeeringConnection) Configuration() interface{} {
	return vpc.Args
}

// Attributes returns the attributes for [DataVpcPeeringConnection].
func (vpc *DataVpcPeeringConnection) Attributes() dataVpcPeeringConnectionAttributes {
	return dataVpcPeeringConnectionAttributes{ref: terra.ReferenceDataResource(vpc)}
}

// DataVpcPeeringConnectionArgs contains the configurations for aws_vpc_peering_connection.
type DataVpcPeeringConnectionArgs struct {
	// CidrBlock: string, optional
	CidrBlock terra.StringValue `hcl:"cidr_block,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OwnerId: string, optional
	OwnerId terra.StringValue `hcl:"owner_id,attr"`
	// PeerCidrBlock: string, optional
	PeerCidrBlock terra.StringValue `hcl:"peer_cidr_block,attr"`
	// PeerOwnerId: string, optional
	PeerOwnerId terra.StringValue `hcl:"peer_owner_id,attr"`
	// PeerRegion: string, optional
	PeerRegion terra.StringValue `hcl:"peer_region,attr"`
	// PeerVpcId: string, optional
	PeerVpcId terra.StringValue `hcl:"peer_vpc_id,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Status: string, optional
	Status terra.StringValue `hcl:"status,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// CidrBlockSet: min=0
	CidrBlockSet []datavpcpeeringconnection.CidrBlockSet `hcl:"cidr_block_set,block" validate:"min=0"`
	// PeerCidrBlockSet: min=0
	PeerCidrBlockSet []datavpcpeeringconnection.PeerCidrBlockSet `hcl:"peer_cidr_block_set,block" validate:"min=0"`
	// Filter: min=0
	Filter []datavpcpeeringconnection.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datavpcpeeringconnection.Timeouts `hcl:"timeouts,block"`
}
type dataVpcPeeringConnectionAttributes struct {
	ref terra.Reference
}

// Accepter returns a reference to field accepter of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Accepter() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceAsMap[terra.BoolValue](vpc.ref.Append("accepter"))
}

// CidrBlock returns a reference to field cidr_block of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) CidrBlock() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("cidr_block"))
}

// Id returns a reference to field id of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("owner_id"))
}

// PeerCidrBlock returns a reference to field peer_cidr_block of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) PeerCidrBlock() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_cidr_block"))
}

// PeerOwnerId returns a reference to field peer_owner_id of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) PeerOwnerId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_owner_id"))
}

// PeerRegion returns a reference to field peer_region of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_region"))
}

// PeerVpcId returns a reference to field peer_vpc_id of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) PeerVpcId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_vpc_id"))
}

// Region returns a reference to field region of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("region"))
}

// Requester returns a reference to field requester of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Requester() terra.MapValue[terra.BoolValue] {
	return terra.ReferenceAsMap[terra.BoolValue](vpc.ref.Append("requester"))
}

// Status returns a reference to field status of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("status"))
}

// Tags returns a reference to field tags of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpc.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_peering_connection.
func (vpc dataVpcPeeringConnectionAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vpc_id"))
}

func (vpc dataVpcPeeringConnectionAttributes) CidrBlockSet() terra.ListValue[datavpcpeeringconnection.CidrBlockSetAttributes] {
	return terra.ReferenceAsList[datavpcpeeringconnection.CidrBlockSetAttributes](vpc.ref.Append("cidr_block_set"))
}

func (vpc dataVpcPeeringConnectionAttributes) PeerCidrBlockSet() terra.ListValue[datavpcpeeringconnection.PeerCidrBlockSetAttributes] {
	return terra.ReferenceAsList[datavpcpeeringconnection.PeerCidrBlockSetAttributes](vpc.ref.Append("peer_cidr_block_set"))
}

func (vpc dataVpcPeeringConnectionAttributes) Filter() terra.SetValue[datavpcpeeringconnection.FilterAttributes] {
	return terra.ReferenceAsSet[datavpcpeeringconnection.FilterAttributes](vpc.ref.Append("filter"))
}

func (vpc dataVpcPeeringConnectionAttributes) Timeouts() datavpcpeeringconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datavpcpeeringconnection.TimeoutsAttributes](vpc.ref.Append("timeouts"))
}
