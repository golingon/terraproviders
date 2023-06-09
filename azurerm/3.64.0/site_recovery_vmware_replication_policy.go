// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryvmwarereplicationpolicy "github.com/golingon/terraproviders/azurerm/3.64.0/siterecoveryvmwarereplicationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryVmwareReplicationPolicy creates a new instance of [SiteRecoveryVmwareReplicationPolicy].
func NewSiteRecoveryVmwareReplicationPolicy(name string, args SiteRecoveryVmwareReplicationPolicyArgs) *SiteRecoveryVmwareReplicationPolicy {
	return &SiteRecoveryVmwareReplicationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryVmwareReplicationPolicy)(nil)

// SiteRecoveryVmwareReplicationPolicy represents the Terraform resource azurerm_site_recovery_vmware_replication_policy.
type SiteRecoveryVmwareReplicationPolicy struct {
	Name      string
	Args      SiteRecoveryVmwareReplicationPolicyArgs
	state     *siteRecoveryVmwareReplicationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) Type() string {
	return "azurerm_site_recovery_vmware_replication_policy"
}

// LocalName returns the local name for [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) LocalName() string {
	return srvrp.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) Configuration() interface{} {
	return srvrp.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(srvrp)
}

// Dependencies returns the list of resources [SiteRecoveryVmwareReplicationPolicy] depends_on.
func (srvrp *SiteRecoveryVmwareReplicationPolicy) Dependencies() terra.Dependencies {
	return srvrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) LifecycleManagement() *terra.Lifecycle {
	return srvrp.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryVmwareReplicationPolicy].
func (srvrp *SiteRecoveryVmwareReplicationPolicy) Attributes() siteRecoveryVmwareReplicationPolicyAttributes {
	return siteRecoveryVmwareReplicationPolicyAttributes{ref: terra.ReferenceResource(srvrp)}
}

// ImportState imports the given attribute values into [SiteRecoveryVmwareReplicationPolicy]'s state.
func (srvrp *SiteRecoveryVmwareReplicationPolicy) ImportState(av io.Reader) error {
	srvrp.state = &siteRecoveryVmwareReplicationPolicyState{}
	if err := json.NewDecoder(av).Decode(srvrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srvrp.Type(), srvrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryVmwareReplicationPolicy] has state.
func (srvrp *SiteRecoveryVmwareReplicationPolicy) State() (*siteRecoveryVmwareReplicationPolicyState, bool) {
	return srvrp.state, srvrp.state != nil
}

// StateMust returns the state for [SiteRecoveryVmwareReplicationPolicy]. Panics if the state is nil.
func (srvrp *SiteRecoveryVmwareReplicationPolicy) StateMust() *siteRecoveryVmwareReplicationPolicyState {
	if srvrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srvrp.Type(), srvrp.LocalName()))
	}
	return srvrp.state
}

// SiteRecoveryVmwareReplicationPolicyArgs contains the configurations for azurerm_site_recovery_vmware_replication_policy.
type SiteRecoveryVmwareReplicationPolicyArgs struct {
	// ApplicationConsistentSnapshotFrequencyInMinutes: number, required
	ApplicationConsistentSnapshotFrequencyInMinutes terra.NumberValue `hcl:"application_consistent_snapshot_frequency_in_minutes,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryPointRetentionInMinutes: number, required
	RecoveryPointRetentionInMinutes terra.NumberValue `hcl:"recovery_point_retention_in_minutes,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryvmwarereplicationpolicy.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryVmwareReplicationPolicyAttributes struct {
	ref terra.Reference
}

// ApplicationConsistentSnapshotFrequencyInMinutes returns a reference to field application_consistent_snapshot_frequency_in_minutes of azurerm_site_recovery_vmware_replication_policy.
func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) ApplicationConsistentSnapshotFrequencyInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srvrp.ref.Append("application_consistent_snapshot_frequency_in_minutes"))
}

// Id returns a reference to field id of azurerm_site_recovery_vmware_replication_policy.
func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srvrp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_vmware_replication_policy.
func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srvrp.ref.Append("name"))
}

// RecoveryPointRetentionInMinutes returns a reference to field recovery_point_retention_in_minutes of azurerm_site_recovery_vmware_replication_policy.
func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) RecoveryPointRetentionInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srvrp.ref.Append("recovery_point_retention_in_minutes"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_vmware_replication_policy.
func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srvrp.ref.Append("recovery_vault_id"))
}

func (srvrp siteRecoveryVmwareReplicationPolicyAttributes) Timeouts() siterecoveryvmwarereplicationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryvmwarereplicationpolicy.TimeoutsAttributes](srvrp.ref.Append("timeouts"))
}

type siteRecoveryVmwareReplicationPolicyState struct {
	ApplicationConsistentSnapshotFrequencyInMinutes float64                                            `json:"application_consistent_snapshot_frequency_in_minutes"`
	Id                                              string                                             `json:"id"`
	Name                                            string                                             `json:"name"`
	RecoveryPointRetentionInMinutes                 float64                                            `json:"recovery_point_retention_in_minutes"`
	RecoveryVaultId                                 string                                             `json:"recovery_vault_id"`
	Timeouts                                        *siterecoveryvmwarereplicationpolicy.TimeoutsState `json:"timeouts"`
}
