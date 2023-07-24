// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryreplicationrecoveryplan "github.com/golingon/terraproviders/azurerm/3.66.0/siterecoveryreplicationrecoveryplan"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryReplicationRecoveryPlan creates a new instance of [SiteRecoveryReplicationRecoveryPlan].
func NewSiteRecoveryReplicationRecoveryPlan(name string, args SiteRecoveryReplicationRecoveryPlanArgs) *SiteRecoveryReplicationRecoveryPlan {
	return &SiteRecoveryReplicationRecoveryPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryReplicationRecoveryPlan)(nil)

// SiteRecoveryReplicationRecoveryPlan represents the Terraform resource azurerm_site_recovery_replication_recovery_plan.
type SiteRecoveryReplicationRecoveryPlan struct {
	Name      string
	Args      SiteRecoveryReplicationRecoveryPlanArgs
	state     *siteRecoveryReplicationRecoveryPlanState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) Type() string {
	return "azurerm_site_recovery_replication_recovery_plan"
}

// LocalName returns the local name for [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) LocalName() string {
	return srrrp.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) Configuration() interface{} {
	return srrrp.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) DependOn() terra.Reference {
	return terra.ReferenceResource(srrrp)
}

// Dependencies returns the list of resources [SiteRecoveryReplicationRecoveryPlan] depends_on.
func (srrrp *SiteRecoveryReplicationRecoveryPlan) Dependencies() terra.Dependencies {
	return srrrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) LifecycleManagement() *terra.Lifecycle {
	return srrrp.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryReplicationRecoveryPlan].
func (srrrp *SiteRecoveryReplicationRecoveryPlan) Attributes() siteRecoveryReplicationRecoveryPlanAttributes {
	return siteRecoveryReplicationRecoveryPlanAttributes{ref: terra.ReferenceResource(srrrp)}
}

// ImportState imports the given attribute values into [SiteRecoveryReplicationRecoveryPlan]'s state.
func (srrrp *SiteRecoveryReplicationRecoveryPlan) ImportState(av io.Reader) error {
	srrrp.state = &siteRecoveryReplicationRecoveryPlanState{}
	if err := json.NewDecoder(av).Decode(srrrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srrrp.Type(), srrrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryReplicationRecoveryPlan] has state.
func (srrrp *SiteRecoveryReplicationRecoveryPlan) State() (*siteRecoveryReplicationRecoveryPlanState, bool) {
	return srrrp.state, srrrp.state != nil
}

// StateMust returns the state for [SiteRecoveryReplicationRecoveryPlan]. Panics if the state is nil.
func (srrrp *SiteRecoveryReplicationRecoveryPlan) StateMust() *siteRecoveryReplicationRecoveryPlanState {
	if srrrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srrrp.Type(), srrrp.LocalName()))
	}
	return srrrp.state
}

// SiteRecoveryReplicationRecoveryPlanArgs contains the configurations for azurerm_site_recovery_replication_recovery_plan.
type SiteRecoveryReplicationRecoveryPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// SourceRecoveryFabricId: string, required
	SourceRecoveryFabricId terra.StringValue `hcl:"source_recovery_fabric_id,attr" validate:"required"`
	// TargetRecoveryFabricId: string, required
	TargetRecoveryFabricId terra.StringValue `hcl:"target_recovery_fabric_id,attr" validate:"required"`
	// AzureToAzureSettings: optional
	AzureToAzureSettings *siterecoveryreplicationrecoveryplan.AzureToAzureSettings `hcl:"azure_to_azure_settings,block"`
	// BootRecoveryGroup: min=0
	BootRecoveryGroup []siterecoveryreplicationrecoveryplan.BootRecoveryGroup `hcl:"boot_recovery_group,block" validate:"min=0"`
	// FailoverRecoveryGroup: optional
	FailoverRecoveryGroup *siterecoveryreplicationrecoveryplan.FailoverRecoveryGroup `hcl:"failover_recovery_group,block"`
	// RecoveryGroup: min=0
	RecoveryGroup []siterecoveryreplicationrecoveryplan.RecoveryGroup `hcl:"recovery_group,block" validate:"min=0"`
	// ShutdownRecoveryGroup: optional
	ShutdownRecoveryGroup *siterecoveryreplicationrecoveryplan.ShutdownRecoveryGroup `hcl:"shutdown_recovery_group,block"`
	// Timeouts: optional
	Timeouts *siterecoveryreplicationrecoveryplan.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryReplicationRecoveryPlanAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_replication_recovery_plan.
