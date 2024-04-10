// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networkwatcher "github.com/golingon/terraproviders/azurerm/3.98.0/networkwatcher"
	"io"
)

// NewNetworkWatcher creates a new instance of [NetworkWatcher].
func NewNetworkWatcher(name string, args NetworkWatcherArgs) *NetworkWatcher {
	return &NetworkWatcher{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkWatcher)(nil)

// NetworkWatcher represents the Terraform resource azurerm_network_watcher.
type NetworkWatcher struct {
	Name      string
	Args      NetworkWatcherArgs
	state     *networkWatcherState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkWatcher].
func (nw *NetworkWatcher) Type() string {
	return "azurerm_network_watcher"
}

// LocalName returns the local name for [NetworkWatcher].
func (nw *NetworkWatcher) LocalName() string {
	return nw.Name
}

// Configuration returns the configuration (args) for [NetworkWatcher].
func (nw *NetworkWatcher) Configuration() interface{} {
	return nw.Args
}

// DependOn is used for other resources to depend on [NetworkWatcher].
func (nw *NetworkWatcher) DependOn() terra.Reference {
	return terra.ReferenceResource(nw)
}

// Dependencies returns the list of resources [NetworkWatcher] depends_on.
func (nw *NetworkWatcher) Dependencies() terra.Dependencies {
	return nw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkWatcher].
func (nw *NetworkWatcher) LifecycleManagement() *terra.Lifecycle {
	return nw.Lifecycle
}

// Attributes returns the attributes for [NetworkWatcher].
func (nw *NetworkWatcher) Attributes() networkWatcherAttributes {
	return networkWatcherAttributes{ref: terra.ReferenceResource(nw)}
}

// ImportState imports the given attribute values into [NetworkWatcher]'s state.
func (nw *NetworkWatcher) ImportState(av io.Reader) error {
	nw.state = &networkWatcherState{}
	if err := json.NewDecoder(av).Decode(nw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nw.Type(), nw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkWatcher] has state.
func (nw *NetworkWatcher) State() (*networkWatcherState, bool) {
	return nw.state, nw.state != nil
}

// StateMust returns the state for [NetworkWatcher]. Panics if the state is nil.
func (nw *NetworkWatcher) StateMust() *networkWatcherState {
	if nw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nw.Type(), nw.LocalName()))
	}
	return nw.state
}

// NetworkWatcherArgs contains the configurations for azurerm_network_watcher.
type NetworkWatcherArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *networkwatcher.Timeouts `hcl:"timeouts,block"`
}
type networkWatcherAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_watcher.
func (nw networkWatcherAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_watcher.
func (nw networkWatcherAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_watcher.
func (nw networkWatcherAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_watcher.
func (nw networkWatcherAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nw.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_watcher.
func (nw networkWatcherAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nw.ref.Append("tags"))
}

func (nw networkWatcherAttributes) Timeouts() networkwatcher.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkwatcher.TimeoutsAttributes](nw.ref.Append("timeouts"))
}

type networkWatcherState struct {
	Id                string                        `json:"id"`
	Location          string                        `json:"location"`
	Name              string                        `json:"name"`
	ResourceGroupName string                        `json:"resource_group_name"`
	Tags              map[string]string             `json:"tags"`
	Timeouts          *networkwatcher.TimeoutsState `json:"timeouts"`
}
