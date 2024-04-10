// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import "github.com/golingon/lingon/pkg/terra"

// NewDataEndpointsServiceIamPolicy creates a new instance of [DataEndpointsServiceIamPolicy].
func NewDataEndpointsServiceIamPolicy(name string, args DataEndpointsServiceIamPolicyArgs) *DataEndpointsServiceIamPolicy {
	return &DataEndpointsServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEndpointsServiceIamPolicy)(nil)

// DataEndpointsServiceIamPolicy represents the Terraform data resource google_endpoints_service_iam_policy.
type DataEndpointsServiceIamPolicy struct {
	Name string
	Args DataEndpointsServiceIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataEndpointsServiceIamPolicy].
func (esip *DataEndpointsServiceIamPolicy) DataSource() string {
	return "google_endpoints_service_iam_policy"
}

// LocalName returns the local name for [DataEndpointsServiceIamPolicy].
func (esip *DataEndpointsServiceIamPolicy) LocalName() string {
	return esip.Name
}

// Configuration returns the configuration (args) for [DataEndpointsServiceIamPolicy].
func (esip *DataEndpointsServiceIamPolicy) Configuration() interface{} {
	return esip.Args
}

// Attributes returns the attributes for [DataEndpointsServiceIamPolicy].
func (esip *DataEndpointsServiceIamPolicy) Attributes() dataEndpointsServiceIamPolicyAttributes {
	return dataEndpointsServiceIamPolicyAttributes{ref: terra.ReferenceDataResource(esip)}
}

// DataEndpointsServiceIamPolicyArgs contains the configurations for google_endpoints_service_iam_policy.
type DataEndpointsServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}
type dataEndpointsServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_policy.
func (esip dataEndpointsServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_policy.
func (esip dataEndpointsServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_endpoints_service_iam_policy.
func (esip dataEndpointsServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("policy_data"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_policy.
func (esip dataEndpointsServiceIamPolicyAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("service_name"))
}
