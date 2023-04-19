// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerconnection "github.com/golingon/terraproviders/aws/4.63.0/networkmanagerconnection"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerConnection creates a new instance of [NetworkmanagerConnection].
func NewNetworkmanagerConnection(name string, args NetworkmanagerConnectionArgs) *NetworkmanagerConnection {
	return &NetworkmanagerConnection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerConnection)(nil)

// NetworkmanagerConnection represents the Terraform resource aws_networkmanager_connection.
type NetworkmanagerConnection struct {
	Name      string
	Args      NetworkmanagerConnectionArgs
	state     *networkmanagerConnectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) Type() string {
	return "aws_networkmanager_connection"
}

// LocalName returns the local name for [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) LocalName() string {
	return nc.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) Configuration() interface{} {
	return nc.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) DependOn() terra.Reference {
	return terra.ReferenceResource(nc)
}

// Dependencies returns the list of resources [NetworkmanagerConnection] depends_on.
func (nc *NetworkmanagerConnection) Dependencies() terra.Dependencies {
	return nc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) LifecycleManagement() *terra.Lifecycle {
	return nc.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerConnection].
func (nc *NetworkmanagerConnection) Attributes() networkmanagerConnectionAttributes {
	return networkmanagerConnectionAttributes{ref: terra.ReferenceResource(nc)}
}

// ImportState imports the given attribute values into [NetworkmanagerConnection]'s state.
func (nc *NetworkmanagerConnection) ImportState(av io.Reader) error {
	nc.state = &networkmanagerConnectionState{}
	if err := json.NewDecoder(av).Decode(nc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nc.Type(), nc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerConnection] has state.
func (nc *NetworkmanagerConnection) State() (*networkmanagerConnectionState, bool) {
	return nc.state, nc.state != nil
}

// StateMust returns the state for [NetworkmanagerConnection]. Panics if the state is nil.
func (nc *NetworkmanagerConnection) StateMust() *networkmanagerConnectionState {
	if nc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nc.Type(), nc.LocalName()))
	}
	return nc.state
}

// NetworkmanagerConnectionArgs contains the configurations for aws_networkmanager_connection.
type NetworkmanagerConnectionArgs struct {
	// ConnectedDeviceId: string, required
	ConnectedDeviceId terra.StringValue `hcl:"connected_device_id,attr" validate:"required"`
	// ConnectedLinkId: string, optional
	ConnectedLinkId terra.StringValue `hcl:"connected_link_id,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DeviceId: string, required
	DeviceId terra.StringValue `hcl:"device_id,attr" validate:"required"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LinkId: string, optional
	LinkId terra.StringValue `hcl:"link_id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *networkmanagerconnection.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerConnectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("arn"))
}

// ConnectedDeviceId returns a reference to field connected_device_id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) ConnectedDeviceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("connected_device_id"))
}

// ConnectedLinkId returns a reference to field connected_link_id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) ConnectedLinkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("connected_link_id"))
}

// Description returns a reference to field description of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("description"))
}

// DeviceId returns a reference to field device_id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) DeviceId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("device_id"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("id"))
}

// LinkId returns a reference to field link_id of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) LinkId() terra.StringValue {
	return terra.ReferenceAsString(nc.ref.Append("link_id"))
}

// Tags returns a reference to field tags of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_connection.
func (nc networkmanagerConnectionAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nc.ref.Append("tags_all"))
}

func (nc networkmanagerConnectionAttributes) Timeouts() networkmanagerconnection.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerconnection.TimeoutsAttributes](nc.ref.Append("timeouts"))
}

type networkmanagerConnectionState struct {
	Arn               string                                  `json:"arn"`
	ConnectedDeviceId string                                  `json:"connected_device_id"`
	ConnectedLinkId   string                                  `json:"connected_link_id"`
	Description       string                                  `json:"description"`
	DeviceId          string                                  `json:"device_id"`
	GlobalNetworkId   string                                  `json:"global_network_id"`
	Id                string                                  `json:"id"`
	LinkId            string                                  `json:"link_id"`
	Tags              map[string]string                       `json:"tags"`
	TagsAll           map[string]string                       `json:"tags_all"`
	Timeouts          *networkmanagerconnection.TimeoutsState `json:"timeouts"`
}
