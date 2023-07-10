// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryservicesvaulthypervsite "github.com/golingon/terraproviders/azurerm/3.64.0/siterecoveryservicesvaulthypervsite"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryServicesVaultHypervSite creates a new instance of [SiteRecoveryServicesVaultHypervSite].
func NewSiteRecoveryServicesVaultHypervSite(name string, args SiteRecoveryServicesVaultHypervSiteArgs) *SiteRecoveryServicesVaultHypervSite {
	return &SiteRecoveryServicesVaultHypervSite{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryServicesVaultHypervSite)(nil)

// SiteRecoveryServicesVaultHypervSite represents the Terraform resource azurerm_site_recovery_services_vault_hyperv_site.
type SiteRecoveryServicesVaultHypervSite struct {
	Name      string
	Args      SiteRecoveryServicesVaultHypervSiteArgs
	state     *siteRecoveryServicesVaultHypervSiteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) Type() string {
	return "azurerm_site_recovery_services_vault_hyperv_site"
}

// LocalName returns the local name for [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) LocalName() string {
	return srsvhs.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) Configuration() interface{} {
	return srsvhs.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) DependOn() terra.Reference {
	return terra.ReferenceResource(srsvhs)
}

// Dependencies returns the list of resources [SiteRecoveryServicesVaultHypervSite] depends_on.
func (srsvhs *SiteRecoveryServicesVaultHypervSite) Dependencies() terra.Dependencies {
	return srsvhs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) LifecycleManagement() *terra.Lifecycle {
	return srsvhs.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryServicesVaultHypervSite].
func (srsvhs *SiteRecoveryServicesVaultHypervSite) Attributes() siteRecoveryServicesVaultHypervSiteAttributes {
	return siteRecoveryServicesVaultHypervSiteAttributes{ref: terra.ReferenceResource(srsvhs)}
}

// ImportState imports the given attribute values into [SiteRecoveryServicesVaultHypervSite]'s state.
func (srsvhs *SiteRecoveryServicesVaultHypervSite) ImportState(av io.Reader) error {
	srsvhs.state = &siteRecoveryServicesVaultHypervSiteState{}
	if err := json.NewDecoder(av).Decode(srsvhs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srsvhs.Type(), srsvhs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryServicesVaultHypervSite] has state.
func (srsvhs *SiteRecoveryServicesVaultHypervSite) State() (*siteRecoveryServicesVaultHypervSiteState, bool) {
	return srsvhs.state, srsvhs.state != nil
}

// StateMust returns the state for [SiteRecoveryServicesVaultHypervSite]. Panics if the state is nil.
func (srsvhs *SiteRecoveryServicesVaultHypervSite) StateMust() *siteRecoveryServicesVaultHypervSiteState {
	if srsvhs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srsvhs.Type(), srsvhs.LocalName()))
	}
	return srsvhs.state
}

// SiteRecoveryServicesVaultHypervSiteArgs contains the configurations for azurerm_site_recovery_services_vault_hyperv_site.
type SiteRecoveryServicesVaultHypervSiteArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryservicesvaulthypervsite.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryServicesVaultHypervSiteAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_services_vault_hyperv_site.
func (srsvhs siteRecoveryServicesVaultHypervSiteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srsvhs.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_services_vault_hyperv_site.
func (srsvhs siteRecoveryServicesVaultHypervSiteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srsvhs.ref.Append("name"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_services_vault_hyperv_site.
func (srsvhs siteRecoveryServicesVaultHypervSiteAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srsvhs.ref.Append("recovery_vault_id"))
}

func (srsvhs siteRecoveryServicesVaultHypervSiteAttributes) Timeouts() siterecoveryservicesvaulthypervsite.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryservicesvaulthypervsite.TimeoutsAttributes](srsvhs.ref.Append("timeouts"))
}

type siteRecoveryServicesVaultHypervSiteState struct {
	Id              string                                             `json:"id"`
	Name            string                                             `json:"name"`
	RecoveryVaultId string                                             `json:"recovery_vault_id"`
	Timeouts        *siterecoveryservicesvaulthypervsite.TimeoutsState `json:"timeouts"`
}
