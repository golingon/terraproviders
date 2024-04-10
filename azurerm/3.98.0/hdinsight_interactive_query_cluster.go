// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	hdinsightinteractivequerycluster "github.com/golingon/terraproviders/azurerm/3.98.0/hdinsightinteractivequerycluster"
	"io"
)

// NewHdinsightInteractiveQueryCluster creates a new instance of [HdinsightInteractiveQueryCluster].
func NewHdinsightInteractiveQueryCluster(name string, args HdinsightInteractiveQueryClusterArgs) *HdinsightInteractiveQueryCluster {
	return &HdinsightInteractiveQueryCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*HdinsightInteractiveQueryCluster)(nil)

// HdinsightInteractiveQueryCluster represents the Terraform resource azurerm_hdinsight_interactive_query_cluster.
type HdinsightInteractiveQueryCluster struct {
	Name      string
	Args      HdinsightInteractiveQueryClusterArgs
	state     *hdinsightInteractiveQueryClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) Type() string {
	return "azurerm_hdinsight_interactive_query_cluster"
}

// LocalName returns the local name for [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) LocalName() string {
	return hiqc.Name
}

// Configuration returns the configuration (args) for [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) Configuration() interface{} {
	return hiqc.Args
}

// DependOn is used for other resources to depend on [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(hiqc)
}

// Dependencies returns the list of resources [HdinsightInteractiveQueryCluster] depends_on.
func (hiqc *HdinsightInteractiveQueryCluster) Dependencies() terra.Dependencies {
	return hiqc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) LifecycleManagement() *terra.Lifecycle {
	return hiqc.Lifecycle
}

// Attributes returns the attributes for [HdinsightInteractiveQueryCluster].
func (hiqc *HdinsightInteractiveQueryCluster) Attributes() hdinsightInteractiveQueryClusterAttributes {
	return hdinsightInteractiveQueryClusterAttributes{ref: terra.ReferenceResource(hiqc)}
}

// ImportState imports the given attribute values into [HdinsightInteractiveQueryCluster]'s state.
func (hiqc *HdinsightInteractiveQueryCluster) ImportState(av io.Reader) error {
	hiqc.state = &hdinsightInteractiveQueryClusterState{}
	if err := json.NewDecoder(av).Decode(hiqc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", hiqc.Type(), hiqc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [HdinsightInteractiveQueryCluster] has state.
func (hiqc *HdinsightInteractiveQueryCluster) State() (*hdinsightInteractiveQueryClusterState, bool) {
	return hiqc.state, hiqc.state != nil
}

// StateMust returns the state for [HdinsightInteractiveQueryCluster]. Panics if the state is nil.
func (hiqc *HdinsightInteractiveQueryCluster) StateMust() *hdinsightInteractiveQueryClusterState {
	if hiqc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", hiqc.Type(), hiqc.LocalName()))
	}
	return hiqc.state
}

// HdinsightInteractiveQueryClusterArgs contains the configurations for azurerm_hdinsight_interactive_query_cluster.
type HdinsightInteractiveQueryClusterArgs struct {
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
	ComponentVersion *hdinsightinteractivequerycluster.ComponentVersion `hcl:"component_version,block" validate:"required"`
	// ComputeIsolation: optional
	ComputeIsolation *hdinsightinteractivequerycluster.ComputeIsolation `hcl:"compute_isolation,block"`
	// DiskEncryption: min=0
	DiskEncryption []hdinsightinteractivequerycluster.DiskEncryption `hcl:"disk_encryption,block" validate:"min=0"`
	// Extension: optional
	Extension *hdinsightinteractivequerycluster.Extension `hcl:"extension,block"`
	// Gateway: required
	Gateway *hdinsightinteractivequerycluster.Gateway `hcl:"gateway,block" validate:"required"`
	// Metastores: optional
	Metastores *hdinsightinteractivequerycluster.Metastores `hcl:"metastores,block"`
	// Monitor: optional
	Monitor *hdinsightinteractivequerycluster.Monitor `hcl:"monitor,block"`
	// Network: optional
	Network *hdinsightinteractivequerycluster.Network `hcl:"network,block"`
	// Roles: required
	Roles *hdinsightinteractivequerycluster.Roles `hcl:"roles,block" validate:"required"`
	// SecurityProfile: optional
	SecurityProfile *hdinsightinteractivequerycluster.SecurityProfile `hcl:"security_profile,block"`
	// StorageAccount: min=0
	StorageAccount []hdinsightinteractivequerycluster.StorageAccount `hcl:"storage_account,block" validate:"min=0"`
	// StorageAccountGen2: optional
	StorageAccountGen2 *hdinsightinteractivequerycluster.StorageAccountGen2 `hcl:"storage_account_gen2,block"`
	// Timeouts: optional
	Timeouts *hdinsightinteractivequerycluster.Timeouts `hcl:"timeouts,block"`
}
type hdinsightInteractiveQueryClusterAttributes struct {
	ref terra.Reference
}

// ClusterVersion returns a reference to field cluster_version of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) ClusterVersion() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("cluster_version"))
}

// EncryptionInTransitEnabled returns a reference to field encryption_in_transit_enabled of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) EncryptionInTransitEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(hiqc.ref.Append("encryption_in_transit_enabled"))
}

