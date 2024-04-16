// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_apigatewayv2_api_mapping

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

// Resource represents the Terraform resource aws_apigatewayv2_api_mapping.
type Resource struct {
	Name      string
	Args      Args
	state     *awsApigatewayv2ApiMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aaam *Resource) Type() string {
	return "aws_apigatewayv2_api_mapping"
}

// LocalName returns the local name for [Resource].
func (aaam *Resource) LocalName() string {
	return aaam.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aaam *Resource) Configuration() interface{} {
	return aaam.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aaam *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aaam)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aaam *Resource) Dependencies() terra.Dependencies {
	return aaam.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aaam *Resource) LifecycleManagement() *terra.Lifecycle {
	return aaam.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aaam *Resource) Attributes() awsApigatewayv2ApiMappingAttributes {
	return awsApigatewayv2ApiMappingAttributes{ref: terra.ReferenceResource(aaam)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aaam *Resource) ImportState(state io.Reader) error {
	aaam.state = &awsApigatewayv2ApiMappingState{}
	if err := json.NewDecoder(state).Decode(aaam.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aaam.Type(), aaam.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aaam *Resource) State() (*awsApigatewayv2ApiMappingState, bool) {
	return aaam.state, aaam.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aaam *Resource) StateMust() *awsApigatewayv2ApiMappingState {
	if aaam.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aaam.Type(), aaam.LocalName()))
	}
	return aaam.state
}

// Args contains the configurations for aws_apigatewayv2_api_mapping.
type Args struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// ApiMappingKey: string, optional
	ApiMappingKey terra.StringValue `hcl:"api_mapping_key,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Stage: string, required
	Stage terra.StringValue `hcl:"stage,attr" validate:"required"`
}

type awsApigatewayv2ApiMappingAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_api_mapping.
func (aaam awsApigatewayv2ApiMappingAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(aaam.ref.Append("api_id"))
}

// ApiMappingKey returns a reference to field api_mapping_key of aws_apigatewayv2_api_mapping.
func (aaam awsApigatewayv2ApiMappingAttributes) ApiMappingKey() terra.StringValue {
	return terra.ReferenceAsString(aaam.ref.Append("api_mapping_key"))
}

// DomainName returns a reference to field domain_name of aws_apigatewayv2_api_mapping.
func (aaam awsApigatewayv2ApiMappingAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(aaam.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_apigatewayv2_api_mapping.
func (aaam awsApigatewayv2ApiMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aaam.ref.Append("id"))
}

// Stage returns a reference to field stage of aws_apigatewayv2_api_mapping.
func (aaam awsApigatewayv2ApiMappingAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(aaam.ref.Append("stage"))
}

type awsApigatewayv2ApiMappingState struct {
	ApiId         string `json:"api_id"`
	ApiMappingKey string `json:"api_mapping_key"`
	DomainName    string `json:"domain_name"`
	Id            string `json:"id"`
	Stage         string `json:"stage"`
}
