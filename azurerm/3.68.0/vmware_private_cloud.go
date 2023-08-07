// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	vmwareprivatecloud "github.com/golingon/terraproviders/azurerm/3.68.0/vmwareprivatecloud"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVmwarePrivateCloud creates a new instance of [VmwarePrivateCloud].
func NewVmwarePrivateCloud(name string, args VmwarePrivateCloudArgs) *VmwarePrivateCloud {
	return &VmwarePrivateCloud{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VmwarePrivateCloud)(nil)

// VmwarePrivateCloud represents the Terraform resource azurerm_vmware_private_cloud.
type VmwarePrivateCloud struct {
	Name      string
	Args      VmwarePrivateCloudArgs
	state     *vmwarePrivateCloudState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) Type() string {
	return "azurerm_vmware_private_cloud"
}

// LocalName returns the local name for [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) LocalName() string {
	return vpc.Name
}

// Configuration returns the configuration (args) for [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) Configuration() interface{} {
	return vpc.Args
}

// DependOn is used for other resources to depend on [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) DependOn() terra.Reference {
	return terra.ReferenceResource(vpc)
}

// Dependencies returns the list of resources [VmwarePrivateCloud] depends_on.
func (vpc *VmwarePrivateCloud) Dependencies() terra.Dependencies {
	return vpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) LifecycleManagement() *terra.Lifecycle {
	return vpc.Lifecycle
}

// Attributes returns the attributes for [VmwarePrivateCloud].
func (vpc *VmwarePrivateCloud) Attributes() vmwarePrivateCloudAttributes {
	return vmwarePrivateCloudAttributes{ref: terra.ReferenceResource(vpc)}
}

// ImportState imports the given attribute values into [VmwarePrivateCloud]'s state.
func (vpc *VmwarePrivateCloud) ImportState(av io.Reader) error {
	vpc.state = &vmwarePrivateCloudState{}
	if err := json.NewDecoder(av).Decode(vpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpc.Type(), vpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VmwarePrivateCloud] has state.
func (vpc *VmwarePrivateCloud) State() (*vmwarePrivateCloudState, bool) {
	return vpc.state, vpc.state != nil
}

// StateMust returns the state for [VmwarePrivateCloud]. Panics if the state is nil.
func (vpc *VmwarePrivateCloud) StateMust() *vmwarePrivateCloudState {
	if vpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpc.Type(), vpc.LocalName()))
	}
	return vpc.state
}

// VmwarePrivateCloudArgs contains the configurations for azurerm_vmware_private_cloud.
type VmwarePrivateCloudArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternetConnectionEnabled: bool, optional
	InternetConnectionEnabled terra.BoolValue `hcl:"internet_connection_enabled,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkSubnetCidr: string, required
	NetworkSubnetCidr terra.StringValue `hcl:"network_subnet_cidr,attr" validate:"required"`
	// NsxtPassword: string, optional
	NsxtPassword terra.StringValue `hcl:"nsxt_password,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VcenterPassword: string, optional
	VcenterPassword terra.StringValue `hcl:"vcenter_password,attr"`
	// Circuit: min=0
	Circuit []vmwareprivatecloud.Circuit `hcl:"circuit,block" validate:"min=0"`
	// ManagementCluster: required
	ManagementCluster *vmwareprivatecloud.ManagementCluster `hcl:"management_cluster,block" validate:"required"`
	// Timeouts: optional
	Timeouts *vmwareprivatecloud.Timeouts `hcl:"timeouts,block"`
}
type vmwarePrivateCloudAttributes struct {
	ref terra.Reference
}

// HcxCloudManagerEndpoint returns a reference to field hcx_cloud_manager_endpoint of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) HcxCloudManagerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("hcx_cloud_manager_endpoint"))
}

// Id returns a reference to field id of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("id"))
}

// InternetConnectionEnabled returns a reference to field internet_connection_enabled of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) InternetConnectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(vpc.ref.Append("internet_connection_enabled"))
}

// Location returns a reference to field location of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("location"))
}

// ManagementSubnetCidr returns a reference to field management_subnet_cidr of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) ManagementSubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("management_subnet_cidr"))
}

