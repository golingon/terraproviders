// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_backup_policy_vm_workload

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_backup_policy_vm_workload.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermBackupPolicyVmWorkloadState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (abpvw *Resource) Type() string {
	return "azurerm_backup_policy_vm_workload"
}

// LocalName returns the local name for [Resource].
func (abpvw *Resource) LocalName() string {
	return abpvw.Name
}

// Configuration returns the configuration (args) for [Resource].
func (abpvw *Resource) Configuration() interface{} {
	return abpvw.Args
}

// DependOn is used for other resources to depend on [Resource].
func (abpvw *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(abpvw)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (abpvw *Resource) Dependencies() terra.Dependencies {
	return abpvw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (abpvw *Resource) LifecycleManagement() *terra.Lifecycle {
	return abpvw.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (abpvw *Resource) Attributes() azurermBackupPolicyVmWorkloadAttributes {
	return azurermBackupPolicyVmWorkloadAttributes{ref: terra.ReferenceResource(abpvw)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (abpvw *Resource) ImportState(state io.Reader) error {
	abpvw.state = &azurermBackupPolicyVmWorkloadState{}
	if err := json.NewDecoder(state).Decode(abpvw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", abpvw.Type(), abpvw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (abpvw *Resource) State() (*azurermBackupPolicyVmWorkloadState, bool) {
	return abpvw.state, abpvw.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (abpvw *Resource) StateMust() *azurermBackupPolicyVmWorkloadState {
	if abpvw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", abpvw.Type(), abpvw.LocalName()))
	}
	return abpvw.state
}

// Args contains the configurations for azurerm_backup_policy_vm_workload.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// WorkloadType: string, required
	WorkloadType terra.StringValue `hcl:"workload_type,attr" validate:"required"`
	// ProtectionPolicy: min=1
	ProtectionPolicy []ProtectionPolicy `hcl:"protection_policy,block" validate:"min=1"`
	// Settings: required
	Settings *Settings `hcl:"settings,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermBackupPolicyVmWorkloadAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_backup_policy_vm_workload.
func (abpvw azurermBackupPolicyVmWorkloadAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(abpvw.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_backup_policy_vm_workload.
func (abpvw azurermBackupPolicyVmWorkloadAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(abpvw.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_backup_policy_vm_workload.
func (abpvw azurermBackupPolicyVmWorkloadAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(abpvw.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_backup_policy_vm_workload.
func (abpvw azurermBackupPolicyVmWorkloadAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(abpvw.ref.Append("resource_group_name"))
}

// WorkloadType returns a reference to field workload_type of azurerm_backup_policy_vm_workload.
func (abpvw azurermBackupPolicyVmWorkloadAttributes) WorkloadType() terra.StringValue {
	return terra.ReferenceAsString(abpvw.ref.Append("workload_type"))
}

func (abpvw azurermBackupPolicyVmWorkloadAttributes) ProtectionPolicy() terra.SetValue[ProtectionPolicyAttributes] {
	return terra.ReferenceAsSet[ProtectionPolicyAttributes](abpvw.ref.Append("protection_policy"))
}

func (abpvw azurermBackupPolicyVmWorkloadAttributes) Settings() terra.ListValue[SettingsAttributes] {
	return terra.ReferenceAsList[SettingsAttributes](abpvw.ref.Append("settings"))
}

func (abpvw azurermBackupPolicyVmWorkloadAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](abpvw.ref.Append("timeouts"))
}

type azurermBackupPolicyVmWorkloadState struct {
	Id                string                  `json:"id"`
	Name              string                  `json:"name"`
	RecoveryVaultName string                  `json:"recovery_vault_name"`
	ResourceGroupName string                  `json:"resource_group_name"`
	WorkloadType      string                  `json:"workload_type"`
	ProtectionPolicy  []ProtectionPolicyState `json:"protection_policy"`
	Settings          []SettingsState         `json:"settings"`
	Timeouts          *TimeoutsState          `json:"timeouts"`
}
