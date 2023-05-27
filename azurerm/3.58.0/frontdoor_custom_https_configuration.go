// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	frontdoorcustomhttpsconfiguration "github.com/golingon/terraproviders/azurerm/3.58.0/frontdoorcustomhttpsconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFrontdoorCustomHttpsConfiguration creates a new instance of [FrontdoorCustomHttpsConfiguration].
func NewFrontdoorCustomHttpsConfiguration(name string, args FrontdoorCustomHttpsConfigurationArgs) *FrontdoorCustomHttpsConfiguration {
	return &FrontdoorCustomHttpsConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FrontdoorCustomHttpsConfiguration)(nil)

// FrontdoorCustomHttpsConfiguration represents the Terraform resource azurerm_frontdoor_custom_https_configuration.
type FrontdoorCustomHttpsConfiguration struct {
	Name      string
	Args      FrontdoorCustomHttpsConfigurationArgs
	state     *frontdoorCustomHttpsConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) Type() string {
	return "azurerm_frontdoor_custom_https_configuration"
}

// LocalName returns the local name for [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) LocalName() string {
	return fchc.Name
}

// Configuration returns the configuration (args) for [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) Configuration() interface{} {
	return fchc.Args
}

// DependOn is used for other resources to depend on [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(fchc)
}

// Dependencies returns the list of resources [FrontdoorCustomHttpsConfiguration] depends_on.
func (fchc *FrontdoorCustomHttpsConfiguration) Dependencies() terra.Dependencies {
	return fchc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) LifecycleManagement() *terra.Lifecycle {
	return fchc.Lifecycle
}

// Attributes returns the attributes for [FrontdoorCustomHttpsConfiguration].
func (fchc *FrontdoorCustomHttpsConfiguration) Attributes() frontdoorCustomHttpsConfigurationAttributes {
	return frontdoorCustomHttpsConfigurationAttributes{ref: terra.ReferenceResource(fchc)}
}

// ImportState imports the given attribute values into [FrontdoorCustomHttpsConfiguration]'s state.
func (fchc *FrontdoorCustomHttpsConfiguration) ImportState(av io.Reader) error {
	fchc.state = &frontdoorCustomHttpsConfigurationState{}
	if err := json.NewDecoder(av).Decode(fchc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fchc.Type(), fchc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FrontdoorCustomHttpsConfiguration] has state.
func (fchc *FrontdoorCustomHttpsConfiguration) State() (*frontdoorCustomHttpsConfigurationState, bool) {
	return fchc.state, fchc.state != nil
}

// StateMust returns the state for [FrontdoorCustomHttpsConfiguration]. Panics if the state is nil.
func (fchc *FrontdoorCustomHttpsConfiguration) StateMust() *frontdoorCustomHttpsConfigurationState {
	if fchc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fchc.Type(), fchc.LocalName()))
	}
	return fchc.state
}

// FrontdoorCustomHttpsConfigurationArgs contains the configurations for azurerm_frontdoor_custom_https_configuration.
type FrontdoorCustomHttpsConfigurationArgs struct {
	// CustomHttpsProvisioningEnabled: bool, required
	CustomHttpsProvisioningEnabled terra.BoolValue `hcl:"custom_https_provisioning_enabled,attr" validate:"required"`
	// FrontendEndpointId: string, required
	FrontendEndpointId terra.StringValue `hcl:"frontend_endpoint_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// CustomHttpsConfiguration: optional
	CustomHttpsConfiguration *frontdoorcustomhttpsconfiguration.CustomHttpsConfiguration `hcl:"custom_https_configuration,block"`
	// Timeouts: optional
	Timeouts *frontdoorcustomhttpsconfiguration.Timeouts `hcl:"timeouts,block"`
}
type frontdoorCustomHttpsConfigurationAttributes struct {
	ref terra.Reference
}

// CustomHttpsProvisioningEnabled returns a reference to field custom_https_provisioning_enabled of azurerm_frontdoor_custom_https_configuration.
func (fchc frontdoorCustomHttpsConfigurationAttributes) CustomHttpsProvisioningEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(fchc.ref.Append("custom_https_provisioning_enabled"))
}

// FrontendEndpointId returns a reference to field frontend_endpoint_id of azurerm_frontdoor_custom_https_configuration.
func (fchc frontdoorCustomHttpsConfigurationAttributes) FrontendEndpointId() terra.StringValue {
	return terra.ReferenceAsString(fchc.ref.Append("frontend_endpoint_id"))
}

// Id returns a reference to field id of azurerm_frontdoor_custom_https_configuration.
func (fchc frontdoorCustomHttpsConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fchc.ref.Append("id"))
}

func (fchc frontdoorCustomHttpsConfigurationAttributes) CustomHttpsConfiguration() terra.ListValue[frontdoorcustomhttpsconfiguration.CustomHttpsConfigurationAttributes] {
	return terra.ReferenceAsList[frontdoorcustomhttpsconfiguration.CustomHttpsConfigurationAttributes](fchc.ref.Append("custom_https_configuration"))
}

func (fchc frontdoorCustomHttpsConfigurationAttributes) Timeouts() frontdoorcustomhttpsconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[frontdoorcustomhttpsconfiguration.TimeoutsAttributes](fchc.ref.Append("timeouts"))
}

type frontdoorCustomHttpsConfigurationState struct {
	CustomHttpsProvisioningEnabled bool                                                              `json:"custom_https_provisioning_enabled"`
	FrontendEndpointId             string                                                            `json:"frontend_endpoint_id"`
	Id                             string                                                            `json:"id"`
	CustomHttpsConfiguration       []frontdoorcustomhttpsconfiguration.CustomHttpsConfigurationState `json:"custom_https_configuration"`
	Timeouts                       *frontdoorcustomhttpsconfiguration.TimeoutsState                  `json:"timeouts"`
}
