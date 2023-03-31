// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupVaultPolicy creates a new instance of [BackupVaultPolicy].
func NewBackupVaultPolicy(name string, args BackupVaultPolicyArgs) *BackupVaultPolicy {
	return &BackupVaultPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupVaultPolicy)(nil)

// BackupVaultPolicy represents the Terraform resource aws_backup_vault_policy.
type BackupVaultPolicy struct {
	Name      string
	Args      BackupVaultPolicyArgs
	state     *backupVaultPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) Type() string {
	return "aws_backup_vault_policy"
}

// LocalName returns the local name for [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) LocalName() string {
	return bvp.Name
}

// Configuration returns the configuration (args) for [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) Configuration() interface{} {
	return bvp.Args
}

// DependOn is used for other resources to depend on [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(bvp)
}

// Dependencies returns the list of resources [BackupVaultPolicy] depends_on.
func (bvp *BackupVaultPolicy) Dependencies() terra.Dependencies {
	return bvp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) LifecycleManagement() *terra.Lifecycle {
	return bvp.Lifecycle
}

// Attributes returns the attributes for [BackupVaultPolicy].
func (bvp *BackupVaultPolicy) Attributes() backupVaultPolicyAttributes {
	return backupVaultPolicyAttributes{ref: terra.ReferenceResource(bvp)}
}

// ImportState imports the given attribute values into [BackupVaultPolicy]'s state.
func (bvp *BackupVaultPolicy) ImportState(av io.Reader) error {
	bvp.state = &backupVaultPolicyState{}
	if err := json.NewDecoder(av).Decode(bvp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bvp.Type(), bvp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupVaultPolicy] has state.
func (bvp *BackupVaultPolicy) State() (*backupVaultPolicyState, bool) {
	return bvp.state, bvp.state != nil
}

// StateMust returns the state for [BackupVaultPolicy]. Panics if the state is nil.
func (bvp *BackupVaultPolicy) StateMust() *backupVaultPolicyState {
	if bvp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bvp.Type(), bvp.LocalName()))
	}
	return bvp.state
}

// BackupVaultPolicyArgs contains the configurations for aws_backup_vault_policy.
type BackupVaultPolicyArgs struct {
	// BackupVaultName: string, required
	BackupVaultName terra.StringValue `hcl:"backup_vault_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type backupVaultPolicyAttributes struct {
	ref terra.Reference
}

// BackupVaultArn returns a reference to field backup_vault_arn of aws_backup_vault_policy.
func (bvp backupVaultPolicyAttributes) BackupVaultArn() terra.StringValue {
	return terra.ReferenceAsString(bvp.ref.Append("backup_vault_arn"))
}

// BackupVaultName returns a reference to field backup_vault_name of aws_backup_vault_policy.
func (bvp backupVaultPolicyAttributes) BackupVaultName() terra.StringValue {
	return terra.ReferenceAsString(bvp.ref.Append("backup_vault_name"))
}

// Id returns a reference to field id of aws_backup_vault_policy.
func (bvp backupVaultPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bvp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_backup_vault_policy.
func (bvp backupVaultPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(bvp.ref.Append("policy"))
}

type backupVaultPolicyState struct {
	BackupVaultArn  string `json:"backup_vault_arn"`
	BackupVaultName string `json:"backup_vault_name"`
	Id              string `json:"id"`
	Policy          string `json:"policy"`
}