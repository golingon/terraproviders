// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datadatasharedatasetkustocluster "github.com/golingon/terraproviders/azurerm/3.98.0/datadatasharedatasetkustocluster"
)

// NewDataDataShareDatasetKustoCluster creates a new instance of [DataDataShareDatasetKustoCluster].
func NewDataDataShareDatasetKustoCluster(name string, args DataDataShareDatasetKustoClusterArgs) *DataDataShareDatasetKustoCluster {
	return &DataDataShareDatasetKustoCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDataShareDatasetKustoCluster)(nil)

// DataDataShareDatasetKustoCluster represents the Terraform data resource azurerm_data_share_dataset_kusto_cluster.
type DataDataShareDatasetKustoCluster struct {
	Name string
	Args DataDataShareDatasetKustoClusterArgs
}

// DataSource returns the Terraform object type for [DataDataShareDatasetKustoCluster].
func (dsdkc *DataDataShareDatasetKustoCluster) DataSource() string {
	return "azurerm_data_share_dataset_kusto_cluster"
}

// LocalName returns the local name for [DataDataShareDatasetKustoCluster].
func (dsdkc *DataDataShareDatasetKustoCluster) LocalName() string {
	return dsdkc.Name
}

// Configuration returns the configuration (args) for [DataDataShareDatasetKustoCluster].
func (dsdkc *DataDataShareDatasetKustoCluster) Configuration() interface{} {
	return dsdkc.Args
}

// Attributes returns the attributes for [DataDataShareDatasetKustoCluster].
func (dsdkc *DataDataShareDatasetKustoCluster) Attributes() dataDataShareDatasetKustoClusterAttributes {
	return dataDataShareDatasetKustoClusterAttributes{ref: terra.ReferenceDataResource(dsdkc)}
}

// DataDataShareDatasetKustoClusterArgs contains the configurations for azurerm_data_share_dataset_kusto_cluster.
type DataDataShareDatasetKustoClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareId: string, required
	ShareId terra.StringValue `hcl:"share_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadatasharedatasetkustocluster.Timeouts `hcl:"timeouts,block"`
}
type dataDataShareDatasetKustoClusterAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("id"))
}

// KustoClusterId returns a reference to field kusto_cluster_id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) KustoClusterId() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("kusto_cluster_id"))
}

// KustoClusterLocation returns a reference to field kusto_cluster_location of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) KustoClusterLocation() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("kusto_cluster_location"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("name"))
}

// ShareId returns a reference to field share_id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataDataShareDatasetKustoClusterAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("share_id"))
}

func (dsdkc dataDataShareDatasetKustoClusterAttributes) Timeouts() datadatasharedatasetkustocluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadatasharedatasetkustocluster.TimeoutsAttributes](dsdkc.ref.Append("timeouts"))
}
