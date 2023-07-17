// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataGkeHubMembershipIamPolicy creates a new instance of [DataGkeHubMembershipIamPolicy].
func NewDataGkeHubMembershipIamPolicy(name string, args DataGkeHubMembershipIamPolicyArgs) *DataGkeHubMembershipIamPolicy {
	return &DataGkeHubMembershipIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataGkeHubMembershipIamPolicy)(nil)

// DataGkeHubMembershipIamPolicy represents the Terraform data resource google_gke_hub_membership_iam_policy.
type DataGkeHubMembershipIamPolicy struct {
	Name string
	Args DataGkeHubMembershipIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataGkeHubMembershipIamPolicy].
func (ghmip *DataGkeHubMembershipIamPolicy) DataSource() string {
	return "google_gke_hub_membership_iam_policy"
}

// LocalName returns the local name for [DataGkeHubMembershipIamPolicy].
func (ghmip *DataGkeHubMembershipIamPolicy) LocalName() string {
	return ghmip.Name
}

// Configuration returns the configuration (args) for [DataGkeHubMembershipIamPolicy].
func (ghmip *DataGkeHubMembershipIamPolicy) Configuration() interface{} {
	return ghmip.Args
}

// Attributes returns the attributes for [DataGkeHubMembershipIamPolicy].
func (ghmip *DataGkeHubMembershipIamPolicy) Attributes() dataGkeHubMembershipIamPolicyAttributes {
	return dataGkeHubMembershipIamPolicyAttributes{ref: terra.ReferenceDataResource(ghmip)}
}

// DataGkeHubMembershipIamPolicyArgs contains the configurations for google_gke_hub_membership_iam_policy.
type DataGkeHubMembershipIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MembershipId: string, required
	MembershipId terra.StringValue `hcl:"membership_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataGkeHubMembershipIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_hub_membership_iam_policy.
func (ghmip dataGkeHubMembershipIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghmip.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_hub_membership_iam_policy.
func (ghmip dataGkeHubMembershipIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghmip.ref.Append("id"))
}

// MembershipId returns a reference to field membership_id of google_gke_hub_membership_iam_policy.
func (ghmip dataGkeHubMembershipIamPolicyAttributes) MembershipId() terra.StringValue {
	return terra.ReferenceAsString(ghmip.ref.Append("membership_id"))
}

// PolicyData returns a reference to field policy_data of google_gke_hub_membership_iam_policy.
func (ghmip dataGkeHubMembershipIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ghmip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_gke_hub_membership_iam_policy.
func (ghmip dataGkeHubMembershipIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ghmip.ref.Append("project"))
}