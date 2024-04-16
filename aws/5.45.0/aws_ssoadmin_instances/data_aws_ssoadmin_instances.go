// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ssoadmin_instances

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_ssoadmin_instances.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asi *DataSource) DataSource() string {
	return "aws_ssoadmin_instances"
}

// LocalName returns the local name for [DataSource].
func (asi *DataSource) LocalName() string {
	return asi.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asi *DataSource) Configuration() interface{} {
	return asi.Args
}

// Attributes returns the attributes for [DataSource].
func (asi *DataSource) Attributes() dataAwsSsoadminInstancesAttributes {
	return dataAwsSsoadminInstancesAttributes{ref: terra.ReferenceDataSource(asi)}
}

// DataArgs contains the configurations for aws_ssoadmin_instances.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataAwsSsoadminInstancesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_ssoadmin_instances.
func (asi dataAwsSsoadminInstancesAttributes) Arns() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asi.ref.Append("arns"))
}

// Id returns a reference to field id of aws_ssoadmin_instances.
func (asi dataAwsSsoadminInstancesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asi.ref.Append("id"))
}

// IdentityStoreIds returns a reference to field identity_store_ids of aws_ssoadmin_instances.
func (asi dataAwsSsoadminInstancesAttributes) IdentityStoreIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asi.ref.Append("identity_store_ids"))
}
