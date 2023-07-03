// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagesharedirectory "github.com/golingon/terraproviders/azurerm/3.63.0/storagesharedirectory"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageShareDirectory creates a new instance of [StorageShareDirectory].
func NewStorageShareDirectory(name string, args StorageShareDirectoryArgs) *StorageShareDirectory {
	return &StorageShareDirectory{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageShareDirectory)(nil)

// StorageShareDirectory represents the Terraform resource azurerm_storage_share_directory.
type StorageShareDirectory struct {
	Name      string
	Args      StorageShareDirectoryArgs
	state     *storageShareDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageShareDirectory].
func (ssd *StorageShareDirectory) Type() string {
	return "azurerm_storage_share_directory"
}

// LocalName returns the local name for [StorageShareDirectory].
func (ssd *StorageShareDirectory) LocalName() string {
	return ssd.Name
}

// Configuration returns the configuration (args) for [StorageShareDirectory].
func (ssd *StorageShareDirectory) Configuration() interface{} {
	return ssd.Args
}

// DependOn is used for other resources to depend on [StorageShareDirectory].
func (ssd *StorageShareDirectory) DependOn() terra.Reference {
	return terra.ReferenceResource(ssd)
}

// Dependencies returns the list of resources [StorageShareDirectory] depends_on.
func (ssd *StorageShareDirectory) Dependencies() terra.Dependencies {
	return ssd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageShareDirectory].
func (ssd *StorageShareDirectory) LifecycleManagement() *terra.Lifecycle {
	return ssd.Lifecycle
}

// Attributes returns the attributes for [StorageShareDirectory].
func (ssd *StorageShareDirectory) Attributes() storageShareDirectoryAttributes {
	return storageShareDirectoryAttributes{ref: terra.ReferenceResource(ssd)}
}

// ImportState imports the given attribute values into [StorageShareDirectory]'s state.
func (ssd *StorageShareDirectory) ImportState(av io.Reader) error {
	ssd.state = &storageShareDirectoryState{}
	if err := json.NewDecoder(av).Decode(ssd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssd.Type(), ssd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageShareDirectory] has state.
func (ssd *StorageShareDirectory) State() (*storageShareDirectoryState, bool) {
	return ssd.state, ssd.state != nil
}

// StateMust returns the state for [StorageShareDirectory]. Panics if the state is nil.
func (ssd *StorageShareDirectory) StateMust() *storageShareDirectoryState {
	if ssd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssd.Type(), ssd.LocalName()))
	}
	return ssd.state
}

// StorageShareDirectoryArgs contains the configurations for azurerm_storage_share_directory.
type StorageShareDirectoryArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareName: string, required
	ShareName terra.StringValue `hcl:"share_name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *storagesharedirectory.Timeouts `hcl:"timeouts,block"`
}
type storageShareDirectoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_share_directory.
func (ssd storageShareDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssd.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_share_directory.
func (ssd storageShareDirectoryAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssd.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_share_directory.
func (ssd storageShareDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssd.ref.Append("name"))
}

// ShareName returns a reference to field share_name of azurerm_storage_share_directory.
func (ssd storageShareDirectoryAttributes) ShareName() terra.StringValue {
	return terra.ReferenceAsString(ssd.ref.Append("share_name"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_share_directory.
func (ssd storageShareDirectoryAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(ssd.ref.Append("storage_account_name"))
}

func (ssd storageShareDirectoryAttributes) Timeouts() storagesharedirectory.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagesharedirectory.TimeoutsAttributes](ssd.ref.Append("timeouts"))
}

type storageShareDirectoryState struct {
	Id                 string                               `json:"id"`
	Metadata           map[string]string                    `json:"metadata"`
	Name               string                               `json:"name"`
	ShareName          string                               `json:"share_name"`
	StorageAccountName string                               `json:"storage_account_name"`
	Timeouts           *storagesharedirectory.TimeoutsState `json:"timeouts"`
}