func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("name"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("recovery_vault_id"))
}

// SourceRecoveryFabricId returns a reference to field source_recovery_fabric_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) SourceRecoveryFabricId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("source_recovery_fabric_id"))
}

// TargetRecoveryFabricId returns a reference to field target_recovery_fabric_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) TargetRecoveryFabricId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("target_recovery_fabric_id"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) AzureToAzureSettings() terra.ListValue[siterecoveryreplicationrecoveryplan.AzureToAzureSettingsAttributes] {
	return terra.ReferenceAsList[siterecoveryreplicationrecoveryplan.AzureToAzureSettingsAttributes](srrrp.ref.Append("azure_to_azure_settings"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) BootRecoveryGroup() terra.ListValue[siterecoveryreplicationrecoveryplan.BootRecoveryGroupAttributes] {
	return terra.ReferenceAsList[siterecoveryreplicationrecoveryplan.BootRecoveryGroupAttributes](srrrp.ref.Append("boot_recovery_group"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) FailoverRecoveryGroup() terra.ListValue[siterecoveryreplicationrecoveryplan.FailoverRecoveryGroupAttributes] {
	return terra.ReferenceAsList[siterecoveryreplicationrecoveryplan.FailoverRecoveryGroupAttributes](srrrp.ref.Append("failover_recovery_group"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) RecoveryGroup() terra.SetValue[siterecoveryreplicationrecoveryplan.RecoveryGroupAttributes] {
	return terra.ReferenceAsSet[siterecoveryreplicationrecoveryplan.RecoveryGroupAttributes](srrrp.ref.Append("recovery_group"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) ShutdownRecoveryGroup() terra.ListValue[siterecoveryreplicationrecoveryplan.ShutdownRecoveryGroupAttributes] {
	return terra.ReferenceAsList[siterecoveryreplicationrecoveryplan.ShutdownRecoveryGroupAttributes](srrrp.ref.Append("shutdown_recovery_group"))
}

func (srrrp siteRecoveryReplicationRecoveryPlanAttributes) Timeouts() siterecoveryreplicationrecoveryplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryreplicationrecoveryplan.TimeoutsAttributes](srrrp.ref.Append("timeouts"))
}

type siteRecoveryReplicationRecoveryPlanState struct {
	Id                     string                                                           `json:"id"`
	Name                   string                                                           `json:"name"`
	RecoveryVaultId        string                                                           `json:"recovery_vault_id"`
	SourceRecoveryFabricId string                                                           `json:"source_recovery_fabric_id"`
	TargetRecoveryFabricId string                                                           `json:"target_recovery_fabric_id"`
	AzureToAzureSettings   []siterecoveryreplicationrecoveryplan.AzureToAzureSettingsState  `json:"azure_to_azure_settings"`
	BootRecoveryGroup      []siterecoveryreplicationrecoveryplan.BootRecoveryGroupState     `json:"boot_recovery_group"`
	FailoverRecoveryGroup  []siterecoveryreplicationrecoveryplan.FailoverRecoveryGroupState `json:"failover_recovery_group"`
	RecoveryGroup          []siterecoveryreplicationrecoveryplan.RecoveryGroupState         `json:"recovery_group"`
	ShutdownRecoveryGroup  []siterecoveryreplicationrecoveryplan.ShutdownRecoveryGroupState `json:"shutdown_recovery_group"`
	Timeouts               *siterecoveryreplicationrecoveryplan.TimeoutsState               `json:"timeouts"`
}
