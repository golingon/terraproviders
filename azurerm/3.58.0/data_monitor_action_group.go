// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datamonitoractiongroup "github.com/golingon/terraproviders/azurerm/3.58.0/datamonitoractiongroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataMonitorActionGroup creates a new instance of [DataMonitorActionGroup].
func NewDataMonitorActionGroup(name string, args DataMonitorActionGroupArgs) *DataMonitorActionGroup {
	return &DataMonitorActionGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataMonitorActionGroup)(nil)

// DataMonitorActionGroup represents the Terraform data resource azurerm_monitor_action_group.
type DataMonitorActionGroup struct {
	Name string
	Args DataMonitorActionGroupArgs
}

// DataSource returns the Terraform object type for [DataMonitorActionGroup].
func (mag *DataMonitorActionGroup) DataSource() string {
	return "azurerm_monitor_action_group"
}

// LocalName returns the local name for [DataMonitorActionGroup].
func (mag *DataMonitorActionGroup) LocalName() string {
	return mag.Name
}

// Configuration returns the configuration (args) for [DataMonitorActionGroup].
func (mag *DataMonitorActionGroup) Configuration() interface{} {
	return mag.Args
}

// Attributes returns the attributes for [DataMonitorActionGroup].
func (mag *DataMonitorActionGroup) Attributes() dataMonitorActionGroupAttributes {
	return dataMonitorActionGroupAttributes{ref: terra.ReferenceDataResource(mag)}
}

// DataMonitorActionGroupArgs contains the configurations for azurerm_monitor_action_group.
type DataMonitorActionGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ArmRoleReceiver: min=0
	ArmRoleReceiver []datamonitoractiongroup.ArmRoleReceiver `hcl:"arm_role_receiver,block" validate:"min=0"`
	// AutomationRunbookReceiver: min=0
	AutomationRunbookReceiver []datamonitoractiongroup.AutomationRunbookReceiver `hcl:"automation_runbook_receiver,block" validate:"min=0"`
	// AzureAppPushReceiver: min=0
	AzureAppPushReceiver []datamonitoractiongroup.AzureAppPushReceiver `hcl:"azure_app_push_receiver,block" validate:"min=0"`
	// AzureFunctionReceiver: min=0
	AzureFunctionReceiver []datamonitoractiongroup.AzureFunctionReceiver `hcl:"azure_function_receiver,block" validate:"min=0"`
	// EmailReceiver: min=0
	EmailReceiver []datamonitoractiongroup.EmailReceiver `hcl:"email_receiver,block" validate:"min=0"`
	// EventHubReceiver: min=0
	EventHubReceiver []datamonitoractiongroup.EventHubReceiver `hcl:"event_hub_receiver,block" validate:"min=0"`
	// ItsmReceiver: min=0
	ItsmReceiver []datamonitoractiongroup.ItsmReceiver `hcl:"itsm_receiver,block" validate:"min=0"`
	// LogicAppReceiver: min=0
	LogicAppReceiver []datamonitoractiongroup.LogicAppReceiver `hcl:"logic_app_receiver,block" validate:"min=0"`
	// SmsReceiver: min=0
	SmsReceiver []datamonitoractiongroup.SmsReceiver `hcl:"sms_receiver,block" validate:"min=0"`
	// VoiceReceiver: min=0
	VoiceReceiver []datamonitoractiongroup.VoiceReceiver `hcl:"voice_receiver,block" validate:"min=0"`
	// WebhookReceiver: min=0
	WebhookReceiver []datamonitoractiongroup.WebhookReceiver `hcl:"webhook_receiver,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datamonitoractiongroup.Timeouts `hcl:"timeouts,block"`
}
type dataMonitorActionGroupAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_monitor_action_group.
func (mag dataMonitorActionGroupAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mag.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_monitor_action_group.
func (mag dataMonitorActionGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_monitor_action_group.
func (mag dataMonitorActionGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_monitor_action_group.
func (mag dataMonitorActionGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("resource_group_name"))
}

// ShortName returns a reference to field short_name of azurerm_monitor_action_group.
func (mag dataMonitorActionGroupAttributes) ShortName() terra.StringValue {
	return terra.ReferenceAsString(mag.ref.Append("short_name"))
}

func (mag dataMonitorActionGroupAttributes) ArmRoleReceiver() terra.ListValue[datamonitoractiongroup.ArmRoleReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.ArmRoleReceiverAttributes](mag.ref.Append("arm_role_receiver"))
}

func (mag dataMonitorActionGroupAttributes) AutomationRunbookReceiver() terra.ListValue[datamonitoractiongroup.AutomationRunbookReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.AutomationRunbookReceiverAttributes](mag.ref.Append("automation_runbook_receiver"))
}

func (mag dataMonitorActionGroupAttributes) AzureAppPushReceiver() terra.ListValue[datamonitoractiongroup.AzureAppPushReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.AzureAppPushReceiverAttributes](mag.ref.Append("azure_app_push_receiver"))
}

func (mag dataMonitorActionGroupAttributes) AzureFunctionReceiver() terra.ListValue[datamonitoractiongroup.AzureFunctionReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.AzureFunctionReceiverAttributes](mag.ref.Append("azure_function_receiver"))
}

func (mag dataMonitorActionGroupAttributes) EmailReceiver() terra.ListValue[datamonitoractiongroup.EmailReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.EmailReceiverAttributes](mag.ref.Append("email_receiver"))
}

func (mag dataMonitorActionGroupAttributes) EventHubReceiver() terra.ListValue[datamonitoractiongroup.EventHubReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.EventHubReceiverAttributes](mag.ref.Append("event_hub_receiver"))
}

func (mag dataMonitorActionGroupAttributes) ItsmReceiver() terra.ListValue[datamonitoractiongroup.ItsmReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.ItsmReceiverAttributes](mag.ref.Append("itsm_receiver"))
}

func (mag dataMonitorActionGroupAttributes) LogicAppReceiver() terra.ListValue[datamonitoractiongroup.LogicAppReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.LogicAppReceiverAttributes](mag.ref.Append("logic_app_receiver"))
}

func (mag dataMonitorActionGroupAttributes) SmsReceiver() terra.ListValue[datamonitoractiongroup.SmsReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.SmsReceiverAttributes](mag.ref.Append("sms_receiver"))
}

func (mag dataMonitorActionGroupAttributes) VoiceReceiver() terra.ListValue[datamonitoractiongroup.VoiceReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.VoiceReceiverAttributes](mag.ref.Append("voice_receiver"))
}

func (mag dataMonitorActionGroupAttributes) WebhookReceiver() terra.ListValue[datamonitoractiongroup.WebhookReceiverAttributes] {
	return terra.ReferenceAsList[datamonitoractiongroup.WebhookReceiverAttributes](mag.ref.Append("webhook_receiver"))
}

func (mag dataMonitorActionGroupAttributes) Timeouts() datamonitoractiongroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datamonitoractiongroup.TimeoutsAttributes](mag.ref.Append("timeouts"))
}
