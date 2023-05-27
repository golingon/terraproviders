// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagertransitgatewayroutetableattachment "github.com/golingon/terraproviders/aws/5.0.1/networkmanagertransitgatewayroutetableattachment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerTransitGatewayRouteTableAttachment creates a new instance of [NetworkmanagerTransitGatewayRouteTableAttachment].
func NewNetworkmanagerTransitGatewayRouteTableAttachment(name string, args NetworkmanagerTransitGatewayRouteTableAttachmentArgs) *NetworkmanagerTransitGatewayRouteTableAttachment {
	return &NetworkmanagerTransitGatewayRouteTableAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerTransitGatewayRouteTableAttachment)(nil)

// NetworkmanagerTransitGatewayRouteTableAttachment represents the Terraform resource aws_networkmanager_transit_gateway_route_table_attachment.
type NetworkmanagerTransitGatewayRouteTableAttachment struct {
	Name      string
	Args      NetworkmanagerTransitGatewayRouteTableAttachmentArgs
	state     *networkmanagerTransitGatewayRouteTableAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) Type() string {
	return "aws_networkmanager_transit_gateway_route_table_attachment"
}

// LocalName returns the local name for [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) LocalName() string {
	return ntgrta.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) Configuration() interface{} {
	return ntgrta.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(ntgrta)
}

// Dependencies returns the list of resources [NetworkmanagerTransitGatewayRouteTableAttachment] depends_on.
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) Dependencies() terra.Dependencies {
	return ntgrta.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) LifecycleManagement() *terra.Lifecycle {
	return ntgrta.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerTransitGatewayRouteTableAttachment].
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) Attributes() networkmanagerTransitGatewayRouteTableAttachmentAttributes {
	return networkmanagerTransitGatewayRouteTableAttachmentAttributes{ref: terra.ReferenceResource(ntgrta)}
}

// ImportState imports the given attribute values into [NetworkmanagerTransitGatewayRouteTableAttachment]'s state.
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) ImportState(av io.Reader) error {
	ntgrta.state = &networkmanagerTransitGatewayRouteTableAttachmentState{}
	if err := json.NewDecoder(av).Decode(ntgrta.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ntgrta.Type(), ntgrta.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerTransitGatewayRouteTableAttachment] has state.
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) State() (*networkmanagerTransitGatewayRouteTableAttachmentState, bool) {
	return ntgrta.state, ntgrta.state != nil
}

// StateMust returns the state for [NetworkmanagerTransitGatewayRouteTableAttachment]. Panics if the state is nil.
func (ntgrta *NetworkmanagerTransitGatewayRouteTableAttachment) StateMust() *networkmanagerTransitGatewayRouteTableAttachmentState {
	if ntgrta.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ntgrta.Type(), ntgrta.LocalName()))
	}
	return ntgrta.state
}

// NetworkmanagerTransitGatewayRouteTableAttachmentArgs contains the configurations for aws_networkmanager_transit_gateway_route_table_attachment.
type NetworkmanagerTransitGatewayRouteTableAttachmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PeeringId: string, required
	PeeringId terra.StringValue `hcl:"peering_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TransitGatewayRouteTableArn: string, required
	TransitGatewayRouteTableArn terra.StringValue `hcl:"transit_gateway_route_table_arn,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *networkmanagertransitgatewayroutetableattachment.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerTransitGatewayRouteTableAttachmentAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("arn"))
}

// AttachmentPolicyRuleNumber returns a reference to field attachment_policy_rule_number of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) AttachmentPolicyRuleNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ntgrta.ref.Append("attachment_policy_rule_number"))
}

// AttachmentType returns a reference to field attachment_type of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) AttachmentType() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("attachment_type"))
}

// CoreNetworkArn returns a reference to field core_network_arn of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) CoreNetworkArn() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("core_network_arn"))
}

// CoreNetworkId returns a reference to field core_network_id of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) CoreNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("core_network_id"))
}

// EdgeLocation returns a reference to field edge_location of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) EdgeLocation() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("edge_location"))
}

// Id returns a reference to field id of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("id"))
}

// OwnerAccountId returns a reference to field owner_account_id of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) OwnerAccountId() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("owner_account_id"))
}

// PeeringId returns a reference to field peering_id of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) PeeringId() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("peering_id"))
}

// ResourceArn returns a reference to field resource_arn of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) ResourceArn() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("resource_arn"))
}

// SegmentName returns a reference to field segment_name of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) SegmentName() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("segment_name"))
}

// State returns a reference to field state of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ntgrta.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ntgrta.ref.Append("tags_all"))
}

// TransitGatewayRouteTableArn returns a reference to field transit_gateway_route_table_arn of aws_networkmanager_transit_gateway_route_table_attachment.
func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) TransitGatewayRouteTableArn() terra.StringValue {
	return terra.ReferenceAsString(ntgrta.ref.Append("transit_gateway_route_table_arn"))
}

func (ntgrta networkmanagerTransitGatewayRouteTableAttachmentAttributes) Timeouts() networkmanagertransitgatewayroutetableattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagertransitgatewayroutetableattachment.TimeoutsAttributes](ntgrta.ref.Append("timeouts"))
}

type networkmanagerTransitGatewayRouteTableAttachmentState struct {
	Arn                         string                                                          `json:"arn"`
	AttachmentPolicyRuleNumber  float64                                                         `json:"attachment_policy_rule_number"`
	AttachmentType              string                                                          `json:"attachment_type"`
	CoreNetworkArn              string                                                          `json:"core_network_arn"`
	CoreNetworkId               string                                                          `json:"core_network_id"`
	EdgeLocation                string                                                          `json:"edge_location"`
	Id                          string                                                          `json:"id"`
	OwnerAccountId              string                                                          `json:"owner_account_id"`
	PeeringId                   string                                                          `json:"peering_id"`
	ResourceArn                 string                                                          `json:"resource_arn"`
	SegmentName                 string                                                          `json:"segment_name"`
	State                       string                                                          `json:"state"`
	Tags                        map[string]string                                               `json:"tags"`
	TagsAll                     map[string]string                                               `json:"tags_all"`
	TransitGatewayRouteTableArn string                                                          `json:"transit_gateway_route_table_arn"`
	Timeouts                    *networkmanagertransitgatewayroutetableattachment.TimeoutsState `json:"timeouts"`
}
