// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	appservicesourcecontroltoken "github.com/golingon/terraproviders/azurerm/3.66.0/appservicesourcecontroltoken"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppServiceSourceControlToken creates a new instance of [AppServiceSourceControlToken].
func NewAppServiceSourceControlToken(name string, args AppServiceSourceControlTokenArgs) *AppServiceSourceControlToken {
	return &AppServiceSourceControlToken{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppServiceSourceControlToken)(nil)

// AppServiceSourceControlToken represents the Terraform resource azurerm_app_service_source_control_token.
type AppServiceSourceControlToken struct {
	Name      string
	Args      AppServiceSourceControlTokenArgs
	state     *appServiceSourceControlTokenState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) Type() string {
	return "azurerm_app_service_source_control_token"
}

// LocalName returns the local name for [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) LocalName() string {
	return assct.Name
}

// Configuration returns the configuration (args) for [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) Configuration() interface{} {
	return assct.Args
}

// DependOn is used for other resources to depend on [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) DependOn() terra.Reference {
	return terra.ReferenceResource(assct)
}

// Dependencies returns the list of resources [AppServiceSourceControlToken] depends_on.
func (assct *AppServiceSourceControlToken) Dependencies() terra.Dependencies {
	return assct.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) LifecycleManagement() *terra.Lifecycle {
	return assct.Lifecycle
}

// Attributes returns the attributes for [AppServiceSourceControlToken].
func (assct *AppServiceSourceControlToken) Attributes() appServiceSourceControlTokenAttributes {
	return appServiceSourceControlTokenAttributes{ref: terra.ReferenceResource(assct)}
}

// ImportState imports the given attribute values into [AppServiceSourceControlToken]'s state.
func (assct *AppServiceSourceControlToken) ImportState(av io.Reader) error {
	assct.state = &appServiceSourceControlTokenState{}
	if err := json.NewDecoder(av).Decode(assct.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", assct.Type(), assct.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppServiceSourceControlToken] has state.
func (assct *AppServiceSourceControlToken) State() (*appServiceSourceControlTokenState, bool) {
	return assct.state, assct.state != nil
}

// StateMust returns the state for [AppServiceSourceControlToken]. Panics if the state is nil.
func (assct *AppServiceSourceControlToken) StateMust() *appServiceSourceControlTokenState {
	if assct.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", assct.Type(), assct.LocalName()))
	}
	return assct.state
}

// AppServiceSourceControlTokenArgs contains the configurations for azurerm_app_service_source_control_token.
type AppServiceSourceControlTokenArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Token: string, required
	Token terra.StringValue `hcl:"token,attr" validate:"required"`
	// TokenSecret: string, optional
	TokenSecret terra.StringValue `hcl:"token_secret,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *appservicesourcecontroltoken.Timeouts `hcl:"timeouts,block"`
}
type appServiceSourceControlTokenAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_app_service_source_control_token.
func (assct appServiceSourceControlTokenAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assct.ref.Append("id"))
}

// Token returns a reference to field token of azurerm_app_service_source_control_token.
func (assct appServiceSourceControlTokenAttributes) Token() terra.StringValue {
	return terra.ReferenceAsString(assct.ref.Append("token"))
}

// TokenSecret returns a reference to field token_secret of azurerm_app_service_source_control_token.
func (assct appServiceSourceControlTokenAttributes) TokenSecret() terra.StringValue {
	return terra.ReferenceAsString(assct.ref.Append("token_secret"))
}

// Type returns a reference to field type of azurerm_app_service_source_control_token.
func (assct appServiceSourceControlTokenAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(assct.ref.Append("type"))
}

func (assct appServiceSourceControlTokenAttributes) Timeouts() appservicesourcecontroltoken.TimeoutsAttributes {
	return terra.ReferenceAsSingle[appservicesourcecontroltoken.TimeoutsAttributes](assct.ref.Append("timeouts"))
}

type appServiceSourceControlTokenState struct {
	Id          string                                      `json:"id"`
	Token       string                                      `json:"token"`
	TokenSecret string                                      `json:"token_secret"`
	Type        string                                      `json:"type"`
	Timeouts    *appservicesourcecontroltoken.TimeoutsState `json:"timeouts"`
}
