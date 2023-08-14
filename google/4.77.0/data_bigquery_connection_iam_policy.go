// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataBigqueryConnectionIamPolicy creates a new instance of [DataBigqueryConnectionIamPolicy].
func NewDataBigqueryConnectionIamPolicy(name string, args DataBigqueryConnectionIamPolicyArgs) *DataBigqueryConnectionIamPolicy {
	return &DataBigqueryConnectionIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigqueryConnectionIamPolicy)(nil)

// DataBigqueryConnectionIamPolicy represents the Terraform data resource google_bigquery_connection_iam_policy.
type DataBigqueryConnectionIamPolicy struct {
	Name string
	Args DataBigqueryConnectionIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBigqueryConnectionIamPolicy].
func (bcip *DataBigqueryConnectionIamPolicy) DataSource() string {
	return "google_bigquery_connection_iam_policy"
}

// LocalName returns the local name for [DataBigqueryConnectionIamPolicy].
func (bcip *DataBigqueryConnectionIamPolicy) LocalName() string {
	return bcip.Name
}

// Configuration returns the configuration (args) for [DataBigqueryConnectionIamPolicy].
func (bcip *DataBigqueryConnectionIamPolicy) Configuration() interface{} {
	return bcip.Args
}

// Attributes returns the attributes for [DataBigqueryConnectionIamPolicy].
func (bcip *DataBigqueryConnectionIamPolicy) Attributes() dataBigqueryConnectionIamPolicyAttributes {
	return dataBigqueryConnectionIamPolicyAttributes{ref: terra.ReferenceDataResource(bcip)}
}

// DataBigqueryConnectionIamPolicyArgs contains the configurations for google_bigquery_connection_iam_policy.
type DataBigqueryConnectionIamPolicyArgs struct {
	// ConnectionId: string, required
	ConnectionId terra.StringValue `hcl:"connection_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataBigqueryConnectionIamPolicyAttributes struct {
	ref terra.Reference
}

// ConnectionId returns a reference to field connection_id of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) ConnectionId() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("connection_id"))
}

// Etag returns a reference to field etag of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("id"))
}

// Location returns a reference to field location of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("location"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_connection_iam_policy.
func (bcip dataBigqueryConnectionIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bcip.ref.Append("project"))
}
