// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataHealthcareHl7V2StoreIamPolicy creates a new instance of [DataHealthcareHl7V2StoreIamPolicy].
func NewDataHealthcareHl7V2StoreIamPolicy(name string, args DataHealthcareHl7V2StoreIamPolicyArgs) *DataHealthcareHl7V2StoreIamPolicy {
	return &DataHealthcareHl7V2StoreIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHealthcareHl7V2StoreIamPolicy)(nil)

// DataHealthcareHl7V2StoreIamPolicy represents the Terraform data resource google_healthcare_hl7_v2_store_iam_policy.
type DataHealthcareHl7V2StoreIamPolicy struct {
	Name string
	Args DataHealthcareHl7V2StoreIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataHealthcareHl7V2StoreIamPolicy].
func (hhvsip *DataHealthcareHl7V2StoreIamPolicy) DataSource() string {
	return "google_healthcare_hl7_v2_store_iam_policy"
}

// LocalName returns the local name for [DataHealthcareHl7V2StoreIamPolicy].
func (hhvsip *DataHealthcareHl7V2StoreIamPolicy) LocalName() string {
	return hhvsip.Name
}

// Configuration returns the configuration (args) for [DataHealthcareHl7V2StoreIamPolicy].
func (hhvsip *DataHealthcareHl7V2StoreIamPolicy) Configuration() interface{} {
	return hhvsip.Args
}

// Attributes returns the attributes for [DataHealthcareHl7V2StoreIamPolicy].
func (hhvsip *DataHealthcareHl7V2StoreIamPolicy) Attributes() dataHealthcareHl7V2StoreIamPolicyAttributes {
	return dataHealthcareHl7V2StoreIamPolicyAttributes{ref: terra.ReferenceDataResource(hhvsip)}
}

// DataHealthcareHl7V2StoreIamPolicyArgs contains the configurations for google_healthcare_hl7_v2_store_iam_policy.
type DataHealthcareHl7V2StoreIamPolicyArgs struct {
	// Hl7V2StoreId: string, required
	Hl7V2StoreId terra.StringValue `hcl:"hl7_v2_store_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}
type dataHealthcareHl7V2StoreIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_healthcare_hl7_v2_store_iam_policy.
func (hhvsip dataHealthcareHl7V2StoreIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(hhvsip.ref.Append("etag"))
}

// Hl7V2StoreId returns a reference to field hl7_v2_store_id of google_healthcare_hl7_v2_store_iam_policy.
func (hhvsip dataHealthcareHl7V2StoreIamPolicyAttributes) Hl7V2StoreId() terra.StringValue {
	return terra.ReferenceAsString(hhvsip.ref.Append("hl7_v2_store_id"))
}

// Id returns a reference to field id of google_healthcare_hl7_v2_store_iam_policy.
func (hhvsip dataHealthcareHl7V2StoreIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hhvsip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_healthcare_hl7_v2_store_iam_policy.
func (hhvsip dataHealthcareHl7V2StoreIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(hhvsip.ref.Append("policy_data"))
}
