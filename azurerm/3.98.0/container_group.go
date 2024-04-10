// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	containergroup "github.com/golingon/terraproviders/azurerm/3.98.0/containergroup"
	"io"
)

// NewContainerGroup creates a new instance of [ContainerGroup].
func NewContainerGroup(name string, args ContainerGroupArgs) *ContainerGroup {
	return &ContainerGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerGroup)(nil)

// ContainerGroup represents the Terraform resource azurerm_container_group.
type ContainerGroup struct {
	Name      string
	Args      ContainerGroupArgs
	state     *containerGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ContainerGroup].
func (cg *ContainerGroup) Type() string {
	return "azurerm_container_group"
}

// LocalName returns the local name for [ContainerGroup].
func (cg *ContainerGroup) LocalName() string {
	return cg.Name
}

// Configuration returns the configuration (args) for [ContainerGroup].
func (cg *ContainerGroup) Configuration() interface{} {
	return cg.Args
}

// DependOn is used for other resources to depend on [ContainerGroup].
func (cg *ContainerGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(cg)
}

// Dependencies returns the list of resources [ContainerGroup] depends_on.
func (cg *ContainerGroup) Dependencies() terra.Dependencies {
	return cg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ContainerGroup].
func (cg *ContainerGroup) LifecycleManagement() *terra.Lifecycle {
	return cg.Lifecycle
}

// Attributes returns the attributes for [ContainerGroup].
func (cg *ContainerGroup) Attributes() containerGroupAttributes {
	return containerGroupAttributes{ref: terra.ReferenceResource(cg)}
}

// ImportState imports the given attribute values into [ContainerGroup]'s state.
func (cg *ContainerGroup) ImportState(av io.Reader) error {
	cg.state = &containerGroupState{}
	if err := json.NewDecoder(av).Decode(cg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cg.Type(), cg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ContainerGroup] has state.
func (cg *ContainerGroup) State() (*containerGroupState, bool) {
	return cg.state, cg.state != nil
}

// StateMust returns the state for [ContainerGroup]. Panics if the state is nil.
func (cg *ContainerGroup) StateMust() *containerGroupState {
	if cg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cg.Type(), cg.LocalName()))
	}
	return cg.state
}

// ContainerGroupArgs contains the configurations for azurerm_container_group.
type ContainerGroupArgs struct {
	// DnsNameLabel: string, optional
	DnsNameLabel terra.StringValue `hcl:"dns_name_label,attr"`
	// DnsNameLabelReusePolicy: string, optional
	DnsNameLabelReusePolicy terra.StringValue `hcl:"dns_name_label_reuse_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddressType: string, optional
	IpAddressType terra.StringValue `hcl:"ip_address_type,attr"`
	// KeyVaultKeyId: string, optional
	KeyVaultKeyId terra.StringValue `hcl:"key_vault_key_id,attr"`
	// KeyVaultUserAssignedIdentityId: string, optional
	KeyVaultUserAssignedIdentityId terra.StringValue `hcl:"key_vault_user_assigned_identity_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkProfileId: string, optional
	NetworkProfileId terra.StringValue `hcl:"network_profile_id,attr"`
	// OsType: string, required
	OsType terra.StringValue `hcl:"os_type,attr" validate:"required"`
	// Priority: string, optional
	Priority terra.StringValue `hcl:"priority,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RestartPolicy: string, optional
	RestartPolicy terra.StringValue `hcl:"restart_policy,attr"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// SubnetIds: set of string, optional
	SubnetIds terra.SetValue[terra.StringValue] `hcl:"subnet_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// ExposedPort: min=0
	ExposedPort []containergroup.ExposedPort `hcl:"exposed_port,block" validate:"min=0"`
	// Container: min=1
	Container []containergroup.Container `hcl:"container,block" validate:"min=1"`
	// Diagnostics: optional
	Diagnostics *containergroup.Diagnostics `hcl:"diagnostics,block"`
	// DnsConfig: optional
	DnsConfig *containergroup.DnsConfig `hcl:"dns_config,block"`
	// Identity: optional
	Identity *containergroup.Identity `hcl:"identity,block"`
	// ImageRegistryCredential: min=0
	ImageRegistryCredential []containergroup.ImageRegistryCredential `hcl:"image_registry_credential,block" validate:"min=0"`
	// InitContainer: min=0
	InitContainer []containergroup.InitContainer `hcl:"init_container,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *containergroup.Timeouts `hcl:"timeouts,block"`
}
type containerGroupAttributes struct {
	ref terra.Reference
}

// DnsNameLabel returns a reference to field dns_name_label of azurerm_container_group.
func (cg containerGroupAttributes) DnsNameLabel() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("dns_name_label"))
}

// DnsNameLabelReusePolicy returns a reference to field dns_name_label_reuse_policy of azurerm_container_group.
func (cg containerGroupAttributes) DnsNameLabelReusePolicy() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("dns_name_label_reuse_policy"))
}

// Fqdn returns a reference to field fqdn of azurerm_container_group.
func (cg containerGroupAttributes) Fqdn() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("fqdn"))
}

// Id returns a reference to field id of azurerm_container_group.
func (cg containerGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of azurerm_container_group.
func (cg containerGroupAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("ip_address"))
}

// IpAddressType returns a reference to field ip_address_type of azurerm_container_group.
func (cg containerGroupAttributes) IpAddressType() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("ip_address_type"))
}

