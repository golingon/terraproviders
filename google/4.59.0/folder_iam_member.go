// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	folderiammember "github.com/golingon/terraproviders/google/4.59.0/folderiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewFolderIamMember(name string, args FolderIamMemberArgs) *FolderIamMember {
	return &FolderIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderIamMember)(nil)

type FolderIamMember struct {
	Name  string
	Args  FolderIamMemberArgs
	state *folderIamMemberState
}

func (fim *FolderIamMember) Type() string {
	return "google_folder_iam_member"
}

func (fim *FolderIamMember) LocalName() string {
	return fim.Name
}

func (fim *FolderIamMember) Configuration() interface{} {
	return fim.Args
}

func (fim *FolderIamMember) Attributes() folderIamMemberAttributes {
	return folderIamMemberAttributes{ref: terra.ReferenceResource(fim)}
}

func (fim *FolderIamMember) ImportState(av io.Reader) error {
	fim.state = &folderIamMemberState{}
	if err := json.NewDecoder(av).Decode(fim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fim.Type(), fim.LocalName(), err)
	}
	return nil
}

func (fim *FolderIamMember) State() (*folderIamMemberState, bool) {
	return fim.state, fim.state != nil
}

func (fim *FolderIamMember) StateMust() *folderIamMemberState {
	if fim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fim.Type(), fim.LocalName()))
	}
	return fim.state
}

func (fim *FolderIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(fim)
}

type FolderIamMemberArgs struct {
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Member: string, required
	Member terra.StringValue `hcl:"member,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *folderiammember.Condition `hcl:"condition,block"`
	// DependsOn contains resources that FolderIamMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type folderIamMemberAttributes struct {
	ref terra.Reference
}

func (fim folderIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(fim.ref.Append("etag"))
}

func (fim folderIamMemberAttributes) Folder() terra.StringValue {
	return terra.ReferenceString(fim.ref.Append("folder"))
}

func (fim folderIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(fim.ref.Append("id"))
}

func (fim folderIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceString(fim.ref.Append("member"))
}

func (fim folderIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceString(fim.ref.Append("role"))
}

func (fim folderIamMemberAttributes) Condition() terra.ListValue[folderiammember.ConditionAttributes] {
	return terra.ReferenceList[folderiammember.ConditionAttributes](fim.ref.Append("condition"))
}

type folderIamMemberState struct {
	Etag      string                           `json:"etag"`
	Folder    string                           `json:"folder"`
	Id        string                           `json:"id"`
	Member    string                           `json:"member"`
	Role      string                           `json:"role"`
	Condition []folderiammember.ConditionState `json:"condition"`
}
