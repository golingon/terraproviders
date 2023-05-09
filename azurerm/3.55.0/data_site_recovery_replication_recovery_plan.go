// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasiterecoveryreplicationrecoveryplan "github.com/golingon/terraproviders/azurerm/3.55.0/datasiterecoveryreplicationrecoveryplan"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSiteRecoveryReplicationRecoveryPlan creates a new instance of [DataSiteRecoveryReplicationRecoveryPlan].
func NewDataSiteRecoveryReplicationRecoveryPlan(name string, args DataSiteRecoveryReplicationRecoveryPlanArgs) *DataSiteRecoveryReplicationRecoveryPlan {
	return &DataSiteRecoveryReplicationRecoveryPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSiteRecoveryReplicationRecoveryPlan)(nil)

// DataSiteRecoveryReplicationRecoveryPlan represents the Terraform data resource azurerm_site_recovery_replication_recovery_plan.
type DataSiteRecoveryReplicationRecoveryPlan struct {
	Name string
	Args DataSiteRecoveryReplicationRecoveryPlanArgs
}

// DataSource returns the Terraform object type for [DataSiteRecoveryReplicationRecoveryPlan].
func (srrrp *DataSiteRecoveryReplicationRecoveryPlan) DataSource() string {
	return "azurerm_site_recovery_replication_recovery_plan"
}

// LocalName returns the local name for [DataSiteRecoveryReplicationRecoveryPlan].
func (srrrp *DataSiteRecoveryReplicationRecoveryPlan) LocalName() string {
	return srrrp.Name
}

// Configuration returns the configuration (args) for [DataSiteRecoveryReplicationRecoveryPlan].
func (srrrp *DataSiteRecoveryReplicationRecoveryPlan) Configuration() interface{} {
	return srrrp.Args
}

// Attributes returns the attributes for [DataSiteRecoveryReplicationRecoveryPlan].
func (srrrp *DataSiteRecoveryReplicationRecoveryPlan) Attributes() dataSiteRecoveryReplicationRecoveryPlanAttributes {
	return dataSiteRecoveryReplicationRecoveryPlanAttributes{ref: terra.ReferenceDataResource(srrrp)}
}

// DataSiteRecoveryReplicationRecoveryPlanArgs contains the configurations for azurerm_site_recovery_replication_recovery_plan.
type DataSiteRecoveryReplicationRecoveryPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// RecoveryGroup: min=0
	RecoveryGroup []datasiterecoveryreplicationrecoveryplan.RecoveryGroup `hcl:"recovery_group,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasiterecoveryreplicationrecoveryplan.Timeouts `hcl:"timeouts,block"`
}
type dataSiteRecoveryReplicationRecoveryPlanAttributes struct {
	ref terra.Reference
}

// FailoverDeploymentModel returns a reference to field failover_deployment_model of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) FailoverDeploymentModel() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("failover_deployment_model"))
}

// Id returns a reference to field id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("name"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("recovery_vault_id"))
}

// SourceRecoveryFabricId returns a reference to field source_recovery_fabric_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) SourceRecoveryFabricId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("source_recovery_fabric_id"))
}

// TargetRecoveryFabricId returns a reference to field target_recovery_fabric_id of azurerm_site_recovery_replication_recovery_plan.
func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) TargetRecoveryFabricId() terra.StringValue {
	return terra.ReferenceAsString(srrrp.ref.Append("target_recovery_fabric_id"))
}

func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) RecoveryGroup() terra.SetValue[datasiterecoveryreplicationrecoveryplan.RecoveryGroupAttributes] {
	return terra.ReferenceAsSet[datasiterecoveryreplicationrecoveryplan.RecoveryGroupAttributes](srrrp.ref.Append("recovery_group"))
}

func (srrrp dataSiteRecoveryReplicationRecoveryPlanAttributes) Timeouts() datasiterecoveryreplicationrecoveryplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasiterecoveryreplicationrecoveryplan.TimeoutsAttributes](srrrp.ref.Append("timeouts"))
}
