// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datadogmonitorssoconfiguration "github.com/golingon/terraproviders/azurerm/3.67.0/datadogmonitorssoconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatadogMonitorSsoConfiguration creates a new instance of [DatadogMonitorSsoConfiguration].
func NewDatadogMonitorSsoConfiguration(name string, args DatadogMonitorSsoConfigurationArgs) *DatadogMonitorSsoConfiguration {
	return &DatadogMonitorSsoConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatadogMonitorSsoConfiguration)(nil)

// DatadogMonitorSsoConfiguration represents the Terraform resource azurerm_datadog_monitor_sso_configuration.
type DatadogMonitorSsoConfiguration struct {
	Name      string
	Args      DatadogMonitorSsoConfigurationArgs
	state     *datadogMonitorSsoConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) Type() string {
	return "azurerm_datadog_monitor_sso_configuration"
}

// LocalName returns the local name for [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) LocalName() string {
	return dmsc.Name
}

// Configuration returns the configuration (args) for [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) Configuration() interface{} {
	return dmsc.Args
}

// DependOn is used for other resources to depend on [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(dmsc)
}

// Dependencies returns the list of resources [DatadogMonitorSsoConfiguration] depends_on.
func (dmsc *DatadogMonitorSsoConfiguration) Dependencies() terra.Dependencies {
	return dmsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) LifecycleManagement() *terra.Lifecycle {
	return dmsc.Lifecycle
}

// Attributes returns the attributes for [DatadogMonitorSsoConfiguration].
func (dmsc *DatadogMonitorSsoConfiguration) Attributes() datadogMonitorSsoConfigurationAttributes {
	return datadogMonitorSsoConfigurationAttributes{ref: terra.ReferenceResource(dmsc)}
}

// ImportState imports the given attribute values into [DatadogMonitorSsoConfiguration]'s state.
func (dmsc *DatadogMonitorSsoConfiguration) ImportState(av io.Reader) error {
	dmsc.state = &datadogMonitorSsoConfigurationState{}
	if err := json.NewDecoder(av).Decode(dmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dmsc.Type(), dmsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatadogMonitorSsoConfiguration] has state.
func (dmsc *DatadogMonitorSsoConfiguration) State() (*datadogMonitorSsoConfigurationState, bool) {
	return dmsc.state, dmsc.state != nil
}

// StateMust returns the state for [DatadogMonitorSsoConfiguration]. Panics if the state is nil.
func (dmsc *DatadogMonitorSsoConfiguration) StateMust() *datadogMonitorSsoConfigurationState {
	if dmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dmsc.Type(), dmsc.LocalName()))
	}
	return dmsc.state
}

// DatadogMonitorSsoConfigurationArgs contains the configurations for azurerm_datadog_monitor_sso_configuration.
type DatadogMonitorSsoConfigurationArgs struct {
	// DatadogMonitorId: string, required
	DatadogMonitorId terra.StringValue `hcl:"datadog_monitor_id,attr" validate:"required"`
	// EnterpriseApplicationId: string, required
	EnterpriseApplicationId terra.StringValue `hcl:"enterprise_application_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// SingleSignOnEnabled: string, required
	SingleSignOnEnabled terra.StringValue `hcl:"single_sign_on_enabled,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadogmonitorssoconfiguration.Timeouts `hcl:"timeouts,block"`
}
type datadogMonitorSsoConfigurationAttributes struct {
	ref terra.Reference
}

// DatadogMonitorId returns a reference to field datadog_monitor_id of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) DatadogMonitorId() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("datadog_monitor_id"))
}

// EnterpriseApplicationId returns a reference to field enterprise_application_id of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) EnterpriseApplicationId() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("enterprise_application_id"))
}

// Id returns a reference to field id of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("id"))
}

// LoginUrl returns a reference to field login_url of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) LoginUrl() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("login_url"))
}

// Name returns a reference to field name of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("name"))
}

// SingleSignOnEnabled returns a reference to field single_sign_on_enabled of azurerm_datadog_monitor_sso_configuration.
func (dmsc datadogMonitorSsoConfigurationAttributes) SingleSignOnEnabled() terra.StringValue {
	return terra.ReferenceAsString(dmsc.ref.Append("single_sign_on_enabled"))
}

func (dmsc datadogMonitorSsoConfigurationAttributes) Timeouts() datadogmonitorssoconfiguration.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadogmonitorssoconfiguration.TimeoutsAttributes](dmsc.ref.Append("timeouts"))
}

type datadogMonitorSsoConfigurationState struct {
	DatadogMonitorId        string                                        `json:"datadog_monitor_id"`
	EnterpriseApplicationId string                                        `json:"enterprise_application_id"`
	Id                      string                                        `json:"id"`
	LoginUrl                string                                        `json:"login_url"`
	Name                    string                                        `json:"name"`
	SingleSignOnEnabled     string                                        `json:"single_sign_on_enabled"`
	Timeouts                *datadogmonitorssoconfiguration.TimeoutsState `json:"timeouts"`
}
