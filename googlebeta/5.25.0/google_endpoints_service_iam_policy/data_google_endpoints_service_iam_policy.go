// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_endpoints_service_iam_policy

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_endpoints_service_iam_policy.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gesip *DataSource) DataSource() string {
	return "google_endpoints_service_iam_policy"
}

// LocalName returns the local name for [DataSource].
func (gesip *DataSource) LocalName() string {
	return gesip.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gesip *DataSource) Configuration() interface{} {
	return gesip.Args
}

// Attributes returns the attributes for [DataSource].
func (gesip *DataSource) Attributes() dataGoogleEndpointsServiceIamPolicyAttributes {
	return dataGoogleEndpointsServiceIamPolicyAttributes{ref: terra.ReferenceDataSource(gesip)}
}

// DataArgs contains the configurations for google_endpoints_service_iam_policy.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}

type dataGoogleEndpointsServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_policy.
func (gesip dataGoogleEndpointsServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gesip.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_policy.
func (gesip dataGoogleEndpointsServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gesip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_endpoints_service_iam_policy.
func (gesip dataGoogleEndpointsServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gesip.ref.Append("policy_data"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_policy.
func (gesip dataGoogleEndpointsServiceIamPolicyAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(gesip.ref.Append("service_name"))
}
