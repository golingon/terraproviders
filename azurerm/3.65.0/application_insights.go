// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	applicationinsights "github.com/golingon/terraproviders/azurerm/3.65.0/applicationinsights"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewApplicationInsights creates a new instance of [ApplicationInsights].
func NewApplicationInsights(name string, args ApplicationInsightsArgs) *ApplicationInsights {
	return &ApplicationInsights{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ApplicationInsights)(nil)

// ApplicationInsights represents the Terraform resource azurerm_application_insights.
type ApplicationInsights struct {
	Name      string
	Args      ApplicationInsightsArgs
	state     *applicationInsightsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ApplicationInsights].
func (ai *ApplicationInsights) Type() string {
	return "azurerm_application_insights"
}

// LocalName returns the local name for [ApplicationInsights].
func (ai *ApplicationInsights) LocalName() string {
	return ai.Name
}

// Configuration returns the configuration (args) for [ApplicationInsights].
func (ai *ApplicationInsights) Configuration() interface{} {
	return ai.Args
}

// DependOn is used for other resources to depend on [ApplicationInsights].
func (ai *ApplicationInsights) DependOn() terra.Reference {
	return terra.ReferenceResource(ai)
}

// Dependencies returns the list of resources [ApplicationInsights] depends_on.
func (ai *ApplicationInsights) Dependencies() terra.Dependencies {
	return ai.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ApplicationInsights].
func (ai *ApplicationInsights) LifecycleManagement() *terra.Lifecycle {
	return ai.Lifecycle
}

// Attributes returns the attributes for [ApplicationInsights].
func (ai *ApplicationInsights) Attributes() applicationInsightsAttributes {
	return applicationInsightsAttributes{ref: terra.ReferenceResource(ai)}
}

// ImportState imports the given attribute values into [ApplicationInsights]'s state.
func (ai *ApplicationInsights) ImportState(av io.Reader) error {
	ai.state = &applicationInsightsState{}
	if err := json.NewDecoder(av).Decode(ai.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ai.Type(), ai.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ApplicationInsights] has state.
func (ai *ApplicationInsights) State() (*applicationInsightsState, bool) {
	return ai.state, ai.state != nil
}

// StateMust returns the state for [ApplicationInsights]. Panics if the state is nil.
func (ai *ApplicationInsights) StateMust() *applicationInsightsState {
	if ai.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ai.Type(), ai.LocalName()))
	}
	return ai.state
}

// ApplicationInsightsArgs contains the configurations for azurerm_application_insights.
type ApplicationInsightsArgs struct {
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
	Timeouts *applicationinsights.Timeouts `hcl:"timeouts,block"`
}
type applicationInsightsAttributes struct {
	ref terra.Reference
}

// AppId returns a reference to field app_id of azurerm_application_insights.
func (ai applicationInsightsAttributes) AppId() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("app_id"))
}

// ApplicationType returns a reference to field application_type of azurerm_application_insights.
func (ai applicationInsightsAttributes) ApplicationType() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("application_type"))
}

// ConnectionString returns a reference to field connection_string of azurerm_application_insights.
func (ai applicationInsightsAttributes) ConnectionString() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("connection_string"))
}

// DailyDataCapInGb returns a reference to field daily_data_cap_in_gb of azurerm_application_insights.
func (ai applicationInsightsAttributes) DailyDataCapInGb() terra.NumberValue {
	return terra.ReferenceAsNumber(ai.ref.Append("daily_data_cap_in_gb"))
}

// DailyDataCapNotificationsDisabled returns a reference to field daily_data_cap_notifications_disabled of azurerm_application_insights.
func (ai applicationInsightsAttributes) DailyDataCapNotificationsDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("daily_data_cap_notifications_disabled"))
}

// DisableIpMasking returns a reference to field disable_ip_masking of azurerm_application_insights.
func (ai applicationInsightsAttributes) DisableIpMasking() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("disable_ip_masking"))
}

// ForceCustomerStorageForProfiler returns a reference to field force_customer_storage_for_profiler of azurerm_application_insights.
func (ai applicationInsightsAttributes) ForceCustomerStorageForProfiler() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("force_customer_storage_for_profiler"))
}

// Id returns a reference to field id of azurerm_application_insights.
func (ai applicationInsightsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("id"))
}

// InstrumentationKey returns a reference to field instrumentation_key of azurerm_application_insights.
func (ai applicationInsightsAttributes) InstrumentationKey() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("instrumentation_key"))
}

// InternetIngestionEnabled returns a reference to field internet_ingestion_enabled of azurerm_application_insights.
func (ai applicationInsightsAttributes) InternetIngestionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("internet_ingestion_enabled"))
}

// InternetQueryEnabled returns a reference to field internet_query_enabled of azurerm_application_insights.
func (ai applicationInsightsAttributes) InternetQueryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("internet_query_enabled"))
}

// LocalAuthenticationDisabled returns a reference to field local_authentication_disabled of azurerm_application_insights.
func (ai applicationInsightsAttributes) LocalAuthenticationDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(ai.ref.Append("local_authentication_disabled"))
}

// Location returns a reference to field location of azurerm_application_insights.
func (ai applicationInsightsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_insights.
func (ai applicationInsightsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_insights.
func (ai applicationInsightsAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("resource_group_name"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_application_insights.
func (ai applicationInsightsAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(ai.ref.Append("retention_in_days"))
}

// SamplingPercentage returns a reference to field sampling_percentage of azurerm_application_insights.
func (ai applicationInsightsAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(ai.ref.Append("sampling_percentage"))
}

// Tags returns a reference to field tags of azurerm_application_insights.
func (ai applicationInsightsAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ai.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_application_insights.
func (ai applicationInsightsAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(ai.ref.Append("workspace_id"))
}

func (ai applicationInsightsAttributes) Timeouts() applicationinsights.TimeoutsAttributes {
	return terra.ReferenceAsSingle[applicationinsights.TimeoutsAttributes](ai.ref.Append("timeouts"))
}

type applicationInsightsState struct {
	AppId                             string                             `json:"app_id"`
	ApplicationType                   string                             `json:"application_type"`
	ConnectionString                  string                             `json:"connection_string"`
	DailyDataCapInGb                  float64                            `json:"daily_data_cap_in_gb"`
	DailyDataCapNotificationsDisabled bool                               `json:"daily_data_cap_notifications_disabled"`
	DisableIpMasking                  bool                               `json:"disable_ip_masking"`
	ForceCustomerStorageForProfiler   bool                               `json:"force_customer_storage_for_profiler"`
	Id                                string                             `json:"id"`
	InstrumentationKey                string                             `json:"instrumentation_key"`
	InternetIngestionEnabled          bool                               `json:"internet_ingestion_enabled"`
	InternetQueryEnabled              bool                               `json:"internet_query_enabled"`
	LocalAuthenticationDisabled       bool                               `json:"local_authentication_disabled"`
	Location                          string                             `json:"location"`
	Name                              string                             `json:"name"`
	ResourceGroupName                 string                             `json:"resource_group_name"`
	RetentionInDays                   float64                            `json:"retention_in_days"`
	SamplingPercentage                float64                            `json:"sampling_percentage"`
	Tags                              map[string]string                  `json:"tags"`
	WorkspaceId                       string                             `json:"workspace_id"`
	Timeouts                          *applicationinsights.TimeoutsState `json:"timeouts"`
}
