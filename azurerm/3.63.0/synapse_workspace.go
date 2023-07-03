// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseworkspace "github.com/golingon/terraproviders/azurerm/3.63.0/synapseworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseWorkspace creates a new instance of [SynapseWorkspace].
func NewSynapseWorkspace(name string, args SynapseWorkspaceArgs) *SynapseWorkspace {
	return &SynapseWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseWorkspace)(nil)

// SynapseWorkspace represents the Terraform resource azurerm_synapse_workspace.
type SynapseWorkspace struct {
	Name      string
	Args      SynapseWorkspaceArgs
	state     *synapseWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseWorkspace].
func (sw *SynapseWorkspace) Type() string {
	return "azurerm_synapse_workspace"
}

// LocalName returns the local name for [SynapseWorkspace].
func (sw *SynapseWorkspace) LocalName() string {
	return sw.Name
}

// Configuration returns the configuration (args) for [SynapseWorkspace].
func (sw *SynapseWorkspace) Configuration() interface{} {
	return sw.Args
}

// DependOn is used for other resources to depend on [SynapseWorkspace].
func (sw *SynapseWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(sw)
}

// Dependencies returns the list of resources [SynapseWorkspace] depends_on.
func (sw *SynapseWorkspace) Dependencies() terra.Dependencies {
	return sw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseWorkspace].
func (sw *SynapseWorkspace) LifecycleManagement() *terra.Lifecycle {
	return sw.Lifecycle
}

// Attributes returns the attributes for [SynapseWorkspace].
func (sw *SynapseWorkspace) Attributes() synapseWorkspaceAttributes {
	return synapseWorkspaceAttributes{ref: terra.ReferenceResource(sw)}
}

// ImportState imports the given attribute values into [SynapseWorkspace]'s state.
func (sw *SynapseWorkspace) ImportState(av io.Reader) error {
	sw.state = &synapseWorkspaceState{}
	if err := json.NewDecoder(av).Decode(sw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sw.Type(), sw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseWorkspace] has state.
func (sw *SynapseWorkspace) State() (*synapseWorkspaceState, bool) {
	return sw.state, sw.state != nil
}

// StateMust returns the state for [SynapseWorkspace]. Panics if the state is nil.
func (sw *SynapseWorkspace) StateMust() *synapseWorkspaceState {
	if sw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sw.Type(), sw.LocalName()))
	}
	return sw.state
}

