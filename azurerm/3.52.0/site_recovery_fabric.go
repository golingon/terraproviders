// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryfabric "github.com/golingon/terraproviders/azurerm/3.52.0/siterecoveryfabric"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryFabric creates a new instance of [SiteRecoveryFabric].
func NewSiteRecoveryFabric(name string, args SiteRecoveryFabricArgs) *SiteRecoveryFabric {
	return &SiteRecoveryFabric{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryFabric)(nil)

// SiteRecoveryFabric represents the Terraform resource azurerm_site_recovery_fabric.
type SiteRecoveryFabric struct {
	Name      string
	Args      SiteRecoveryFabricArgs
	state     *siteRecoveryFabricState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) Type() string {
	return "azurerm_site_recovery_fabric"
}

// LocalName returns the local name for [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) LocalName() string {
	return srf.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) Configuration() interface{} {
	return srf.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) DependOn() terra.Reference {
	return terra.ReferenceResource(srf)
}

// Dependencies returns the list of resources [SiteRecoveryFabric] depends_on.
func (srf *SiteRecoveryFabric) Dependencies() terra.Dependencies {
	return srf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) LifecycleManagement() *terra.Lifecycle {
	return srf.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryFabric].
func (srf *SiteRecoveryFabric) Attributes() siteRecoveryFabricAttributes {
	return siteRecoveryFabricAttributes{ref: terra.ReferenceResource(srf)}
}

// ImportState imports the given attribute values into [SiteRecoveryFabric]'s state.
func (srf *SiteRecoveryFabric) ImportState(av io.Reader) error {
	srf.state = &siteRecoveryFabricState{}
	if err := json.NewDecoder(av).Decode(srf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srf.Type(), srf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryFabric] has state.
func (srf *SiteRecoveryFabric) State() (*siteRecoveryFabricState, bool) {
	return srf.state, srf.state != nil
}

// StateMust returns the state for [SiteRecoveryFabric]. Panics if the state is nil.
func (srf *SiteRecoveryFabric) StateMust() *siteRecoveryFabricState {
	if srf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srf.Type(), srf.LocalName()))
	}
	return srf.state
}

// SiteRecoveryFabricArgs contains the configurations for azurerm_site_recovery_fabric.
type SiteRecoveryFabricArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryfabric.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryFabricAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_fabric.
func (srf siteRecoveryFabricAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_site_recovery_fabric.
func (srf siteRecoveryFabricAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_site_recovery_fabric.
func (srf siteRecoveryFabricAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_fabric.
func (srf siteRecoveryFabricAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_fabric.
func (srf siteRecoveryFabricAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("resource_group_name"))
}

func (srf siteRecoveryFabricAttributes) Timeouts() siterecoveryfabric.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryfabric.TimeoutsAttributes](srf.ref.Append("timeouts"))
}

type siteRecoveryFabricState struct {
	Id                string                            `json:"id"`
	Location          string                            `json:"location"`
	Name              string                            `json:"name"`
	RecoveryVaultName string                            `json:"recovery_vault_name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	Timeouts          *siterecoveryfabric.TimeoutsState `json:"timeouts"`
}
