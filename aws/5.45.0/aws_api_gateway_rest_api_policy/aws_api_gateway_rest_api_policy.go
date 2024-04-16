// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_api_gateway_rest_api_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_api_gateway_rest_api_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApiGatewayRestApiPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aagrap *Resource) Type() string {
	return "aws_api_gateway_rest_api_policy"
}

// LocalName returns the local name for [Resource].
func (aagrap *Resource) LocalName() string {
	return aagrap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aagrap *Resource) Configuration() interface{} {
	return aagrap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aagrap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aagrap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aagrap *Resource) Dependencies() terra.Dependencies {
	return aagrap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aagrap *Resource) LifecycleManagement() *terra.Lifecycle {
	return aagrap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aagrap *Resource) Attributes() awsApiGatewayRestApiPolicyAttributes {
	return awsApiGatewayRestApiPolicyAttributes{ref: terra.ReferenceResource(aagrap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aagrap *Resource) ImportState(state io.Reader) error {
	aagrap.state = &awsApiGatewayRestApiPolicyState{}
	if err := json.NewDecoder(state).Decode(aagrap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aagrap.Type(), aagrap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aagrap *Resource) State() (*awsApiGatewayRestApiPolicyState, bool) {
	return aagrap.state, aagrap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aagrap *Resource) StateMust() *awsApiGatewayRestApiPolicyState {
	if aagrap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aagrap.Type(), aagrap.LocalName()))
	}
	return aagrap.state
}

// Args contains the configurations for aws_api_gateway_rest_api_policy.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
}

type awsApiGatewayRestApiPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_api_gateway_rest_api_policy.
func (aagrap awsApiGatewayRestApiPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aagrap.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_api_gateway_rest_api_policy.
func (aagrap awsApiGatewayRestApiPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(aagrap.ref.Append("policy"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_rest_api_policy.
func (aagrap awsApiGatewayRestApiPolicyAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(aagrap.ref.Append("rest_api_id"))
}

type awsApiGatewayRestApiPolicyState struct {
	Id        string `json:"id"`
	Policy    string `json:"policy"`
	RestApiId string `json:"rest_api_id"`
}
