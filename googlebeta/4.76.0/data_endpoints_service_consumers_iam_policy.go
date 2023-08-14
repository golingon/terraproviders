// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataEndpointsServiceConsumersIamPolicy creates a new instance of [DataEndpointsServiceConsumersIamPolicy].
func NewDataEndpointsServiceConsumersIamPolicy(name string, args DataEndpointsServiceConsumersIamPolicyArgs) *DataEndpointsServiceConsumersIamPolicy {
	return &DataEndpointsServiceConsumersIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEndpointsServiceConsumersIamPolicy)(nil)

// DataEndpointsServiceConsumersIamPolicy represents the Terraform data resource google_endpoints_service_consumers_iam_policy.
type DataEndpointsServiceConsumersIamPolicy struct {
	Name string
	Args DataEndpointsServiceConsumersIamPolicyArgs
}

// DataSource returns the Terraform object type for [DataEndpointsServiceConsumersIamPolicy].
func (escip *DataEndpointsServiceConsumersIamPolicy) DataSource() string {
	return "google_endpoints_service_consumers_iam_policy"
}

// LocalName returns the local name for [DataEndpointsServiceConsumersIamPolicy].
func (escip *DataEndpointsServiceConsumersIamPolicy) LocalName() string {
	return escip.Name
}

// Configuration returns the configuration (args) for [DataEndpointsServiceConsumersIamPolicy].
func (escip *DataEndpointsServiceConsumersIamPolicy) Configuration() interface{} {
	return escip.Args
}

// Attributes returns the attributes for [DataEndpointsServiceConsumersIamPolicy].
func (escip *DataEndpointsServiceConsumersIamPolicy) Attributes() dataEndpointsServiceConsumersIamPolicyAttributes {
	return dataEndpointsServiceConsumersIamPolicyAttributes{ref: terra.ReferenceDataResource(escip)}
}

// DataEndpointsServiceConsumersIamPolicyArgs contains the configurations for google_endpoints_service_consumers_iam_policy.
type DataEndpointsServiceConsumersIamPolicyArgs struct {
	// ConsumerProject: string, required
	ConsumerProject terra.StringValue `hcl:"consumer_project,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}
type dataEndpointsServiceConsumersIamPolicyAttributes struct {
	ref terra.Reference
}

// ConsumerProject returns a reference to field consumer_project of google_endpoints_service_consumers_iam_policy.
func (escip dataEndpointsServiceConsumersIamPolicyAttributes) ConsumerProject() terra.StringValue {
	return terra.ReferenceAsString(escip.ref.Append("consumer_project"))
}

// Etag returns a reference to field etag of google_endpoints_service_consumers_iam_policy.
func (escip dataEndpointsServiceConsumersIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(escip.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_consumers_iam_policy.
func (escip dataEndpointsServiceConsumersIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(escip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_endpoints_service_consumers_iam_policy.
func (escip dataEndpointsServiceConsumersIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(escip.ref.Append("policy_data"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_consumers_iam_policy.
func (escip dataEndpointsServiceConsumersIamPolicyAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(escip.ref.Append("service_name"))
}