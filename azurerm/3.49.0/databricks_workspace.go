// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	databricksworkspace "github.com/golingon/terraproviders/azurerm/3.49.0/databricksworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDatabricksWorkspace creates a new instance of [DatabricksWorkspace].
func NewDatabricksWorkspace(name string, args DatabricksWorkspaceArgs) *DatabricksWorkspace {
	return &DatabricksWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatabricksWorkspace)(nil)

// DatabricksWorkspace represents the Terraform resource azurerm_databricks_workspace.
type DatabricksWorkspace struct {
	Name      string
	Args      DatabricksWorkspaceArgs
	state     *databricksWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DatabricksWorkspace].
func (dw *DatabricksWorkspace) Type() string {
	return "azurerm_databricks_workspace"
}

// LocalName returns the local name for [DatabricksWorkspace].
func (dw *DatabricksWorkspace) LocalName() string {
	return dw.Name
}

// Configuration returns the configuration (args) for [DatabricksWorkspace].
func (dw *DatabricksWorkspace) Configuration() interface{} {
	return dw.Args
}

// DependOn is used for other resources to depend on [DatabricksWorkspace].
func (dw *DatabricksWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(dw)
}

// Dependencies returns the list of resources [DatabricksWorkspace] depends_on.
func (dw *DatabricksWorkspace) Dependencies() terra.Dependencies {
	return dw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DatabricksWorkspace].
func (dw *DatabricksWorkspace) LifecycleManagement() *terra.Lifecycle {
	return dw.Lifecycle
}

// Attributes returns the attributes for [DatabricksWorkspace].
func (dw *DatabricksWorkspace) Attributes() databricksWorkspaceAttributes {
	return databricksWorkspaceAttributes{ref: terra.ReferenceResource(dw)}
}

// ImportState imports the given attribute values into [DatabricksWorkspace]'s state.
func (dw *DatabricksWorkspace) ImportState(av io.Reader) error {
	dw.state = &databricksWorkspaceState{}
	if err := json.NewDecoder(av).Decode(dw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dw.Type(), dw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DatabricksWorkspace] has state.
func (dw *DatabricksWorkspace) State() (*databricksWorkspaceState, bool) {
	return dw.state, dw.state != nil
}

// StateMust returns the state for [DatabricksWorkspace]. Panics if the state is nil.
func (dw *DatabricksWorkspace) StateMust() *databricksWorkspaceState {
	if dw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dw.Type(), dw.LocalName()))
	}
	return dw.state
}

