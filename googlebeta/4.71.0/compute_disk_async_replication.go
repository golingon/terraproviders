// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computediskasyncreplication "github.com/golingon/terraproviders/googlebeta/4.71.0/computediskasyncreplication"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeDiskAsyncReplication creates a new instance of [ComputeDiskAsyncReplication].
func NewComputeDiskAsyncReplication(name string, args ComputeDiskAsyncReplicationArgs) *ComputeDiskAsyncReplication {
	return &ComputeDiskAsyncReplication{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeDiskAsyncReplication)(nil)

// ComputeDiskAsyncReplication represents the Terraform resource google_compute_disk_async_replication.
type ComputeDiskAsyncReplication struct {
	Name      string
	Args      ComputeDiskAsyncReplicationArgs
	state     *computeDiskAsyncReplicationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) Type() string {
	return "google_compute_disk_async_replication"
}

// LocalName returns the local name for [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) LocalName() string {
	return cdar.Name
}

// Configuration returns the configuration (args) for [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) Configuration() interface{} {
	return cdar.Args
}

// DependOn is used for other resources to depend on [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) DependOn() terra.Reference {
	return terra.ReferenceResource(cdar)
}

// Dependencies returns the list of resources [ComputeDiskAsyncReplication] depends_on.
func (cdar *ComputeDiskAsyncReplication) Dependencies() terra.Dependencies {
	return cdar.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) LifecycleManagement() *terra.Lifecycle {
	return cdar.Lifecycle
}

// Attributes returns the attributes for [ComputeDiskAsyncReplication].
func (cdar *ComputeDiskAsyncReplication) Attributes() computeDiskAsyncReplicationAttributes {
	return computeDiskAsyncReplicationAttributes{ref: terra.ReferenceResource(cdar)}
}

// ImportState imports the given attribute values into [ComputeDiskAsyncReplication]'s state.
func (cdar *ComputeDiskAsyncReplication) ImportState(av io.Reader) error {
	cdar.state = &computeDiskAsyncReplicationState{}
	if err := json.NewDecoder(av).Decode(cdar.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdar.Type(), cdar.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeDiskAsyncReplication] has state.
func (cdar *ComputeDiskAsyncReplication) State() (*computeDiskAsyncReplicationState, bool) {
	return cdar.state, cdar.state != nil
}

// StateMust returns the state for [ComputeDiskAsyncReplication]. Panics if the state is nil.
func (cdar *ComputeDiskAsyncReplication) StateMust() *computeDiskAsyncReplicationState {
	if cdar.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdar.Type(), cdar.LocalName()))
	}
	return cdar.state
}

// ComputeDiskAsyncReplicationArgs contains the configurations for google_compute_disk_async_replication.
type ComputeDiskAsyncReplicationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PrimaryDisk: string, required
	PrimaryDisk terra.StringValue `hcl:"primary_disk,attr" validate:"required"`
	// SecondaryDisk: required
	SecondaryDisk *computediskasyncreplication.SecondaryDisk `hcl:"secondary_disk,block" validate:"required"`
	// Timeouts: optional
	Timeouts *computediskasyncreplication.Timeouts `hcl:"timeouts,block"`
}
type computeDiskAsyncReplicationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_disk_async_replication.
func (cdar computeDiskAsyncReplicationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdar.ref.Append("id"))
}

// PrimaryDisk returns a reference to field primary_disk of google_compute_disk_async_replication.
func (cdar computeDiskAsyncReplicationAttributes) PrimaryDisk() terra.StringValue {
	return terra.ReferenceAsString(cdar.ref.Append("primary_disk"))
}

func (cdar computeDiskAsyncReplicationAttributes) SecondaryDisk() terra.ListValue[computediskasyncreplication.SecondaryDiskAttributes] {
	return terra.ReferenceAsList[computediskasyncreplication.SecondaryDiskAttributes](cdar.ref.Append("secondary_disk"))
}

func (cdar computeDiskAsyncReplicationAttributes) Timeouts() computediskasyncreplication.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computediskasyncreplication.TimeoutsAttributes](cdar.ref.Append("timeouts"))
}

type computeDiskAsyncReplicationState struct {
	Id            string                                           `json:"id"`
	PrimaryDisk   string                                           `json:"primary_disk"`
	SecondaryDisk []computediskasyncreplication.SecondaryDiskState `json:"secondary_disk"`
	Timeouts      *computediskasyncreplication.TimeoutsState       `json:"timeouts"`
}
