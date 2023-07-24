// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	hdinsightsparkcluster "github.com/golingon/terraproviders/azurerm/3.66.0/hdinsightsparkcluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewHdinsightSparkCluster creates a new instance of [HdinsightSparkCluster].
func NewHdinsightSparkCluster(name string, args HdinsightSparkClusterArgs) *HdinsightSparkCluster {
	return &HdinsightSparkCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HdinsightSparkCluster)(nil)

// HdinsightSparkCluster represents the Terraform resource azurerm_hdinsight_spark_cluster.
type HdinsightSparkCluster struct {
	Name      string
	Args      HdinsightSparkClusterArgs
	state     *hdinsightSparkClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) Type() string {
	return "azurerm_hdinsight_spark_cluster"
}

// LocalName returns the local name for [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) LocalName() string {
	return hsc.Name
}

// Configuration returns the configuration (args) for [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) Configuration() interface{} {
	return hsc.Args
}

// DependOn is used for other resources to depend on [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(hsc)
}

// Dependencies returns the list of resources [HdinsightSparkCluster] depends_on.
func (hsc *HdinsightSparkCluster) Dependencies() terra.Dependencies {
	return hsc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) LifecycleManagement() *terra.Lifecycle {
	return hsc.Lifecycle
}

// Attributes returns the attributes for [HdinsightSparkCluster].
func (hsc *HdinsightSparkCluster) Attributes() hdinsightSparkClusterAttributes {
	return hdinsightSparkClusterAttributes{ref: terra.ReferenceResource(hsc)}
}

// ImportState imports the given attribute values into [HdinsightSparkCluster]'s state.
func (hsc *HdinsightSparkCluster) ImportState(av io.Reader) error {
	hsc.state = &hdinsightSparkClusterState{}
	if err := json.NewDecoder(av).Decode(hsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hsc.Type(), hsc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HdinsightSparkCluster] has state.
func (hsc *HdinsightSparkCluster) State() (*hdinsightSparkClusterState, bool) {
	return hsc.state, hsc.state != nil
}

// StateMust returns the state for [HdinsightSparkCluster]. Panics if the state is nil.
func (hsc *HdinsightSparkCluster) StateMust() *hdinsightSparkClusterState {
	if hsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hsc.Type(), hsc.LocalName()))
	}
	return hsc.state
}

// HdinsightSparkClusterArgs contains the configurations for azurerm_hdinsight_spark_cluster.
type HdinsightSparkClusterArgs struct {
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
	ComponentVersion *hdinsightsparkcluster.ComponentVersion `hcl:"component_version,block" validate:"required"`
	// ComputeIsolation: optional
	ComputeIsolation *hdinsightsparkcluster.ComputeIsolation `hcl:"compute_isolation,block"`
	// DiskEncryption: min=0
	DiskEncryption []hdinsightsparkcluster.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extension: optional
	Extension *hdinsightsparkcluster.Extension `hcl:"extension,block"`
	// Gateway: required
	Gateway *hdinsightsparkcluster.Gateway `hcl:"gateway,block" validate:"required"`
	// Metastores: optional
	Metastores *hdinsightsparkcluster.Metastores `hcl:"metastores,block"`
	// Monitor: optional
	Monitor *hdinsightsparkcluster.Monitor `hcl:"monitor,block"`
	// Network: optional
	Network *hdinsightsparkcluster.Network `hcl:"network,block"`
	// Roles: required
	Roles *hdinsightsparkcluster.Roles `hcl:"roles,block" validate:"required"`
	// SecurityProfile: optional
	SecurityProfile *hdinsightsparkcluster.SecurityProfile `hcl:"security_profile,block"`
	// StorageAccount: min=0
	StorageAccount []hdinsightsparkcluster.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// StorageAccountGen2: optional
	StorageAccountGen2 *hdinsightsparkcluster.StorageAccountGen2 `hcl:"storage_account_gen2,block"`
	// Timeouts: optional
	Timeouts *hdinsightsparkcluster.Timeouts `hcl:"timeouts,block"`
}
type hdinsightSparkClusterAttributes struct {
	ref terra.Reference
}

// ClusterVersion returns a reference to field cluster_version of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("cluster_version"))
}

// EncryptionInTransitEnabled returns a reference to field encryption_in_transit_enabled of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) EncryptionInTransitEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hsc.ref.Append("encryption_in_transit_enabled"))
}

// HttpsEndpoint returns a reference to field https_endpoint of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) HttpsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("https_endpoint"))
}

// Id returns a reference to field id of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("resource_group_name"))
}

// SshEndpoint returns a reference to field ssh_endpoint of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) SshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("ssh_endpoint"))
}

// Tags returns a reference to field tags of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hsc.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("tier"))
}

// TlsMinVersion returns a reference to field tls_min_version of azurerm_hdinsight_spark_cluster.
func (hsc hdinsightSparkClusterAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(hsc.ref.Append("tls_min_version"))
}

