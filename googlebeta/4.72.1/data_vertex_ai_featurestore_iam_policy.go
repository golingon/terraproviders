// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataVertexAiFeaturestoreIamPolicy creates a new instance of [DataVertexAiFeaturestoreIamPolicy].
func NewDataVertexAiFeaturestoreIamPolicy(name string, args DataVertexAiFeaturestoreIamPolicyArgs) *DataVertexAiFeaturestoreIamPolicy {
	return &DataVertexAiFeaturestoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataVertexAiFeaturestoreIamPolicy)(nil)

// DataVertexAiFeaturestoreIamPolicy represents the Terraform data resource google_vertex_ai_featurestore_iam_policy.
type DataVertexAiFeaturestoreIamPolicy struct {
	Name string
	Args DataVertexAiFeaturestoreIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataVertexAiFeaturestoreIamPolicy].
func (vafip *DataVertexAiFeaturestoreIamPolicy) DataSource() string {
	return "google_vertex_ai_featurestore_iam_policy"
}

// LocalName returns the local name for [DataVertexAiFeaturestoreIamPolicy].
func (vafip *DataVertexAiFeaturestoreIamPolicy) LocalName() string {
	return vafip.Name
}

// Configuration returns the configuration (args) for [DataVertexAiFeaturestoreIamPolicy].
func (vafip *DataVertexAiFeaturestoreIamPolicy) Configuration() interface{} {
	return vafip.Args
}

// Attributes returns the attributes for [DataVertexAiFeaturestoreIamPolicy].
func (vafip *DataVertexAiFeaturestoreIamPolicy) Attributes() dataVertexAiFeaturestoreIamPolicyAttributes {
	return dataVertexAiFeaturestoreIamPolicyAttributes{ref: terra.ReferenceDataResource(vafip)}
}

// DataVertexAiFeaturestoreIamPolicyArgs contains the configurations for google_vertex_ai_featurestore_iam_policy.
type DataVertexAiFeaturestoreIamPolicyArgs struct {
	// Featurestore: string, required
	Featurestore terra.StringValue `hcl:"featurestore,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
}
type dataVertexAiFeaturestoreIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("etag"))
}

// Featurestore returns a reference to field featurestore of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) Featurestore() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("featurestore"))
}

// Id returns a reference to field id of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("project"))
}

// Region returns a reference to field region of google_vertex_ai_featurestore_iam_policy.
func (vafip dataVertexAiFeaturestoreIamPolicyAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(vafip.ref.Append("region"))
}
