// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcpeeringconnectionaccepter "github.com/golingon/terraproviders/aws/5.8.0/vpcpeeringconnectionaccepter"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcPeeringConnectionAccepter creates a new instance of [VpcPeeringConnectionAccepter].
func NewVpcPeeringConnectionAccepter(name string, args VpcPeeringConnectionAccepterArgs) *VpcPeeringConnectionAccepter {
	return &VpcPeeringConnectionAccepter{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcPeeringConnectionAccepter)(nil)

// VpcPeeringConnectionAccepter represents the Terraform resource aws_vpc_peering_connection_accepter.
type VpcPeeringConnectionAccepter struct {
	Name      string
	Args      VpcPeeringConnectionAccepterArgs
	state     *vpcPeeringConnectionAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) Type() string {
	return "aws_vpc_peering_connection_accepter"
}

// LocalName returns the local name for [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) LocalName() string {
	return vpca.Name
}

// Configuration returns the configuration (args) for [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) Configuration() interface{} {
	return vpca.Args
}

// DependOn is used for other resources to depend on [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) DependOn() terra.Reference {
	return terra.ReferenceResource(vpca)
}

// Dependencies returns the list of resources [VpcPeeringConnectionAccepter] depends_on.
func (vpca *VpcPeeringConnectionAccepter) Dependencies() terra.Dependencies {
	return vpca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) LifecycleManagement() *terra.Lifecycle {
	return vpca.Lifecycle
}

// Attributes returns the attributes for [VpcPeeringConnectionAccepter].
func (vpca *VpcPeeringConnectionAccepter) Attributes() vpcPeeringConnectionAccepterAttributes {
	return vpcPeeringConnectionAccepterAttributes{ref: terra.ReferenceResource(vpca)}
}

// ImportState imports the given attribute values into [VpcPeeringConnectionAccepter]'s state.
func (vpca *VpcPeeringConnectionAccepter) ImportState(av io.Reader) error {
	vpca.state = &vpcPeeringConnectionAccepterState{}
	if err := json.NewDecoder(av).Decode(vpca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpca.Type(), vpca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcPeeringConnectionAccepter] has state.
func (vpca *VpcPeeringConnectionAccepter) State() (*vpcPeeringConnectionAccepterState, bool) {
	return vpca.state, vpca.state != nil
}

// StateMust returns the state for [VpcPeeringConnectionAccepter]. Panics if the state is nil.
func (vpca *VpcPeeringConnectionAccepter) StateMust() *vpcPeeringConnectionAccepterState {
	if vpca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpca.Type(), vpca.LocalName()))
	}
	return vpca.state
}

// VpcPeeringConnectionAccepterArgs contains the configurations for aws_vpc_peering_connection_accepter.
type VpcPeeringConnectionAccepterArgs struct {
	// AutoAccept: bool, optional
	AutoAccept terra.BoolValue `hcl:"auto_accept,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcPeeringConnectionId: string, required
	VpcPeeringConnectionId terra.StringValue `hcl:"vpc_peering_connection_id,attr" validate:"required"`
	// Accepter: optional
	Accepter *vpcpeeringconnectionaccepter.Accepter `hcl:"accepter,block"`
	// Requester: optional
	Requester *vpcpeeringconnectionaccepter.Requester `hcl:"requester,block"`
	// Timeouts: optional
	Timeouts *vpcpeeringconnectionaccepter.Timeouts `hcl:"timeouts,block"`
}
type vpcPeeringConnectionAccepterAttributes struct {
	ref terra.Reference
}

// AcceptStatus returns a reference to field accept_status of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) AcceptStatus() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("accept_status"))
}

// AutoAccept returns a reference to field auto_accept of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) AutoAccept() terra.BoolValue {
	return terra.ReferenceAsBool(vpca.ref.Append("auto_accept"))
}

// Id returns a reference to field id of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("id"))
}

// PeerOwnerId returns a reference to field peer_owner_id of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) PeerOwnerId() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("peer_owner_id"))
}

// PeerRegion returns a reference to field peer_region of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("peer_region"))
}

// PeerVpcId returns a reference to field peer_vpc_id of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) PeerVpcId() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("peer_vpc_id"))
}

// Tags returns a reference to field tags of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpca.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("vpc_id"))
}

// VpcPeeringConnectionId returns a reference to field vpc_peering_connection_id of aws_vpc_peering_connection_accepter.
func (vpca vpcPeeringConnectionAccepterAttributes) VpcPeeringConnectionId() terra.StringValue {
	return terra.ReferenceAsString(vpca.ref.Append("vpc_peering_connection_id"))
}

func (vpca vpcPeeringConnectionAccepterAttributes) Accepter() terra.ListValue[vpcpeeringconnectionaccepter.AccepterAttributes] {
	return terra.ReferenceAsList[vpcpeeringconnectionaccepter.AccepterAttributes](vpca.ref.Append("accepter"))
}

func (vpca vpcPeeringConnectionAccepterAttributes) Requester() terra.ListValue[vpcpeeringconnectionaccepter.RequesterAttributes] {
	return terra.ReferenceAsList[vpcpeeringconnectionaccepter.RequesterAttributes](vpca.ref.Append("requester"))
}

func (vpca vpcPeeringConnectionAccepterAttributes) Timeouts() vpcpeeringconnectionaccepter.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcpeeringconnectionaccepter.TimeoutsAttributes](vpca.ref.Append("timeouts"))
}

type vpcPeeringConnectionAccepterState struct {
	AcceptStatus           string                                        `json:"accept_status"`
	AutoAccept             bool                                          `json:"auto_accept"`
	Id                     string                                        `json:"id"`
	PeerOwnerId            string                                        `json:"peer_owner_id"`
	PeerRegion             string                                        `json:"peer_region"`
	PeerVpcId              string                                        `json:"peer_vpc_id"`
	Tags                   map[string]string                             `json:"tags"`
	TagsAll                map[string]string                             `json:"tags_all"`
	VpcId                  string                                        `json:"vpc_id"`
	VpcPeeringConnectionId string                                        `json:"vpc_peering_connection_id"`
	Accepter               []vpcpeeringconnectionaccepter.AccepterState  `json:"accepter"`
	Requester              []vpcpeeringconnectionaccepter.RequesterState `json:"requester"`
	Timeouts               *vpcpeeringconnectionaccepter.TimeoutsState   `json:"timeouts"`
}
