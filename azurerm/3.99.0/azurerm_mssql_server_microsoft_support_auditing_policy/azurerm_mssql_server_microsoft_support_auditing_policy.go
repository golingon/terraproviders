// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mssql_server_microsoft_support_auditing_policy

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

// Resource represents the Terraform resource azurerm_mssql_server_microsoft_support_auditing_policy.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMssqlServerMicrosoftSupportAuditingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amsmsap *Resource) Type() string {
	return "azurerm_mssql_server_microsoft_support_auditing_policy"
}

// LocalName returns the local name for [Resource].
func (amsmsap *Resource) LocalName() string {
	return amsmsap.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amsmsap *Resource) Configuration() interface{} {
	return amsmsap.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amsmsap *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amsmsap)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amsmsap *Resource) Dependencies() terra.Dependencies {
	return amsmsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amsmsap *Resource) LifecycleManagement() *terra.Lifecycle {
	return amsmsap.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amsmsap *Resource) Attributes() azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes {
	return azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes{ref: terra.ReferenceResource(amsmsap)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amsmsap *Resource) ImportState(state io.Reader) error {
	amsmsap.state = &azurermMssqlServerMicrosoftSupportAuditingPolicyState{}
	if err := json.NewDecoder(state).Decode(amsmsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amsmsap.Type(), amsmsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amsmsap *Resource) State() (*azurermMssqlServerMicrosoftSupportAuditingPolicyState, bool) {
	return amsmsap.state, amsmsap.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amsmsap *Resource) StateMust() *azurermMssqlServerMicrosoftSupportAuditingPolicyState {
	if amsmsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amsmsap.Type(), amsmsap.LocalName()))
	}
	return amsmsap.state
}

// Args contains the configurations for azurerm_mssql_server_microsoft_support_auditing_policy.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes struct {
	ref terra.Reference
}

// BlobStorageEndpoint returns a reference to field blob_storage_endpoint of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) BlobStorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(amsmsap.ref.Append("blob_storage_endpoint"))
}

// Enabled returns a reference to field enabled of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(amsmsap.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amsmsap.ref.Append("id"))
}

// LogMonitoringEnabled returns a reference to field log_monitoring_enabled of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) LogMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(amsmsap.ref.Append("log_monitoring_enabled"))
}

// ServerId returns a reference to field server_id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) ServerId() terra.StringValue {
	return terra.ReferenceAsString(amsmsap.ref.Append("server_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(amsmsap.ref.Append("storage_account_access_key"))
}

// StorageAccountSubscriptionId returns a reference to field storage_account_subscription_id of azurerm_mssql_server_microsoft_support_auditing_policy.
func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) StorageAccountSubscriptionId() terra.StringValue {
	return terra.ReferenceAsString(amsmsap.ref.Append("storage_account_subscription_id"))
}

func (amsmsap azurermMssqlServerMicrosoftSupportAuditingPolicyAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amsmsap.ref.Append("timeouts"))
}

type azurermMssqlServerMicrosoftSupportAuditingPolicyState struct {
	BlobStorageEndpoint          string         `json:"blob_storage_endpoint"`
	Enabled                      bool           `json:"enabled"`
	Id                           string         `json:"id"`
	LogMonitoringEnabled         bool           `json:"log_monitoring_enabled"`
	ServerId                     string         `json:"server_id"`
	StorageAccountAccessKey      string         `json:"storage_account_access_key"`
	StorageAccountSubscriptionId string         `json:"storage_account_subscription_id"`
	Timeouts                     *TimeoutsState `json:"timeouts"`
}
