// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayDeployment creates a new instance of [ApiGatewayDeployment].
func NewApiGatewayDeployment(name string, args ApiGatewayDeploymentArgs) *ApiGatewayDeployment {
	return &ApiGatewayDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayDeployment)(nil)

// ApiGatewayDeployment represents the Terraform resource aws_api_gateway_deployment.
type ApiGatewayDeployment struct {
	Name      string
	Args      ApiGatewayDeploymentArgs
	state     *apiGatewayDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) Type() string {
	return "aws_api_gateway_deployment"
}

// LocalName returns the local name for [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) LocalName() string {
	return agd.Name
}

// Configuration returns the configuration (args) for [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) Configuration() interface{} {
	return agd.Args
}

// DependOn is used for other resources to depend on [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(agd)
}

// Dependencies returns the list of resources [ApiGatewayDeployment] depends_on.
func (agd *ApiGatewayDeployment) Dependencies() terra.Dependencies {
	return agd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) LifecycleManagement() *terra.Lifecycle {
	return agd.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayDeployment].
func (agd *ApiGatewayDeployment) Attributes() apiGatewayDeploymentAttributes {
	return apiGatewayDeploymentAttributes{ref: terra.ReferenceResource(agd)}
}

// ImportState imports the given attribute values into [ApiGatewayDeployment]'s state.
func (agd *ApiGatewayDeployment) ImportState(av io.Reader) error {
	agd.state = &apiGatewayDeploymentState{}
	if err := json.NewDecoder(av).Decode(agd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agd.Type(), agd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayDeployment] has state.
func (agd *ApiGatewayDeployment) State() (*apiGatewayDeploymentState, bool) {
	return agd.state, agd.state != nil
}

// StateMust returns the state for [ApiGatewayDeployment]. Panics if the state is nil.
func (agd *ApiGatewayDeployment) StateMust() *apiGatewayDeploymentState {
	if agd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agd.Type(), agd.LocalName()))
	}
	return agd.state
}

// ApiGatewayDeploymentArgs contains the configurations for aws_api_gateway_deployment.
type ApiGatewayDeploymentArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// StageDescription: string, optional
	StageDescription terra.StringValue `hcl:"stage_description,attr"`
	// StageName: string, optional
	StageName terra.StringValue `hcl:"stage_name,attr"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
	// Variables: map of string, optional
	Variables terra.MapValue[terra.StringValue] `hcl:"variables,attr"`
}
type apiGatewayDeploymentAttributes struct {
	ref terra.Reference
}

// CreatedDate returns a reference to field created_date of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) CreatedDate() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("created_date"))
}

// Description returns a reference to field description of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("description"))
}

// ExecutionArn returns a reference to field execution_arn of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) ExecutionArn() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("execution_arn"))
}

// Id returns a reference to field id of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("id"))
}

// InvokeUrl returns a reference to field invoke_url of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) InvokeUrl() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("invoke_url"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("rest_api_id"))
}

// StageDescription returns a reference to field stage_description of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) StageDescription() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("stage_description"))
}

// StageName returns a reference to field stage_name of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) StageName() terra.StringValue {
	return terra.ReferenceAsString(agd.ref.Append("stage_name"))
}

// Triggers returns a reference to field triggers of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agd.ref.Append("triggers"))
}

// Variables returns a reference to field variables of aws_api_gateway_deployment.
func (agd apiGatewayDeploymentAttributes) Variables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](agd.ref.Append("variables"))
}

type apiGatewayDeploymentState struct {
	CreatedDate      string            `json:"created_date"`
	Description      string            `json:"description"`
	ExecutionArn     string            `json:"execution_arn"`
	Id               string            `json:"id"`
	InvokeUrl        string            `json:"invoke_url"`
	RestApiId        string            `json:"rest_api_id"`
	StageDescription string            `json:"stage_description"`
	StageName        string            `json:"stage_name"`
	Triggers         map[string]string `json:"triggers"`
	Variables        map[string]string `json:"variables"`
}
