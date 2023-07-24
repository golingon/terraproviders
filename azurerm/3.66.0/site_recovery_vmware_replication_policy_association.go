// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	siterecoveryvmwarereplicationpolicyassociation "github.com/golingon/terraproviders/azurerm/3.66.0/siterecoveryvmwarereplicationpolicyassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSiteRecoveryVmwareReplicationPolicyAssociation creates a new instance of [SiteRecoveryVmwareReplicationPolicyAssociation].
func NewSiteRecoveryVmwareReplicationPolicyAssociation(name string, args SiteRecoveryVmwareReplicationPolicyAssociationArgs) *SiteRecoveryVmwareReplicationPolicyAssociation {
	return &SiteRecoveryVmwareReplicationPolicyAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SiteRecoveryVmwareReplicationPolicyAssociation)(nil)

// SiteRecoveryVmwareReplicationPolicyAssociation represents the Terraform resource azurerm_site_recovery_vmware_replication_policy_association.
type SiteRecoveryVmwareReplicationPolicyAssociation struct {
	Name      string
	Args      SiteRecoveryVmwareReplicationPolicyAssociationArgs
	state     *siteRecoveryVmwareReplicationPolicyAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) Type() string {
	return "azurerm_site_recovery_vmware_replication_policy_association"
}

// LocalName returns the local name for [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) LocalName() string {
	return srvrpa.Name
}

// Configuration returns the configuration (args) for [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) Configuration() interface{} {
	return srvrpa.Args
}

// DependOn is used for other resources to depend on [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(srvrpa)
}

// Dependencies returns the list of resources [SiteRecoveryVmwareReplicationPolicyAssociation] depends_on.
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) Dependencies() terra.Dependencies {
	return srvrpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) LifecycleManagement() *terra.Lifecycle {
	return srvrpa.Lifecycle
}

// Attributes returns the attributes for [SiteRecoveryVmwareReplicationPolicyAssociation].
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) Attributes() siteRecoveryVmwareReplicationPolicyAssociationAttributes {
	return siteRecoveryVmwareReplicationPolicyAssociationAttributes{ref: terra.ReferenceResource(srvrpa)}
}

// ImportState imports the given attribute values into [SiteRecoveryVmwareReplicationPolicyAssociation]'s state.
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) ImportState(av io.Reader) error {
	srvrpa.state = &siteRecoveryVmwareReplicationPolicyAssociationState{}
	if err := json.NewDecoder(av).Decode(srvrpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", srvrpa.Type(), srvrpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SiteRecoveryVmwareReplicationPolicyAssociation] has state.
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) State() (*siteRecoveryVmwareReplicationPolicyAssociationState, bool) {
	return srvrpa.state, srvrpa.state != nil
}

// StateMust returns the state for [SiteRecoveryVmwareReplicationPolicyAssociation]. Panics if the state is nil.
func (srvrpa *SiteRecoveryVmwareReplicationPolicyAssociation) StateMust() *siteRecoveryVmwareReplicationPolicyAssociationState {
	if srvrpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", srvrpa.Type(), srvrpa.LocalName()))
	}
	return srvrpa.state
}

// SiteRecoveryVmwareReplicationPolicyAssociationArgs contains the configurations for azurerm_site_recovery_vmware_replication_policy_association.
type SiteRecoveryVmwareReplicationPolicyAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyId: string, required
	PolicyId terra.StringValue `hcl:"policy_id,attr" validate:"required"`
	// RecoveryVaultId: string, required
	RecoveryVaultId terra.StringValue `hcl:"recovery_vault_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *siterecoveryvmwarereplicationpolicyassociation.Timeouts `hcl:"timeouts,block"`
}
type siteRecoveryVmwareReplicationPolicyAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_vmware_replication_policy_association.
func (srvrpa siteRecoveryVmwareReplicationPolicyAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srvrpa.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_vmware_replication_policy_association.
func (srvrpa siteRecoveryVmwareReplicationPolicyAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srvrpa.ref.Append("name"))
}

// PolicyId returns a reference to field policy_id of azurerm_site_recovery_vmware_replication_policy_association.
func (srvrpa siteRecoveryVmwareReplicationPolicyAssociationAttributes) PolicyId() terra.StringValue {
	return terra.ReferenceAsString(srvrpa.ref.Append("policy_id"))
}

// RecoveryVaultId returns a reference to field recovery_vault_id of azurerm_site_recovery_vmware_replication_policy_association.
func (srvrpa siteRecoveryVmwareReplicationPolicyAssociationAttributes) RecoveryVaultId() terra.StringValue {
	return terra.ReferenceAsString(srvrpa.ref.Append("recovery_vault_id"))
}

func (srvrpa siteRecoveryVmwareReplicationPolicyAssociationAttributes) Timeouts() siterecoveryvmwarereplicationpolicyassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[siterecoveryvmwarereplicationpolicyassociation.TimeoutsAttributes](srvrpa.ref.Append("timeouts"))
}

type siteRecoveryVmwareReplicationPolicyAssociationState struct {
	Id              string                                                        `json:"id"`
	Name            string                                                        `json:"name"`
	PolicyId        string                                                        `json:"policy_id"`
	RecoveryVaultId string                                                        `json:"recovery_vault_id"`
	Timeouts        *siterecoveryvmwarereplicationpolicyassociation.TimeoutsState `json:"timeouts"`
}
