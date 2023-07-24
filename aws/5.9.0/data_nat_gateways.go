// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanatgateways "github.com/golingon/terraproviders/aws/5.9.0/datanatgateways"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNatGateways creates a new instance of [DataNatGateways].
func NewDataNatGateways(name string, args DataNatGatewaysArgs) *DataNatGateways {
	return &DataNatGateways{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNatGateways)(nil)

// DataNatGateways represents the Terraform data resource aws_nat_gateways.
type DataNatGateways struct {
	Name string
	Args DataNatGatewaysArgs
}

// DataSource returns the Terraform object type for [DataNatGateways].
func (ng *DataNatGateways) DataSource() string {
	return "aws_nat_gateways"
}

// LocalName returns the local name for [DataNatGateways].
func (ng *DataNatGateways) LocalName() string {
	return ng.Name
}

// Configuration returns the configuration (args) for [DataNatGateways].
func (ng *DataNatGateways) Configuration() interface{} {
	return ng.Args
}

// Attributes returns the attributes for [DataNatGateways].
func (ng *DataNatGateways) Attributes() dataNatGatewaysAttributes {
	return dataNatGatewaysAttributes{ref: terra.ReferenceDataResource(ng)}
}

// DataNatGatewaysArgs contains the configurations for aws_nat_gateways.
type DataNatGatewaysArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VpcId: string, optional
	VpcId terra.StringValue `hcl:"vpc_id,attr"`
	// Filter: min=0
	Filter []datanatgateways.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanatgateways.Timeouts `hcl:"timeouts,block"`
}
type dataNatGatewaysAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_nat_gateways.
func (ng dataNatGatewaysAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_nat_gateways.
func (ng dataNatGatewaysAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ng.ref.Append("ids"))
}

// Tags returns a reference to field tags of aws_nat_gateways.
func (ng dataNatGatewaysAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ng.ref.Append("tags"))
}

// VpcId returns a reference to field vpc_id of aws_nat_gateways.
func (ng dataNatGatewaysAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(ng.ref.Append("vpc_id"))
}

func (ng dataNatGatewaysAttributes) Filter() terra.SetValue[datanatgateways.FilterAttributes] {
	return terra.ReferenceAsSet[datanatgateways.FilterAttributes](ng.ref.Append("filter"))
}

func (ng dataNatGatewaysAttributes) Timeouts() datanatgateways.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanatgateways.TimeoutsAttributes](ng.ref.Append("timeouts"))
}
