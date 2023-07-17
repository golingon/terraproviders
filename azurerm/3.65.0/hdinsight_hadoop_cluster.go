// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	hdinsighthadoopcluster "github.com/golingon/terraproviders/azurerm/3.65.0/hdinsighthadoopcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHdinsightHadoopCluster creates a new instance of [HdinsightHadoopCluster].
func NewHdinsightHadoopCluster(name string, args HdinsightHadoopClusterArgs) *HdinsightHadoopCluster {
	return &HdinsightHadoopCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HdinsightHadoopCluster)(nil)

// HdinsightHadoopCluster represents the Terraform resource azurerm_hdinsight_hadoop_cluster.
type HdinsightHadoopCluster struct {
	Name      string
	Args      HdinsightHadoopClusterArgs
	state     *hdinsightHadoopClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) Type() string {
	return "azurerm_hdinsight_hadoop_cluster"
}

// LocalName returns the local name for [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) LocalName() string {
	return hhc.Name
}

// Configuration returns the configuration (args) for [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) Configuration() interface{} {
	return hhc.Args
}

// DependOn is used for other resources to depend on [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(hhc)
}

// Dependencies returns the list of resources [HdinsightHadoopCluster] depends_on.
func (hhc *HdinsightHadoopCluster) Dependencies() terra.Dependencies {
	return hhc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) LifecycleManagement() *terra.Lifecycle {
	return hhc.Lifecycle
}

// Attributes returns the attributes for [HdinsightHadoopCluster].
func (hhc *HdinsightHadoopCluster) Attributes() hdinsightHadoopClusterAttributes {
	return hdinsightHadoopClusterAttributes{ref: terra.ReferenceResource(hhc)}
}

// ImportState imports the given attribute values into [HdinsightHadoopCluster]'s state.
func (hhc *HdinsightHadoopCluster) ImportState(av io.Reader) error {
	hhc.state = &hdinsightHadoopClusterState{}
	if err := json.NewDecoder(av).Decode(hhc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hhc.Type(), hhc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HdinsightHadoopCluster] has state.
func (hhc *HdinsightHadoopCluster) State() (*hdinsightHadoopClusterState, bool) {
	return hhc.state, hhc.state != nil
}

// StateMust returns the state for [HdinsightHadoopCluster]. Panics if the state is nil.
func (hhc *HdinsightHadoopCluster) StateMust() *hdinsightHadoopClusterState {
	if hhc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hhc.Type(), hhc.LocalName()))
	}
	return hhc.state
}

// HdinsightHadoopClusterArgs contains the configurations for azurerm_hdinsight_hadoop_cluster.
type HdinsightHadoopClusterArgs struct {
	// ClusterVersion: string, required
	ClusterVersion terra.StringValue `hcl:"cluster_version,attr" validate:"required"`
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
	ComponentVersion *hdinsighthadoopcluster.ComponentVersion `hcl:"component_version,block" validate:"required"`
	// ComputeIsolation: optional
	ComputeIsolation *hdinsighthadoopcluster.ComputeIsolation `hcl:"compute_isolation,block"`
	// DiskEncryption: min=0
	DiskEncryption []hdinsighthadoopcluster.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extension: optional
	Extension *hdinsighthadoopcluster.Extension `hcl:"extension,block"`
	// Gateway: required
	Gateway *hdinsighthadoopcluster.Gateway `hcl:"gateway,block" validate:"required"`
	// Metastores: optional
	Metastores *hdinsighthadoopcluster.Metastores `hcl:"metastores,block"`
	// Monitor: optional
	Monitor *hdinsighthadoopcluster.Monitor `hcl:"monitor,block"`
	// Network: optional
	Network *hdinsighthadoopcluster.Network `hcl:"network,block"`
	// Roles: required
	Roles *hdinsighthadoopcluster.Roles `hcl:"roles,block" validate:"required"`
	// SecurityProfile: optional
	SecurityProfile *hdinsighthadoopcluster.SecurityProfile `hcl:"security_profile,block"`
	// StorageAccount: min=0
	StorageAccount []hdinsighthadoopcluster.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// StorageAccountGen2: optional
	StorageAccountGen2 *hdinsighthadoopcluster.StorageAccountGen2 `hcl:"storage_account_gen2,block"`
	// Timeouts: optional
	Timeouts *hdinsighthadoopcluster.Timeouts `hcl:"timeouts,block"`
}
type hdinsightHadoopClusterAttributes struct {
	ref terra.Reference
}

// ClusterVersion returns a reference to field cluster_version of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("cluster_version"))
}

// HttpsEndpoint returns a reference to field https_endpoint of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) HttpsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("https_endpoint"))
}

// Id returns a reference to field id of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("resource_group_name"))
}

// SshEndpoint returns a reference to field ssh_endpoint of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) SshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("ssh_endpoint"))
}

// Tags returns a reference to field tags of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hhc.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("tier"))
}

