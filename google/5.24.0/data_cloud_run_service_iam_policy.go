// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataCloudRunServiceIamPolicy creates a new instance of [DataCloudRunServiceIamPolicy].
func NewDataCloudRunServiceIamPolicy(name string, args DataCloudRunServiceIamPolicyArgs) *DataCloudRunServiceIamPolicy {
	return &DataCloudRunServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudRunServiceIamPolicy)(nil)

// DataCloudRunServiceIamPolicy represents the Terraform data resource google_cloud_run_service_iam_policy.
type DataCloudRunServiceIamPolicy struct {
	Name string
	Args DataCloudRunServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataCloudRunServiceIamPolicy].
func (crsip *DataCloudRunServiceIamPolicy) DataSource() string {
	return "google_cloud_run_service_iam_policy"
}

// LocalName returns the local name for [DataCloudRunServiceIamPolicy].
func (crsip *DataCloudRunServiceIamPolicy) LocalName() string {
	return crsip.Name
}

// Configuration returns the configuration (args) for [DataCloudRunServiceIamPolicy].
func (crsip *DataCloudRunServiceIamPolicy) Configuration() interface{} {
	return crsip.Args
}

// Attributes returns the attributes for [DataCloudRunServiceIamPolicy].
func (crsip *DataCloudRunServiceIamPolicy) Attributes() dataCloudRunServiceIamPolicyAttributes {
	return dataCloudRunServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(crsip)}
}

// DataCloudRunServiceIamPolicyArgs contains the configurations for google_cloud_run_service_iam_policy.
type DataCloudRunServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Service: string, required
	Service terra.StringValue `hcl:"service,attr" validate:"required"`
}
type dataCloudRunServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("project"))
}

// Service returns a reference to field service of google_cloud_run_service_iam_policy.
func (crsip dataCloudRunServiceIamPolicyAttributes) Service() terra.StringValue {
	return terra.ReferenceAsString(crsip.ref.Append("service"))
}
