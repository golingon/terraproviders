// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryhypervreplicationpolicy "github.com/golingon/terraproviders/azurerm/3.49.0/siterecoveryhypervreplicationpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSiteRecoveryHypervReplicationPolicy(name string, args SiteRecoveryHypervReplicationPolicyArgs) *SiteRecoveryHypervReplicationPolicy {
	return &SiteRecoveryHypervReplicationPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryHypervReplicationPolicy)(nil)

type SiteRecoveryHypervReplicationPolicy struct {
	Name  string
	Args  SiteRecoveryHypervReplicationPolicyArgs
	state *siteRecoveryHypervReplicationPolicyState
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) Type() string {
	return "azurerm_site_recovery_hyperv_replication_policy"
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) LocalName() string {
	return srhrp.Name
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) Configuration() interface{} {
	return srhrp.Args
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) Attributes() siteRecoveryHypervReplicationPolicyAttributes {
	return siteRecoveryHypervReplicationPolicyAttributes{ref: terra.ReferenceResource(srhrp)}
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) ImportState(av io.Reader) error {
	srhrp.state = &siteRecoveryHypervReplicationPolicyState{}
	if err := json.NewDecoder(av).Decode(srhrp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srhrp.Type(), srhrp.LocalName(), err)
	}
	return nil
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) State() (*siteRecoveryHypervReplicationPolicyState, bool) {
	return srhrp.state, srhrp.state != nil
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) StateMust() *siteRecoveryHypervReplicationPolicyState {
	if srhrp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srhrp.Type(), srhrp.LocalName()))
	}
	return srhrp.state
}

func (srhrp *SiteRecoveryHypervReplicationPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(srhrp)
}

type SiteRecoveryHypervReplicationPolicyArgs struct {
	// ApplicationConsistentSnapshotFrequencyInHours: number, required
	ApplicationConsistentSnapshotFrequencyInHours terra.NumberValue `hcl:"application_consistent_snapshot_frequency_in_hours,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryPointRetentionInHours: number, required
	RecoveryPointRetentionInHours terra.NumberValue `hcl:"recovery_point_retention_in_hours,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// ReplicationIntervalInSeconds: number, required
	ReplicationIntervalInSeconds terra.NumberValue `hcl:"replication_interval_in_seconds,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryhypervreplicationpolicy.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that SiteRecoveryHypervReplicationPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type siteRecoveryHypervReplicationPolicyAttributes struct {
	ref terra.Reference
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) ApplicationConsistentSnapshotFrequencyInHours() terra.NumberValue {
	return terra.ReferenceNumber(srhrp.ref.Append("application_consistent_snapshot_frequency_in_hours"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(srhrp.ref.Append("id"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceString(srhrp.ref.Append("name"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) RecoveryPointRetentionInHours() terra.NumberValue {
	return terra.ReferenceNumber(srhrp.ref.Append("recovery_point_retention_in_hours"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceString(srhrp.ref.Append("recovery_vault_id"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) ReplicationIntervalInSeconds() terra.NumberValue {
	return terra.ReferenceNumber(srhrp.ref.Append("replication_interval_in_seconds"))
}

func (srhrp siteRecoveryHypervReplicationPolicyAttributes) Timeouts() siterecoveryhypervreplicationpolicy.TimeoutsAttributes {
	return terra.ReferenceSingle[siterecoveryhypervreplicationpolicy.TimeoutsAttributes](srhrp.ref.Append("timeouts"))
}

type siteRecoveryHypervReplicationPolicyState struct {
	ApplicationConsistentSnapshotFrequencyInHours float64                                            `json:"application_consistent_snapshot_frequency_in_hours"`
	Id                                            string                                             `json:"id"`
	Name                                          string                                             `json:"name"`
	RecoveryPointRetentionInHours                 float64                                            `json:"recovery_point_retention_in_hours"`
	RecoveryVaultId                               string                                             `json:"recovery_vault_id"`
	ReplicationIntervalInSeconds                  float64                                            `json:"replication_interval_in_seconds"`
	Timeouts                                      *siterecoveryhypervreplicationpolicy.TimeoutsState `json:"timeouts"`
}
