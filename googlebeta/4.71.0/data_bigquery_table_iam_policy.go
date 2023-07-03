// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataBigqueryTableIamPolicy creates a new instance of [DataBigqueryTableIamPolicy].
func NewDataBigqueryTableIamPolicy(name string, args DataBigqueryTableIamPolicyArgs) *DataBigqueryTableIamPolicy {
	return &DataBigqueryTableIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataBigqueryTableIamPolicy)(nil)

// DataBigqueryTableIamPolicy represents the Terraform data resource google_bigquery_table_iam_policy.
type DataBigqueryTableIamPolicy struct {
	Name string
	Args DataBigqueryTableIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataBigqueryTableIamPolicy].
func (btip *DataBigqueryTableIamPolicy) DataSource() string {
	return "google_bigquery_table_iam_policy"
}

// LocalName returns the local name for [DataBigqueryTableIamPolicy].
func (btip *DataBigqueryTableIamPolicy) LocalName() string {
	return btip.Name
}

// Configuration returns the configuration (args) for [DataBigqueryTableIamPolicy].
func (btip *DataBigqueryTableIamPolicy) Configuration() interface{} {
	return btip.Args
}

// Attributes returns the attributes for [DataBigqueryTableIamPolicy].
func (btip *DataBigqueryTableIamPolicy) Attributes() dataBigqueryTableIamPolicyAttributes {
	return dataBigqueryTableIamPolicyAttributes{ref: terra.ReferenceDataResource(btip)}
}

// DataBigqueryTableIamPolicyArgs contains the configurations for google_bigquery_table_iam_policy.
type DataBigqueryTableIamPolicyArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// TableId: string, required
	TableId terra.StringValue `hcl:"table_id,attr" validate:"required"`
}
type dataBigqueryTableIamPolicyAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("project"))
}

// TableId returns a reference to field table_id of google_bigquery_table_iam_policy.
func (btip dataBigqueryTableIamPolicyAttributes) TableId() terra.StringValue {
	return terra.ReferenceAsString(btip.ref.Append("table_id"))
}
