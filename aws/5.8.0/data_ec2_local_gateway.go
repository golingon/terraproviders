// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgateway "github.com/golingon/terraproviders/aws/5.8.0/dataec2localgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGateway creates a new instance of [DataEc2LocalGateway].
func NewDataEc2LocalGateway(name string, args DataEc2LocalGatewayArgs) *DataEc2LocalGateway {
	return &DataEc2LocalGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGateway)(nil)

// DataEc2LocalGateway represents the Terraform data resource aws_ec2_local_gateway.
type DataEc2LocalGateway struct {
	Name string
	Args DataEc2LocalGatewayArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGateway].
func (elg *DataEc2LocalGateway) DataSource() string {
	return "aws_ec2_local_gateway"
}

// LocalName returns the local name for [DataEc2LocalGateway].
func (elg *DataEc2LocalGateway) LocalName() string {
	return elg.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGateway].
func (elg *DataEc2LocalGateway) Configuration() interface{} {
	return elg.Args
}

// Attributes returns the attributes for [DataEc2LocalGateway].
func (elg *DataEc2LocalGateway) Attributes() dataEc2LocalGatewayAttributes {
	return dataEc2LocalGatewayAttributes{ref: terra.ReferenceDataResource(elg)}
}

// DataEc2LocalGatewayArgs contains the configurations for aws_ec2_local_gateway.
type DataEc2LocalGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// State: string, optional
	State terra.StringValue `hcl:"state,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgateway.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgateway.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewayAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateway.
func (elg dataEc2LocalGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elg.ref.Append("id"))
}

// OutpostArn returns a reference to field outpost_arn of aws_ec2_local_gateway.
func (elg dataEc2LocalGatewayAttributes) OutpostArn() terra.StringValue {
	return terra.ReferenceAsString(elg.ref.Append("outpost_arn"))
}

// OwnerId returns a reference to field owner_id of aws_ec2_local_gateway.
func (elg dataEc2LocalGatewayAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(elg.ref.Append("owner_id"))
}

// State returns a reference to field state of aws_ec2_local_gateway.
func (elg dataEc2LocalGatewayAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(elg.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateway.
func (elg dataEc2LocalGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elg.ref.Append("tags"))
}

func (elg dataEc2LocalGatewayAttributes) Filter() terra.SetValue[dataec2localgateway.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgateway.FilterAttributes](elg.ref.Append("filter"))
}

func (elg dataEc2LocalGatewayAttributes) Timeouts() dataec2localgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgateway.TimeoutsAttributes](elg.ref.Append("timeouts"))
}
