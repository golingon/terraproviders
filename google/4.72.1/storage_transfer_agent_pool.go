// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	storagetransferagentpool "github.com/golingon/terraproviders/google/4.72.1/storagetransferagentpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageTransferAgentPool creates a new instance of [StorageTransferAgentPool].
func NewStorageTransferAgentPool(name string, args StorageTransferAgentPoolArgs) *StorageTransferAgentPool {
	return &StorageTransferAgentPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageTransferAgentPool)(nil)

// StorageTransferAgentPool represents the Terraform resource google_storage_transfer_agent_pool.
type StorageTransferAgentPool struct {
	Name      string
	Args      StorageTransferAgentPoolArgs
	state     *storageTransferAgentPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) Type() string {
	return "google_storage_transfer_agent_pool"
}

// LocalName returns the local name for [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) LocalName() string {
	return stap.Name
}

// Configuration returns the configuration (args) for [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) Configuration() interface{} {
	return stap.Args
}

// DependOn is used for other resources to depend on [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) DependOn() terra.Reference {
	return terra.ReferenceResource(stap)
}

// Dependencies returns the list of resources [StorageTransferAgentPool] depends_on.
func (stap *StorageTransferAgentPool) Dependencies() terra.Dependencies {
	return stap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) LifecycleManagement() *terra.Lifecycle {
	return stap.Lifecycle
}

// Attributes returns the attributes for [StorageTransferAgentPool].
func (stap *StorageTransferAgentPool) Attributes() storageTransferAgentPoolAttributes {
	return storageTransferAgentPoolAttributes{ref: terra.ReferenceResource(stap)}
}

// ImportState imports the given attribute values into [StorageTransferAgentPool]'s state.
func (stap *StorageTransferAgentPool) ImportState(av io.Reader) error {
	stap.state = &storageTransferAgentPoolState{}
	if err := json.NewDecoder(av).Decode(stap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", stap.Type(), stap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageTransferAgentPool] has state.
func (stap *StorageTransferAgentPool) State() (*storageTransferAgentPoolState, bool) {
	return stap.state, stap.state != nil
}

// StateMust returns the state for [StorageTransferAgentPool]. Panics if the state is nil.
func (stap *StorageTransferAgentPool) StateMust() *storageTransferAgentPoolState {
	if stap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", stap.Type(), stap.LocalName()))
	}
	return stap.state
}

// StorageTransferAgentPoolArgs contains the configurations for google_storage_transfer_agent_pool.
type StorageTransferAgentPoolArgs struct {
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BandwidthLimit: optional
	BandwidthLimit *storagetransferagentpool.BandwidthLimit `hcl:"bandwidth_limit,block"`
	// Timeouts: optional
	Timeouts *storagetransferagentpool.Timeouts `hcl:"timeouts,block"`
}
type storageTransferAgentPoolAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_storage_transfer_agent_pool.
func (stap storageTransferAgentPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(stap.ref.Append("display_name"))
}

// Id returns a reference to field id of google_storage_transfer_agent_pool.
func (stap storageTransferAgentPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(stap.ref.Append("id"))
}

// Name returns a reference to field name of google_storage_transfer_agent_pool.
func (stap storageTransferAgentPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(stap.ref.Append("name"))
}

// Project returns a reference to field project of google_storage_transfer_agent_pool.
func (stap storageTransferAgentPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(stap.ref.Append("project"))
}

// State returns a reference to field state of google_storage_transfer_agent_pool.
func (stap storageTransferAgentPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(stap.ref.Append("state"))
}

func (stap storageTransferAgentPoolAttributes) BandwidthLimit() terra.ListValue[storagetransferagentpool.BandwidthLimitAttributes] {
	return terra.ReferenceAsList[storagetransferagentpool.BandwidthLimitAttributes](stap.ref.Append("bandwidth_limit"))
}

func (stap storageTransferAgentPoolAttributes) Timeouts() storagetransferagentpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagetransferagentpool.TimeoutsAttributes](stap.ref.Append("timeouts"))
}

type storageTransferAgentPoolState struct {
	DisplayName    string                                         `json:"display_name"`
	Id             string                                         `json:"id"`
	Name           string                                         `json:"name"`
	Project        string                                         `json:"project"`
	State          string                                         `json:"state"`
	BandwidthLimit []storagetransferagentpool.BandwidthLimitState `json:"bandwidth_limit"`
	Timeouts       *storagetransferagentpool.TimeoutsState        `json:"timeouts"`
}
