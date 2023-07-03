// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataProjectIamPolicy creates a new instance of [DataProjectIamPolicy].
func NewDataProjectIamPolicy(name string, args DataProjectIamPolicyArgs) *DataProjectIamPolicy {
	return &DataProjectIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataProjectIamPolicy)(nil)

// DataProjectIamPolicy represents the Terraform data resource google_project_iam_policy.
type DataProjectIamPolicy struct {
	Name string
	Args DataProjectIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataProjectIamPolicy].
func (pip *DataProjectIamPolicy) DataSource() string {
	return "google_project_iam_policy"
}

// LocalName returns the local name for [DataProjectIamPolicy].
func (pip *DataProjectIamPolicy) LocalName() string {
	return pip.Name
}

// Configuration returns the configuration (args) for [DataProjectIamPolicy].
func (pip *DataProjectIamPolicy) Configuration() interface{} {
	return pip.Args
}

// Attributes returns the attributes for [DataProjectIamPolicy].
func (pip *DataProjectIamPolicy) Attributes() dataProjectIamPolicyAttributes {
	return dataProjectIamPolicyAttributes{ref: terra.ReferenceDataResource(pip)}
}

// DataProjectIamPolicyArgs contains the configurations for google_project_iam_policy.
type DataProjectIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}
type dataProjectIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_project_iam_policy.
func (pip dataProjectIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_iam_policy.
func (pip dataProjectIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_project_iam_policy.
func (pip dataProjectIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_project_iam_policy.
func (pip dataProjectIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pip.ref.Append("project"))
}
