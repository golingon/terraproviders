// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataOrganizationIamPolicy creates a new instance of [DataOrganizationIamPolicy].
func NewDataOrganizationIamPolicy(name string, args DataOrganizationIamPolicyArgs) *DataOrganizationIamPolicy {
	return &DataOrganizationIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataOrganizationIamPolicy)(nil)

// DataOrganizationIamPolicy represents the Terraform data resource google_organization_iam_policy.
type DataOrganizationIamPolicy struct {
	Name string
	Args DataOrganizationIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataOrganizationIamPolicy].
func (oip *DataOrganizationIamPolicy) DataSource() string {
	return "google_organization_iam_policy"
}

// LocalName returns the local name for [DataOrganizationIamPolicy].
func (oip *DataOrganizationIamPolicy) LocalName() string {
	return oip.Name
}

// Configuration returns the configuration (args) for [DataOrganizationIamPolicy].
func (oip *DataOrganizationIamPolicy) Configuration() interface{} {
	return oip.Args
}

// Attributes returns the attributes for [DataOrganizationIamPolicy].
func (oip *DataOrganizationIamPolicy) Attributes() dataOrganizationIamPolicyAttributes {
	return dataOrganizationIamPolicyAttributes{ref: terra.ReferenceDataResource(oip)}
}

// DataOrganizationIamPolicyArgs contains the configurations for google_organization_iam_policy.
type DataOrganizationIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// OrgId: string, required
	OrgId terra.StringValue `hcl:"org_id,attr" validate:"required"`
}
type dataOrganizationIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_organization_iam_policy.
func (oip dataOrganizationIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(oip.ref.Append("etag"))
}

// Id returns a reference to field id of google_organization_iam_policy.
func (oip dataOrganizationIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(oip.ref.Append("id"))
}

// OrgId returns a reference to field org_id of google_organization_iam_policy.
func (oip dataOrganizationIamPolicyAttributes) OrgId() terra.StringValue {
	return terra.ReferenceAsString(oip.ref.Append("org_id"))
}

// PolicyData returns a reference to field policy_data of google_organization_iam_policy.
func (oip dataOrganizationIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(oip.ref.Append("policy_data"))
}
