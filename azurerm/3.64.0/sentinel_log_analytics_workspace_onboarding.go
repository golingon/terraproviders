// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	sentinelloganalyticsworkspaceonboarding "github.com/golingon/terraproviders/azurerm/3.64.0/sentinelloganalyticsworkspaceonboarding"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSentinelLogAnalyticsWorkspaceOnboarding creates a new instance of [SentinelLogAnalyticsWorkspaceOnboarding].
func NewSentinelLogAnalyticsWorkspaceOnboarding(name string, args SentinelLogAnalyticsWorkspaceOnboardingArgs) *SentinelLogAnalyticsWorkspaceOnboarding {
	return &SentinelLogAnalyticsWorkspaceOnboarding{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SentinelLogAnalyticsWorkspaceOnboarding)(nil)

// SentinelLogAnalyticsWorkspaceOnboarding represents the Terraform resource azurerm_sentinel_log_analytics_workspace_onboarding.
type SentinelLogAnalyticsWorkspaceOnboarding struct {
	Name      string
	Args      SentinelLogAnalyticsWorkspaceOnboardingArgs
	state     *sentinelLogAnalyticsWorkspaceOnboardingState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) Type() string {
	return "azurerm_sentinel_log_analytics_workspace_onboarding"
}

// LocalName returns the local name for [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) LocalName() string {
	return slawo.Name
}

// Configuration returns the configuration (args) for [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) Configuration() interface{} {
	return slawo.Args
}

// DependOn is used for other resources to depend on [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) DependOn() terra.Reference {
	return terra.ReferenceResource(slawo)
}

// Dependencies returns the list of resources [SentinelLogAnalyticsWorkspaceOnboarding] depends_on.
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) Dependencies() terra.Dependencies {
	return slawo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) LifecycleManagement() *terra.Lifecycle {
	return slawo.Lifecycle
}

// Attributes returns the attributes for [SentinelLogAnalyticsWorkspaceOnboarding].
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) Attributes() sentinelLogAnalyticsWorkspaceOnboardingAttributes {
	return sentinelLogAnalyticsWorkspaceOnboardingAttributes{ref: terra.ReferenceResource(slawo)}
}

// ImportState imports the given attribute values into [SentinelLogAnalyticsWorkspaceOnboarding]'s state.
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) ImportState(av io.Reader) error {
	slawo.state = &sentinelLogAnalyticsWorkspaceOnboardingState{}
	if err := json.NewDecoder(av).Decode(slawo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", slawo.Type(), slawo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SentinelLogAnalyticsWorkspaceOnboarding] has state.
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) State() (*sentinelLogAnalyticsWorkspaceOnboardingState, bool) {
	return slawo.state, slawo.state != nil
}

// StateMust returns the state for [SentinelLogAnalyticsWorkspaceOnboarding]. Panics if the state is nil.
func (slawo *SentinelLogAnalyticsWorkspaceOnboarding) StateMust() *sentinelLogAnalyticsWorkspaceOnboardingState {
	if slawo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", slawo.Type(), slawo.LocalName()))
	}
	return slawo.state
}

// SentinelLogAnalyticsWorkspaceOnboardingArgs contains the configurations for azurerm_sentinel_log_analytics_workspace_onboarding.
type SentinelLogAnalyticsWorkspaceOnboardingArgs struct {
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
	Timeouts *sentinelloganalyticsworkspaceonboarding.Timeouts `hcl:"timeouts,block"`
}
type sentinelLogAnalyticsWorkspaceOnboardingAttributes struct {
	ref terra.Reference
}

// CustomerManagedKeyEnabled returns a reference to field customer_managed_key_enabled of azurerm_sentinel_log_analytics_workspace_onboarding.
func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) CustomerManagedKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(slawo.ref.Append("customer_managed_key_enabled"))
}

// Id returns a reference to field id of azurerm_sentinel_log_analytics_workspace_onboarding.
func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(slawo.ref.Append("id"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_sentinel_log_analytics_workspace_onboarding.
func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(slawo.ref.Append("resource_group_name"))
}

// WorkspaceId returns a reference to field workspace_id of azurerm_sentinel_log_analytics_workspace_onboarding.
func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) WorkspaceId() terra.StringValue {
	return terra.ReferenceAsString(slawo.ref.Append("workspace_id"))
}

// WorkspaceName returns a reference to field workspace_name of azurerm_sentinel_log_analytics_workspace_onboarding.
func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) WorkspaceName() terra.StringValue {
	return terra.ReferenceAsString(slawo.ref.Append("workspace_name"))
}

func (slawo sentinelLogAnalyticsWorkspaceOnboardingAttributes) Timeouts() sentinelloganalyticsworkspaceonboarding.TimeoutsAttributes {
	return terra.ReferenceAsSingle[sentinelloganalyticsworkspaceonboarding.TimeoutsAttributes](slawo.ref.Append("timeouts"))
}

type sentinelLogAnalyticsWorkspaceOnboardingState struct {
	CustomerManagedKeyEnabled bool                                                   `json:"customer_managed_key_enabled"`
	Id                        string                                                 `json:"id"`
	ResourceGroupName         string                                                 `json:"resource_group_name"`
	WorkspaceId               string                                                 `json:"workspace_id"`
	WorkspaceName             string                                                 `json:"workspace_name"`
	Timeouts                  *sentinelloganalyticsworkspaceonboarding.TimeoutsState `json:"timeouts"`
}
