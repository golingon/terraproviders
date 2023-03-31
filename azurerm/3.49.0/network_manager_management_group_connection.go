// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagermanagementgroupconnection "github.com/golingon/terraproviders/azurerm/3.49.0/networkmanagermanagementgroupconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkManagerManagementGroupConnection(name string, args NetworkManagerManagementGroupConnectionArgs) *NetworkManagerManagementGroupConnection {
	return &NetworkManagerManagementGroupConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerManagementGroupConnection)(nil)

type NetworkManagerManagementGroupConnection struct {
	Name  string
	Args  NetworkManagerManagementGroupConnectionArgs
	state *networkManagerManagementGroupConnectionState
}

func (nmmgc *NetworkManagerManagementGroupConnection) Type() string {
	return "azurerm_network_manager_management_group_connection"
}

func (nmmgc *NetworkManagerManagementGroupConnection) LocalName() string {
	return nmmgc.Name
}

func (nmmgc *NetworkManagerManagementGroupConnection) Configuration() interface{} {
	return nmmgc.Args
}

func (nmmgc *NetworkManagerManagementGroupConnection) Attributes() networkManagerManagementGroupConnectionAttributes {
	return networkManagerManagementGroupConnectionAttributes{ref: terra.ReferenceResource(nmmgc)}
}

func (nmmgc *NetworkManagerManagementGroupConnection) ImportState(av io.Reader) error {
	nmmgc.state = &networkManagerManagementGroupConnectionState{}
	if err := json.NewDecoder(av).Decode(nmmgc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmmgc.Type(), nmmgc.LocalName(), err)
	}
	return nil
}

func (nmmgc *NetworkManagerManagementGroupConnection) State() (*networkManagerManagementGroupConnectionState, bool) {
	return nmmgc.state, nmmgc.state != nil
}

func (nmmgc *NetworkManagerManagementGroupConnection) StateMust() *networkManagerManagementGroupConnectionState {
	if nmmgc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmmgc.Type(), nmmgc.LocalName()))
	}
	return nmmgc.state
}

func (nmmgc *NetworkManagerManagementGroupConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmmgc)
}

type NetworkManagerManagementGroupConnectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagermanagementgroupconnection.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NetworkManagerManagementGroupConnection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkManagerManagementGroupConnectionAttributes struct {
	ref terra.Reference
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("connection_state"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("description"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("id"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("management_group_id"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("name"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceString(nmmgc.ref.Append("network_manager_id"))
}

func (nmmgc networkManagerManagementGroupConnectionAttributes) Timeouts() networkmanagermanagementgroupconnection.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanagermanagementgroupconnection.TimeoutsAttributes](nmmgc.ref.Append("timeouts"))
}

type networkManagerManagementGroupConnectionState struct {
	ConnectionState   string                                                 `json:"connection_state"`
	Description       string                                                 `json:"description"`
	Id                string                                                 `json:"id"`
	ManagementGroupId string                                                 `json:"management_group_id"`
	Name              string                                                 `json:"name"`
	NetworkManagerId  string                                                 `json:"network_manager_id"`
	Timeouts          *networkmanagermanagementgroupconnection.TimeoutsState `json:"timeouts"`
}
