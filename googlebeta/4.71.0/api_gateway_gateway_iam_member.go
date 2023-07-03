// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	apigatewaygatewayiammember "github.com/golingon/terraproviders/googlebeta/4.71.0/apigatewaygatewayiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayGatewayIamMember creates a new instance of [ApiGatewayGatewayIamMember].
func NewApiGatewayGatewayIamMember(name string, args ApiGatewayGatewayIamMemberArgs) *ApiGatewayGatewayIamMember {
	return &ApiGatewayGatewayIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayGatewayIamMember)(nil)

// ApiGatewayGatewayIamMember represents the Terraform resource google_api_gateway_gateway_iam_member.
type ApiGatewayGatewayIamMember struct {
	Name      string
	Args      ApiGatewayGatewayIamMemberArgs
	state     *apiGatewayGatewayIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) Type() string {
	return "google_api_gateway_gateway_iam_member"
}

// LocalName returns the local name for [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) LocalName() string {
	return aggim.Name
}

// Configuration returns the configuration (args) for [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) Configuration() interface{} {
	return aggim.Args
}

// DependOn is used for other resources to depend on [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(aggim)
}

// Dependencies returns the list of resources [ApiGatewayGatewayIamMember] depends_on.
func (aggim *ApiGatewayGatewayIamMember) Dependencies() terra.Dependencies {
	return aggim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) LifecycleManagement() *terra.Lifecycle {
	return aggim.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayGatewayIamMember].
func (aggim *ApiGatewayGatewayIamMember) Attributes() apiGatewayGatewayIamMemberAttributes {
	return apiGatewayGatewayIamMemberAttributes{ref: terra.ReferenceResource(aggim)}
}

// ImportState imports the given attribute values into [ApiGatewayGatewayIamMember]'s state.
func (aggim *ApiGatewayGatewayIamMember) ImportState(av io.Reader) error {
	aggim.state = &apiGatewayGatewayIamMemberState{}
	if err := json.NewDecoder(av).Decode(aggim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aggim.Type(), aggim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayGatewayIamMember] has state.
func (aggim *ApiGatewayGatewayIamMember) State() (*apiGatewayGatewayIamMemberState, bool) {
	return aggim.state, aggim.state != nil
}

// StateMust returns the state for [ApiGatewayGatewayIamMember]. Panics if the state is nil.
func (aggim *ApiGatewayGatewayIamMember) StateMust() *apiGatewayGatewayIamMemberState {
	if aggim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aggim.Type(), aggim.LocalName()))
	}
	return aggim.state
}

// ApiGatewayGatewayIamMemberArgs contains the configurations for google_api_gateway_gateway_iam_member.
type ApiGatewayGatewayIamMemberArgs struct {
	// Gateway: string, required
	Gateway terra.StringValue `hcl:"gateway,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *apigatewaygatewayiammember.Condition `hcl:"condition,block"`
}
type apiGatewayGatewayIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("etag"))
}

// Gateway returns a reference to field gateway of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Gateway() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("gateway"))
}

// Id returns a reference to field id of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("id"))
}

// Member returns a reference to field member of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("member"))
}

// Project returns a reference to field project of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("project"))
}

// Region returns a reference to field region of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("region"))
}

// Role returns a reference to field role of google_api_gateway_gateway_iam_member.
func (aggim apiGatewayGatewayIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(aggim.ref.Append("role"))
}

func (aggim apiGatewayGatewayIamMemberAttributes) Condition() terra.ListValue[apigatewaygatewayiammember.ConditionAttributes] {
	return terra.ReferenceAsList[apigatewaygatewayiammember.ConditionAttributes](aggim.ref.Append("condition"))
}

type apiGatewayGatewayIamMemberState struct {
	Etag      string                                      `json:"etag"`
	Gateway   string                                      `json:"gateway"`
	Id        string                                      `json:"id"`
	Member    string                                      `json:"member"`
	Project   string                                      `json:"project"`
	Region    string                                      `json:"region"`
	Role      string                                      `json:"role"`
	Condition []apigatewaygatewayiammember.ConditionState `json:"condition"`
}
