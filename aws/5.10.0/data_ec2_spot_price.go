// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2spotprice "github.com/golingon/terraproviders/aws/5.10.0/dataec2spotprice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2SpotPrice creates a new instance of [DataEc2SpotPrice].
func NewDataEc2SpotPrice(name string, args DataEc2SpotPriceArgs) *DataEc2SpotPrice {
	return &DataEc2SpotPrice{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2SpotPrice)(nil)

// DataEc2SpotPrice represents the Terraform data resource aws_ec2_spot_price.
type DataEc2SpotPrice struct {
	Name string
	Args DataEc2SpotPriceArgs
}

// DataSource returns the Terraform object type for [DataEc2SpotPrice].
func (esp *DataEc2SpotPrice) DataSource() string {
	return "aws_ec2_spot_price"
}

// LocalName returns the local name for [DataEc2SpotPrice].
func (esp *DataEc2SpotPrice) LocalName() string {
	return esp.Name
}

// Configuration returns the configuration (args) for [DataEc2SpotPrice].
func (esp *DataEc2SpotPrice) Configuration() interface{} {
	return esp.Args
}

// Attributes returns the attributes for [DataEc2SpotPrice].
func (esp *DataEc2SpotPrice) Attributes() dataEc2SpotPriceAttributes {
	return dataEc2SpotPriceAttributes{ref: terra.ReferenceDataResource(esp)}
}

// DataEc2SpotPriceArgs contains the configurations for aws_ec2_spot_price.
type DataEc2SpotPriceArgs struct {
	// AvailabilityZone: string, optional
	AvailabilityZone terra.StringValue `hcl:"availability_zone,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// Filter: min=0
	Filter []dataec2spotprice.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2spotprice.Timeouts `hcl:"timeouts,block"`
}
type dataEc2SpotPriceAttributes struct {
	ref terra.Reference
}

// AvailabilityZone returns a reference to field availability_zone of aws_ec2_spot_price.
func (esp dataEc2SpotPriceAttributes) AvailabilityZone() terra.StringValue {
	return terra.ReferenceAsString(esp.ref.Append("availability_zone"))
}

// Id returns a reference to field id of aws_ec2_spot_price.
func (esp dataEc2SpotPriceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esp.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of aws_ec2_spot_price.
func (esp dataEc2SpotPriceAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(esp.ref.Append("instance_type"))
}

// SpotPrice returns a reference to field spot_price of aws_ec2_spot_price.
func (esp dataEc2SpotPriceAttributes) SpotPrice() terra.StringValue {
	return terra.ReferenceAsString(esp.ref.Append("spot_price"))
}

// SpotPriceTimestamp returns a reference to field spot_price_timestamp of aws_ec2_spot_price.
func (esp dataEc2SpotPriceAttributes) SpotPriceTimestamp() terra.StringValue {
	return terra.ReferenceAsString(esp.ref.Append("spot_price_timestamp"))
}

func (esp dataEc2SpotPriceAttributes) Filter() terra.SetValue[dataec2spotprice.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2spotprice.FilterAttributes](esp.ref.Append("filter"))
}

func (esp dataEc2SpotPriceAttributes) Timeouts() dataec2spotprice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2spotprice.TimeoutsAttributes](esp.ref.Append("timeouts"))
}