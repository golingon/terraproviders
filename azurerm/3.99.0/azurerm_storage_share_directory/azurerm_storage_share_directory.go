// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_storage_share_directory

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_storage_share_directory.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermStorageShareDirectoryState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (assd *Resource) Type() string {
	return "azurerm_storage_share_directory"
}

// LocalName returns the local name for [Resource].
func (assd *Resource) LocalName() string {
	return assd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (assd *Resource) Configuration() interface{} {
	return assd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (assd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(assd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (assd *Resource) Dependencies() terra.Dependencies {
	return assd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (assd *Resource) LifecycleManagement() *terra.Lifecycle {
	return assd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (assd *Resource) Attributes() azurermStorageShareDirectoryAttributes {
	return azurermStorageShareDirectoryAttributes{ref: terra.ReferenceResource(assd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (assd *Resource) ImportState(state io.Reader) error {
	assd.state = &azurermStorageShareDirectoryState{}
	if err := json.NewDecoder(state).Decode(assd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", assd.Type(), assd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (assd *Resource) State() (*azurermStorageShareDirectoryState, bool) {
	return assd.state, assd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (assd *Resource) StateMust() *azurermStorageShareDirectoryState {
	if assd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", assd.Type(), assd.LocalName()))
	}
	return assd.state
}

// Args contains the configurations for azurerm_storage_share_directory.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareName: string, optional
	ShareName terra.StringValue `hcl:"share_name,attr"`
	// StorageAccountName: string, optional
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr"`
	// StorageShareId: string, optional
	StorageShareId terra.StringValue `hcl:"storage_share_id,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermStorageShareDirectoryAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(assd.ref.Append("id"))
}

// Metadata returns a reference to field metadata of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](assd.ref.Append("metadata"))
}

// Name returns a reference to field name of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(assd.ref.Append("name"))
}

// ShareName returns a reference to field share_name of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) ShareName() terra.StringValue {
	return terra.ReferenceAsString(assd.ref.Append("share_name"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(assd.ref.Append("storage_account_name"))
}

// StorageShareId returns a reference to field storage_share_id of azurerm_storage_share_directory.
func (assd azurermStorageShareDirectoryAttributes) StorageShareId() terra.StringValue {
	return terra.ReferenceAsString(assd.ref.Append("storage_share_id"))
}

func (assd azurermStorageShareDirectoryAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](assd.ref.Append("timeouts"))
}

type azurermStorageShareDirectoryState struct {
	Id                 string            `json:"id"`
	Metadata           map[string]string `json:"metadata"`
	Name               string            `json:"name"`
	ShareName          string            `json:"share_name"`
	StorageAccountName string            `json:"storage_account_name"`
	StorageShareId     string            `json:"storage_share_id"`
	Timeouts           *TimeoutsState    `json:"timeouts"`
}
