// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_project_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_project_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gpip *DataSource) DataSource() string {
	return "google_project_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gpip *DataSource) LocalName() string {
	return gpip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gpip *DataSource) Configuration() interface{} {
	return gpip.Args
}

// Attributes returns the attributes for [DataSource].
func (gpip *DataSource) Attributes() dataGoogleProjectIamPolicyAttributes {
	return dataGoogleProjectIamPolicyAttributes{ref: terra.ReferenceDataSource(gpip)}
}

// DataArgs contains the configurations for google_project_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
}

type dataGoogleProjectIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_project_iam_policy.
func (gpip dataGoogleProjectIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_iam_policy.
func (gpip dataGoogleProjectIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gpip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_project_iam_policy.
func (gpip dataGoogleProjectIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_project_iam_policy.
func (gpip dataGoogleProjectIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gpip.ref.Append("project"))
}
