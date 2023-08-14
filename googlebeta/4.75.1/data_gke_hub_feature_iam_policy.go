// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataGkeHubFeatureIamPolicy creates a new instance of [DataGkeHubFeatureIamPolicy].
func NewDataGkeHubFeatureIamPolicy(name string, args DataGkeHubFeatureIamPolicyArgs) *DataGkeHubFeatureIamPolicy {
	return &DataGkeHubFeatureIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGkeHubFeatureIamPolicy)(nil)

// DataGkeHubFeatureIamPolicy represents the Terraform data resource google_gke_hub_feature_iam_policy.
type DataGkeHubFeatureIamPolicy struct {
	Name string
	Args DataGkeHubFeatureIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataGkeHubFeatureIamPolicy].
func (ghfip *DataGkeHubFeatureIamPolicy) DataSource() string {
	return "google_gke_hub_feature_iam_policy"
}

// LocalName returns the local name for [DataGkeHubFeatureIamPolicy].
func (ghfip *DataGkeHubFeatureIamPolicy) LocalName() string {
	return ghfip.Name
}

// Configuration returns the configuration (args) for [DataGkeHubFeatureIamPolicy].
func (ghfip *DataGkeHubFeatureIamPolicy) Configuration() interface{} {
	return ghfip.Args
}

// Attributes returns the attributes for [DataGkeHubFeatureIamPolicy].
func (ghfip *DataGkeHubFeatureIamPolicy) Attributes() dataGkeHubFeatureIamPolicyAttributes {
	return dataGkeHubFeatureIamPolicyAttributes{ref: terra.ReferenceDataResource(ghfip)}
}

// DataGkeHubFeatureIamPolicyArgs contains the configurations for google_gke_hub_feature_iam_policy.
type DataGkeHubFeatureIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataGkeHubFeatureIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("location"))
}

// Name returns a reference to field name of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_gke_hub_feature_iam_policy.
func (ghfip dataGkeHubFeatureIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghfip.ref.Append("project"))
}