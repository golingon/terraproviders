// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapsesqlpoolextendedauditingpolicy "github.com/golingon/terraproviders/azurerm/3.55.0/synapsesqlpoolextendedauditingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseSqlPoolExtendedAuditingPolicy creates a new instance of [SynapseSqlPoolExtendedAuditingPolicy].
func NewSynapseSqlPoolExtendedAuditingPolicy(name string, args SynapseSqlPoolExtendedAuditingPolicyArgs) *SynapseSqlPoolExtendedAuditingPolicy {
	return &SynapseSqlPoolExtendedAuditingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseSqlPoolExtendedAuditingPolicy)(nil)

// SynapseSqlPoolExtendedAuditingPolicy represents the Terraform resource azurerm_synapse_sql_pool_extended_auditing_policy.
type SynapseSqlPoolExtendedAuditingPolicy struct {
	Name      string
	Args      SynapseSqlPoolExtendedAuditingPolicyArgs
	state     *synapseSqlPoolExtendedAuditingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) Type() string {
	return "azurerm_synapse_sql_pool_extended_auditing_policy"
}

// LocalName returns the local name for [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) LocalName() string {
	return sspeap.Name
}

// Configuration returns the configuration (args) for [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) Configuration() interface{} {
	return sspeap.Args
}

// DependOn is used for other resources to depend on [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sspeap)
}

// Dependencies returns the list of resources [SynapseSqlPoolExtendedAuditingPolicy] depends_on.
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) Dependencies() terra.Dependencies {
	return sspeap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) LifecycleManagement() *terra.Lifecycle {
	return sspeap.Lifecycle
}

// Attributes returns the attributes for [SynapseSqlPoolExtendedAuditingPolicy].
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) Attributes() synapseSqlPoolExtendedAuditingPolicyAttributes {
	return synapseSqlPoolExtendedAuditingPolicyAttributes{ref: terra.ReferenceResource(sspeap)}
}

// ImportState imports the given attribute values into [SynapseSqlPoolExtendedAuditingPolicy]'s state.
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) ImportState(av io.Reader) error {
	sspeap.state = &synapseSqlPoolExtendedAuditingPolicyState{}
	if err := json.NewDecoder(av).Decode(sspeap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sspeap.Type(), sspeap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseSqlPoolExtendedAuditingPolicy] has state.
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) State() (*synapseSqlPoolExtendedAuditingPolicyState, bool) {
	return sspeap.state, sspeap.state != nil
}

// StateMust returns the state for [SynapseSqlPoolExtendedAuditingPolicy]. Panics if the state is nil.
func (sspeap *SynapseSqlPoolExtendedAuditingPolicy) StateMust() *synapseSqlPoolExtendedAuditingPolicyState {
	if sspeap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sspeap.Type(), sspeap.LocalName()))
	}
	return sspeap.state
}

// SynapseSqlPoolExtendedAuditingPolicyArgs contains the configurations for azurerm_synapse_sql_pool_extended_auditing_policy.
type SynapseSqlPoolExtendedAuditingPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogMonitoringEnabled: bool, optional
	LogMonitoringEnabled terra.BoolValue `hcl:"log_monitoring_enabled,attr"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// SqlPoolId: string, required
	SqlPoolId terra.StringValue `hcl:"sql_pool_id,attr" validate:"required"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageAccountAccessKeyIsSecondary: bool, optional
	StorageAccountAccessKeyIsSecondary terra.BoolValue `hcl:"storage_account_access_key_is_secondary,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// Timeouts: optional
	Timeouts *synapsesqlpoolextendedauditingpolicy.Timeouts `hcl:"timeouts,block"`
}
type synapseSqlPoolExtendedAuditingPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sspeap.ref.Append("id"))
}

// LogMonitoringEnabled returns a reference to field log_monitoring_enabled of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) LogMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sspeap.ref.Append("log_monitoring_enabled"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sspeap.ref.Append("retention_in_days"))
}

// SqlPoolId returns a reference to field sql_pool_id of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) SqlPoolId() terra.StringValue {
	return terra.ReferenceAsString(sspeap.ref.Append("sql_pool_id"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sspeap.ref.Append("storage_account_access_key"))
}

// StorageAccountAccessKeyIsSecondary returns a reference to field storage_account_access_key_is_secondary of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) StorageAccountAccessKeyIsSecondary() terra.BoolValue {
	return terra.ReferenceAsBool(sspeap.ref.Append("storage_account_access_key_is_secondary"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_synapse_sql_pool_extended_auditing_policy.
func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sspeap.ref.Append("storage_endpoint"))
}

func (sspeap synapseSqlPoolExtendedAuditingPolicyAttributes) Timeouts() synapsesqlpoolextendedauditingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapsesqlpoolextendedauditingpolicy.TimeoutsAttributes](sspeap.ref.Append("timeouts"))
}

type synapseSqlPoolExtendedAuditingPolicyState struct {
	Id                                 string                                              `json:"id"`
	LogMonitoringEnabled               bool                                                `json:"log_monitoring_enabled"`
	RetentionInDays                    float64                                             `json:"retention_in_days"`
	SqlPoolId                          string                                              `json:"sql_pool_id"`
	StorageAccountAccessKey            string                                              `json:"storage_account_access_key"`
	StorageAccountAccessKeyIsSecondary bool                                                `json:"storage_account_access_key_is_secondary"`
	StorageEndpoint                    string                                              `json:"storage_endpoint"`
	Timeouts                           *synapsesqlpoolextendedauditingpolicy.TimeoutsState `json:"timeouts"`
}
