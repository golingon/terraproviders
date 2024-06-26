// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_apigee_environment_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_apigee_environment_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gaeip *DataSource) DataSource() string {
	return "google_apigee_environment_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gaeip *DataSource) LocalName() string {
	return gaeip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gaeip *DataSource) Configuration() interface{} {
	return gaeip.Args
}

// Attributes returns the attributes for [DataSource].
func (gaeip *DataSource) Attributes() dataGoogleApigeeEnvironmentIamPolicyAttributes {
	return dataGoogleApigeeEnvironmentIamPolicyAttributes{ref: terra.ReferenceDataSource(gaeip)}
}

// DataArgs contains the configurations for google_apigee_environment_iam_policy.
type DataArgs struct {
	// EnvId: string, required
	EnvId terra.StringValue `hcl:"env_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
}

type dataGoogleApigeeEnvironmentIamPolicyAttributes struct {
	ref terra.Reference
}

// EnvId returns a reference to field env_id of google_apigee_environment_iam_policy.
func (gaeip dataGoogleApigeeEnvironmentIamPolicyAttributes) EnvId() terra.StringValue {
	return terra.ReferenceAsString(gaeip.ref.Append("env_id"))
}

// Etag returns a reference to field etag of google_apigee_environment_iam_policy.
func (gaeip dataGoogleApigeeEnvironmentIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gaeip.ref.Append("etag"))
}

// Id returns a reference to field id of google_apigee_environment_iam_policy.
func (gaeip dataGoogleApigeeEnvironmentIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gaeip.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_apigee_environment_iam_policy.
func (gaeip dataGoogleApigeeEnvironmentIamPolicyAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(gaeip.ref.Append("org_id"))
}

// PolicyData returns a reference to field policy_data of google_apigee_environment_iam_policy.
func (gaeip dataGoogleApigeeEnvironmentIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gaeip.ref.Append("policy_data"))
}
