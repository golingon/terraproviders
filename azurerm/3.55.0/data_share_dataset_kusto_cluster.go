// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datasharedatasetkustocluster "github.com/golingon/terraproviders/azurerm/3.55.0/datasharedatasetkustocluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataShareDatasetKustoCluster creates a new instance of [DataShareDatasetKustoCluster].
func NewDataShareDatasetKustoCluster(name string, args DataShareDatasetKustoClusterArgs) *DataShareDatasetKustoCluster {
	return &DataShareDatasetKustoCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataShareDatasetKustoCluster)(nil)

// DataShareDatasetKustoCluster represents the Terraform resource azurerm_data_share_dataset_kusto_cluster.
type DataShareDatasetKustoCluster struct {
	Name      string
	Args      DataShareDatasetKustoClusterArgs
	state     *dataShareDatasetKustoClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) Type() string {
	return "azurerm_data_share_dataset_kusto_cluster"
}

// LocalName returns the local name for [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) LocalName() string {
	return dsdkc.Name
}

// Configuration returns the configuration (args) for [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) Configuration() interface{} {
	return dsdkc.Args
}

// DependOn is used for other resources to depend on [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(dsdkc)
}

// Dependencies returns the list of resources [DataShareDatasetKustoCluster] depends_on.
func (dsdkc *DataShareDatasetKustoCluster) Dependencies() terra.Dependencies {
	return dsdkc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) LifecycleManagement() *terra.Lifecycle {
	return dsdkc.Lifecycle
}

// Attributes returns the attributes for [DataShareDatasetKustoCluster].
func (dsdkc *DataShareDatasetKustoCluster) Attributes() dataShareDatasetKustoClusterAttributes {
	return dataShareDatasetKustoClusterAttributes{ref: terra.ReferenceResource(dsdkc)}
}

// ImportState imports the given attribute values into [DataShareDatasetKustoCluster]'s state.
func (dsdkc *DataShareDatasetKustoCluster) ImportState(av io.Reader) error {
	dsdkc.state = &dataShareDatasetKustoClusterState{}
	if err := json.NewDecoder(av).Decode(dsdkc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dsdkc.Type(), dsdkc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataShareDatasetKustoCluster] has state.
func (dsdkc *DataShareDatasetKustoCluster) State() (*dataShareDatasetKustoClusterState, bool) {
	return dsdkc.state, dsdkc.state != nil
}

// StateMust returns the state for [DataShareDatasetKustoCluster]. Panics if the state is nil.
func (dsdkc *DataShareDatasetKustoCluster) StateMust() *dataShareDatasetKustoClusterState {
	if dsdkc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dsdkc.Type(), dsdkc.LocalName()))
	}
	return dsdkc.state
}

// DataShareDatasetKustoClusterArgs contains the configurations for azurerm_data_share_dataset_kusto_cluster.
type DataShareDatasetKustoClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KustoClusterId: string, required
	KustoClusterId terra.StringValue `hcl:"kusto_cluster_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ShareId: string, required
	ShareId terra.StringValue `hcl:"share_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasharedatasetkustocluster.Timeouts `hcl:"timeouts,block"`
}
type dataShareDatasetKustoClusterAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("display_name"))
}

// Id returns a reference to field id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("id"))
}

// KustoClusterId returns a reference to field kusto_cluster_id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) KustoClusterId() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("kusto_cluster_id"))
}

// KustoClusterLocation returns a reference to field kusto_cluster_location of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) KustoClusterLocation() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("kusto_cluster_location"))
}

// Name returns a reference to field name of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("name"))
}

// ShareId returns a reference to field share_id of azurerm_data_share_dataset_kusto_cluster.
func (dsdkc dataShareDatasetKustoClusterAttributes) ShareId() terra.StringValue {
	return terra.ReferenceAsString(dsdkc.ref.Append("share_id"))
}

func (dsdkc dataShareDatasetKustoClusterAttributes) Timeouts() datasharedatasetkustocluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasharedatasetkustocluster.TimeoutsAttributes](dsdkc.ref.Append("timeouts"))
}

type dataShareDatasetKustoClusterState struct {
	DisplayName          string                                      `json:"display_name"`
	Id                   string                                      `json:"id"`
	KustoClusterId       string                                      `json:"kusto_cluster_id"`
	KustoClusterLocation string                                      `json:"kusto_cluster_location"`
	Name                 string                                      `json:"name"`
	ShareId              string                                      `json:"share_id"`
	Timeouts             *datasharedatasetkustocluster.TimeoutsState `json:"timeouts"`
}
