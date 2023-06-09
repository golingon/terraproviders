// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	gkebackupbackupplan "github.com/golingon/terraproviders/googlebeta/4.66.0/gkebackupbackupplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeBackupBackupPlan creates a new instance of [GkeBackupBackupPlan].
func NewGkeBackupBackupPlan(name string, args GkeBackupBackupPlanArgs) *GkeBackupBackupPlan {
	return &GkeBackupBackupPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeBackupBackupPlan)(nil)

// GkeBackupBackupPlan represents the Terraform resource google_gke_backup_backup_plan.
type GkeBackupBackupPlan struct {
	Name      string
	Args      GkeBackupBackupPlanArgs
	state     *gkeBackupBackupPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) Type() string {
	return "google_gke_backup_backup_plan"
}

// LocalName returns the local name for [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) LocalName() string {
	return gbbp.Name
}

// Configuration returns the configuration (args) for [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) Configuration() interface{} {
	return gbbp.Args
}

// DependOn is used for other resources to depend on [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(gbbp)
}

// Dependencies returns the list of resources [GkeBackupBackupPlan] depends_on.
func (gbbp *GkeBackupBackupPlan) Dependencies() terra.Dependencies {
	return gbbp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) LifecycleManagement() *terra.Lifecycle {
	return gbbp.Lifecycle
}

// Attributes returns the attributes for [GkeBackupBackupPlan].
func (gbbp *GkeBackupBackupPlan) Attributes() gkeBackupBackupPlanAttributes {
	return gkeBackupBackupPlanAttributes{ref: terra.ReferenceResource(gbbp)}
}

// ImportState imports the given attribute values into [GkeBackupBackupPlan]'s state.
func (gbbp *GkeBackupBackupPlan) ImportState(av io.Reader) error {
	gbbp.state = &gkeBackupBackupPlanState{}
	if err := json.NewDecoder(av).Decode(gbbp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbbp.Type(), gbbp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeBackupBackupPlan] has state.
func (gbbp *GkeBackupBackupPlan) State() (*gkeBackupBackupPlanState, bool) {
	return gbbp.state, gbbp.state != nil
}

// StateMust returns the state for [GkeBackupBackupPlan]. Panics if the state is nil.
func (gbbp *GkeBackupBackupPlan) StateMust() *gkeBackupBackupPlanState {
	if gbbp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbbp.Type(), gbbp.LocalName()))
	}
	return gbbp.state
}

// GkeBackupBackupPlanArgs contains the configurations for google_gke_backup_backup_plan.
type GkeBackupBackupPlanArgs struct {
	// Cluster: string, required
	Cluster terra.StringValue `hcl:"cluster,attr" validate:"required"`
	// Deactivated: bool, optional
	Deactivated terra.BoolValue `hcl:"deactivated,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BackupConfig: optional
	BackupConfig *gkebackupbackupplan.BackupConfig `hcl:"backup_config,block"`
	// BackupSchedule: optional
	BackupSchedule *gkebackupbackupplan.BackupSchedule `hcl:"backup_schedule,block"`
	// RetentionPolicy: optional
	RetentionPolicy *gkebackupbackupplan.RetentionPolicy `hcl:"retention_policy,block"`
	// Timeouts: optional
	Timeouts *gkebackupbackupplan.Timeouts `hcl:"timeouts,block"`
}
type gkeBackupBackupPlanAttributes struct {
	ref terra.Reference
}

// Cluster returns a reference to field cluster of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Cluster() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("cluster"))
}

// Deactivated returns a reference to field deactivated of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Deactivated() terra.BoolValue {
	return terra.ReferenceAsBool(gbbp.ref.Append("deactivated"))
}

// Description returns a reference to field description of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("description"))
}

// Etag returns a reference to field etag of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("id"))
}

// Labels returns a reference to field labels of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gbbp.ref.Append("labels"))
}

// Location returns a reference to field location of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("location"))
}

// Name returns a reference to field name of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("project"))
}

// ProtectedPodCount returns a reference to field protected_pod_count of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) ProtectedPodCount() terra.NumberValue {
	return terra.ReferenceAsNumber(gbbp.ref.Append("protected_pod_count"))
}

// Uid returns a reference to field uid of google_gke_backup_backup_plan.
func (gbbp gkeBackupBackupPlanAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gbbp.ref.Append("uid"))
}

func (gbbp gkeBackupBackupPlanAttributes) BackupConfig() terra.ListValue[gkebackupbackupplan.BackupConfigAttributes] {
	return terra.ReferenceAsList[gkebackupbackupplan.BackupConfigAttributes](gbbp.ref.Append("backup_config"))
}

func (gbbp gkeBackupBackupPlanAttributes) BackupSchedule() terra.ListValue[gkebackupbackupplan.BackupScheduleAttributes] {
	return terra.ReferenceAsList[gkebackupbackupplan.BackupScheduleAttributes](gbbp.ref.Append("backup_schedule"))
}

func (gbbp gkeBackupBackupPlanAttributes) RetentionPolicy() terra.ListValue[gkebackupbackupplan.RetentionPolicyAttributes] {
	return terra.ReferenceAsList[gkebackupbackupplan.RetentionPolicyAttributes](gbbp.ref.Append("retention_policy"))
}

func (gbbp gkeBackupBackupPlanAttributes) Timeouts() gkebackupbackupplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[gkebackupbackupplan.TimeoutsAttributes](gbbp.ref.Append("timeouts"))
}

type gkeBackupBackupPlanState struct {
	Cluster           string                                     `json:"cluster"`
	Deactivated       bool                                       `json:"deactivated"`
	Description       string                                     `json:"description"`
	Etag              string                                     `json:"etag"`
	Id                string                                     `json:"id"`
	Labels            map[string]string                          `json:"labels"`
	Location          string                                     `json:"location"`
	Name              string                                     `json:"name"`
	Project           string                                     `json:"project"`
	ProtectedPodCount float64                                    `json:"protected_pod_count"`
	Uid               string                                     `json:"uid"`
	BackupConfig      []gkebackupbackupplan.BackupConfigState    `json:"backup_config"`
	BackupSchedule    []gkebackupbackupplan.BackupScheduleState  `json:"backup_schedule"`
	RetentionPolicy   []gkebackupbackupplan.RetentionPolicyState `json:"retention_policy"`
	Timeouts          *gkebackupbackupplan.TimeoutsState         `json:"timeouts"`
}
