// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayBasePathMapping creates a new instance of [ApiGatewayBasePathMapping].
func NewApiGatewayBasePathMapping(name string, args ApiGatewayBasePathMappingArgs) *ApiGatewayBasePathMapping {
	return &ApiGatewayBasePathMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayBasePathMapping)(nil)

// ApiGatewayBasePathMapping represents the Terraform resource aws_api_gateway_base_path_mapping.
type ApiGatewayBasePathMapping struct {
	Name      string
	Args      ApiGatewayBasePathMappingArgs
	state     *apiGatewayBasePathMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) Type() string {
	return "aws_api_gateway_base_path_mapping"
}

// LocalName returns the local name for [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) LocalName() string {
	return agbpm.Name
}

// Configuration returns the configuration (args) for [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) Configuration() interface{} {
	return agbpm.Args
}

// DependOn is used for other resources to depend on [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(agbpm)
}

// Dependencies returns the list of resources [ApiGatewayBasePathMapping] depends_on.
func (agbpm *ApiGatewayBasePathMapping) Dependencies() terra.Dependencies {
	return agbpm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) LifecycleManagement() *terra.Lifecycle {
	return agbpm.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayBasePathMapping].
func (agbpm *ApiGatewayBasePathMapping) Attributes() apiGatewayBasePathMappingAttributes {
	return apiGatewayBasePathMappingAttributes{ref: terra.ReferenceResource(agbpm)}
}

// ImportState imports the given attribute values into [ApiGatewayBasePathMapping]'s state.
func (agbpm *ApiGatewayBasePathMapping) ImportState(av io.Reader) error {
	agbpm.state = &apiGatewayBasePathMappingState{}
	if err := json.NewDecoder(av).Decode(agbpm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agbpm.Type(), agbpm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayBasePathMapping] has state.
func (agbpm *ApiGatewayBasePathMapping) State() (*apiGatewayBasePathMappingState, bool) {
	return agbpm.state, agbpm.state != nil
}

// StateMust returns the state for [ApiGatewayBasePathMapping]. Panics if the state is nil.
func (agbpm *ApiGatewayBasePathMapping) StateMust() *apiGatewayBasePathMappingState {
	if agbpm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agbpm.Type(), agbpm.LocalName()))
	}
	return agbpm.state
}

// ApiGatewayBasePathMappingArgs contains the configurations for aws_api_gateway_base_path_mapping.
type ApiGatewayBasePathMappingArgs struct {
	// ApiId: string, required
	ApiId terra.StringValue `hcl:"api_id,attr" validate:"required"`
	// BasePath: string, optional
	BasePath terra.StringValue `hcl:"base_path,attr"`
	// DomainName: string, required
	DomainName terra.StringValue `hcl:"domain_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StageName: string, optional
	StageName terra.StringValue `hcl:"stage_name,attr"`
}
type apiGatewayBasePathMappingAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_api_gateway_base_path_mapping.
func (agbpm apiGatewayBasePathMappingAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(agbpm.ref.Append("api_id"))
}

// BasePath returns a reference to field base_path of aws_api_gateway_base_path_mapping.
func (agbpm apiGatewayBasePathMappingAttributes) BasePath() terra.StringValue {
	return terra.ReferenceAsString(agbpm.ref.Append("base_path"))
}

// DomainName returns a reference to field domain_name of aws_api_gateway_base_path_mapping.
func (agbpm apiGatewayBasePathMappingAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(agbpm.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_api_gateway_base_path_mapping.
func (agbpm apiGatewayBasePathMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agbpm.ref.Append("id"))
}

// StageName returns a reference to field stage_name of aws_api_gateway_base_path_mapping.
func (agbpm apiGatewayBasePathMappingAttributes) StageName() terra.StringValue {
	return terra.ReferenceAsString(agbpm.ref.Append("stage_name"))
}

type apiGatewayBasePathMappingState struct {
	ApiId      string `json:"api_id"`
	BasePath   string `json:"base_path"`
	DomainName string `json:"domain_name"`
	Id         string `json:"id"`
	StageName  string `json:"stage_name"`
}