// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_monitor_log_profile

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_monitor_log_profile.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMonitorLogProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amlp *Resource) Type() string {
	return "azurerm_monitor_log_profile"
}

// LocalName returns the local name for [Resource].
func (amlp *Resource) LocalName() string {
	return amlp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amlp *Resource) Configuration() interface{} {
	return amlp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amlp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amlp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amlp *Resource) Dependencies() terra.Dependencies {
	return amlp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amlp *Resource) LifecycleManagement() *terra.Lifecycle {
	return amlp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amlp *Resource) Attributes() azurermMonitorLogProfileAttributes {
	return azurermMonitorLogProfileAttributes{ref: terra.ReferenceResource(amlp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amlp *Resource) ImportState(state io.Reader) error {
	amlp.state = &azurermMonitorLogProfileState{}
	if err := json.NewDecoder(state).Decode(amlp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amlp.Type(), amlp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amlp *Resource) State() (*azurermMonitorLogProfileState, bool) {
	return amlp.state, amlp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amlp *Resource) StateMust() *azurermMonitorLogProfileState {
	if amlp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amlp.Type(), amlp.LocalName()))
	}
	return amlp.state
}

// Args contains the configurations for azurerm_monitor_log_profile.
type Args struct {
	// Categories: set of string, required
	Categories terra.SetValue[terra.StringValue] `hcl:"categories,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Locations: set of string, required
	Locations terra.SetValue[terra.StringValue] `hcl:"locations,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ServicebusRuleId: string, optional
	ServicebusRuleId terra.StringValue `hcl:"servicebus_rule_id,attr"`
	// StorageAccountId: string, optional
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr"`
	// RetentionPolicy: required
	RetentionPolicy *RetentionPolicy `hcl:"retention_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMonitorLogProfileAttributes struct {
	ref terra.Reference
}

// Categories returns a reference to field categories of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) Categories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amlp.ref.Append("categories"))
}

// Id returns a reference to field id of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amlp.ref.Append("id"))
}

// Locations returns a reference to field locations of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) Locations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](amlp.ref.Append("locations"))
}

// Name returns a reference to field name of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amlp.ref.Append("name"))
}

// ServicebusRuleId returns a reference to field servicebus_rule_id of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) ServicebusRuleId() terra.StringValue {
	return terra.ReferenceAsString(amlp.ref.Append("servicebus_rule_id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_monitor_log_profile.
func (amlp azurermMonitorLogProfileAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(amlp.ref.Append("storage_account_id"))
}

func (amlp azurermMonitorLogProfileAttributes) RetentionPolicy() terra.ListValue[RetentionPolicyAttributes] {
	return terra.ReferenceAsList[RetentionPolicyAttributes](amlp.ref.Append("retention_policy"))
}

func (amlp azurermMonitorLogProfileAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amlp.ref.Append("timeouts"))
}

type azurermMonitorLogProfileState struct {
	Categories       []string               `json:"categories"`
	Id               string                 `json:"id"`
	Locations        []string               `json:"locations"`
	Name             string                 `json:"name"`
	ServicebusRuleId string                 `json:"servicebus_rule_id"`
	StorageAccountId string                 `json:"storage_account_id"`
	RetentionPolicy  []RetentionPolicyState `json:"retention_policy"`
	Timeouts         *TimeoutsState         `json:"timeouts"`
}
