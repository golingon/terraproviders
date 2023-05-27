// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	hdinsightkafkacluster "github.com/golingon/terraproviders/azurerm/3.58.0/hdinsightkafkacluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHdinsightKafkaCluster creates a new instance of [HdinsightKafkaCluster].
func NewHdinsightKafkaCluster(name string, args HdinsightKafkaClusterArgs) *HdinsightKafkaCluster {
	return &HdinsightKafkaCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HdinsightKafkaCluster)(nil)

// HdinsightKafkaCluster represents the Terraform resource azurerm_hdinsight_kafka_cluster.
type HdinsightKafkaCluster struct {
	Name      string
	Args      HdinsightKafkaClusterArgs
	state     *hdinsightKafkaClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) Type() string {
	return "azurerm_hdinsight_kafka_cluster"
}

// LocalName returns the local name for [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) LocalName() string {
	return hkc.Name
}

// Configuration returns the configuration (args) for [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) Configuration() interface{} {
	return hkc.Args
}

// DependOn is used for other resources to depend on [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(hkc)
}

// Dependencies returns the list of resources [HdinsightKafkaCluster] depends_on.
func (hkc *HdinsightKafkaCluster) Dependencies() terra.Dependencies {
	return hkc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) LifecycleManagement() *terra.Lifecycle {
	return hkc.Lifecycle
}

// Attributes returns the attributes for [HdinsightKafkaCluster].
func (hkc *HdinsightKafkaCluster) Attributes() hdinsightKafkaClusterAttributes {
	return hdinsightKafkaClusterAttributes{ref: terra.ReferenceResource(hkc)}
}

// ImportState imports the given attribute values into [HdinsightKafkaCluster]'s state.
func (hkc *HdinsightKafkaCluster) ImportState(av io.Reader) error {
	hkc.state = &hdinsightKafkaClusterState{}
	if err := json.NewDecoder(av).Decode(hkc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hkc.Type(), hkc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HdinsightKafkaCluster] has state.
func (hkc *HdinsightKafkaCluster) State() (*hdinsightKafkaClusterState, bool) {
	return hkc.state, hkc.state != nil
}

// StateMust returns the state for [HdinsightKafkaCluster]. Panics if the state is nil.
func (hkc *HdinsightKafkaCluster) StateMust() *hdinsightKafkaClusterState {
	if hkc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hkc.Type(), hkc.LocalName()))
	}
	return hkc.state
}

