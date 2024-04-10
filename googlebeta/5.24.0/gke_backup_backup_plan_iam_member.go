// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	gkebackupbackupplaniammember "github.com/golingon/terraproviders/googlebeta/5.24.0/gkebackupbackupplaniammember"
	"io"
)

// NewGkeBackupBackupPlanIamMember creates a new instance of [GkeBackupBackupPlanIamMember].
func NewGkeBackupBackupPlanIamMember(name string, args GkeBackupBackupPlanIamMemberArgs) *GkeBackupBackupPlanIamMember {
	return &GkeBackupBackupPlanIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeBackupBackupPlanIamMember)(nil)

// GkeBackupBackupPlanIamMember represents the Terraform resource google_gke_backup_backup_plan_iam_member.
type GkeBackupBackupPlanIamMember struct {
	Name      string
	Args      GkeBackupBackupPlanIamMemberArgs
	state     *gkeBackupBackupPlanIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) Type() string {
	return "google_gke_backup_backup_plan_iam_member"
}

// LocalName returns the local name for [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) LocalName() string {
	return gbbpim.Name
}

// Configuration returns the configuration (args) for [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) Configuration() interface{} {
	return gbbpim.Args
}

// DependOn is used for other resources to depend on [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(gbbpim)
}

// Dependencies returns the list of resources [GkeBackupBackupPlanIamMember] depends_on.
func (gbbpim *GkeBackupBackupPlanIamMember) Dependencies() terra.Dependencies {
	return gbbpim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) LifecycleManagement() *terra.Lifecycle {
	return gbbpim.Lifecycle
}

// Attributes returns the attributes for [GkeBackupBackupPlanIamMember].
func (gbbpim *GkeBackupBackupPlanIamMember) Attributes() gkeBackupBackupPlanIamMemberAttributes {
	return gkeBackupBackupPlanIamMemberAttributes{ref: terra.ReferenceResource(gbbpim)}
}

// ImportState imports the given attribute values into [GkeBackupBackupPlanIamMember]'s state.
func (gbbpim *GkeBackupBackupPlanIamMember) ImportState(av io.Reader) error {
	gbbpim.state = &gkeBackupBackupPlanIamMemberState{}
	if err := json.NewDecoder(av).Decode(gbbpim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbbpim.Type(), gbbpim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeBackupBackupPlanIamMember] has state.
func (gbbpim *GkeBackupBackupPlanIamMember) State() (*gkeBackupBackupPlanIamMemberState, bool) {
	return gbbpim.state, gbbpim.state != nil
}

// StateMust returns the state for [GkeBackupBackupPlanIamMember]. Panics if the state is nil.
func (gbbpim *GkeBackupBackupPlanIamMember) StateMust() *gkeBackupBackupPlanIamMemberState {
	if gbbpim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbbpim.Type(), gbbpim.LocalName()))
	}
	return gbbpim.state
}

// GkeBackupBackupPlanIamMemberArgs contains the configurations for google_gke_backup_backup_plan_iam_member.
type GkeBackupBackupPlanIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *gkebackupbackupplaniammember.Condition `hcl:"condition,block"`
}
type gkeBackupBackupPlanIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("location"))
}

// Member returns a reference to field member of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("member"))
}

// Name returns a reference to field name of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("name"))
}

// Project returns a reference to field project of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("project"))
}

// Role returns a reference to field role of google_gke_backup_backup_plan_iam_member.
func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(gbbpim.ref.Append("role"))
}

func (gbbpim gkeBackupBackupPlanIamMemberAttributes) Condition() terra.ListValue[gkebackupbackupplaniammember.ConditionAttributes] {
	return terra.ReferenceAsList[gkebackupbackupplaniammember.ConditionAttributes](gbbpim.ref.Append("condition"))
}

type gkeBackupBackupPlanIamMemberState struct {
	Etag      string                                        `json:"etag"`
	Id        string                                        `json:"id"`
	Location  string                                        `json:"location"`
	Member    string                                        `json:"member"`
	Name      string                                        `json:"name"`
	Project   string                                        `json:"project"`
	Role      string                                        `json:"role"`
	Condition []gkebackupbackupplaniammember.ConditionState `json:"condition"`
}
