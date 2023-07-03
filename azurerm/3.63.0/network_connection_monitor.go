// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkconnectionmonitor "github.com/golingon/terraproviders/azurerm/3.63.0/networkconnectionmonitor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkConnectionMonitor creates a new instance of [NetworkConnectionMonitor].
func NewNetworkConnectionMonitor(name string, args NetworkConnectionMonitorArgs) *NetworkConnectionMonitor {
	return &NetworkConnectionMonitor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkConnectionMonitor)(nil)

// NetworkConnectionMonitor represents the Terraform resource azurerm_network_connection_monitor.
type NetworkConnectionMonitor struct {
	Name      string
	Args      NetworkConnectionMonitorArgs
	state     *networkConnectionMonitorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) Type() string {
	return "azurerm_network_connection_monitor"
}

// LocalName returns the local name for [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) LocalName() string {
	return ncm.Name
}

// Configuration returns the configuration (args) for [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) Configuration() interface{} {
	return ncm.Args
}

// DependOn is used for other resources to depend on [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) DependOn() terra.Reference {
	return terra.ReferenceResource(ncm)
}

// Dependencies returns the list of resources [NetworkConnectionMonitor] depends_on.
func (ncm *NetworkConnectionMonitor) Dependencies() terra.Dependencies {
	return ncm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) LifecycleManagement() *terra.Lifecycle {
	return ncm.Lifecycle
}

// Attributes returns the attributes for [NetworkConnectionMonitor].
func (ncm *NetworkConnectionMonitor) Attributes() networkConnectionMonitorAttributes {
	return networkConnectionMonitorAttributes{ref: terra.ReferenceResource(ncm)}
}

// ImportState imports the given attribute values into [NetworkConnectionMonitor]'s state.
func (ncm *NetworkConnectionMonitor) ImportState(av io.Reader) error {
	ncm.state = &networkConnectionMonitorState{}
	if err := json.NewDecoder(av).Decode(ncm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncm.Type(), ncm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkConnectionMonitor] has state.
func (ncm *NetworkConnectionMonitor) State() (*networkConnectionMonitorState, bool) {
	return ncm.state, ncm.state != nil
}

// StateMust returns the state for [NetworkConnectionMonitor]. Panics if the state is nil.
func (ncm *NetworkConnectionMonitor) StateMust() *networkConnectionMonitorState {
	if ncm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncm.Type(), ncm.LocalName()))
	}
	return ncm.state
}

// NetworkConnectionMonitorArgs contains the configurations for azurerm_network_connection_monitor.
type NetworkConnectionMonitorArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkWatcherId: string, required
	NetworkWatcherId terra.StringValue `hcl:"network_watcher_id,attr" validate:"required"`
	// Notes: string, optional
	Notes terra.StringValue `hcl:"notes,attr"`
	// OutputWorkspaceResourceIds: set of string, optional
	OutputWorkspaceResourceIds terra.SetValue[terra.StringValue] `hcl:"output_workspace_resource_ids,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Endpoint: min=1
	Endpoint []networkconnectionmonitor.Endpoint `hcl:"endpoint,block" validate:"min=1"`
	// TestConfiguration: min=1
	TestConfiguration []networkconnectionmonitor.TestConfiguration `hcl:"test_configuration,block" validate:"min=1"`
	// TestGroup: min=1
	TestGroup []networkconnectionmonitor.TestGroup `hcl:"test_group,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *networkconnectionmonitor.Timeouts `hcl:"timeouts,block"`
}
type networkConnectionMonitorAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("name"))
}

// NetworkWatcherId returns a reference to field network_watcher_id of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) NetworkWatcherId() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("network_watcher_id"))
}

// Notes returns a reference to field notes of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) Notes() terra.StringValue {
	return terra.ReferenceAsString(ncm.ref.Append("notes"))
}

// OutputWorkspaceResourceIds returns a reference to field output_workspace_resource_ids of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) OutputWorkspaceResourceIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ncm.ref.Append("output_workspace_resource_ids"))
}

// Tags returns a reference to field tags of azurerm_network_connection_monitor.
func (ncm networkConnectionMonitorAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncm.ref.Append("tags"))
}

func (ncm networkConnectionMonitorAttributes) Endpoint() terra.SetValue[networkconnectionmonitor.EndpointAttributes] {
	return terra.ReferenceAsSet[networkconnectionmonitor.EndpointAttributes](ncm.ref.Append("endpoint"))
}

func (ncm networkConnectionMonitorAttributes) TestConfiguration() terra.SetValue[networkconnectionmonitor.TestConfigurationAttributes] {
	return terra.ReferenceAsSet[networkconnectionmonitor.TestConfigurationAttributes](ncm.ref.Append("test_configuration"))
}

func (ncm networkConnectionMonitorAttributes) TestGroup() terra.SetValue[networkconnectionmonitor.TestGroupAttributes] {
	return terra.ReferenceAsSet[networkconnectionmonitor.TestGroupAttributes](ncm.ref.Append("test_group"))
}

func (ncm networkConnectionMonitorAttributes) Timeouts() networkconnectionmonitor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkconnectionmonitor.TimeoutsAttributes](ncm.ref.Append("timeouts"))
}

type networkConnectionMonitorState struct {
	Id                         string                                            `json:"id"`
	Location                   string                                            `json:"location"`
	Name                       string                                            `json:"name"`
	NetworkWatcherId           string                                            `json:"network_watcher_id"`
	Notes                      string                                            `json:"notes"`
	OutputWorkspaceResourceIds []string                                          `json:"output_workspace_resource_ids"`
	Tags                       map[string]string                                 `json:"tags"`
	Endpoint                   []networkconnectionmonitor.EndpointState          `json:"endpoint"`
	TestConfiguration          []networkconnectionmonitor.TestConfigurationState `json:"test_configuration"`
	TestGroup                  []networkconnectionmonitor.TestGroupState         `json:"test_group"`
	Timeouts                   *networkconnectionmonitor.TimeoutsState           `json:"timeouts"`
}
