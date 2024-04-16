// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_api_gateway_documentation_version

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

// Resource represents the Terraform resource aws_api_gateway_documentation_version.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApiGatewayDocumentationVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aagdv *Resource) Type() string {
	return "aws_api_gateway_documentation_version"
}

// LocalName returns the local name for [Resource].
func (aagdv *Resource) LocalName() string {
	return aagdv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aagdv *Resource) Configuration() interface{} {
	return aagdv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aagdv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aagdv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aagdv *Resource) Dependencies() terra.Dependencies {
	return aagdv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aagdv *Resource) LifecycleManagement() *terra.Lifecycle {
	return aagdv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aagdv *Resource) Attributes() awsApiGatewayDocumentationVersionAttributes {
	return awsApiGatewayDocumentationVersionAttributes{ref: terra.ReferenceResource(aagdv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aagdv *Resource) ImportState(state io.Reader) error {
	aagdv.state = &awsApiGatewayDocumentationVersionState{}
	if err := json.NewDecoder(state).Decode(aagdv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aagdv.Type(), aagdv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aagdv *Resource) State() (*awsApiGatewayDocumentationVersionState, bool) {
	return aagdv.state, aagdv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aagdv *Resource) StateMust() *awsApiGatewayDocumentationVersionState {
	if aagdv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aagdv.Type(), aagdv.LocalName()))
	}
	return aagdv.state
}

// Args contains the configurations for aws_api_gateway_documentation_version.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}

type awsApiGatewayDocumentationVersionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_api_gateway_documentation_version.
func (aagdv awsApiGatewayDocumentationVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(aagdv.ref.Append("description"))
}

// Id returns a reference to field id of aws_api_gateway_documentation_version.
func (aagdv awsApiGatewayDocumentationVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aagdv.ref.Append("id"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_documentation_version.
func (aagdv awsApiGatewayDocumentationVersionAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(aagdv.ref.Append("rest_api_id"))
}

// Version returns a reference to field version of aws_api_gateway_documentation_version.
func (aagdv awsApiGatewayDocumentationVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(aagdv.ref.Append("version"))
}

type awsApiGatewayDocumentationVersionState struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	RestApiId   string `json:"rest_api_id"`
	Version     string `json:"version"`
}
