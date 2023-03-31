// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	networkmanagersubscriptionconnection "github.com/golingon/terraproviders/azurerm/3.49.0/networkmanagersubscriptionconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewNetworkManagerSubscriptionConnection(name string, args NetworkManagerSubscriptionConnectionArgs) *NetworkManagerSubscriptionConnection {
	return &NetworkManagerSubscriptionConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkManagerSubscriptionConnection)(nil)

type NetworkManagerSubscriptionConnection struct {
	Name  string
	Args  NetworkManagerSubscriptionConnectionArgs
	state *networkManagerSubscriptionConnectionState
}

func (nmsc *NetworkManagerSubscriptionConnection) Type() string {
	return "azurerm_network_manager_subscription_connection"
}

func (nmsc *NetworkManagerSubscriptionConnection) LocalName() string {
	return nmsc.Name
}

func (nmsc *NetworkManagerSubscriptionConnection) Configuration() interface{} {
	return nmsc.Args
}

func (nmsc *NetworkManagerSubscriptionConnection) Attributes() networkManagerSubscriptionConnectionAttributes {
	return networkManagerSubscriptionConnectionAttributes{ref: terra.ReferenceResource(nmsc)}
}

func (nmsc *NetworkManagerSubscriptionConnection) ImportState(av io.Reader) error {
	nmsc.state = &networkManagerSubscriptionConnectionState{}
	if err := json.NewDecoder(av).Decode(nmsc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nmsc.Type(), nmsc.LocalName(), err)
	}
	return nil
}

func (nmsc *NetworkManagerSubscriptionConnection) State() (*networkManagerSubscriptionConnectionState, bool) {
	return nmsc.state, nmsc.state != nil
}

func (nmsc *NetworkManagerSubscriptionConnection) StateMust() *networkManagerSubscriptionConnectionState {
	if nmsc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nmsc.Type(), nmsc.LocalName()))
	}
	return nmsc.state
}

func (nmsc *NetworkManagerSubscriptionConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(nmsc)
}

type NetworkManagerSubscriptionConnectionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkManagerId: string, required
	NetworkManagerId terra.StringValue `hcl:"network_manager_id,attr" validate:"required"`
	// SubscriptionId: string, required
	SubscriptionId terra.StringValue `hcl:"subscription_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagersubscriptionconnection.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that NetworkManagerSubscriptionConnection depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type networkManagerSubscriptionConnectionAttributes struct {
	ref terra.Reference
}

func (nmsc networkManagerSubscriptionConnectionAttributes) ConnectionState() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("connection_state"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("description"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("id"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) Name() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("name"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) NetworkManagerId() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("network_manager_id"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) SubscriptionId() terra.StringValue {
	return terra.ReferenceString(nmsc.ref.Append("subscription_id"))
}

func (nmsc networkManagerSubscriptionConnectionAttributes) Timeouts() networkmanagersubscriptionconnection.TimeoutsAttributes {
	return terra.ReferenceSingle[networkmanagersubscriptionconnection.TimeoutsAttributes](nmsc.ref.Append("timeouts"))
}

type networkManagerSubscriptionConnectionState struct {
	ConnectionState  string                                              `json:"connection_state"`
	Description      string                                              `json:"description"`
	Id               string                                              `json:"id"`
	Name             string                                              `json:"name"`
	NetworkManagerId string                                              `json:"network_manager_id"`
	SubscriptionId   string                                              `json:"subscription_id"`
	Timeouts         *networkmanagersubscriptionconnection.TimeoutsState `json:"timeouts"`
}
