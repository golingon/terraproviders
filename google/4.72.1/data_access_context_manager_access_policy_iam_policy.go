// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataAccessContextManagerAccessPolicyIamPolicy creates a new instance of [DataAccessContextManagerAccessPolicyIamPolicy].
func NewDataAccessContextManagerAccessPolicyIamPolicy(name string, args DataAccessContextManagerAccessPolicyIamPolicyArgs) *DataAccessContextManagerAccessPolicyIamPolicy {
	return &DataAccessContextManagerAccessPolicyIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAccessContextManagerAccessPolicyIamPolicy)(nil)

// DataAccessContextManagerAccessPolicyIamPolicy represents the Terraform data resource google_access_context_manager_access_policy_iam_policy.
type DataAccessContextManagerAccessPolicyIamPolicy struct {
	Name string
	Args DataAccessContextManagerAccessPolicyIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataAccessContextManagerAccessPolicyIamPolicy].
func (acmapip *DataAccessContextManagerAccessPolicyIamPolicy) DataSource() string {
	return "google_access_context_manager_access_policy_iam_policy"
}

// LocalName returns the local name for [DataAccessContextManagerAccessPolicyIamPolicy].
func (acmapip *DataAccessContextManagerAccessPolicyIamPolicy) LocalName() string {
	return acmapip.Name
}

// Configuration returns the configuration (args) for [DataAccessContextManagerAccessPolicyIamPolicy].
func (acmapip *DataAccessContextManagerAccessPolicyIamPolicy) Configuration() interface{} {
	return acmapip.Args
}

// Attributes returns the attributes for [DataAccessContextManagerAccessPolicyIamPolicy].
func (acmapip *DataAccessContextManagerAccessPolicyIamPolicy) Attributes() dataAccessContextManagerAccessPolicyIamPolicyAttributes {
	return dataAccessContextManagerAccessPolicyIamPolicyAttributes{ref: terra.ReferenceDataResource(acmapip)}
}

// DataAccessContextManagerAccessPolicyIamPolicyArgs contains the configurations for google_access_context_manager_access_policy_iam_policy.
type DataAccessContextManagerAccessPolicyIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type dataAccessContextManagerAccessPolicyIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_access_context_manager_access_policy_iam_policy.
func (acmapip dataAccessContextManagerAccessPolicyIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(acmapip.ref.Append("etag"))
}

// Id returns a reference to field id of google_access_context_manager_access_policy_iam_policy.
func (acmapip dataAccessContextManagerAccessPolicyIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acmapip.ref.Append("id"))
}

// Name returns a reference to field name of google_access_context_manager_access_policy_iam_policy.
func (acmapip dataAccessContextManagerAccessPolicyIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acmapip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_access_context_manager_access_policy_iam_policy.
func (acmapip dataAccessContextManagerAccessPolicyIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(acmapip.ref.Append("policy_data"))
}
