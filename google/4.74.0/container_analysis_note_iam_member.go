// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	containeranalysisnoteiammember "github.com/golingon/terraproviders/google/4.74.0/containeranalysisnoteiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewContainerAnalysisNoteIamMember creates a new instance of [ContainerAnalysisNoteIamMember].
func NewContainerAnalysisNoteIamMember(name string, args ContainerAnalysisNoteIamMemberArgs) *ContainerAnalysisNoteIamMember {
	return &ContainerAnalysisNoteIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerAnalysisNoteIamMember)(nil)

// ContainerAnalysisNoteIamMember represents the Terraform resource google_container_analysis_note_iam_member.
type ContainerAnalysisNoteIamMember struct {
	Name      string
	Args      ContainerAnalysisNoteIamMemberArgs
	state     *containerAnalysisNoteIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) Type() string {
	return "google_container_analysis_note_iam_member"
}

// LocalName returns the local name for [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) LocalName() string {
	return canim.Name
}

// Configuration returns the configuration (args) for [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) Configuration() interface{} {
	return canim.Args
}

// DependOn is used for other resources to depend on [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(canim)
}

// Dependencies returns the list of resources [ContainerAnalysisNoteIamMember] depends_on.
func (canim *ContainerAnalysisNoteIamMember) Dependencies() terra.Dependencies {
	return canim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) LifecycleManagement() *terra.Lifecycle {
	return canim.Lifecycle
}

// Attributes returns the attributes for [ContainerAnalysisNoteIamMember].
func (canim *ContainerAnalysisNoteIamMember) Attributes() containerAnalysisNoteIamMemberAttributes {
	return containerAnalysisNoteIamMemberAttributes{ref: terra.ReferenceResource(canim)}
}

// ImportState imports the given attribute values into [ContainerAnalysisNoteIamMember]'s state.
func (canim *ContainerAnalysisNoteIamMember) ImportState(av io.Reader) error {
	canim.state = &containerAnalysisNoteIamMemberState{}
	if err := json.NewDecoder(av).Decode(canim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", canim.Type(), canim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerAnalysisNoteIamMember] has state.
func (canim *ContainerAnalysisNoteIamMember) State() (*containerAnalysisNoteIamMemberState, bool) {
	return canim.state, canim.state != nil
}

// StateMust returns the state for [ContainerAnalysisNoteIamMember]. Panics if the state is nil.
func (canim *ContainerAnalysisNoteIamMember) StateMust() *containerAnalysisNoteIamMemberState {
	if canim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", canim.Type(), canim.LocalName()))
	}
	return canim.state
}

// ContainerAnalysisNoteIamMemberArgs contains the configurations for google_container_analysis_note_iam_member.
type ContainerAnalysisNoteIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Note: string, required
	Note terra.StringValue `hcl:"note,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *containeranalysisnoteiammember.Condition `hcl:"condition,block"`
}
type containerAnalysisNoteIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("etag"))
}

// Id returns a reference to field id of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("id"))
}

// Member returns a reference to field member of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("member"))
}

// Note returns a reference to field note of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Note() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("note"))
}

// Project returns a reference to field project of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("project"))
}

// Role returns a reference to field role of google_container_analysis_note_iam_member.
func (canim containerAnalysisNoteIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(canim.ref.Append("role"))
}

func (canim containerAnalysisNoteIamMemberAttributes) Condition() terra.ListValue[containeranalysisnoteiammember.ConditionAttributes] {
	return terra.ReferenceAsList[containeranalysisnoteiammember.ConditionAttributes](canim.ref.Append("condition"))
}

type containerAnalysisNoteIamMemberState struct {
	Etag      string                                          `json:"etag"`
	Id        string                                          `json:"id"`
	Member    string                                          `json:"member"`
	Note      string                                          `json:"note"`
	Project   string                                          `json:"project"`
	Role      string                                          `json:"role"`
	Condition []containeranalysisnoteiammember.ConditionState `json:"condition"`
}