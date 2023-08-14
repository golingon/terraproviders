// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	folderiammember "github.com/golingon/terraproviders/googlebeta/4.77.0/folderiammember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolderIamMember creates a new instance of [FolderIamMember].
func NewFolderIamMember(name string, args FolderIamMemberArgs) *FolderIamMember {
	return &FolderIamMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderIamMember)(nil)

// FolderIamMember represents the Terraform resource google_folder_iam_member.
type FolderIamMember struct {
	Name      string
	Args      FolderIamMemberArgs
	state     *folderIamMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FolderIamMember].
func (fim *FolderIamMember) Type() string {
	return "google_folder_iam_member"
}

// LocalName returns the local name for [FolderIamMember].
func (fim *FolderIamMember) LocalName() string {
	return fim.Name
}

// Configuration returns the configuration (args) for [FolderIamMember].
func (fim *FolderIamMember) Configuration() interface{} {
	return fim.Args
}

// DependOn is used for other resources to depend on [FolderIamMember].
func (fim *FolderIamMember) DependOn() terra.Reference {
	return terra.ReferenceResource(fim)
}

// Dependencies returns the list of resources [FolderIamMember] depends_on.
func (fim *FolderIamMember) Dependencies() terra.Dependencies {
	return fim.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FolderIamMember].
func (fim *FolderIamMember) LifecycleManagement() *terra.Lifecycle {
	return fim.Lifecycle
}

// Attributes returns the attributes for [FolderIamMember].
func (fim *FolderIamMember) Attributes() folderIamMemberAttributes {
	return folderIamMemberAttributes{ref: terra.ReferenceResource(fim)}
}

// ImportState imports the given attribute values into [FolderIamMember]'s state.
func (fim *FolderIamMember) ImportState(av io.Reader) error {
	fim.state = &folderIamMemberState{}
	if err := json.NewDecoder(av).Decode(fim.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fim.Type(), fim.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FolderIamMember] has state.
func (fim *FolderIamMember) State() (*folderIamMemberState, bool) {
	return fim.state, fim.state != nil
}

// StateMust returns the state for [FolderIamMember]. Panics if the state is nil.
func (fim *FolderIamMember) StateMust() *folderIamMemberState {
	if fim.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fim.Type(), fim.LocalName()))
	}
	return fim.state
}

// FolderIamMemberArgs contains the configurations for google_folder_iam_member.
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
}
type folderIamMemberAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_folder_iam_member.
func (fim folderIamMemberAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fim.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_iam_member.
func (fim folderIamMemberAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fim.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_iam_member.
func (fim folderIamMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fim.ref.Append("id"))
}

// Member returns a reference to field member of google_folder_iam_member.
func (fim folderIamMemberAttributes) Member() terra.StringValue {
	return terra.ReferenceAsString(fim.ref.Append("member"))
}

// Role returns a reference to field role of google_folder_iam_member.
func (fim folderIamMemberAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(fim.ref.Append("role"))
}

func (fim folderIamMemberAttributes) Condition() terra.ListValue[folderiammember.ConditionAttributes] {
	return terra.ReferenceAsList[folderiammember.ConditionAttributes](fim.ref.Append("condition"))
}

type folderIamMemberState struct {
	Etag      string                           `json:"etag"`
	Folder    string                           `json:"folder"`
	Id        string                           `json:"id"`
	Member    string                           `json:"member"`
	Role      string                           `json:"role"`
	Condition []folderiammember.ConditionState `json:"condition"`
}
