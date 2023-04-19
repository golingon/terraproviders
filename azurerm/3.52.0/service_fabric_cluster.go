// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	servicefabriccluster "github.com/golingon/terraproviders/azurerm/3.52.0/servicefabriccluster"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewServiceFabricCluster creates a new instance of [ServiceFabricCluster].
func NewServiceFabricCluster(name string, args ServiceFabricClusterArgs) *ServiceFabricCluster {
	return &ServiceFabricCluster{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ServiceFabricCluster)(nil)

// ServiceFabricCluster represents the Terraform resource azurerm_service_fabric_cluster.
type ServiceFabricCluster struct {
	Name      string
	Args      ServiceFabricClusterArgs
	state     *serviceFabricClusterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) Type() string {
	return "azurerm_service_fabric_cluster"
}

// LocalName returns the local name for [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) LocalName() string {
	return sfc.Name
}

// Configuration returns the configuration (args) for [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) Configuration() interface{} {
	return sfc.Args
}

// DependOn is used for other resources to depend on [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) DependOn() terra.Reference {
	return terra.ReferenceResource(sfc)
}

// Dependencies returns the list of resources [ServiceFabricCluster] depends_on.
func (sfc *ServiceFabricCluster) Dependencies() terra.Dependencies {
	return sfc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) LifecycleManagement() *terra.Lifecycle {
	return sfc.Lifecycle
}

// Attributes returns the attributes for [ServiceFabricCluster].
func (sfc *ServiceFabricCluster) Attributes() serviceFabricClusterAttributes {
	return serviceFabricClusterAttributes{ref: terra.ReferenceResource(sfc)}
}

// ImportState imports the given attribute values into [ServiceFabricCluster]'s state.
func (sfc *ServiceFabricCluster) ImportState(av io.Reader) error {
	sfc.state = &serviceFabricClusterState{}
	if err := json.NewDecoder(av).Decode(sfc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sfc.Type(), sfc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ServiceFabricCluster] has state.
func (sfc *ServiceFabricCluster) State() (*serviceFabricClusterState, bool) {
	return sfc.state, sfc.state != nil
}

// StateMust returns the state for [ServiceFabricCluster]. Panics if the state is nil.
func (sfc *ServiceFabricCluster) StateMust() *serviceFabricClusterState {
	if sfc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sfc.Type(), sfc.LocalName()))
	}
	return sfc.state
}

// ServiceFabricClusterArgs contains the configurations for azurerm_service_fabric_cluster.
type ServiceFabricClusterArgs struct {
	// AddOnFeatures: set of string, optional
	AddOnFeatures terra.SetValue[terra.StringValue] `hcl:"add_on_features,attr"`
	// ClusterCodeVersion: string, optional
	ClusterCodeVersion terra.StringValue `hcl:"cluster_code_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagementEndpoint: string, required
	ManagementEndpoint terra.StringValue `hcl:"management_endpoint,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ReliabilityLevel: string, required
	ReliabilityLevel terra.StringValue `hcl:"reliability_level,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ServiceFabricZonalUpgradeMode: string, optional
	ServiceFabricZonalUpgradeMode terra.StringValue `hcl:"service_fabric_zonal_upgrade_mode,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UpgradeMode: string, required
	UpgradeMode terra.StringValue `hcl:"upgrade_mode,attr" validate:"required"`
	// VmImage: string, required
	VmImage terra.StringValue `hcl:"vm_image,attr" validate:"required"`
	// VmssZonalUpgradeMode: string, optional
	VmssZonalUpgradeMode terra.StringValue `hcl:"vmss_zonal_upgrade_mode,attr"`
	// AzureActiveDirectory: optional
	AzureActiveDirectory *servicefabriccluster.AzureActiveDirectory `hcl:"azure_active_directory,block"`
	// Certificate: optional
	Certificate *servicefabriccluster.Certificate `hcl:"certificate,block"`
	// CertificateCommonNames: optional
	CertificateCommonNames *servicefabriccluster.CertificateCommonNames `hcl:"certificate_common_names,block"`
	// ClientCertificateCommonName: min=0
	ClientCertificateCommonName []servicefabriccluster.ClientCertificateCommonName `hcl:"client_certificate_common_name,block" validate:"min=0"`
	// ClientCertificateThumbprint: min=0
	ClientCertificateThumbprint []servicefabriccluster.ClientCertificateThumbprint `hcl:"client_certificate_thumbprint,block" validate:"min=0"`
	// DiagnosticsConfig: optional
	DiagnosticsConfig *servicefabriccluster.DiagnosticsConfig `hcl:"diagnostics_config,block"`
	// FabricSettings: min=0
	FabricSettings []servicefabriccluster.FabricSettings `hcl:"fabric_settings,block" validate:"min=0"`
	// NodeType: min=1
	NodeType []servicefabriccluster.NodeType `hcl:"node_type,block" validate:"min=1"`
	// ReverseProxyCertificate: optional
	ReverseProxyCertificate *servicefabriccluster.ReverseProxyCertificate `hcl:"reverse_proxy_certificate,block"`
	// ReverseProxyCertificateCommonNames: optional
	ReverseProxyCertificateCommonNames *servicefabriccluster.ReverseProxyCertificateCommonNames `hcl:"reverse_proxy_certificate_common_names,block"`
	// Timeouts: optional
	Timeouts *servicefabriccluster.Timeouts `hcl:"timeouts,block"`
	// UpgradePolicy: optional
	UpgradePolicy *servicefabriccluster.UpgradePolicy `hcl:"upgrade_policy,block"`
}
type serviceFabricClusterAttributes struct {
	ref terra.Reference
}

