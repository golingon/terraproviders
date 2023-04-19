// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryreplicationpolicy "github.com/golingon/terraproviders/azurerm/3.52.0/siterecoveryreplicationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryReplicationPolicy creates a new instance of [SiteRecoveryReplicationPolicy].
func NewSiteRecoveryReplicationPolicy(name string, args SiteRecoveryReplicationPolicyArgs) *SiteRecoveryReplicationPolicy {
	return &SiteRecoveryReplicationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryReplicationPolicy)(nil)

// SiteRecoveryReplicationPolicy represents the Terraform resource azurerm_site_recovery_replication_policy.
type SiteRecoveryReplicationPolicy struct {
	Name      string
	Args      SiteRecoveryReplicationPolicyArgs
	state     *siteRecoveryReplicationPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) Type() string {
	return "azurerm_site_recovery_replication_policy"
}

// LocalName returns the local name for [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) LocalName() string {
	return srrp.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) Configuration() interface{} {
	return srrp.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(srrp)
}

// Dependencies returns the list of resources [SiteRecoveryReplicationPolicy] depends_on.
func (srrp *SiteRecoveryReplicationPolicy) Dependencies() terra.Dependencies {
	return srrp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) LifecycleManagement() *terra.Lifecycle {
	return srrp.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryReplicationPolicy].
func (srrp *SiteRecoveryReplicationPolicy) Attributes() siteRecoveryReplicationPolicyAttributes {
	return siteRecoveryReplicationPolicyAttributes{ref: terra.ReferenceResource(srrp)}
}

// ImportState imports the given attribute values into [SiteRecoveryReplicationPolicy]'s state.
func (srrp *SiteRecoveryReplicationPolicy) ImportState(av io.Reader) error {
	srrp.state = &siteRecoveryReplicationPolicyState{}
	if err := json.NewDecoder(av).Decode(srrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srrp.Type(), srrp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryReplicationPolicy] has state.
func (srrp *SiteRecoveryReplicationPolicy) State() (*siteRecoveryReplicationPolicyState, bool) {
	return srrp.state, srrp.state != nil
}

// StateMust returns the state for [SiteRecoveryReplicationPolicy]. Panics if the state is nil.
func (srrp *SiteRecoveryReplicationPolicy) StateMust() *siteRecoveryReplicationPolicyState {
	if srrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srrp.Type(), srrp.LocalName()))
	}
	return srrp.state
}

// SiteRecoveryReplicationPolicyArgs contains the configurations for azurerm_site_recovery_replication_policy.
type SiteRecoveryReplicationPolicyArgs struct {
	// ApplicationConsistentSnapshotFrequencyInMinutes: number, required
	ApplicationConsistentSnapshotFrequencyInMinutes terra.NumberValue `hcl:"application_consistent_snapshot_frequency_in_minutes,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryPointRetentionInMinutes: number, required
	RecoveryPointRetentionInMinutes terra.NumberValue `hcl:"recovery_point_retention_in_minutes,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryreplicationpolicy.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryReplicationPolicyAttributes struct {
	ref terra.Reference
}

// ApplicationConsistentSnapshotFrequencyInMinutes returns a reference to field application_consistent_snapshot_frequency_in_minutes of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) ApplicationConsistentSnapshotFrequencyInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srrp.ref.Append("application_consistent_snapshot_frequency_in_minutes"))
}

// Id returns a reference to field id of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("name"))
}

// RecoveryPointRetentionInMinutes returns a reference to field recovery_point_retention_in_minutes of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) RecoveryPointRetentionInMinutes() terra.NumberValue {
	return terra.ReferenceAsNumber(srrp.ref.Append("recovery_point_retention_in_minutes"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_replication_policy.
func (srrp siteRecoveryReplicationPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srrp.ref.Append("resource_group_name"))
}

func (srrp siteRecoveryReplicationPolicyAttributes) Timeouts() siterecoveryreplicationpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryreplicationpolicy.TimeoutsAttributes](srrp.ref.Append("timeouts"))
}

type siteRecoveryReplicationPolicyState struct {
	ApplicationConsistentSnapshotFrequencyInMinutes float64                                      `json:"application_consistent_snapshot_frequency_in_minutes"`
	Id                                              string                                       `json:"id"`
	Name                                            string                                       `json:"name"`
	RecoveryPointRetentionInMinutes                 float64                                      `json:"recovery_point_retention_in_minutes"`
	RecoveryVaultName                               string                                       `json:"recovery_vault_name"`
	ResourceGroupName                               string                                       `json:"resource_group_name"`
	Timeouts                                        *siterecoveryreplicationpolicy.TimeoutsState `json:"timeouts"`
}
