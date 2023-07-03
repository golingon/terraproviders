// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataHealthcareDatasetIamPolicy creates a new instance of [DataHealthcareDatasetIamPolicy].
func NewDataHealthcareDatasetIamPolicy(name string, args DataHealthcareDatasetIamPolicyArgs) *DataHealthcareDatasetIamPolicy {
	return &DataHealthcareDatasetIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareDatasetIamPolicy)(nil)

// DataHealthcareDatasetIamPolicy represents the Terraform data resource google_healthcare_dataset_iam_policy.
type DataHealthcareDatasetIamPolicy struct {
	Name string
	Args DataHealthcareDatasetIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataHealthcareDatasetIamPolicy].
func (hdip *DataHealthcareDatasetIamPolicy) DataSource() string {
	return "google_healthcare_dataset_iam_policy"
}

// LocalName returns the local name for [DataHealthcareDatasetIamPolicy].
func (hdip *DataHealthcareDatasetIamPolicy) LocalName() string {
	return hdip.Name
}

// Configuration returns the configuration (args) for [DataHealthcareDatasetIamPolicy].
func (hdip *DataHealthcareDatasetIamPolicy) Configuration() interface{} {
	return hdip.Args
}

// Attributes returns the attributes for [DataHealthcareDatasetIamPolicy].
func (hdip *DataHealthcareDatasetIamPolicy) Attributes() dataHealthcareDatasetIamPolicyAttributes {
	return dataHealthcareDatasetIamPolicyAttributes{ref: terra.ReferenceDataResource(hdip)}
}

// DataHealthcareDatasetIamPolicyArgs contains the configurations for google_healthcare_dataset_iam_policy.
type DataHealthcareDatasetIamPolicyArgs struct {
	// DatasetId: string, required
	DatasetId terra.StringValue `hcl:"dataset_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataHealthcareDatasetIamPolicyAttributes struct {
	ref terra.Reference
}

// DatasetId returns a reference to field dataset_id of google_healthcare_dataset_iam_policy.
func (hdip dataHealthcareDatasetIamPolicyAttributes) DatasetId() terra.StringValue {
	return terra.ReferenceAsString(hdip.ref.Append("dataset_id"))
}

// Etag returns a reference to field etag of google_healthcare_dataset_iam_policy.
func (hdip dataHealthcareDatasetIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dataset_iam_policy.
func (hdip dataHealthcareDatasetIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_dataset_iam_policy.
func (hdip dataHealthcareDatasetIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hdip.ref.Append("policy_data"))
}
