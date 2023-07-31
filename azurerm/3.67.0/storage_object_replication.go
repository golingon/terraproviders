// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	storageobjectreplication "github.com/golingon/terraproviders/azurerm/3.67.0/storageobjectreplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageObjectReplication creates a new instance of [StorageObjectReplication].
func NewStorageObjectReplication(name string, args StorageObjectReplicationArgs) *StorageObjectReplication {
	return &StorageObjectReplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageObjectReplication)(nil)

// StorageObjectReplication represents the Terraform resource azurerm_storage_object_replication.
type StorageObjectReplication struct {
	Name      string
	Args      StorageObjectReplicationArgs
	state     *storageObjectReplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageObjectReplication].
func (sor *StorageObjectReplication) Type() string {
	return "azurerm_storage_object_replication"
}

// LocalName returns the local name for [StorageObjectReplication].
func (sor *StorageObjectReplication) LocalName() string {
	return sor.Name
}

// Configuration returns the configuration (args) for [StorageObjectReplication].
func (sor *StorageObjectReplication) Configuration() interface{} {
	return sor.Args
}

// DependOn is used for other resources to depend on [StorageObjectReplication].
func (sor *StorageObjectReplication) DependOn() terra.Reference {
	return terra.ReferenceResource(sor)
}

// Dependencies returns the list of resources [StorageObjectReplication] depends_on.
func (sor *StorageObjectReplication) Dependencies() terra.Dependencies {
	return sor.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageObjectReplication].
func (sor *StorageObjectReplication) LifecycleManagement() *terra.Lifecycle {
	return sor.Lifecycle
}

// Attributes returns the attributes for [StorageObjectReplication].
func (sor *StorageObjectReplication) Attributes() storageObjectReplicationAttributes {
	return storageObjectReplicationAttributes{ref: terra.ReferenceResource(sor)}
}

// ImportState imports the given attribute values into [StorageObjectReplication]'s state.
func (sor *StorageObjectReplication) ImportState(av io.Reader) error {
	sor.state = &storageObjectReplicationState{}
	if err := json.NewDecoder(av).Decode(sor.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sor.Type(), sor.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageObjectReplication] has state.
func (sor *StorageObjectReplication) State() (*storageObjectReplicationState, bool) {
	return sor.state, sor.state != nil
}

// StateMust returns the state for [StorageObjectReplication]. Panics if the state is nil.
func (sor *StorageObjectReplication) StateMust() *storageObjectReplicationState {
	if sor.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sor.Type(), sor.LocalName()))
	}
	return sor.state
}

// StorageObjectReplicationArgs contains the configurations for azurerm_storage_object_replication.
type StorageObjectReplicationArgs struct {
	// DestinationStorageAccountId: string, required
	DestinationStorageAccountId terra.StringValue `hcl:"destination_storage_account_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SourceStorageAccountId: string, required
	SourceStorageAccountId terra.StringValue `hcl:"source_storage_account_id,attr" validate:"required"`
	// Rules: min=1
	Rules []storageobjectreplication.Rules `hcl:"rules,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *storageobjectreplication.Timeouts `hcl:"timeouts,block"`
}
type storageObjectReplicationAttributes struct {
	ref terra.Reference
}

// DestinationObjectReplicationId returns a reference to field destination_object_replication_id of azurerm_storage_object_replication.
func (sor storageObjectReplicationAttributes) DestinationObjectReplicationId() terra.StringValue {
	return terra.ReferenceAsString(sor.ref.Append("destination_object_replication_id"))
}

// DestinationStorageAccountId returns a reference to field destination_storage_account_id of azurerm_storage_object_replication.
func (sor storageObjectReplicationAttributes) DestinationStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sor.ref.Append("destination_storage_account_id"))
}

// Id returns a reference to field id of azurerm_storage_object_replication.
func (sor storageObjectReplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sor.ref.Append("id"))
}

// SourceObjectReplicationId returns a reference to field source_object_replication_id of azurerm_storage_object_replication.
func (sor storageObjectReplicationAttributes) SourceObjectReplicationId() terra.StringValue {
	return terra.ReferenceAsString(sor.ref.Append("source_object_replication_id"))
}

// SourceStorageAccountId returns a reference to field source_storage_account_id of azurerm_storage_object_replication.
func (sor storageObjectReplicationAttributes) SourceStorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(sor.ref.Append("source_storage_account_id"))
}

func (sor storageObjectReplicationAttributes) Rules() terra.SetValue[storageobjectreplication.RulesAttributes] {
	return terra.ReferenceAsSet[storageobjectreplication.RulesAttributes](sor.ref.Append("rules"))
}

func (sor storageObjectReplicationAttributes) Timeouts() storageobjectreplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storageobjectreplication.TimeoutsAttributes](sor.ref.Append("timeouts"))
}

type storageObjectReplicationState struct {
	DestinationObjectReplicationId string                                  `json:"destination_object_replication_id"`
	DestinationStorageAccountId    string                                  `json:"destination_storage_account_id"`
	Id                             string                                  `json:"id"`
	SourceObjectReplicationId      string                                  `json:"source_object_replication_id"`
	SourceStorageAccountId         string                                  `json:"source_storage_account_id"`
	Rules                          []storageobjectreplication.RulesState   `json:"rules"`
	Timeouts                       *storageobjectreplication.TimeoutsState `json:"timeouts"`
}
