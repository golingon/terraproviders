// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	monitoractiongroup "github.com/golingon/terraproviders/azurerm/3.69.0/monitoractiongroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMonitorActionGroup creates a new instance of [MonitorActionGroup].
func NewMonitorActionGroup(name string, args MonitorActionGroupArgs) *MonitorActionGroup {
	return &MonitorActionGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MonitorActionGroup)(nil)

// MonitorActionGroup represents the Terraform resource azurerm_monitor_action_group.
type MonitorActionGroup struct {
	Name      string
	Args      MonitorActionGroupArgs
	state     *monitorActionGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MonitorActionGroup].
func (mag *MonitorActionGroup) Type() string {
	return "azurerm_monitor_action_group"
}

// LocalName returns the local name for [MonitorActionGroup].
func (mag *MonitorActionGroup) LocalName() string {
	return mag.Name
}

// Configuration returns the configuration (args) for [MonitorActionGroup].
func (mag *MonitorActionGroup) Configuration() interface{} {
	return mag.Args
}

// DependOn is used for other resources to depend on [MonitorActionGroup].
func (mag *MonitorActionGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(mag)
}

// Dependencies returns the list of resources [MonitorActionGroup] depends_on.
func (mag *MonitorActionGroup) Dependencies() terra.Dependencies {
	return mag.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MonitorActionGroup].
func (mag *MonitorActionGroup) LifecycleManagement() *terra.Lifecycle {
	return mag.Lifecycle
}

// Attributes returns the attributes for [MonitorActionGroup].
func (mag *MonitorActionGroup) Attributes() monitorActionGroupAttributes {
	return monitorActionGroupAttributes{ref: terra.ReferenceResource(mag)}
}

// ImportState imports the given attribute values into [MonitorActionGroup]'s state.
func (mag *MonitorActionGroup) ImportState(av io.Reader) error {
	mag.state = &monitorActionGroupState{}
	if err := json.NewDecoder(av).Decode(mag.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mag.Type(), mag.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MonitorActionGroup] has state.
func (mag *MonitorActionGroup) State() (*monitorActionGroupState, bool) {
	return mag.state, mag.state != nil
}

// StateMust returns the state for [MonitorActionGroup]. Panics if the state is nil.
func (mag *MonitorActionGroup) StateMust() *monitorActionGroupState {
	if mag.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mag.Type(), mag.LocalName()))
	}
	return mag.state
}

// MonitorActionGroupArgs contains the configurations for azurerm_monitor_action_group.
type MonitorActionGroupArgs struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ShortName: string, required
	ShortName terra.StringValue `hcl:"short_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ArmRoleReceiver: min=0
	ArmRoleReceiver []monitoractiongroup.ArmRoleReceiver `hcl:"arm_role_receiver,block" validate:"min=0"`
	// AutomationRunbookReceiver: min=0
	AutomationRunbookReceiver []monitoractiongroup.AutomationRunbookReceiver `hcl:"automation_runbook_receiver,block" validate:"min=0"`
	// AzureAppPushReceiver: min=0
	AzureAppPushReceiver []monitoractiongroup.AzureAppPushReceiver `hcl:"azure_app_push_receiver,block" validate:"min=0"`
	// AzureFunctionReceiver: min=0
	AzureFunctionReceiver []monitoractiongroup.AzureFunctionReceiver `hcl:"azure_function_receiver,block" validate:"min=0"`
	// EmailReceiver: min=0
	EmailReceiver []monitoractiongroup.EmailReceiver `hcl:"email_receiver,block" validate:"min=0"`
	// EventHubReceiver: min=0
	EventHubReceiver []monitoractiongroup.EventHubReceiver `hcl:"event_hub_receiver,block" validate:"min=0"`
	// ItsmReceiver: min=0
	ItsmReceiver []monitoractiongroup.ItsmReceiver `hcl:"itsm_receiver,block" validate:"min=0"`
	// LogicAppReceiver: min=0
	LogicAppReceiver []monitoractiongroup.LogicAppReceiver `hcl:"logic_app_receiver,block" validate:"min=0"`
	// SmsReceiver: min=0
	SmsReceiver []monitoractiongroup.SmsReceiver `hcl:"sms_receiver,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *monitoractiongroup.Timeouts `hcl:"timeouts,block"`
	// VoiceReceiver: min=0
	VoiceReceiver []monitoractiongroup.VoiceReceiver `hcl:"voice_receiver,block" validate:"min=0"`
	// WebhookReceiver: min=0
	WebhookReceiver []monitoractiongroup.WebhookReceiver `hcl:"webhook_receiver,block" validate:"min=0"`
}
type monitorActionGroupAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mag.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("resource_group_name"))
}

