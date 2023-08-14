// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	mssqlservermicrosoftsupportauditingpolicy "github.com/golingon/terraproviders/azurerm/3.69.0/mssqlservermicrosoftsupportauditingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMssqlServerMicrosoftSupportAuditingPolicy creates a new instance of [MssqlServerMicrosoftSupportAuditingPolicy].
func NewMssqlServerMicrosoftSupportAuditingPolicy(name string, args MssqlServerMicrosoftSupportAuditingPolicyArgs) *MssqlServerMicrosoftSupportAuditingPolicy {
	return &MssqlServerMicrosoftSupportAuditingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*MssqlServerMicrosoftSupportAuditingPolicy)(nil)

// MssqlServerMicrosoftSupportAuditingPolicy represents the Terraform resource azurerm_mssql_server_microsoft_support_auditing_policy.
type MssqlServerMicrosoftSupportAuditingPolicy struct {
	Name      string
	Args      MssqlServerMicrosoftSupportAuditingPolicyArgs
	state     *mssqlServerMicrosoftSupportAuditingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) Type() string {
	return "azurerm_mssql_server_microsoft_support_auditing_policy"
}

// LocalName returns the local name for [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) LocalName() string {
	return msmsap.Name
}

// Configuration returns the configuration (args) for [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) Configuration() interface{} {
	return msmsap.Args
}

// DependOn is used for other resources to depend on [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(msmsap)
}

// Dependencies returns the list of resources [MssqlServerMicrosoftSupportAuditingPolicy] depends_on.
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) Dependencies() terra.Dependencies {
	return msmsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) LifecycleManagement() *terra.Lifecycle {
	return msmsap.Lifecycle
}

// Attributes returns the attributes for [MssqlServerMicrosoftSupportAuditingPolicy].
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) Attributes() mssqlServerMicrosoftSupportAuditingPolicyAttributes {
	return mssqlServerMicrosoftSupportAuditingPolicyAttributes{ref: terra.ReferenceResource(msmsap)}
}

// ImportState imports the given attribute values into [MssqlServerMicrosoftSupportAuditingPolicy]'s state.
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) ImportState(av io.Reader) error {
	msmsap.state = &mssqlServerMicrosoftSupportAuditingPolicyState{}
	if err := json.NewDecoder(av).Decode(msmsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", msmsap.Type(), msmsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [MssqlServerMicrosoftSupportAuditingPolicy] has state.
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) State() (*mssqlServerMicrosoftSupportAuditingPolicyState, bool) {
	return msmsap.state, msmsap.state != nil
}

// StateMust returns the state for [MssqlServerMicrosoftSupportAuditingPolicy]. Panics if the state is nil.
func (msmsap *MssqlServerMicrosoftSupportAuditingPolicy) StateMust() *mssqlServerMicrosoftSupportAuditingPolicyState {
	if msmsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", msmsap.Type(), msmsap.LocalName()))
	}
	return msmsap.state
}

// MssqlServerMicrosoftSupportAuditingPolicyArgs contains the configurations for azurerm_mssql_server_microsoft_support_auditing_policy.
type MssqlServerMicrosoftSupportAuditingPolicyArgs struct {
	// BlobStorageEndpoint: string, optional
	BlobStorageEndpoint terra.StringValue `hcl:"blob_storage_endpoint,attr"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogMonitoringEnabled: bool, optional
	LogMonitoringEnabled terra.BoolValue `hcl:"log_monitoring_enabled,attr"`
	// ServerId: string, required
	ServerId terra.StringValue `hcl:"server_id,attr" validate:"required"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageAccountSubscriptionId: string, optional
	StorageAccountSubscriptionId terra.StringValue `hcl:"storage_account_subscription_id,attr"`
	// Timeouts: optional
	Timeouts *mssqlservermicrosoftsupportauditingpolicy.Timeouts `hcl:"timeouts,block"`
}
type mssqlServerMicrosoftSupportAuditingPolicyAttributes struct {
	ref terra.Reference
}

// BlobStorageEndpoint returns a reference to field blob_storage_endpoint of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) BlobStorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(msmsap.ref.Append("blob_storage_endpoint"))
}

// Enabled returns a reference to field enabled of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(msmsap.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(msmsap.ref.Append("id"))
}

// LogMonitoringEnabled returns a reference to field log_monitoring_enabled of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) LogMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(msmsap.ref.Append("log_monitoring_enabled"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(msmsap.ref.Append("server_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(msmsap.ref.Append("storage_account_access_key"))
}

// StorageAccountSubscriptionId returns a reference to field storage_account_subscription_id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) StorageAccountSubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(msmsap.ref.Append("storage_account_subscription_id"))
}

func (msmsap mssqlServerMicrosoftSupportAuditingPolicyAttributes) Timeouts() mssqlservermicrosoftsupportauditingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[mssqlservermicrosoftsupportauditingpolicy.TimeoutsAttributes](msmsap.ref.Append("timeouts"))
}

type mssqlServerMicrosoftSupportAuditingPolicyState struct {
	BlobStorageEndpoint          string                                                   `json:"blob_storage_endpoint"`
	Enabled                      bool                                                     `json:"enabled"`
	Id                           string                                                   `json:"id"`
	LogMonitoringEnabled         bool                                                     `json:"log_monitoring_enabled"`
	ServerId                     string                                                   `json:"server_id"`
	StorageAccountAccessKey      string                                                   `json:"storage_account_access_key"`
	StorageAccountSubscriptionId string                                                   `json:"storage_account_subscription_id"`
	Timeouts                     *mssqlservermicrosoftsupportauditingpolicy.TimeoutsState `json:"timeouts"`
}
