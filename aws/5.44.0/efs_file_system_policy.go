// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewEfsFileSystemPolicy creates a new instance of [EfsFileSystemPolicy].
func NewEfsFileSystemPolicy(name string, args EfsFileSystemPolicyArgs) *EfsFileSystemPolicy {
	return &EfsFileSystemPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsFileSystemPolicy)(nil)

// EfsFileSystemPolicy represents the Terraform resource aws_efs_file_system_policy.
type EfsFileSystemPolicy struct {
	Name      string
	Args      EfsFileSystemPolicyArgs
	state     *efsFileSystemPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) Type() string {
	return "aws_efs_file_system_policy"
}

// LocalName returns the local name for [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) LocalName() string {
	return efsp.Name
}

// Configuration returns the configuration (args) for [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) Configuration() interface{} {
	return efsp.Args
}

// DependOn is used for other resources to depend on [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(efsp)
}

// Dependencies returns the list of resources [EfsFileSystemPolicy] depends_on.
func (efsp *EfsFileSystemPolicy) Dependencies() terra.Dependencies {
	return efsp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) LifecycleManagement() *terra.Lifecycle {
	return efsp.Lifecycle
}

// Attributes returns the attributes for [EfsFileSystemPolicy].
func (efsp *EfsFileSystemPolicy) Attributes() efsFileSystemPolicyAttributes {
	return efsFileSystemPolicyAttributes{ref: terra.ReferenceResource(efsp)}
}

// ImportState imports the given attribute values into [EfsFileSystemPolicy]'s state.
func (efsp *EfsFileSystemPolicy) ImportState(av io.Reader) error {
	efsp.state = &efsFileSystemPolicyState{}
	if err := json.NewDecoder(av).Decode(efsp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", efsp.Type(), efsp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EfsFileSystemPolicy] has state.
func (efsp *EfsFileSystemPolicy) State() (*efsFileSystemPolicyState, bool) {
	return efsp.state, efsp.state != nil
}

// StateMust returns the state for [EfsFileSystemPolicy]. Panics if the state is nil.
func (efsp *EfsFileSystemPolicy) StateMust() *efsFileSystemPolicyState {
	if efsp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", efsp.Type(), efsp.LocalName()))
	}
	return efsp.state
}

// EfsFileSystemPolicyArgs contains the configurations for aws_efs_file_system_policy.
type EfsFileSystemPolicyArgs struct {
	// BypassPolicyLockoutSafetyCheck: bool, optional
	BypassPolicyLockoutSafetyCheck terra.BoolValue `hcl:"bypass_policy_lockout_safety_check,attr"`
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Policy: string, required
	Policy terra.StringValue `hcl:"policy,attr" validate:"required"`
}
type efsFileSystemPolicyAttributes struct {
	ref terra.Reference
}

// BypassPolicyLockoutSafetyCheck returns a reference to field bypass_policy_lockout_safety_check of aws_efs_file_system_policy.
func (efsp efsFileSystemPolicyAttributes) BypassPolicyLockoutSafetyCheck() terra.BoolValue {
	return terra.ReferenceAsBool(efsp.ref.Append("bypass_policy_lockout_safety_check"))
}

// FileSystemId returns a reference to field file_system_id of aws_efs_file_system_policy.
func (efsp efsFileSystemPolicyAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceAsString(efsp.ref.Append("file_system_id"))
}

// Id returns a reference to field id of aws_efs_file_system_policy.
func (efsp efsFileSystemPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(efsp.ref.Append("id"))
}

// Policy returns a reference to field policy of aws_efs_file_system_policy.
func (efsp efsFileSystemPolicyAttributes) Policy() terra.StringValue {
	return terra.ReferenceAsString(efsp.ref.Append("policy"))
}

type efsFileSystemPolicyState struct {
	BypassPolicyLockoutSafetyCheck bool   `json:"bypass_policy_lockout_safety_check"`
	FileSystemId                   string `json:"file_system_id"`
	Id                             string `json:"id"`
	Policy                         string `json:"policy"`
}
