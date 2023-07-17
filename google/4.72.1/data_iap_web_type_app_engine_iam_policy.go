// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataIapWebTypeAppEngineIamPolicy creates a new instance of [DataIapWebTypeAppEngineIamPolicy].
func NewDataIapWebTypeAppEngineIamPolicy(name string, args DataIapWebTypeAppEngineIamPolicyArgs) *DataIapWebTypeAppEngineIamPolicy {
	return &DataIapWebTypeAppEngineIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataIapWebTypeAppEngineIamPolicy)(nil)

// DataIapWebTypeAppEngineIamPolicy represents the Terraform data resource google_iap_web_type_app_engine_iam_policy.
type DataIapWebTypeAppEngineIamPolicy struct {
	Name string
	Args DataIapWebTypeAppEngineIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataIapWebTypeAppEngineIamPolicy].
func (iwtaeip *DataIapWebTypeAppEngineIamPolicy) DataSource() string {
	return "google_iap_web_type_app_engine_iam_policy"
}

// LocalName returns the local name for [DataIapWebTypeAppEngineIamPolicy].
func (iwtaeip *DataIapWebTypeAppEngineIamPolicy) LocalName() string {
	return iwtaeip.Name
}

// Configuration returns the configuration (args) for [DataIapWebTypeAppEngineIamPolicy].
func (iwtaeip *DataIapWebTypeAppEngineIamPolicy) Configuration() interface{} {
	return iwtaeip.Args
}

// Attributes returns the attributes for [DataIapWebTypeAppEngineIamPolicy].
func (iwtaeip *DataIapWebTypeAppEngineIamPolicy) Attributes() dataIapWebTypeAppEngineIamPolicyAttributes {
	return dataIapWebTypeAppEngineIamPolicyAttributes{ref: terra.ReferenceDataResource(iwtaeip)}
}

// DataIapWebTypeAppEngineIamPolicyArgs contains the configurations for google_iap_web_type_app_engine_iam_policy.
type DataIapWebTypeAppEngineIamPolicyArgs struct {
	// AppId: string, required
	AppId terra.StringValue `hcl:"app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type dataIapWebTypeAppEngineIamPolicyAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of google_iap_web_type_app_engine_iam_policy.
func (iwtaeip dataIapWebTypeAppEngineIamPolicyAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(iwtaeip.ref.Append("app_id"))
}

// Etag returns a reference to field etag of google_iap_web_type_app_engine_iam_policy.
func (iwtaeip dataIapWebTypeAppEngineIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(iwtaeip.ref.Append("etag"))
}

// Id returns a reference to field id of google_iap_web_type_app_engine_iam_policy.
func (iwtaeip dataIapWebTypeAppEngineIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(iwtaeip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_iap_web_type_app_engine_iam_policy.
func (iwtaeip dataIapWebTypeAppEngineIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(iwtaeip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_iap_web_type_app_engine_iam_policy.
func (iwtaeip dataIapWebTypeAppEngineIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(iwtaeip.ref.Append("project"))
}