func (hsc hdinsightSparkClusterAttributes) ComponentVersion() terra.ListValue[hdinsightsparkcluster.ComponentVersionAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.ComponentVersionAttributes](hsc.ref.Append("component_version"))
}

func (hsc hdinsightSparkClusterAttributes) ComputeIsolation() terra.ListValue[hdinsightsparkcluster.ComputeIsolationAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.ComputeIsolationAttributes](hsc.ref.Append("compute_isolation"))
}

func (hsc hdinsightSparkClusterAttributes) DiskEncryption() terra.ListValue[hdinsightsparkcluster.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.DiskEncryptionAttributes](hsc.ref.Append("disk_encryption"))
}

func (hsc hdinsightSparkClusterAttributes) Extension() terra.ListValue[hdinsightsparkcluster.ExtensionAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.ExtensionAttributes](hsc.ref.Append("extension"))
}

func (hsc hdinsightSparkClusterAttributes) Gateway() terra.ListValue[hdinsightsparkcluster.GatewayAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.GatewayAttributes](hsc.ref.Append("gateway"))
}

func (hsc hdinsightSparkClusterAttributes) Metastores() terra.ListValue[hdinsightsparkcluster.MetastoresAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.MetastoresAttributes](hsc.ref.Append("metastores"))
}

func (hsc hdinsightSparkClusterAttributes) Monitor() terra.ListValue[hdinsightsparkcluster.MonitorAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.MonitorAttributes](hsc.ref.Append("monitor"))
}

func (hsc hdinsightSparkClusterAttributes) Network() terra.ListValue[hdinsightsparkcluster.NetworkAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.NetworkAttributes](hsc.ref.Append("network"))
}

func (hsc hdinsightSparkClusterAttributes) Roles() terra.ListValue[hdinsightsparkcluster.RolesAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.RolesAttributes](hsc.ref.Append("roles"))
}

func (hsc hdinsightSparkClusterAttributes) SecurityProfile() terra.ListValue[hdinsightsparkcluster.SecurityProfileAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.SecurityProfileAttributes](hsc.ref.Append("security_profile"))
}

func (hsc hdinsightSparkClusterAttributes) StorageAccount() terra.ListValue[hdinsightsparkcluster.StorageAccountAttributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.StorageAccountAttributes](hsc.ref.Append("storage_account"))
}

func (hsc hdinsightSparkClusterAttributes) StorageAccountGen2() terra.ListValue[hdinsightsparkcluster.StorageAccountGen2Attributes] {
	return terra.ReferenceAsList[hdinsightsparkcluster.StorageAccountGen2Attributes](hsc.ref.Append("storage_account_gen2"))
}

func (hsc hdinsightSparkClusterAttributes) Timeouts() hdinsightsparkcluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[hdinsightsparkcluster.TimeoutsAttributes](hsc.ref.Append("timeouts"))
}

type hdinsightSparkClusterState struct {
	ClusterVersion             string                                          `json:"cluster_version"`
	EncryptionInTransitEnabled bool                                            `json:"encryption_in_transit_enabled"`
	HttpsEndpoint              string                                          `json:"https_endpoint"`
	Id                         string                                          `json:"id"`
	Location                   string                                          `json:"location"`
	Name                       string                                          `json:"name"`
	ResourceGroupName          string                                          `json:"resource_group_name"`
	SshEndpoint                string                                          `json:"ssh_endpoint"`
	Tags                       map[string]string                               `json:"tags"`
	Tier                       string                                          `json:"tier"`
	TlsMinVersion              string                                          `json:"tls_min_version"`
	ComponentVersion           []hdinsightsparkcluster.ComponentVersionState   `json:"component_version"`
	ComputeIsolation           []hdinsightsparkcluster.ComputeIsolationState   `json:"compute_isolation"`
	DiskEncryption             []hdinsightsparkcluster.DiskEncryptionState     `json:"disk_encryption"`
	Extension                  []hdinsightsparkcluster.ExtensionState          `json:"extension"`
	Gateway                    []hdinsightsparkcluster.GatewayState            `json:"gateway"`
	Metastores                 []hdinsightsparkcluster.MetastoresState         `json:"metastores"`
	Monitor                    []hdinsightsparkcluster.MonitorState            `json:"monitor"`
	Network                    []hdinsightsparkcluster.NetworkState            `json:"network"`
	Roles                      []hdinsightsparkcluster.RolesState              `json:"roles"`
	SecurityProfile            []hdinsightsparkcluster.SecurityProfileState    `json:"security_profile"`
	StorageAccount             []hdinsightsparkcluster.StorageAccountState     `json:"storage_account"`
	StorageAccountGen2         []hdinsightsparkcluster.StorageAccountGen2State `json:"storage_account_gen2"`
	Timeouts                   *hdinsightsparkcluster.TimeoutsState            `json:"timeouts"`
}
