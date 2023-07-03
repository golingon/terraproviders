// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataTagsTagValueIamPolicy creates a new instance of [DataTagsTagValueIamPolicy].
func NewDataTagsTagValueIamPolicy(name string, args DataTagsTagValueIamPolicyArgs) *DataTagsTagValueIamPolicy {
	return &DataTagsTagValueIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataTagsTagValueIamPolicy)(nil)

// DataTagsTagValueIamPolicy represents the Terraform data resource google_tags_tag_value_iam_policy.
type DataTagsTagValueIamPolicy struct {
	Name string
	Args DataTagsTagValueIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataTagsTagValueIamPolicy].
func (ttvip *DataTagsTagValueIamPolicy) DataSource() string {
	return "google_tags_tag_value_iam_policy"
}

// LocalName returns the local name for [DataTagsTagValueIamPolicy].
func (ttvip *DataTagsTagValueIamPolicy) LocalName() string {
	return ttvip.Name
}

// Configuration returns the configuration (args) for [DataTagsTagValueIamPolicy].
func (ttvip *DataTagsTagValueIamPolicy) Configuration() interface{} {
	return ttvip.Args
}

// Attributes returns the attributes for [DataTagsTagValueIamPolicy].
func (ttvip *DataTagsTagValueIamPolicy) Attributes() dataTagsTagValueIamPolicyAttributes {
	return dataTagsTagValueIamPolicyAttributes{ref: terra.ReferenceDataResource(ttvip)}
}

// DataTagsTagValueIamPolicyArgs contains the configurations for google_tags_tag_value_iam_policy.
type DataTagsTagValueIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// TagValue: string, required
	TagValue terra.StringValue `hcl:"tag_value,attr" validate:"required"`
}
type dataTagsTagValueIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_tags_tag_value_iam_policy.
func (ttvip dataTagsTagValueIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("etag"))
}

// Id returns a reference to field id of google_tags_tag_value_iam_policy.
func (ttvip dataTagsTagValueIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_tags_tag_value_iam_policy.
func (ttvip dataTagsTagValueIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("policy_data"))
}

// TagValue returns a reference to field tag_value of google_tags_tag_value_iam_policy.
func (ttvip dataTagsTagValueIamPolicyAttributes) TagValue() terra.StringValue {
	return terra.ReferenceAsString(ttvip.ref.Append("tag_value"))
}
