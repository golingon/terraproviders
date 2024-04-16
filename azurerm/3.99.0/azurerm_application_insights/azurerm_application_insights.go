// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_application_insights

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

// Resource represents the Terraform resource azurerm_application_insights.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermApplicationInsightsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aai *Resource) Type() string {
	return "azurerm_application_insights"
}

// LocalName returns the local name for [Resource].
func (aai *Resource) LocalName() string {
	return aai.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aai *Resource) Configuration() interface{} {
	return aai.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aai *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aai)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aai *Resource) Dependencies() terra.Dependencies {
	return aai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aai *Resource) LifecycleManagement() *terra.Lifecycle {
	return aai.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aai *Resource) Attributes() azurermApplicationInsightsAttributes {
	return azurermApplicationInsightsAttributes{ref: terra.ReferenceResource(aai)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aai *Resource) ImportState(state io.Reader) error {
	aai.state = &azurermApplicationInsightsState{}
	if err := json.NewDecoder(state).Decode(aai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aai.Type(), aai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aai *Resource) State() (*azurermApplicationInsightsState, bool) {
	return aai.state, aai.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aai *Resource) StateMust() *azurermApplicationInsightsState {
	if aai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aai.Type(), aai.LocalName()))
	}
	return aai.state
}

// Args contains the configurations for azurerm_application_insights.
type Args struct {
	// ApplicationType: string, required
	ApplicationType terra.StringValue `hcl:"application_type,attr" validate:"required"`
	// DailyDataCapInGb: number, optional
	DailyDataCapInGb terra.NumberValue `hcl:"daily_data_cap_in_gb,attr"`
	// DailyDataCapNotificationsDisabled: bool, optional
	DailyDataCapNotificationsDisabled terra.BoolValue `hcl:"daily_data_cap_notifications_disabled,attr"`
	// DisableIpMasking: bool, optional
	DisableIpMasking terra.BoolValue `hcl:"disable_ip_masking,attr"`
	// ForceCustomerStorageForProfiler: bool, optional
	ForceCustomerStorageForProfiler terra.BoolValue `hcl:"force_customer_storage_for_profiler,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetIngestionEnabled: bool, optional
	InternetIngestionEnabled terra.BoolValue `hcl:"internet_ingestion_enabled,attr"`
	// InternetQueryEnabled: bool, optional
	InternetQueryEnabled terra.BoolValue `hcl:"internet_query_enabled,attr"`
	// LocalAuthenticationDisabled: bool, optional
	LocalAuthenticationDisabled terra.BoolValue `hcl:"local_authentication_disabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// SamplingPercentage: number, optional
	SamplingPercentage terra.NumberValue `hcl:"sampling_percentage,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// WorkspaceId: string, optional
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermApplicationInsightsAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("app_id"))
}

// ApplicationType returns a reference to field application_type of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) ApplicationType() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("application_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("connection_string"))
}

// DailyDataCapInGb returns a reference to field daily_data_cap_in_gb of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) DailyDataCapInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(aai.ref.Append("daily_data_cap_in_gb"))
}

// DailyDataCapNotificationsDisabled returns a reference to field daily_data_cap_notifications_disabled of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) DailyDataCapNotificationsDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("daily_data_cap_notifications_disabled"))
}

// DisableIpMasking returns a reference to field disable_ip_masking of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) DisableIpMasking() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("disable_ip_masking"))
}

// ForceCustomerStorageForProfiler returns a reference to field force_customer_storage_for_profiler of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) ForceCustomerStorageForProfiler() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("force_customer_storage_for_profiler"))
}

// Id returns a reference to field id of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("id"))
}

// InstrumentationKey returns a reference to field instrumentation_key of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) InstrumentationKey() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("instrumentation_key"))
}

// InternetIngestionEnabled returns a reference to field internet_ingestion_enabled of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) InternetIngestionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("internet_ingestion_enabled"))
}

// InternetQueryEnabled returns a reference to field internet_query_enabled of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) InternetQueryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("internet_query_enabled"))
}

// LocalAuthenticationDisabled returns a reference to field local_authentication_disabled of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) LocalAuthenticationDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(aai.ref.Append("local_authentication_disabled"))
}

// Location returns a reference to field location of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("resource_group_name"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(aai.ref.Append("retention_in_days"))
}

// SamplingPercentage returns a reference to field sampling_percentage of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(aai.ref.Append("sampling_percentage"))
}

// Tags returns a reference to field tags of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](aai.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_application_insights.
func (aai azurermApplicationInsightsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(aai.ref.Append("workspace_id"))
}

func (aai azurermApplicationInsightsAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aai.ref.Append("timeouts"))
}

type azurermApplicationInsightsState struct {
	AppId                             string            `json:"app_id"`
	ApplicationType                   string            `json:"application_type"`
	ConnectionString                  string            `json:"connection_string"`
	DailyDataCapInGb                  float64           `json:"daily_data_cap_in_gb"`
	DailyDataCapNotificationsDisabled bool              `json:"daily_data_cap_notifications_disabled"`
	DisableIpMasking                  bool              `json:"disable_ip_masking"`
	ForceCustomerStorageForProfiler   bool              `json:"force_customer_storage_for_profiler"`
	Id                                string            `json:"id"`
	InstrumentationKey                string            `json:"instrumentation_key"`
	InternetIngestionEnabled          bool              `json:"internet_ingestion_enabled"`
	InternetQueryEnabled              bool              `json:"internet_query_enabled"`
	LocalAuthenticationDisabled       bool              `json:"local_authentication_disabled"`
	Location                          string            `json:"location"`
	Name                              string            `json:"name"`
	ResourceGroupName                 string            `json:"resource_group_name"`
	RetentionInDays                   float64           `json:"retention_in_days"`
	SamplingPercentage                float64           `json:"sampling_percentage"`
	Tags                              map[string]string `json:"tags"`
	WorkspaceId                       string            `json:"workspace_id"`
	Timeouts                          *TimeoutsState    `json:"timeouts"`
}