// AddOnFeatures returns a reference to field add_on_features of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) AddOnFeatures() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sfc.ref.Append("add_on_features"))
}

// ClusterCodeVersion returns a reference to field cluster_code_version of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ClusterCodeVersion() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("cluster_code_version"))
}

// ClusterEndpoint returns a reference to field cluster_endpoint of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ClusterEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("cluster_endpoint"))
}

// Id returns a reference to field id of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("location"))
}

// ManagementEndpoint returns a reference to field management_endpoint of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ManagementEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("management_endpoint"))
}

// Name returns a reference to field name of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("name"))
}

// ReliabilityLevel returns a reference to field reliability_level of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ReliabilityLevel() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("reliability_level"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("resource_group_name"))
}

// ServiceFabricZonalUpgradeMode returns a reference to field service_fabric_zonal_upgrade_mode of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) ServiceFabricZonalUpgradeMode() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("service_fabric_zonal_upgrade_mode"))
}

// Tags returns a reference to field tags of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sfc.ref.Append("tags"))
}

// UpgradeMode returns a reference to field upgrade_mode of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) UpgradeMode() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("upgrade_mode"))
}

// VmImage returns a reference to field vm_image of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) VmImage() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("vm_image"))
}

// VmssZonalUpgradeMode returns a reference to field vmss_zonal_upgrade_mode of azurerm_service_fabric_cluster.
func (sfc serviceFabricClusterAttributes) VmssZonalUpgradeMode() terra.StringValue {
	return terra.ReferenceAsString(sfc.ref.Append("vmss_zonal_upgrade_mode"))
}

func (sfc serviceFabricClusterAttributes) AzureActiveDirectory() terra.ListValue[servicefabriccluster.AzureActiveDirectoryAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.AzureActiveDirectoryAttributes](sfc.ref.Append("azure_active_directory"))
}

func (sfc serviceFabricClusterAttributes) Certificate() terra.ListValue[servicefabriccluster.CertificateAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.CertificateAttributes](sfc.ref.Append("certificate"))
}

func (sfc serviceFabricClusterAttributes) CertificateCommonNames() terra.ListValue[servicefabriccluster.CertificateCommonNamesAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.CertificateCommonNamesAttributes](sfc.ref.Append("certificate_common_names"))
}

func (sfc serviceFabricClusterAttributes) ClientCertificateCommonName() terra.ListValue[servicefabriccluster.ClientCertificateCommonNameAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.ClientCertificateCommonNameAttributes](sfc.ref.Append("client_certificate_common_name"))
}

func (sfc serviceFabricClusterAttributes) ClientCertificateThumbprint() terra.ListValue[servicefabriccluster.ClientCertificateThumbprintAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.ClientCertificateThumbprintAttributes](sfc.ref.Append("client_certificate_thumbprint"))
}

