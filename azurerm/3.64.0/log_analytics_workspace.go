// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	loganalyticsworkspace "github.com/golingon/terraproviders/azurerm/3.64.0/loganalyticsworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLogAnalyticsWorkspace creates a new instance of [LogAnalyticsWorkspace].
func NewLogAnalyticsWorkspace(name string, args LogAnalyticsWorkspaceArgs) *LogAnalyticsWorkspace {
	return &LogAnalyticsWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LogAnalyticsWorkspace)(nil)

// LogAnalyticsWorkspace represents the Terraform resource azurerm_log_analytics_workspace.
type LogAnalyticsWorkspace struct {
	Name      string
	Args      LogAnalyticsWorkspaceArgs
	state     *logAnalyticsWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) Type() string {
	return "azurerm_log_analytics_workspace"
}

// LocalName returns the local name for [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) LocalName() string {
	return law.Name
}

// Configuration returns the configuration (args) for [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) Configuration() interface{} {
	return law.Args
}

// DependOn is used for other resources to depend on [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(law)
}

// Dependencies returns the list of resources [LogAnalyticsWorkspace] depends_on.
func (law *LogAnalyticsWorkspace) Dependencies() terra.Dependencies {
	return law.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) LifecycleManagement() *terra.Lifecycle {
	return law.Lifecycle
}

// Attributes returns the attributes for [LogAnalyticsWorkspace].
func (law *LogAnalyticsWorkspace) Attributes() logAnalyticsWorkspaceAttributes {
	return logAnalyticsWorkspaceAttributes{ref: terra.ReferenceResource(law)}
}

// ImportState imports the given attribute values into [LogAnalyticsWorkspace]'s state.
func (law *LogAnalyticsWorkspace) ImportState(av io.Reader) error {
	law.state = &logAnalyticsWorkspaceState{}
	if err := json.NewDecoder(av).Decode(law.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", law.Type(), law.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LogAnalyticsWorkspace] has state.
func (law *LogAnalyticsWorkspace) State() (*logAnalyticsWorkspaceState, bool) {
	return law.state, law.state != nil
}

// StateMust returns the state for [LogAnalyticsWorkspace]. Panics if the state is nil.
func (law *LogAnalyticsWorkspace) StateMust() *logAnalyticsWorkspaceState {
	if law.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", law.Type(), law.LocalName()))
	}
	return law.state
}

// LogAnalyticsWorkspaceArgs contains the configurations for azurerm_log_analytics_workspace.
type LogAnalyticsWorkspaceArgs struct {
	// AllowResourceOnlyPermissions: bool, optional
	AllowResourceOnlyPermissions terra.BoolValue `hcl:"allow_resource_only_permissions,attr"`
	// CmkForQueryForced: bool, optional
	CmkForQueryForced terra.BoolValue `hcl:"cmk_for_query_forced,attr"`
	// DailyQuotaGb: number, optional
	DailyQuotaGb terra.NumberValue `hcl:"daily_quota_gb,attr"`
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
	// ReservationCapacityInGbPerDay: number, optional
	ReservationCapacityInGbPerDay terra.NumberValue `hcl:"reservation_capacity_in_gb_per_day,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *loganalyticsworkspace.Timeouts `hcl:"timeouts,block"`
}
type logAnalyticsWorkspaceAttributes struct {
	ref terra.Reference
}

// AllowResourceOnlyPermissions returns a reference to field allow_resource_only_permissions of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) AllowResourceOnlyPermissions() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("allow_resource_only_permissions"))
}

// CmkForQueryForced returns a reference to field cmk_for_query_forced of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) CmkForQueryForced() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("cmk_for_query_forced"))
}

// DailyQuotaGb returns a reference to field daily_quota_gb of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) DailyQuotaGb() terra.NumberValue {
	return terra.ReferenceAsNumber(law.ref.Append("daily_quota_gb"))
}

// Id returns a reference to field id of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("id"))
}

// InternetIngestionEnabled returns a reference to field internet_ingestion_enabled of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) InternetIngestionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("internet_ingestion_enabled"))
}

// InternetQueryEnabled returns a reference to field internet_query_enabled of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) InternetQueryEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("internet_query_enabled"))
}

// LocalAuthenticationDisabled returns a reference to field local_authentication_disabled of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) LocalAuthenticationDisabled() terra.BoolValue {
	return terra.ReferenceAsBool(law.ref.Append("local_authentication_disabled"))
}

// Location returns a reference to field location of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("name"))
}

// PrimarySharedKey returns a reference to field primary_shared_key of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) PrimarySharedKey() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("primary_shared_key"))
}

// ReservationCapacityInGbPerDay returns a reference to field reservation_capacity_in_gb_per_day of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) ReservationCapacityInGbPerDay() terra.NumberValue {
	return terra.ReferenceAsNumber(law.ref.Append("reservation_capacity_in_gb_per_day"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("resource_group_name"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(law.ref.Append("retention_in_days"))
}

// SecondarySharedKey returns a reference to field secondary_shared_key of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) SecondarySharedKey() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("secondary_shared_key"))
}

// Sku returns a reference to field sku of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](law.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_log_analytics_workspace.
func (law logAnalyticsWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(law.ref.Append("workspace_id"))
}

func (law logAnalyticsWorkspaceAttributes) Timeouts() loganalyticsworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[loganalyticsworkspace.TimeoutsAttributes](law.ref.Append("timeouts"))
}

type logAnalyticsWorkspaceState struct {
	AllowResourceOnlyPermissions  bool                                 `json:"allow_resource_only_permissions"`
	CmkForQueryForced             bool                                 `json:"cmk_for_query_forced"`
	DailyQuotaGb                  float64                              `json:"daily_quota_gb"`
	Id                            string                               `json:"id"`
	InternetIngestionEnabled      bool                                 `json:"internet_ingestion_enabled"`
	InternetQueryEnabled          bool                                 `json:"internet_query_enabled"`
	LocalAuthenticationDisabled   bool                                 `json:"local_authentication_disabled"`
	Location                      string                               `json:"location"`
	Name                          string                               `json:"name"`
	PrimarySharedKey              string                               `json:"primary_shared_key"`
	ReservationCapacityInGbPerDay float64                              `json:"reservation_capacity_in_gb_per_day"`
	ResourceGroupName             string                               `json:"resource_group_name"`
	RetentionInDays               float64                              `json:"retention_in_days"`
	SecondarySharedKey            string                               `json:"secondary_shared_key"`
	Sku                           string                               `json:"sku"`
	Tags                          map[string]string                    `json:"tags"`
	WorkspaceId                   string                               `json:"workspace_id"`
	Timeouts                      *loganalyticsworkspace.TimeoutsState `json:"timeouts"`
}
