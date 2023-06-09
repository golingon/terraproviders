// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApiIamPolicy creates a new instance of [ApiGatewayApiIamPolicy].
func NewApiGatewayApiIamPolicy(name string, args ApiGatewayApiIamPolicyArgs) *ApiGatewayApiIamPolicy {
	return &ApiGatewayApiIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiIamPolicy)(nil)

// ApiGatewayApiIamPolicy represents the Terraform resource google_api_gateway_api_iam_policy.
type ApiGatewayApiIamPolicy struct {
	Name      string
	Args      ApiGatewayApiIamPolicyArgs
	state     *apiGatewayApiIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) Type() string {
	return "google_api_gateway_api_iam_policy"
}

// LocalName returns the local name for [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) LocalName() string {
	return agaip.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) Configuration() interface{} {
	return agaip.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(agaip)
}

// Dependencies returns the list of resources [ApiGatewayApiIamPolicy] depends_on.
func (agaip *ApiGatewayApiIamPolicy) Dependencies() terra.Dependencies {
	return agaip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return agaip.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiIamPolicy].
func (agaip *ApiGatewayApiIamPolicy) Attributes() apiGatewayApiIamPolicyAttributes {
	return apiGatewayApiIamPolicyAttributes{ref: terra.ReferenceResource(agaip)}
}

// ImportState imports the given attribute values into [ApiGatewayApiIamPolicy]'s state.
func (agaip *ApiGatewayApiIamPolicy) ImportState(av io.Reader) error {
	agaip.state = &apiGatewayApiIamPolicyState{}
	if err := json.NewDecoder(av).Decode(agaip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agaip.Type(), agaip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiIamPolicy] has state.
func (agaip *ApiGatewayApiIamPolicy) State() (*apiGatewayApiIamPolicyState, bool) {
	return agaip.state, agaip.state != nil
}

// StateMust returns the state for [ApiGatewayApiIamPolicy]. Panics if the state is nil.
func (agaip *ApiGatewayApiIamPolicy) StateMust() *apiGatewayApiIamPolicyState {
	if agaip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agaip.Type(), agaip.LocalName()))
	}
	return agaip.state
}

// ApiGatewayApiIamPolicyArgs contains the configurations for google_api_gateway_api_iam_policy.
type ApiGatewayApiIamPolicyArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type apiGatewayApiIamPolicyAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_iam_policy.
func (agaip apiGatewayApiIamPolicyAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agaip.ref.Append("api"))
}

// Etag returns a reference to field etag of google_api_gateway_api_iam_policy.
func (agaip apiGatewayApiIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agaip.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_iam_policy.
func (agaip apiGatewayApiIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agaip.ref.Append("id"))
}

// PolicyData returns a reference to field policy_data of google_api_gateway_api_iam_policy.
func (agaip apiGatewayApiIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(agaip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_api_gateway_api_iam_policy.
func (agaip apiGatewayApiIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agaip.ref.Append("project"))
}

type apiGatewayApiIamPolicyState struct {
	Api        string `json:"api"`
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
