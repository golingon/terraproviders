// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetappsnapshot "github.com/golingon/terraproviders/azurerm/3.68.0/datanetappsnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetappSnapshot creates a new instance of [DataNetappSnapshot].
func NewDataNetappSnapshot(name string, args DataNetappSnapshotArgs) *DataNetappSnapshot {
	return &DataNetappSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetappSnapshot)(nil)

// DataNetappSnapshot represents the Terraform data resource azurerm_netapp_snapshot.
type DataNetappSnapshot struct {
	Name string
	Args DataNetappSnapshotArgs
}

// DataSource returns the Terraform object type for [DataNetappSnapshot].
func (ns *DataNetappSnapshot) DataSource() string {
	return "azurerm_netapp_snapshot"
}

// LocalName returns the local name for [DataNetappSnapshot].
func (ns *DataNetappSnapshot) LocalName() string {
	return ns.Name
}

// Configuration returns the configuration (args) for [DataNetappSnapshot].
func (ns *DataNetappSnapshot) Configuration() interface{} {
	return ns.Args
}

// Attributes returns the attributes for [DataNetappSnapshot].
func (ns *DataNetappSnapshot) Attributes() dataNetappSnapshotAttributes {
	return dataNetappSnapshotAttributes{ref: terra.ReferenceDataResource(ns)}
}

// DataNetappSnapshotArgs contains the configurations for azurerm_netapp_snapshot.
type DataNetappSnapshotArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PoolName: string, required
	PoolName terra.StringValue `hcl:"pool_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VolumeName: string, required
	VolumeName terra.StringValue `hcl:"volume_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datanetappsnapshot.Timeouts `hcl:"timeouts,block"`
}
type dataNetappSnapshotAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("account_name"))
}

// Id returns a reference to field id of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("name"))
}

// PoolName returns a reference to field pool_name of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) PoolName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("pool_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("resource_group_name"))
}

// VolumeName returns a reference to field volume_name of azurerm_netapp_snapshot.
func (ns dataNetappSnapshotAttributes) VolumeName() terra.StringValue {
	return terra.ReferenceAsString(ns.ref.Append("volume_name"))
}

func (ns dataNetappSnapshotAttributes) Timeouts() datanetappsnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetappsnapshot.TimeoutsAttributes](ns.ref.Append("timeouts"))
}
