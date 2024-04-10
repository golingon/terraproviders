// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	apigatewayapiconfigiambinding "github.com/golingon/terraproviders/googlebeta/5.24.0/apigatewayapiconfigiambinding"
	"io"
)

// NewApiGatewayApiConfigIamBinding creates a new instance of [ApiGatewayApiConfigIamBinding].
func NewApiGatewayApiConfigIamBinding(name string, args ApiGatewayApiConfigIamBindingArgs) *ApiGatewayApiConfigIamBinding {
	return &ApiGatewayApiConfigIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiConfigIamBinding)(nil)

// ApiGatewayApiConfigIamBinding represents the Terraform resource google_api_gateway_api_config_iam_binding.
type ApiGatewayApiConfigIamBinding struct {
	Name      string
	Args      ApiGatewayApiConfigIamBindingArgs
	state     *apiGatewayApiConfigIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) Type() string {
	return "google_api_gateway_api_config_iam_binding"
}

// LocalName returns the local name for [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) LocalName() string {
	return agacib.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) Configuration() interface{} {
	return agacib.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(agacib)
}

// Dependencies returns the list of resources [ApiGatewayApiConfigIamBinding] depends_on.
func (agacib *ApiGatewayApiConfigIamBinding) Dependencies() terra.Dependencies {
	return agacib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) LifecycleManagement() *terra.Lifecycle {
	return agacib.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiConfigIamBinding].
func (agacib *ApiGatewayApiConfigIamBinding) Attributes() apiGatewayApiConfigIamBindingAttributes {
	return apiGatewayApiConfigIamBindingAttributes{ref: terra.ReferenceResource(agacib)}
}

// ImportState imports the given attribute values into [ApiGatewayApiConfigIamBinding]'s state.
func (agacib *ApiGatewayApiConfigIamBinding) ImportState(av io.Reader) error {
	agacib.state = &apiGatewayApiConfigIamBindingState{}
	if err := json.NewDecoder(av).Decode(agacib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agacib.Type(), agacib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiConfigIamBinding] has state.
func (agacib *ApiGatewayApiConfigIamBinding) State() (*apiGatewayApiConfigIamBindingState, bool) {
	return agacib.state, agacib.state != nil
}

// StateMust returns the state for [ApiGatewayApiConfigIamBinding]. Panics if the state is nil.
func (agacib *ApiGatewayApiConfigIamBinding) StateMust() *apiGatewayApiConfigIamBindingState {
	if agacib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agacib.Type(), agacib.LocalName()))
	}
	return agacib.state
}

// ApiGatewayApiConfigIamBindingArgs contains the configurations for google_api_gateway_api_config_iam_binding.
type ApiGatewayApiConfigIamBindingArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// ApiConfig: string, required
	ApiConfig terra.StringValue `hcl:"api_config,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigatewayapiconfigiambinding.Condition `hcl:"condition,block"`
}
type apiGatewayApiConfigIamBindingAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("api"))
}

// ApiConfig returns a reference to field api_config of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) ApiConfig() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("api_config"))
}

// Etag returns a reference to field etag of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("id"))
}

// Members returns a reference to field members of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](agacib.ref.Append("members"))
}

// Project returns a reference to field project of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("project"))
}

// Role returns a reference to field role of google_api_gateway_api_config_iam_binding.
func (agacib apiGatewayApiConfigIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(agacib.ref.Append("role"))
}

func (agacib apiGatewayApiConfigIamBindingAttributes) Condition() terra.ListValue[apigatewayapiconfigiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfigiambinding.ConditionAttributes](agacib.ref.Append("condition"))
}

type apiGatewayApiConfigIamBindingState struct {
	Api       string                                         `json:"api"`
	ApiConfig string                                         `json:"api_config"`
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Members   []string                                       `json:"members"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Condition []apigatewayapiconfigiambinding.ConditionState `json:"condition"`
}
