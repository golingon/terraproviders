// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datakustocluster "github.com/golingon/terraproviders/azurerm/3.98.0/datakustocluster"
)

// NewDataKustoCluster creates a new instance of [DataKustoCluster].
func NewDataKustoCluster(name string, args DataKustoClusterArgs) *DataKustoCluster {
	return &DataKustoCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataKustoCluster)(nil)

// DataKustoCluster represents the Terraform data resource azurerm_kusto_cluster.
type DataKustoCluster struct {
	Name string
	Args DataKustoClusterArgs
}

// DataSource returns the Terraform object type for [DataKustoCluster].
func (kc *DataKustoCluster) DataSource() string {
	return "azurerm_kusto_cluster"
}

// LocalName returns the local name for [DataKustoCluster].
func (kc *DataKustoCluster) LocalName() string {
	return kc.Name
}

// Configuration returns the configuration (args) for [DataKustoCluster].
func (kc *DataKustoCluster) Configuration() interface{} {
	return kc.Args
}

// Attributes returns the attributes for [DataKustoCluster].
func (kc *DataKustoCluster) Attributes() dataKustoClusterAttributes {
	return dataKustoClusterAttributes{ref: terra.ReferenceDataResource(kc)}
}

// DataKustoClusterArgs contains the configurations for azurerm_kusto_cluster.
type DataKustoClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datakustocluster.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datakustocluster.Timeouts `hcl:"timeouts,block"`
}
type dataKustoClusterAttributes struct {
	ref terra.Reference
}

// DataIngestionUri returns a reference to field data_ingestion_uri of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) DataIngestionUri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("data_ingestion_uri"))
}

// Id returns a reference to field id of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](kc.ref.Append("tags"))
}

// Uri returns a reference to field uri of azurerm_kusto_cluster.
func (kc dataKustoClusterAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(kc.ref.Append("uri"))
}

func (kc dataKustoClusterAttributes) Identity() terra.ListValue[datakustocluster.IdentityAttributes] {
	return terra.ReferenceAsList[datakustocluster.IdentityAttributes](kc.ref.Append("identity"))
}

func (kc dataKustoClusterAttributes) Timeouts() datakustocluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datakustocluster.TimeoutsAttributes](kc.ref.Append("timeouts"))
}
