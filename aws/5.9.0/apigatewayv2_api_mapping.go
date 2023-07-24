// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApigatewayv2ApiMapping creates a new instance of [Apigatewayv2ApiMapping].
func NewApigatewayv2ApiMapping(name string, args Apigatewayv2ApiMappingArgs) *Apigatewayv2ApiMapping {
	return &Apigatewayv2ApiMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Apigatewayv2ApiMapping)(nil)

// Apigatewayv2ApiMapping represents the Terraform resource aws_apigatewayv2_api_mapping.
type Apigatewayv2ApiMapping struct {
	Name      string
	Args      Apigatewayv2ApiMappingArgs
	state     *apigatewayv2ApiMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) Type() string {
	return "aws_apigatewayv2_api_mapping"
}

// LocalName returns the local name for [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) LocalName() string {
	return aam.Name
}

// Configuration returns the configuration (args) for [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) Configuration() interface{} {
	return aam.Args
}

// DependOn is used for other resources to depend on [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(aam)
}

// Dependencies returns the list of resources [Apigatewayv2ApiMapping] depends_on.
func (aam *Apigatewayv2ApiMapping) Dependencies() terra.Dependencies {
	return aam.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) LifecycleManagement() *terra.Lifecycle {
	return aam.Lifecycle
}

// Attributes returns the attributes for [Apigatewayv2ApiMapping].
func (aam *Apigatewayv2ApiMapping) Attributes() apigatewayv2ApiMappingAttributes {
	return apigatewayv2ApiMappingAttributes{ref: terra.ReferenceResource(aam)}
}

// ImportState imports the given attribute values into [Apigatewayv2ApiMapping]'s state.
func (aam *Apigatewayv2ApiMapping) ImportState(av io.Reader) error {
	aam.state = &apigatewayv2ApiMappingState{}
	if err := json.NewDecoder(av).Decode(aam.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aam.Type(), aam.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Apigatewayv2ApiMapping] has state.
func (aam *Apigatewayv2ApiMapping) State() (*apigatewayv2ApiMappingState, bool) {
	return aam.state, aam.state != nil
}

// StateMust returns the state for [Apigatewayv2ApiMapping]. Panics if the state is nil.
func (aam *Apigatewayv2ApiMapping) StateMust() *apigatewayv2ApiMappingState {
	if aam.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aam.Type(), aam.LocalName()))
	}
	return aam.state
}

// Apigatewayv2ApiMappingArgs contains the configurations for aws_apigatewayv2_api_mapping.
type Apigatewayv2ApiMappingArgs struct {
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
type apigatewayv2ApiMappingAttributes struct {
	ref terra.Reference
}

// ApiId returns a reference to field api_id of aws_apigatewayv2_api_mapping.
func (aam apigatewayv2ApiMappingAttributes) ApiId() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("api_id"))
}

// ApiMappingKey returns a reference to field api_mapping_key of aws_apigatewayv2_api_mapping.
func (aam apigatewayv2ApiMappingAttributes) ApiMappingKey() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("api_mapping_key"))
}

// DomainName returns a reference to field domain_name of aws_apigatewayv2_api_mapping.
func (aam apigatewayv2ApiMappingAttributes) DomainName() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("domain_name"))
}

// Id returns a reference to field id of aws_apigatewayv2_api_mapping.
func (aam apigatewayv2ApiMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("id"))
}

// Stage returns a reference to field stage of aws_apigatewayv2_api_mapping.
func (aam apigatewayv2ApiMappingAttributes) Stage() terra.StringValue {
	return terra.ReferenceAsString(aam.ref.Append("stage"))
}

type apigatewayv2ApiMappingState struct {
	ApiId         string `json:"api_id"`
	ApiMappingKey string `json:"api_mapping_key"`
	DomainName    string `json:"domain_name"`
	Id            string `json:"id"`
	Stage         string `json:"stage"`
}
