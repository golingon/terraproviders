// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryhypervnetworkmapping "github.com/golingon/terraproviders/azurerm/3.58.0/siterecoveryhypervnetworkmapping"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryHypervNetworkMapping creates a new instance of [SiteRecoveryHypervNetworkMapping].
func NewSiteRecoveryHypervNetworkMapping(name string, args SiteRecoveryHypervNetworkMappingArgs) *SiteRecoveryHypervNetworkMapping {
	return &SiteRecoveryHypervNetworkMapping{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryHypervNetworkMapping)(nil)

// SiteRecoveryHypervNetworkMapping represents the Terraform resource azurerm_site_recovery_hyperv_network_mapping.
type SiteRecoveryHypervNetworkMapping struct {
	Name      string
	Args      SiteRecoveryHypervNetworkMappingArgs
	state     *siteRecoveryHypervNetworkMappingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) Type() string {
	return "azurerm_site_recovery_hyperv_network_mapping"
}

// LocalName returns the local name for [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) LocalName() string {
	return srhnm.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) Configuration() interface{} {
	return srhnm.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) DependOn() terra.Reference {
	return terra.ReferenceResource(srhnm)
}

// Dependencies returns the list of resources [SiteRecoveryHypervNetworkMapping] depends_on.
func (srhnm *SiteRecoveryHypervNetworkMapping) Dependencies() terra.Dependencies {
	return srhnm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) LifecycleManagement() *terra.Lifecycle {
	return srhnm.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryHypervNetworkMapping].
func (srhnm *SiteRecoveryHypervNetworkMapping) Attributes() siteRecoveryHypervNetworkMappingAttributes {
	return siteRecoveryHypervNetworkMappingAttributes{ref: terra.ReferenceResource(srhnm)}
}

// ImportState imports the given attribute values into [SiteRecoveryHypervNetworkMapping]'s state.
func (srhnm *SiteRecoveryHypervNetworkMapping) ImportState(av io.Reader) error {
	srhnm.state = &siteRecoveryHypervNetworkMappingState{}
	if err := json.NewDecoder(av).Decode(srhnm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srhnm.Type(), srhnm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryHypervNetworkMapping] has state.
func (srhnm *SiteRecoveryHypervNetworkMapping) State() (*siteRecoveryHypervNetworkMappingState, bool) {
	return srhnm.state, srhnm.state != nil
}

// StateMust returns the state for [SiteRecoveryHypervNetworkMapping]. Panics if the state is nil.
func (srhnm *SiteRecoveryHypervNetworkMapping) StateMust() *siteRecoveryHypervNetworkMappingState {
	if srhnm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srhnm.Type(), srhnm.LocalName()))
	}
	return srhnm.state
}

// SiteRecoveryHypervNetworkMappingArgs contains the configurations for azurerm_site_recovery_hyperv_network_mapping.
type SiteRecoveryHypervNetworkMappingArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// SourceNetworkName: string, required
	SourceNetworkName terra.StringValue `hcl:"source_network_name,attr" validate:"required"`
	// SourceSystemCenterVirtualMachineManagerName: string, required
	SourceSystemCenterVirtualMachineManagerName terra.StringValue `hcl:"source_system_center_virtual_machine_manager_name,attr" validate:"required"`
	// TargetNetworkId: string, required
	TargetNetworkId terra.StringValue `hcl:"target_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryhypervnetworkmapping.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryHypervNetworkMappingAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("name"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("recovery_vault_id"))
}

// SourceNetworkName returns a reference to field source_network_name of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) SourceNetworkName() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("source_network_name"))
}

// SourceSystemCenterVirtualMachineManagerName returns a reference to field source_system_center_virtual_machine_manager_name of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) SourceSystemCenterVirtualMachineManagerName() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("source_system_center_virtual_machine_manager_name"))
}

// TargetNetworkId returns a reference to field target_network_id of azurerm_site_recovery_hyperv_network_mapping.
func (srhnm siteRecoveryHypervNetworkMappingAttributes) TargetNetworkId() terra.StringValue {
	return terra.ReferenceAsString(srhnm.ref.Append("target_network_id"))
}

func (srhnm siteRecoveryHypervNetworkMappingAttributes) Timeouts() siterecoveryhypervnetworkmapping.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryhypervnetworkmapping.TimeoutsAttributes](srhnm.ref.Append("timeouts"))
}

type siteRecoveryHypervNetworkMappingState struct {
	Id                                          string                                          `json:"id"`
	Name                                        string                                          `json:"name"`
	RecoveryVaultId                             string                                          `json:"recovery_vault_id"`
	SourceNetworkName                           string                                          `json:"source_network_name"`
	SourceSystemCenterVirtualMachineManagerName string                                          `json:"source_system_center_virtual_machine_manager_name"`
	TargetNetworkId                             string                                          `json:"target_network_id"`
	Timeouts                                    *siterecoveryhypervnetworkmapping.TimeoutsState `json:"timeouts"`
}
