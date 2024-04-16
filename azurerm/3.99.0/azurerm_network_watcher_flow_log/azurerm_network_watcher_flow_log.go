// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_network_watcher_flow_log

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

// Resource represents the Terraform resource azurerm_network_watcher_flow_log.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermNetworkWatcherFlowLogState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (anwfl *Resource) Type() string {
	return "azurerm_network_watcher_flow_log"
}

// LocalName returns the local name for [Resource].
func (anwfl *Resource) LocalName() string {
	return anwfl.Name
}

// Configuration returns the configuration (args) for [Resource].
func (anwfl *Resource) Configuration() interface{} {
	return anwfl.Args
}

// DependOn is used for other resources to depend on [Resource].
func (anwfl *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(anwfl)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (anwfl *Resource) Dependencies() terra.Dependencies {
	return anwfl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (anwfl *Resource) LifecycleManagement() *terra.Lifecycle {
	return anwfl.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (anwfl *Resource) Attributes() azurermNetworkWatcherFlowLogAttributes {
	return azurermNetworkWatcherFlowLogAttributes{ref: terra.ReferenceResource(anwfl)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (anwfl *Resource) ImportState(state io.Reader) error {
	anwfl.state = &azurermNetworkWatcherFlowLogState{}
	if err := json.NewDecoder(state).Decode(anwfl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", anwfl.Type(), anwfl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (anwfl *Resource) State() (*azurermNetworkWatcherFlowLogState, bool) {
	return anwfl.state, anwfl.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (anwfl *Resource) StateMust() *azurermNetworkWatcherFlowLogState {
	if anwfl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", anwfl.Type(), anwfl.LocalName()))
	}
	return anwfl.state
}

// Args contains the configurations for azurerm_network_watcher_flow_log.
type Args struct {
	// Enabled: bool, required
	Enabled terra.BoolValue `hcl:"enabled,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkSecurityGroupId: string, required
	NetworkSecurityGroupId terra.StringValue `hcl:"network_security_group_id,attr" validate:"required"`
	// NetworkWatcherName: string, required
	NetworkWatcherName terra.StringValue `hcl:"network_watcher_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StorageAccountId: string, required
	StorageAccountId terra.StringValue `hcl:"storage_account_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Version: number, optional
	Version terra.NumberValue `hcl:"version,attr"`
	// RetentionPolicy: required
	RetentionPolicy *RetentionPolicy `hcl:"retention_policy,block" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// TrafficAnalytics: optional
	TrafficAnalytics *TrafficAnalytics `hcl:"traffic_analytics,block"`
}

type azurermNetworkWatcherFlowLogAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(anwfl.ref.Append("enabled"))
}

// Id returns a reference to field id of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("name"))
}

// NetworkSecurityGroupId returns a reference to field network_security_group_id of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) NetworkSecurityGroupId() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("network_security_group_id"))
}

// NetworkWatcherName returns a reference to field network_watcher_name of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) NetworkWatcherName() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("network_watcher_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("resource_group_name"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(anwfl.ref.Append("storage_account_id"))
}

// Tags returns a reference to field tags of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](anwfl.ref.Append("tags"))
}

// Version returns a reference to field version of azurerm_network_watcher_flow_log.
func (anwfl azurermNetworkWatcherFlowLogAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(anwfl.ref.Append("version"))
}

func (anwfl azurermNetworkWatcherFlowLogAttributes) RetentionPolicy() terra.ListValue[RetentionPolicyAttributes] {
	return terra.ReferenceAsList[RetentionPolicyAttributes](anwfl.ref.Append("retention_policy"))
}

func (anwfl azurermNetworkWatcherFlowLogAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](anwfl.ref.Append("timeouts"))
}

func (anwfl azurermNetworkWatcherFlowLogAttributes) TrafficAnalytics() terra.ListValue[TrafficAnalyticsAttributes] {
	return terra.ReferenceAsList[TrafficAnalyticsAttributes](anwfl.ref.Append("traffic_analytics"))
}

type azurermNetworkWatcherFlowLogState struct {
	Enabled                bool                    `json:"enabled"`
	Id                     string                  `json:"id"`
	Location               string                  `json:"location"`
	Name                   string                  `json:"name"`
	NetworkSecurityGroupId string                  `json:"network_security_group_id"`
	NetworkWatcherName     string                  `json:"network_watcher_name"`
	ResourceGroupName      string                  `json:"resource_group_name"`
	StorageAccountId       string                  `json:"storage_account_id"`
	Tags                   map[string]string       `json:"tags"`
	Version                float64                 `json:"version"`
	RetentionPolicy        []RetentionPolicyState  `json:"retention_policy"`
	Timeouts               *TimeoutsState          `json:"timeouts"`
	TrafficAnalytics       []TrafficAnalyticsState `json:"traffic_analytics"`
}