// Name returns a reference to field name of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("name"))
}

// NetworkSubnetCidr returns a reference to field network_subnet_cidr of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) NetworkSubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("network_subnet_cidr"))
}

// NsxtCertificateThumbprint returns a reference to field nsxt_certificate_thumbprint of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) NsxtCertificateThumbprint() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("nsxt_certificate_thumbprint"))
}

// NsxtManagerEndpoint returns a reference to field nsxt_manager_endpoint of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) NsxtManagerEndpoint() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("nsxt_manager_endpoint"))
}

// NsxtPassword returns a reference to field nsxt_password of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) NsxtPassword() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("nsxt_password"))
}

// ProvisioningSubnetCidr returns a reference to field provisioning_subnet_cidr of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) ProvisioningSubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("provisioning_subnet_cidr"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("sku_name"))
}

// Tags returns a reference to field tags of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpc.ref.Append("tags"))
}

// VcenterCertificateThumbprint returns a reference to field vcenter_certificate_thumbprint of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) VcenterCertificateThumbprint() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vcenter_certificate_thumbprint"))
}

// VcenterPassword returns a reference to field vcenter_password of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) VcenterPassword() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vcenter_password"))
}

// VcsaEndpoint returns a reference to field vcsa_endpoint of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) VcsaEndpoint() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vcsa_endpoint"))
}

// VmotionSubnetCidr returns a reference to field vmotion_subnet_cidr of azurerm_vmware_private_cloud.
func (vpc vmwarePrivateCloudAttributes) VmotionSubnetCidr() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vmotion_subnet_cidr"))
}

func (vpc vmwarePrivateCloudAttributes) Circuit() terra.ListValue[vmwareprivatecloud.CircuitAttributes] {
	return terra.ReferenceAsList[vmwareprivatecloud.CircuitAttributes](vpc.ref.Append("circuit"))
}

func (vpc vmwarePrivateCloudAttributes) ManagementCluster() terra.ListValue[vmwareprivatecloud.ManagementClusterAttributes] {
	return terra.ReferenceAsList[vmwareprivatecloud.ManagementClusterAttributes](vpc.ref.Append("management_cluster"))
}

func (vpc vmwarePrivateCloudAttributes) Timeouts() vmwareprivatecloud.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vmwareprivatecloud.TimeoutsAttributes](vpc.ref.Append("timeouts"))
}

type vmwarePrivateCloudState struct {
	HcxCloudManagerEndpoint      string                                      `json:"hcx_cloud_manager_endpoint"`
	Id                           string                                      `json:"id"`
	InternetConnectionEnabled    bool                                        `json:"internet_connection_enabled"`
	Location                     string                                      `json:"location"`
	ManagementSubnetCidr         string                                      `json:"management_subnet_cidr"`
	Name                         string                                      `json:"name"`
	NetworkSubnetCidr            string                                      `json:"network_subnet_cidr"`
	NsxtCertificateThumbprint    string                                      `json:"nsxt_certificate_thumbprint"`
	NsxtManagerEndpoint          string                                      `json:"nsxt_manager_endpoint"`
	NsxtPassword                 string                                      `json:"nsxt_password"`
	ProvisioningSubnetCidr       string                                      `json:"provisioning_subnet_cidr"`
	ResourceGroupName            string                                      `json:"resource_group_name"`
	SkuName                      string                                      `json:"sku_name"`
	Tags                         map[string]string                           `json:"tags"`
	VcenterCertificateThumbprint string                                      `json:"vcenter_certificate_thumbprint"`
	VcenterPassword              string                                      `json:"vcenter_password"`
	VcsaEndpoint                 string                                      `json:"vcsa_endpoint"`
	VmotionSubnetCidr            string                                      `json:"vmotion_subnet_cidr"`
	Circuit                      []vmwareprivatecloud.CircuitState           `json:"circuit"`
	ManagementCluster            []vmwareprivatecloud.ManagementClusterState `json:"management_cluster"`
	Timeouts                     *vmwareprivatecloud.TimeoutsState           `json:"timeouts"`
}
