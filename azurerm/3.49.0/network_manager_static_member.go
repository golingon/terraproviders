// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagerstaticmember "github.com/golingon/terraproviders/azurerm/3.49.0/networkmanagerstaticmember"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkManagerStaticMember(name string, args NetworkManagerStaticMemberArgs) *NetworkManagerStaticMember {
	return &NetworkManagerStaticMember{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerStaticMember)(nil)

type NetworkManagerStaticMember struct {
	Name  string
	Args  NetworkManagerStaticMemberArgs
	state *networkManagerStaticMemberState
}

func (nmsm *NetworkManagerStaticMember) Type() string {
	return "azurerm_network_manager_static_member"
}

func (nmsm *NetworkManagerStaticMember) LocalName() string {
	return nmsm.Name
}

func (nmsm *NetworkManagerStaticMember) Configuration() interface{} {
	return nmsm.Args
}

func (nmsm *NetworkManagerStaticMember) Attributes() networkManagerStaticMemberAttributes {
	return networkManagerStaticMemberAttributes{ref: terra.ReferenceResource(nmsm)}
}

func (nmsm *NetworkManagerStaticMember) ImportState(av io.Reader) error {
	nmsm.state = &networkManagerStaticMemberState{}
	if err := json.NewDecoder(av).Decode(nmsm.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsm.Type(), nmsm.LocalName(), err)
	}
	return nil
}

func (nmsm *NetworkManagerStaticMember) State() (*networkManagerStaticMemberState, bool) {
	return nmsm.state, nmsm.state != nil
}

func (nmsm *NetworkManagerStaticMember) StateMust() *networkManagerStaticMemberState {
	if nmsm.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsm.Type(), nmsm.LocalName()))
	}
	return nmsm.state
}

func (nmsm *NetworkManagerStaticMember) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsm)
}

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
	// DependsOn contains resources that NetworkManagerStaticMember depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkManagerStaticMemberAttributes struct {
	ref terra.Reference
}

func (nmsm networkManagerStaticMemberAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nmsm.ref.Append("id"))
}

func (nmsm networkManagerStaticMemberAttributes) Name() terra.StringValue {
	return terra.ReferenceString(nmsm.ref.Append("name"))
}

func (nmsm networkManagerStaticMemberAttributes) NetworkGroupId() terra.StringValue {
	return terra.ReferenceString(nmsm.ref.Append("network_group_id"))
}

func (nmsm networkManagerStaticMemberAttributes) Region() terra.StringValue {
	return terra.ReferenceString(nmsm.ref.Append("region"))
}

func (nmsm networkManagerStaticMemberAttributes) TargetVirtualNetworkId() terra.StringValue {
	return terra.ReferenceString(nmsm.ref.Append("target_virtual_network_id"))
}

func (nmsm networkManagerStaticMemberAttributes) Timeouts() networkmanagerstaticmember.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanagerstaticmember.TimeoutsAttributes](nmsm.ref.Append("timeouts"))
}

type networkManagerStaticMemberState struct {
	Id                     string                                    `json:"id"`
	Name                   string                                    `json:"name"`
	NetworkGroupId         string                                    `json:"network_group_id"`
	Region                 string                                    `json:"region"`
	TargetVirtualNetworkId string                                    `json:"target_virtual_network_id"`
	Timeouts               *networkmanagerstaticmember.TimeoutsState `json:"timeouts"`
}