func (sfc serviceFabricClusterAttributes) DiagnosticsConfig() terra.ListValue[servicefabriccluster.DiagnosticsConfigAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.DiagnosticsConfigAttributes](sfc.ref.Append("diagnostics_config"))
}

func (sfc serviceFabricClusterAttributes) FabricSettings() terra.ListValue[servicefabriccluster.FabricSettingsAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.FabricSettingsAttributes](sfc.ref.Append("fabric_settings"))
}

func (sfc serviceFabricClusterAttributes) NodeType() terra.ListValue[servicefabriccluster.NodeTypeAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.NodeTypeAttributes](sfc.ref.Append("node_type"))
}

func (sfc serviceFabricClusterAttributes) ReverseProxyCertificate() terra.ListValue[servicefabriccluster.ReverseProxyCertificateAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.ReverseProxyCertificateAttributes](sfc.ref.Append("reverse_proxy_certificate"))
}

func (sfc serviceFabricClusterAttributes) ReverseProxyCertificateCommonNames() terra.ListValue[servicefabriccluster.ReverseProxyCertificateCommonNamesAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.ReverseProxyCertificateCommonNamesAttributes](sfc.ref.Append("reverse_proxy_certificate_common_names"))
}

func (sfc serviceFabricClusterAttributes) Timeouts() servicefabriccluster.TimeoutsAttributes {
	return terra.ReferenceAsSingle[servicefabriccluster.TimeoutsAttributes](sfc.ref.Append("timeouts"))
}

func (sfc serviceFabricClusterAttributes) UpgradePolicy() terra.ListValue[servicefabriccluster.UpgradePolicyAttributes] {
	return terra.ReferenceAsList[servicefabriccluster.UpgradePolicyAttributes](sfc.ref.Append("upgrade_policy"))
}

type serviceFabricClusterState struct {
	AddOnFeatures                      []string                                                       `json:"add_on_features"`
	ClusterCodeVersion                 string                                                         `json:"cluster_code_version"`
	ClusterEndpoint                    string                                                         `json:"cluster_endpoint"`
	Id                                 string                                                         `json:"id"`
	Location                           string                                                         `json:"location"`
	ManagementEndpoint                 string                                                         `json:"management_endpoint"`
	Name                               string                                                         `json:"name"`
	ReliabilityLevel                   string                                                         `json:"reliability_level"`
	ResourceGroupName                  string                                                         `json:"resource_group_name"`
	ServiceFabricZonalUpgradeMode      string                                                         `json:"service_fabric_zonal_upgrade_mode"`
	Tags                               map[string]string                                              `json:"tags"`
	UpgradeMode                        string                                                         `json:"upgrade_mode"`
	VmImage                            string                                                         `json:"vm_image"`
	VmssZonalUpgradeMode               string                                                         `json:"vmss_zonal_upgrade_mode"`
	AzureActiveDirectory               []servicefabriccluster.AzureActiveDirectoryState               `json:"azure_active_directory"`
	Certificate                        []servicefabriccluster.CertificateState                        `json:"certificate"`
	CertificateCommonNames             []servicefabriccluster.CertificateCommonNamesState             `json:"certificate_common_names"`
	ClientCertificateCommonName        []servicefabriccluster.ClientCertificateCommonNameState        `json:"client_certificate_common_name"`
	ClientCertificateThumbprint        []servicefabriccluster.ClientCertificateThumbprintState        `json:"client_certificate_thumbprint"`
	DiagnosticsConfig                  []servicefabriccluster.DiagnosticsConfigState                  `json:"diagnostics_config"`
	FabricSettings                     []servicefabriccluster.FabricSettingsState                     `json:"fabric_settings"`
	NodeType                           []servicefabriccluster.NodeTypeState                           `json:"node_type"`
	ReverseProxyCertificate            []servicefabriccluster.ReverseProxyCertificateState            `json:"reverse_proxy_certificate"`
	ReverseProxyCertificateCommonNames []servicefabriccluster.ReverseProxyCertificateCommonNamesState `json:"reverse_proxy_certificate_common_names"`
	Timeouts                           *servicefabriccluster.TimeoutsState                            `json:"timeouts"`
	UpgradePolicy                      []servicefabriccluster.UpgradePolicyState                      `json:"upgrade_policy"`
}
