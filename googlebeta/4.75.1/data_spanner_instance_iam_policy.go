// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataSpannerInstanceIamPolicy creates a new instance of [DataSpannerInstanceIamPolicy].
func NewDataSpannerInstanceIamPolicy(name string, args DataSpannerInstanceIamPolicyArgs) *DataSpannerInstanceIamPolicy {
	return &DataSpannerInstanceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSpannerInstanceIamPolicy)(nil)

// DataSpannerInstanceIamPolicy represents the Terraform data resource google_spanner_instance_iam_policy.
type DataSpannerInstanceIamPolicy struct {
	Name string
	Args DataSpannerInstanceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataSpannerInstanceIamPolicy].
func (siip *DataSpannerInstanceIamPolicy) DataSource() string {
	return "google_spanner_instance_iam_policy"
}

// LocalName returns the local name for [DataSpannerInstanceIamPolicy].
func (siip *DataSpannerInstanceIamPolicy) LocalName() string {
	return siip.Name
}

// Configuration returns the configuration (args) for [DataSpannerInstanceIamPolicy].
func (siip *DataSpannerInstanceIamPolicy) Configuration() interface{} {
	return siip.Args
}

// Attributes returns the attributes for [DataSpannerInstanceIamPolicy].
func (siip *DataSpannerInstanceIamPolicy) Attributes() dataSpannerInstanceIamPolicyAttributes {
	return dataSpannerInstanceIamPolicyAttributes{ref: terra.ReferenceDataResource(siip)}
}

// DataSpannerInstanceIamPolicyArgs contains the configurations for google_spanner_instance_iam_policy.
type DataSpannerInstanceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Instance: string, required
	Instance terra.StringValue `hcl:"instance,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataSpannerInstanceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_spanner_instance_iam_policy.
func (siip dataSpannerInstanceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(siip.ref.Append("etag"))
}

// Id returns a reference to field id of google_spanner_instance_iam_policy.
func (siip dataSpannerInstanceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(siip.ref.Append("id"))
}

// Instance returns a reference to field instance of google_spanner_instance_iam_policy.
func (siip dataSpannerInstanceIamPolicyAttributes) Instance() terra.StringValue {
	return terra.ReferenceAsString(siip.ref.Append("instance"))
}

// PolicyData returns a reference to field policy_data of google_spanner_instance_iam_policy.
func (siip dataSpannerInstanceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(siip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_spanner_instance_iam_policy.
func (siip dataSpannerInstanceIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(siip.ref.Append("project"))
}
