// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_os_config_patch_deployment

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

// Resource represents the Terraform resource google_os_config_patch_deployment.
type Resource struct {
	Name      string
	Args      Args
	state     *googleOsConfigPatchDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gocpd *Resource) Type() string {
	return "google_os_config_patch_deployment"
}

// LocalName returns the local name for [Resource].
func (gocpd *Resource) LocalName() string {
	return gocpd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gocpd *Resource) Configuration() interface{} {
	return gocpd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gocpd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gocpd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gocpd *Resource) Dependencies() terra.Dependencies {
	return gocpd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gocpd *Resource) LifecycleManagement() *terra.Lifecycle {
	return gocpd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gocpd *Resource) Attributes() googleOsConfigPatchDeploymentAttributes {
	return googleOsConfigPatchDeploymentAttributes{ref: terra.ReferenceResource(gocpd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gocpd *Resource) ImportState(state io.Reader) error {
	gocpd.state = &googleOsConfigPatchDeploymentState{}
	if err := json.NewDecoder(state).Decode(gocpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gocpd.Type(), gocpd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gocpd *Resource) State() (*googleOsConfigPatchDeploymentState, bool) {
	return gocpd.state, gocpd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gocpd *Resource) StateMust() *googleOsConfigPatchDeploymentState {
	if gocpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gocpd.Type(), gocpd.LocalName()))
	}
	return gocpd.state
}

// Args contains the configurations for google_os_config_patch_deployment.
type Args struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PatchDeploymentId: string, required
	PatchDeploymentId terra.StringValue `hcl:"patch_deployment_id,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// InstanceFilter: required
	InstanceFilter *InstanceFilter `hcl:"instance_filter,block" validate:"required"`
	// OneTimeSchedule: optional
	OneTimeSchedule *OneTimeSchedule `hcl:"one_time_schedule,block"`
	// PatchConfig: optional
	PatchConfig *PatchConfig `hcl:"patch_config,block"`
	// RecurringSchedule: optional
	RecurringSchedule *RecurringSchedule `hcl:"recurring_schedule,block"`
	// Rollout: optional
	Rollout *Rollout `hcl:"rollout,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleOsConfigPatchDeploymentAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("create_time"))
}

// Description returns a reference to field description of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("description"))
}

// Duration returns a reference to field duration of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("duration"))
}

// Id returns a reference to field id of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("id"))
}

// LastExecuteTime returns a reference to field last_execute_time of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) LastExecuteTime() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("last_execute_time"))
}

// Name returns a reference to field name of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("name"))
}

// PatchDeploymentId returns a reference to field patch_deployment_id of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) PatchDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("patch_deployment_id"))
}

// Project returns a reference to field project of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_os_config_patch_deployment.
func (gocpd googleOsConfigPatchDeploymentAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gocpd.ref.Append("update_time"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) InstanceFilter() terra.ListValue[InstanceFilterAttributes] {
	return terra.ReferenceAsList[InstanceFilterAttributes](gocpd.ref.Append("instance_filter"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) OneTimeSchedule() terra.ListValue[OneTimeScheduleAttributes] {
	return terra.ReferenceAsList[OneTimeScheduleAttributes](gocpd.ref.Append("one_time_schedule"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) PatchConfig() terra.ListValue[PatchConfigAttributes] {
	return terra.ReferenceAsList[PatchConfigAttributes](gocpd.ref.Append("patch_config"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) RecurringSchedule() terra.ListValue[RecurringScheduleAttributes] {
	return terra.ReferenceAsList[RecurringScheduleAttributes](gocpd.ref.Append("recurring_schedule"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) Rollout() terra.ListValue[RolloutAttributes] {
	return terra.ReferenceAsList[RolloutAttributes](gocpd.ref.Append("rollout"))
}

func (gocpd googleOsConfigPatchDeploymentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gocpd.ref.Append("timeouts"))
}

type googleOsConfigPatchDeploymentState struct {
	CreateTime        string                   `json:"create_time"`
	Description       string                   `json:"description"`
	Duration          string                   `json:"duration"`
	Id                string                   `json:"id"`
	LastExecuteTime   string                   `json:"last_execute_time"`
	Name              string                   `json:"name"`
	PatchDeploymentId string                   `json:"patch_deployment_id"`
	Project           string                   `json:"project"`
	UpdateTime        string                   `json:"update_time"`
	InstanceFilter    []InstanceFilterState    `json:"instance_filter"`
	OneTimeSchedule   []OneTimeScheduleState   `json:"one_time_schedule"`
	PatchConfig       []PatchConfigState       `json:"patch_config"`
	RecurringSchedule []RecurringScheduleState `json:"recurring_schedule"`
	Rollout           []RolloutState           `json:"rollout"`
	Timeouts          *TimeoutsState           `json:"timeouts"`
}
