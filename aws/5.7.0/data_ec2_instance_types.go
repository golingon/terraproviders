// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	dataec2instancetypes "github.com/golingon/terraproviders/aws/5.7.0/dataec2instancetypes"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEc2InstanceTypes creates a new instance of [DataEc2InstanceTypes].
func NewDataEc2InstanceTypes(name string, args DataEc2InstanceTypesArgs) *DataEc2InstanceTypes {
	return &DataEc2InstanceTypes{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEc2InstanceTypes)(nil)

// DataEc2InstanceTypes represents the Terraform data resource aws_ec2_instance_types.
type DataEc2InstanceTypes struct {
	Name string
	Args DataEc2InstanceTypesArgs
}

// DataSource returns the Terraform object type for [DataEc2InstanceTypes].
func (eit *DataEc2InstanceTypes) DataSource() string {
	return "aws_ec2_instance_types"
}

// LocalName returns the local name for [DataEc2InstanceTypes].
func (eit *DataEc2InstanceTypes) LocalName() string {
	return eit.Name
}

// Configuration returns the configuration (args) for [DataEc2InstanceTypes].
func (eit *DataEc2InstanceTypes) Configuration() interface{} {
	return eit.Args
}

// Attributes returns the attributes for [DataEc2InstanceTypes].
func (eit *DataEc2InstanceTypes) Attributes() dataEc2InstanceTypesAttributes {
	return dataEc2InstanceTypesAttributes{ref: terra.ReferenceDataResource(eit)}
}

// DataEc2InstanceTypesArgs contains the configurations for aws_ec2_instance_types.
type DataEc2InstanceTypesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []dataec2instancetypes.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataec2instancetypes.Timeouts `hcl:"timeouts,block"`
}
type dataEc2InstanceTypesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_instance_types.
func (eit dataEc2InstanceTypesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(eit.ref.Append("id"))
}

// InstanceTypes returns a reference to field instance_types of aws_ec2_instance_types.
func (eit dataEc2InstanceTypesAttributes) InstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](eit.ref.Append("instance_types"))
}

func (eit dataEc2InstanceTypesAttributes) Filter() terra.SetValue[dataec2instancetypes.FilterAttributes] {
	return terra.ReferenceAsSet[dataec2instancetypes.FilterAttributes](eit.ref.Append("filter"))
}

func (eit dataEc2InstanceTypesAttributes) Timeouts() dataec2instancetypes.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataec2instancetypes.TimeoutsAttributes](eit.ref.Append("timeouts"))
}