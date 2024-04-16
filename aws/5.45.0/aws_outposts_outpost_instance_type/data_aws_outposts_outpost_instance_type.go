// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_outposts_outpost_instance_type

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_outposts_outpost_instance_type.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aooit *DataSource) DataSource() string {
	return "aws_outposts_outpost_instance_type"
}

// LocalName returns the local name for [DataSource].
func (aooit *DataSource) LocalName() string {
	return aooit.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aooit *DataSource) Configuration() interface{} {
	return aooit.Args
}

// Attributes returns the attributes for [DataSource].
func (aooit *DataSource) Attributes() dataAwsOutpostsOutpostInstanceTypeAttributes {
	return dataAwsOutpostsOutpostInstanceTypeAttributes{ref: terra.ReferenceDataSource(aooit)}
}

// DataArgs contains the configurations for aws_outposts_outpost_instance_type.
type DataArgs struct {
	// Arn: string, required
	Arn terra.StringValue `hcl:"arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceType: string, optional
	InstanceType terra.StringValue `hcl:"instance_type,attr"`
	// PreferredInstanceTypes: list of string, optional
	PreferredInstanceTypes terra.ListValue[terra.StringValue] `hcl:"preferred_instance_types,attr"`
}

type dataAwsOutpostsOutpostInstanceTypeAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_outposts_outpost_instance_type.
func (aooit dataAwsOutpostsOutpostInstanceTypeAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(aooit.ref.Append("arn"))
}

// Id returns a reference to field id of aws_outposts_outpost_instance_type.
func (aooit dataAwsOutpostsOutpostInstanceTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aooit.ref.Append("id"))
}

// InstanceType returns a reference to field instance_type of aws_outposts_outpost_instance_type.
func (aooit dataAwsOutpostsOutpostInstanceTypeAttributes) InstanceType() terra.StringValue {
	return terra.ReferenceAsString(aooit.ref.Append("instance_type"))
}

// PreferredInstanceTypes returns a reference to field preferred_instance_types of aws_outposts_outpost_instance_type.
func (aooit dataAwsOutpostsOutpostInstanceTypeAttributes) PreferredInstanceTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](aooit.ref.Append("preferred_instance_types"))
}
