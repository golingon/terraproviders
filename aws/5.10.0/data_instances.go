// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datainstances "github.com/golingon/terraproviders/aws/5.10.0/datainstances"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataInstances creates a new instance of [DataInstances].
func NewDataInstances(name string, args DataInstancesArgs) *DataInstances {
	return &DataInstances{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataInstances)(nil)

// DataInstances represents the Terraform data resource aws_instances.
type DataInstances struct {
	Name string
	Args DataInstancesArgs
}

// DataSource returns the Terraform object type for [DataInstances].
func (i *DataInstances) DataSource() string {
	return "aws_instances"
}

// LocalName returns the local name for [DataInstances].
func (i *DataInstances) LocalName() string {
	return i.Name
}

// Configuration returns the configuration (args) for [DataInstances].
func (i *DataInstances) Configuration() interface{} {
	return i.Args
}

// Attributes returns the attributes for [DataInstances].
func (i *DataInstances) Attributes() dataInstancesAttributes {
	return dataInstancesAttributes{ref: terra.ReferenceDataResource(i)}
}

// DataInstancesArgs contains the configurations for aws_instances.
type DataInstancesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceStateNames: set of string, optional
	InstanceStateNames terra.SetValue[terra.StringValue] `hcl:"instance_state_names,attr"`
	// InstanceTags: map of string, optional
	InstanceTags terra.MapValue[terra.StringValue] `hcl:"instance_tags,attr"`
	// Filter: min=0
	Filter []datainstances.Filter `hcl:"filter,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datainstances.Timeouts `hcl:"timeouts,block"`
}
type dataInstancesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_instances.
func (i dataInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(i.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_instances.
func (i dataInstancesAttributes) Ids() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("ids"))
}

// InstanceStateNames returns a reference to field instance_state_names of aws_instances.
func (i dataInstancesAttributes) InstanceStateNames() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](i.ref.Append("instance_state_names"))
}

// InstanceTags returns a reference to field instance_tags of aws_instances.
func (i dataInstancesAttributes) InstanceTags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](i.ref.Append("instance_tags"))
}

// Ipv6Addresses returns a reference to field ipv6_addresses of aws_instances.
func (i dataInstancesAttributes) Ipv6Addresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("ipv6_addresses"))
}

// PrivateIps returns a reference to field private_ips of aws_instances.
func (i dataInstancesAttributes) PrivateIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("private_ips"))
}

// PublicIps returns a reference to field public_ips of aws_instances.
func (i dataInstancesAttributes) PublicIps() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](i.ref.Append("public_ips"))
}

func (i dataInstancesAttributes) Filter() terra.SetValue[datainstances.FilterAttributes] {
	return terra.ReferenceAsSet[datainstances.FilterAttributes](i.ref.Append("filter"))
}

func (i dataInstancesAttributes) Timeouts() datainstances.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datainstances.TimeoutsAttributes](i.ref.Append("timeouts"))
}
