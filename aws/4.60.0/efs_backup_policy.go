// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	efsbackuppolicy "github.com/golingon/terraproviders/aws/4.60.0/efsbackuppolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewEfsBackupPolicy(name string, args EfsBackupPolicyArgs) *EfsBackupPolicy {
	return &EfsBackupPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EfsBackupPolicy)(nil)

type EfsBackupPolicy struct {
	Name  string
	Args  EfsBackupPolicyArgs
	state *efsBackupPolicyState
}

func (ebp *EfsBackupPolicy) Type() string {
	return "aws_efs_backup_policy"
}

func (ebp *EfsBackupPolicy) LocalName() string {
	return ebp.Name
}

func (ebp *EfsBackupPolicy) Configuration() interface{} {
	return ebp.Args
}

func (ebp *EfsBackupPolicy) Attributes() efsBackupPolicyAttributes {
	return efsBackupPolicyAttributes{ref: terra.ReferenceResource(ebp)}
}

func (ebp *EfsBackupPolicy) ImportState(av io.Reader) error {
	ebp.state = &efsBackupPolicyState{}
	if err := json.NewDecoder(av).Decode(ebp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ebp.Type(), ebp.LocalName(), err)
	}
	return nil
}

func (ebp *EfsBackupPolicy) State() (*efsBackupPolicyState, bool) {
	return ebp.state, ebp.state != nil
}

func (ebp *EfsBackupPolicy) StateMust() *efsBackupPolicyState {
	if ebp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ebp.Type(), ebp.LocalName()))
	}
	return ebp.state
}

func (ebp *EfsBackupPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(ebp)
}

type EfsBackupPolicyArgs struct {
	// FileSystemId: string, required
	FileSystemId terra.StringValue `hcl:"file_system_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// BackupPolicy: required
	BackupPolicy *efsbackuppolicy.BackupPolicy `hcl:"backup_policy,block" validate:"required"`
	// DependsOn contains resources that EfsBackupPolicy depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type efsBackupPolicyAttributes struct {
	ref terra.Reference
}

func (ebp efsBackupPolicyAttributes) FileSystemId() terra.StringValue {
	return terra.ReferenceString(ebp.ref.Append("file_system_id"))
}

func (ebp efsBackupPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ebp.ref.Append("id"))
}

func (ebp efsBackupPolicyAttributes) BackupPolicy() terra.ListValue[efsbackuppolicy.BackupPolicyAttributes] {
	return terra.ReferenceList[efsbackuppolicy.BackupPolicyAttributes](ebp.ref.Append("backup_policy"))
}

type efsBackupPolicyState struct {
	FileSystemId string                              `json:"file_system_id"`
	Id           string                              `json:"id"`
	BackupPolicy []efsbackuppolicy.BackupPolicyState `json:"backup_policy"`
}
