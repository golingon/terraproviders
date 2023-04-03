// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetworkwatcher "github.com/golingon/terraproviders/azurerm/3.49.0/datanetworkwatcher"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkWatcher creates a new instance of [DataNetworkWatcher].
func NewDataNetworkWatcher(name string, args DataNetworkWatcherArgs) *DataNetworkWatcher {
	return &DataNetworkWatcher{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkWatcher)(nil)

// DataNetworkWatcher represents the Terraform data resource azurerm_network_watcher.
type DataNetworkWatcher struct {
	Name string
	Args DataNetworkWatcherArgs
}

// DataSource returns the Terraform object type for [DataNetworkWatcher].
func (nw *DataNetworkWatcher) DataSource() string {
	return "azurerm_network_watcher"
}

// LocalName returns the local name for [DataNetworkWatcher].
func (nw *DataNetworkWatcher) LocalName() string {
	return nw.Name
}

// Configuration returns the configuration (args) for [DataNetworkWatcher].
func (nw *DataNetworkWatcher) Configuration() interface{} {
	return nw.Args
}

// Attributes returns the attributes for [DataNetworkWatcher].
func (nw *DataNetworkWatcher) Attributes() dataNetworkWatcherAttributes {
	return dataNetworkWatcherAttributes{ref: terra.ReferenceDataResource(nw)}
}

// DataNetworkWatcherArgs contains the configurations for azurerm_network_watcher.
type DataNetworkWatcherArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datanetworkwatcher.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkWatcherAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_watcher.
func (nw dataNetworkWatcherAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_watcher.
func (nw dataNetworkWatcherAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_watcher.
func (nw dataNetworkWatcherAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_watcher.
func (nw dataNetworkWatcherAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_watcher.
func (nw dataNetworkWatcherAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nw.ref.Append("tags"))
}

func (nw dataNetworkWatcherAttributes) Timeouts() datanetworkwatcher.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkwatcher.TimeoutsAttributes](nw.ref.Append("timeouts"))
}
