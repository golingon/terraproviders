// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasiterecoveryreplicationpolicy "github.com/golingon/terraproviders/azurerm/3.49.0/datasiterecoveryreplicationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSiteRecoveryReplicationPolicy creates a new instance of [DataSiteRecoveryReplicationPolicy].
func NewDataSiteRecoveryReplicationPolicy(name string, args DataSiteRecoveryReplicationPolicyArgs) *DataSiteRecoveryReplicationPolicy {
	return &DataSiteRecoveryReplicationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSiteRecoveryReplicationPolicy)(nil)

// DataSiteRecoveryReplicationPolicy represents the Terraform data resource azurerm_site_recovery_replication_policy.
type DataSiteRecoveryReplicationPolicy struct {
	Name string
	Args DataSiteRecoveryReplicationPolicyArgs
}

// DataSource returns the Terraform object type for [DataSiteRecoveryReplicationPolicy].
func (srrp *DataSiteRecoveryReplicationPolicy) DataSource() string {
	return "azurerm_site_recovery_replication_policy"
}

// LocalName returns the local name for [DataSiteRecoveryReplicationPolicy].
func (srrp *DataSiteRecoveryReplicationPolicy) LocalName() string {
	return srrp.Name
}

// Configuration returns the configuration (args) for [DataSiteRecoveryReplicationPolicy].
func (srrp *DataSiteRecoveryReplicationPolicy) Configuration() interface{} {
	return srrp.Args
}

// Attributes returns the attributes for [DataSiteRecoveryReplicationPolicy].
func (srrp *DataSiteRecoveryReplicationPolicy) Attributes() dataSiteRecoveryReplicationPolicyAttributes {
	return dataSiteRecoveryReplicationPolicyAttributes{ref: terra.ReferenceDataResource(srrp)}
}

// DataSiteRecoveryReplicationPolicyArgs contains the configurations for azurerm_site_recovery_replication_policy.
type DataSiteRecoveryReplicationPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasiterecoveryreplicationpolicy.Timeouts `hcl:"timeouts,block"`
}
type dataSiteRecoveryReplicationPolicyAttributes struct {
	ref terra.Reference
}

// ApplicationConsistentSnapshotFrequencyInMinutes returns a reference to field application_consistent_snapshot_frequency_in_minutes of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) ApplicationConsistentSnapshotFrequencyInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srrp.ref.Append("application_consistent_snapshot_frequency_in_minutes"))
}

// Id returns a reference to field id of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("name"))
}

// RecoveryPointRetentionInMinutes returns a reference to field recovery_point_retention_in_minutes of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) RecoveryPointRetentionInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srrp.ref.Append("recovery_point_retention_in_minutes"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_replication_policy.
func (srrp dataSiteRecoveryReplicationPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("resource_group_name"))
}

func (srrp dataSiteRecoveryReplicationPolicyAttributes) Timeouts() datasiterecoveryreplicationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasiterecoveryreplicationpolicy.TimeoutsAttributes](srrp.ref.Append("timeouts"))
}
