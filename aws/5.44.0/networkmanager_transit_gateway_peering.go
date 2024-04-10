// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	networkmanagertransitgatewaypeering "github.com/golingon/terraproviders/aws/5.44.0/networkmanagertransitgatewaypeering"
	"io"
)

// NewNetworkmanagerTransitGatewayPeering creates a new instance of [NetworkmanagerTransitGatewayPeering].
func NewNetworkmanagerTransitGatewayPeering(name string, args NetworkmanagerTransitGatewayPeeringArgs) *NetworkmanagerTransitGatewayPeering {
	return &NetworkmanagerTransitGatewayPeering{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerTransitGatewayPeering)(nil)

// NetworkmanagerTransitGatewayPeering represents the Terraform resource aws_networkmanager_transit_gateway_peering.
type NetworkmanagerTransitGatewayPeering struct {
	Name      string
	Args      NetworkmanagerTransitGatewayPeeringArgs
	state     *networkmanagerTransitGatewayPeeringState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) Type() string {
	return "aws_networkmanager_transit_gateway_peering"
}

// LocalName returns the local name for [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) LocalName() string {
	return ntgp.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) Configuration() interface{} {
	return ntgp.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) DependOn() terra.Reference {
	return terra.ReferenceResource(ntgp)
}

// Dependencies returns the list of resources [NetworkmanagerTransitGatewayPeering] depends_on.
func (ntgp *NetworkmanagerTransitGatewayPeering) Dependencies() terra.Dependencies {
	return ntgp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) LifecycleManagement() *terra.Lifecycle {
	return ntgp.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerTransitGatewayPeering].
func (ntgp *NetworkmanagerTransitGatewayPeering) Attributes() networkmanagerTransitGatewayPeeringAttributes {
	return networkmanagerTransitGatewayPeeringAttributes{ref: terra.ReferenceResource(ntgp)}
}

// ImportState imports the given attribute values into [NetworkmanagerTransitGatewayPeering]'s state.
func (ntgp *NetworkmanagerTransitGatewayPeering) ImportState(av io.Reader) error {
	ntgp.state = &networkmanagerTransitGatewayPeeringState{}
	if err := json.NewDecoder(av).Decode(ntgp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ntgp.Type(), ntgp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerTransitGatewayPeering] has state.
func (ntgp *NetworkmanagerTransitGatewayPeering) State() (*networkmanagerTransitGatewayPeeringState, bool) {
	return ntgp.state, ntgp.state != nil
}

// StateMust returns the state for [NetworkmanagerTransitGatewayPeering]. Panics if the state is nil.
func (ntgp *NetworkmanagerTransitGatewayPeering) StateMust() *networkmanagerTransitGatewayPeeringState {
	if ntgp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ntgp.Type(), ntgp.LocalName()))
	}
	return ntgp.state
}

// NetworkmanagerTransitGatewayPeeringArgs contains the configurations for aws_networkmanager_transit_gateway_peering.
type NetworkmanagerTransitGatewayPeeringArgs struct {
	// CoreNetworkId: string, required
	CoreNetworkId terra.StringValue `hcl:"core_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayArn: string, required
	TransitGatewayArn terra.StringValue `hcl:"transit_gateway_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagertransitgatewaypeering.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerTransitGatewayPeeringAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("arn"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("owner_account_id"))
}

// PeeringType returns a reference to field peering_type of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) PeeringType() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("peering_type"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("resource_arn"))
}

// Tags returns a reference to field tags of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ntgp.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ntgp.ref.Append("tags_all"))
}

// TransitGatewayArn returns a reference to field transit_gateway_arn of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) TransitGatewayArn() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("transit_gateway_arn"))
}

// TransitGatewayPeeringAttachmentId returns a reference to field transit_gateway_peering_attachment_id of aws_networkmanager_transit_gateway_peering.
func (ntgp networkmanagerTransitGatewayPeeringAttributes) TransitGatewayPeeringAttachmentId() terra.StringValue {
	return terra.ReferenceAsString(ntgp.ref.Append("transit_gateway_peering_attachment_id"))
}

func (ntgp networkmanagerTransitGatewayPeeringAttributes) Timeouts() networkmanagertransitgatewaypeering.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagertransitgatewaypeering.TimeoutsAttributes](ntgp.ref.Append("timeouts"))
}

type networkmanagerTransitGatewayPeeringState struct {
	Arn                               string                                             `json:"arn"`
	CoreNetworkArn                    string                                             `json:"core_network_arn"`
	CoreNetworkId                     string                                             `json:"core_network_id"`
	EdgeLocation                      string                                             `json:"edge_location"`
	Id                                string                                             `json:"id"`
	OwnerAccountId                    string                                             `json:"owner_account_id"`
	PeeringType                       string                                             `json:"peering_type"`
	ResourceArn                       string                                             `json:"resource_arn"`
	Tags                              map[string]string                                  `json:"tags"`
	TagsAll                           map[string]string                                  `json:"tags_all"`
	TransitGatewayArn                 string                                             `json:"transit_gateway_arn"`
	TransitGatewayPeeringAttachmentId string                                             `json:"transit_gateway_peering_attachment_id"`
	Timeouts                          *networkmanagertransitgatewaypeering.TimeoutsState `json:"timeouts"`
}
