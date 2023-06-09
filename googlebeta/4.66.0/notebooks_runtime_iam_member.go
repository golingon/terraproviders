// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	notebooksruntimeiammember "github.com/golingon/terraproviders/googlebeta/4.66.0/notebooksruntimeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNotebooksRuntimeIamMember creates a new instance of [NotebooksRuntimeIamMember].
func NewNotebooksRuntimeIamMember(name string, args NotebooksRuntimeIamMemberArgs) *NotebooksRuntimeIamMember {
	return &NotebooksRuntimeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksRuntimeIamMember)(nil)

// NotebooksRuntimeIamMember represents the Terraform resource google_notebooks_runtime_iam_member.
type NotebooksRuntimeIamMember struct {
	Name      string
	Args      NotebooksRuntimeIamMemberArgs
	state     *notebooksRuntimeIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) Type() string {
	return "google_notebooks_runtime_iam_member"
}

// LocalName returns the local name for [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) LocalName() string {
	return nrim.Name
}

// Configuration returns the configuration (args) for [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) Configuration() interface{} {
	return nrim.Args
}

// DependOn is used for other resources to depend on [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(nrim)
}

// Dependencies returns the list of resources [NotebooksRuntimeIamMember] depends_on.
func (nrim *NotebooksRuntimeIamMember) Dependencies() terra.Dependencies {
	return nrim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) LifecycleManagement() *terra.Lifecycle {
	return nrim.Lifecycle
}

// Attributes returns the attributes for [NotebooksRuntimeIamMember].
func (nrim *NotebooksRuntimeIamMember) Attributes() notebooksRuntimeIamMemberAttributes {
	return notebooksRuntimeIamMemberAttributes{ref: terra.ReferenceResource(nrim)}
}

// ImportState imports the given attribute values into [NotebooksRuntimeIamMember]'s state.
func (nrim *NotebooksRuntimeIamMember) ImportState(av io.Reader) error {
	nrim.state = &notebooksRuntimeIamMemberState{}
	if err := json.NewDecoder(av).Decode(nrim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrim.Type(), nrim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NotebooksRuntimeIamMember] has state.
func (nrim *NotebooksRuntimeIamMember) State() (*notebooksRuntimeIamMemberState, bool) {
	return nrim.state, nrim.state != nil
}

// StateMust returns the state for [NotebooksRuntimeIamMember]. Panics if the state is nil.
func (nrim *NotebooksRuntimeIamMember) StateMust() *notebooksRuntimeIamMemberState {
	if nrim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrim.Type(), nrim.LocalName()))
	}
	return nrim.state
}

// NotebooksRuntimeIamMemberArgs contains the configurations for google_notebooks_runtime_iam_member.
type NotebooksRuntimeIamMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// RuntimeName: string, required
	RuntimeName terra.StringValue `hcl:"runtime_name,attr" validate:"required"`
	// Condition: optional
	Condition *notebooksruntimeiammember.Condition `hcl:"condition,block"`
}
type notebooksRuntimeIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("etag"))
}

// Id returns a reference to field id of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("id"))
}

// Location returns a reference to field location of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("location"))
}

// Member returns a reference to field member of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("member"))
}

// Project returns a reference to field project of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("project"))
}

// Role returns a reference to field role of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("role"))
}

// RuntimeName returns a reference to field runtime_name of google_notebooks_runtime_iam_member.
func (nrim notebooksRuntimeIamMemberAttributes) RuntimeName() terra.StringValue {
	return terra.ReferenceAsString(nrim.ref.Append("runtime_name"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Condition() terra.ListValue[notebooksruntimeiammember.ConditionAttributes] {
	return terra.ReferenceAsList[notebooksruntimeiammember.ConditionAttributes](nrim.ref.Append("condition"))
}

type notebooksRuntimeIamMemberState struct {
	Etag        string                                     `json:"etag"`
	Id          string                                     `json:"id"`
	Location    string                                     `json:"location"`
	Member      string                                     `json:"member"`
	Project     string                                     `json:"project"`
	Role        string                                     `json:"role"`
	RuntimeName string                                     `json:"runtime_name"`
	Condition   []notebooksruntimeiammember.ConditionState `json:"condition"`
}
