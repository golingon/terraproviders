// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagerstaticmember "github.com/golingon/terraproviders/azurerm/3.64.0/networkmanagerstaticmember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkManagerStaticMember creates a new instance of [NetworkManagerStaticMember].
func NewNetworkManagerStaticMember(name string, args NetworkManagerStaticMemberArgs) *NetworkManagerStaticMember {
	return &NetworkManagerStaticMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerStaticMember)(nil)

// NetworkManagerStaticMember represents the Terraform resource azurerm_network_manager_static_member.
type NetworkManagerStaticMember struct {
	Name      string
	Args      NetworkManagerStaticMemberArgs
	state     *networkManagerStaticMemberState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) Type() string {
	return "azurerm_network_manager_static_member"
}

// LocalName returns the local name for [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) LocalName() string {
	return nmsm.Name
}

// Configuration returns the configuration (args) for [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) Configuration() interface{} {
	return nmsm.Args
}

// DependOn is used for other resources to depend on [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsm)
}

// Dependencies returns the list of resources [NetworkManagerStaticMember] depends_on.
func (nmsm *NetworkManagerStaticMember) Dependencies() terra.Dependencies {
	return nmsm.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) LifecycleManagement() *terra.Lifecycle {
	return nmsm.Lifecycle
}

// Attributes returns the attributes for [NetworkManagerStaticMember].
func (nmsm *NetworkManagerStaticMember) Attributes() networkManagerStaticMemberAttributes {
	return networkManagerStaticMemberAttributes{ref: terra.ReferenceResource(nmsm)}
}

// ImportState imports the given attribute values into [NetworkManagerStaticMember]'s state.
func (nmsm *NetworkManagerStaticMember) ImportState(av io.Reader) error {
	nmsm.state = &networkManagerStaticMemberState{}
	if err := json.NewDecoder(av).Decode(nmsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsm.Type(), nmsm.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkManagerStaticMember] has state.
func (nmsm *NetworkManagerStaticMember) State() (*networkManagerStaticMemberState, bool) {
	return nmsm.state, nmsm.state != nil
}

// StateMust returns the state for [NetworkManagerStaticMember]. Panics if the state is nil.
func (nmsm *NetworkManagerStaticMember) StateMust() *networkManagerStaticMemberState {
	if nmsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsm.Type(), nmsm.LocalName()))
	}
	return nmsm.state
}

// NetworkManagerStaticMemberArgs contains the configurations for azurerm_network_manager_static_member.
type NetworkManagerStaticMemberArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkGroupId: string, required
	NetworkGroupId terra.StringValue `hcl:"network_group_id,attr" validate:"required"`
	// TargetVirtualNetworkId: string, required
	TargetVirtualNetworkId terra.StringValue `hcl:"target_virtual_network_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagerstaticmember.Timeouts `hcl:"timeouts,block"`
}
type networkManagerStaticMemberAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_manager_static_member.
func (nmsm networkManagerStaticMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nmsm.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_network_manager_static_member.
func (nmsm networkManagerStaticMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nmsm.ref.Append("name"))
}

// NetworkGroupId returns a reference to field network_group_id of azurerm_network_manager_static_member.
func (nmsm networkManagerStaticMemberAttributes) NetworkGroupId() terra.StringValue {
	return terra.ReferenceAsString(nmsm.ref.Append("network_group_id"))
}

// Region returns a reference to field region of azurerm_network_manager_static_member.
func (nmsm networkManagerStaticMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(nmsm.ref.Append("region"))
}

// TargetVirtualNetworkId returns a reference to field target_virtual_network_id of azurerm_network_manager_static_member.
func (nmsm networkManagerStaticMemberAttributes) TargetVirtualNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nmsm.ref.Append("target_virtual_network_id"))
}

func (nmsm networkManagerStaticMemberAttributes) Timeouts() networkmanagerstaticmember.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerstaticmember.TimeoutsAttributes](nmsm.ref.Append("timeouts"))
}

type networkManagerStaticMemberState struct {
	Id                     string                                    `json:"id"`
	Name                   string                                    `json:"name"`
	NetworkGroupId         string                                    `json:"network_group_id"`
	Region                 string                                    `json:"region"`
	TargetVirtualNetworkId string                                    `json:"target_virtual_network_id"`
	Timeouts               *networkmanagerstaticmember.TimeoutsState `json:"timeouts"`
}