// TlsMinVersion returns a reference to field tls_min_version of azurerm_hdinsight_hadoop_cluster.
func (hhc hdinsightHadoopClusterAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(hhc.ref.Append("tls_min_version"))
}

func (hhc hdinsightHadoopClusterAttributes) ComponentVersion() terra.ListValue[hdinsighthadoopcluster.ComponentVersionAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.ComponentVersionAttributes](hhc.ref.Append("component_version"))
}

func (hhc hdinsightHadoopClusterAttributes) ComputeIsolation() terra.ListValue[hdinsighthadoopcluster.ComputeIsolationAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.ComputeIsolationAttributes](hhc.ref.Append("compute_isolation"))
}

func (hhc hdinsightHadoopClusterAttributes) DiskEncryption() terra.ListValue[hdinsighthadoopcluster.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.DiskEncryptionAttributes](hhc.ref.Append("disk_encryption"))
}

func (hhc hdinsightHadoopClusterAttributes) Extension() terra.ListValue[hdinsighthadoopcluster.ExtensionAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.ExtensionAttributes](hhc.ref.Append("extension"))
}

func (hhc hdinsightHadoopClusterAttributes) Gateway() terra.ListValue[hdinsighthadoopcluster.GatewayAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.GatewayAttributes](hhc.ref.Append("gateway"))
}

func (hhc hdinsightHadoopClusterAttributes) Metastores() terra.ListValue[hdinsighthadoopcluster.MetastoresAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.MetastoresAttributes](hhc.ref.Append("metastores"))
}

func (hhc hdinsightHadoopClusterAttributes) Monitor() terra.ListValue[hdinsighthadoopcluster.MonitorAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.MonitorAttributes](hhc.ref.Append("monitor"))
}

func (hhc hdinsightHadoopClusterAttributes) Network() terra.ListValue[hdinsighthadoopcluster.NetworkAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.NetworkAttributes](hhc.ref.Append("network"))
}

func (hhc hdinsightHadoopClusterAttributes) Roles() terra.ListValue[hdinsighthadoopcluster.RolesAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.RolesAttributes](hhc.ref.Append("roles"))
}

func (hhc hdinsightHadoopClusterAttributes) SecurityProfile() terra.ListValue[hdinsighthadoopcluster.SecurityProfileAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.SecurityProfileAttributes](hhc.ref.Append("security_profile"))
}

func (hhc hdinsightHadoopClusterAttributes) StorageAccount() terra.ListValue[hdinsighthadoopcluster.StorageAccountAttributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.StorageAccountAttributes](hhc.ref.Append("storage_account"))
}

func (hhc hdinsightHadoopClusterAttributes) StorageAccountGen2() terra.ListValue[hdinsighthadoopcluster.StorageAccountGen2Attributes] {
	return terra.ReferenceAsList[hdinsighthadoopcluster.StorageAccountGen2Attributes](hhc.ref.Append("storage_account_gen2"))
}

func (hhc hdinsightHadoopClusterAttributes) Timeouts() hdinsighthadoopcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[hdinsighthadoopcluster.TimeoutsAttributes](hhc.ref.Append("timeouts"))
}

type hdinsightHadoopClusterState struct {
	ClusterVersion     string                                           `json:"cluster_version"`
	HttpsEndpoint      string                                           `json:"https_endpoint"`
	Id                 string                                           `json:"id"`
	Location           string                                           `json:"location"`
	Name               string                                           `json:"name"`
	ResourceGroupName  string                                           `json:"resource_group_name"`
	SshEndpoint        string                                           `json:"ssh_endpoint"`
	Tags               map[string]string                                `json:"tags"`
	Tier               string                                           `json:"tier"`
	TlsMinVersion      string                                           `json:"tls_min_version"`
	ComponentVersion   []hdinsighthadoopcluster.ComponentVersionState   `json:"component_version"`
	ComputeIsolation   []hdinsighthadoopcluster.ComputeIsolationState   `json:"compute_isolation"`
	DiskEncryption     []hdinsighthadoopcluster.DiskEncryptionState     `json:"disk_encryption"`
	Extension          []hdinsighthadoopcluster.ExtensionState          `json:"extension"`
	Gateway            []hdinsighthadoopcluster.GatewayState            `json:"gateway"`
	Metastores         []hdinsighthadoopcluster.MetastoresState         `json:"metastores"`
	Monitor            []hdinsighthadoopcluster.MonitorState            `json:"monitor"`
	Network            []hdinsighthadoopcluster.NetworkState            `json:"network"`
	Roles              []hdinsighthadoopcluster.RolesState              `json:"roles"`
	SecurityProfile    []hdinsighthadoopcluster.SecurityProfileState    `json:"security_profile"`
	StorageAccount     []hdinsighthadoopcluster.StorageAccountState     `json:"storage_account"`
	StorageAccountGen2 []hdinsighthadoopcluster.StorageAccountGen2State `json:"storage_account_gen2"`
	Timeouts           *hdinsighthadoopcluster.TimeoutsState            `json:"timeouts"`
}
