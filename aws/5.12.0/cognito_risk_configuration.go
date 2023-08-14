// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cognitoriskconfiguration "github.com/golingon/terraproviders/aws/5.12.0/cognitoriskconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCognitoRiskConfiguration creates a new instance of [CognitoRiskConfiguration].
func NewCognitoRiskConfiguration(name string, args CognitoRiskConfigurationArgs) *CognitoRiskConfiguration {
	return &CognitoRiskConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CognitoRiskConfiguration)(nil)

// CognitoRiskConfiguration represents the Terraform resource aws_cognito_risk_configuration.
type CognitoRiskConfiguration struct {
	Name      string
	Args      CognitoRiskConfigurationArgs
	state     *cognitoRiskConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) Type() string {
	return "aws_cognito_risk_configuration"
}

// LocalName returns the local name for [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) LocalName() string {
	return crc.Name
}

// Configuration returns the configuration (args) for [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) Configuration() interface{} {
	return crc.Args
}

// DependOn is used for other resources to depend on [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(crc)
}

// Dependencies returns the list of resources [CognitoRiskConfiguration] depends_on.
func (crc *CognitoRiskConfiguration) Dependencies() terra.Dependencies {
	return crc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) LifecycleManagement() *terra.Lifecycle {
	return crc.Lifecycle
}

// Attributes returns the attributes for [CognitoRiskConfiguration].
func (crc *CognitoRiskConfiguration) Attributes() cognitoRiskConfigurationAttributes {
	return cognitoRiskConfigurationAttributes{ref: terra.ReferenceResource(crc)}
}

// ImportState imports the given attribute values into [CognitoRiskConfiguration]'s state.
func (crc *CognitoRiskConfiguration) ImportState(av io.Reader) error {
	crc.state = &cognitoRiskConfigurationState{}
	if err := json.NewDecoder(av).Decode(crc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crc.Type(), crc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CognitoRiskConfiguration] has state.
func (crc *CognitoRiskConfiguration) State() (*cognitoRiskConfigurationState, bool) {
	return crc.state, crc.state != nil
}

// StateMust returns the state for [CognitoRiskConfiguration]. Panics if the state is nil.
func (crc *CognitoRiskConfiguration) StateMust() *cognitoRiskConfigurationState {
	if crc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crc.Type(), crc.LocalName()))
	}
	return crc.state
}

// CognitoRiskConfigurationArgs contains the configurations for aws_cognito_risk_configuration.
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
}
type cognitoRiskConfigurationAttributes struct {
	ref terra.Reference
}

// ClientId returns a reference to field client_id of aws_cognito_risk_configuration.
func (crc cognitoRiskConfigurationAttributes) ClientId() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("client_id"))
}

// Id returns a reference to field id of aws_cognito_risk_configuration.
func (crc cognitoRiskConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("id"))
}

// UserPoolId returns a reference to field user_pool_id of aws_cognito_risk_configuration.
func (crc cognitoRiskConfigurationAttributes) UserPoolId() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("user_pool_id"))
}

func (crc cognitoRiskConfigurationAttributes) AccountTakeoverRiskConfiguration() terra.ListValue[cognitoriskconfiguration.AccountTakeoverRiskConfigurationAttributes] {
	return terra.ReferenceAsList[cognitoriskconfiguration.AccountTakeoverRiskConfigurationAttributes](crc.ref.Append("account_takeover_risk_configuration"))
}

func (crc cognitoRiskConfigurationAttributes) CompromisedCredentialsRiskConfiguration() terra.ListValue[cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationAttributes] {
	return terra.ReferenceAsList[cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationAttributes](crc.ref.Append("compromised_credentials_risk_configuration"))
}

func (crc cognitoRiskConfigurationAttributes) RiskExceptionConfiguration() terra.ListValue[cognitoriskconfiguration.RiskExceptionConfigurationAttributes] {
	return terra.ReferenceAsList[cognitoriskconfiguration.RiskExceptionConfigurationAttributes](crc.ref.Append("risk_exception_configuration"))
}

type cognitoRiskConfigurationState struct {
	ClientId                                string                                                                  `json:"client_id"`
	Id                                      string                                                                  `json:"id"`
	UserPoolId                              string                                                                  `json:"user_pool_id"`
	AccountTakeoverRiskConfiguration        []cognitoriskconfiguration.AccountTakeoverRiskConfigurationState        `json:"account_takeover_risk_configuration"`
	CompromisedCredentialsRiskConfiguration []cognitoriskconfiguration.CompromisedCredentialsRiskConfigurationState `json:"compromised_credentials_risk_configuration"`
	RiskExceptionConfiguration              []cognitoriskconfiguration.RiskExceptionConfigurationState              `json:"risk_exception_configuration"`
}
