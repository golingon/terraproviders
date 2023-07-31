// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataDataprocAutoscalingPolicyIamPolicy creates a new instance of [DataDataprocAutoscalingPolicyIamPolicy].
func NewDataDataprocAutoscalingPolicyIamPolicy(name string, args DataDataprocAutoscalingPolicyIamPolicyArgs) *DataDataprocAutoscalingPolicyIamPolicy {
	return &DataDataprocAutoscalingPolicyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataprocAutoscalingPolicyIamPolicy)(nil)

// DataDataprocAutoscalingPolicyIamPolicy represents the Terraform data resource google_dataproc_autoscaling_policy_iam_policy.
type DataDataprocAutoscalingPolicyIamPolicy struct {
	Name string
	Args DataDataprocAutoscalingPolicyIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataDataprocAutoscalingPolicyIamPolicy].
func (dapip *DataDataprocAutoscalingPolicyIamPolicy) DataSource() string {
	return "google_dataproc_autoscaling_policy_iam_policy"
}

// LocalName returns the local name for [DataDataprocAutoscalingPolicyIamPolicy].
func (dapip *DataDataprocAutoscalingPolicyIamPolicy) LocalName() string {
	return dapip.Name
}

// Configuration returns the configuration (args) for [DataDataprocAutoscalingPolicyIamPolicy].
func (dapip *DataDataprocAutoscalingPolicyIamPolicy) Configuration() interface{} {
	return dapip.Args
}

// Attributes returns the attributes for [DataDataprocAutoscalingPolicyIamPolicy].
func (dapip *DataDataprocAutoscalingPolicyIamPolicy) Attributes() dataDataprocAutoscalingPolicyIamPolicyAttributes {
	return dataDataprocAutoscalingPolicyIamPolicyAttributes{ref: terra.ReferenceDataResource(dapip)}
}

// DataDataprocAutoscalingPolicyIamPolicyArgs contains the configurations for google_dataproc_autoscaling_policy_iam_policy.
type DataDataprocAutoscalingPolicyIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataDataprocAutoscalingPolicyIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("etag"))
}

// Id returns a reference to field id of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("id"))
}

// Location returns a reference to field location of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("policy_data"))
}

// PolicyId returns a reference to field policy_id of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("policy_id"))
}

// Project returns a reference to field project of google_dataproc_autoscaling_policy_iam_policy.
func (dapip dataDataprocAutoscalingPolicyIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dapip.ref.Append("project"))
}