// KeyVaultKeyId returns a reference to field key_vault_key_id of azurerm_container_group.
func (cg containerGroupAttributes) KeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("key_vault_key_id"))
}

// KeyVaultUserAssignedIdentityId returns a reference to field key_vault_user_assigned_identity_id of azurerm_container_group.
func (cg containerGroupAttributes) KeyVaultUserAssignedIdentityId() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("key_vault_user_assigned_identity_id"))
}

// Location returns a reference to field location of azurerm_container_group.
func (cg containerGroupAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_container_group.
func (cg containerGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("name"))
}

// NetworkProfileId returns a reference to field network_profile_id of azurerm_container_group.
func (cg containerGroupAttributes) NetworkProfileId() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("network_profile_id"))
}

// OsType returns a reference to field os_type of azurerm_container_group.
func (cg containerGroupAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("os_type"))
}

// Priority returns a reference to field priority of azurerm_container_group.
func (cg containerGroupAttributes) Priority() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("priority"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_group.
func (cg containerGroupAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("resource_group_name"))
}

// RestartPolicy returns a reference to field restart_policy of azurerm_container_group.
func (cg containerGroupAttributes) RestartPolicy() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("restart_policy"))
}

// Sku returns a reference to field sku of azurerm_container_group.
func (cg containerGroupAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(cg.ref.Append("sku"))
}

// SubnetIds returns a reference to field subnet_ids of azurerm_container_group.
func (cg containerGroupAttributes) SubnetIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cg.ref.Append("subnet_ids"))
}

// Tags returns a reference to field tags of azurerm_container_group.
func (cg containerGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cg.ref.Append("tags"))
}

// Zones returns a reference to field zones of azurerm_container_group.
func (cg containerGroupAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cg.ref.Append("zones"))
}

func (cg containerGroupAttributes) ExposedPort() terra.SetValue[containergroup.ExposedPortAttributes] {
	return terra.ReferenceAsSet[containergroup.ExposedPortAttributes](cg.ref.Append("exposed_port"))
}

func (cg containerGroupAttributes) Container() terra.ListValue[containergroup.ContainerAttributes] {
	return terra.ReferenceAsList[containergroup.ContainerAttributes](cg.ref.Append("container"))
}

func (cg containerGroupAttributes) Diagnostics() terra.ListValue[containergroup.DiagnosticsAttributes] {
	return terra.ReferenceAsList[containergroup.DiagnosticsAttributes](cg.ref.Append("diagnostics"))
}

func (cg containerGroupAttributes) DnsConfig() terra.ListValue[containergroup.DnsConfigAttributes] {
	return terra.ReferenceAsList[containergroup.DnsConfigAttributes](cg.ref.Append("dns_config"))
}

func (cg containerGroupAttributes) Identity() terra.ListValue[containergroup.IdentityAttributes] {
	return terra.ReferenceAsList[containergroup.IdentityAttributes](cg.ref.Append("identity"))
}

func (cg containerGroupAttributes) ImageRegistryCredential() terra.ListValue[containergroup.ImageRegistryCredentialAttributes] {
	return terra.ReferenceAsList[containergroup.ImageRegistryCredentialAttributes](cg.ref.Append("image_registry_credential"))
}

func (cg containerGroupAttributes) InitContainer() terra.ListValue[containergroup.InitContainerAttributes] {
	return terra.ReferenceAsList[containergroup.InitContainerAttributes](cg.ref.Append("init_container"))
}

func (cg containerGroupAttributes) Timeouts() containergroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[containergroup.TimeoutsAttributes](cg.ref.Append("timeouts"))
}

type containerGroupState struct {
	DnsNameLabel                   string                                        `json:"dns_name_label"`
	DnsNameLabelReusePolicy        string                                        `json:"dns_name_label_reuse_policy"`
	Fqdn                           string                                        `json:"fqdn"`
	Id                             string                                        `json:"id"`
	IpAddress                      string                                        `json:"ip_address"`
	IpAddressType                  string                                        `json:"ip_address_type"`
	KeyVaultKeyId                  string                                        `json:"key_vault_key_id"`
	KeyVaultUserAssignedIdentityId string                                        `json:"key_vault_user_assigned_identity_id"`
	Location                       string                                        `json:"location"`
	Name                           string                                        `json:"name"`
	NetworkProfileId               string                                        `json:"network_profile_id"`
	OsType                         string                                        `json:"os_type"`
	Priority                       string                                        `json:"priority"`
	ResourceGroupName              string                                        `json:"resource_group_name"`
	RestartPolicy                  string                                        `json:"restart_policy"`
	Sku                            string                                        `json:"sku"`
	SubnetIds                      []string                                      `json:"subnet_ids"`
	Tags                           map[string]string                             `json:"tags"`
	Zones                          []string                                      `json:"zones"`
	ExposedPort                    []containergroup.ExposedPortState             `json:"exposed_port"`
	Container                      []containergroup.ContainerState               `json:"container"`
	Diagnostics                    []containergroup.DiagnosticsState             `json:"diagnostics"`
	DnsConfig                      []containergroup.DnsConfigState               `json:"dns_config"`
	Identity                       []containergroup.IdentityState                `json:"identity"`
	ImageRegistryCredential        []containergroup.ImageRegistryCredentialState `json:"image_registry_credential"`
	InitContainer                  []containergroup.InitContainerState           `json:"init_container"`
	Timeouts                       *containergroup.TimeoutsState                 `json:"timeouts"`
}
