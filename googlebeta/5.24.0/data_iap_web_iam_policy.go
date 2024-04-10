// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataIapWebIamPolicy creates a new instance of [DataIapWebIamPolicy].
func NewDataIapWebIamPolicy(name string, args DataIapWebIamPolicyArgs) *DataIapWebIamPolicy {
	return &DataIapWebIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIapWebIamPolicy)(nil)

// DataIapWebIamPolicy represents the Terraform data resource google_iap_web_iam_policy.
type DataIapWebIamPolicy struct {
	Name string
	Args DataIapWebIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIapWebIamPolicy].
func (iwip *DataIapWebIamPolicy) DataSource() string {
	return "google_iap_web_iam_policy"
}

// LocalName returns the local name for [DataIapWebIamPolicy].
func (iwip *DataIapWebIamPolicy) LocalName() string {
	return iwip.Name
}

// Configuration returns the configuration (args) for [DataIapWebIamPolicy].
func (iwip *DataIapWebIamPolicy) Configuration() interface{} {
	return iwip.Args
}

// Attributes returns the attributes for [DataIapWebIamPolicy].
func (iwip *DataIapWebIamPolicy) Attributes() dataIapWebIamPolicyAttributes {
	return dataIapWebIamPolicyAttributes{ref: terra.ReferenceDataResource(iwip)}
}

// DataIapWebIamPolicyArgs contains the configurations for google_iap_web_iam_policy.
type DataIapWebIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataIapWebIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_iap_web_iam_policy.
func (iwip dataIapWebIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_iam_policy.
func (iwip dataIapWebIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_iam_policy.
func (iwip dataIapWebIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_iam_policy.
func (iwip dataIapWebIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwip.ref.Append("project"))
}
