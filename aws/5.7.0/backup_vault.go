// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	backupvault "github.com/golingon/terraproviders/aws/5.7.0/backupvault"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBackupVault creates a new instance of [BackupVault].
func NewBackupVault(name string, args BackupVaultArgs) *BackupVault {
	return &BackupVault{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BackupVault)(nil)

// BackupVault represents the Terraform resource aws_backup_vault.
type BackupVault struct {
	Name      string
	Args      BackupVaultArgs
	state     *backupVaultState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BackupVault].
func (bv *BackupVault) Type() string {
	return "aws_backup_vault"
}

// LocalName returns the local name for [BackupVault].
func (bv *BackupVault) LocalName() string {
	return bv.Name
}

// Configuration returns the configuration (args) for [BackupVault].
func (bv *BackupVault) Configuration() interface{} {
	return bv.Args
}

// DependOn is used for other resources to depend on [BackupVault].
func (bv *BackupVault) DependOn() terra.Reference {
	return terra.ReferenceResource(bv)
}

// Dependencies returns the list of resources [BackupVault] depends_on.
func (bv *BackupVault) Dependencies() terra.Dependencies {
	return bv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BackupVault].
func (bv *BackupVault) LifecycleManagement() *terra.Lifecycle {
	return bv.Lifecycle
}

// Attributes returns the attributes for [BackupVault].
func (bv *BackupVault) Attributes() backupVaultAttributes {
	return backupVaultAttributes{ref: terra.ReferenceResource(bv)}
}

// ImportState imports the given attribute values into [BackupVault]'s state.
func (bv *BackupVault) ImportState(av io.Reader) error {
	bv.state = &backupVaultState{}
	if err := json.NewDecoder(av).Decode(bv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bv.Type(), bv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BackupVault] has state.
func (bv *BackupVault) State() (*backupVaultState, bool) {
	return bv.state, bv.state != nil
}

// StateMust returns the state for [BackupVault]. Panics if the state is nil.
func (bv *BackupVault) StateMust() *backupVaultState {
	if bv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bv.Type(), bv.LocalName()))
	}
	return bv.state
}

// BackupVaultArgs contains the configurations for aws_backup_vault.
type BackupVaultArgs struct {
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyArn: string, optional
	KmsKeyArn terra.StringValue `hcl:"kms_key_arn,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *backupvault.Timeouts `hcl:"timeouts,block"`
}
type backupVaultAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_backup_vault.
func (bv backupVaultAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(bv.ref.Append("arn"))
}

// ForceDestroy returns a reference to field force_destroy of aws_backup_vault.
func (bv backupVaultAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(bv.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_backup_vault.
func (bv backupVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bv.ref.Append("id"))
}

// KmsKeyArn returns a reference to field kms_key_arn of aws_backup_vault.
func (bv backupVaultAttributes) KmsKeyArn() terra.StringValue {
	return terra.ReferenceAsString(bv.ref.Append("kms_key_arn"))
}

// Name returns a reference to field name of aws_backup_vault.
func (bv backupVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(bv.ref.Append("name"))
}

// RecoveryPoints returns a reference to field recovery_points of aws_backup_vault.
func (bv backupVaultAttributes) RecoveryPoints() terra.NumberValue {
	return terra.ReferenceAsNumber(bv.ref.Append("recovery_points"))
}

// Tags returns a reference to field tags of aws_backup_vault.
func (bv backupVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bv.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_backup_vault.
func (bv backupVaultAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bv.ref.Append("tags_all"))
}

func (bv backupVaultAttributes) Timeouts() backupvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[backupvault.TimeoutsAttributes](bv.ref.Append("timeouts"))
}

type backupVaultState struct {
	Arn            string                     `json:"arn"`
	ForceDestroy   bool                       `json:"force_destroy"`
	Id             string                     `json:"id"`
	KmsKeyArn      string                     `json:"kms_key_arn"`
	Name           string                     `json:"name"`
	RecoveryPoints float64                    `json:"recovery_points"`
	Tags           map[string]string          `json:"tags"`
	TagsAll        map[string]string          `json:"tags_all"`
	Timeouts       *backupvault.TimeoutsState `json:"timeouts"`
}