// HdinsightKafkaClusterArgs contains the configurations for azurerm_hdinsight_kafka_cluster.
type HdinsightKafkaClusterArgs struct {
	// ClusterVersion: string, required
	ClusterVersion terra.StringValue `hcl:"cluster_version,attr" validate:"required"`
	// EncryptionInTransitEnabled: bool, optional
	EncryptionInTransitEnabled terra.BoolValue `hcl:"encryption_in_transit_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Tier: string, required
	Tier terra.StringValue `hcl:"tier,attr" validate:"required"`
	// TlsMinVersion: string, optional
	TlsMinVersion terra.StringValue `hcl:"tls_min_version,attr"`
	// ComponentVersion: required
	ComponentVersion *hdinsightkafkacluster.ComponentVersion `hcl:"component_version,block" validate:"required"`
	// ComputeIsolation: optional
	ComputeIsolation *hdinsightkafkacluster.ComputeIsolation `hcl:"compute_isolation,block"`
	// DiskEncryption: min=0
	DiskEncryption []hdinsightkafkacluster.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extension: optional
	Extension *hdinsightkafkacluster.Extension `hcl:"extension,block"`
	// Gateway: required
	Gateway *hdinsightkafkacluster.Gateway `hcl:"gateway,block" validate:"required"`
	// Metastores: optional
	Metastores *hdinsightkafkacluster.Metastores `hcl:"metastores,block"`
	// Monitor: optional
	Monitor *hdinsightkafkacluster.Monitor `hcl:"monitor,block"`
	// Network: optional
	Network *hdinsightkafkacluster.Network `hcl:"network,block"`
	// RestProxy: optional
	RestProxy *hdinsightkafkacluster.RestProxy `hcl:"rest_proxy,block"`
	// Roles: required
	Roles *hdinsightkafkacluster.Roles `hcl:"roles,block" validate:"required"`
	// SecurityProfile: optional
	SecurityProfile *hdinsightkafkacluster.SecurityProfile `hcl:"security_profile,block"`
	// StorageAccount: min=0
	StorageAccount []hdinsightkafkacluster.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// StorageAccountGen2: optional
	StorageAccountGen2 *hdinsightkafkacluster.StorageAccountGen2 `hcl:"storage_account_gen2,block"`
	// Timeouts: optional
	Timeouts *hdinsightkafkacluster.Timeouts `hcl:"timeouts,block"`
}
type hdinsightKafkaClusterAttributes struct {
	ref terra.Reference
}

// ClusterVersion returns a reference to field cluster_version of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("cluster_version"))
}

// EncryptionInTransitEnabled returns a reference to field encryption_in_transit_enabled of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) EncryptionInTransitEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hkc.ref.Append("encryption_in_transit_enabled"))
}

// HttpsEndpoint returns a reference to field https_endpoint of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) HttpsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("https_endpoint"))
}

// Id returns a reference to field id of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("id"))
}

// KafkaRestProxyEndpoint returns a reference to field kafka_rest_proxy_endpoint of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) KafkaRestProxyEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("kafka_rest_proxy_endpoint"))
}

// Location returns a reference to field location of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("resource_group_name"))
}

// SshEndpoint returns a reference to field ssh_endpoint of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) SshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("ssh_endpoint"))
}

// Tags returns a reference to field tags of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hkc.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("tier"))
}

// TlsMinVersion returns a reference to field tls_min_version of azurerm_hdinsight_kafka_cluster.
func (hkc hdinsightKafkaClusterAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(hkc.ref.Append("tls_min_version"))
}

func (hkc hdinsightKafkaClusterAttributes) ComponentVersion() terra.ListValue[hdinsightkafkacluster.ComponentVersionAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.ComponentVersionAttributes](hkc.ref.Append("component_version"))
}

func (hkc hdinsightKafkaClusterAttributes) ComputeIsolation() terra.ListValue[hdinsightkafkacluster.ComputeIsolationAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.ComputeIsolationAttributes](hkc.ref.Append("compute_isolation"))
}

func (hkc hdinsightKafkaClusterAttributes) DiskEncryption() terra.ListValue[hdinsightkafkacluster.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.DiskEncryptionAttributes](hkc.ref.Append("disk_encryption"))
}

func (hkc hdinsightKafkaClusterAttributes) Extension() terra.ListValue[hdinsightkafkacluster.ExtensionAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.ExtensionAttributes](hkc.ref.Append("extension"))
}

func (hkc hdinsightKafkaClusterAttributes) Gateway() terra.ListValue[hdinsightkafkacluster.GatewayAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.GatewayAttributes](hkc.ref.Append("gateway"))
}

func (hkc hdinsightKafkaClusterAttributes) Metastores() terra.ListValue[hdinsightkafkacluster.MetastoresAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.MetastoresAttributes](hkc.ref.Append("metastores"))
}

func (hkc hdinsightKafkaClusterAttributes) Monitor() terra.ListValue[hdinsightkafkacluster.MonitorAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.MonitorAttributes](hkc.ref.Append("monitor"))
}

func (hkc hdinsightKafkaClusterAttributes) Network() terra.ListValue[hdinsightkafkacluster.NetworkAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.NetworkAttributes](hkc.ref.Append("network"))
}

func (hkc hdinsightKafkaClusterAttributes) RestProxy() terra.ListValue[hdinsightkafkacluster.RestProxyAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.RestProxyAttributes](hkc.ref.Append("rest_proxy"))
}

func (hkc hdinsightKafkaClusterAttributes) Roles() terra.ListValue[hdinsightkafkacluster.RolesAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.RolesAttributes](hkc.ref.Append("roles"))
}

func (hkc hdinsightKafkaClusterAttributes) SecurityProfile() terra.ListValue[hdinsightkafkacluster.SecurityProfileAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.SecurityProfileAttributes](hkc.ref.Append("security_profile"))
}

func (hkc hdinsightKafkaClusterAttributes) StorageAccount() terra.ListValue[hdinsightkafkacluster.StorageAccountAttributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.StorageAccountAttributes](hkc.ref.Append("storage_account"))
}

func (hkc hdinsightKafkaClusterAttributes) StorageAccountGen2() terra.ListValue[hdinsightkafkacluster.StorageAccountGen2Attributes] {
	return terra.ReferenceAsList[hdinsightkafkacluster.StorageAccountGen2Attributes](hkc.ref.Append("storage_account_gen2"))
}

func (hkc hdinsightKafkaClusterAttributes) Timeouts() hdinsightkafkacluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[hdinsightkafkacluster.TimeoutsAttributes](hkc.ref.Append("timeouts"))
}

type hdinsightKafkaClusterState struct {
	ClusterVersion             string                                          `json:"cluster_version"`
	EncryptionInTransitEnabled bool                                            `json:"encryption_in_transit_enabled"`
	HttpsEndpoint              string                                          `json:"https_endpoint"`
	Id                         string                                          `json:"id"`
	KafkaRestProxyEndpoint     string                                          `json:"kafka_rest_proxy_endpoint"`
	Location                   string                                          `json:"location"`
	Name                       string                                          `json:"name"`
	ResourceGroupName          string                                          `json:"resource_group_name"`
	SshEndpoint                string                                          `json:"ssh_endpoint"`
	Tags                       map[string]string                               `json:"tags"`
	Tier                       string                                          `json:"tier"`
	TlsMinVersion              string                                          `json:"tls_min_version"`
	ComponentVersion           []hdinsightkafkacluster.ComponentVersionState   `json:"component_version"`
	ComputeIsolation           []hdinsightkafkacluster.ComputeIsolationState   `json:"compute_isolation"`
	DiskEncryption             []hdinsightkafkacluster.DiskEncryptionState     `json:"disk_encryption"`
	Extension                  []hdinsightkafkacluster.ExtensionState          `json:"extension"`
	Gateway                    []hdinsightkafkacluster.GatewayState            `json:"gateway"`
	Metastores                 []hdinsightkafkacluster.MetastoresState         `json:"metastores"`
	Monitor                    []hdinsightkafkacluster.MonitorState            `json:"monitor"`
	Network                    []hdinsightkafkacluster.NetworkState            `json:"network"`
	RestProxy                  []hdinsightkafkacluster.RestProxyState          `json:"rest_proxy"`
	Roles                      []hdinsightkafkacluster.RolesState              `json:"roles"`
	SecurityProfile            []hdinsightkafkacluster.SecurityProfileState    `json:"security_profile"`
	StorageAccount             []hdinsightkafkacluster.StorageAccountState     `json:"storage_account"`
	StorageAccountGen2         []hdinsightkafkacluster.StorageAccountGen2State `json:"storage_account_gen2"`
	Timeouts                   *hdinsightkafkacluster.TimeoutsState            `json:"timeouts"`
}
