// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	notebooksruntimeiammember "github.com/golingon/terraproviders/google/4.59.0/notebooksruntimeiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNotebooksRuntimeIamMember(name string, args NotebooksRuntimeIamMemberArgs) *NotebooksRuntimeIamMember {
	return &NotebooksRuntimeIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NotebooksRuntimeIamMember)(nil)

type NotebooksRuntimeIamMember struct {
	Name  string
	Args  NotebooksRuntimeIamMemberArgs
	state *notebooksRuntimeIamMemberState
}

func (nrim *NotebooksRuntimeIamMember) Type() string {
	return "google_notebooks_runtime_iam_member"
}

func (nrim *NotebooksRuntimeIamMember) LocalName() string {
	return nrim.Name
}

func (nrim *NotebooksRuntimeIamMember) Configuration() interface{} {
	return nrim.Args
}

func (nrim *NotebooksRuntimeIamMember) Attributes() notebooksRuntimeIamMemberAttributes {
	return notebooksRuntimeIamMemberAttributes{ref: terra.ReferenceResource(nrim)}
}

func (nrim *NotebooksRuntimeIamMember) ImportState(av io.Reader) error {
	nrim.state = &notebooksRuntimeIamMemberState{}
	if err := json.NewDecoder(av).Decode(nrim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nrim.Type(), nrim.LocalName(), err)
	}
	return nil
}

func (nrim *NotebooksRuntimeIamMember) State() (*notebooksRuntimeIamMemberState, bool) {
	return nrim.state, nrim.state != nil
}

func (nrim *NotebooksRuntimeIamMember) StateMust() *notebooksRuntimeIamMemberState {
	if nrim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nrim.Type(), nrim.LocalName()))
	}
	return nrim.state
}

func (nrim *NotebooksRuntimeIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(nrim)
}

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
	// DependsOn contains resources that NotebooksRuntimeIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type notebooksRuntimeIamMemberAttributes struct {
	ref terra.Reference
}

func (nrim notebooksRuntimeIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("etag"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("id"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Location() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("location"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("member"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Project() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("project"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("role"))
}

func (nrim notebooksRuntimeIamMemberAttributes) RuntimeName() terra.StringValue {
	return terra.ReferenceString(nrim.ref.Append("runtime_name"))
}

func (nrim notebooksRuntimeIamMemberAttributes) Condition() terra.ListValue[notebooksruntimeiammember.ConditionAttributes] {
	return terra.ReferenceList[notebooksruntimeiammember.ConditionAttributes](nrim.ref.Append("condition"))
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
