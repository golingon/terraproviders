// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsbackuppolicy "github.com/golingon/terraproviders/aws/5.9.0/efsbackuppolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEfsBackupPolicy creates a new instance of [EfsBackupPolicy].
func NewEfsBackupPolicy(name string, args EfsBackupPolicyArgs) *EfsBackupPolicy {
	return &EfsBackupPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsBackupPolicy)(nil)

// EfsBackupPolicy represents the Terraform resource aws_efs_backup_policy.
type EfsBackupPolicy struct {
	Name      string
	Args      EfsBackupPolicyArgs
	state     *efsBackupPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) Type() string {
	return "aws_efs_backup_policy"
}

// LocalName returns the local name for [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) LocalName() string {
	return ebp.Name
}

// Configuration returns the configuration (args) for [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) Configuration() interface{} {
	return ebp.Args
}

// DependOn is used for other resources to depend on [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ebp)
}

// Dependencies returns the list of resources [EfsBackupPolicy] depends_on.
func (ebp *EfsBackupPolicy) Dependencies() terra.Dependencies {
	return ebp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) LifecycleManagement() *terra.Lifecycle {
	return ebp.Lifecycle
}

// Attributes returns the attributes for [EfsBackupPolicy].
func (ebp *EfsBackupPolicy) Attributes() efsBackupPolicyAttributes {
	return efsBackupPolicyAttributes{ref: terra.ReferenceResource(ebp)}
}

// ImportState imports the given attribute values into [EfsBackupPolicy]'s state.
func (ebp *EfsBackupPolicy) ImportState(av io.Reader) error {
	ebp.state = &efsBackupPolicyState{}
	if err := json.NewDecoder(av).Decode(ebp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ebp.Type(), ebp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EfsBackupPolicy] has state.
func (ebp *EfsBackupPolicy) State() (*efsBackupPolicyState, bool) {
	return ebp.state, ebp.state != nil
}

// StateMust returns the state for [EfsBackupPolicy]. Panics if the state is nil.
func (ebp *EfsBackupPolicy) StateMust() *efsBackupPolicyState {
	if ebp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ebp.Type(), ebp.LocalName()))
	}
	return ebp.state
}

// EfsBackupPolicyArgs contains the configurations for aws_efs_backup_policy.
type EfsBackupPolicyArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// BackupPolicy: required
	BackupPolicy *efsbackuppolicy.BackupPolicy `hcl:"backup_policy,block" validate:"required"`
}
type efsBackupPolicyAttributes struct {
	ref terra.Reference
}

// FileSystemId returns a reference to field file_system_id of aws_efs_backup_policy.
func (ebp efsBackupPolicyAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(ebp.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_backup_policy.
func (ebp efsBackupPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ebp.ref.Append("id"))
}

func (ebp efsBackupPolicyAttributes) BackupPolicy() terra.ListValue[efsbackuppolicy.BackupPolicyAttributes] {
	return terra.ReferenceAsList[efsbackuppolicy.BackupPolicyAttributes](ebp.ref.Append("backup_policy"))
}

type efsBackupPolicyState struct {
	FileSystemId string                              `json:"file_system_id"`
	Id           string                              `json:"id"`
	BackupPolicy []efsbackuppolicy.BackupPolicyState `json:"backup_policy"`
}
