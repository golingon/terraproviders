// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewayapiiammember "github.com/golingon/terraproviders/googlebeta/4.77.0/apigatewayapiiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApiIamMember creates a new instance of [ApiGatewayApiIamMember].
func NewApiGatewayApiIamMember(name string, args ApiGatewayApiIamMemberArgs) *ApiGatewayApiIamMember {
	return &ApiGatewayApiIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiIamMember)(nil)

// ApiGatewayApiIamMember represents the Terraform resource google_api_gateway_api_iam_member.
type ApiGatewayApiIamMember struct {
	Name      string
	Args      ApiGatewayApiIamMemberArgs
	state     *apiGatewayApiIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) Type() string {
	return "google_api_gateway_api_iam_member"
}

// LocalName returns the local name for [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) LocalName() string {
	return agaim.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) Configuration() interface{} {
	return agaim.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(agaim)
}

// Dependencies returns the list of resources [ApiGatewayApiIamMember] depends_on.
func (agaim *ApiGatewayApiIamMember) Dependencies() terra.Dependencies {
	return agaim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) LifecycleManagement() *terra.Lifecycle {
	return agaim.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiIamMember].
func (agaim *ApiGatewayApiIamMember) Attributes() apiGatewayApiIamMemberAttributes {
	return apiGatewayApiIamMemberAttributes{ref: terra.ReferenceResource(agaim)}
}

// ImportState imports the given attribute values into [ApiGatewayApiIamMember]'s state.
func (agaim *ApiGatewayApiIamMember) ImportState(av io.Reader) error {
	agaim.state = &apiGatewayApiIamMemberState{}
	if err := json.NewDecoder(av).Decode(agaim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agaim.Type(), agaim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiIamMember] has state.
func (agaim *ApiGatewayApiIamMember) State() (*apiGatewayApiIamMemberState, bool) {
	return agaim.state, agaim.state != nil
}

// StateMust returns the state for [ApiGatewayApiIamMember]. Panics if the state is nil.
func (agaim *ApiGatewayApiIamMember) StateMust() *apiGatewayApiIamMemberState {
	if agaim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agaim.Type(), agaim.LocalName()))
	}
	return agaim.state
}

// ApiGatewayApiIamMemberArgs contains the configurations for google_api_gateway_api_iam_member.
type ApiGatewayApiIamMemberArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigatewayapiiammember.Condition `hcl:"condition,block"`
}
type apiGatewayApiIamMemberAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("api"))
}

// Etag returns a reference to field etag of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("id"))
}

// Member returns a reference to field member of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("member"))
}

// Project returns a reference to field project of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("project"))
}

// Role returns a reference to field role of google_api_gateway_api_iam_member.
func (agaim apiGatewayApiIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(agaim.ref.Append("role"))
}

func (agaim apiGatewayApiIamMemberAttributes) Condition() terra.ListValue[apigatewayapiiammember.ConditionAttributes] {
	return terra.ReferenceAsList[apigatewayapiiammember.ConditionAttributes](agaim.ref.Append("condition"))
}

type apiGatewayApiIamMemberState struct {
	Api       string                                  `json:"api"`
	Etag      string                                  `json:"etag"`
	Id        string                                  `json:"id"`
	Member    string                                  `json:"member"`
	Project   string                                  `json:"project"`
	Role      string                                  `json:"role"`
	Condition []apigatewayapiiammember.ConditionState `json:"condition"`
}
