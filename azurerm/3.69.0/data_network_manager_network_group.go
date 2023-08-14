// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetworkmanagernetworkgroup "github.com/golingon/terraproviders/azurerm/3.69.0/datanetworkmanagernetworkgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkManagerNetworkGroup creates a new instance of [DataNetworkManagerNetworkGroup].
func NewDataNetworkManagerNetworkGroup(name string, args DataNetworkManagerNetworkGroupArgs) *DataNetworkManagerNetworkGroup {
	return &DataNetworkManagerNetworkGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkManagerNetworkGroup)(nil)

// DataNetworkManagerNetworkGroup represents the Terraform data resource azurerm_network_manager_network_group.
type DataNetworkManagerNetworkGroup struct {
	Name string
	Args DataNetworkManagerNetworkGroupArgs
}

// DataSource returns the Terraform object type for [DataNetworkManagerNetworkGroup].
func (nmng *DataNetworkManagerNetworkGroup) DataSource() string {
	return "azurerm_network_manager_network_group"
}

// LocalName returns the local name for [DataNetworkManagerNetworkGroup].
func (nmng *DataNetworkManagerNetworkGroup) LocalName() string {
	return nmng.Name
}

// Configuration returns the configuration (args) for [DataNetworkManagerNetworkGroup].
func (nmng *DataNetworkManagerNetworkGroup) Configuration() interface{} {
	return nmng.Args
}

// Attributes returns the attributes for [DataNetworkManagerNetworkGroup].
func (nmng *DataNetworkManagerNetworkGroup) Attributes() dataNetworkManagerNetworkGroupAttributes {
	return dataNetworkManagerNetworkGroupAttributes{ref: terra.ReferenceDataResource(nmng)}
}

// DataNetworkManagerNetworkGroupArgs contains the configurations for azurerm_network_manager_network_group.
type DataNetworkManagerNetworkGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datanetworkmanagernetworkgroup.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkManagerNetworkGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_network_manager_network_group.
func (nmng dataNetworkManagerNetworkGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_network_group.
func (nmng dataNetworkManagerNetworkGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_network_group.
func (nmng dataNetworkManagerNetworkGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_network_group.
func (nmng dataNetworkManagerNetworkGroupAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("network_manager_id"))
}

func (nmng dataNetworkManagerNetworkGroupAttributes) Timeouts() datanetworkmanagernetworkgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkmanagernetworkgroup.TimeoutsAttributes](nmng.ref.Append("timeouts"))
}