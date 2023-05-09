// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkprofile "github.com/golingon/terraproviders/azurerm/3.55.0/networkprofile"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkProfile creates a new instance of [NetworkProfile].
func NewNetworkProfile(name string, args NetworkProfileArgs) *NetworkProfile {
	return &NetworkProfile{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkProfile)(nil)

// NetworkProfile represents the Terraform resource azurerm_network_profile.
type NetworkProfile struct {
	Name      string
	Args      NetworkProfileArgs
	state     *networkProfileState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkProfile].
func (np *NetworkProfile) Type() string {
	return "azurerm_network_profile"
}

// LocalName returns the local name for [NetworkProfile].
func (np *NetworkProfile) LocalName() string {
	return np.Name
}

// Configuration returns the configuration (args) for [NetworkProfile].
func (np *NetworkProfile) Configuration() interface{} {
	return np.Args
}

// DependOn is used for other resources to depend on [NetworkProfile].
func (np *NetworkProfile) DependOn() terra.Reference {
	return terra.ReferenceResource(np)
}

// Dependencies returns the list of resources [NetworkProfile] depends_on.
func (np *NetworkProfile) Dependencies() terra.Dependencies {
	return np.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkProfile].
func (np *NetworkProfile) LifecycleManagement() *terra.Lifecycle {
	return np.Lifecycle
}

// Attributes returns the attributes for [NetworkProfile].
func (np *NetworkProfile) Attributes() networkProfileAttributes {
	return networkProfileAttributes{ref: terra.ReferenceResource(np)}
}

// ImportState imports the given attribute values into [NetworkProfile]'s state.
func (np *NetworkProfile) ImportState(av io.Reader) error {
	np.state = &networkProfileState{}
	if err := json.NewDecoder(av).Decode(np.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", np.Type(), np.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkProfile] has state.
func (np *NetworkProfile) State() (*networkProfileState, bool) {
	return np.state, np.state != nil
}

// StateMust returns the state for [NetworkProfile]. Panics if the state is nil.
func (np *NetworkProfile) StateMust() *networkProfileState {
	if np.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", np.Type(), np.LocalName()))
	}
	return np.state
}

// NetworkProfileArgs contains the configurations for azurerm_network_profile.
type NetworkProfileArgs struct {
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
	// ContainerNetworkInterface: required
	ContainerNetworkInterface *networkprofile.ContainerNetworkInterface `hcl:"container_network_interface,block" validate:"required"`
	// Timeouts: optional
	Timeouts *networkprofile.Timeouts `hcl:"timeouts,block"`
}
type networkProfileAttributes struct {
	ref terra.Reference
}

// ContainerNetworkInterfaceIds returns a reference to field container_network_interface_ids of azurerm_network_profile.
func (np networkProfileAttributes) ContainerNetworkInterfaceIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](np.ref.Append("container_network_interface_ids"))
}

// Id returns a reference to field id of azurerm_network_profile.
func (np networkProfileAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_profile.
func (np networkProfileAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_profile.
func (np networkProfileAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_profile.
func (np networkProfileAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(np.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_profile.
func (np networkProfileAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](np.ref.Append("tags"))
}

func (np networkProfileAttributes) ContainerNetworkInterface() terra.ListValue[networkprofile.ContainerNetworkInterfaceAttributes] {
	return terra.ReferenceAsList[networkprofile.ContainerNetworkInterfaceAttributes](np.ref.Append("container_network_interface"))
}

func (np networkProfileAttributes) Timeouts() networkprofile.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkprofile.TimeoutsAttributes](np.ref.Append("timeouts"))
}

type networkProfileState struct {
	ContainerNetworkInterfaceIds []string                                        `json:"container_network_interface_ids"`
	Id                           string                                          `json:"id"`
	Location                     string                                          `json:"location"`
	Name                         string                                          `json:"name"`
	ResourceGroupName            string                                          `json:"resource_group_name"`
	Tags                         map[string]string                               `json:"tags"`
	ContainerNetworkInterface    []networkprofile.ContainerNetworkInterfaceState `json:"container_network_interface"`
	Timeouts                     *networkprofile.TimeoutsState                   `json:"timeouts"`
}
