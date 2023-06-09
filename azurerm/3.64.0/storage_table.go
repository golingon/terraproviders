// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storagetable "github.com/golingon/terraproviders/azurerm/3.64.0/storagetable"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageTable creates a new instance of [StorageTable].
func NewStorageTable(name string, args StorageTableArgs) *StorageTable {
	return &StorageTable{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageTable)(nil)

// StorageTable represents the Terraform resource azurerm_storage_table.
type StorageTable struct {
	Name      string
	Args      StorageTableArgs
	state     *storageTableState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageTable].
func (st *StorageTable) Type() string {
	return "azurerm_storage_table"
}

// LocalName returns the local name for [StorageTable].
func (st *StorageTable) LocalName() string {
	return st.Name
}

// Configuration returns the configuration (args) for [StorageTable].
func (st *StorageTable) Configuration() interface{} {
	return st.Args
}

// DependOn is used for other resources to depend on [StorageTable].
func (st *StorageTable) DependOn() terra.Reference {
	return terra.ReferenceResource(st)
}

// Dependencies returns the list of resources [StorageTable] depends_on.
func (st *StorageTable) Dependencies() terra.Dependencies {
	return st.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageTable].
func (st *StorageTable) LifecycleManagement() *terra.Lifecycle {
	return st.Lifecycle
}

// Attributes returns the attributes for [StorageTable].
func (st *StorageTable) Attributes() storageTableAttributes {
	return storageTableAttributes{ref: terra.ReferenceResource(st)}
}

// ImportState imports the given attribute values into [StorageTable]'s state.
func (st *StorageTable) ImportState(av io.Reader) error {
	st.state = &storageTableState{}
	if err := json.NewDecoder(av).Decode(st.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", st.Type(), st.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageTable] has state.
func (st *StorageTable) State() (*storageTableState, bool) {
	return st.state, st.state != nil
}

// StateMust returns the state for [StorageTable]. Panics if the state is nil.
func (st *StorageTable) StateMust() *storageTableState {
	if st.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", st.Type(), st.LocalName()))
	}
	return st.state
}

// StorageTableArgs contains the configurations for azurerm_storage_table.
type StorageTableArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// StorageAccountName: string, required
	StorageAccountName terra.StringValue `hcl:"storage_account_name,attr" validate:"required"`
	// Acl: min=0
	Acl []storagetable.Acl `hcl:"acl,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *storagetable.Timeouts `hcl:"timeouts,block"`
}
type storageTableAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_storage_table.
func (st storageTableAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_storage_table.
func (st storageTableAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("name"))
}

// StorageAccountName returns a reference to field storage_account_name of azurerm_storage_table.
func (st storageTableAttributes) StorageAccountName() terra.StringValue {
	return terra.ReferenceAsString(st.ref.Append("storage_account_name"))
}

func (st storageTableAttributes) Acl() terra.SetValue[storagetable.AclAttributes] {
	return terra.ReferenceAsSet[storagetable.AclAttributes](st.ref.Append("acl"))
}

func (st storageTableAttributes) Timeouts() storagetable.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagetable.TimeoutsAttributes](st.ref.Append("timeouts"))
}

type storageTableState struct {
	Id                 string                      `json:"id"`
	Name               string                      `json:"name"`
	StorageAccountName string                      `json:"storage_account_name"`
	Acl                []storagetable.AclState     `json:"acl"`
	Timeouts           *storagetable.TimeoutsState `json:"timeouts"`
}
