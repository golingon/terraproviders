// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagertransitgatewayconnectpeerassociation "github.com/golingon/terraproviders/aws/5.11.0/networkmanagertransitgatewayconnectpeerassociation"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerTransitGatewayConnectPeerAssociation creates a new instance of [NetworkmanagerTransitGatewayConnectPeerAssociation].
func NewNetworkmanagerTransitGatewayConnectPeerAssociation(name string, args NetworkmanagerTransitGatewayConnectPeerAssociationArgs) *NetworkmanagerTransitGatewayConnectPeerAssociation {
	return &NetworkmanagerTransitGatewayConnectPeerAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerTransitGatewayConnectPeerAssociation)(nil)

// NetworkmanagerTransitGatewayConnectPeerAssociation represents the Terraform resource aws_networkmanager_transit_gateway_connect_peer_association.
type NetworkmanagerTransitGatewayConnectPeerAssociation struct {
	Name      string
	Args      NetworkmanagerTransitGatewayConnectPeerAssociationArgs
	state     *networkmanagerTransitGatewayConnectPeerAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) Type() string {
	return "aws_networkmanager_transit_gateway_connect_peer_association"
}

// LocalName returns the local name for [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) LocalName() string {
	return ntgcpa.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) Configuration() interface{} {
	return ntgcpa.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(ntgcpa)
}

// Dependencies returns the list of resources [NetworkmanagerTransitGatewayConnectPeerAssociation] depends_on.
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) Dependencies() terra.Dependencies {
	return ntgcpa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) LifecycleManagement() *terra.Lifecycle {
	return ntgcpa.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerTransitGatewayConnectPeerAssociation].
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) Attributes() networkmanagerTransitGatewayConnectPeerAssociationAttributes {
	return networkmanagerTransitGatewayConnectPeerAssociationAttributes{ref: terra.ReferenceResource(ntgcpa)}
}

// ImportState imports the given attribute values into [NetworkmanagerTransitGatewayConnectPeerAssociation]'s state.
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) ImportState(av io.Reader) error {
	ntgcpa.state = &networkmanagerTransitGatewayConnectPeerAssociationState{}
	if err := json.NewDecoder(av).Decode(ntgcpa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ntgcpa.Type(), ntgcpa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerTransitGatewayConnectPeerAssociation] has state.
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) State() (*networkmanagerTransitGatewayConnectPeerAssociationState, bool) {
	return ntgcpa.state, ntgcpa.state != nil
}

// StateMust returns the state for [NetworkmanagerTransitGatewayConnectPeerAssociation]. Panics if the state is nil.
func (ntgcpa *NetworkmanagerTransitGatewayConnectPeerAssociation) StateMust() *networkmanagerTransitGatewayConnectPeerAssociationState {
	if ntgcpa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ntgcpa.Type(), ntgcpa.LocalName()))
	}
	return ntgcpa.state
}

// NetworkmanagerTransitGatewayConnectPeerAssociationArgs contains the configurations for aws_networkmanager_transit_gateway_connect_peer_association.
type NetworkmanagerTransitGatewayConnectPeerAssociationArgs struct {
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, optional
	LinkId terra.StringValue `hcl:"link_id,attr"`
	// TransitGatewayConnectPeerArn: string, required
	TransitGatewayConnectPeerArn terra.StringValue `hcl:"transit_gateway_connect_peer_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagertransitgatewayconnectpeerassociation.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerTransitGatewayConnectPeerAssociationAttributes struct {
	ref terra.Reference
}

// DeviceId returns a reference to field device_id of aws_networkmanager_transit_gateway_connect_peer_association.
func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(ntgcpa.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_transit_gateway_connect_peer_association.
func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ntgcpa.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_transit_gateway_connect_peer_association.
func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ntgcpa.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_transit_gateway_connect_peer_association.
func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(ntgcpa.ref.Append("link_id"))
}

// TransitGatewayConnectPeerArn returns a reference to field transit_gateway_connect_peer_arn of aws_networkmanager_transit_gateway_connect_peer_association.
func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) TransitGatewayConnectPeerArn() terra.StringValue {
	return terra.ReferenceAsString(ntgcpa.ref.Append("transit_gateway_connect_peer_arn"))
}

func (ntgcpa networkmanagerTransitGatewayConnectPeerAssociationAttributes) Timeouts() networkmanagertransitgatewayconnectpeerassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagertransitgatewayconnectpeerassociation.TimeoutsAttributes](ntgcpa.ref.Append("timeouts"))
}

type networkmanagerTransitGatewayConnectPeerAssociationState struct {
	DeviceId                     string                                                            `json:"device_id"`
	GlobalNetworkId              string                                                            `json:"global_network_id"`
	Id                           string                                                            `json:"id"`
	LinkId                       string                                                            `json:"link_id"`
	TransitGatewayConnectPeerArn string                                                            `json:"transit_gateway_connect_peer_arn"`
	Timeouts                     *networkmanagertransitgatewayconnectpeerassociation.TimeoutsState `json:"timeouts"`
}
