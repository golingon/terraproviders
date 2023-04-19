// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	osconfigpatchdeployment "github.com/golingon/terraproviders/googlebeta/4.62.0/osconfigpatchdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewOsConfigPatchDeployment creates a new instance of [OsConfigPatchDeployment].
func NewOsConfigPatchDeployment(name string, args OsConfigPatchDeploymentArgs) *OsConfigPatchDeployment {
	return &OsConfigPatchDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OsConfigPatchDeployment)(nil)

// OsConfigPatchDeployment represents the Terraform resource google_os_config_patch_deployment.
type OsConfigPatchDeployment struct {
	Name      string
	Args      OsConfigPatchDeploymentArgs
	state     *osConfigPatchDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) Type() string {
	return "google_os_config_patch_deployment"
}

// LocalName returns the local name for [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) LocalName() string {
	return ocpd.Name
}

// Configuration returns the configuration (args) for [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) Configuration() interface{} {
	return ocpd.Args
}

// DependOn is used for other resources to depend on [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(ocpd)
}

// Dependencies returns the list of resources [OsConfigPatchDeployment] depends_on.
func (ocpd *OsConfigPatchDeployment) Dependencies() terra.Dependencies {
	return ocpd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) LifecycleManagement() *terra.Lifecycle {
	return ocpd.Lifecycle
}

// Attributes returns the attributes for [OsConfigPatchDeployment].
func (ocpd *OsConfigPatchDeployment) Attributes() osConfigPatchDeploymentAttributes {
	return osConfigPatchDeploymentAttributes{ref: terra.ReferenceResource(ocpd)}
}

// ImportState imports the given attribute values into [OsConfigPatchDeployment]'s state.
func (ocpd *OsConfigPatchDeployment) ImportState(av io.Reader) error {
	ocpd.state = &osConfigPatchDeploymentState{}
	if err := json.NewDecoder(av).Decode(ocpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ocpd.Type(), ocpd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [OsConfigPatchDeployment] has state.
func (ocpd *OsConfigPatchDeployment) State() (*osConfigPatchDeploymentState, bool) {
	return ocpd.state, ocpd.state != nil
}

// StateMust returns the state for [OsConfigPatchDeployment]. Panics if the state is nil.
func (ocpd *OsConfigPatchDeployment) StateMust() *osConfigPatchDeploymentState {
	if ocpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ocpd.Type(), ocpd.LocalName()))
	}
	return ocpd.state
}

// OsConfigPatchDeploymentArgs contains the configurations for google_os_config_patch_deployment.
type OsConfigPatchDeploymentArgs struct {
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
	InstanceFilter *osconfigpatchdeployment.InstanceFilter `hcl:"instance_filter,block" validate:"required"`
	// OneTimeSchedule: optional
	OneTimeSchedule *osconfigpatchdeployment.OneTimeSchedule `hcl:"one_time_schedule,block"`
	// PatchConfig: optional
	PatchConfig *osconfigpatchdeployment.PatchConfig `hcl:"patch_config,block"`
	// RecurringSchedule: optional
	RecurringSchedule *osconfigpatchdeployment.RecurringSchedule `hcl:"recurring_schedule,block"`
	// Rollout: optional
	Rollout *osconfigpatchdeployment.Rollout `hcl:"rollout,block"`
	// Timeouts: optional
	Timeouts *osconfigpatchdeployment.Timeouts `hcl:"timeouts,block"`
}
type osConfigPatchDeploymentAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("create_time"))
}

// Description returns a reference to field description of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("description"))
}

// Duration returns a reference to field duration of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("duration"))
}

// Id returns a reference to field id of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("id"))
}

// LastExecuteTime returns a reference to field last_execute_time of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) LastExecuteTime() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("last_execute_time"))
}

// Name returns a reference to field name of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("name"))
}

// PatchDeploymentId returns a reference to field patch_deployment_id of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) PatchDeploymentId() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("patch_deployment_id"))
}

// Project returns a reference to field project of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_os_config_patch_deployment.
func (ocpd osConfigPatchDeploymentAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ocpd.ref.Append("update_time"))
}

func (ocpd osConfigPatchDeploymentAttributes) InstanceFilter() terra.ListValue[osconfigpatchdeployment.InstanceFilterAttributes] {
	return terra.ReferenceAsList[osconfigpatchdeployment.InstanceFilterAttributes](ocpd.ref.Append("instance_filter"))
}

func (ocpd osConfigPatchDeploymentAttributes) OneTimeSchedule() terra.ListValue[osconfigpatchdeployment.OneTimeScheduleAttributes] {
	return terra.ReferenceAsList[osconfigpatchdeployment.OneTimeScheduleAttributes](ocpd.ref.Append("one_time_schedule"))
}

func (ocpd osConfigPatchDeploymentAttributes) PatchConfig() terra.ListValue[osconfigpatchdeployment.PatchConfigAttributes] {
	return terra.ReferenceAsList[osconfigpatchdeployment.PatchConfigAttributes](ocpd.ref.Append("patch_config"))
}

func (ocpd osConfigPatchDeploymentAttributes) RecurringSchedule() terra.ListValue[osconfigpatchdeployment.RecurringScheduleAttributes] {
	return terra.ReferenceAsList[osconfigpatchdeployment.RecurringScheduleAttributes](ocpd.ref.Append("recurring_schedule"))
}

func (ocpd osConfigPatchDeploymentAttributes) Rollout() terra.ListValue[osconfigpatchdeployment.RolloutAttributes] {
	return terra.ReferenceAsList[osconfigpatchdeployment.RolloutAttributes](ocpd.ref.Append("rollout"))
}

func (ocpd osConfigPatchDeploymentAttributes) Timeouts() osconfigpatchdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[osconfigpatchdeployment.TimeoutsAttributes](ocpd.ref.Append("timeouts"))
}

type osConfigPatchDeploymentState struct {
	CreateTime        string                                           `json:"create_time"`
	Description       string                                           `json:"description"`
	Duration          string                                           `json:"duration"`
	Id                string                                           `json:"id"`
	LastExecuteTime   string                                           `json:"last_execute_time"`
	Name              string                                           `json:"name"`
	PatchDeploymentId string                                           `json:"patch_deployment_id"`
	Project           string                                           `json:"project"`
	UpdateTime        string                                           `json:"update_time"`
	InstanceFilter    []osconfigpatchdeployment.InstanceFilterState    `json:"instance_filter"`
	OneTimeSchedule   []osconfigpatchdeployment.OneTimeScheduleState   `json:"one_time_schedule"`
	PatchConfig       []osconfigpatchdeployment.PatchConfigState       `json:"patch_config"`
	RecurringSchedule []osconfigpatchdeployment.RecurringScheduleState `json:"recurring_schedule"`
	Rollout           []osconfigpatchdeployment.RolloutState           `json:"rollout"`
	Timeouts          *osconfigpatchdeployment.TimeoutsState           `json:"timeouts"`
}
