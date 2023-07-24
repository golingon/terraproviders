// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	projectiammember "github.com/golingon/terraproviders/google/4.74.0/projectiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewProjectIamMember creates a new instance of [ProjectIamMember].
func NewProjectIamMember(name string, args ProjectIamMemberArgs) *ProjectIamMember {
	return &ProjectIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ProjectIamMember)(nil)

// ProjectIamMember represents the Terraform resource google_project_iam_member.
type ProjectIamMember struct {
	Name      string
	Args      ProjectIamMemberArgs
	state     *projectIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ProjectIamMember].
func (pim *ProjectIamMember) Type() string {
	return "google_project_iam_member"
}

// LocalName returns the local name for [ProjectIamMember].
func (pim *ProjectIamMember) LocalName() string {
	return pim.Name
}

// Configuration returns the configuration (args) for [ProjectIamMember].
func (pim *ProjectIamMember) Configuration() interface{} {
	return pim.Args
}

// DependOn is used for other resources to depend on [ProjectIamMember].
func (pim *ProjectIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(pim)
}

// Dependencies returns the list of resources [ProjectIamMember] depends_on.
func (pim *ProjectIamMember) Dependencies() terra.Dependencies {
	return pim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ProjectIamMember].
func (pim *ProjectIamMember) LifecycleManagement() *terra.Lifecycle {
	return pim.Lifecycle
}

// Attributes returns the attributes for [ProjectIamMember].
func (pim *ProjectIamMember) Attributes() projectIamMemberAttributes {
	return projectIamMemberAttributes{ref: terra.ReferenceResource(pim)}
}

// ImportState imports the given attribute values into [ProjectIamMember]'s state.
func (pim *ProjectIamMember) ImportState(av io.Reader) error {
	pim.state = &projectIamMemberState{}
	if err := json.NewDecoder(av).Decode(pim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", pim.Type(), pim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ProjectIamMember] has state.
func (pim *ProjectIamMember) State() (*projectIamMemberState, bool) {
	return pim.state, pim.state != nil
}

// StateMust returns the state for [ProjectIamMember]. Panics if the state is nil.
func (pim *ProjectIamMember) StateMust() *projectIamMemberState {
	if pim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", pim.Type(), pim.LocalName()))
	}
	return pim.state
}

// ProjectIamMemberArgs contains the configurations for google_project_iam_member.
type ProjectIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *projectiammember.Condition `hcl:"condition,block"`
}
type projectIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_project_iam_member.
func (pim projectIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(pim.ref.Append("etag"))
}

// Id returns a reference to field id of google_project_iam_member.
func (pim projectIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pim.ref.Append("id"))
}

// Member returns a reference to field member of google_project_iam_member.
func (pim projectIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(pim.ref.Append("member"))
}

// Project returns a reference to field project of google_project_iam_member.
func (pim projectIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(pim.ref.Append("project"))
}

// Role returns a reference to field role of google_project_iam_member.
func (pim projectIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(pim.ref.Append("role"))
}

func (pim projectIamMemberAttributes) Condition() terra.ListValue[projectiammember.ConditionAttributes] {
	return terra.ReferenceAsList[projectiammember.ConditionAttributes](pim.ref.Append("condition"))
}

type projectIamMemberState struct {
	Etag      string                            `json:"etag"`
	Id        string                            `json:"id"`
	Member    string                            `json:"member"`
	Project   string                            `json:"project"`
	Role      string                            `json:"role"`
	Condition []projectiammember.ConditionState `json:"condition"`
}
