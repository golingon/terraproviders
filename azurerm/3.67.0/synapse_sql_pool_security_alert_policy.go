// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsesqlpoolsecurityalertpolicy "github.com/golingon/terraproviders/azurerm/3.67.0/synapsesqlpoolsecurityalertpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseSqlPoolSecurityAlertPolicy creates a new instance of [SynapseSqlPoolSecurityAlertPolicy].
func NewSynapseSqlPoolSecurityAlertPolicy(name string, args SynapseSqlPoolSecurityAlertPolicyArgs) *SynapseSqlPoolSecurityAlertPolicy {
	return &SynapseSqlPoolSecurityAlertPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseSqlPoolSecurityAlertPolicy)(nil)

// SynapseSqlPoolSecurityAlertPolicy represents the Terraform resource azurerm_synapse_sql_pool_security_alert_policy.
type SynapseSqlPoolSecurityAlertPolicy struct {
	Name      string
	Args      SynapseSqlPoolSecurityAlertPolicyArgs
	state     *synapseSqlPoolSecurityAlertPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) Type() string {
	return "azurerm_synapse_sql_pool_security_alert_policy"
}

// LocalName returns the local name for [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) LocalName() string {
	return sspsap.Name
}

// Configuration returns the configuration (args) for [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) Configuration() interface{} {
	return sspsap.Args
}

// DependOn is used for other resources to depend on [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sspsap)
}

// Dependencies returns the list of resources [SynapseSqlPoolSecurityAlertPolicy] depends_on.
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) Dependencies() terra.Dependencies {
	return sspsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) LifecycleManagement() *terra.Lifecycle {
	return sspsap.Lifecycle
}

// Attributes returns the attributes for [SynapseSqlPoolSecurityAlertPolicy].
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) Attributes() synapseSqlPoolSecurityAlertPolicyAttributes {
	return synapseSqlPoolSecurityAlertPolicyAttributes{ref: terra.ReferenceResource(sspsap)}
}

// ImportState imports the given attribute values into [SynapseSqlPoolSecurityAlertPolicy]'s state.
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) ImportState(av io.Reader) error {
	sspsap.state = &synapseSqlPoolSecurityAlertPolicyState{}
	if err := json.NewDecoder(av).Decode(sspsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sspsap.Type(), sspsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseSqlPoolSecurityAlertPolicy] has state.
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) State() (*synapseSqlPoolSecurityAlertPolicyState, bool) {
	return sspsap.state, sspsap.state != nil
}

// StateMust returns the state for [SynapseSqlPoolSecurityAlertPolicy]. Panics if the state is nil.
func (sspsap *SynapseSqlPoolSecurityAlertPolicy) StateMust() *synapseSqlPoolSecurityAlertPolicyState {
	if sspsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sspsap.Type(), sspsap.LocalName()))
	}
	return sspsap.state
}

// SynapseSqlPoolSecurityAlertPolicyArgs contains the configurations for azurerm_synapse_sql_pool_security_alert_policy.
type SynapseSqlPoolSecurityAlertPolicyArgs struct {
	// DisabledAlerts: set of string, optional
	DisabledAlerts terra.SetValue[terra.StringValue] `hcl:"disabled_alerts,attr"`
	// EmailAccountAdminsEnabled: bool, optional
	EmailAccountAdminsEnabled terra.BoolValue `hcl:"email_account_admins_enabled,attr"`
	// EmailAddresses: set of string, optional
	EmailAddresses terra.SetValue[terra.StringValue] `hcl:"email_addresses,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PolicyState: string, required
	PolicyState terra.StringValue `hcl:"policy_state,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// SqlPoolId: string, required
	SqlPoolId terra.StringValue `hcl:"sql_pool_id,attr" validate:"required"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// Timeouts: optional
	Timeouts *synapsesqlpoolsecurityalertpolicy.Timeouts `hcl:"timeouts,block"`
}
type synapseSqlPoolSecurityAlertPolicyAttributes struct {
	ref terra.Reference
}

// DisabledAlerts returns a reference to field disabled_alerts of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sspsap.ref.Append("disabled_alerts"))
}

// EmailAccountAdminsEnabled returns a reference to field email_account_admins_enabled of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) EmailAccountAdminsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sspsap.ref.Append("email_account_admins_enabled"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](sspsap.ref.Append("email_addresses"))
}

// Id returns a reference to field id of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sspsap.ref.Append("id"))
}

// PolicyState returns a reference to field policy_state of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) PolicyState() terra.StringValue {
	return terra.ReferenceAsString(sspsap.ref.Append("policy_state"))
}

// RetentionDays returns a reference to field retention_days of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sspsap.ref.Append("retention_days"))
}

// SqlPoolId returns a reference to field sql_pool_id of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) SqlPoolId() terra.StringValue {
	return terra.ReferenceAsString(sspsap.ref.Append("sql_pool_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sspsap.ref.Append("storage_account_access_key"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_synapse_sql_pool_security_alert_policy.
func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sspsap.ref.Append("storage_endpoint"))
}

func (sspsap synapseSqlPoolSecurityAlertPolicyAttributes) Timeouts() synapsesqlpoolsecurityalertpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapsesqlpoolsecurityalertpolicy.TimeoutsAttributes](sspsap.ref.Append("timeouts"))
}

type synapseSqlPoolSecurityAlertPolicyState struct {
	DisabledAlerts            []string                                         `json:"disabled_alerts"`
	EmailAccountAdminsEnabled bool                                             `json:"email_account_admins_enabled"`
	EmailAddresses            []string                                         `json:"email_addresses"`
	Id                        string                                           `json:"id"`
	PolicyState               string                                           `json:"policy_state"`
	RetentionDays             float64                                          `json:"retention_days"`
	SqlPoolId                 string                                           `json:"sql_pool_id"`
	StorageAccountAccessKey   string                                           `json:"storage_account_access_key"`
	StorageEndpoint           string                                           `json:"storage_endpoint"`
	Timeouts                  *synapsesqlpoolsecurityalertpolicy.TimeoutsState `json:"timeouts"`
}
