// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseworkspaceextendedauditingpolicy "github.com/golingon/terraproviders/azurerm/3.66.0/synapseworkspaceextendedauditingpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseWorkspaceExtendedAuditingPolicy creates a new instance of [SynapseWorkspaceExtendedAuditingPolicy].
func NewSynapseWorkspaceExtendedAuditingPolicy(name string, args SynapseWorkspaceExtendedAuditingPolicyArgs) *SynapseWorkspaceExtendedAuditingPolicy {
	return &SynapseWorkspaceExtendedAuditingPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseWorkspaceExtendedAuditingPolicy)(nil)

// SynapseWorkspaceExtendedAuditingPolicy represents the Terraform resource azurerm_synapse_workspace_extended_auditing_policy.
type SynapseWorkspaceExtendedAuditingPolicy struct {
	Name      string
	Args      SynapseWorkspaceExtendedAuditingPolicyArgs
	state     *synapseWorkspaceExtendedAuditingPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) Type() string {
	return "azurerm_synapse_workspace_extended_auditing_policy"
}

// LocalName returns the local name for [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) LocalName() string {
	return sweap.Name
}

// Configuration returns the configuration (args) for [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) Configuration() interface{} {
	return sweap.Args
}

// DependOn is used for other resources to depend on [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(sweap)
}

// Dependencies returns the list of resources [SynapseWorkspaceExtendedAuditingPolicy] depends_on.
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) Dependencies() terra.Dependencies {
	return sweap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) LifecycleManagement() *terra.Lifecycle {
	return sweap.Lifecycle
}

// Attributes returns the attributes for [SynapseWorkspaceExtendedAuditingPolicy].
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) Attributes() synapseWorkspaceExtendedAuditingPolicyAttributes {
	return synapseWorkspaceExtendedAuditingPolicyAttributes{ref: terra.ReferenceResource(sweap)}
}

// ImportState imports the given attribute values into [SynapseWorkspaceExtendedAuditingPolicy]'s state.
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) ImportState(av io.Reader) error {
	sweap.state = &synapseWorkspaceExtendedAuditingPolicyState{}
	if err := json.NewDecoder(av).Decode(sweap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sweap.Type(), sweap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseWorkspaceExtendedAuditingPolicy] has state.
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) State() (*synapseWorkspaceExtendedAuditingPolicyState, bool) {
	return sweap.state, sweap.state != nil
}

// StateMust returns the state for [SynapseWorkspaceExtendedAuditingPolicy]. Panics if the state is nil.
func (sweap *SynapseWorkspaceExtendedAuditingPolicy) StateMust() *synapseWorkspaceExtendedAuditingPolicyState {
	if sweap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sweap.Type(), sweap.LocalName()))
	}
	return sweap.state
}

// SynapseWorkspaceExtendedAuditingPolicyArgs contains the configurations for azurerm_synapse_workspace_extended_auditing_policy.
type SynapseWorkspaceExtendedAuditingPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogMonitoringEnabled: bool, optional
	LogMonitoringEnabled terra.BoolValue `hcl:"log_monitoring_enabled,attr"`
	// RetentionInDays: number, optional
	RetentionInDays terra.NumberValue `hcl:"retention_in_days,attr"`
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageAccountAccessKeyIsSecondary: bool, optional
	StorageAccountAccessKeyIsSecondary terra.BoolValue `hcl:"storage_account_access_key_is_secondary,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *synapseworkspaceextendedauditingpolicy.Timeouts `hcl:"timeouts,block"`
}
type synapseWorkspaceExtendedAuditingPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sweap.ref.Append("id"))
}

// LogMonitoringEnabled returns a reference to field log_monitoring_enabled of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) LogMonitoringEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sweap.ref.Append("log_monitoring_enabled"))
}

// RetentionInDays returns a reference to field retention_in_days of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) RetentionInDays() terra.NumberValue {
	return terra.ReferenceAsNumber(sweap.ref.Append("retention_in_days"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(sweap.ref.Append("storage_account_access_key"))
}

// StorageAccountAccessKeyIsSecondary returns a reference to field storage_account_access_key_is_secondary of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) StorageAccountAccessKeyIsSecondary() terra.BoolValue {
	return terra.ReferenceAsBool(sweap.ref.Append("storage_account_access_key_is_secondary"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(sweap.ref.Append("storage_endpoint"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_workspace_extended_auditing_policy.
func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(sweap.ref.Append("synapse_workspace_id"))
}

func (sweap synapseWorkspaceExtendedAuditingPolicyAttributes) Timeouts() synapseworkspaceextendedauditingpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseworkspaceextendedauditingpolicy.TimeoutsAttributes](sweap.ref.Append("timeouts"))
}

type synapseWorkspaceExtendedAuditingPolicyState struct {
	Id                                 string                                                `json:"id"`
	LogMonitoringEnabled               bool                                                  `json:"log_monitoring_enabled"`
	RetentionInDays                    float64                                               `json:"retention_in_days"`
	StorageAccountAccessKey            string                                                `json:"storage_account_access_key"`
	StorageAccountAccessKeyIsSecondary bool                                                  `json:"storage_account_access_key_is_secondary"`
	StorageEndpoint                    string                                                `json:"storage_endpoint"`
	SynapseWorkspaceId                 string                                                `json:"synapse_workspace_id"`
	Timeouts                           *synapseworkspaceextendedauditingpolicy.TimeoutsState `json:"timeouts"`
}
