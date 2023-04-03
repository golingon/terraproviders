// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGkeBackupBackupPlanIamPolicy creates a new instance of [GkeBackupBackupPlanIamPolicy].
func NewGkeBackupBackupPlanIamPolicy(name string, args GkeBackupBackupPlanIamPolicyArgs) *GkeBackupBackupPlanIamPolicy {
	return &GkeBackupBackupPlanIamPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GkeBackupBackupPlanIamPolicy)(nil)

// GkeBackupBackupPlanIamPolicy represents the Terraform resource google_gke_backup_backup_plan_iam_policy.
type GkeBackupBackupPlanIamPolicy struct {
	Name      string
	Args      GkeBackupBackupPlanIamPolicyArgs
	state     *gkeBackupBackupPlanIamPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) Type() string {
	return "google_gke_backup_backup_plan_iam_policy"
}

// LocalName returns the local name for [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) LocalName() string {
	return gbbpip.Name
}

// Configuration returns the configuration (args) for [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) Configuration() interface{} {
	return gbbpip.Args
}

// DependOn is used for other resources to depend on [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(gbbpip)
}

// Dependencies returns the list of resources [GkeBackupBackupPlanIamPolicy] depends_on.
func (gbbpip *GkeBackupBackupPlanIamPolicy) Dependencies() terra.Dependencies {
	return gbbpip.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) LifecycleManagement() *terra.Lifecycle {
	return gbbpip.Lifecycle
}

// Attributes returns the attributes for [GkeBackupBackupPlanIamPolicy].
func (gbbpip *GkeBackupBackupPlanIamPolicy) Attributes() gkeBackupBackupPlanIamPolicyAttributes {
	return gkeBackupBackupPlanIamPolicyAttributes{ref: terra.ReferenceResource(gbbpip)}
}

// ImportState imports the given attribute values into [GkeBackupBackupPlanIamPolicy]'s state.
func (gbbpip *GkeBackupBackupPlanIamPolicy) ImportState(av io.Reader) error {
	gbbpip.state = &gkeBackupBackupPlanIamPolicyState{}
	if err := json.NewDecoder(av).Decode(gbbpip.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gbbpip.Type(), gbbpip.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GkeBackupBackupPlanIamPolicy] has state.
func (gbbpip *GkeBackupBackupPlanIamPolicy) State() (*gkeBackupBackupPlanIamPolicyState, bool) {
	return gbbpip.state, gbbpip.state != nil
}

// StateMust returns the state for [GkeBackupBackupPlanIamPolicy]. Panics if the state is nil.
func (gbbpip *GkeBackupBackupPlanIamPolicy) StateMust() *gkeBackupBackupPlanIamPolicyState {
	if gbbpip.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gbbpip.Type(), gbbpip.LocalName()))
	}
	return gbbpip.state
}

// GkeBackupBackupPlanIamPolicyArgs contains the configurations for google_gke_backup_backup_plan_iam_policy.
type GkeBackupBackupPlanIamPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyData: string, required
	PolicyData terra.StringValue `hcl:"policy_data,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
}
type gkeBackupBackupPlanIamPolicyAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("etag"))
}

// Id returns a reference to field id of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("id"))
}

// Location returns a reference to field location of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("location"))
}

// Name returns a reference to field name of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("name"))
}

// PolicyData returns a reference to field policy_data of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) PolicyData() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("policy_data"))
}

// Project returns a reference to field project of google_gke_backup_backup_plan_iam_policy.
func (gbbpip gkeBackupBackupPlanIamPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gbbpip.ref.Append("project"))
}

type gkeBackupBackupPlanIamPolicyState struct {
	Etag       string `json:"etag"`
	Id         string `json:"id"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PolicyData string `json:"policy_data"`
	Project    string `json:"project"`
}
