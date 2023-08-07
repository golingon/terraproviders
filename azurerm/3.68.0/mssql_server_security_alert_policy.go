// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlserversecurityalertpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/mssqlserversecurityalertpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlServerSecurityAlertPolicy creates a new instance of [MssqlServerSecurityAlertPolicy].
func NewMssqlServerSecurityAlertPolicy(name string, args MssqlServerSecurityAlertPolicyArgs) *MssqlServerSecurityAlertPolicy {
	return &MssqlServerSecurityAlertPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServerSecurityAlertPolicy)(nil)

// MssqlServerSecurityAlertPolicy represents the Terraform resource azurerm_mssql_server_security_alert_policy.
type MssqlServerSecurityAlertPolicy struct {
	Name      string
	Args      MssqlServerSecurityAlertPolicyArgs
	state     *mssqlServerSecurityAlertPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) Type() string {
	return "azurerm_mssql_server_security_alert_policy"
}

// LocalName returns the local name for [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) LocalName() string {
	return mssap.Name
}

// Configuration returns the configuration (args) for [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) Configuration() interface{} {
	return mssap.Args
}

// DependOn is used for other resources to depend on [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(mssap)
}

// Dependencies returns the list of resources [MssqlServerSecurityAlertPolicy] depends_on.
func (mssap *MssqlServerSecurityAlertPolicy) Dependencies() terra.Dependencies {
	return mssap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) LifecycleManagement() *terra.Lifecycle {
	return mssap.Lifecycle
}

// Attributes returns the attributes for [MssqlServerSecurityAlertPolicy].
func (mssap *MssqlServerSecurityAlertPolicy) Attributes() mssqlServerSecurityAlertPolicyAttributes {
	return mssqlServerSecurityAlertPolicyAttributes{ref: terra.ReferenceResource(mssap)}
}

// ImportState imports the given attribute values into [MssqlServerSecurityAlertPolicy]'s state.
func (mssap *MssqlServerSecurityAlertPolicy) ImportState(av io.Reader) error {
	mssap.state = &mssqlServerSecurityAlertPolicyState{}
	if err := json.NewDecoder(av).Decode(mssap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mssap.Type(), mssap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServerSecurityAlertPolicy] has state.
func (mssap *MssqlServerSecurityAlertPolicy) State() (*mssqlServerSecurityAlertPolicyState, bool) {
	return mssap.state, mssap.state != nil
}

// StateMust returns the state for [MssqlServerSecurityAlertPolicy]. Panics if the state is nil.
func (mssap *MssqlServerSecurityAlertPolicy) StateMust() *mssqlServerSecurityAlertPolicyState {
	if mssap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mssap.Type(), mssap.LocalName()))
	}
	return mssap.state
}

// MssqlServerSecurityAlertPolicyArgs contains the configurations for azurerm_mssql_server_security_alert_policy.
type MssqlServerSecurityAlertPolicyArgs struct {
	// DisabledAlerts: set of string, optional
	DisabledAlerts terra.SetValue[terra.StringValue] `hcl:"disabled_alerts,attr"`
	// EmailAccountAdmins: bool, optional
	EmailAccountAdmins terra.BoolValue `hcl:"email_account_admins,attr"`
	// EmailAddresses: set of string, optional
	EmailAddresses terra.SetValue[terra.StringValue] `hcl:"email_addresses,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RetentionDays: number, optional
	RetentionDays terra.NumberValue `hcl:"retention_days,attr"`
	// ServerName: string, required
	ServerName terra.StringValue `hcl:"server_name,attr" validate:"required"`
	// State: string, required
	State terra.StringValue `hcl:"state,attr" validate:"required"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// Timeouts: optional
	Timeouts *mssqlserversecurityalertpolicy.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerSecurityAlertPolicyAttributes struct {
	ref terra.Reference
}

// DisabledAlerts returns a reference to field disabled_alerts of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mssap.ref.Append("disabled_alerts"))
}

// EmailAccountAdmins returns a reference to field email_account_admins of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) EmailAccountAdmins() terra.BoolValue {
	return terra.ReferenceAsBool(mssap.ref.Append("email_account_admins"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](mssap.ref.Append("email_addresses"))
}

// Id returns a reference to field id of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("resource_group_name"))
}

// RetentionDays returns a reference to field retention_days of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(mssap.ref.Append("retention_days"))
}

// ServerName returns a reference to field server_name of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) ServerName() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("server_name"))
}

// State returns a reference to field state of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("state"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("storage_account_access_key"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_mssql_server_security_alert_policy.
func (mssap mssqlServerSecurityAlertPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(mssap.ref.Append("storage_endpoint"))
}

func (mssap mssqlServerSecurityAlertPolicyAttributes) Timeouts() mssqlserversecurityalertpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlserversecurityalertpolicy.TimeoutsAttributes](mssap.ref.Append("timeouts"))
}

type mssqlServerSecurityAlertPolicyState struct {
	DisabledAlerts          []string                                      `json:"disabled_alerts"`
	EmailAccountAdmins      bool                                          `json:"email_account_admins"`
	EmailAddresses          []string                                      `json:"email_addresses"`
	Id                      string                                        `json:"id"`
	ResourceGroupName       string                                        `json:"resource_group_name"`
	RetentionDays           float64                                       `json:"retention_days"`
	ServerName              string                                        `json:"server_name"`
	State                   string                                        `json:"state"`
	StorageAccountAccessKey string                                        `json:"storage_account_access_key"`
	StorageEndpoint         string                                        `json:"storage_endpoint"`
	Timeouts                *mssqlserversecurityalertpolicy.TimeoutsState `json:"timeouts"`
}