// HttpsEndpoint returns a reference to field https_endpoint of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) HttpsEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("https_endpoint"))
}

// Id returns a reference to field id of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("resource_group_name"))
}

// SshEndpoint returns a reference to field ssh_endpoint of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) SshEndpoint() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("ssh_endpoint"))
}

// Tags returns a reference to field tags of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](hiqc.ref.Append("tags"))
}

// Tier returns a reference to field tier of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) Tier() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("tier"))
}

// TlsMinVersion returns a reference to field tls_min_version of azurerm_hdinsight_interactive_query_cluster.
func (hiqc hdinsightInteractiveQueryClusterAttributes) TlsMinVersion() terra.StringValue {
	return terra.ReferenceAsString(hiqc.ref.Append("tls_min_version"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) ComponentVersion() terra.ListValue[hdinsightinteractivequerycluster.ComponentVersionAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.ComponentVersionAttributes](hiqc.ref.Append("component_version"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) ComputeIsolation() terra.ListValue[hdinsightinteractivequerycluster.ComputeIsolationAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.ComputeIsolationAttributes](hiqc.ref.Append("compute_isolation"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) DiskEncryption() terra.ListValue[hdinsightinteractivequerycluster.DiskEncryptionAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.DiskEncryptionAttributes](hiqc.ref.Append("disk_encryption"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Extension() terra.ListValue[hdinsightinteractivequerycluster.ExtensionAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.ExtensionAttributes](hiqc.ref.Append("extension"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Gateway() terra.ListValue[hdinsightinteractivequerycluster.GatewayAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.GatewayAttributes](hiqc.ref.Append("gateway"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Metastores() terra.ListValue[hdinsightinteractivequerycluster.MetastoresAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.MetastoresAttributes](hiqc.ref.Append("metastores"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Monitor() terra.ListValue[hdinsightinteractivequerycluster.MonitorAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.MonitorAttributes](hiqc.ref.Append("monitor"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Network() terra.ListValue[hdinsightinteractivequerycluster.NetworkAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.NetworkAttributes](hiqc.ref.Append("network"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Roles() terra.ListValue[hdinsightinteractivequerycluster.RolesAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.RolesAttributes](hiqc.ref.Append("roles"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) SecurityProfile() terra.ListValue[hdinsightinteractivequerycluster.SecurityProfileAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.SecurityProfileAttributes](hiqc.ref.Append("security_profile"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) StorageAccount() terra.ListValue[hdinsightinteractivequerycluster.StorageAccountAttributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.StorageAccountAttributes](hiqc.ref.Append("storage_account"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) StorageAccountGen2() terra.ListValue[hdinsightinteractivequerycluster.StorageAccountGen2Attributes] {
	return terra.ReferenceAsList[hdinsightinteractivequerycluster.StorageAccountGen2Attributes](hiqc.ref.Append("storage_account_gen2"))
}

func (hiqc hdinsightInteractiveQueryClusterAttributes) Timeouts() hdinsightinteractivequerycluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[hdinsightinteractivequerycluster.TimeoutsAttributes](hiqc.ref.Append("timeouts"))
}

type hdinsightInteractiveQueryClusterState struct {
	ClusterVersion             string                                                     `json:"cluster_version"`
	EncryptionInTransitEnabled bool                                                       `json:"encryption_in_transit_enabled"`
	HttpsEndpoint              string                                                     `json:"https_endpoint"`
	Id                         string                                                     `json:"id"`
	Location                   string                                                     `json:"location"`
	Name                       string                                                     `json:"name"`
	ResourceGroupName          string                                                     `json:"resource_group_name"`
	SshEndpoint                string                                                     `json:"ssh_endpoint"`
	Tags                       map[string]string                                          `json:"tags"`
	Tier                       string                                                     `json:"tier"`
	TlsMinVersion              string                                                     `json:"tls_min_version"`
	ComponentVersion           []hdinsightinteractivequerycluster.ComponentVersionState   `json:"component_version"`
	ComputeIsolation           []hdinsightinteractivequerycluster.ComputeIsolationState   `json:"compute_isolation"`
	DiskEncryption             []hdinsightinteractivequerycluster.DiskEncryptionState     `json:"disk_encryption"`
	Extension                  []hdinsightinteractivequerycluster.ExtensionState          `json:"extension"`
	Gateway                    []hdinsightinteractivequerycluster.GatewayState            `json:"gateway"`
	Metastores                 []hdinsightinteractivequerycluster.MetastoresState         `json:"metastores"`
	Monitor                    []hdinsightinteractivequerycluster.MonitorState            `json:"monitor"`
	Network                    []hdinsightinteractivequerycluster.NetworkState            `json:"network"`
	Roles                      []hdinsightinteractivequerycluster.RolesState              `json:"roles"`
	SecurityProfile            []hdinsightinteractivequerycluster.SecurityProfileState    `json:"security_profile"`
	StorageAccount             []hdinsightinteractivequerycluster.StorageAccountState     `json:"storage_account"`
	StorageAccountGen2         []hdinsightinteractivequerycluster.StorageAccountGen2State `json:"storage_account_gen2"`
	Timeouts                   *hdinsightinteractivequerycluster.TimeoutsState            `json:"timeouts"`
}
