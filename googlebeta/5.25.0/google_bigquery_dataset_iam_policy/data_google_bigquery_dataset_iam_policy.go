// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_bigquery_dataset_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_bigquery_dataset_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gbdip *DataSource) DataSource() string {
	return "google_bigquery_dataset_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gbdip *DataSource) LocalName() string {
	return gbdip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gbdip *DataSource) Configuration() interface{} {
	return gbdip.Args
}

// Attributes returns the attributes for [DataSource].
func (gbdip *DataSource) Attributes() dataGoogleBigqueryDatasetIamPolicyAttributes {
	return dataGoogleBigqueryDatasetIamPolicyAttributes{ref: terra.ReferenceDataSource(gbdip)}
}

// DataArgs contains the configurations for google_bigquery_dataset_iam_policy.
type DataArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}

type dataGoogleBigqueryDatasetIamPolicyAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_bigquery_dataset_iam_policy.
func (gbdip dataGoogleBigqueryDatasetIamPolicyAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(gbdip.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_bigquery_dataset_iam_policy.
func (gbdip dataGoogleBigqueryDatasetIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbdip.ref.Append("etag"))
}

// Id returns a reference to field id of google_bigquery_dataset_iam_policy.
func (gbdip dataGoogleBigqueryDatasetIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbdip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_bigquery_dataset_iam_policy.
func (gbdip dataGoogleBigqueryDatasetIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gbdip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_bigquery_dataset_iam_policy.
func (gbdip dataGoogleBigqueryDatasetIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbdip.ref.Append("project"))
}
