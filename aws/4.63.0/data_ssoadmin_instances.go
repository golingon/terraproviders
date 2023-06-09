// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSsoadminInstances creates a new instance of [DataSsoadminInstances].
func NewDataSsoadminInstances(name string, args DataSsoadminInstancesArgs) *DataSsoadminInstances {
	return &DataSsoadminInstances{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSsoadminInstances)(nil)

// DataSsoadminInstances represents the Terraform data resource aws_ssoadmin_instances.
type DataSsoadminInstances struct {
	Name string
	Args DataSsoadminInstancesArgs
}

// DataSource returns the Terraform object type for [DataSsoadminInstances].
func (si *DataSsoadminInstances) DataSource() string {
	return "aws_ssoadmin_instances"
}

// LocalName returns the local name for [DataSsoadminInstances].
func (si *DataSsoadminInstances) LocalName() string {
	return si.Name
}

// Configuration returns the configuration (args) for [DataSsoadminInstances].
func (si *DataSsoadminInstances) Configuration() interface{} {
	return si.Args
}

// Attributes returns the attributes for [DataSsoadminInstances].
func (si *DataSsoadminInstances) Attributes() dataSsoadminInstancesAttributes {
	return dataSsoadminInstancesAttributes{ref: terra.ReferenceDataResource(si)}
}

// DataSsoadminInstancesArgs contains the configurations for aws_ssoadmin_instances.
type DataSsoadminInstancesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataSsoadminInstancesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_ssoadmin_instances.
func (si dataSsoadminInstancesAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](si.ref.Append("arns"))
}

// Id returns a reference to field id of aws_ssoadmin_instances.
func (si dataSsoadminInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(si.ref.Append("id"))
}

// IdentityStoreIds returns a reference to field identity_store_ids of aws_ssoadmin_instances.
func (si dataSsoadminInstancesAttributes) IdentityStoreIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](si.ref.Append("identity_store_ids"))
}
