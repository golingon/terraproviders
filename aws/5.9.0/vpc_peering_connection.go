// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	vpcpeeringconnection "github.com/golingon/terraproviders/aws/5.9.0/vpcpeeringconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewVpcPeeringConnection creates a new instance of [VpcPeeringConnection].
func NewVpcPeeringConnection(name string, args VpcPeeringConnectionArgs) *VpcPeeringConnection {
	return &VpcPeeringConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*VpcPeeringConnection)(nil)

// VpcPeeringConnection represents the Terraform resource aws_vpc_peering_connection.
type VpcPeeringConnection struct {
	Name      string
	Args      VpcPeeringConnectionArgs
	state     *vpcPeeringConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) Type() string {
	return "aws_vpc_peering_connection"
}

// LocalName returns the local name for [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) LocalName() string {
	return vpc.Name
}

// Configuration returns the configuration (args) for [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) Configuration() interface{} {
	return vpc.Args
}

// DependOn is used for other resources to depend on [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(vpc)
}

// Dependencies returns the list of resources [VpcPeeringConnection] depends_on.
func (vpc *VpcPeeringConnection) Dependencies() terra.Dependencies {
	return vpc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) LifecycleManagement() *terra.Lifecycle {
	return vpc.Lifecycle
}

// Attributes returns the attributes for [VpcPeeringConnection].
func (vpc *VpcPeeringConnection) Attributes() vpcPeeringConnectionAttributes {
	return vpcPeeringConnectionAttributes{ref: terra.ReferenceResource(vpc)}
}

// ImportState imports the given attribute values into [VpcPeeringConnection]'s state.
func (vpc *VpcPeeringConnection) ImportState(av io.Reader) error {
	vpc.state = &vpcPeeringConnectionState{}
	if err := json.NewDecoder(av).Decode(vpc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", vpc.Type(), vpc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [VpcPeeringConnection] has state.
func (vpc *VpcPeeringConnection) State() (*vpcPeeringConnectionState, bool) {
	return vpc.state, vpc.state != nil
}

// StateMust returns the state for [VpcPeeringConnection]. Panics if the state is nil.
func (vpc *VpcPeeringConnection) StateMust() *vpcPeeringConnectionState {
	if vpc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", vpc.Type(), vpc.LocalName()))
	}
	return vpc.state
}

// VpcPeeringConnectionArgs contains the configurations for aws_vpc_peering_connection.
type VpcPeeringConnectionArgs struct {
	// AutoAccept: bool, optional
	AutoAccept terra.BoolValue `hcl:"auto_accept,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// PeerOwnerId: string, optional
	PeerOwnerId terra.StringValue `hcl:"peer_owner_id,attr"`
	// PeerRegion: string, optional
	PeerRegion terra.StringValue `hcl:"peer_region,attr"`
	// PeerVpcId: string, required
	PeerVpcId terra.StringValue `hcl:"peer_vpc_id,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Accepter: optional
	Accepter *vpcpeeringconnection.Accepter `hcl:"accepter,block"`
	// Requester: optional
	Requester *vpcpeeringconnection.Requester `hcl:"requester,block"`
	// Timeouts: optional
	Timeouts *vpcpeeringconnection.Timeouts `hcl:"timeouts,block"`
}
type vpcPeeringConnectionAttributes struct {
	ref terra.Reference
}

// AcceptStatus returns a reference to field accept_status of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) AcceptStatus() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("accept_status"))
}

// AutoAccept returns a reference to field auto_accept of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) AutoAccept() terra.BoolValue {
	return terra.ReferenceAsBool(vpc.ref.Append("auto_accept"))
}

// Id returns a reference to field id of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("id"))
}

// PeerOwnerId returns a reference to field peer_owner_id of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) PeerOwnerId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_owner_id"))
}

// PeerRegion returns a reference to field peer_region of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_region"))
}

// PeerVpcId returns a reference to field peer_vpc_id of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) PeerVpcId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("peer_vpc_id"))
}

// Tags returns a reference to field tags of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](vpc.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_peering_connection.
func (vpc vpcPeeringConnectionAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(vpc.ref.Append("vpc_id"))
}

func (vpc vpcPeeringConnectionAttributes) Accepter() terra.ListValue[vpcpeeringconnection.AccepterAttributes] {
	return terra.ReferenceAsList[vpcpeeringconnection.AccepterAttributes](vpc.ref.Append("accepter"))
}

func (vpc vpcPeeringConnectionAttributes) Requester() terra.ListValue[vpcpeeringconnection.RequesterAttributes] {
	return terra.ReferenceAsList[vpcpeeringconnection.RequesterAttributes](vpc.ref.Append("requester"))
}

func (vpc vpcPeeringConnectionAttributes) Timeouts() vpcpeeringconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[vpcpeeringconnection.TimeoutsAttributes](vpc.ref.Append("timeouts"))
}

type vpcPeeringConnectionState struct {
	AcceptStatus string                                `json:"accept_status"`
	AutoAccept   bool                                  `json:"auto_accept"`
	Id           string                                `json:"id"`
	PeerOwnerId  string                                `json:"peer_owner_id"`
	PeerRegion   string                                `json:"peer_region"`
	PeerVpcId    string                                `json:"peer_vpc_id"`
	Tags         map[string]string                     `json:"tags"`
	TagsAll      map[string]string                     `json:"tags_all"`
	VpcId        string                                `json:"vpc_id"`
	Accepter     []vpcpeeringconnection.AccepterState  `json:"accepter"`
	Requester    []vpcpeeringconnection.RequesterState `json:"requester"`
	Timeouts     *vpcpeeringconnection.TimeoutsState   `json:"timeouts"`
}
