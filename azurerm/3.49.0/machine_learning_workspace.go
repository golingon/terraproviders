// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	machinelearningworkspace "github.com/golingon/terraproviders/azurerm/3.49.0/machinelearningworkspace"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewMachineLearningWorkspace(name string, args MachineLearningWorkspaceArgs) *MachineLearningWorkspace {
	return &MachineLearningWorkspace{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MachineLearningWorkspace)(nil)

type MachineLearningWorkspace struct {
	Name  string
	Args  MachineLearningWorkspaceArgs
	state *machineLearningWorkspaceState
}

func (mlw *MachineLearningWorkspace) Type() string {
	return "azurerm_machine_learning_workspace"
}

func (mlw *MachineLearningWorkspace) LocalName() string {
	return mlw.Name
}

func (mlw *MachineLearningWorkspace) Configuration() interface{} {
	return mlw.Args
}

func (mlw *MachineLearningWorkspace) Attributes() machineLearningWorkspaceAttributes {
	return machineLearningWorkspaceAttributes{ref: terra.ReferenceResource(mlw)}
}

func (mlw *MachineLearningWorkspace) ImportState(av io.Reader) error {
	mlw.state = &machineLearningWorkspaceState{}
	if err := json.NewDecoder(av).Decode(mlw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mlw.Type(), mlw.LocalName(), err)
	}
	return nil
}

func (mlw *MachineLearningWorkspace) State() (*machineLearningWorkspaceState, bool) {
	return mlw.state, mlw.state != nil
}

func (mlw *MachineLearningWorkspace) StateMust() *machineLearningWorkspaceState {
	if mlw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mlw.Type(), mlw.LocalName()))
	}
	return mlw.state
}

func (mlw *MachineLearningWorkspace) DependOn() terra.Reference {
	return terra.ReferenceResource(mlw)
}

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
	// DependsOn contains resources that MachineLearningWorkspace depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type machineLearningWorkspaceAttributes struct {
	ref terra.Reference
}

func (mlw machineLearningWorkspaceAttributes) ApplicationInsightsId() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("application_insights_id"))
}

func (mlw machineLearningWorkspaceAttributes) ContainerRegistryId() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("container_registry_id"))
}

func (mlw machineLearningWorkspaceAttributes) Description() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("description"))
}

func (mlw machineLearningWorkspaceAttributes) DiscoveryUrl() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("discovery_url"))
}

func (mlw machineLearningWorkspaceAttributes) FriendlyName() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("friendly_name"))
}

func (mlw machineLearningWorkspaceAttributes) HighBusinessImpact() terra.BoolValue {
	return terra.ReferenceBool(mlw.ref.Append("high_business_impact"))
}

func (mlw machineLearningWorkspaceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("id"))
}

func (mlw machineLearningWorkspaceAttributes) ImageBuildComputeName() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("image_build_compute_name"))
}

func (mlw machineLearningWorkspaceAttributes) KeyVaultId() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("key_vault_id"))
}

func (mlw machineLearningWorkspaceAttributes) Location() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("location"))
}

func (mlw machineLearningWorkspaceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("name"))
}

func (mlw machineLearningWorkspaceAttributes) PrimaryUserAssignedIdentity() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("primary_user_assigned_identity"))
}

func (mlw machineLearningWorkspaceAttributes) PublicAccessBehindVirtualNetworkEnabled() terra.BoolValue {
	return terra.ReferenceBool(mlw.ref.Append("public_access_behind_virtual_network_enabled"))
}

func (mlw machineLearningWorkspaceAttributes) PublicNetworkAccessEnabled() terra.BoolValue {
	return terra.ReferenceBool(mlw.ref.Append("public_network_access_enabled"))
}

func (mlw machineLearningWorkspaceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("resource_group_name"))
}

func (mlw machineLearningWorkspaceAttributes) SkuName() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("sku_name"))
}

func (mlw machineLearningWorkspaceAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceString(mlw.ref.Append("storage_account_id"))
}

func (mlw machineLearningWorkspaceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](mlw.ref.Append("tags"))
}

func (mlw machineLearningWorkspaceAttributes) V1LegacyModeEnabled() terra.BoolValue {
	return terra.ReferenceBool(mlw.ref.Append("v1_legacy_mode_enabled"))
}

func (mlw machineLearningWorkspaceAttributes) Encryption() terra.ListValue[machinelearningworkspace.EncryptionAttributes] {
	return terra.ReferenceList[machinelearningworkspace.EncryptionAttributes](mlw.ref.Append("encryption"))
}

func (mlw machineLearningWorkspaceAttributes) Identity() terra.ListValue[machinelearningworkspace.IdentityAttributes] {
	return terra.ReferenceList[machinelearningworkspace.IdentityAttributes](mlw.ref.Append("identity"))
}

func (mlw machineLearningWorkspaceAttributes) Timeouts() machinelearningworkspace.TimeoutsAttributes {
	return terra.ReferenceSingle[machinelearningworkspace.TimeoutsAttributes](mlw.ref.Append("timeouts"))
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
	Encryption                              []machinelearningworkspace.EncryptionState `json:"encryption"`
	Identity                                []machinelearningworkspace.IdentityState   `json:"identity"`
	Timeouts                                *machinelearningworkspace.TimeoutsState    `json:"timeouts"`
}
