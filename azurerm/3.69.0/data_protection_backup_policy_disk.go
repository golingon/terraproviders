// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackuppolicydisk "github.com/golingon/terraproviders/azurerm/3.69.0/dataprotectionbackuppolicydisk"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupPolicyDisk creates a new instance of [DataProtectionBackupPolicyDisk].
func NewDataProtectionBackupPolicyDisk(name string, args DataProtectionBackupPolicyDiskArgs) *DataProtectionBackupPolicyDisk {
	return &DataProtectionBackupPolicyDisk{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupPolicyDisk)(nil)

// DataProtectionBackupPolicyDisk represents the Terraform resource azurerm_data_protection_backup_policy_disk.
type DataProtectionBackupPolicyDisk struct {
	Name      string
	Args      DataProtectionBackupPolicyDiskArgs
	state     *dataProtectionBackupPolicyDiskState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) Type() string {
	return "azurerm_data_protection_backup_policy_disk"
}

// LocalName returns the local name for [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) LocalName() string {
	return dpbpd.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) Configuration() interface{} {
	return dpbpd.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbpd)
}

// Dependencies returns the list of resources [DataProtectionBackupPolicyDisk] depends_on.
func (dpbpd *DataProtectionBackupPolicyDisk) Dependencies() terra.Dependencies {
	return dpbpd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) LifecycleManagement() *terra.Lifecycle {
	return dpbpd.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupPolicyDisk].
func (dpbpd *DataProtectionBackupPolicyDisk) Attributes() dataProtectionBackupPolicyDiskAttributes {
	return dataProtectionBackupPolicyDiskAttributes{ref: terra.ReferenceResource(dpbpd)}
}

// ImportState imports the given attribute values into [DataProtectionBackupPolicyDisk]'s state.
func (dpbpd *DataProtectionBackupPolicyDisk) ImportState(av io.Reader) error {
	dpbpd.state = &dataProtectionBackupPolicyDiskState{}
	if err := json.NewDecoder(av).Decode(dpbpd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbpd.Type(), dpbpd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupPolicyDisk] has state.
func (dpbpd *DataProtectionBackupPolicyDisk) State() (*dataProtectionBackupPolicyDiskState, bool) {
	return dpbpd.state, dpbpd.state != nil
}

// StateMust returns the state for [DataProtectionBackupPolicyDisk]. Panics if the state is nil.
func (dpbpd *DataProtectionBackupPolicyDisk) StateMust() *dataProtectionBackupPolicyDiskState {
	if dpbpd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbpd.Type(), dpbpd.LocalName()))
	}
	return dpbpd.state
}

// DataProtectionBackupPolicyDiskArgs contains the configurations for azurerm_data_protection_backup_policy_disk.
type DataProtectionBackupPolicyDiskArgs struct {
	// BackupRepeatingTimeIntervals: list of string, required
	BackupRepeatingTimeIntervals terra.ListValue[terra.StringValue] `hcl:"backup_repeating_time_intervals,attr" validate:"required"`
	// DefaultRetentionDuration: string, required
	DefaultRetentionDuration terra.StringValue `hcl:"default_retention_duration,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// VaultId: string, required
	VaultId terra.StringValue `hcl:"vault_id,attr" validate:"required"`
	// RetentionRule: min=0
	RetentionRule []dataprotectionbackuppolicydisk.RetentionRule `hcl:"retention_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprotectionbackuppolicydisk.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupPolicyDiskAttributes struct {
	ref terra.Reference
}

// BackupRepeatingTimeIntervals returns a reference to field backup_repeating_time_intervals of azurerm_data_protection_backup_policy_disk.
func (dpbpd dataProtectionBackupPolicyDiskAttributes) BackupRepeatingTimeIntervals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dpbpd.ref.Append("backup_repeating_time_intervals"))
}

// DefaultRetentionDuration returns a reference to field default_retention_duration of azurerm_data_protection_backup_policy_disk.
func (dpbpd dataProtectionBackupPolicyDiskAttributes) DefaultRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(dpbpd.ref.Append("default_retention_duration"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_policy_disk.
func (dpbpd dataProtectionBackupPolicyDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbpd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_policy_disk.
func (dpbpd dataProtectionBackupPolicyDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbpd.ref.Append("name"))
}

// VaultId returns a reference to field vault_id of azurerm_data_protection_backup_policy_disk.
func (dpbpd dataProtectionBackupPolicyDiskAttributes) VaultId() terra.StringValue {
	return terra.ReferenceAsString(dpbpd.ref.Append("vault_id"))
}

func (dpbpd dataProtectionBackupPolicyDiskAttributes) RetentionRule() terra.ListValue[dataprotectionbackuppolicydisk.RetentionRuleAttributes] {
	return terra.ReferenceAsList[dataprotectionbackuppolicydisk.RetentionRuleAttributes](dpbpd.ref.Append("retention_rule"))
}

func (dpbpd dataProtectionBackupPolicyDiskAttributes) Timeouts() dataprotectionbackuppolicydisk.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackuppolicydisk.TimeoutsAttributes](dpbpd.ref.Append("timeouts"))
}

type dataProtectionBackupPolicyDiskState struct {
	BackupRepeatingTimeIntervals []string                                            `json:"backup_repeating_time_intervals"`
	DefaultRetentionDuration     string                                              `json:"default_retention_duration"`
	Id                           string                                              `json:"id"`
	Name                         string                                              `json:"name"`
	VaultId                      string                                              `json:"vault_id"`
	RetentionRule                []dataprotectionbackuppolicydisk.RetentionRuleState `json:"retention_rule"`
	Timeouts                     *dataprotectionbackuppolicydisk.TimeoutsState       `json:"timeouts"`
}
