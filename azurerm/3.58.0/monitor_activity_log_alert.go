// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoractivitylogalert "github.com/golingon/terraproviders/azurerm/3.58.0/monitoractivitylogalert"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorActivityLogAlert creates a new instance of [MonitorActivityLogAlert].
func NewMonitorActivityLogAlert(name string, args MonitorActivityLogAlertArgs) *MonitorActivityLogAlert {
	return &MonitorActivityLogAlert{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorActivityLogAlert)(nil)

// MonitorActivityLogAlert represents the Terraform resource azurerm_monitor_activity_log_alert.
type MonitorActivityLogAlert struct {
	Name      string
	Args      MonitorActivityLogAlertArgs
	state     *monitorActivityLogAlertState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) Type() string {
	return "azurerm_monitor_activity_log_alert"
}

// LocalName returns the local name for [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) LocalName() string {
	return mala.Name
}

// Configuration returns the configuration (args) for [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) Configuration() interface{} {
	return mala.Args
}

// DependOn is used for other resources to depend on [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) DependOn() terra.Reference {
	return terra.ReferenceResource(mala)
}

// Dependencies returns the list of resources [MonitorActivityLogAlert] depends_on.
func (mala *MonitorActivityLogAlert) Dependencies() terra.Dependencies {
	return mala.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) LifecycleManagement() *terra.Lifecycle {
	return mala.Lifecycle
}

// Attributes returns the attributes for [MonitorActivityLogAlert].
func (mala *MonitorActivityLogAlert) Attributes() monitorActivityLogAlertAttributes {
	return monitorActivityLogAlertAttributes{ref: terra.ReferenceResource(mala)}
}

// ImportState imports the given attribute values into [MonitorActivityLogAlert]'s state.
func (mala *MonitorActivityLogAlert) ImportState(av io.Reader) error {
	mala.state = &monitorActivityLogAlertState{}
	if err := json.NewDecoder(av).Decode(mala.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mala.Type(), mala.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorActivityLogAlert] has state.
func (mala *MonitorActivityLogAlert) State() (*monitorActivityLogAlertState, bool) {
	return mala.state, mala.state != nil
}

// StateMust returns the state for [MonitorActivityLogAlert]. Panics if the state is nil.
func (mala *MonitorActivityLogAlert) StateMust() *monitorActivityLogAlertState {
	if mala.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mala.Type(), mala.LocalName()))
	}
	return mala.state
}

// MonitorActivityLogAlertArgs contains the configurations for azurerm_monitor_activity_log_alert.
type MonitorActivityLogAlertArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Scopes: set of string, required
	Scopes terra.SetValue[terra.StringValue] `hcl:"scopes,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Action: min=0
	Action []monitoractivitylogalert.Action `hcl:"action,block" validate:"min=0"`
	// Criteria: required
	Criteria *monitoractivitylogalert.Criteria `hcl:"criteria,block" validate:"required"`
	// Timeouts: optional
	Timeouts *monitoractivitylogalert.Timeouts `hcl:"timeouts,block"`
}
type monitorActivityLogAlertAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mala.ref.Append("description"))
}

// Enabled returns a reference to field enabled of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mala.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mala.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mala.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mala.ref.Append("resource_group_name"))
}

// Scopes returns a reference to field scopes of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Scopes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mala.ref.Append("scopes"))
}

// Tags returns a reference to field tags of azurerm_monitor_activity_log_alert.
func (mala monitorActivityLogAlertAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mala.ref.Append("tags"))
}

func (mala monitorActivityLogAlertAttributes) Action() terra.ListValue[monitoractivitylogalert.ActionAttributes] {
	return terra.ReferenceAsList[monitoractivitylogalert.ActionAttributes](mala.ref.Append("action"))
}

func (mala monitorActivityLogAlertAttributes) Criteria() terra.ListValue[monitoractivitylogalert.CriteriaAttributes] {
	return terra.ReferenceAsList[monitoractivitylogalert.CriteriaAttributes](mala.ref.Append("criteria"))
}

func (mala monitorActivityLogAlertAttributes) Timeouts() monitoractivitylogalert.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoractivitylogalert.TimeoutsAttributes](mala.ref.Append("timeouts"))
}

type monitorActivityLogAlertState struct {
	Description       string                                  `json:"description"`
	Enabled           bool                                    `json:"enabled"`
	Id                string                                  `json:"id"`
	Name              string                                  `json:"name"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	Scopes            []string                                `json:"scopes"`
	Tags              map[string]string                       `json:"tags"`
	Action            []monitoractivitylogalert.ActionState   `json:"action"`
	Criteria          []monitoractivitylogalert.CriteriaState `json:"criteria"`
	Timeouts          *monitoractivitylogalert.TimeoutsState  `json:"timeouts"`
}
