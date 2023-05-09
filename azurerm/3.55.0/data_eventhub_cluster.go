// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataeventhubcluster "github.com/golingon/terraproviders/azurerm/3.55.0/dataeventhubcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataEventhubCluster creates a new instance of [DataEventhubCluster].
func NewDataEventhubCluster(name string, args DataEventhubClusterArgs) *DataEventhubCluster {
	return &DataEventhubCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataEventhubCluster)(nil)

// DataEventhubCluster represents the Terraform data resource azurerm_eventhub_cluster.
type DataEventhubCluster struct {
	Name string
	Args DataEventhubClusterArgs
}

// DataSource returns the Terraform object type for [DataEventhubCluster].
func (ec *DataEventhubCluster) DataSource() string {
	return "azurerm_eventhub_cluster"
}

// LocalName returns the local name for [DataEventhubCluster].
func (ec *DataEventhubCluster) LocalName() string {
	return ec.Name
}

// Configuration returns the configuration (args) for [DataEventhubCluster].
func (ec *DataEventhubCluster) Configuration() interface{} {
	return ec.Args
}

// Attributes returns the attributes for [DataEventhubCluster].
func (ec *DataEventhubCluster) Attributes() dataEventhubClusterAttributes {
	return dataEventhubClusterAttributes{ref: terra.ReferenceDataResource(ec)}
}

// DataEventhubClusterArgs contains the configurations for azurerm_eventhub_cluster.
type DataEventhubClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *dataeventhubcluster.Timeouts `hcl:"timeouts,block"`
}
type dataEventhubClusterAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_eventhub_cluster.
func (ec dataEventhubClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_eventhub_cluster.
func (ec dataEventhubClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_eventhub_cluster.
func (ec dataEventhubClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_eventhub_cluster.
func (ec dataEventhubClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_eventhub_cluster.
func (ec dataEventhubClusterAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("sku_name"))
}

func (ec dataEventhubClusterAttributes) Timeouts() dataeventhubcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataeventhubcluster.TimeoutsAttributes](ec.ref.Append("timeouts"))
}
