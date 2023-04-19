// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagerglobalnetwork "github.com/golingon/terraproviders/aws/4.63.0/networkmanagerglobalnetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerGlobalNetwork creates a new instance of [NetworkmanagerGlobalNetwork].
func NewNetworkmanagerGlobalNetwork(name string, args NetworkmanagerGlobalNetworkArgs) *NetworkmanagerGlobalNetwork {
	return &NetworkmanagerGlobalNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerGlobalNetwork)(nil)

// NetworkmanagerGlobalNetwork represents the Terraform resource aws_networkmanager_global_network.
type NetworkmanagerGlobalNetwork struct {
	Name      string
	Args      NetworkmanagerGlobalNetworkArgs
	state     *networkmanagerGlobalNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) Type() string {
	return "aws_networkmanager_global_network"
}

// LocalName returns the local name for [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) LocalName() string {
	return ngn.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) Configuration() interface{} {
	return ngn.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(ngn)
}

// Dependencies returns the list of resources [NetworkmanagerGlobalNetwork] depends_on.
func (ngn *NetworkmanagerGlobalNetwork) Dependencies() terra.Dependencies {
	return ngn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) LifecycleManagement() *terra.Lifecycle {
	return ngn.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerGlobalNetwork].
func (ngn *NetworkmanagerGlobalNetwork) Attributes() networkmanagerGlobalNetworkAttributes {
	return networkmanagerGlobalNetworkAttributes{ref: terra.ReferenceResource(ngn)}
}

// ImportState imports the given attribute values into [NetworkmanagerGlobalNetwork]'s state.
func (ngn *NetworkmanagerGlobalNetwork) ImportState(av io.Reader) error {
	ngn.state = &networkmanagerGlobalNetworkState{}
	if err := json.NewDecoder(av).Decode(ngn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ngn.Type(), ngn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerGlobalNetwork] has state.
func (ngn *NetworkmanagerGlobalNetwork) State() (*networkmanagerGlobalNetworkState, bool) {
	return ngn.state, ngn.state != nil
}

// StateMust returns the state for [NetworkmanagerGlobalNetwork]. Panics if the state is nil.
func (ngn *NetworkmanagerGlobalNetwork) StateMust() *networkmanagerGlobalNetworkState {
	if ngn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ngn.Type(), ngn.LocalName()))
	}
	return ngn.state
}

// NetworkmanagerGlobalNetworkArgs contains the configurations for aws_networkmanager_global_network.
type NetworkmanagerGlobalNetworkArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *networkmanagerglobalnetwork.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerGlobalNetworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_global_network.
func (ngn networkmanagerGlobalNetworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkmanager_global_network.
func (ngn networkmanagerGlobalNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("description"))
}

// Id returns a reference to field id of aws_networkmanager_global_network.
func (ngn networkmanagerGlobalNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ngn.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_networkmanager_global_network.
func (ngn networkmanagerGlobalNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ngn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_global_network.
func (ngn networkmanagerGlobalNetworkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ngn.ref.Append("tags_all"))
}

func (ngn networkmanagerGlobalNetworkAttributes) Timeouts() networkmanagerglobalnetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagerglobalnetwork.TimeoutsAttributes](ngn.ref.Append("timeouts"))
}

type networkmanagerGlobalNetworkState struct {
	Arn         string                                     `json:"arn"`
	Description string                                     `json:"description"`
	Id          string                                     `json:"id"`
	Tags        map[string]string                          `json:"tags"`
	TagsAll     map[string]string                          `json:"tags_all"`
	Timeouts    *networkmanagerglobalnetwork.TimeoutsState `json:"timeouts"`
}
