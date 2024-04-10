// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewEndpointsServiceIamPolicy creates a new instance of [EndpointsServiceIamPolicy].
func NewEndpointsServiceIamPolicy(name string, args EndpointsServiceIamPolicyArgs) *EndpointsServiceIamPolicy {
	return &EndpointsServiceIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EndpointsServiceIamPolicy)(nil)

// EndpointsServiceIamPolicy represents the Terraform resource google_endpoints_service_iam_policy.
type EndpointsServiceIamPolicy struct {
	Name      string
	Args      EndpointsServiceIamPolicyArgs
	state     *endpointsServiceIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) Type() string {
	return "google_endpoints_service_iam_policy"
}

// LocalName returns the local name for [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) LocalName() string {
	return esip.Name
}

// Configuration returns the configuration (args) for [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) Configuration() interface{} {
	return esip.Args
}

// DependOn is used for other resources to depend on [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(esip)
}

// Dependencies returns the list of resources [EndpointsServiceIamPolicy] depends_on.
func (esip *EndpointsServiceIamPolicy) Dependencies() terra.Dependencies {
	return esip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return esip.Lifecycle
}

// Attributes returns the attributes for [EndpointsServiceIamPolicy].
func (esip *EndpointsServiceIamPolicy) Attributes() endpointsServiceIamPolicyAttributes {
	return endpointsServiceIamPolicyAttributes{ref: terra.ReferenceResource(esip)}
}

// ImportState imports the given attribute values into [EndpointsServiceIamPolicy]'s state.
func (esip *EndpointsServiceIamPolicy) ImportState(av io.Reader) error {
	esip.state = &endpointsServiceIamPolicyState{}
	if err := json.NewDecoder(av).Decode(esip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", esip.Type(), esip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EndpointsServiceIamPolicy] has state.
func (esip *EndpointsServiceIamPolicy) State() (*endpointsServiceIamPolicyState, bool) {
	return esip.state, esip.state != nil
}

// StateMust returns the state for [EndpointsServiceIamPolicy]. Panics if the state is nil.
func (esip *EndpointsServiceIamPolicy) StateMust() *endpointsServiceIamPolicyState {
	if esip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", esip.Type(), esip.LocalName()))
	}
	return esip.state
}

// EndpointsServiceIamPolicyArgs contains the configurations for google_endpoints_service_iam_policy.
type EndpointsServiceIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}
type endpointsServiceIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_endpoints_service_iam_policy.
func (esip endpointsServiceIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("etag"))
}

// Id returns a reference to field id of google_endpoints_service_iam_policy.
func (esip endpointsServiceIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_endpoints_service_iam_policy.
func (esip endpointsServiceIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("policy_data"))
}

// ServiceName returns a reference to field service_name of google_endpoints_service_iam_policy.
func (esip endpointsServiceIamPolicyAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(esip.ref.Append("service_name"))
}

type endpointsServiceIamPolicyState struct {
	Etag        string `json:"etag"`
	Id          string `json:"id"`
	PolicyData  string `json:"policy_data"`
	ServiceName string `json:"service_name"`
}
