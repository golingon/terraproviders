// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoverynetworkmapping "github.com/golingon/terraproviders/azurerm/3.66.0/siterecoverynetworkmapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryNetworkMapping creates a new instance of [SiteRecoveryNetworkMapping].
func NewSiteRecoveryNetworkMapping(name string, args SiteRecoveryNetworkMappingArgs) *SiteRecoveryNetworkMapping {
	return &SiteRecoveryNetworkMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryNetworkMapping)(nil)

// SiteRecoveryNetworkMapping represents the Terraform resource azurerm_site_recovery_network_mapping.
type SiteRecoveryNetworkMapping struct {
	Name      string
	Args      SiteRecoveryNetworkMappingArgs
	state     *siteRecoveryNetworkMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) Type() string {
	return "azurerm_site_recovery_network_mapping"
}

// LocalName returns the local name for [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) LocalName() string {
	return srnm.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) Configuration() interface{} {
	return srnm.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(srnm)
}

// Dependencies returns the list of resources [SiteRecoveryNetworkMapping] depends_on.
func (srnm *SiteRecoveryNetworkMapping) Dependencies() terra.Dependencies {
	return srnm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) LifecycleManagement() *terra.Lifecycle {
	return srnm.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryNetworkMapping].
func (srnm *SiteRecoveryNetworkMapping) Attributes() siteRecoveryNetworkMappingAttributes {
	return siteRecoveryNetworkMappingAttributes{ref: terra.ReferenceResource(srnm)}
}

// ImportState imports the given attribute values into [SiteRecoveryNetworkMapping]'s state.
func (srnm *SiteRecoveryNetworkMapping) ImportState(av io.Reader) error {
	srnm.state = &siteRecoveryNetworkMappingState{}
	if err := json.NewDecoder(av).Decode(srnm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srnm.Type(), srnm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryNetworkMapping] has state.
func (srnm *SiteRecoveryNetworkMapping) State() (*siteRecoveryNetworkMappingState, bool) {
	return srnm.state, srnm.state != nil
}

// StateMust returns the state for [SiteRecoveryNetworkMapping]. Panics if the state is nil.
func (srnm *SiteRecoveryNetworkMapping) StateMust() *siteRecoveryNetworkMappingState {
	if srnm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srnm.Type(), srnm.LocalName()))
	}
	return srnm.state
}

// SiteRecoveryNetworkMappingArgs contains the configurations for azurerm_site_recovery_network_mapping.
type SiteRecoveryNetworkMappingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SourceNetworkId: string, required
	SourceNetworkId terra.StringValue `hcl:"source_network_id,attr" validate:"required"`
	// SourceRecoveryFabricName: string, required
	SourceRecoveryFabricName terra.StringValue `hcl:"source_recovery_fabric_name,attr" validate:"required"`
	// TargetNetworkId: string, required
	TargetNetworkId terra.StringValue `hcl:"target_network_id,attr" validate:"required"`
	// TargetRecoveryFabricName: string, required
	TargetRecoveryFabricName terra.StringValue `hcl:"target_recovery_fabric_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoverynetworkmapping.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryNetworkMappingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("resource_group_name"))
}

// SourceNetworkId returns a reference to field source_network_id of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) SourceNetworkId() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("source_network_id"))
}

// SourceRecoveryFabricName returns a reference to field source_recovery_fabric_name of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) SourceRecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("source_recovery_fabric_name"))
}

// TargetNetworkId returns a reference to field target_network_id of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) TargetNetworkId() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("target_network_id"))
}

// TargetRecoveryFabricName returns a reference to field target_recovery_fabric_name of azurerm_site_recovery_network_mapping.
func (srnm siteRecoveryNetworkMappingAttributes) TargetRecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srnm.ref.Append("target_recovery_fabric_name"))
}

func (srnm siteRecoveryNetworkMappingAttributes) Timeouts() siterecoverynetworkmapping.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoverynetworkmapping.TimeoutsAttributes](srnm.ref.Append("timeouts"))
}

type siteRecoveryNetworkMappingState struct {
	Id                       string                                    `json:"id"`
	Name                     string                                    `json:"name"`
	RecoveryVaultName        string                                    `json:"recovery_vault_name"`
	ResourceGroupName        string                                    `json:"resource_group_name"`
	SourceNetworkId          string                                    `json:"source_network_id"`
	SourceRecoveryFabricName string                                    `json:"source_recovery_fabric_name"`
	TargetNetworkId          string                                    `json:"target_network_id"`
	TargetRecoveryFabricName string                                    `json:"target_recovery_fabric_name"`
	Timeouts                 *siterecoverynetworkmapping.TimeoutsState `json:"timeouts"`
}
