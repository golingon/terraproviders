// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryprotectioncontainer "github.com/golingon/terraproviders/azurerm/3.49.0/siterecoveryprotectioncontainer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryProtectionContainer creates a new instance of [SiteRecoveryProtectionContainer].
func NewSiteRecoveryProtectionContainer(name string, args SiteRecoveryProtectionContainerArgs) *SiteRecoveryProtectionContainer {
	return &SiteRecoveryProtectionContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryProtectionContainer)(nil)

// SiteRecoveryProtectionContainer represents the Terraform resource azurerm_site_recovery_protection_container.
type SiteRecoveryProtectionContainer struct {
	Name      string
	Args      SiteRecoveryProtectionContainerArgs
	state     *siteRecoveryProtectionContainerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) Type() string {
	return "azurerm_site_recovery_protection_container"
}

// LocalName returns the local name for [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) LocalName() string {
	return srpc.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) Configuration() interface{} {
	return srpc.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) DependOn() terra.Reference {
	return terra.ReferenceResource(srpc)
}

// Dependencies returns the list of resources [SiteRecoveryProtectionContainer] depends_on.
func (srpc *SiteRecoveryProtectionContainer) Dependencies() terra.Dependencies {
	return srpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) LifecycleManagement() *terra.Lifecycle {
	return srpc.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryProtectionContainer].
func (srpc *SiteRecoveryProtectionContainer) Attributes() siteRecoveryProtectionContainerAttributes {
	return siteRecoveryProtectionContainerAttributes{ref: terra.ReferenceResource(srpc)}
}

// ImportState imports the given attribute values into [SiteRecoveryProtectionContainer]'s state.
func (srpc *SiteRecoveryProtectionContainer) ImportState(av io.Reader) error {
	srpc.state = &siteRecoveryProtectionContainerState{}
	if err := json.NewDecoder(av).Decode(srpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srpc.Type(), srpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryProtectionContainer] has state.
func (srpc *SiteRecoveryProtectionContainer) State() (*siteRecoveryProtectionContainerState, bool) {
	return srpc.state, srpc.state != nil
}

// StateMust returns the state for [SiteRecoveryProtectionContainer]. Panics if the state is nil.
func (srpc *SiteRecoveryProtectionContainer) StateMust() *siteRecoveryProtectionContainerState {
	if srpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srpc.Type(), srpc.LocalName()))
	}
	return srpc.state
}

// SiteRecoveryProtectionContainerArgs contains the configurations for azurerm_site_recovery_protection_container.
type SiteRecoveryProtectionContainerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryFabricName: string, required
	RecoveryFabricName terra.StringValue `hcl:"recovery_fabric_name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryprotectioncontainer.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryProtectionContainerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_protection_container.
func (srpc siteRecoveryProtectionContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_protection_container.
func (srpc siteRecoveryProtectionContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("name"))
}

// RecoveryFabricName returns a reference to field recovery_fabric_name of azurerm_site_recovery_protection_container.
func (srpc siteRecoveryProtectionContainerAttributes) RecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("recovery_fabric_name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_protection_container.
func (srpc siteRecoveryProtectionContainerAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_protection_container.
func (srpc siteRecoveryProtectionContainerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("resource_group_name"))
}

func (srpc siteRecoveryProtectionContainerAttributes) Timeouts() siterecoveryprotectioncontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryprotectioncontainer.TimeoutsAttributes](srpc.ref.Append("timeouts"))
}

type siteRecoveryProtectionContainerState struct {
	Id                 string                                         `json:"id"`
	Name               string                                         `json:"name"`
	RecoveryFabricName string                                         `json:"recovery_fabric_name"`
	RecoveryVaultName  string                                         `json:"recovery_vault_name"`
	ResourceGroupName  string                                         `json:"resource_group_name"`
	Timeouts           *siterecoveryprotectioncontainer.TimeoutsState `json:"timeouts"`
}
