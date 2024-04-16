// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_synapse_sql_pool_security_alert_policy

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource azurerm_synapse_sql_pool_security_alert_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSynapseSqlPoolSecurityAlertPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asspsap *Resource) Type() string {
	return "azurerm_synapse_sql_pool_security_alert_policy"
}

// LocalName returns the local name for [Resource].
func (asspsap *Resource) LocalName() string {
	return asspsap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asspsap *Resource) Configuration() interface{} {
	return asspsap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asspsap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asspsap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asspsap *Resource) Dependencies() terra.Dependencies {
	return asspsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asspsap *Resource) LifecycleManagement() *terra.Lifecycle {
	return asspsap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asspsap *Resource) Attributes() azurermSynapseSqlPoolSecurityAlertPolicyAttributes {
	return azurermSynapseSqlPoolSecurityAlertPolicyAttributes{ref: terra.ReferenceResource(asspsap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asspsap *Resource) ImportState(state io.Reader) error {
	asspsap.state = &azurermSynapseSqlPoolSecurityAlertPolicyState{}
	if err := json.NewDecoder(state).Decode(asspsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asspsap.Type(), asspsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asspsap *Resource) State() (*azurermSynapseSqlPoolSecurityAlertPolicyState, bool) {
	return asspsap.state, asspsap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asspsap *Resource) StateMust() *azurermSynapseSqlPoolSecurityAlertPolicyState {
	if asspsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asspsap.Type(), asspsap.LocalName()))
	}
	return asspsap.state
}

// Args contains the configurations for azurerm_synapse_sql_pool_security_alert_policy.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSynapseSqlPoolSecurityAlertPolicyAttributes struct {
	ref terra.Reference
}

// DisabledAlerts returns a reference to field disabled_alerts of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](asspsap.ref.Append("disabled_alerts"))
}

// EmailAccountAdminsEnabled returns a reference to field email_account_admins_enabled of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) EmailAccountAdminsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asspsap.ref.Append("email_account_admins_enabled"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](asspsap.ref.Append("email_addresses"))
}

// Id returns a reference to field id of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asspsap.ref.Append("id"))
}

// PolicyState returns a reference to field policy_state of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) PolicyState() terra.StringValue {
	return terra.ReferenceAsString(asspsap.ref.Append("policy_state"))
}

// RetentionDays returns a reference to field retention_days of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(asspsap.ref.Append("retention_days"))
}

// SqlPoolId returns a reference to field sql_pool_id of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) SqlPoolId() terra.StringValue {
	return terra.ReferenceAsString(asspsap.ref.Append("sql_pool_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(asspsap.ref.Append("storage_account_access_key"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_synapse_sql_pool_security_alert_policy.
func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(asspsap.ref.Append("storage_endpoint"))
}

func (asspsap azurermSynapseSqlPoolSecurityAlertPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asspsap.ref.Append("timeouts"))
}

type azurermSynapseSqlPoolSecurityAlertPolicyState struct {
	DisabledAlerts            []string       `json:"disabled_alerts"`
	EmailAccountAdminsEnabled bool           `json:"email_account_admins_enabled"`
	EmailAddresses            []string       `json:"email_addresses"`
	Id                        string         `json:"id"`
	PolicyState               string         `json:"policy_state"`
	RetentionDays             float64        `json:"retention_days"`
	SqlPoolId                 string         `json:"sql_pool_id"`
	StorageAccountAccessKey   string         `json:"storage_account_access_key"`
	StorageEndpoint           string         `json:"storage_endpoint"`
	Timeouts                  *TimeoutsState `json:"timeouts"`
}