// DatabricksWorkspaceArgs contains the configurations for azurerm_databricks_workspace.
type DatabricksWorkspaceArgs struct {
	// CustomerManagedKeyEnabled: bool, optional
	CustomerManagedKeyEnabled terra.BoolValue `hcl:"customer_managed_key_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InfrastructureEncryptionEnabled: bool, optional
	InfrastructureEncryptionEnabled terra.BoolValue `hcl:"infrastructure_encryption_enabled,attr"`
	// LoadBalancerBackendAddressPoolId: string, optional
	LoadBalancerBackendAddressPoolId terra.StringValue `hcl:"load_balancer_backend_address_pool_id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedDiskCmkKeyVaultKeyId: string, optional
	ManagedDiskCmkKeyVaultKeyId terra.StringValue `hcl:"managed_disk_cmk_key_vault_key_id,attr"`
	// ManagedDiskCmkRotationToLatestVersionEnabled: bool, optional
	ManagedDiskCmkRotationToLatestVersionEnabled terra.BoolValue `hcl:"managed_disk_cmk_rotation_to_latest_version_enabled,attr"`
	// ManagedResourceGroupName: string, optional
	ManagedResourceGroupName terra.StringValue `hcl:"managed_resource_group_name,attr"`
	// ManagedServicesCmkKeyVaultKeyId: string, optional
	ManagedServicesCmkKeyVaultKeyId terra.StringValue `hcl:"managed_services_cmk_key_vault_key_id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkSecurityGroupRulesRequired: string, optional
	NetworkSecurityGroupRulesRequired terra.StringValue `hcl:"network_security_group_rules_required,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, required
	Sku terra.StringValue `hcl:"sku,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ManagedDiskIdentity: min=0
	ManagedDiskIdentity []databricksworkspace.ManagedDiskIdentity `hcl:"managed_disk_identity,block" validate:"min=0"`
	// StorageAccountIdentity: min=0
	StorageAccountIdentity []databricksworkspace.StorageAccountIdentity `hcl:"storage_account_identity,block" validate:"min=0"`
	// CustomParameters: optional
	CustomParameters *databricksworkspace.CustomParameters `hcl:"custom_parameters,block"`
	// Timeouts: optional
	Timeouts *databricksworkspace.Timeouts `hcl:"timeouts,block"`
}
type databricksWorkspaceAttributes struct {
	ref terra.Reference
}

// CustomerManagedKeyEnabled returns a reference to field customer_managed_key_enabled of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) CustomerManagedKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dw.ref.Append("customer_managed_key_enabled"))
}

// DiskEncryptionSetId returns a reference to field disk_encryption_set_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) DiskEncryptionSetId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("disk_encryption_set_id"))
}

// Id returns a reference to field id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("id"))
}

// InfrastructureEncryptionEnabled returns a reference to field infrastructure_encryption_enabled of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) InfrastructureEncryptionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dw.ref.Append("infrastructure_encryption_enabled"))
}

// LoadBalancerBackendAddressPoolId returns a reference to field load_balancer_backend_address_pool_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) LoadBalancerBackendAddressPoolId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("load_balancer_backend_address_pool_id"))
}

// Location returns a reference to field location of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("location"))
}

// ManagedDiskCmkKeyVaultKeyId returns a reference to field managed_disk_cmk_key_vault_key_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ManagedDiskCmkKeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("managed_disk_cmk_key_vault_key_id"))
}

// ManagedDiskCmkRotationToLatestVersionEnabled returns a reference to field managed_disk_cmk_rotation_to_latest_version_enabled of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ManagedDiskCmkRotationToLatestVersionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dw.ref.Append("managed_disk_cmk_rotation_to_latest_version_enabled"))
}

// ManagedResourceGroupId returns a reference to field managed_resource_group_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ManagedResourceGroupId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("managed_resource_group_id"))
}

// ManagedResourceGroupName returns a reference to field managed_resource_group_name of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ManagedResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("managed_resource_group_name"))
}

// ManagedServicesCmkKeyVaultKeyId returns a reference to field managed_services_cmk_key_vault_key_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ManagedServicesCmkKeyVaultKeyId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("managed_services_cmk_key_vault_key_id"))
}

// Name returns a reference to field name of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("name"))
}

// NetworkSecurityGroupRulesRequired returns a reference to field network_security_group_rules_required of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) NetworkSecurityGroupRulesRequired() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("network_security_group_rules_required"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(dw.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dw.ref.Append("tags"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("workspace_id"))
}

// WorkspaceUrl returns a reference to field workspace_url of azurerm_databricks_workspace.
func (dw databricksWorkspaceAttributes) WorkspaceUrl() terra.StringValue {
	return terra.ReferenceAsString(dw.ref.Append("workspace_url"))
}

func (dw databricksWorkspaceAttributes) ManagedDiskIdentity() terra.ListValue[databricksworkspace.ManagedDiskIdentityAttributes] {
	return terra.ReferenceAsList[databricksworkspace.ManagedDiskIdentityAttributes](dw.ref.Append("managed_disk_identity"))
}

func (dw databricksWorkspaceAttributes) StorageAccountIdentity() terra.ListValue[databricksworkspace.StorageAccountIdentityAttributes] {
	return terra.ReferenceAsList[databricksworkspace.StorageAccountIdentityAttributes](dw.ref.Append("storage_account_identity"))
}

func (dw databricksWorkspaceAttributes) CustomParameters() terra.ListValue[databricksworkspace.CustomParametersAttributes] {
	return terra.ReferenceAsList[databricksworkspace.CustomParametersAttributes](dw.ref.Append("custom_parameters"))
}

func (dw databricksWorkspaceAttributes) Timeouts() databricksworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[databricksworkspace.TimeoutsAttributes](dw.ref.Append("timeouts"))
}

type databricksWorkspaceState struct {
	CustomerManagedKeyEnabled                    bool                                              `json:"customer_managed_key_enabled"`
	DiskEncryptionSetId                          string                                            `json:"disk_encryption_set_id"`
	Id                                           string                                            `json:"id"`
	InfrastructureEncryptionEnabled              bool                                              `json:"infrastructure_encryption_enabled"`
	LoadBalancerBackendAddressPoolId             string                                            `json:"load_balancer_backend_address_pool_id"`
	Location                                     string                                            `json:"location"`
	ManagedDiskCmkKeyVaultKeyId                  string                                            `json:"managed_disk_cmk_key_vault_key_id"`
	ManagedDiskCmkRotationToLatestVersionEnabled bool                                              `json:"managed_disk_cmk_rotation_to_latest_version_enabled"`
	ManagedResourceGroupId                       string                                            `json:"managed_resource_group_id"`
	ManagedResourceGroupName                     string                                            `json:"managed_resource_group_name"`
	ManagedServicesCmkKeyVaultKeyId              string                                            `json:"managed_services_cmk_key_vault_key_id"`
	Name                                         string                                            `json:"name"`
	NetworkSecurityGroupRulesRequired            string                                            `json:"network_security_group_rules_required"`
	PublicNetworkAccessEnabled                   bool                                              `json:"public_network_access_enabled"`
	ResourceGroupName                            string                                            `json:"resource_group_name"`
	Sku                                          string                                            `json:"sku"`
	Tags                                         map[string]string                                 `json:"tags"`
	WorkspaceId                                  string                                            `json:"workspace_id"`
	WorkspaceUrl                                 string                                            `json:"workspace_url"`
	ManagedDiskIdentity                          []databricksworkspace.ManagedDiskIdentityState    `json:"managed_disk_identity"`
	StorageAccountIdentity                       []databricksworkspace.StorageAccountIdentityState `json:"storage_account_identity"`
	CustomParameters                             []databricksworkspace.CustomParametersState       `json:"custom_parameters"`
	Timeouts                                     *databricksworkspace.TimeoutsState                `json:"timeouts"`
}