// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataCloudRunV2ServiceIamPolicy creates a new instance of [DataCloudRunV2ServiceIamPolicy].
func NewDataCloudRunV2ServiceIamPolicy(name string, args DataCloudRunV2ServiceIamPolicyArgs) *DataCloudRunV2ServiceIamPolicy {
	return &DataCloudRunV2ServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCloudRunV2ServiceIamPolicy)(nil)

// DataCloudRunV2ServiceIamPolicy represents the Terraform data resource google_cloud_run_v2_service_iam_policy.
type DataCloudRunV2ServiceIamPolicy struct {
	Name string
	Args DataCloudRunV2ServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataCloudRunV2ServiceIamPolicy].
func (crvsip *DataCloudRunV2ServiceIamPolicy) DataSource() string {
	return "google_cloud_run_v2_service_iam_policy"
}

// LocalName returns the local name for [DataCloudRunV2ServiceIamPolicy].
func (crvsip *DataCloudRunV2ServiceIamPolicy) LocalName() string {
	return crvsip.Name
}

// Configuration returns the configuration (args) for [DataCloudRunV2ServiceIamPolicy].
func (crvsip *DataCloudRunV2ServiceIamPolicy) Configuration() interface{} {
	return crvsip.Args
}

// Attributes returns the attributes for [DataCloudRunV2ServiceIamPolicy].
func (crvsip *DataCloudRunV2ServiceIamPolicy) Attributes() dataCloudRunV2ServiceIamPolicyAttributes {
	return dataCloudRunV2ServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(crvsip)}
}

// DataCloudRunV2ServiceIamPolicyArgs contains the configurations for google_cloud_run_v2_service_iam_policy.
type DataCloudRunV2ServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataCloudRunV2ServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_cloud_run_v2_service_iam_policy.
func (crvsip dataCloudRunV2ServiceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvsip.ref.Append("project"))
}
