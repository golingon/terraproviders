// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataHealthcareDicomStoreIamPolicy creates a new instance of [DataHealthcareDicomStoreIamPolicy].
func NewDataHealthcareDicomStoreIamPolicy(name string, args DataHealthcareDicomStoreIamPolicyArgs) *DataHealthcareDicomStoreIamPolicy {
	return &DataHealthcareDicomStoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareDicomStoreIamPolicy)(nil)

// DataHealthcareDicomStoreIamPolicy represents the Terraform data resource google_healthcare_dicom_store_iam_policy.
type DataHealthcareDicomStoreIamPolicy struct {
	Name string
	Args DataHealthcareDicomStoreIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataHealthcareDicomStoreIamPolicy].
func (hdsip *DataHealthcareDicomStoreIamPolicy) DataSource() string {
	return "google_healthcare_dicom_store_iam_policy"
}

// LocalName returns the local name for [DataHealthcareDicomStoreIamPolicy].
func (hdsip *DataHealthcareDicomStoreIamPolicy) LocalName() string {
	return hdsip.Name
}

// Configuration returns the configuration (args) for [DataHealthcareDicomStoreIamPolicy].
func (hdsip *DataHealthcareDicomStoreIamPolicy) Configuration() interface{} {
	return hdsip.Args
}

// Attributes returns the attributes for [DataHealthcareDicomStoreIamPolicy].
func (hdsip *DataHealthcareDicomStoreIamPolicy) Attributes() dataHealthcareDicomStoreIamPolicyAttributes {
	return dataHealthcareDicomStoreIamPolicyAttributes{ref: terra.ReferenceDataResource(hdsip)}
}

// DataHealthcareDicomStoreIamPolicyArgs contains the configurations for google_healthcare_dicom_store_iam_policy.
type DataHealthcareDicomStoreIamPolicyArgs struct {
	// DicomStoreId: string, required
	DicomStoreId terra.StringValue `hcl:"dicom_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataHealthcareDicomStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// DicomStoreId returns a reference to field dicom_store_id of google_healthcare_dicom_store_iam_policy.
func (hdsip dataHealthcareDicomStoreIamPolicyAttributes) DicomStoreId() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("dicom_store_id"))
}

// Etag returns a reference to field etag of google_healthcare_dicom_store_iam_policy.
func (hdsip dataHealthcareDicomStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_dicom_store_iam_policy.
func (hdsip dataHealthcareDicomStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_dicom_store_iam_policy.
func (hdsip dataHealthcareDicomStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hdsip.ref.Append("policy_data"))
}
