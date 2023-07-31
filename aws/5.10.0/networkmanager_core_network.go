// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	networkmanagercorenetwork "github.com/golingon/terraproviders/aws/5.10.0/networkmanagercorenetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkmanagerCoreNetwork creates a new instance of [NetworkmanagerCoreNetwork].
func NewNetworkmanagerCoreNetwork(name string, args NetworkmanagerCoreNetworkArgs) *NetworkmanagerCoreNetwork {
	return &NetworkmanagerCoreNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkmanagerCoreNetwork)(nil)

// NetworkmanagerCoreNetwork represents the Terraform resource aws_networkmanager_core_network.
type NetworkmanagerCoreNetwork struct {
	Name      string
	Args      NetworkmanagerCoreNetworkArgs
	state     *networkmanagerCoreNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) Type() string {
	return "aws_networkmanager_core_network"
}

// LocalName returns the local name for [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) LocalName() string {
	return ncn.Name
}

// Configuration returns the configuration (args) for [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) Configuration() interface{} {
	return ncn.Args
}

// DependOn is used for other resources to depend on [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(ncn)
}

// Dependencies returns the list of resources [NetworkmanagerCoreNetwork] depends_on.
func (ncn *NetworkmanagerCoreNetwork) Dependencies() terra.Dependencies {
	return ncn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) LifecycleManagement() *terra.Lifecycle {
	return ncn.Lifecycle
}

// Attributes returns the attributes for [NetworkmanagerCoreNetwork].
func (ncn *NetworkmanagerCoreNetwork) Attributes() networkmanagerCoreNetworkAttributes {
	return networkmanagerCoreNetworkAttributes{ref: terra.ReferenceResource(ncn)}
}

// ImportState imports the given attribute values into [NetworkmanagerCoreNetwork]'s state.
func (ncn *NetworkmanagerCoreNetwork) ImportState(av io.Reader) error {
	ncn.state = &networkmanagerCoreNetworkState{}
	if err := json.NewDecoder(av).Decode(ncn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncn.Type(), ncn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkmanagerCoreNetwork] has state.
func (ncn *NetworkmanagerCoreNetwork) State() (*networkmanagerCoreNetworkState, bool) {
	return ncn.state, ncn.state != nil
}

// StateMust returns the state for [NetworkmanagerCoreNetwork]. Panics if the state is nil.
func (ncn *NetworkmanagerCoreNetwork) StateMust() *networkmanagerCoreNetworkState {
	if ncn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncn.Type(), ncn.LocalName()))
	}
	return ncn.state
}

// NetworkmanagerCoreNetworkArgs contains the configurations for aws_networkmanager_core_network.
type NetworkmanagerCoreNetworkArgs struct {
	// BasePolicyRegion: string, optional
	BasePolicyRegion terra.StringValue `hcl:"base_policy_region,attr"`
	// BasePolicyRegions: set of string, optional
	BasePolicyRegions terra.SetValue[terra.StringValue] `hcl:"base_policy_regions,attr"`
	// CreateBasePolicy: bool, optional
	CreateBasePolicy terra.BoolValue `hcl:"create_base_policy,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// GlobalNetworkId: string, required
	GlobalNetworkId terra.StringValue `hcl:"global_network_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Edges: min=0
	Edges []networkmanagercorenetwork.Edges `hcl:"edges,block" validate:"min=0"`
	// Segments: min=0
	Segments []networkmanagercorenetwork.Segments `hcl:"segments,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *networkmanagercorenetwork.Timeouts `hcl:"timeouts,block"`
}
type networkmanagerCoreNetworkAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("arn"))
}

// BasePolicyRegion returns a reference to field base_policy_region of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) BasePolicyRegion() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("base_policy_region"))
}

// BasePolicyRegions returns a reference to field base_policy_regions of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) BasePolicyRegions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ncn.ref.Append("base_policy_regions"))
}

// CreateBasePolicy returns a reference to field create_base_policy of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) CreateBasePolicy() terra.BoolValue {
	return terra.ReferenceAsBool(ncn.ref.Append("create_base_policy"))
}

// CreatedAt returns a reference to field created_at of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("created_at"))
}

// Description returns a reference to field description of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("description"))
}

// GlobalNetworkId returns a reference to field global_network_id of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) GlobalNetworkId() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("global_network_id"))
}

// Id returns a reference to field id of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("id"))
}

// State returns a reference to field state of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ncn.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncn.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_networkmanager_core_network.
func (ncn networkmanagerCoreNetworkAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncn.ref.Append("tags_all"))
}

func (ncn networkmanagerCoreNetworkAttributes) Edges() terra.ListValue[networkmanagercorenetwork.EdgesAttributes] {
	return terra.ReferenceAsList[networkmanagercorenetwork.EdgesAttributes](ncn.ref.Append("edges"))
}

func (ncn networkmanagerCoreNetworkAttributes) Segments() terra.ListValue[networkmanagercorenetwork.SegmentsAttributes] {
	return terra.ReferenceAsList[networkmanagercorenetwork.SegmentsAttributes](ncn.ref.Append("segments"))
}

func (ncn networkmanagerCoreNetworkAttributes) Timeouts() networkmanagercorenetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkmanagercorenetwork.TimeoutsAttributes](ncn.ref.Append("timeouts"))
}

type networkmanagerCoreNetworkState struct {
	Arn               string                                    `json:"arn"`
	BasePolicyRegion  string                                    `json:"base_policy_region"`
	BasePolicyRegions []string                                  `json:"base_policy_regions"`
	CreateBasePolicy  bool                                      `json:"create_base_policy"`
	CreatedAt         string                                    `json:"created_at"`
	Description       string                                    `json:"description"`
	GlobalNetworkId   string                                    `json:"global_network_id"`
	Id                string                                    `json:"id"`
	State             string                                    `json:"state"`
	Tags              map[string]string                         `json:"tags"`
	TagsAll           map[string]string                         `json:"tags_all"`
	Edges             []networkmanagercorenetwork.EdgesState    `json:"edges"`
	Segments          []networkmanagercorenetwork.SegmentsState `json:"segments"`
	Timeouts          *networkmanagercorenetwork.TimeoutsState  `json:"timeouts"`
}