// ShortName returns a reference to field short_name of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("short_name"))
}

// Tags returns a reference to field tags of azurerm_monitor_action_group.
func (mag monitorActionGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mag.ref.Append("tags"))
}

func (mag monitorActionGroupAttributes) ArmRoleReceiver() terra.ListValue[monitoractiongroup.ArmRoleReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.ArmRoleReceiverAttributes](mag.ref.Append("arm_role_receiver"))
}

func (mag monitorActionGroupAttributes) AutomationRunbookReceiver() terra.ListValue[monitoractiongroup.AutomationRunbookReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.AutomationRunbookReceiverAttributes](mag.ref.Append("automation_runbook_receiver"))
}

func (mag monitorActionGroupAttributes) AzureAppPushReceiver() terra.ListValue[monitoractiongroup.AzureAppPushReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.AzureAppPushReceiverAttributes](mag.ref.Append("azure_app_push_receiver"))
}

func (mag monitorActionGroupAttributes) AzureFunctionReceiver() terra.ListValue[monitoractiongroup.AzureFunctionReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.AzureFunctionReceiverAttributes](mag.ref.Append("azure_function_receiver"))
}

func (mag monitorActionGroupAttributes) EmailReceiver() terra.ListValue[monitoractiongroup.EmailReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.EmailReceiverAttributes](mag.ref.Append("email_receiver"))
}

func (mag monitorActionGroupAttributes) EventHubReceiver() terra.ListValue[monitoractiongroup.EventHubReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.EventHubReceiverAttributes](mag.ref.Append("event_hub_receiver"))
}

func (mag monitorActionGroupAttributes) ItsmReceiver() terra.ListValue[monitoractiongroup.ItsmReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.ItsmReceiverAttributes](mag.ref.Append("itsm_receiver"))
}

func (mag monitorActionGroupAttributes) LogicAppReceiver() terra.ListValue[monitoractiongroup.LogicAppReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.LogicAppReceiverAttributes](mag.ref.Append("logic_app_receiver"))
}

func (mag monitorActionGroupAttributes) SmsReceiver() terra.ListValue[monitoractiongroup.SmsReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.SmsReceiverAttributes](mag.ref.Append("sms_receiver"))
}

func (mag monitorActionGroupAttributes) Timeouts() monitoractiongroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[monitoractiongroup.TimeoutsAttributes](mag.ref.Append("timeouts"))
}

func (mag monitorActionGroupAttributes) VoiceReceiver() terra.ListValue[monitoractiongroup.VoiceReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.VoiceReceiverAttributes](mag.ref.Append("voice_receiver"))
}

func (mag monitorActionGroupAttributes) WebhookReceiver() terra.ListValue[monitoractiongroup.WebhookReceiverAttributes] {
	return terra.ReferenceAsList[monitoractiongroup.WebhookReceiverAttributes](mag.ref.Append("webhook_receiver"))
}

type monitorActionGroupState struct {
	Enabled                   bool                                                `json:"enabled"`
	Id                        string                                              `json:"id"`
	Location                  string                                              `json:"location"`
	Name                      string                                              `json:"name"`
	ResourceGroupName         string                                              `json:"resource_group_name"`
	ShortName                 string                                              `json:"short_name"`
	Tags                      map[string]string                                   `json:"tags"`
	ArmRoleReceiver           []monitoractiongroup.ArmRoleReceiverState           `json:"arm_role_receiver"`
	AutomationRunbookReceiver []monitoractiongroup.AutomationRunbookReceiverState `json:"automation_runbook_receiver"`
	AzureAppPushReceiver      []monitoractiongroup.AzureAppPushReceiverState      `json:"azure_app_push_receiver"`
	AzureFunctionReceiver     []monitoractiongroup.AzureFunctionReceiverState     `json:"azure_function_receiver"`
	EmailReceiver             []monitoractiongroup.EmailReceiverState             `json:"email_receiver"`
	EventHubReceiver          []monitoractiongroup.EventHubReceiverState          `json:"event_hub_receiver"`
	ItsmReceiver              []monitoractiongroup.ItsmReceiverState              `json:"itsm_receiver"`
	LogicAppReceiver          []monitoractiongroup.LogicAppReceiverState          `json:"logic_app_receiver"`
	SmsReceiver               []monitoractiongroup.SmsReceiverState               `json:"sms_receiver"`
	Timeouts                  *monitoractiongroup.TimeoutsState                   `json:"timeouts"`
	VoiceReceiver             []monitoractiongroup.VoiceReceiverState             `json:"voice_receiver"`
	WebhookReceiver           []monitoractiongroup.WebhookReceiverState           `json:"webhook_receiver"`
}
