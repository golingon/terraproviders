// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	mssqlserverextendedauditingpolicy "github.com/golingon/terraproviders/azurerm/3.98.0/mssqlserverextendedauditingpolicy"
	"io"
)

// NewMssqlServerExtendedAuditingPolicy creates a new instance of [MssqlServerExtendedAuditingPolicy].
func NewMssqlServerExtendedAuditingPolicy(name string, args MssqlServerExtendedAuditingPolicyArgs) *MssqlServerExtendedAuditingPolicy {
	return &MssqlServerExtendedAuditingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServerExtendedAuditingPolicy)(nil)

// MssqlServerExtendedAuditingPolicy represents the Terraform resource azurerm_mssql_server_extended_auditing_policy.
type MssqlServerExtendedAuditingPolicy struct {
	Name      string
	Args      MssqlServerExtendedAuditingPolicyArgs
	state     *mssqlServerExtendedAuditingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) Type() string {
	return "azurerm_mssql_server_extended_auditing_policy"
}

// LocalName returns the local name for [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) LocalName() string {
	return mseap.Name
}

// Configuration returns the configuration (args) for [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) Configuration() interface{} {
	return mseap.Args
}

// DependOn is used for other resources to depend on [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(mseap)
}

// Dependencies returns the list of resources [MssqlServerExtendedAuditingPolicy] depends_on.
func (mseap *MssqlServerExtendedAuditingPolicy) Dependencies() terra.Dependencies {
	return mseap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) LifecycleManagement() *terra.Lifecycle {
	return mseap.Lifecycle
}

// Attributes returns the attributes for [MssqlServerExtendedAuditingPolicy].
func (mseap *MssqlServerExtendedAuditingPolicy) Attributes() mssqlServerExtendedAuditingPolicyAttributes {
	return mssqlServerExtendedAuditingPolicyAttributes{ref: terra.ReferenceResource(mseap)}
}

// ImportState imports the given attribute values into [MssqlServerExtendedAuditingPolicy]'s state.
func (mseap *MssqlServerExtendedAuditingPolicy) ImportState(av io.Reader) error {
	mseap.state = &mssqlServerExtendedAuditingPolicyState{}
	if err := json.NewDecoder(av).Decode(mseap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mseap.Type(), mseap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServerExtendedAuditingPolicy] has state.
func (mseap *MssqlServerExtendedAuditingPolicy) State() (*mssqlServerExtendedAuditingPolicyState, bool) {
	return mseap.state, mseap.state != nil
}

// StateMust returns the state for [MssqlServerExtendedAuditingPolicy]. Panics if the state is nil.
func (mseap *MssqlServerExtendedAuditingPolicy) StateMust() *mssqlServerExtendedAuditingPolicyState {
	if mseap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mseap.Type(), mseap.LocalName()))
	}
	return mseap.state
}

// MssqlServerExtendedAuditingPolicyArgs contains the configurations for azurerm_mssql_server_extended_auditing_policy.
type MssqlServerExtendedAuditingPolicyArgs struct {
	// AuditActionsAndGroups: list of string, optional
	AuditActionsAndGroups terra.ListValue[terra.StringValue] `hcl:"audit_actions_and_groups,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogMonitoringEnabled: bool, optional
	LogMonitoringEnabled terra.BoolValue `hcl:"log_monitoring_enabled,attr"`
	// PredicateExpression: string, optional
	PredicateExpression terra.StringValue `hcl:"predicate_expression,attr"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageAccountAccessKeyIsSecondary: bool, optional
	StorageAccountAccessKeyIsSecondary terra.BoolValue `hcl:"storage_account_access_key_is_secondary,attr"`
	// StorageAccountSubscriptionId: string, optional
	StorageAccountSubscriptionId terra.StringValue `hcl:"storage_account_subscription_id,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// Timeouts: optional
	Timeouts *mssqlserverextendedauditingpolicy.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerExtendedAuditingPolicyAttributes struct {
	ref terra.Reference
}

// AuditActionsAndGroups returns a reference to field audit_actions_and_groups of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) AuditActionsAndGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mseap.ref.Append("audit_actions_and_groups"))
}

// Enabled returns a reference to field enabled of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(mseap.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("id"))
}

// LogMonitoringEnabled returns a reference to field log_monitoring_enabled of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) LogMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(mseap.ref.Append("log_monitoring_enabled"))
}

// PredicateExpression returns a reference to field predicate_expression of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) PredicateExpression() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("predicate_expression"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(mseap.ref.Append("retention_in_days"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("server_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("storage_account_access_key"))
}

// StorageAccountAccessKeyIsSecondary returns a reference to field storage_account_access_key_is_secondary of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) StorageAccountAccessKeyIsSecondary() terra.BoolValue {
	return terra.ReferenceAsBool(mseap.ref.Append("storage_account_access_key_is_secondary"))
}

// StorageAccountSubscriptionId returns a reference to field storage_account_subscription_id of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) StorageAccountSubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("storage_account_subscription_id"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_mssql_server_extended_auditing_policy.
func (mseap mssqlServerExtendedAuditingPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mseap.ref.Append("storage_endpoint"))
}

func (mseap mssqlServerExtendedAuditingPolicyAttributes) Timeouts() mssqlserverextendedauditingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlserverextendedauditingpolicy.TimeoutsAttributes](mseap.ref.Append("timeouts"))
}

type mssqlServerExtendedAuditingPolicyState struct {
	AuditActionsAndGroups              []string                                         `json:"audit_actions_and_groups"`
	Enabled                            bool                                             `json:"enabled"`
	Id                                 string                                           `json:"id"`
	LogMonitoringEnabled               bool                                             `json:"log_monitoring_enabled"`
	PredicateExpression                string                                           `json:"predicate_expression"`
	RetentionInDays                    float64                                          `json:"retention_in_days"`
	ServerId                           string                                           `json:"server_id"`
	StorageAccountAccessKey            string                                           `json:"storage_account_access_key"`
	StorageAccountAccessKeyIsSecondary bool                                             `json:"storage_account_access_key_is_secondary"`
	StorageAccountSubscriptionId       string                                           `json:"storage_account_subscription_id"`
	StorageEndpoint                    string                                           `json:"storage_endpoint"`
	Timeouts                           *mssqlserverextendedauditingpolicy.TimeoutsState `json:"timeouts"`
}
