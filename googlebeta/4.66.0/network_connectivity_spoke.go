// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	networkconnectivityspoke "github.com/golingon/terraproviders/googlebeta/4.66.0/networkconnectivityspoke"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkConnectivitySpoke creates a new instance of [NetworkConnectivitySpoke].
func NewNetworkConnectivitySpoke(name string, args NetworkConnectivitySpokeArgs) *NetworkConnectivitySpoke {
	return &NetworkConnectivitySpoke{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkConnectivitySpoke)(nil)

// NetworkConnectivitySpoke represents the Terraform resource google_network_connectivity_spoke.
type NetworkConnectivitySpoke struct {
	Name      string
	Args      NetworkConnectivitySpokeArgs
	state     *networkConnectivitySpokeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) Type() string {
	return "google_network_connectivity_spoke"
}

// LocalName returns the local name for [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) LocalName() string {
	return ncs.Name
}

// Configuration returns the configuration (args) for [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) Configuration() interface{} {
	return ncs.Args
}

// DependOn is used for other resources to depend on [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) DependOn() terra.Reference {
	return terra.ReferenceResource(ncs)
}

// Dependencies returns the list of resources [NetworkConnectivitySpoke] depends_on.
func (ncs *NetworkConnectivitySpoke) Dependencies() terra.Dependencies {
	return ncs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) LifecycleManagement() *terra.Lifecycle {
	return ncs.Lifecycle
}

// Attributes returns the attributes for [NetworkConnectivitySpoke].
func (ncs *NetworkConnectivitySpoke) Attributes() networkConnectivitySpokeAttributes {
	return networkConnectivitySpokeAttributes{ref: terra.ReferenceResource(ncs)}
}

// ImportState imports the given attribute values into [NetworkConnectivitySpoke]'s state.
func (ncs *NetworkConnectivitySpoke) ImportState(av io.Reader) error {
	ncs.state = &networkConnectivitySpokeState{}
	if err := json.NewDecoder(av).Decode(ncs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ncs.Type(), ncs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkConnectivitySpoke] has state.
func (ncs *NetworkConnectivitySpoke) State() (*networkConnectivitySpokeState, bool) {
	return ncs.state, ncs.state != nil
}

// StateMust returns the state for [NetworkConnectivitySpoke]. Panics if the state is nil.
func (ncs *NetworkConnectivitySpoke) StateMust() *networkConnectivitySpokeState {
	if ncs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ncs.Type(), ncs.LocalName()))
	}
	return ncs.state
}

// NetworkConnectivitySpokeArgs contains the configurations for google_network_connectivity_spoke.
type NetworkConnectivitySpokeArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Hub: string, required
	Hub terra.StringValue `hcl:"hub,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// LinkedInterconnectAttachments: optional
	LinkedInterconnectAttachments *networkconnectivityspoke.LinkedInterconnectAttachments `hcl:"linked_interconnect_attachments,block"`
	// LinkedRouterApplianceInstances: optional
	LinkedRouterApplianceInstances *networkconnectivityspoke.LinkedRouterApplianceInstances `hcl:"linked_router_appliance_instances,block"`
	// LinkedVpnTunnels: optional
	LinkedVpnTunnels *networkconnectivityspoke.LinkedVpnTunnels `hcl:"linked_vpn_tunnels,block"`
	// Timeouts: optional
	Timeouts *networkconnectivityspoke.Timeouts `hcl:"timeouts,block"`
}
type networkConnectivitySpokeAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("description"))
}

// Hub returns a reference to field hub of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Hub() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("hub"))
}

// Id returns a reference to field id of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ncs.ref.Append("labels"))
}

// Location returns a reference to field location of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("location"))
}

// Name returns a reference to field name of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("name"))
}

// Project returns a reference to field project of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("project"))
}

// State returns a reference to field state of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("state"))
}

// UniqueId returns a reference to field unique_id of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("unique_id"))
}

// UpdateTime returns a reference to field update_time of google_network_connectivity_spoke.
func (ncs networkConnectivitySpokeAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(ncs.ref.Append("update_time"))
}

func (ncs networkConnectivitySpokeAttributes) LinkedInterconnectAttachments() terra.ListValue[networkconnectivityspoke.LinkedInterconnectAttachmentsAttributes] {
	return terra.ReferenceAsList[networkconnectivityspoke.LinkedInterconnectAttachmentsAttributes](ncs.ref.Append("linked_interconnect_attachments"))
}

func (ncs networkConnectivitySpokeAttributes) LinkedRouterApplianceInstances() terra.ListValue[networkconnectivityspoke.LinkedRouterApplianceInstancesAttributes] {
	return terra.ReferenceAsList[networkconnectivityspoke.LinkedRouterApplianceInstancesAttributes](ncs.ref.Append("linked_router_appliance_instances"))
}

func (ncs networkConnectivitySpokeAttributes) LinkedVpnTunnels() terra.ListValue[networkconnectivityspoke.LinkedVpnTunnelsAttributes] {
	return terra.ReferenceAsList[networkconnectivityspoke.LinkedVpnTunnelsAttributes](ncs.ref.Append("linked_vpn_tunnels"))
}

func (ncs networkConnectivitySpokeAttributes) Timeouts() networkconnectivityspoke.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkconnectivityspoke.TimeoutsAttributes](ncs.ref.Append("timeouts"))
}

type networkConnectivitySpokeState struct {
	CreateTime                     string                                                         `json:"create_time"`
	Description                    string                                                         `json:"description"`
	Hub                            string                                                         `json:"hub"`
	Id                             string                                                         `json:"id"`
	Labels                         map[string]string                                              `json:"labels"`
	Location                       string                                                         `json:"location"`
	Name                           string                                                         `json:"name"`
	Project                        string                                                         `json:"project"`
	State                          string                                                         `json:"state"`
	UniqueId                       string                                                         `json:"unique_id"`
	UpdateTime                     string                                                         `json:"update_time"`
	LinkedInterconnectAttachments  []networkconnectivityspoke.LinkedInterconnectAttachmentsState  `json:"linked_interconnect_attachments"`
	LinkedRouterApplianceInstances []networkconnectivityspoke.LinkedRouterApplianceInstancesState `json:"linked_router_appliance_instances"`
	LinkedVpnTunnels               []networkconnectivityspoke.LinkedVpnTunnelsState               `json:"linked_vpn_tunnels"`
	Timeouts                       *networkconnectivityspoke.TimeoutsState                        `json:"timeouts"`
}
