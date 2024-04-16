// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_healthcare_dicom_store_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_healthcare_dicom_store_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ghdsip *DataSource) DataSource() string {
	return "google_healthcare_dicom_store_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (ghdsip *DataSource) LocalName() string {
	return ghdsip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ghdsip *DataSource) Configuration() interface{} {
	return ghdsip.Args
}

// Attributes returns the attributes for [DataSource].
func (ghdsip *DataSource) Attributes() dataGoogleHealthcareDicomStoreIamPolicyAttributes {
	return dataGoogleHealthcareDicomStoreIamPolicyAttributes{ref: terra.ReferenceDataSource(ghdsip)}
}

// DataArgs contains the configurations for google_healthcare_dicom_store_iam_policy.
type DataArgs struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataGoogleHealthcareDicomStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_policy.
func (ghdsip dataGoogleHealthcareDicomStoreIamPolicyAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_policy.
func (ghdsip dataGoogleHealthcareDicomStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_policy.
func (ghdsip dataGoogleHealthcareDicomStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_dicom_store_iam_policy.
func (ghdsip dataGoogleHealthcareDicomStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(ghdsip.ref.Append("policy_data"))
}
