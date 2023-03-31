// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datadbinstances "github.com/golingon/terraproviders/aws/4.60.0/datadbinstances"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDbInstances creates a new instance of [DataDbInstances].
func NewDataDbInstances(name string, args DataDbInstancesArgs) *DataDbInstances {
	return &DataDbInstances{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDbInstances)(nil)

// DataDbInstances represents the Terraform data resource aws_db_instances.
type DataDbInstances struct {
	Name string
	Args DataDbInstancesArgs
}

// DataSource returns the Terraform object type for [DataDbInstances].
func (di *DataDbInstances) DataSource() string {
	return "aws_db_instances"
}

// LocalName returns the local name for [DataDbInstances].
func (di *DataDbInstances) LocalName() string {
	return di.Name
}

// Configuration returns the configuration (args) for [DataDbInstances].
func (di *DataDbInstances) Configuration() interface{} {
	return di.Args
}

// Attributes returns the attributes for [DataDbInstances].
func (di *DataDbInstances) Attributes() dataDbInstancesAttributes {
	return dataDbInstancesAttributes{ref: terra.ReferenceDataResource(di)}
}

// DataDbInstancesArgs contains the configurations for aws_db_instances.
type DataDbInstancesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []datadbinstances.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataDbInstancesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_db_instances.
func (di dataDbInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(di.ref.Append("id"))
}

// InstanceArns returns a reference to field instance_arns of aws_db_instances.
func (di dataDbInstancesAttributes) InstanceArns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("instance_arns"))
}

// InstanceIdentifiers returns a reference to field instance_identifiers of aws_db_instances.
func (di dataDbInstancesAttributes) InstanceIdentifiers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](di.ref.Append("instance_identifiers"))
}

func (di dataDbInstancesAttributes) Filter() terra.SetValue[datadbinstances.FilterAttributes] {
	return terra.ReferenceAsSet[datadbinstances.FilterAttributes](di.ref.Append("filter"))
}
