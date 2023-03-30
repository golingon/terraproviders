// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitoriskconfiguration "github.com/golingon/terraproviders/aws/4.60.0/cognitoriskconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewCognitoRiskConfiguration(name string, args CognitoRiskConfigurationArgs) *CognitoRiskConfiguration {
	return &CognitoRiskConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoRiskConfiguration)(nil)

type CognitoRiskConfiguration struct {
	Name  string
	Args  CognitoRiskConfigurationArgs
	state *cognitoRiskConfigurationState
}

func (crc *CognitoRiskConfiguration) Type() string {
	return "aws_cognito_risk_configuration"
}

func (crc *CognitoRiskConfiguration) LocalName() string {
	return crc.Name
}

func (crc *CognitoRiskConfiguration) Configuration() interface{} {
	return crc.Args
}

func (crc *CognitoRiskConfiguration) Attributes() cognitoRiskConfigurationAttributes {
	return cognitoRiskConfigurationAttributes{ref: terra.ReferenceResource(crc)}
}

func (crc *CognitoRiskConfiguration) ImportState(av io.Reader) error {
	crc.state = &cognitoRiskConfigurationState{}
	if err := json.NewDecoder(av).Decode(crc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crc.Type(), crc.LocalName(), err)
	}
	return nil
}

func (crc *CognitoRiskConfiguration) State() (*cognitoRiskConfigurationState, bool) {
	return crc.state, crc.state != nil
}

func (crc *CognitoRiskConfiguration) StateMust() *cognitoRiskConfigurationState {
	if crc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crc.Type(), crc.LocalName()))
	}
	return crc.state
}

func (crc *CognitoRiskConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(crc)
}

type CognitoRiskConfigurationArgs struct {
	// ClientId: string, optional
	ClientId terra.StringValue `hcl:"client_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// UserPoolId: string, required
	UserPoolId terra.StringValue `hcl:"user_pool_id,attr" validate:"required"`
	// AccountTakeoverRiskConfiguration: optional
	AccountTakeoverRiskConfiguration *cognitoriskconfiguration.AccountTakeoverRiskConfiguration `hcl:"account_takeover_risk_configuration,block"`
	// CompromisedCredentialsRiskConfiguration: optional
	CompromisedCredentialsRiskConfiguration *cognitoriskconfiguration.CompromisedCredentialsRiskConfiguration `hcl:"compromised_credentials_risk_configuration,block"`
	// RiskExceptionConfiguration: optional
	RiskExceptionConfiguration *cognitoriskconfiguration.RiskExceptionConfiguration `hcl:"risk_exception_configuration,block"`
	// DependsOn contains resources that CognitoRiskConfiguration depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type cognitoRiskConfigurationAttributes struct {
	ref terra.Reference
}

func (crc cognitoRiskConfigurationAttributes) ClientId() terra.StringValue {
	return terra.ReferenceString(crc.ref.Append("client_id"))
}

func (crc cognitoRiskConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceString(crc.ref.Append("id"))
}

func (crc cognitoRiskConfigurationAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceString(crc.ref.Append("user_pool_id"))
}

func (crc cognitoRiskConfigurationAttributes) AccountTakeoverRiskConfiguration() terra.ListValue[cognitoriskconfiguration.AccountTakeoverRiskConfigurationAttributes] {
	return terra.ReferenceList[cognitoriskconfiguration.AccountTakeoverRiskConfigurationAttributes](crc.ref.Append("account_takeover_risk_configuration"))
}

func (crc cognitoRiskConfigurationAttributes) CompromisedCredentialsRiskConfiguration() terra.ListValue[cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationAttributes] {
	return terra.ReferenceList[cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationAttributes](crc.ref.Append("compromised_credentials_risk_configuration"))
}

func (crc cognitoRiskConfigurationAttributes) RiskExceptionConfiguration() terra.ListValue[cognitoriskconfiguration.RiskExceptionConfigurationAttributes] {
	return terra.ReferenceList[cognitoriskconfiguration.RiskExceptionConfigurationAttributes](crc.ref.Append("risk_exception_configuration"))
}

type cognitoRiskConfigurationState struct {
	ClientId                                string                                                                  `json:"client_id"`
	Id                                      string                                                                  `json:"id"`
	UserPoolId                              string                                                                  `json:"user_pool_id"`
	AccountTakeoverRiskConfiguration        []cognitoriskconfiguration.AccountTakeoverRiskConfigurationState        `json:"account_takeover_risk_configuration"`
	CompromisedCredentialsRiskConfiguration []cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationState `json:"compromised_credentials_risk_configuration"`
	RiskExceptionConfiguration              []cognitoriskconfiguration.RiskExceptionConfigurationState              `json:"risk_exception_configuration"`
}
