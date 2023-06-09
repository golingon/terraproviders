// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataBillingAccountIamPolicy creates a new instance of [DataBillingAccountIamPolicy].
func NewDataBillingAccountIamPolicy(name string, args DataBillingAccountIamPolicyArgs) *DataBillingAccountIamPolicy {
	return &DataBillingAccountIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBillingAccountIamPolicy)(nil)

// DataBillingAccountIamPolicy represents the Terraform data resource google_billing_account_iam_policy.
type DataBillingAccountIamPolicy struct {
	Name string
	Args DataBillingAccountIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBillingAccountIamPolicy].
func (baip *DataBillingAccountIamPolicy) DataSource() string {
	return "google_billing_account_iam_policy"
}

// LocalName returns the local name for [DataBillingAccountIamPolicy].
func (baip *DataBillingAccountIamPolicy) LocalName() string {
	return baip.Name
}

// Configuration returns the configuration (args) for [DataBillingAccountIamPolicy].
func (baip *DataBillingAccountIamPolicy) Configuration() interface{} {
	return baip.Args
}

// Attributes returns the attributes for [DataBillingAccountIamPolicy].
func (baip *DataBillingAccountIamPolicy) Attributes() dataBillingAccountIamPolicyAttributes {
	return dataBillingAccountIamPolicyAttributes{ref: terra.ReferenceDataResource(baip)}
}

// DataBillingAccountIamPolicyArgs contains the configurations for google_billing_account_iam_policy.
type DataBillingAccountIamPolicyArgs struct {
	// BillingAccountId: string, required
	BillingAccountId terra.StringValue `hcl:"billing_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataBillingAccountIamPolicyAttributes struct {
	ref terra.Reference
}

// BillingAccountId returns a reference to field billing_account_id of google_billing_account_iam_policy.
func (baip dataBillingAccountIamPolicyAttributes) BillingAccountId() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("billing_account_id"))
}

// Etag returns a reference to field etag of google_billing_account_iam_policy.
func (baip dataBillingAccountIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("etag"))
}

// Id returns a reference to field id of google_billing_account_iam_policy.
func (baip dataBillingAccountIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_billing_account_iam_policy.
func (baip dataBillingAccountIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(baip.ref.Append("policy_data"))
}
