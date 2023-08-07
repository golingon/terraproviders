// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	notebooksinstanceiammember "github.com/golingon/terraproviders/googlebeta/4.76.0/notebooksinstanceiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksInstanceIamMember creates a new instance of [NotebooksInstanceIamMember].
func NewNotebooksInstanceIamMember(name string, args NotebooksInstanceIamMemberArgs) *NotebooksInstanceIamMember {
	return &NotebooksInstanceIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksInstanceIamMember)(nil)

// NotebooksInstanceIamMember represents the Terraform resource google_notebooks_instance_iam_member.
type NotebooksInstanceIamMember struct {
	Name      string
	Args      NotebooksInstanceIamMemberArgs
	state     *notebooksInstanceIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) Type() string {
	return "google_notebooks_instance_iam_member"
}

// LocalName returns the local name for [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) LocalName() string {
	return niim.Name
}

// Configuration returns the configuration (args) for [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) Configuration() interface{} {
	return niim.Args
}

// DependOn is used for other resources to depend on [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(niim)
}

// Dependencies returns the list of resources [NotebooksInstanceIamMember] depends_on.
func (niim *NotebooksInstanceIamMember) Dependencies() terra.Dependencies {
	return niim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) LifecycleManagement() *terra.Lifecycle {
	return niim.Lifecycle
}

// Attributes returns the attributes for [NotebooksInstanceIamMember].
func (niim *NotebooksInstanceIamMember) Attributes() notebooksInstanceIamMemberAttributes {
	return notebooksInstanceIamMemberAttributes{ref: terra.ReferenceResource(niim)}
}

// ImportState imports the given attribute values into [NotebooksInstanceIamMember]'s state.
func (niim *NotebooksInstanceIamMember) ImportState(av io.Reader) error {
	niim.state = &notebooksInstanceIamMemberState{}
	if err := json.NewDecoder(av).Decode(niim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", niim.Type(), niim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksInstanceIamMember] has state.
func (niim *NotebooksInstanceIamMember) State() (*notebooksInstanceIamMemberState, bool) {
	return niim.state, niim.state != nil
}

// StateMust returns the state for [NotebooksInstanceIamMember]. Panics if the state is nil.
func (niim *NotebooksInstanceIamMember) StateMust() *notebooksInstanceIamMemberState {
	if niim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", niim.Type(), niim.LocalName()))
	}
	return niim.state
}

// NotebooksInstanceIamMemberArgs contains the configurations for google_notebooks_instance_iam_member.
type NotebooksInstanceIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceName: string, required
	InstanceName terra.StringValue `hcl:"instance_name,attr" validate:"required"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *notebooksinstanceiammember.Condition `hcl:"condition,block"`
}
type notebooksInstanceIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("id"))
}

// InstanceName returns a reference to field instance_name of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) InstanceName() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("instance_name"))
}

// Location returns a reference to field location of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("location"))
}

// Member returns a reference to field member of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("member"))
}

// Project returns a reference to field project of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("project"))
}

// Role returns a reference to field role of google_notebooks_instance_iam_member.
func (niim notebooksInstanceIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(niim.ref.Append("role"))
}

func (niim notebooksInstanceIamMemberAttributes) Condition() terra.ListValue[notebooksinstanceiammember.ConditionAttributes] {
	return terra.ReferenceAsList[notebooksinstanceiammember.ConditionAttributes](niim.ref.Append("condition"))
}

type notebooksInstanceIamMemberState struct {
	Etag         string                                      `json:"etag"`
	Id           string                                      `json:"id"`
	InstanceName string                                      `json:"instance_name"`
	Location     string                                      `json:"location"`
	Member       string                                      `json:"member"`
	Project      string                                      `json:"project"`
	Role         string                                      `json:"role"`
	Condition    []notebooksinstanceiammember.ConditionState `json:"condition"`
}
