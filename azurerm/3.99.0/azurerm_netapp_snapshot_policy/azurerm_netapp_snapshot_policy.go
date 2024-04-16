// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_netapp_snapshot_policy

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

// Resource represents the Terraform resource azurerm_netapp_snapshot_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermNetappSnapshotPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ansp *Resource) Type() string {
	return "azurerm_netapp_snapshot_policy"
}

// LocalName returns the local name for [Resource].
func (ansp *Resource) LocalName() string {
	return ansp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ansp *Resource) Configuration() interface{} {
	return ansp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ansp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ansp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ansp *Resource) Dependencies() terra.Dependencies {
	return ansp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ansp *Resource) LifecycleManagement() *terra.Lifecycle {
	return ansp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ansp *Resource) Attributes() azurermNetappSnapshotPolicyAttributes {
	return azurermNetappSnapshotPolicyAttributes{ref: terra.ReferenceResource(ansp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ansp *Resource) ImportState(state io.Reader) error {
	ansp.state = &azurermNetappSnapshotPolicyState{}
	if err := json.NewDecoder(state).Decode(ansp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ansp.Type(), ansp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ansp *Resource) State() (*azurermNetappSnapshotPolicyState, bool) {
	return ansp.state, ansp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ansp *Resource) StateMust() *azurermNetappSnapshotPolicyState {
	if ansp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ansp.Type(), ansp.LocalName()))
	}
	return ansp.state
}

// Args contains the configurations for azurerm_netapp_snapshot_policy.
type Args struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// DailySchedule: optional
	DailySchedule *DailySchedule `hcl:"daily_schedule,block"`
	// HourlySchedule: optional
	HourlySchedule *HourlySchedule `hcl:"hourly_schedule,block"`
	// MonthlySchedule: optional
	MonthlySchedule *MonthlySchedule `hcl:"monthly_schedule,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// WeeklySchedule: optional
	WeeklySchedule *WeeklySchedule `hcl:"weekly_schedule,block"`
}

type azurermNetappSnapshotPolicyAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(ansp.ref.Append("account_name"))
}

// Enabled returns a reference to field enabled of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ansp.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ansp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ansp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ansp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ansp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_netapp_snapshot_policy.
func (ansp azurermNetappSnapshotPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ansp.ref.Append("tags"))
}

func (ansp azurermNetappSnapshotPolicyAttributes) DailySchedule() terra.ListValue[DailyScheduleAttributes] {
	return terra.ReferenceAsList[DailyScheduleAttributes](ansp.ref.Append("daily_schedule"))
}

func (ansp azurermNetappSnapshotPolicyAttributes) HourlySchedule() terra.ListValue[HourlyScheduleAttributes] {
	return terra.ReferenceAsList[HourlyScheduleAttributes](ansp.ref.Append("hourly_schedule"))
}

func (ansp azurermNetappSnapshotPolicyAttributes) MonthlySchedule() terra.ListValue[MonthlyScheduleAttributes] {
	return terra.ReferenceAsList[MonthlyScheduleAttributes](ansp.ref.Append("monthly_schedule"))
}

func (ansp azurermNetappSnapshotPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ansp.ref.Append("timeouts"))
}

func (ansp azurermNetappSnapshotPolicyAttributes) WeeklySchedule() terra.ListValue[WeeklyScheduleAttributes] {
	return terra.ReferenceAsList[WeeklyScheduleAttributes](ansp.ref.Append("weekly_schedule"))
}

type azurermNetappSnapshotPolicyState struct {
	AccountName       string                 `json:"account_name"`
	Enabled           bool                   `json:"enabled"`
	Id                string                 `json:"id"`
	Location          string                 `json:"location"`
	Name              string                 `json:"name"`
	ResourceGroupName string                 `json:"resource_group_name"`
	Tags              map[string]string      `json:"tags"`
	DailySchedule     []DailyScheduleState   `json:"daily_schedule"`
	HourlySchedule    []HourlyScheduleState  `json:"hourly_schedule"`
	MonthlySchedule   []MonthlyScheduleState `json:"monthly_schedule"`
	Timeouts          *TimeoutsState         `json:"timeouts"`
	WeeklySchedule    []WeeklyScheduleState  `json:"weekly_schedule"`
}
