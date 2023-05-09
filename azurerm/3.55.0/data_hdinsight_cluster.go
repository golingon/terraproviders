// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datahdinsightcluster "github.com/golingon/terraproviders/azurerm/3.55.0/datahdinsightcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataHdinsightCluster creates a new instance of [DataHdinsightCluster].
func NewDataHdinsightCluster(name string, args DataHdinsightClusterArgs) *DataHdinsightCluster {
	return &DataHdinsightCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataHdinsightCluster)(nil)

// DataHdinsightCluster represents the Terraform data resource azurerm_hdinsight_cluster.
type DataHdinsightCluster struct {
	Name string
	Args DataHdinsightClusterArgs
}

// DataSource returns the Terraform object type for [DataHdinsightCluster].
func (hc *DataHdinsightCluster) DataSource() string {
	return "azurerm_hdinsight_cluster"
}

// LocalName returns the local name for [DataHdinsightCluster].
func (hc *DataHdinsightCluster) LocalName() string {
	return hc.Name
}

// Configuration returns the configuration (args) for [DataHdinsightCluster].
func (hc *DataHdinsightCluster) Configuration() interface{} {
	return hc.Args
}

// Attributes returns the attributes for [DataHdinsightCluster].
func (hc *DataHdinsightCluster) Attributes() dataHdinsightClusterAttributes {
	return dataHdinsightClusterAttributes{ref: terra.ReferenceDataResource(hc)}
}

// DataHdinsightClusterArgs contains the configurations for azurerm_hdinsight_cluster.
type DataHdinsightClusterArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Gateway: min=0
	Gateway []datahdinsightcluster.Gateway `hcl:"gateway,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datahdinsightcluster.Timeouts `hcl:"timeouts,block"`
}
type dataHdinsightClusterAttributes struct {
	ref terra.Reference
}

// ClusterVersion returns a reference to field cluster_version of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("cluster_version"))
}

// ComponentVersions returns a reference to field component_versions of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) ComponentVersions() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("component_versions"))
}

// EdgeSshEndpoint returns a reference to field edge_ssh_endpoint of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) EdgeSshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("edge_ssh_endpoint"))
}

// HttpsEndpoint returns a reference to field https_endpoint of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) HttpsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("https_endpoint"))
}

// Id returns a reference to field id of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("id"))
}

// KafkaRestProxyEndpoint returns a reference to field kafka_rest_proxy_endpoint of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) KafkaRestProxyEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("kafka_rest_proxy_endpoint"))
}

// Kind returns a reference to field kind of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("kind"))
}

// Location returns a reference to field location of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("resource_group_name"))
}

// SshEndpoint returns a reference to field ssh_endpoint of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) SshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("ssh_endpoint"))
}

// Tags returns a reference to field tags of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hc.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("tier"))
}

// TlsMinVersion returns a reference to field tls_min_version of azurerm_hdinsight_cluster.
func (hc dataHdinsightClusterAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(hc.ref.Append("tls_min_version"))
}

func (hc dataHdinsightClusterAttributes) Gateway() terra.ListValue[datahdinsightcluster.GatewayAttributes] {
	return terra.ReferenceAsList[datahdinsightcluster.GatewayAttributes](hc.ref.Append("gateway"))
}

func (hc dataHdinsightClusterAttributes) Timeouts() datahdinsightcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datahdinsightcluster.TimeoutsAttributes](hc.ref.Append("timeouts"))
}
