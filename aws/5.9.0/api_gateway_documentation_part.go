// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewaydocumentationpart "github.com/golingon/terraproviders/aws/5.9.0/apigatewaydocumentationpart"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayDocumentationPart creates a new instance of [ApiGatewayDocumentationPart].
func NewApiGatewayDocumentationPart(name string, args ApiGatewayDocumentationPartArgs) *ApiGatewayDocumentationPart {
	return &ApiGatewayDocumentationPart{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayDocumentationPart)(nil)

// ApiGatewayDocumentationPart represents the Terraform resource aws_api_gateway_documentation_part.
type ApiGatewayDocumentationPart struct {
	Name      string
	Args      ApiGatewayDocumentationPartArgs
	state     *apiGatewayDocumentationPartState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) Type() string {
	return "aws_api_gateway_documentation_part"
}

// LocalName returns the local name for [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) LocalName() string {
	return agdp.Name
}

// Configuration returns the configuration (args) for [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) Configuration() interface{} {
	return agdp.Args
}

// DependOn is used for other resources to depend on [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) DependOn() terra.Reference {
	return terra.ReferenceResource(agdp)
}

// Dependencies returns the list of resources [ApiGatewayDocumentationPart] depends_on.
func (agdp *ApiGatewayDocumentationPart) Dependencies() terra.Dependencies {
	return agdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) LifecycleManagement() *terra.Lifecycle {
	return agdp.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayDocumentationPart].
func (agdp *ApiGatewayDocumentationPart) Attributes() apiGatewayDocumentationPartAttributes {
	return apiGatewayDocumentationPartAttributes{ref: terra.ReferenceResource(agdp)}
}

// ImportState imports the given attribute values into [ApiGatewayDocumentationPart]'s state.
func (agdp *ApiGatewayDocumentationPart) ImportState(av io.Reader) error {
	agdp.state = &apiGatewayDocumentationPartState{}
	if err := json.NewDecoder(av).Decode(agdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agdp.Type(), agdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayDocumentationPart] has state.
func (agdp *ApiGatewayDocumentationPart) State() (*apiGatewayDocumentationPartState, bool) {
	return agdp.state, agdp.state != nil
}

// StateMust returns the state for [ApiGatewayDocumentationPart]. Panics if the state is nil.
func (agdp *ApiGatewayDocumentationPart) StateMust() *apiGatewayDocumentationPartState {
	if agdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agdp.Type(), agdp.LocalName()))
	}
	return agdp.state
}

// ApiGatewayDocumentationPartArgs contains the configurations for aws_api_gateway_documentation_part.
type ApiGatewayDocumentationPartArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Properties: string, required
	Properties terra.StringValue `hcl:"properties,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// Location: required
	Location *apigatewaydocumentationpart.Location `hcl:"location,block" validate:"required"`
}
type apiGatewayDocumentationPartAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_api_gateway_documentation_part.
func (agdp apiGatewayDocumentationPartAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agdp.ref.Append("id"))
}

// Properties returns a reference to field properties of aws_api_gateway_documentation_part.
func (agdp apiGatewayDocumentationPartAttributes) Properties() terra.StringValue {
	return terra.ReferenceAsString(agdp.ref.Append("properties"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_documentation_part.
func (agdp apiGatewayDocumentationPartAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agdp.ref.Append("rest_api_id"))
}

func (agdp apiGatewayDocumentationPartAttributes) Location() terra.ListValue[apigatewaydocumentationpart.LocationAttributes] {
	return terra.ReferenceAsList[apigatewaydocumentationpart.LocationAttributes](agdp.ref.Append("location"))
}

type apiGatewayDocumentationPartState struct {
	Id         string                                      `json:"id"`
	Properties string                                      `json:"properties"`
	RestApiId  string                                      `json:"rest_api_id"`
	Location   []apigatewaydocumentationpart.LocationState `json:"location"`
}
