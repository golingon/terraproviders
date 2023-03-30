// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewayaccount "github.com/golingon/terraproviders/aws/4.60.0/apigatewayaccount"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewApiGatewayAccount(name string, args ApiGatewayAccountArgs) *ApiGatewayAccount {
	return &ApiGatewayAccount{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayAccount)(nil)

type ApiGatewayAccount struct {
	Name  string
	Args  ApiGatewayAccountArgs
	state *apiGatewayAccountState
}

func (aga *ApiGatewayAccount) Type() string {
	return "aws_api_gateway_account"
}

func (aga *ApiGatewayAccount) LocalName() string {
	return aga.Name
}

func (aga *ApiGatewayAccount) Configuration() interface{} {
	return aga.Args
}

func (aga *ApiGatewayAccount) Attributes() apiGatewayAccountAttributes {
	return apiGatewayAccountAttributes{ref: terra.ReferenceResource(aga)}
}

func (aga *ApiGatewayAccount) ImportState(av io.Reader) error {
	aga.state = &apiGatewayAccountState{}
	if err := json.NewDecoder(av).Decode(aga.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aga.Type(), aga.LocalName(), err)
	}
	return nil
}

func (aga *ApiGatewayAccount) State() (*apiGatewayAccountState, bool) {
	return aga.state, aga.state != nil
}

func (aga *ApiGatewayAccount) StateMust() *apiGatewayAccountState {
	if aga.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aga.Type(), aga.LocalName()))
	}
	return aga.state
}

func (aga *ApiGatewayAccount) DependOn() terra.Reference {
	return terra.ReferenceResource(aga)
}

type ApiGatewayAccountArgs struct {
	// CloudwatchRoleArn: string, optional
	CloudwatchRoleArn terra.StringValue `hcl:"cloudwatch_role_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ThrottleSettings: min=0
	ThrottleSettings []apigatewayaccount.ThrottleSettings `hcl:"throttle_settings,block" validate:"min=0"`
	// DependsOn contains resources that ApiGatewayAccount depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type apiGatewayAccountAttributes struct {
	ref terra.Reference
}

func (aga apiGatewayAccountAttributes) CloudwatchRoleArn() terra.StringValue {
	return terra.ReferenceString(aga.ref.Append("cloudwatch_role_arn"))
}

func (aga apiGatewayAccountAttributes) Id() terra.StringValue {
	return terra.ReferenceString(aga.ref.Append("id"))
}

func (aga apiGatewayAccountAttributes) ThrottleSettings() terra.ListValue[apigatewayaccount.ThrottleSettingsAttributes] {
	return terra.ReferenceList[apigatewayaccount.ThrottleSettingsAttributes](aga.ref.Append("throttle_settings"))
}

type apiGatewayAccountState struct {
	CloudwatchRoleArn string                                    `json:"cloudwatch_role_arn"`
	Id                string                                    `json:"id"`
	ThrottleSettings  []apigatewayaccount.ThrottleSettingsState `json:"throttle_settings"`
}
