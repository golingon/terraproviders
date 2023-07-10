// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningworkspace "github.com/golingon/terraproviders/azurerm/3.64.0/machinelearningworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMachineLearningWorkspace creates a new instance of [MachineLearningWorkspace].
func NewMachineLearningWorkspace(name string, args MachineLearningWorkspaceArgs) *MachineLearningWorkspace {
	return &MachineLearningWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningWorkspace)(nil)

// MachineLearningWorkspace represents the Terraform resource azurerm_machine_learning_workspace.
type MachineLearningWorkspace struct {
	Name      string
	Args      MachineLearningWorkspaceArgs
	state     *machineLearningWorkspaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) Type() string {
	return "azurerm_machine_learning_workspace"
}

// LocalName returns the local name for [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) LocalName() string {
	return mlw.Name
}

// Configuration returns the configuration (args) for [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) Configuration() interface{} {
	return mlw.Args
}

// DependOn is used for other resources to depend on [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(mlw)
}

// Dependencies returns the list of resources [MachineLearningWorkspace] depends_on.
func (mlw *MachineLearningWorkspace) Dependencies() terra.Dependencies {
	return mlw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) LifecycleManagement() *terra.Lifecycle {
	return mlw.Lifecycle
}

// Attributes returns the attributes for [MachineLearningWorkspace].
func (mlw *MachineLearningWorkspace) Attributes() machineLearningWorkspaceAttributes {
	return machineLearningWorkspaceAttributes{ref: terra.ReferenceResource(mlw)}
}

// ImportState imports the given attribute values into [MachineLearningWorkspace]'s state.
func (mlw *MachineLearningWorkspace) ImportState(av io.Reader) error {
	mlw.state = &machineLearningWorkspaceState{}
	if err := json.NewDecoder(av).Decode(mlw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlw.Type(), mlw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MachineLearningWorkspace] has state.
func (mlw *MachineLearningWorkspace) State() (*machineLearningWorkspaceState, bool) {
	return mlw.state, mlw.state != nil
}

// StateMust returns the state for [MachineLearningWorkspace]. Panics if the state is nil.
func (mlw *MachineLearningWorkspace) StateMust() *machineLearningWorkspaceState {
	if mlw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlw.Type(), mlw.LocalName()))
	}
	return mlw.state
}

