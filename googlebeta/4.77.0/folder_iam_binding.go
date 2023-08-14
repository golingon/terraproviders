// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	folderiambinding "github.com/golingon/terraproviders/googlebeta/4.77.0/folderiambinding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFolderIamBinding creates a new instance of [FolderIamBinding].
func NewFolderIamBinding(name string, args FolderIamBindingArgs) *FolderIamBinding {
	return &FolderIamBinding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FolderIamBinding)(nil)

// FolderIamBinding represents the Terraform resource google_folder_iam_binding.
type FolderIamBinding struct {
	Name      string
	Args      FolderIamBindingArgs
	state     *folderIamBindingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FolderIamBinding].
func (fib *FolderIamBinding) Type() string {
	return "google_folder_iam_binding"
}

// LocalName returns the local name for [FolderIamBinding].
func (fib *FolderIamBinding) LocalName() string {
	return fib.Name
}

// Configuration returns the configuration (args) for [FolderIamBinding].
func (fib *FolderIamBinding) Configuration() interface{} {
	return fib.Args
}

// DependOn is used for other resources to depend on [FolderIamBinding].
func (fib *FolderIamBinding) DependOn() terra.Reference {
	return terra.ReferenceResource(fib)
}

// Dependencies returns the list of resources [FolderIamBinding] depends_on.
func (fib *FolderIamBinding) Dependencies() terra.Dependencies {
	return fib.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FolderIamBinding].
func (fib *FolderIamBinding) LifecycleManagement() *terra.Lifecycle {
	return fib.Lifecycle
}

// Attributes returns the attributes for [FolderIamBinding].
func (fib *FolderIamBinding) Attributes() folderIamBindingAttributes {
	return folderIamBindingAttributes{ref: terra.ReferenceResource(fib)}
}

// ImportState imports the given attribute values into [FolderIamBinding]'s state.
func (fib *FolderIamBinding) ImportState(av io.Reader) error {
	fib.state = &folderIamBindingState{}
	if err := json.NewDecoder(av).Decode(fib.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fib.Type(), fib.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FolderIamBinding] has state.
func (fib *FolderIamBinding) State() (*folderIamBindingState, bool) {
	return fib.state, fib.state != nil
}

// StateMust returns the state for [FolderIamBinding]. Panics if the state is nil.
func (fib *FolderIamBinding) StateMust() *folderIamBindingState {
	if fib.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fib.Type(), fib.LocalName()))
	}
	return fib.state
}

// FolderIamBindingArgs contains the configurations for google_folder_iam_binding.
type FolderIamBindingArgs struct {
	// Folder: string, required
	Folder terra.StringValue `hcl:"folder,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Members: set of string, required
	Members terra.SetValue[terra.StringValue] `hcl:"members,attr" validate:"required"`
	// Role: string, required
	Role terra.StringValue `hcl:"role,attr" validate:"required"`
	// Condition: optional
	Condition *folderiambinding.Condition `hcl:"condition,block"`
}
type folderIamBindingAttributes struct {
	ref terra.Reference
}

// Etag returns a reference to field etag of google_folder_iam_binding.
func (fib folderIamBindingAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(fib.ref.Append("etag"))
}

// Folder returns a reference to field folder of google_folder_iam_binding.
func (fib folderIamBindingAttributes) Folder() terra.StringValue {
	return terra.ReferenceAsString(fib.ref.Append("folder"))
}

// Id returns a reference to field id of google_folder_iam_binding.
func (fib folderIamBindingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fib.ref.Append("id"))
}

// Members returns a reference to field members of google_folder_iam_binding.
func (fib folderIamBindingAttributes) Members() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](fib.ref.Append("members"))
}

// Role returns a reference to field role of google_folder_iam_binding.
func (fib folderIamBindingAttributes) Role() terra.StringValue {
	return terra.ReferenceAsString(fib.ref.Append("role"))
}

func (fib folderIamBindingAttributes) Condition() terra.ListValue[folderiambinding.ConditionAttributes] {
	return terra.ReferenceAsList[folderiambinding.ConditionAttributes](fib.ref.Append("condition"))
}

type folderIamBindingState struct {
	Etag      string                            `json:"etag"`
	Folder    string                            `json:"folder"`
	Id        string                            `json:"id"`
	Members   []string                          `json:"members"`
	Role      string                            `json:"role"`
	Condition []folderiambinding.ConditionState `json:"condition"`
}
