// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayDocumentationVersion creates a new instance of [ApiGatewayDocumentationVersion].
func NewApiGatewayDocumentationVersion(name string, args ApiGatewayDocumentationVersionArgs) *ApiGatewayDocumentationVersion {
	return &ApiGatewayDocumentationVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayDocumentationVersion)(nil)

// ApiGatewayDocumentationVersion represents the Terraform resource aws_api_gateway_documentation_version.
type ApiGatewayDocumentationVersion struct {
	Name      string
	Args      ApiGatewayDocumentationVersionArgs
	state     *apiGatewayDocumentationVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) Type() string {
	return "aws_api_gateway_documentation_version"
}

// LocalName returns the local name for [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) LocalName() string {
	return agdv.Name
}

// Configuration returns the configuration (args) for [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) Configuration() interface{} {
	return agdv.Args
}

// DependOn is used for other resources to depend on [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(agdv)
}

// Dependencies returns the list of resources [ApiGatewayDocumentationVersion] depends_on.
func (agdv *ApiGatewayDocumentationVersion) Dependencies() terra.Dependencies {
	return agdv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) LifecycleManagement() *terra.Lifecycle {
	return agdv.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayDocumentationVersion].
func (agdv *ApiGatewayDocumentationVersion) Attributes() apiGatewayDocumentationVersionAttributes {
	return apiGatewayDocumentationVersionAttributes{ref: terra.ReferenceResource(agdv)}
}

// ImportState imports the given attribute values into [ApiGatewayDocumentationVersion]'s state.
func (agdv *ApiGatewayDocumentationVersion) ImportState(av io.Reader) error {
	agdv.state = &apiGatewayDocumentationVersionState{}
	if err := json.NewDecoder(av).Decode(agdv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agdv.Type(), agdv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayDocumentationVersion] has state.
func (agdv *ApiGatewayDocumentationVersion) State() (*apiGatewayDocumentationVersionState, bool) {
	return agdv.state, agdv.state != nil
}

// StateMust returns the state for [ApiGatewayDocumentationVersion]. Panics if the state is nil.
func (agdv *ApiGatewayDocumentationVersion) StateMust() *apiGatewayDocumentationVersionState {
	if agdv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agdv.Type(), agdv.LocalName()))
	}
	return agdv.state
}

// ApiGatewayDocumentationVersionArgs contains the configurations for aws_api_gateway_documentation_version.
type ApiGatewayDocumentationVersionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
}
type apiGatewayDocumentationVersionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of aws_api_gateway_documentation_version.
func (agdv apiGatewayDocumentationVersionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(agdv.ref.Append("description"))
}

// Id returns a reference to field id of aws_api_gateway_documentation_version.
func (agdv apiGatewayDocumentationVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agdv.ref.Append("id"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_documentation_version.
func (agdv apiGatewayDocumentationVersionAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agdv.ref.Append("rest_api_id"))
}

// Version returns a reference to field version of aws_api_gateway_documentation_version.
func (agdv apiGatewayDocumentationVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(agdv.ref.Append("version"))
}

type apiGatewayDocumentationVersionState struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	RestApiId   string `json:"rest_api_id"`
	Version     string `json:"version"`
}
