// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	folder "github.com/golingon/terraproviders/google/4.76.0/folder"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolder creates a new instance of [Folder].
func NewFolder(name string, args FolderArgs) *Folder {
	return &Folder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Folder)(nil)

// Folder represents the Terraform resource google_folder.
type Folder struct {
	Name      string
	Args      FolderArgs
	state     *folderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Folder].
func (f *Folder) Type() string {
	return "google_folder"
}

// LocalName returns the local name for [Folder].
func (f *Folder) LocalName() string {
	return f.Name
}

// Configuration returns the configuration (args) for [Folder].
func (f *Folder) Configuration() interface{} {
	return f.Args
}

// DependOn is used for other resources to depend on [Folder].
func (f *Folder) DependOn() terra.Reference {
	return terra.ReferenceResource(f)
}

// Dependencies returns the list of resources [Folder] depends_on.
func (f *Folder) Dependencies() terra.Dependencies {
	return f.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Folder].
func (f *Folder) LifecycleManagement() *terra.Lifecycle {
	return f.Lifecycle
}

// Attributes returns the attributes for [Folder].
func (f *Folder) Attributes() folderAttributes {
	return folderAttributes{ref: terra.ReferenceResource(f)}
}

// ImportState imports the given attribute values into [Folder]'s state.
func (f *Folder) ImportState(av io.Reader) error {
	f.state = &folderState{}
	if err := json.NewDecoder(av).Decode(f.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", f.Type(), f.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Folder] has state.
func (f *Folder) State() (*folderState, bool) {
	return f.state, f.state != nil
}

// StateMust returns the state for [Folder]. Panics if the state is nil.
func (f *Folder) StateMust() *folderState {
	if f.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", f.Type(), f.LocalName()))
	}
	return f.state
}

// FolderArgs contains the configurations for google_folder.
type FolderArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, required
	Parent terra.StringValue `hcl:"parent,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *folder.Timeouts `hcl:"timeouts,block"`
}
type folderAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_folder.
func (f folderAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("create_time"))
}

// DisplayName returns a reference to field display_name of google_folder.
func (f folderAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("display_name"))
}

// FolderId returns a reference to field folder_id of google_folder.
func (f folderAttributes) FolderId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("folder_id"))
}

// Id returns a reference to field id of google_folder.
func (f folderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("id"))
}

// LifecycleState returns a reference to field lifecycle_state of google_folder.
func (f folderAttributes) LifecycleState() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("lifecycle_state"))
}

// Name returns a reference to field name of google_folder.
func (f folderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

// Parent returns a reference to field parent of google_folder.
func (f folderAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("parent"))
}

func (f folderAttributes) Timeouts() folder.TimeoutsAttributes {
	return terra.ReferenceAsSingle[folder.TimeoutsAttributes](f.ref.Append("timeouts"))
}

type folderState struct {
	CreateTime     string                `json:"create_time"`
	DisplayName    string                `json:"display_name"`
	FolderId       string                `json:"folder_id"`
	Id             string                `json:"id"`
	LifecycleState string                `json:"lifecycle_state"`
	Name           string                `json:"name"`
	Parent         string                `json:"parent"`
	Timeouts       *folder.TimeoutsState `json:"timeouts"`
}
