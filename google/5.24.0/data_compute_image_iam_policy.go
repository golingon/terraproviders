// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataComputeImageIamPolicy creates a new instance of [DataComputeImageIamPolicy].
func NewDataComputeImageIamPolicy(name string, args DataComputeImageIamPolicyArgs) *DataComputeImageIamPolicy {
	return &DataComputeImageIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeImageIamPolicy)(nil)

// DataComputeImageIamPolicy represents the Terraform data resource google_compute_image_iam_policy.
type DataComputeImageIamPolicy struct {
	Name string
	Args DataComputeImageIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataComputeImageIamPolicy].
func (ciip *DataComputeImageIamPolicy) DataSource() string {
	return "google_compute_image_iam_policy"
}

// LocalName returns the local name for [DataComputeImageIamPolicy].
func (ciip *DataComputeImageIamPolicy) LocalName() string {
	return ciip.Name
}

// Configuration returns the configuration (args) for [DataComputeImageIamPolicy].
func (ciip *DataComputeImageIamPolicy) Configuration() interface{} {
	return ciip.Args
}

// Attributes returns the attributes for [DataComputeImageIamPolicy].
func (ciip *DataComputeImageIamPolicy) Attributes() dataComputeImageIamPolicyAttributes {
	return dataComputeImageIamPolicyAttributes{ref: terra.ReferenceDataResource(ciip)}
}

// DataComputeImageIamPolicyArgs contains the configurations for google_compute_image_iam_policy.
type DataComputeImageIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataComputeImageIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_compute_image_iam_policy.
func (ciip dataComputeImageIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("etag"))
}

// Id returns a reference to field id of google_compute_image_iam_policy.
func (ciip dataComputeImageIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("id"))
}

// Image returns a reference to field image of google_compute_image_iam_policy.
func (ciip dataComputeImageIamPolicyAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("image"))
}

// PolicyData returns a reference to field policy_data of google_compute_image_iam_policy.
func (ciip dataComputeImageIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_compute_image_iam_policy.
func (ciip dataComputeImageIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ciip.ref.Append("project"))
}
