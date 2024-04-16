// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_networkfirewall_resource_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_networkfirewall_resource_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (anrp *DataSource) DataSource() string {
	return "aws_networkfirewall_resource_policy"
}

// LocalName returns the local name for [DataSource].
func (anrp *DataSource) LocalName() string {
	return anrp.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (anrp *DataSource) Configuration() interface{} {
	return anrp.Args
}

// Attributes returns the attributes for [DataSource].
func (anrp *DataSource) Attributes() dataAwsNetworkfirewallResourcePolicyAttributes {
	return dataAwsNetworkfirewallResourcePolicyAttributes{ref: terra.ReferenceDataSource(anrp)}
}

// DataArgs contains the configurations for aws_networkfirewall_resource_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceArn: string, required
	ResourceArn terra.StringValue `hcl:"resource_arn,attr" validate:"required"`
}

type dataAwsNetworkfirewallResourcePolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_networkfirewall_resource_policy.
func (anrp dataAwsNetworkfirewallResourcePolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anrp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_networkfirewall_resource_policy.
func (anrp dataAwsNetworkfirewallResourcePolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(anrp.ref.Append("policy"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkfirewall_resource_policy.
func (anrp dataAwsNetworkfirewallResourcePolicyAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(anrp.ref.Append("resource_arn"))
}
