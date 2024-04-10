// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	storagedatalakegen2filesystem "github.com/golingon/terraproviders/azurerm/3.98.0/storagedatalakegen2filesystem"
	"io"
)

// NewStorageDataLakeGen2Filesystem creates a new instance of [StorageDataLakeGen2Filesystem].
func NewStorageDataLakeGen2Filesystem(name string, args StorageDataLakeGen2FilesystemArgs) *StorageDataLakeGen2Filesystem {
	return &StorageDataLakeGen2Filesystem{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageDataLakeGen2Filesystem)(nil)

// StorageDataLakeGen2Filesystem represents the Terraform resource azurerm_storage_data_lake_gen2_filesystem.
type StorageDataLakeGen2Filesystem struct {
	Name      string
	Args      StorageDataLakeGen2FilesystemArgs
	state     *storageDataLakeGen2FilesystemState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) Type() string {
	return "azurerm_storage_data_lake_gen2_filesystem"
}

// LocalName returns the local name for [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) LocalName() string {
	return sdlgf.Name
}

// Configuration returns the configuration (args) for [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) Configuration() interface{} {
	return sdlgf.Args
}

// DependOn is used for other resources to depend on [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) DependOn() terra.Reference {
	return terra.ReferenceResource(sdlgf)
}

// Dependencies returns the list of resources [StorageDataLakeGen2Filesystem] depends_on.
func (sdlgf *StorageDataLakeGen2Filesystem) Dependencies() terra.Dependencies {
	return sdlgf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) LifecycleManagement() *terra.Lifecycle {
	return sdlgf.Lifecycle
}

// Attributes returns the attributes for [StorageDataLakeGen2Filesystem].
func (sdlgf *StorageDataLakeGen2Filesystem) Attributes() storageDataLakeGen2FilesystemAttributes {
	return storageDataLakeGen2FilesystemAttributes{ref: terra.ReferenceResource(sdlgf)}
}

// ImportState imports the given attribute values into [StorageDataLakeGen2Filesystem]'s state.
func (sdlgf *StorageDataLakeGen2Filesystem) ImportState(av io.Reader) error {
	sdlgf.state = &storageDataLakeGen2FilesystemState{}
	if err := json.NewDecoder(av).Decode(sdlgf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdlgf.Type(), sdlgf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageDataLakeGen2Filesystem] has state.
func (sdlgf *StorageDataLakeGen2Filesystem) State() (*storageDataLakeGen2FilesystemState, bool) {
	return sdlgf.state, sdlgf.state != nil
}

// StateMust returns the state for [StorageDataLakeGen2Filesystem]. Panics if the state is nil.
func (sdlgf *StorageDataLakeGen2Filesystem) StateMust() *storageDataLakeGen2FilesystemState {
	if sdlgf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdlgf.Type(), sdlgf.LocalName()))
	}
	return sdlgf.state
}

// StorageDataLakeGen2FilesystemArgs contains the configurations for azurerm_storage_data_lake_gen2_filesystem.
type StorageDataLakeGen2FilesystemArgs struct {
	// Group: string, optional
	Group terra.StringValue `hcl:"group,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Owner: string, optional
	Owner terra.StringValue `hcl:"owner,attr"`
	// Properties: map of string, optional
	Properties terra.MapValue[terra.StringValue] `hcl:"properties,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Ace: min=0
	Ace []storagedatalakegen2filesystem.Ace `hcl:"ace,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storagedatalakegen2filesystem.Timeouts `hcl:"timeouts,block"`
}
type storageDataLakeGen2FilesystemAttributes struct {
	ref terra.Reference
}

// Group returns a reference to field group of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) Group() terra.StringValue {
	return terra.ReferenceAsString(sdlgf.ref.Append("group"))
}

// Id returns a reference to field id of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdlgf.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sdlgf.ref.Append("name"))
}

// Owner returns a reference to field owner of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) Owner() terra.StringValue {
	return terra.ReferenceAsString(sdlgf.ref.Append("owner"))
}

// Properties returns a reference to field properties of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) Properties() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdlgf.ref.Append("properties"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_storage_data_lake_gen2_filesystem.
func (sdlgf storageDataLakeGen2FilesystemAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sdlgf.ref.Append("storage_account_id"))
}

func (sdlgf storageDataLakeGen2FilesystemAttributes) Ace() terra.SetValue[storagedatalakegen2filesystem.AceAttributes] {
	return terra.ReferenceAsSet[storagedatalakegen2filesystem.AceAttributes](sdlgf.ref.Append("ace"))
}

func (sdlgf storageDataLakeGen2FilesystemAttributes) Timeouts() storagedatalakegen2filesystem.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagedatalakegen2filesystem.TimeoutsAttributes](sdlgf.ref.Append("timeouts"))
}

type storageDataLakeGen2FilesystemState struct {
	Group            string                                       `json:"group"`
	Id               string                                       `json:"id"`
	Name             string                                       `json:"name"`
	Owner            string                                       `json:"owner"`
	Properties       map[string]string                            `json:"properties"`
	StorageAccountId string                                       `json:"storage_account_id"`
	Ace              []storagedatalakegen2filesystem.AceState     `json:"ace"`
	Timeouts         *storagedatalakegen2filesystem.TimeoutsState `json:"timeouts"`
}
