// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagedatalakegen2path "github.com/golingon/terraproviders/azurerm/3.64.0/storagedatalakegen2path"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageDataLakeGen2Path creates a new instance of [StorageDataLakeGen2Path].
func NewStorageDataLakeGen2Path(name string, args StorageDataLakeGen2PathArgs) *StorageDataLakeGen2Path {
	return &StorageDataLakeGen2Path{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageDataLakeGen2Path)(nil)

// StorageDataLakeGen2Path represents the Terraform resource azurerm_storage_data_lake_gen2_path.
type StorageDataLakeGen2Path struct {
	Name      string
	Args      StorageDataLakeGen2PathArgs
	state     *storageDataLakeGen2PathState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) Type() string {
	return "azurerm_storage_data_lake_gen2_path"
}

// LocalName returns the local name for [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) LocalName() string {
	return sdlgp.Name
}

// Configuration returns the configuration (args) for [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) Configuration() interface{} {
	return sdlgp.Args
}

// DependOn is used for other resources to depend on [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) DependOn() terra.Reference {
	return terra.ReferenceResource(sdlgp)
}

// Dependencies returns the list of resources [StorageDataLakeGen2Path] depends_on.
func (sdlgp *StorageDataLakeGen2Path) Dependencies() terra.Dependencies {
	return sdlgp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) LifecycleManagement() *terra.Lifecycle {
	return sdlgp.Lifecycle
}

// Attributes returns the attributes for [StorageDataLakeGen2Path].
func (sdlgp *StorageDataLakeGen2Path) Attributes() storageDataLakeGen2PathAttributes {
	return storageDataLakeGen2PathAttributes{ref: terra.ReferenceResource(sdlgp)}
}

// ImportState imports the given attribute values into [StorageDataLakeGen2Path]'s state.
func (sdlgp *StorageDataLakeGen2Path) ImportState(av io.Reader) error {
	sdlgp.state = &storageDataLakeGen2PathState{}
	if err := json.NewDecoder(av).Decode(sdlgp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdlgp.Type(), sdlgp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageDataLakeGen2Path] has state.
func (sdlgp *StorageDataLakeGen2Path) State() (*storageDataLakeGen2PathState, bool) {
	return sdlgp.state, sdlgp.state != nil
}

// StateMust returns the state for [StorageDataLakeGen2Path]. Panics if the state is nil.
func (sdlgp *StorageDataLakeGen2Path) StateMust() *storageDataLakeGen2PathState {
	if sdlgp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdlgp.Type(), sdlgp.LocalName()))
	}
	return sdlgp.state
}

// StorageDataLakeGen2PathArgs contains the configurations for azurerm_storage_data_lake_gen2_path.
type StorageDataLakeGen2PathArgs struct {
	// FilesystemName: string, required
	FilesystemName terra.StringValue `hcl:"filesystem_name,attr" validate:"required"`
	// Group: string, optional
	Group terra.StringValue `hcl:"group,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
	// Resource: string, required
	Resource terra.StringValue `hcl:"resource,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Ace: min=0
	Ace []storagedatalakegen2path.Ace `hcl:"ace,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storagedatalakegen2path.Timeouts `hcl:"timeouts,block"`
}
type storageDataLakeGen2PathAttributes struct {
	ref terra.Reference
}

// FilesystemName returns a reference to field filesystem_name of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) FilesystemName() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("filesystem_name"))
}

// Group returns a reference to field group of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("group"))
}

// Id returns a reference to field id of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("id"))
}

// Owner returns a reference to field owner of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("owner"))
}

// Path returns a reference to field path of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("path"))
}

// Resource returns a reference to field resource of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) Resource() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("resource"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_data_lake_gen2_path.
func (sdlgp storageDataLakeGen2PathAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sdlgp.ref.Append("storage_account_id"))
}

func (sdlgp storageDataLakeGen2PathAttributes) Ace() terra.SetValue[storagedatalakegen2path.AceAttributes] {
	return terra.ReferenceAsSet[storagedatalakegen2path.AceAttributes](sdlgp.ref.Append("ace"))
}

func (sdlgp storageDataLakeGen2PathAttributes) Timeouts() storagedatalakegen2path.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagedatalakegen2path.TimeoutsAttributes](sdlgp.ref.Append("timeouts"))
}

type storageDataLakeGen2PathState struct {
	FilesystemName   string                                 `json:"filesystem_name"`
	Group            string                                 `json:"group"`
	Id               string                                 `json:"id"`
	Owner            string                                 `json:"owner"`
	Path             string                                 `json:"path"`
	Resource         string                                 `json:"resource"`
	StorageAccountId string                                 `json:"storage_account_id"`
	Ace              []storagedatalakegen2path.AceState     `json:"ace"`
	Timeouts         *storagedatalakegen2path.TimeoutsState `json:"timeouts"`
}
