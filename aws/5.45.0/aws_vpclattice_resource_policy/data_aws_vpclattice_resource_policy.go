// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpclattice_resource_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_vpclattice_resource_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (avrp *DataSource) DataSource() string {
	return "aws_vpclattice_resource_policy"
}

// LocalName returns the local name for [DataSource].
func (avrp *DataSource) LocalName() string {
	return avrp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (avrp *DataSource) Configuration() interface{} {
	return avrp.Args
}

// Attributes returns the attributes for [DataSource].
func (avrp *DataSource) Attributes() dataAwsVpclatticeResourcePolicyAttributes {
	return dataAwsVpclatticeResourcePolicyAttributes{ref: terra.ReferenceDataSource(avrp)}
}

// DataArgs contains the configurations for aws_vpclattice_resource_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}

type dataAwsVpclatticeResourcePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_vpclattice_resource_policy.
func (avrp dataAwsVpclatticeResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avrp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_vpclattice_resource_policy.
func (avrp dataAwsVpclatticeResourcePolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(avrp.ref.Append("policy"))
}

// ResourceArn returns a reference to field resource_arn of aws_vpclattice_resource_policy.
func (avrp dataAwsVpclatticeResourcePolicyAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(avrp.ref.Append("resource_arn"))
}
