// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataHealthcareConsentStoreIamPolicy creates a new instance of [DataHealthcareConsentStoreIamPolicy].
func NewDataHealthcareConsentStoreIamPolicy(name string, args DataHealthcareConsentStoreIamPolicyArgs) *DataHealthcareConsentStoreIamPolicy {
	return &DataHealthcareConsentStoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareConsentStoreIamPolicy)(nil)

// DataHealthcareConsentStoreIamPolicy represents the Terraform data resource google_healthcare_consent_store_iam_policy.
type DataHealthcareConsentStoreIamPolicy struct {
	Name string
	Args DataHealthcareConsentStoreIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataHealthcareConsentStoreIamPolicy].
func (hcsip *DataHealthcareConsentStoreIamPolicy) DataSource() string {
	return "google_healthcare_consent_store_iam_policy"
}

// LocalName returns the local name for [DataHealthcareConsentStoreIamPolicy].
func (hcsip *DataHealthcareConsentStoreIamPolicy) LocalName() string {
	return hcsip.Name
}

// Configuration returns the configuration (args) for [DataHealthcareConsentStoreIamPolicy].
func (hcsip *DataHealthcareConsentStoreIamPolicy) Configuration() interface{} {
	return hcsip.Args
}

// Attributes returns the attributes for [DataHealthcareConsentStoreIamPolicy].
func (hcsip *DataHealthcareConsentStoreIamPolicy) Attributes() dataHealthcareConsentStoreIamPolicyAttributes {
	return dataHealthcareConsentStoreIamPolicyAttributes{ref: terra.ReferenceDataResource(hcsip)}
}

// DataHealthcareConsentStoreIamPolicyArgs contains the configurations for google_healthcare_consent_store_iam_policy.
type DataHealthcareConsentStoreIamPolicyArgs struct {
	// ConsentStoreId: string, required
	ConsentStoreId terra.StringValue `hcl:"consent_store_id,attr" validate:"required"`
	// Dataset: string, required
	Dataset terra.StringValue `hcl:"dataset,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataHealthcareConsentStoreIamPolicyAttributes struct {
	ref terra.Reference
}

// ConsentStoreId returns a reference to field consent_store_id of google_healthcare_consent_store_iam_policy.
func (hcsip dataHealthcareConsentStoreIamPolicyAttributes) ConsentStoreId() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("consent_store_id"))
}

// Dataset returns a reference to field dataset of google_healthcare_consent_store_iam_policy.
func (hcsip dataHealthcareConsentStoreIamPolicyAttributes) Dataset() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("dataset"))
}

// Etag returns a reference to field etag of google_healthcare_consent_store_iam_policy.
func (hcsip dataHealthcareConsentStoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("etag"))
}

// Id returns a reference to field id of google_healthcare_consent_store_iam_policy.
func (hcsip dataHealthcareConsentStoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_consent_store_iam_policy.
func (hcsip dataHealthcareConsentStoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hcsip.ref.Append("policy_data"))
}
