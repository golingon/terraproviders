// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datassminstances "github.com/golingon/terraproviders/aws/5.7.0/datassminstances"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSsmInstances creates a new instance of [DataSsmInstances].
func NewDataSsmInstances(name string, args DataSsmInstancesArgs) *DataSsmInstances {
	return &DataSsmInstances{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsmInstances)(nil)

// DataSsmInstances represents the Terraform data resource aws_ssm_instances.
type DataSsmInstances struct {
	Name string
	Args DataSsmInstancesArgs
}

// DataSource returns the Terraform object type for [DataSsmInstances].
func (si *DataSsmInstances) DataSource() string {
	return "aws_ssm_instances"
}

// LocalName returns the local name for [DataSsmInstances].
func (si *DataSsmInstances) LocalName() string {
	return si.Name
}

// Configuration returns the configuration (args) for [DataSsmInstances].
func (si *DataSsmInstances) Configuration() interface{} {
	return si.Args
}

// Attributes returns the attributes for [DataSsmInstances].
func (si *DataSsmInstances) Attributes() dataSsmInstancesAttributes {
	return dataSsmInstancesAttributes{ref: terra.ReferenceDataResource(si)}
}

// DataSsmInstancesArgs contains the configurations for aws_ssm_instances.
type DataSsmInstancesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Filter: min=0
	Filter []datassminstances.Filter `hcl:"filter,block" validate:"min=0"`
}
type dataSsmInstancesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ssm_instances.
func (si dataSsmInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("id"))
}

// Ids returns a reference to field ids of aws_ssm_instances.
func (si dataSsmInstancesAttributes) Ids() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](si.ref.Append("ids"))
}

func (si dataSsmInstancesAttributes) Filter() terra.SetValue[datassminstances.FilterAttributes] {
	return terra.ReferenceAsSet[datassminstances.FilterAttributes](si.ref.Append("filter"))
}
