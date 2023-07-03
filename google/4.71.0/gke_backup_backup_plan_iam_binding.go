// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	gkebackupbackupplaniambinding "github.com/golingon/terraproviders/google/4.71.0/gkebackupbackupplaniambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeBackupBackupPlanIamBinding creates a new instance of [GkeBackupBackupPlanIamBinding].
func NewGkeBackupBackupPlanIamBinding(name string, args GkeBackupBackupPlanIamBindingArgs) *GkeBackupBackupPlanIamBinding {
	return &GkeBackupBackupPlanIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeBackupBackupPlanIamBinding)(nil)

// GkeBackupBackupPlanIamBinding represents the Terraform resource google_gke_backup_backup_plan_iam_binding.
type GkeBackupBackupPlanIamBinding struct {
	Name      string
	Args      GkeBackupBackupPlanIamBindingArgs
	state     *gkeBackupBackupPlanIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) Type() string {
	return "google_gke_backup_backup_plan_iam_binding"
}

// LocalName returns the local name for [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) LocalName() string {
	return gbbpib.Name
}

// Configuration returns the configuration (args) for [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) Configuration() interface{} {
	return gbbpib.Args
}

// DependOn is used for other resources to depend on [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(gbbpib)
}

// Dependencies returns the list of resources [GkeBackupBackupPlanIamBinding] depends_on.
func (gbbpib *GkeBackupBackupPlanIamBinding) Dependencies() terra.Dependencies {
	return gbbpib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) LifecycleManagement() *terra.Lifecycle {
	return gbbpib.Lifecycle
}

// Attributes returns the attributes for [GkeBackupBackupPlanIamBinding].
func (gbbpib *GkeBackupBackupPlanIamBinding) Attributes() gkeBackupBackupPlanIamBindingAttributes {
	return gkeBackupBackupPlanIamBindingAttributes{ref: terra.ReferenceResource(gbbpib)}
}

// ImportState imports the given attribute values into [GkeBackupBackupPlanIamBinding]'s state.
func (gbbpib *GkeBackupBackupPlanIamBinding) ImportState(av io.Reader) error {
	gbbpib.state = &gkeBackupBackupPlanIamBindingState{}
	if err := json.NewDecoder(av).Decode(gbbpib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbbpib.Type(), gbbpib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeBackupBackupPlanIamBinding] has state.
func (gbbpib *GkeBackupBackupPlanIamBinding) State() (*gkeBackupBackupPlanIamBindingState, bool) {
	return gbbpib.state, gbbpib.state != nil
}

// StateMust returns the state for [GkeBackupBackupPlanIamBinding]. Panics if the state is nil.
func (gbbpib *GkeBackupBackupPlanIamBinding) StateMust() *gkeBackupBackupPlanIamBindingState {
	if gbbpib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbbpib.Type(), gbbpib.LocalName()))
	}
	return gbbpib.state
}

// GkeBackupBackupPlanIamBindingArgs contains the configurations for google_gke_backup_backup_plan_iam_binding.
type GkeBackupBackupPlanIamBindingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *gkebackupbackupplaniambinding.Condition `hcl:"condition,block"`
}
type gkeBackupBackupPlanIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("location"))
}

// Members returns a reference to field members of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gbbpib.ref.Append("members"))
}

// Name returns a reference to field name of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_backup_backup_plan_iam_binding.
func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gbbpib.ref.Append("role"))
}

func (gbbpib gkeBackupBackupPlanIamBindingAttributes) Condition() terra.ListValue[gkebackupbackupplaniambinding.ConditionAttributes] {
	return terra.ReferenceAsList[gkebackupbackupplaniambinding.ConditionAttributes](gbbpib.ref.Append("condition"))
}

type gkeBackupBackupPlanIamBindingState struct {
	Etag      string                                         `json:"etag"`
	Id        string                                         `json:"id"`
	Location  string                                         `json:"location"`
	Members   []string                                       `json:"members"`
	Name      string                                         `json:"name"`
	Project   string                                         `json:"project"`
	Role      string                                         `json:"role"`
	Condition []gkebackupbackupplaniambinding.ConditionState `json:"condition"`
}