// SynapseWorkspaceArgs contains the configurations for azurerm_synapse_workspace.
type SynapseWorkspaceArgs struct {
	// ComputeSubnetId: string, optional
	ComputeSubnetId terra.StringValue `hcl:"compute_subnet_id,attr"`
	// DataExfiltrationProtectionEnabled: bool, optional
	DataExfiltrationProtectionEnabled terra.BoolValue `hcl:"data_exfiltration_protection_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkingAllowedForAadTenantIds: list of string, optional
	LinkingAllowedForAadTenantIds terra.ListValue[terra.StringValue] `hcl:"linking_allowed_for_aad_tenant_ids,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagedResourceGroupName: string, optional
	ManagedResourceGroupName terra.StringValue `hcl:"managed_resource_group_name,attr"`
	// ManagedVirtualNetworkEnabled: bool, optional
	ManagedVirtualNetworkEnabled terra.BoolValue `hcl:"managed_virtual_network_enabled,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// PurviewId: string, optional
	PurviewId terra.StringValue `hcl:"purview_id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SqlAdministratorLogin: string, optional
	SqlAdministratorLogin terra.StringValue `hcl:"sql_administrator_login,attr"`
	// SqlAdministratorLoginPassword: string, optional
	SqlAdministratorLoginPassword terra.StringValue `hcl:"sql_administrator_login_password,attr"`
	// SqlIdentityControlEnabled: bool, optional
	SqlIdentityControlEnabled terra.BoolValue `hcl:"sql_identity_control_enabled,attr"`
	// StorageDataLakeGen2FilesystemId: string, required
	StorageDataLakeGen2FilesystemId terra.StringValue `hcl:"storage_data_lake_gen2_filesystem_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// AadAdmin: min=0
	AadAdmin []synapseworkspace.AadAdmin `hcl:"aad_admin,block" validate:"min=0"`
	// SqlAadAdmin: min=0
	SqlAadAdmin []synapseworkspace.SqlAadAdmin `hcl:"sql_aad_admin,block" validate:"min=0"`
	// AzureDevopsRepo: optional
	AzureDevopsRepo *synapseworkspace.AzureDevopsRepo `hcl:"azure_devops_repo,block"`
	// CustomerManagedKey: optional
	CustomerManagedKey *synapseworkspace.CustomerManagedKey `hcl:"customer_managed_key,block"`
	// GithubRepo: optional
	GithubRepo *synapseworkspace.GithubRepo `hcl:"github_repo,block"`
	// Identity: optional
	Identity *synapseworkspace.Identity `hcl:"identity,block"`
	// Timeouts: optional
	Timeouts *synapseworkspace.Timeouts `hcl:"timeouts,block"`
}
type synapseWorkspaceAttributes struct {
	ref terra.Reference
}

// ComputeSubnetId returns a reference to field compute_subnet_id of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) ComputeSubnetId() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("compute_subnet_id"))
}

// ConnectivityEndpoints returns a reference to field connectivity_endpoints of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) ConnectivityEndpoints() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("connectivity_endpoints"))
}

// DataExfiltrationProtectionEnabled returns a reference to field data_exfiltration_protection_enabled of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) DataExfiltrationProtectionEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sw.ref.Append("data_exfiltration_protection_enabled"))
}

// Id returns a reference to field id of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("id"))
}

// LinkingAllowedForAadTenantIds returns a reference to field linking_allowed_for_aad_tenant_ids of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) LinkingAllowedForAadTenantIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sw.ref.Append("linking_allowed_for_aad_tenant_ids"))
}

// Location returns a reference to field location of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("location"))
}

// ManagedResourceGroupName returns a reference to field managed_resource_group_name of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) ManagedResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("managed_resource_group_name"))
}

// ManagedVirtualNetworkEnabled returns a reference to field managed_virtual_network_enabled of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) ManagedVirtualNetworkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sw.ref.Append("managed_virtual_network_enabled"))
}

// Name returns a reference to field name of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("name"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sw.ref.Append("public_network_access_enabled"))
}

// PurviewId returns a reference to field purview_id of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) PurviewId() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("purview_id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("resource_group_name"))
}

// SqlAdministratorLogin returns a reference to field sql_administrator_login of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) SqlAdministratorLogin() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("sql_administrator_login"))
}

// SqlAdministratorLoginPassword returns a reference to field sql_administrator_login_password of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) SqlAdministratorLoginPassword() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("sql_administrator_login_password"))
}

// SqlIdentityControlEnabled returns a reference to field sql_identity_control_enabled of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) SqlIdentityControlEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sw.ref.Append("sql_identity_control_enabled"))
}

// StorageDataLakeGen2FilesystemId returns a reference to field storage_data_lake_gen2_filesystem_id of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) StorageDataLakeGen2FilesystemId() terra.StringValue {
	return terra.ReferenceAsString(sw.ref.Append("storage_data_lake_gen2_filesystem_id"))
}

// Tags returns a reference to field tags of azurerm_synapse_workspace.
func (sw synapseWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sw.ref.Append("tags"))
}

func (sw synapseWorkspaceAttributes) AadAdmin() terra.ListValue[synapseworkspace.AadAdminAttributes] {
	return terra.ReferenceAsList[synapseworkspace.AadAdminAttributes](sw.ref.Append("aad_admin"))
}

func (sw synapseWorkspaceAttributes) SqlAadAdmin() terra.ListValue[synapseworkspace.SqlAadAdminAttributes] {
	return terra.ReferenceAsList[synapseworkspace.SqlAadAdminAttributes](sw.ref.Append("sql_aad_admin"))
}

func (sw synapseWorkspaceAttributes) AzureDevopsRepo() terra.ListValue[synapseworkspace.AzureDevopsRepoAttributes] {
	return terra.ReferenceAsList[synapseworkspace.AzureDevopsRepoAttributes](sw.ref.Append("azure_devops_repo"))
}

func (sw synapseWorkspaceAttributes) CustomerManagedKey() terra.ListValue[synapseworkspace.CustomerManagedKeyAttributes] {
	return terra.ReferenceAsList[synapseworkspace.CustomerManagedKeyAttributes](sw.ref.Append("customer_managed_key"))
}

func (sw synapseWorkspaceAttributes) GithubRepo() terra.ListValue[synapseworkspace.GithubRepoAttributes] {
	return terra.ReferenceAsList[synapseworkspace.GithubRepoAttributes](sw.ref.Append("github_repo"))
}

func (sw synapseWorkspaceAttributes) Identity() terra.ListValue[synapseworkspace.IdentityAttributes] {
	return terra.ReferenceAsList[synapseworkspace.IdentityAttributes](sw.ref.Append("identity"))
}

func (sw synapseWorkspaceAttributes) Timeouts() synapseworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseworkspace.TimeoutsAttributes](sw.ref.Append("timeouts"))
}

type synapseWorkspaceState struct {
	ComputeSubnetId                   string                                     `json:"compute_subnet_id"`
	ConnectivityEndpoints             map[string]string                          `json:"connectivity_endpoints"`
	DataExfiltrationProtectionEnabled bool                                       `json:"data_exfiltration_protection_enabled"`
	Id                                string                                     `json:"id"`
	LinkingAllowedForAadTenantIds     []string                                   `json:"linking_allowed_for_aad_tenant_ids"`
	Location                          string                                     `json:"location"`
	ManagedResourceGroupName          string                                     `json:"managed_resource_group_name"`
	ManagedVirtualNetworkEnabled      bool                                       `json:"managed_virtual_network_enabled"`
	Name                              string                                     `json:"name"`
	PublicNetworkAccessEnabled        bool                                       `json:"public_network_access_enabled"`
	PurviewId                         string                                     `json:"purview_id"`
	ResourceGroupName                 string                                     `json:"resource_group_name"`
	SqlAdministratorLogin             string                                     `json:"sql_administrator_login"`
	SqlAdministratorLoginPassword     string                                     `json:"sql_administrator_login_password"`
	SqlIdentityControlEnabled         bool                                       `json:"sql_identity_control_enabled"`
	StorageDataLakeGen2FilesystemId   string                                     `json:"storage_data_lake_gen2_filesystem_id"`
	Tags                              map[string]string                          `json:"tags"`
	AadAdmin                          []synapseworkspace.AadAdminState           `json:"aad_admin"`
	SqlAadAdmin                       []synapseworkspace.SqlAadAdminState        `json:"sql_aad_admin"`
	AzureDevopsRepo                   []synapseworkspace.AzureDevopsRepoState    `json:"azure_devops_repo"`
	CustomerManagedKey                []synapseworkspace.CustomerManagedKeyState `json:"customer_managed_key"`
	GithubRepo                        []synapseworkspace.GithubRepoState         `json:"github_repo"`
	Identity                          []synapseworkspace.IdentityState           `json:"identity"`
	Timeouts                          *synapseworkspace.TimeoutsState            `json:"timeouts"`
}
