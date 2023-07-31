// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2localgateways "github.com/golingon/terraproviders/aws/5.10.0/dataec2localgateways"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2LocalGateways creates a new instance of [DataEc2LocalGateways].
func NewDataEc2LocalGateways(name string, args DataEc2LocalGatewaysArgs) *DataEc2LocalGateways {
	return &DataEc2LocalGateways{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2LocalGateways)(nil)

// DataEc2LocalGateways represents the Terraform data resource aws_ec2_local_gateways.
type DataEc2LocalGateways struct {
	Name string
	Args DataEc2LocalGatewaysArgs
}

// DataSource returns the Terraform object type for [DataEc2LocalGateways].
func (elg *DataEc2LocalGateways) DataSource() string {
	return "aws_ec2_local_gateways"
}

// LocalName returns the local name for [DataEc2LocalGateways].
func (elg *DataEc2LocalGateways) LocalName() string {
	return elg.Name
}

// Configuration returns the configuration (args) for [DataEc2LocalGateways].
func (elg *DataEc2LocalGateways) Configuration() interface{} {
	return elg.Args
}

// Attributes returns the attributes for [DataEc2LocalGateways].
func (elg *DataEc2LocalGateways) Attributes() dataEc2LocalGatewaysAttributes {
	return dataEc2LocalGatewaysAttributes{ref: terra.ReferenceDataResource(elg)}
}

// DataEc2LocalGatewaysArgs contains the configurations for aws_ec2_local_gateways.
type DataEc2LocalGatewaysArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Filter: min=0
	Filter []dataec2localgateways.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2localgateways.Timeouts `hcl:"timeouts,block"`
}
type dataEc2LocalGatewaysAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_local_gateways.
func (elg dataEc2LocalGatewaysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(elg.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ec2_local_gateways.
func (elg dataEc2LocalGatewaysAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](elg.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_ec2_local_gateways.
func (elg dataEc2LocalGatewaysAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](elg.ref.Append("tags"))
}

func (elg dataEc2LocalGatewaysAttributes) Filter() terra.SetValue[dataec2localgateways.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2localgateways.FilterAttributes](elg.ref.Append("filter"))
}

func (elg dataEc2LocalGatewaysAttributes) Timeouts() dataec2localgateways.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2localgateways.TimeoutsAttributes](elg.ref.Append("timeouts"))
}
