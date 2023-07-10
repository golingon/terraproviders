// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	synapseworkspacesecurityalertpolicy "github.com/golingon/terraproviders/azurerm/3.64.0/synapseworkspacesecurityalertpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSynapseWorkspaceSecurityAlertPolicy creates a new instance of [SynapseWorkspaceSecurityAlertPolicy].
func NewSynapseWorkspaceSecurityAlertPolicy(name string, args SynapseWorkspaceSecurityAlertPolicyArgs) *SynapseWorkspaceSecurityAlertPolicy {
	return &SynapseWorkspaceSecurityAlertPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SynapseWorkspaceSecurityAlertPolicy)(nil)

// SynapseWorkspaceSecurityAlertPolicy represents the Terraform resource azurerm_synapse_workspace_security_alert_policy.
type SynapseWorkspaceSecurityAlertPolicy struct {
	Name      string
	Args      SynapseWorkspaceSecurityAlertPolicyArgs
	state     *synapseWorkspaceSecurityAlertPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) Type() string {
	return "azurerm_synapse_workspace_security_alert_policy"
}

// LocalName returns the local name for [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) LocalName() string {
	return swsap.Name
}

// Configuration returns the configuration (args) for [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) Configuration() interface{} {
	return swsap.Args
}

// DependOn is used for other resources to depend on [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(swsap)
}

// Dependencies returns the list of resources [SynapseWorkspaceSecurityAlertPolicy] depends_on.
func (swsap *SynapseWorkspaceSecurityAlertPolicy) Dependencies() terra.Dependencies {
	return swsap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) LifecycleManagement() *terra.Lifecycle {
	return swsap.Lifecycle
}

// Attributes returns the attributes for [SynapseWorkspaceSecurityAlertPolicy].
func (swsap *SynapseWorkspaceSecurityAlertPolicy) Attributes() synapseWorkspaceSecurityAlertPolicyAttributes {
	return synapseWorkspaceSecurityAlertPolicyAttributes{ref: terra.ReferenceResource(swsap)}
}

// ImportState imports the given attribute values into [SynapseWorkspaceSecurityAlertPolicy]'s state.
func (swsap *SynapseWorkspaceSecurityAlertPolicy) ImportState(av io.Reader) error {
	swsap.state = &synapseWorkspaceSecurityAlertPolicyState{}
	if err := json.NewDecoder(av).Decode(swsap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", swsap.Type(), swsap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SynapseWorkspaceSecurityAlertPolicy] has state.
func (swsap *SynapseWorkspaceSecurityAlertPolicy) State() (*synapseWorkspaceSecurityAlertPolicyState, bool) {
	return swsap.state, swsap.state != nil
}

// StateMust returns the state for [SynapseWorkspaceSecurityAlertPolicy]. Panics if the state is nil.
func (swsap *SynapseWorkspaceSecurityAlertPolicy) StateMust() *synapseWorkspaceSecurityAlertPolicyState {
	if swsap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", swsap.Type(), swsap.LocalName()))
	}
	return swsap.state
}

// SynapseWorkspaceSecurityAlertPolicyArgs contains the configurations for azurerm_synapse_workspace_security_alert_policy.
type SynapseWorkspaceSecurityAlertPolicyArgs struct {
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
	// StorageAccountAccessKey: string, optional
	StorageAccountAccessKey terra.StringValue `hcl:"storage_account_access_key,attr"`
	// StorageEndpoint: string, optional
	StorageEndpoint terra.StringValue `hcl:"storage_endpoint,attr"`
	// SynapseWorkspaceId: string, required
	SynapseWorkspaceId terra.StringValue `hcl:"synapse_workspace_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *synapseworkspacesecurityalertpolicy.Timeouts `hcl:"timeouts,block"`
}
type synapseWorkspaceSecurityAlertPolicyAttributes struct {
	ref terra.Reference
}

// DisabledAlerts returns a reference to field disabled_alerts of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) DisabledAlerts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](swsap.ref.Append("disabled_alerts"))
}

// EmailAccountAdminsEnabled returns a reference to field email_account_admins_enabled of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) EmailAccountAdminsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(swsap.ref.Append("email_account_admins_enabled"))
}

// EmailAddresses returns a reference to field email_addresses of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) EmailAddresses() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](swsap.ref.Append("email_addresses"))
}

// Id returns a reference to field id of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(swsap.ref.Append("id"))
}

// PolicyState returns a reference to field policy_state of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) PolicyState() terra.StringValue {
	return terra.ReferenceAsString(swsap.ref.Append("policy_state"))
}

// RetentionDays returns a reference to field retention_days of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) RetentionDays() terra.NumberValue {
	return terra.ReferenceAsNumber(swsap.ref.Append("retention_days"))
}

// StorageAccountAccessKey returns a reference to field storage_account_access_key of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) StorageAccountAccessKey() terra.StringValue {
	return terra.ReferenceAsString(swsap.ref.Append("storage_account_access_key"))
}

// StorageEndpoint returns a reference to field storage_endpoint of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) StorageEndpoint() terra.StringValue {
	return terra.ReferenceAsString(swsap.ref.Append("storage_endpoint"))
}

// SynapseWorkspaceId returns a reference to field synapse_workspace_id of azurerm_synapse_workspace_security_alert_policy.
func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) SynapseWorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(swsap.ref.Append("synapse_workspace_id"))
}

func (swsap synapseWorkspaceSecurityAlertPolicyAttributes) Timeouts() synapseworkspacesecurityalertpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[synapseworkspacesecurityalertpolicy.TimeoutsAttributes](swsap.ref.Append("timeouts"))
}

type synapseWorkspaceSecurityAlertPolicyState struct {
	DisabledAlerts            []string                                           `json:"disabled_alerts"`
	EmailAccountAdminsEnabled bool                                               `json:"email_account_admins_enabled"`
	EmailAddresses            []string                                           `json:"email_addresses"`
	Id                        string                                             `json:"id"`
	PolicyState               string                                             `json:"policy_state"`
	RetentionDays             float64                                            `json:"retention_days"`
	StorageAccountAccessKey   string                                             `json:"storage_account_access_key"`
	StorageEndpoint           string                                             `json:"storage_endpoint"`
	SynapseWorkspaceId        string                                             `json:"synapse_workspace_id"`
	Timeouts                  *synapseworkspacesecurityalertpolicy.TimeoutsState `json:"timeouts"`
}
