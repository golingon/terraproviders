// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIamInstanceProfiles creates a new instance of [DataIamInstanceProfiles].
func NewDataIamInstanceProfiles(name string, args DataIamInstanceProfilesArgs) *DataIamInstanceProfiles {
	return &DataIamInstanceProfiles{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIamInstanceProfiles)(nil)

// DataIamInstanceProfiles represents the Terraform data resource aws_iam_instance_profiles.
type DataIamInstanceProfiles struct {
	Name string
	Args DataIamInstanceProfilesArgs
}

// DataSource returns the Terraform object type for [DataIamInstanceProfiles].
func (iip *DataIamInstanceProfiles) DataSource() string {
	return "aws_iam_instance_profiles"
}

// LocalName returns the local name for [DataIamInstanceProfiles].
func (iip *DataIamInstanceProfiles) LocalName() string {
	return iip.Name
}

// Configuration returns the configuration (args) for [DataIamInstanceProfiles].
func (iip *DataIamInstanceProfiles) Configuration() interface{} {
	return iip.Args
}

// Attributes returns the attributes for [DataIamInstanceProfiles].
func (iip *DataIamInstanceProfiles) Attributes() dataIamInstanceProfilesAttributes {
	return dataIamInstanceProfilesAttributes{ref: terra.ReferenceDataResource(iip)}
}

// DataIamInstanceProfilesArgs contains the configurations for aws_iam_instance_profiles.
type DataIamInstanceProfilesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleName: string, required
	RoleName terra.StringValue `hcl:"role_name,attr" validate:"required"`
}
type dataIamInstanceProfilesAttributes struct {
	ref terra.Reference
}

// Arns returns a reference to field arns of aws_iam_instance_profiles.
func (iip dataIamInstanceProfilesAttributes) Arns() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iip.ref.Append("arns"))
}

// Id returns a reference to field id of aws_iam_instance_profiles.
func (iip dataIamInstanceProfilesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("id"))
}

// Names returns a reference to field names of aws_iam_instance_profiles.
func (iip dataIamInstanceProfilesAttributes) Names() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iip.ref.Append("names"))
}

// Paths returns a reference to field paths of aws_iam_instance_profiles.
func (iip dataIamInstanceProfilesAttributes) Paths() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](iip.ref.Append("paths"))
}

// RoleName returns a reference to field role_name of aws_iam_instance_profiles.
func (iip dataIamInstanceProfilesAttributes) RoleName() terra.StringValue {
	return terra.ReferenceAsString(iip.ref.Append("role_name"))
}
