// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpc_peering_connection_accepter

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_vpc_peering_connection_accepter.
type Resource struct {
	Name      string
	Args      Args
	state     *awsVpcPeeringConnectionAccepterState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avpca *Resource) Type() string {
	return "aws_vpc_peering_connection_accepter"
}

// LocalName returns the local name for [Resource].
func (avpca *Resource) LocalName() string {
	return avpca.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avpca *Resource) Configuration() interface{} {
	return avpca.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avpca *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avpca)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avpca *Resource) Dependencies() terra.Dependencies {
	return avpca.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avpca *Resource) LifecycleManagement() *terra.Lifecycle {
	return avpca.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avpca *Resource) Attributes() awsVpcPeeringConnectionAccepterAttributes {
	return awsVpcPeeringConnectionAccepterAttributes{ref: terra.ReferenceResource(avpca)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avpca *Resource) ImportState(state io.Reader) error {
	avpca.state = &awsVpcPeeringConnectionAccepterState{}
	if err := json.NewDecoder(state).Decode(avpca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avpca.Type(), avpca.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avpca *Resource) State() (*awsVpcPeeringConnectionAccepterState, bool) {
	return avpca.state, avpca.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avpca *Resource) StateMust() *awsVpcPeeringConnectionAccepterState {
	if avpca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avpca.Type(), avpca.LocalName()))
	}
	return avpca.state
}

// Args contains the configurations for aws_vpc_peering_connection_accepter.
type Args struct {
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
	Accepter *Accepter `hcl:"accepter,block"`
	// Requester: optional
	Requester *Requester `hcl:"requester,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsVpcPeeringConnectionAccepterAttributes struct {
	ref terra.Reference
}

// AcceptStatus returns a reference to field accept_status of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) AcceptStatus() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("accept_status"))
}

// AutoAccept returns a reference to field auto_accept of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) AutoAccept() terra.BoolValue {
	return terra.ReferenceAsBool(avpca.ref.Append("auto_accept"))
}

// Id returns a reference to field id of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("id"))
}

// PeerOwnerId returns a reference to field peer_owner_id of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) PeerOwnerId() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("peer_owner_id"))
}

// PeerRegion returns a reference to field peer_region of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) PeerRegion() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("peer_region"))
}

// PeerVpcId returns a reference to field peer_vpc_id of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) PeerVpcId() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("peer_vpc_id"))
}

// Tags returns a reference to field tags of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avpca.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](avpca.ref.Append("tags_all"))
}

// VpcId returns a reference to field vpc_id of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("vpc_id"))
}

// VpcPeeringConnectionId returns a reference to field vpc_peering_connection_id of aws_vpc_peering_connection_accepter.
func (avpca awsVpcPeeringConnectionAccepterAttributes) VpcPeeringConnectionId() terra.StringValue {
	return terra.ReferenceAsString(avpca.ref.Append("vpc_peering_connection_id"))
}

func (avpca awsVpcPeeringConnectionAccepterAttributes) Accepter() terra.ListValue[AccepterAttributes] {
	return terra.ReferenceAsList[AccepterAttributes](avpca.ref.Append("accepter"))
}

func (avpca awsVpcPeeringConnectionAccepterAttributes) Requester() terra.ListValue[RequesterAttributes] {
	return terra.ReferenceAsList[RequesterAttributes](avpca.ref.Append("requester"))
}

func (avpca awsVpcPeeringConnectionAccepterAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](avpca.ref.Append("timeouts"))
}

type awsVpcPeeringConnectionAccepterState struct {
	AcceptStatus           string            `json:"accept_status"`
	AutoAccept             bool              `json:"auto_accept"`
	Id                     string            `json:"id"`
	PeerOwnerId            string            `json:"peer_owner_id"`
	PeerRegion             string            `json:"peer_region"`
	PeerVpcId              string            `json:"peer_vpc_id"`
	Tags                   map[string]string `json:"tags"`
	TagsAll                map[string]string `json:"tags_all"`
	VpcId                  string            `json:"vpc_id"`
	VpcPeeringConnectionId string            `json:"vpc_peering_connection_id"`
	Accepter               []AccepterState   `json:"accepter"`
	Requester              []RequesterState  `json:"requester"`
	Timeouts               *TimeoutsState    `json:"timeouts"`
}
