// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryprotectioncontainermapping "github.com/golingon/terraproviders/azurerm/3.66.0/siterecoveryprotectioncontainermapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryProtectionContainerMapping creates a new instance of [SiteRecoveryProtectionContainerMapping].
func NewSiteRecoveryProtectionContainerMapping(name string, args SiteRecoveryProtectionContainerMappingArgs) *SiteRecoveryProtectionContainerMapping {
	return &SiteRecoveryProtectionContainerMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryProtectionContainerMapping)(nil)

// SiteRecoveryProtectionContainerMapping represents the Terraform resource azurerm_site_recovery_protection_container_mapping.
type SiteRecoveryProtectionContainerMapping struct {
	Name      string
	Args      SiteRecoveryProtectionContainerMappingArgs
	state     *siteRecoveryProtectionContainerMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) Type() string {
	return "azurerm_site_recovery_protection_container_mapping"
}

// LocalName returns the local name for [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) LocalName() string {
	return srpcm.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) Configuration() interface{} {
	return srpcm.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(srpcm)
}

// Dependencies returns the list of resources [SiteRecoveryProtectionContainerMapping] depends_on.
func (srpcm *SiteRecoveryProtectionContainerMapping) Dependencies() terra.Dependencies {
	return srpcm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) LifecycleManagement() *terra.Lifecycle {
	return srpcm.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryProtectionContainerMapping].
func (srpcm *SiteRecoveryProtectionContainerMapping) Attributes() siteRecoveryProtectionContainerMappingAttributes {
	return siteRecoveryProtectionContainerMappingAttributes{ref: terra.ReferenceResource(srpcm)}
}

// ImportState imports the given attribute values into [SiteRecoveryProtectionContainerMapping]'s state.
func (srpcm *SiteRecoveryProtectionContainerMapping) ImportState(av io.Reader) error {
	srpcm.state = &siteRecoveryProtectionContainerMappingState{}
	if err := json.NewDecoder(av).Decode(srpcm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srpcm.Type(), srpcm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryProtectionContainerMapping] has state.
func (srpcm *SiteRecoveryProtectionContainerMapping) State() (*siteRecoveryProtectionContainerMappingState, bool) {
	return srpcm.state, srpcm.state != nil
}

// StateMust returns the state for [SiteRecoveryProtectionContainerMapping]. Panics if the state is nil.
func (srpcm *SiteRecoveryProtectionContainerMapping) StateMust() *siteRecoveryProtectionContainerMappingState {
	if srpcm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srpcm.Type(), srpcm.LocalName()))
	}
	return srpcm.state
}

// SiteRecoveryProtectionContainerMappingArgs contains the configurations for azurerm_site_recovery_protection_container_mapping.
type SiteRecoveryProtectionContainerMappingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryFabricName: string, required
	RecoveryFabricName terra.StringValue `hcl:"recovery_fabric_name,attr" validate:"required"`
	// RecoveryReplicationPolicyId: string, required
	RecoveryReplicationPolicyId terra.StringValue `hcl:"recovery_replication_policy_id,attr" validate:"required"`
	// RecoverySourceProtectionContainerName: string, required
	RecoverySourceProtectionContainerName terra.StringValue `hcl:"recovery_source_protection_container_name,attr" validate:"required"`
	// RecoveryTargetProtectionContainerId: string, required
	RecoveryTargetProtectionContainerId terra.StringValue `hcl:"recovery_target_protection_container_id,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// AutomaticUpdate: optional
	AutomaticUpdate *siterecoveryprotectioncontainermapping.AutomaticUpdate `hcl:"automatic_update,block"`
	// Timeouts: optional
	Timeouts *siterecoveryprotectioncontainermapping.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryProtectionContainerMappingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("name"))
}

// RecoveryFabricName returns a reference to field recovery_fabric_name of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) RecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("recovery_fabric_name"))
}

// RecoveryReplicationPolicyId returns a reference to field recovery_replication_policy_id of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) RecoveryReplicationPolicyId() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("recovery_replication_policy_id"))
}

// RecoverySourceProtectionContainerName returns a reference to field recovery_source_protection_container_name of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) RecoverySourceProtectionContainerName() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("recovery_source_protection_container_name"))
}

// RecoveryTargetProtectionContainerId returns a reference to field recovery_target_protection_container_id of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) RecoveryTargetProtectionContainerId() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("recovery_target_protection_container_id"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_protection_container_mapping.
func (srpcm siteRecoveryProtectionContainerMappingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srpcm.ref.Append("resource_group_name"))
}

func (srpcm siteRecoveryProtectionContainerMappingAttributes) AutomaticUpdate() terra.ListValue[siterecoveryprotectioncontainermapping.AutomaticUpdateAttributes] {
	return terra.ReferenceAsList[siterecoveryprotectioncontainermapping.AutomaticUpdateAttributes](srpcm.ref.Append("automatic_update"))
}

func (srpcm siteRecoveryProtectionContainerMappingAttributes) Timeouts() siterecoveryprotectioncontainermapping.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryprotectioncontainermapping.TimeoutsAttributes](srpcm.ref.Append("timeouts"))
}

type siteRecoveryProtectionContainerMappingState struct {
	Id                                    string                                                        `json:"id"`
	Name                                  string                                                        `json:"name"`
	RecoveryFabricName                    string                                                        `json:"recovery_fabric_name"`
	RecoveryReplicationPolicyId           string                                                        `json:"recovery_replication_policy_id"`
	RecoverySourceProtectionContainerName string                                                        `json:"recovery_source_protection_container_name"`
	RecoveryTargetProtectionContainerId   string                                                        `json:"recovery_target_protection_container_id"`
	RecoveryVaultName                     string                                                        `json:"recovery_vault_name"`
	ResourceGroupName                     string                                                        `json:"resource_group_name"`
	AutomaticUpdate                       []siterecoveryprotectioncontainermapping.AutomaticUpdateState `json:"automatic_update"`
	Timeouts                              *siterecoveryprotectioncontainermapping.TimeoutsState         `json:"timeouts"`
}