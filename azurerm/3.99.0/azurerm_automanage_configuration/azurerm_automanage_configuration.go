// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_automanage_configuration

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

// Resource represents the Terraform resource azurerm_automanage_configuration.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermAutomanageConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aac *Resource) Type() string {
	return "azurerm_automanage_configuration"
}

// LocalName returns the local name for [Resource].
func (aac *Resource) LocalName() string {
	return aac.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aac *Resource) Configuration() interface{} {
	return aac.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aac *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aac)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aac *Resource) Dependencies() terra.Dependencies {
	return aac.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aac *Resource) LifecycleManagement() *terra.Lifecycle {
	return aac.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aac *Resource) Attributes() azurermAutomanageConfigurationAttributes {
	return azurermAutomanageConfigurationAttributes{ref: terra.ReferenceResource(aac)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aac *Resource) ImportState(state io.Reader) error {
	aac.state = &azurermAutomanageConfigurationState{}
	if err := json.NewDecoder(state).Decode(aac.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aac.Type(), aac.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aac *Resource) State() (*azurermAutomanageConfigurationState, bool) {
	return aac.state, aac.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aac *Resource) StateMust() *azurermAutomanageConfigurationState {
	if aac.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aac.Type(), aac.LocalName()))
	}
	return aac.state
}

// Args contains the configurations for azurerm_automanage_configuration.
type Args struct {
	// AutomationAccountEnabled: bool, optional
	AutomationAccountEnabled terra.BoolValue `hcl:"automation_account_enabled,attr"`
	// BootDiagnosticsEnabled: bool, optional
	BootDiagnosticsEnabled terra.BoolValue `hcl:"boot_diagnostics_enabled,attr"`
	// DefenderForCloudEnabled: bool, optional
	DefenderForCloudEnabled terra.BoolValue `hcl:"defender_for_cloud_enabled,attr"`
	// GuestConfigurationEnabled: bool, optional
	GuestConfigurationEnabled terra.BoolValue `hcl:"guest_configuration_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// LogAnalyticsEnabled: bool, optional
	LogAnalyticsEnabled terra.BoolValue `hcl:"log_analytics_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StatusChangeAlertEnabled: bool, optional
	StatusChangeAlertEnabled terra.BoolValue `hcl:"status_change_alert_enabled,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Antimalware: optional
	Antimalware *Antimalware `hcl:"antimalware,block"`
	// AzureSecurityBaseline: optional
	AzureSecurityBaseline *AzureSecurityBaseline `hcl:"azure_security_baseline,block"`
	// Backup: optional
	Backup *Backup `hcl:"backup,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermAutomanageConfigurationAttributes struct {
	ref terra.Reference
}

// AutomationAccountEnabled returns a reference to field automation_account_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) AutomationAccountEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("automation_account_enabled"))
}

// BootDiagnosticsEnabled returns a reference to field boot_diagnostics_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) BootDiagnosticsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("boot_diagnostics_enabled"))
}

// DefenderForCloudEnabled returns a reference to field defender_for_cloud_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) DefenderForCloudEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("defender_for_cloud_enabled"))
}

// GuestConfigurationEnabled returns a reference to field guest_configuration_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) GuestConfigurationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("guest_configuration_enabled"))
}

// Id returns a reference to field id of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("location"))
}

// LogAnalyticsEnabled returns a reference to field log_analytics_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) LogAnalyticsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("log_analytics_enabled"))
}

// Name returns a reference to field name of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aac.ref.Append("resource_group_name"))
}

// StatusChangeAlertEnabled returns a reference to field status_change_alert_enabled of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) StatusChangeAlertEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aac.ref.Append("status_change_alert_enabled"))
}

// Tags returns a reference to field tags of azurerm_automanage_configuration.
func (aac azurermAutomanageConfigurationAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aac.ref.Append("tags"))
}

func (aac azurermAutomanageConfigurationAttributes) Antimalware() terra.ListValue[AntimalwareAttributes] {
	return terra.ReferenceAsList[AntimalwareAttributes](aac.ref.Append("antimalware"))
}

func (aac azurermAutomanageConfigurationAttributes) AzureSecurityBaseline() terra.ListValue[AzureSecurityBaselineAttributes] {
	return terra.ReferenceAsList[AzureSecurityBaselineAttributes](aac.ref.Append("azure_security_baseline"))
}

func (aac azurermAutomanageConfigurationAttributes) Backup() terra.ListValue[BackupAttributes] {
	return terra.ReferenceAsList[BackupAttributes](aac.ref.Append("backup"))
}

func (aac azurermAutomanageConfigurationAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aac.ref.Append("timeouts"))
}

type azurermAutomanageConfigurationState struct {
	AutomationAccountEnabled  bool                         `json:"automation_account_enabled"`
	BootDiagnosticsEnabled    bool                         `json:"boot_diagnostics_enabled"`
	DefenderForCloudEnabled   bool                         `json:"defender_for_cloud_enabled"`
	GuestConfigurationEnabled bool                         `json:"guest_configuration_enabled"`
	Id                        string                       `json:"id"`
	Location                  string                       `json:"location"`
	LogAnalyticsEnabled       bool                         `json:"log_analytics_enabled"`
	Name                      string                       `json:"name"`
	ResourceGroupName         string                       `json:"resource_group_name"`
	StatusChangeAlertEnabled  bool                         `json:"status_change_alert_enabled"`
	Tags                      map[string]string            `json:"tags"`
	Antimalware               []AntimalwareState           `json:"antimalware"`
	AzureSecurityBaseline     []AzureSecurityBaselineState `json:"azure_security_baseline"`
	Backup                    []BackupState                `json:"backup"`
	Timeouts                  *TimeoutsState               `json:"timeouts"`
}
