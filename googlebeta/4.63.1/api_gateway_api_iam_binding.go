// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewayapiiambinding "github.com/golingon/terraproviders/googlebeta/4.63.1/apigatewayapiiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApiIamBinding creates a new instance of [ApiGatewayApiIamBinding].
func NewApiGatewayApiIamBinding(name string, args ApiGatewayApiIamBindingArgs) *ApiGatewayApiIamBinding {
	return &ApiGatewayApiIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiIamBinding)(nil)

// ApiGatewayApiIamBinding represents the Terraform resource google_api_gateway_api_iam_binding.
type ApiGatewayApiIamBinding struct {
	Name      string
	Args      ApiGatewayApiIamBindingArgs
	state     *apiGatewayApiIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) Type() string {
	return "google_api_gateway_api_iam_binding"
}

// LocalName returns the local name for [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) LocalName() string {
	return agaib.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) Configuration() interface{} {
	return agaib.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(agaib)
}

// Dependencies returns the list of resources [ApiGatewayApiIamBinding] depends_on.
func (agaib *ApiGatewayApiIamBinding) Dependencies() terra.Dependencies {
	return agaib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) LifecycleManagement() *terra.Lifecycle {
	return agaib.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiIamBinding].
func (agaib *ApiGatewayApiIamBinding) Attributes() apiGatewayApiIamBindingAttributes {
	return apiGatewayApiIamBindingAttributes{ref: terra.ReferenceResource(agaib)}
}

// ImportState imports the given attribute values into [ApiGatewayApiIamBinding]'s state.
func (agaib *ApiGatewayApiIamBinding) ImportState(av io.Reader) error {
	agaib.state = &apiGatewayApiIamBindingState{}
	if err := json.NewDecoder(av).Decode(agaib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agaib.Type(), agaib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiIamBinding] has state.
func (agaib *ApiGatewayApiIamBinding) State() (*apiGatewayApiIamBindingState, bool) {
	return agaib.state, agaib.state != nil
}

// StateMust returns the state for [ApiGatewayApiIamBinding]. Panics if the state is nil.
func (agaib *ApiGatewayApiIamBinding) StateMust() *apiGatewayApiIamBindingState {
	if agaib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agaib.Type(), agaib.LocalName()))
	}
	return agaib.state
}

// ApiGatewayApiIamBindingArgs contains the configurations for google_api_gateway_api_iam_binding.
type ApiGatewayApiIamBindingArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigatewayapiiambinding.Condition `hcl:"condition,block"`
}
type apiGatewayApiIamBindingAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agaib.ref.Append("api"))
}

// Etag returns a reference to field etag of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agaib.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agaib.ref.Append("id"))
}

// Members returns a reference to field members of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](agaib.ref.Append("members"))
}

// Project returns a reference to field project of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agaib.ref.Append("project"))
}

// Role returns a reference to field role of google_api_gateway_api_iam_binding.
func (agaib apiGatewayApiIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(agaib.ref.Append("role"))
}

func (agaib apiGatewayApiIamBindingAttributes) Condition() terra.ListValue[apigatewayapiiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[apigatewayapiiambinding.ConditionAttributes](agaib.ref.Append("condition"))
}

type apiGatewayApiIamBindingState struct {
	Api       string                                   `json:"api"`
	Etag      string                                   `json:"etag"`
	Id        string                                   `json:"id"`
	Members   []string                                 `json:"members"`
	Project   string                                   `json:"project"`
	Role      string                                   `json:"role"`
	Condition []apigatewayapiiambinding.ConditionState `json:"condition"`
}
