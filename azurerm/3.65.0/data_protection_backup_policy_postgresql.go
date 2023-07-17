// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackuppolicypostgresql "github.com/golingon/terraproviders/azurerm/3.65.0/dataprotectionbackuppolicypostgresql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataProtectionBackupPolicyPostgresql creates a new instance of [DataProtectionBackupPolicyPostgresql].
func NewDataProtectionBackupPolicyPostgresql(name string, args DataProtectionBackupPolicyPostgresqlArgs) *DataProtectionBackupPolicyPostgresql {
	return &DataProtectionBackupPolicyPostgresql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupPolicyPostgresql)(nil)

// DataProtectionBackupPolicyPostgresql represents the Terraform resource azurerm_data_protection_backup_policy_postgresql.
type DataProtectionBackupPolicyPostgresql struct {
	Name      string
	Args      DataProtectionBackupPolicyPostgresqlArgs
	state     *dataProtectionBackupPolicyPostgresqlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) Type() string {
	return "azurerm_data_protection_backup_policy_postgresql"
}

// LocalName returns the local name for [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) LocalName() string {
	return dpbpp.Name
}

// Configuration returns the configuration (args) for [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) Configuration() interface{} {
	return dpbpp.Args
}

// DependOn is used for other resources to depend on [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbpp)
}

// Dependencies returns the list of resources [DataProtectionBackupPolicyPostgresql] depends_on.
func (dpbpp *DataProtectionBackupPolicyPostgresql) Dependencies() terra.Dependencies {
	return dpbpp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) LifecycleManagement() *terra.Lifecycle {
	return dpbpp.Lifecycle
}

// Attributes returns the attributes for [DataProtectionBackupPolicyPostgresql].
func (dpbpp *DataProtectionBackupPolicyPostgresql) Attributes() dataProtectionBackupPolicyPostgresqlAttributes {
	return dataProtectionBackupPolicyPostgresqlAttributes{ref: terra.ReferenceResource(dpbpp)}
}

// ImportState imports the given attribute values into [DataProtectionBackupPolicyPostgresql]'s state.
func (dpbpp *DataProtectionBackupPolicyPostgresql) ImportState(av io.Reader) error {
	dpbpp.state = &dataProtectionBackupPolicyPostgresqlState{}
	if err := json.NewDecoder(av).Decode(dpbpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbpp.Type(), dpbpp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataProtectionBackupPolicyPostgresql] has state.
func (dpbpp *DataProtectionBackupPolicyPostgresql) State() (*dataProtectionBackupPolicyPostgresqlState, bool) {
	return dpbpp.state, dpbpp.state != nil
}

// StateMust returns the state for [DataProtectionBackupPolicyPostgresql]. Panics if the state is nil.
func (dpbpp *DataProtectionBackupPolicyPostgresql) StateMust() *dataProtectionBackupPolicyPostgresqlState {
	if dpbpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbpp.Type(), dpbpp.LocalName()))
	}
	return dpbpp.state
}

// DataProtectionBackupPolicyPostgresqlArgs contains the configurations for azurerm_data_protection_backup_policy_postgresql.
type DataProtectionBackupPolicyPostgresqlArgs struct {
	// BackupRepeatingTimeIntervals: list of string, required
	BackupRepeatingTimeIntervals terra.ListValue[terra.StringValue] `hcl:"backup_repeating_time_intervals,attr" validate:"required"`
	// DefaultRetentionDuration: string, required
	DefaultRetentionDuration terra.StringValue `hcl:"default_retention_duration,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// VaultName: string, required
	VaultName terra.StringValue `hcl:"vault_name,attr" validate:"required"`
	// RetentionRule: min=0
	RetentionRule []dataprotectionbackuppolicypostgresql.RetentionRule `hcl:"retention_rule,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprotectionbackuppolicypostgresql.Timeouts `hcl:"timeouts,block"`
}
type dataProtectionBackupPolicyPostgresqlAttributes struct {
	ref terra.Reference
}

// BackupRepeatingTimeIntervals returns a reference to field backup_repeating_time_intervals of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) BackupRepeatingTimeIntervals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dpbpp.ref.Append("backup_repeating_time_intervals"))
}

// DefaultRetentionDuration returns a reference to field default_retention_duration of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) DefaultRetentionDuration() terra.StringValue {
	return terra.ReferenceAsString(dpbpp.ref.Append("default_retention_duration"))
}

// Id returns a reference to field id of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dpbpp.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dpbpp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(dpbpp.ref.Append("resource_group_name"))
}

// VaultName returns a reference to field vault_name of azurerm_data_protection_backup_policy_postgresql.
func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) VaultName() terra.StringValue {
	return terra.ReferenceAsString(dpbpp.ref.Append("vault_name"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) RetentionRule() terra.ListValue[dataprotectionbackuppolicypostgresql.RetentionRuleAttributes] {
	return terra.ReferenceAsList[dataprotectionbackuppolicypostgresql.RetentionRuleAttributes](dpbpp.ref.Append("retention_rule"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Timeouts() dataprotectionbackuppolicypostgresql.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprotectionbackuppolicypostgresql.TimeoutsAttributes](dpbpp.ref.Append("timeouts"))
}

type dataProtectionBackupPolicyPostgresqlState struct {
	BackupRepeatingTimeIntervals []string                                                  `json:"backup_repeating_time_intervals"`
	DefaultRetentionDuration     string                                                    `json:"default_retention_duration"`
	Id                           string                                                    `json:"id"`
	Name                         string                                                    `json:"name"`
	ResourceGroupName            string                                                    `json:"resource_group_name"`
	VaultName                    string                                                    `json:"vault_name"`
	RetentionRule                []dataprotectionbackuppolicypostgresql.RetentionRuleState `json:"retention_rule"`
	Timeouts                     *dataprotectionbackuppolicypostgresql.TimeoutsState       `json:"timeouts"`
}
