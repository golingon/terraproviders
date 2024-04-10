// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	staticwebappfunctionappregistration "github.com/golingon/terraproviders/azurerm/3.98.0/staticwebappfunctionappregistration"
	"io"
)

// NewStaticWebAppFunctionAppRegistration creates a new instance of [StaticWebAppFunctionAppRegistration].
func NewStaticWebAppFunctionAppRegistration(name string, args StaticWebAppFunctionAppRegistrationArgs) *StaticWebAppFunctionAppRegistration {
	return &StaticWebAppFunctionAppRegistration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StaticWebAppFunctionAppRegistration)(nil)

// StaticWebAppFunctionAppRegistration represents the Terraform resource azurerm_static_web_app_function_app_registration.
type StaticWebAppFunctionAppRegistration struct {
	Name      string
	Args      StaticWebAppFunctionAppRegistrationArgs
	state     *staticWebAppFunctionAppRegistrationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) Type() string {
	return "azurerm_static_web_app_function_app_registration"
}

// LocalName returns the local name for [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) LocalName() string {
	return swafar.Name
}

// Configuration returns the configuration (args) for [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) Configuration() interface{} {
	return swafar.Args
}

// DependOn is used for other resources to depend on [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) DependOn() terra.Reference {
	return terra.ReferenceResource(swafar)
}

// Dependencies returns the list of resources [StaticWebAppFunctionAppRegistration] depends_on.
func (swafar *StaticWebAppFunctionAppRegistration) Dependencies() terra.Dependencies {
	return swafar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) LifecycleManagement() *terra.Lifecycle {
	return swafar.Lifecycle
}

// Attributes returns the attributes for [StaticWebAppFunctionAppRegistration].
func (swafar *StaticWebAppFunctionAppRegistration) Attributes() staticWebAppFunctionAppRegistrationAttributes {
	return staticWebAppFunctionAppRegistrationAttributes{ref: terra.ReferenceResource(swafar)}
}

// ImportState imports the given attribute values into [StaticWebAppFunctionAppRegistration]'s state.
func (swafar *StaticWebAppFunctionAppRegistration) ImportState(av io.Reader) error {
	swafar.state = &staticWebAppFunctionAppRegistrationState{}
	if err := json.NewDecoder(av).Decode(swafar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", swafar.Type(), swafar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StaticWebAppFunctionAppRegistration] has state.
func (swafar *StaticWebAppFunctionAppRegistration) State() (*staticWebAppFunctionAppRegistrationState, bool) {
	return swafar.state, swafar.state != nil
}

// StateMust returns the state for [StaticWebAppFunctionAppRegistration]. Panics if the state is nil.
func (swafar *StaticWebAppFunctionAppRegistration) StateMust() *staticWebAppFunctionAppRegistrationState {
	if swafar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", swafar.Type(), swafar.LocalName()))
	}
	return swafar.state
}

// StaticWebAppFunctionAppRegistrationArgs contains the configurations for azurerm_static_web_app_function_app_registration.
type StaticWebAppFunctionAppRegistrationArgs struct {
	// FunctionAppId: string, required
	FunctionAppId terra.StringValue `hcl:"function_app_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StaticWebAppId: string, required
	StaticWebAppId terra.StringValue `hcl:"static_web_app_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *staticwebappfunctionappregistration.Timeouts `hcl:"timeouts,block"`
}
type staticWebAppFunctionAppRegistrationAttributes struct {
	ref terra.Reference
}

// FunctionAppId returns a reference to field function_app_id of azurerm_static_web_app_function_app_registration.
func (swafar staticWebAppFunctionAppRegistrationAttributes) FunctionAppId() terra.StringValue {
	return terra.ReferenceAsString(swafar.ref.Append("function_app_id"))
}

// Id returns a reference to field id of azurerm_static_web_app_function_app_registration.
func (swafar staticWebAppFunctionAppRegistrationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(swafar.ref.Append("id"))
}

// StaticWebAppId returns a reference to field static_web_app_id of azurerm_static_web_app_function_app_registration.
func (swafar staticWebAppFunctionAppRegistrationAttributes) StaticWebAppId() terra.StringValue {
	return terra.ReferenceAsString(swafar.ref.Append("static_web_app_id"))
}

func (swafar staticWebAppFunctionAppRegistrationAttributes) Timeouts() staticwebappfunctionappregistration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[staticwebappfunctionappregistration.TimeoutsAttributes](swafar.ref.Append("timeouts"))
}

type staticWebAppFunctionAppRegistrationState struct {
	FunctionAppId  string                                             `json:"function_app_id"`
	Id             string                                             `json:"id"`
	StaticWebAppId string                                             `json:"static_web_app_id"`
	Timeouts       *staticwebappfunctionappregistration.TimeoutsState `json:"timeouts"`
}
