// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagernetworkgroup "github.com/golingon/terraproviders/azurerm/3.67.0/networkmanagernetworkgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerNetworkGroup creates a new instance of [NetworkManagerNetworkGroup].
func NewNetworkManagerNetworkGroup(name string, args NetworkManagerNetworkGroupArgs) *NetworkManagerNetworkGroup {
	return &NetworkManagerNetworkGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerNetworkGroup)(nil)

// NetworkManagerNetworkGroup represents the Terraform resource azurerm_network_manager_network_group.
type NetworkManagerNetworkGroup struct {
	Name      string
	Args      NetworkManagerNetworkGroupArgs
	state     *networkManagerNetworkGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) Type() string {
	return "azurerm_network_manager_network_group"
}

// LocalName returns the local name for [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) LocalName() string {
	return nmng.Name
}

// Configuration returns the configuration (args) for [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) Configuration() interface{} {
	return nmng.Args
}

// DependOn is used for other resources to depend on [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(nmng)
}

// Dependencies returns the list of resources [NetworkManagerNetworkGroup] depends_on.
func (nmng *NetworkManagerNetworkGroup) Dependencies() terra.Dependencies {
	return nmng.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) LifecycleManagement() *terra.Lifecycle {
	return nmng.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerNetworkGroup].
func (nmng *NetworkManagerNetworkGroup) Attributes() networkManagerNetworkGroupAttributes {
	return networkManagerNetworkGroupAttributes{ref: terra.ReferenceResource(nmng)}
}

// ImportState imports the given attribute values into [NetworkManagerNetworkGroup]'s state.
func (nmng *NetworkManagerNetworkGroup) ImportState(av io.Reader) error {
	nmng.state = &networkManagerNetworkGroupState{}
	if err := json.NewDecoder(av).Decode(nmng.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmng.Type(), nmng.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerNetworkGroup] has state.
func (nmng *NetworkManagerNetworkGroup) State() (*networkManagerNetworkGroupState, bool) {
	return nmng.state, nmng.state != nil
}

// StateMust returns the state for [NetworkManagerNetworkGroup]. Panics if the state is nil.
func (nmng *NetworkManagerNetworkGroup) StateMust() *networkManagerNetworkGroupState {
	if nmng.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmng.Type(), nmng.LocalName()))
	}
	return nmng.state
}

// NetworkManagerNetworkGroupArgs contains the configurations for azurerm_network_manager_network_group.
type NetworkManagerNetworkGroupArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagernetworkgroup.Timeouts `hcl:"timeouts,block"`
}
type networkManagerNetworkGroupAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of azurerm_network_manager_network_group.
func (nmng networkManagerNetworkGroupAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_network_manager_network_group.
func (nmng networkManagerNetworkGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_network_group.
func (nmng networkManagerNetworkGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("name"))
}

// NetworkManagerId returns a reference to field network_manager_id of azurerm_network_manager_network_group.
func (nmng networkManagerNetworkGroupAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceAsString(nmng.ref.Append("network_manager_id"))
}

func (nmng networkManagerNetworkGroupAttributes) Timeouts() networkmanagernetworkgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagernetworkgroup.TimeoutsAttributes](nmng.ref.Append("timeouts"))
}

type networkManagerNetworkGroupState struct {
	Description      string                                    `json:"description"`
	Id               string                                    `json:"id"`
	Name             string                                    `json:"name"`
	NetworkManagerId string                                    `json:"network_manager_id"`
	Timeouts         *networkmanagernetworkgroup.TimeoutsState `json:"timeouts"`
}
