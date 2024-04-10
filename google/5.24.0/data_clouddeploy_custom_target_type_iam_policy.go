// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataClouddeployCustomTargetTypeIamPolicy creates a new instance of [DataClouddeployCustomTargetTypeIamPolicy].
func NewDataClouddeployCustomTargetTypeIamPolicy(name string, args DataClouddeployCustomTargetTypeIamPolicyArgs) *DataClouddeployCustomTargetTypeIamPolicy {
	return &DataClouddeployCustomTargetTypeIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataClouddeployCustomTargetTypeIamPolicy)(nil)

// DataClouddeployCustomTargetTypeIamPolicy represents the Terraform data resource google_clouddeploy_custom_target_type_iam_policy.
type DataClouddeployCustomTargetTypeIamPolicy struct {
	Name string
	Args DataClouddeployCustomTargetTypeIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataClouddeployCustomTargetTypeIamPolicy].
func (ccttip *DataClouddeployCustomTargetTypeIamPolicy) DataSource() string {
	return "google_clouddeploy_custom_target_type_iam_policy"
}

// LocalName returns the local name for [DataClouddeployCustomTargetTypeIamPolicy].
func (ccttip *DataClouddeployCustomTargetTypeIamPolicy) LocalName() string {
	return ccttip.Name
}

// Configuration returns the configuration (args) for [DataClouddeployCustomTargetTypeIamPolicy].
func (ccttip *DataClouddeployCustomTargetTypeIamPolicy) Configuration() interface{} {
	return ccttip.Args
}

// Attributes returns the attributes for [DataClouddeployCustomTargetTypeIamPolicy].
func (ccttip *DataClouddeployCustomTargetTypeIamPolicy) Attributes() dataClouddeployCustomTargetTypeIamPolicyAttributes {
	return dataClouddeployCustomTargetTypeIamPolicyAttributes{ref: terra.ReferenceDataResource(ccttip)}
}

// DataClouddeployCustomTargetTypeIamPolicyArgs contains the configurations for google_clouddeploy_custom_target_type_iam_policy.
type DataClouddeployCustomTargetTypeIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataClouddeployCustomTargetTypeIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("etag"))
}

// Id returns a reference to field id of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("id"))
}

// Location returns a reference to field location of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("location"))
}

// Name returns a reference to field name of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_clouddeploy_custom_target_type_iam_policy.
func (ccttip dataClouddeployCustomTargetTypeIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ccttip.ref.Append("project"))
}
