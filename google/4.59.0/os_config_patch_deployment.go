// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	osconfigpatchdeployment "github.com/golingon/terraproviders/google/4.59.0/osconfigpatchdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewOsConfigPatchDeployment(name string, args OsConfigPatchDeploymentArgs) *OsConfigPatchDeployment {
	return &OsConfigPatchDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*OsConfigPatchDeployment)(nil)

type OsConfigPatchDeployment struct {
	Name  string
	Args  OsConfigPatchDeploymentArgs
	state *osConfigPatchDeploymentState
}

func (ocpd *OsConfigPatchDeployment) Type() string {
	return "google_os_config_patch_deployment"
}

func (ocpd *OsConfigPatchDeployment) LocalName() string {
	return ocpd.Name
}

func (ocpd *OsConfigPatchDeployment) Configuration() interface{} {
	return ocpd.Args
}

func (ocpd *OsConfigPatchDeployment) Attributes() osConfigPatchDeploymentAttributes {
	return osConfigPatchDeploymentAttributes{ref: terra.ReferenceResource(ocpd)}
}

func (ocpd *OsConfigPatchDeployment) ImportState(av io.Reader) error {
	ocpd.state = &osConfigPatchDeploymentState{}
	if err := json.NewDecoder(av).Decode(ocpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ocpd.Type(), ocpd.LocalName(), err)
	}
	return nil
}

func (ocpd *OsConfigPatchDeployment) State() (*osConfigPatchDeploymentState, bool) {
	return ocpd.state, ocpd.state != nil
}

func (ocpd *OsConfigPatchDeployment) StateMust() *osConfigPatchDeploymentState {
	if ocpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ocpd.Type(), ocpd.LocalName()))
	}
	return ocpd.state
}

func (ocpd *OsConfigPatchDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(ocpd)
}

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
	// DependsOn contains resources that OsConfigPatchDeployment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type osConfigPatchDeploymentAttributes struct {
	ref terra.Reference
}

func (ocpd osConfigPatchDeploymentAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("create_time"))
}

func (ocpd osConfigPatchDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("description"))
}

func (ocpd osConfigPatchDeploymentAttributes) Duration() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("duration"))
}

func (ocpd osConfigPatchDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("id"))
}

func (ocpd osConfigPatchDeploymentAttributes) LastExecuteTime() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("last_execute_time"))
}

func (ocpd osConfigPatchDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("name"))
}

func (ocpd osConfigPatchDeploymentAttributes) PatchDeploymentId() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("patch_deployment_id"))
}

func (ocpd osConfigPatchDeploymentAttributes) Project() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("project"))
}

func (ocpd osConfigPatchDeploymentAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceString(ocpd.ref.Append("update_time"))
}

func (ocpd osConfigPatchDeploymentAttributes) InstanceFilter() terra.ListValue[osconfigpatchdeployment.InstanceFilterAttributes] {
	return terra.ReferenceList[osconfigpatchdeployment.InstanceFilterAttributes](ocpd.ref.Append("instance_filter"))
}

func (ocpd osConfigPatchDeploymentAttributes) OneTimeSchedule() terra.ListValue[osconfigpatchdeployment.OneTimeScheduleAttributes] {
	return terra.ReferenceList[osconfigpatchdeployment.OneTimeScheduleAttributes](ocpd.ref.Append("one_time_schedule"))
}

func (ocpd osConfigPatchDeploymentAttributes) PatchConfig() terra.ListValue[osconfigpatchdeployment.PatchConfigAttributes] {
	return terra.ReferenceList[osconfigpatchdeployment.PatchConfigAttributes](ocpd.ref.Append("patch_config"))
}

func (ocpd osConfigPatchDeploymentAttributes) RecurringSchedule() terra.ListValue[osconfigpatchdeployment.RecurringScheduleAttributes] {
	return terra.ReferenceList[osconfigpatchdeployment.RecurringScheduleAttributes](ocpd.ref.Append("recurring_schedule"))
}

func (ocpd osConfigPatchDeploymentAttributes) Rollout() terra.ListValue[osconfigpatchdeployment.RolloutAttributes] {
	return terra.ReferenceList[osconfigpatchdeployment.RolloutAttributes](ocpd.ref.Append("rollout"))
}

func (ocpd osConfigPatchDeploymentAttributes) Timeouts() osconfigpatchdeployment.TimeoutsAttributes {
	return terra.ReferenceSingle[osconfigpatchdeployment.TimeoutsAttributes](ocpd.ref.Append("timeouts"))
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