// MachineLearningWorkspaceArgs contains the configurations for azurerm_machine_learning_workspace.
type MachineLearningWorkspaceArgs struct {
	// ApplicationInsightsId: string, required
	ApplicationInsightsId terra.StringValue `hcl:"application_insights_id,attr" validate:"required"`
	// ContainerRegistryId: string, optional
	ContainerRegistryId terra.StringValue `hcl:"container_registry_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// FriendlyName: string, optional
	FriendlyName terra.StringValue `hcl:"friendly_name,attr"`
	// HighBusinessImpact: bool, optional
	HighBusinessImpact terra.BoolValue `hcl:"high_business_impact,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ImageBuildComputeName: string, optional
	ImageBuildComputeName terra.StringValue `hcl:"image_build_compute_name,attr"`
	// KeyVaultId: string, required
	KeyVaultId terra.StringValue `hcl:"key_vault_id,attr" validate:"required"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrimaryUserAssignedIdentity: string, optional
	PrimaryUserAssignedIdentity terra.StringValue `hcl:"primary_user_assigned_identity,attr"`
	// PublicAccessBehindVirtualNetworkEnabled: bool, optional
	PublicAccessBehindVirtualNetworkEnabled terra.BoolValue `hcl:"public_access_behind_virtual_network_enabled,attr"`
	// PublicNetworkAccessEnabled: bool, optional
	PublicNetworkAccessEnabled terra.BoolValue `hcl:"public_network_access_enabled,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, optional
	SkuName terra.StringValue `hcl:"sku_name,attr"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// V1LegacyModeEnabled: bool, optional
	V1LegacyModeEnabled terra.BoolValue `hcl:"v1_legacy_mode_enabled,attr"`
	// Encryption: optional
	Encryption *machinelearningworkspace.Encryption `hcl:"encryption,block"`
	// Identity: required
	Identity *machinelearningworkspace.Identity `hcl:"identity,block" validate:"required"`
	// Timeouts: optional
	Timeouts *machinelearningworkspace.Timeouts `hcl:"timeouts,block"`
}
type machineLearningWorkspaceAttributes struct {
	ref terra.Reference
}

// ApplicationInsightsId returns a reference to field application_insights_id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("application_insights_id"))
}

// ContainerRegistryId returns a reference to field container_registry_id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) ContainerRegistryId() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("container_registry_id"))
}

// Description returns a reference to field description of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("description"))
}

// DiscoveryUrl returns a reference to field discovery_url of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) DiscoveryUrl() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("discovery_url"))
}

// FriendlyName returns a reference to field friendly_name of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("friendly_name"))
}

// HighBusinessImpact returns a reference to field high_business_impact of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) HighBusinessImpact() terra.BoolValue {
	return terra.ReferenceAsBool(mlw.ref.Append("high_business_impact"))
}

// Id returns a reference to field id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("id"))
}

// ImageBuildComputeName returns a reference to field image_build_compute_name of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) ImageBuildComputeName() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("image_build_compute_name"))
}

// KeyVaultId returns a reference to field key_vault_id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("key_vault_id"))
}

// Location returns a reference to field location of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("name"))
}

// PrimaryUserAssignedIdentity returns a reference to field primary_user_assigned_identity of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) PrimaryUserAssignedIdentity() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("primary_user_assigned_identity"))
}

// PublicAccessBehindVirtualNetworkEnabled returns a reference to field public_access_behind_virtual_network_enabled of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) PublicAccessBehindVirtualNetworkEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mlw.ref.Append("public_access_behind_virtual_network_enabled"))
}

// PublicNetworkAccessEnabled returns a reference to field public_network_access_enabled of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mlw.ref.Append("public_network_access_enabled"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("sku_name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("storage_account_id"))
}

// Tags returns a reference to field tags of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mlw.ref.Append("tags"))
}

// V1LegacyModeEnabled returns a reference to field v1_legacy_mode_enabled of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) V1LegacyModeEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mlw.ref.Append("v1_legacy_mode_enabled"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_machine_learning_workspace.
func (mlw machineLearningWorkspaceAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(mlw.ref.Append("workspace_id"))
}

func (mlw machineLearningWorkspaceAttributes) Encryption() terra.ListValue[machinelearningworkspace.EncryptionAttributes] {
	return terra.ReferenceAsList[machinelearningworkspace.EncryptionAttributes](mlw.ref.Append("encryption"))
}

func (mlw machineLearningWorkspaceAttributes) Identity() terra.ListValue[machinelearningworkspace.IdentityAttributes] {
	return terra.ReferenceAsList[machinelearningworkspace.IdentityAttributes](mlw.ref.Append("identity"))
}

func (mlw machineLearningWorkspaceAttributes) Timeouts() machinelearningworkspace.TimeoutsAttributes {
	return terra.ReferenceAsSingle[machinelearningworkspace.TimeoutsAttributes](mlw.ref.Append("timeouts"))
}

type machineLearningWorkspaceState struct {
	ApplicationInsightsId                   string                                     `json:"application_insights_id"`
	ContainerRegistryId                     string                                     `json:"container_registry_id"`
	Description                             string                                     `json:"description"`
	DiscoveryUrl                            string                                     `json:"discovery_url"`
	FriendlyName                            string                                     `json:"friendly_name"`
	HighBusinessImpact                      bool                                       `json:"high_business_impact"`
	Id                                      string                                     `json:"id"`
	ImageBuildComputeName                   string                                     `json:"image_build_compute_name"`
	KeyVaultId                              string                                     `json:"key_vault_id"`
	Location                                string                                     `json:"location"`
	Name                                    string                                     `json:"name"`
	PrimaryUserAssignedIdentity             string                                     `json:"primary_user_assigned_identity"`
	PublicAccessBehindVirtualNetworkEnabled bool                                       `json:"public_access_behind_virtual_network_enabled"`
	PublicNetworkAccessEnabled              bool                                       `json:"public_network_access_enabled"`
	ResourceGroupName                       string                                     `json:"resource_group_name"`
	SkuName                                 string                                     `json:"sku_name"`
	StorageAccountId                        string                                     `json:"storage_account_id"`
	Tags                                    map[string]string                          `json:"tags"`
	V1LegacyModeEnabled                     bool                                       `json:"v1_legacy_mode_enabled"`
	WorkspaceId                             string                                     `json:"workspace_id"`
	Encryption                              []machinelearningworkspace.EncryptionState `json:"encryption"`
	Identity                                []machinelearningworkspace.IdentityState   `json:"identity"`
	Timeouts                                *machinelearningworkspace.TimeoutsState    `json:"timeouts"`
}
