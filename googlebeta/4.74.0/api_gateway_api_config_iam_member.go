// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewayapiconfigiammember "github.com/golingon/terraproviders/googlebeta/4.74.0/apigatewayapiconfigiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayApiConfigIamMember creates a new instance of [ApiGatewayApiConfigIamMember].
func NewApiGatewayApiConfigIamMember(name string, args ApiGatewayApiConfigIamMemberArgs) *ApiGatewayApiConfigIamMember {
	return &ApiGatewayApiConfigIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayApiConfigIamMember)(nil)

// ApiGatewayApiConfigIamMember represents the Terraform resource google_api_gateway_api_config_iam_member.
type ApiGatewayApiConfigIamMember struct {
	Name      string
	Args      ApiGatewayApiConfigIamMemberArgs
	state     *apiGatewayApiConfigIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) Type() string {
	return "google_api_gateway_api_config_iam_member"
}

// LocalName returns the local name for [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) LocalName() string {
	return agacim.Name
}

// Configuration returns the configuration (args) for [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) Configuration() interface{} {
	return agacim.Args
}

// DependOn is used for other resources to depend on [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(agacim)
}

// Dependencies returns the list of resources [ApiGatewayApiConfigIamMember] depends_on.
func (agacim *ApiGatewayApiConfigIamMember) Dependencies() terra.Dependencies {
	return agacim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) LifecycleManagement() *terra.Lifecycle {
	return agacim.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayApiConfigIamMember].
func (agacim *ApiGatewayApiConfigIamMember) Attributes() apiGatewayApiConfigIamMemberAttributes {
	return apiGatewayApiConfigIamMemberAttributes{ref: terra.ReferenceResource(agacim)}
}

// ImportState imports the given attribute values into [ApiGatewayApiConfigIamMember]'s state.
func (agacim *ApiGatewayApiConfigIamMember) ImportState(av io.Reader) error {
	agacim.state = &apiGatewayApiConfigIamMemberState{}
	if err := json.NewDecoder(av).Decode(agacim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agacim.Type(), agacim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayApiConfigIamMember] has state.
func (agacim *ApiGatewayApiConfigIamMember) State() (*apiGatewayApiConfigIamMemberState, bool) {
	return agacim.state, agacim.state != nil
}

// StateMust returns the state for [ApiGatewayApiConfigIamMember]. Panics if the state is nil.
func (agacim *ApiGatewayApiConfigIamMember) StateMust() *apiGatewayApiConfigIamMemberState {
	if agacim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agacim.Type(), agacim.LocalName()))
	}
	return agacim.state
}

// ApiGatewayApiConfigIamMemberArgs contains the configurations for google_api_gateway_api_config_iam_member.
type ApiGatewayApiConfigIamMemberArgs struct {
	// Api: string, required
	Api terra.StringValue `hcl:"api,attr" validate:"required"`
	// ApiConfig: string, required
	ApiConfig terra.StringValue `hcl:"api_config,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigatewayapiconfigiammember.Condition `hcl:"condition,block"`
}
type apiGatewayApiConfigIamMemberAttributes struct {
	ref terra.Reference
}

// Api returns a reference to field api of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Api() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("api"))
}

// ApiConfig returns a reference to field api_config of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) ApiConfig() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("api_config"))
}

// Etag returns a reference to field etag of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("etag"))
}

// Id returns a reference to field id of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("id"))
}

// Member returns a reference to field member of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("member"))
}

// Project returns a reference to field project of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("project"))
}

// Role returns a reference to field role of google_api_gateway_api_config_iam_member.
func (agacim apiGatewayApiConfigIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(agacim.ref.Append("role"))
}

func (agacim apiGatewayApiConfigIamMemberAttributes) Condition() terra.ListValue[apigatewayapiconfigiammember.ConditionAttributes] {
	return terra.ReferenceAsList[apigatewayapiconfigiammember.ConditionAttributes](agacim.ref.Append("condition"))
}

type apiGatewayApiConfigIamMemberState struct {
	Api       string                                        `json:"api"`
	ApiConfig string                                        `json:"api_config"`
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Member    string                                        `json:"member"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Condition []apigatewayapiconfigiammember.ConditionState `json:"condition"`
}
