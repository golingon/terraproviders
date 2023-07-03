// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	projectorganizationpolicy "github.com/golingon/terraproviders/google/4.71.0/projectorganizationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectOrganizationPolicy creates a new instance of [ProjectOrganizationPolicy].
func NewProjectOrganizationPolicy(name string, args ProjectOrganizationPolicyArgs) *ProjectOrganizationPolicy {
	return &ProjectOrganizationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectOrganizationPolicy)(nil)

// ProjectOrganizationPolicy represents the Terraform resource google_project_organization_policy.
type ProjectOrganizationPolicy struct {
	Name      string
	Args      ProjectOrganizationPolicyArgs
	state     *projectOrganizationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) Type() string {
	return "google_project_organization_policy"
}

// LocalName returns the local name for [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) LocalName() string {
	return pop.Name
}

// Configuration returns the configuration (args) for [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) Configuration() interface{} {
	return pop.Args
}

// DependOn is used for other resources to depend on [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(pop)
}

// Dependencies returns the list of resources [ProjectOrganizationPolicy] depends_on.
func (pop *ProjectOrganizationPolicy) Dependencies() terra.Dependencies {
	return pop.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) LifecycleManagement() *terra.Lifecycle {
	return pop.Lifecycle
}

// Attributes returns the attributes for [ProjectOrganizationPolicy].
func (pop *ProjectOrganizationPolicy) Attributes() projectOrganizationPolicyAttributes {
	return projectOrganizationPolicyAttributes{ref: terra.ReferenceResource(pop)}
}

// ImportState imports the given attribute values into [ProjectOrganizationPolicy]'s state.
func (pop *ProjectOrganizationPolicy) ImportState(av io.Reader) error {
	pop.state = &projectOrganizationPolicyState{}
	if err := json.NewDecoder(av).Decode(pop.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pop.Type(), pop.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectOrganizationPolicy] has state.
func (pop *ProjectOrganizationPolicy) State() (*projectOrganizationPolicyState, bool) {
	return pop.state, pop.state != nil
}

// StateMust returns the state for [ProjectOrganizationPolicy]. Panics if the state is nil.
func (pop *ProjectOrganizationPolicy) StateMust() *projectOrganizationPolicyState {
	if pop.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pop.Type(), pop.LocalName()))
	}
	return pop.state
}

// ProjectOrganizationPolicyArgs contains the configurations for google_project_organization_policy.
type ProjectOrganizationPolicyArgs struct {
	// Constraint: string, required
	Constraint terra.StringValue `hcl:"constraint,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// BooleanPolicy: optional
	BooleanPolicy *projectorganizationpolicy.BooleanPolicy `hcl:"boolean_policy,block"`
	// ListPolicy: optional
	ListPolicy *projectorganizationpolicy.ListPolicy `hcl:"list_policy,block"`
	// RestorePolicy: optional
	RestorePolicy *projectorganizationpolicy.RestorePolicy `hcl:"restore_policy,block"`
	// Timeouts: optional
	Timeouts *projectorganizationpolicy.Timeouts `hcl:"timeouts,block"`
}
type projectOrganizationPolicyAttributes struct {
	ref terra.Reference
}

// Constraint returns a reference to field constraint of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) Constraint() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("constraint"))
}

// Etag returns a reference to field etag of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("id"))
}

// Project returns a reference to field project of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("project"))
}

// UpdateTime returns a reference to field update_time of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(pop.ref.Append("update_time"))
}

// Version returns a reference to field version of google_project_organization_policy.
func (pop projectOrganizationPolicyAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(pop.ref.Append("version"))
}

func (pop projectOrganizationPolicyAttributes) BooleanPolicy() terra.ListValue[projectorganizationpolicy.BooleanPolicyAttributes] {
	return terra.ReferenceAsList[projectorganizationpolicy.BooleanPolicyAttributes](pop.ref.Append("boolean_policy"))
}

func (pop projectOrganizationPolicyAttributes) ListPolicy() terra.ListValue[projectorganizationpolicy.ListPolicyAttributes] {
	return terra.ReferenceAsList[projectorganizationpolicy.ListPolicyAttributes](pop.ref.Append("list_policy"))
}

func (pop projectOrganizationPolicyAttributes) RestorePolicy() terra.ListValue[projectorganizationpolicy.RestorePolicyAttributes] {
	return terra.ReferenceAsList[projectorganizationpolicy.RestorePolicyAttributes](pop.ref.Append("restore_policy"))
}

func (pop projectOrganizationPolicyAttributes) Timeouts() projectorganizationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[projectorganizationpolicy.TimeoutsAttributes](pop.ref.Append("timeouts"))
}

type projectOrganizationPolicyState struct {
	Constraint    string                                         `json:"constraint"`
	Etag          string                                         `json:"etag"`
	Id            string                                         `json:"id"`
	Project       string                                         `json:"project"`
	UpdateTime    string                                         `json:"update_time"`
	Version       float64                                        `json:"version"`
	BooleanPolicy []projectorganizationpolicy.BooleanPolicyState `json:"boolean_policy"`
	ListPolicy    []projectorganizationpolicy.ListPolicyState    `json:"list_policy"`
	RestorePolicy []projectorganizationpolicy.RestorePolicyState `json:"restore_policy"`
	Timeouts      *projectorganizationpolicy.TimeoutsState       `json:"timeouts"`
}
