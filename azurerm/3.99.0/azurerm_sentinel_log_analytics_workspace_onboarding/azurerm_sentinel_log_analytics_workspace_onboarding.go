// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_sentinel_log_analytics_workspace_onboarding

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

// Resource represents the Terraform resource azurerm_sentinel_log_analytics_workspace_onboarding.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSentinelLogAnalyticsWorkspaceOnboardingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aslawo *Resource) Type() string {
	return "azurerm_sentinel_log_analytics_workspace_onboarding"
}

// LocalName returns the local name for [Resource].
func (aslawo *Resource) LocalName() string {
	return aslawo.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aslawo *Resource) Configuration() interface{} {
	return aslawo.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aslawo *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aslawo)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aslawo *Resource) Dependencies() terra.Dependencies {
	return aslawo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aslawo *Resource) LifecycleManagement() *terra.Lifecycle {
	return aslawo.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aslawo *Resource) Attributes() azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes {
	return azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes{ref: terra.ReferenceResource(aslawo)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aslawo *Resource) ImportState(state io.Reader) error {
	aslawo.state = &azurermSentinelLogAnalyticsWorkspaceOnboardingState{}
	if err := json.NewDecoder(state).Decode(aslawo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aslawo.Type(), aslawo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aslawo *Resource) State() (*azurermSentinelLogAnalyticsWorkspaceOnboardingState, bool) {
	return aslawo.state, aslawo.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aslawo *Resource) StateMust() *azurermSentinelLogAnalyticsWorkspaceOnboardingState {
	if aslawo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aslawo.Type(), aslawo.LocalName()))
	}
	return aslawo.state
}

// Args contains the configurations for azurerm_sentinel_log_analytics_workspace_onboarding.
type Args struct {
	// CustomerManagedKeyEnabled: bool, optional
	CustomerManagedKeyEnabled terra.BoolValue `hcl:"customer_managed_key_enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceGroupName: string, optional
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr"`
	// WorkspaceId: string, optional
	WorkspaceId terra.StringValue `hcl:"workspace_id,attr"`
	// WorkspaceName: string, optional
	WorkspaceName terra.StringValue `hcl:"workspace_name,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes struct {
	ref terra.Reference
}

// CustomerManagedKeyEnabled returns a reference to field customer_managed_key_enabled of azurerm_sentinel_log_analytics_workspace_onboarding.
func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) CustomerManagedKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(aslawo.ref.Append("customer_managed_key_enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_log_analytics_workspace_onboarding.
func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aslawo.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sentinel_log_analytics_workspace_onboarding.
func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aslawo.ref.Append("resource_group_name"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_sentinel_log_analytics_workspace_onboarding.
func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(aslawo.ref.Append("workspace_id"))
}

// WorkspaceName returns a reference to field workspace_name of azurerm_sentinel_log_analytics_workspace_onboarding.
func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) WorkspaceName() terra.StringValue {
	return terra.ReferenceAsString(aslawo.ref.Append("workspace_name"))
}

func (aslawo azurermSentinelLogAnalyticsWorkspaceOnboardingAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](aslawo.ref.Append("timeouts"))
}

type azurermSentinelLogAnalyticsWorkspaceOnboardingState struct {
	CustomerManagedKeyEnabled bool           `json:"customer_managed_key_enabled"`
	Id                        string         `json:"id"`
	ResourceGroupName         string         `json:"resource_group_name"`
	WorkspaceId               string         `json:"workspace_id"`
	WorkspaceName             string         `json:"workspace_name"`
	Timeouts                  *TimeoutsState `json:"timeouts"`
}
