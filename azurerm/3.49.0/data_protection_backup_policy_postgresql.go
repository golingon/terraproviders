// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	dataprotectionbackuppolicypostgresql "github.com/golingon/terraproviders/azurerm/3.49.0/dataprotectionbackuppolicypostgresql"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataProtectionBackupPolicyPostgresql(name string, args DataProtectionBackupPolicyPostgresqlArgs) *DataProtectionBackupPolicyPostgresql {
	return &DataProtectionBackupPolicyPostgresql{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataProtectionBackupPolicyPostgresql)(nil)

type DataProtectionBackupPolicyPostgresql struct {
	Name  string
	Args  DataProtectionBackupPolicyPostgresqlArgs
	state *dataProtectionBackupPolicyPostgresqlState
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) Type() string {
	return "azurerm_data_protection_backup_policy_postgresql"
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) LocalName() string {
	return dpbpp.Name
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) Configuration() interface{} {
	return dpbpp.Args
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) Attributes() dataProtectionBackupPolicyPostgresqlAttributes {
	return dataProtectionBackupPolicyPostgresqlAttributes{ref: terra.ReferenceResource(dpbpp)}
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) ImportState(av io.Reader) error {
	dpbpp.state = &dataProtectionBackupPolicyPostgresqlState{}
	if err := json.NewDecoder(av).Decode(dpbpp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dpbpp.Type(), dpbpp.LocalName(), err)
	}
	return nil
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) State() (*dataProtectionBackupPolicyPostgresqlState, bool) {
	return dpbpp.state, dpbpp.state != nil
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) StateMust() *dataProtectionBackupPolicyPostgresqlState {
	if dpbpp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dpbpp.Type(), dpbpp.LocalName()))
	}
	return dpbpp.state
}

func (dpbpp *DataProtectionBackupPolicyPostgresql) DependOn() terra.Reference {
	return terra.ReferenceResource(dpbpp)
}

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
	// DependsOn contains resources that DataProtectionBackupPolicyPostgresql depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataProtectionBackupPolicyPostgresqlAttributes struct {
	ref terra.Reference
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) BackupRepeatingTimeIntervals() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dpbpp.ref.Append("backup_repeating_time_intervals"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) DefaultRetentionDuration() terra.StringValue {
	return terra.ReferenceString(dpbpp.ref.Append("default_retention_duration"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dpbpp.ref.Append("id"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dpbpp.ref.Append("name"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(dpbpp.ref.Append("resource_group_name"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) VaultName() terra.StringValue {
	return terra.ReferenceString(dpbpp.ref.Append("vault_name"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) RetentionRule() terra.ListValue[dataprotectionbackuppolicypostgresql.RetentionRuleAttributes] {
	return terra.ReferenceList[dataprotectionbackuppolicypostgresql.RetentionRuleAttributes](dpbpp.ref.Append("retention_rule"))
}

func (dpbpp dataProtectionBackupPolicyPostgresqlAttributes) Timeouts() dataprotectionbackuppolicypostgresql.TimeoutsAttributes {
	return terra.ReferenceSingle[dataprotectionbackuppolicypostgresql.TimeoutsAttributes](dpbpp.ref.Append("timeouts"))
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
