// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	netappsnapshot "github.com/golingon/terraproviders/azurerm/3.65.0/netappsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetappSnapshot creates a new instance of [NetappSnapshot].
func NewNetappSnapshot(name string, args NetappSnapshotArgs) *NetappSnapshot {
	return &NetappSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetappSnapshot)(nil)

// NetappSnapshot represents the Terraform resource azurerm_netapp_snapshot.
type NetappSnapshot struct {
	Name      string
	Args      NetappSnapshotArgs
	state     *netappSnapshotState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetappSnapshot].
func (ns *NetappSnapshot) Type() string {
	return "azurerm_netapp_snapshot"
}

// LocalName returns the local name for [NetappSnapshot].
func (ns *NetappSnapshot) LocalName() string {
	return ns.Name
}

// Configuration returns the configuration (args) for [NetappSnapshot].
func (ns *NetappSnapshot) Configuration() interface{} {
	return ns.Args
}

// DependOn is used for other resources to depend on [NetappSnapshot].
func (ns *NetappSnapshot) DependOn() terra.Reference {
	return terra.ReferenceResource(ns)
}

// Dependencies returns the list of resources [NetappSnapshot] depends_on.
func (ns *NetappSnapshot) Dependencies() terra.Dependencies {
	return ns.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetappSnapshot].
func (ns *NetappSnapshot) LifecycleManagement() *terra.Lifecycle {
	return ns.Lifecycle
}

// Attributes returns the attributes for [NetappSnapshot].
func (ns *NetappSnapshot) Attributes() netappSnapshotAttributes {
	return netappSnapshotAttributes{ref: terra.ReferenceResource(ns)}
}

// ImportState imports the given attribute values into [NetappSnapshot]'s state.
func (ns *NetappSnapshot) ImportState(av io.Reader) error {
	ns.state = &netappSnapshotState{}
	if err := json.NewDecoder(av).Decode(ns.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ns.Type(), ns.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetappSnapshot] has state.
func (ns *NetappSnapshot) State() (*netappSnapshotState, bool) {
	return ns.state, ns.state != nil
}

// StateMust returns the state for [NetappSnapshot]. Panics if the state is nil.
func (ns *NetappSnapshot) StateMust() *netappSnapshotState {
	if ns.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ns.Type(), ns.LocalName()))
	}
	return ns.state
}

// NetappSnapshotArgs contains the configurations for azurerm_netapp_snapshot.
type NetappSnapshotArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PoolName: string, required
	PoolName terra.StringValue `hcl:"pool_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VolumeName: string, required
	VolumeName terra.StringValue `hcl:"volume_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *netappsnapshot.Timeouts `hcl:"timeouts,block"`
}
type netappSnapshotAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("name"))
}

// PoolName returns a reference to field pool_name of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) PoolName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("pool_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("resource_group_name"))
}

// VolumeName returns a reference to field volume_name of azurerm_netapp_snapshot.
func (ns netappSnapshotAttributes) VolumeName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("volume_name"))
}

func (ns netappSnapshotAttributes) Timeouts() netappsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[netappsnapshot.TimeoutsAttributes](ns.ref.Append("timeouts"))
}

type netappSnapshotState struct {
	AccountName       string                        `json:"account_name"`
	Id                string                        `json:"id"`
	Location          string                        `json:"location"`
	Name              string                        `json:"name"`
	PoolName          string                        `json:"pool_name"`
	ResourceGroupName string                        `json:"resource_group_name"`
	VolumeName        string                        `json:"volume_name"`
	Timeouts          *netappsnapshot.TimeoutsState `json:"timeouts"`
}
