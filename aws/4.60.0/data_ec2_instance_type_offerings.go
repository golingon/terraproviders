// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2instancetypeofferings "github.com/golingon/terraproviders/aws/4.60.0/dataec2instancetypeofferings"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2InstanceTypeOfferings creates a new instance of [DataEc2InstanceTypeOfferings].
func NewDataEc2InstanceTypeOfferings(name string, args DataEc2InstanceTypeOfferingsArgs) *DataEc2InstanceTypeOfferings {
	return &DataEc2InstanceTypeOfferings{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2InstanceTypeOfferings)(nil)

// DataEc2InstanceTypeOfferings represents the Terraform data resource aws_ec2_instance_type_offerings.
type DataEc2InstanceTypeOfferings struct {
	Name string
	Args DataEc2InstanceTypeOfferingsArgs
}

// DataSource returns the Terraform object type for [DataEc2InstanceTypeOfferings].
func (eito *DataEc2InstanceTypeOfferings) DataSource() string {
	return "aws_ec2_instance_type_offerings"
}

// LocalName returns the local name for [DataEc2InstanceTypeOfferings].
func (eito *DataEc2InstanceTypeOfferings) LocalName() string {
	return eito.Name
}

// Configuration returns the configuration (args) for [DataEc2InstanceTypeOfferings].
func (eito *DataEc2InstanceTypeOfferings) Configuration() interface{} {
	return eito.Args
}

// Attributes returns the attributes for [DataEc2InstanceTypeOfferings].
func (eito *DataEc2InstanceTypeOfferings) Attributes() dataEc2InstanceTypeOfferingsAttributes {
	return dataEc2InstanceTypeOfferingsAttributes{ref: terra.ReferenceDataResource(eito)}
}

// DataEc2InstanceTypeOfferingsArgs contains the configurations for aws_ec2_instance_type_offerings.
type DataEc2InstanceTypeOfferingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LocationType: string, optional
	LocationType terra.StringValue `hcl:"location_type,attr"`
	// Filter: min=0
	Filter []dataec2instancetypeofferings.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2instancetypeofferings.Timeouts `hcl:"timeouts,block"`
}
type dataEc2InstanceTypeOfferingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_instance_type_offerings.
func (eito dataEc2InstanceTypeOfferingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eito.ref.Append("id"))
}

// InstanceTypes returns a reference to field instance_types of aws_ec2_instance_type_offerings.
func (eito dataEc2InstanceTypeOfferingsAttributes) InstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eito.ref.Append("instance_types"))
}

// LocationType returns a reference to field location_type of aws_ec2_instance_type_offerings.
func (eito dataEc2InstanceTypeOfferingsAttributes) LocationType() terra.StringValue {
	return terra.ReferenceAsString(eito.ref.Append("location_type"))
}

// LocationTypes returns a reference to field location_types of aws_ec2_instance_type_offerings.
func (eito dataEc2InstanceTypeOfferingsAttributes) LocationTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eito.ref.Append("location_types"))
}

// Locations returns a reference to field locations of aws_ec2_instance_type_offerings.
func (eito dataEc2InstanceTypeOfferingsAttributes) Locations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eito.ref.Append("locations"))
}

func (eito dataEc2InstanceTypeOfferingsAttributes) Filter() terra.SetValue[dataec2instancetypeofferings.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2instancetypeofferings.FilterAttributes](eito.ref.Append("filter"))
}

func (eito dataEc2InstanceTypeOfferingsAttributes) Timeouts() dataec2instancetypeofferings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2instancetypeofferings.TimeoutsAttributes](eito.ref.Append("timeouts"))
}
