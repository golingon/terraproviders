// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	apigatewaymethodsettings "github.com/golingon/terraproviders/aws/5.0.1/apigatewaymethodsettings"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApiGatewayMethodSettings creates a new instance of [ApiGatewayMethodSettings].
func NewApiGatewayMethodSettings(name string, args ApiGatewayMethodSettingsArgs) *ApiGatewayMethodSettings {
	return &ApiGatewayMethodSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApiGatewayMethodSettings)(nil)

// ApiGatewayMethodSettings represents the Terraform resource aws_api_gateway_method_settings.
type ApiGatewayMethodSettings struct {
	Name      string
	Args      ApiGatewayMethodSettingsArgs
	state     *apiGatewayMethodSettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) Type() string {
	return "aws_api_gateway_method_settings"
}

// LocalName returns the local name for [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) LocalName() string {
	return agms.Name
}

// Configuration returns the configuration (args) for [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) Configuration() interface{} {
	return agms.Args
}

// DependOn is used for other resources to depend on [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) DependOn() terra.Reference {
	return terra.ReferenceResource(agms)
}

// Dependencies returns the list of resources [ApiGatewayMethodSettings] depends_on.
func (agms *ApiGatewayMethodSettings) Dependencies() terra.Dependencies {
	return agms.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) LifecycleManagement() *terra.Lifecycle {
	return agms.Lifecycle
}

// Attributes returns the attributes for [ApiGatewayMethodSettings].
func (agms *ApiGatewayMethodSettings) Attributes() apiGatewayMethodSettingsAttributes {
	return apiGatewayMethodSettingsAttributes{ref: terra.ReferenceResource(agms)}
}

// ImportState imports the given attribute values into [ApiGatewayMethodSettings]'s state.
func (agms *ApiGatewayMethodSettings) ImportState(av io.Reader) error {
	agms.state = &apiGatewayMethodSettingsState{}
	if err := json.NewDecoder(av).Decode(agms.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", agms.Type(), agms.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApiGatewayMethodSettings] has state.
func (agms *ApiGatewayMethodSettings) State() (*apiGatewayMethodSettingsState, bool) {
	return agms.state, agms.state != nil
}

// StateMust returns the state for [ApiGatewayMethodSettings]. Panics if the state is nil.
func (agms *ApiGatewayMethodSettings) StateMust() *apiGatewayMethodSettingsState {
	if agms.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", agms.Type(), agms.LocalName()))
	}
	return agms.state
}

// ApiGatewayMethodSettingsArgs contains the configurations for aws_api_gateway_method_settings.
type ApiGatewayMethodSettingsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MethodPath: string, required
	MethodPath terra.StringValue `hcl:"method_path,attr" validate:"required"`
	// RestApiId: string, required
	RestApiId terra.StringValue `hcl:"rest_api_id,attr" validate:"required"`
	// StageName: string, required
	StageName terra.StringValue `hcl:"stage_name,attr" validate:"required"`
	// Settings: required
	Settings *apigatewaymethodsettings.Settings `hcl:"settings,block" validate:"required"`
}
type apiGatewayMethodSettingsAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_api_gateway_method_settings.
func (agms apiGatewayMethodSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(agms.ref.Append("id"))
}

// MethodPath returns a reference to field method_path of aws_api_gateway_method_settings.
func (agms apiGatewayMethodSettingsAttributes) MethodPath() terra.StringValue {
	return terra.ReferenceAsString(agms.ref.Append("method_path"))
}

// RestApiId returns a reference to field rest_api_id of aws_api_gateway_method_settings.
func (agms apiGatewayMethodSettingsAttributes) RestApiId() terra.StringValue {
	return terra.ReferenceAsString(agms.ref.Append("rest_api_id"))
}

// StageName returns a reference to field stage_name of aws_api_gateway_method_settings.
func (agms apiGatewayMethodSettingsAttributes) StageName() terra.StringValue {
	return terra.ReferenceAsString(agms.ref.Append("stage_name"))
}

func (agms apiGatewayMethodSettingsAttributes) Settings() terra.ListValue[apigatewaymethodsettings.SettingsAttributes] {
	return terra.ReferenceAsList[apigatewaymethodsettings.SettingsAttributes](agms.ref.Append("settings"))
}

type apiGatewayMethodSettingsState struct {
	Id         string                                   `json:"id"`
	MethodPath string                                   `json:"method_path"`
	RestApiId  string                                   `json:"rest_api_id"`
	StageName  string                                   `json:"stage_name"`
	Settings   []apigatewaymethodsettings.SettingsState `json:"settings"`
}