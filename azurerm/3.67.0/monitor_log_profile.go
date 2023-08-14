// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitorlogprofile "github.com/golingon/terraproviders/azurerm/3.67.0/monitorlogprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorLogProfile creates a new instance of [MonitorLogProfile].
func NewMonitorLogProfile(name string, args MonitorLogProfileArgs) *MonitorLogProfile {
	return &MonitorLogProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorLogProfile)(nil)

// MonitorLogProfile represents the Terraform resource azurerm_monitor_log_profile.
type MonitorLogProfile struct {
	Name      string
	Args      MonitorLogProfileArgs
	state     *monitorLogProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorLogProfile].
func (mlp *MonitorLogProfile) Type() string {
	return "azurerm_monitor_log_profile"
}

// LocalName returns the local name for [MonitorLogProfile].
func (mlp *MonitorLogProfile) LocalName() string {
	return mlp.Name
}

// Configuration returns the configuration (args) for [MonitorLogProfile].
func (mlp *MonitorLogProfile) Configuration() interface{} {
	return mlp.Args
}

// DependOn is used for other resources to depend on [MonitorLogProfile].
func (mlp *MonitorLogProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(mlp)
}

// Dependencies returns the list of resources [MonitorLogProfile] depends_on.
func (mlp *MonitorLogProfile) Dependencies() terra.Dependencies {
	return mlp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorLogProfile].
func (mlp *MonitorLogProfile) LifecycleManagement() *terra.Lifecycle {
	return mlp.Lifecycle
}

// Attributes returns the attributes for [MonitorLogProfile].
func (mlp *MonitorLogProfile) Attributes() monitorLogProfileAttributes {
	return monitorLogProfileAttributes{ref: terra.ReferenceResource(mlp)}
}

// ImportState imports the given attribute values into [MonitorLogProfile]'s state.
func (mlp *MonitorLogProfile) ImportState(av io.Reader) error {
	mlp.state = &monitorLogProfileState{}
	if err := json.NewDecoder(av).Decode(mlp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlp.Type(), mlp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorLogProfile] has state.
func (mlp *MonitorLogProfile) State() (*monitorLogProfileState, bool) {
	return mlp.state, mlp.state != nil
}

// StateMust returns the state for [MonitorLogProfile]. Panics if the state is nil.
func (mlp *MonitorLogProfile) StateMust() *monitorLogProfileState {
	if mlp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlp.Type(), mlp.LocalName()))
	}
	return mlp.state
}

// MonitorLogProfileArgs contains the configurations for azurerm_monitor_log_profile.
type MonitorLogProfileArgs struct {
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
	RetentionPolicy *monitorlogprofile.RetentionPolicy `hcl:"retention_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitorlogprofile.Timeouts `hcl:"timeouts,block"`
}
type monitorLogProfileAttributes struct {
	ref terra.Reference
}

// Categories returns a reference to field categories of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) Categories() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mlp.ref.Append("categories"))
}

// Id returns a reference to field id of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("id"))
}

// Locations returns a reference to field locations of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) Locations() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mlp.ref.Append("locations"))
}

// Name returns a reference to field name of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("name"))
}

// ServicebusRuleId returns a reference to field servicebus_rule_id of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) ServicebusRuleId() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("servicebus_rule_id"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_monitor_log_profile.
func (mlp monitorLogProfileAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(mlp.ref.Append("storage_account_id"))
}

func (mlp monitorLogProfileAttributes) RetentionPolicy() terra.ListValue[monitorlogprofile.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[monitorlogprofile.RetentionPolicyAttributes](mlp.ref.Append("retention_policy"))
}

func (mlp monitorLogProfileAttributes) Timeouts() monitorlogprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitorlogprofile.TimeoutsAttributes](mlp.ref.Append("timeouts"))
}

type monitorLogProfileState struct {
	Categories       []string                                 `json:"categories"`
	Id               string                                   `json:"id"`
	Locations        []string                                 `json:"locations"`
	Name             string                                   `json:"name"`
	ServicebusRuleId string                                   `json:"servicebus_rule_id"`
	StorageAccountId string                                   `json:"storage_account_id"`
	RetentionPolicy  []monitorlogprofile.RetentionPolicyState `json:"retention_policy"`
	Timeouts         *monitorlogprofile.TimeoutsState         `json:"timeouts"`
}