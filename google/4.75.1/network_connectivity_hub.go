// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	networkconnectivityhub "github.com/golingon/terraproviders/google/4.75.1/networkconnectivityhub"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewNetworkConnectivityHub creates a new instance of [NetworkConnectivityHub].
func NewNetworkConnectivityHub(name string, args NetworkConnectivityHubArgs) *NetworkConnectivityHub {
	return &NetworkConnectivityHub{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*NetworkConnectivityHub)(nil)

// NetworkConnectivityHub represents the Terraform resource google_network_connectivity_hub.
type NetworkConnectivityHub struct {
	Name      string
	Args      NetworkConnectivityHubArgs
	state     *networkConnectivityHubState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) Type() string {
	return "google_network_connectivity_hub"
}

// LocalName returns the local name for [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) LocalName() string {
	return nch.Name
}

// Configuration returns the configuration (args) for [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) Configuration() interface{} {
	return nch.Args
}

// DependOn is used for other resources to depend on [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) DependOn() terra.Reference {
	return terra.ReferenceResource(nch)
}

// Dependencies returns the list of resources [NetworkConnectivityHub] depends_on.
func (nch *NetworkConnectivityHub) Dependencies() terra.Dependencies {
	return nch.DependsOn
}

// LifecycleManagement returns the lifecycle block for [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) LifecycleManagement() *terra.Lifecycle {
	return nch.Lifecycle
}

// Attributes returns the attributes for [NetworkConnectivityHub].
func (nch *NetworkConnectivityHub) Attributes() networkConnectivityHubAttributes {
	return networkConnectivityHubAttributes{ref: terra.ReferenceResource(nch)}
}

// ImportState imports the given attribute values into [NetworkConnectivityHub]'s state.
func (nch *NetworkConnectivityHub) ImportState(av io.Reader) error {
	nch.state = &networkConnectivityHubState{}
	if err := json.NewDecoder(av).Decode(nch.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nch.Type(), nch.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [NetworkConnectivityHub] has state.
func (nch *NetworkConnectivityHub) State() (*networkConnectivityHubState, bool) {
	return nch.state, nch.state != nil
}

// StateMust returns the state for [NetworkConnectivityHub]. Panics if the state is nil.
func (nch *NetworkConnectivityHub) StateMust() *networkConnectivityHubState {
	if nch.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nch.Type(), nch.LocalName()))
	}
	return nch.state
}

// NetworkConnectivityHubArgs contains the configurations for google_network_connectivity_hub.
type NetworkConnectivityHubArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RoutingVpcs: min=0
	RoutingVpcs []networkconnectivityhub.RoutingVpcs `hcl:"routing_vpcs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *networkconnectivityhub.Timeouts `hcl:"timeouts,block"`
}
type networkConnectivityHubAttributes struct {
	ref terra.Reference
}

// CreateTime returns a reference to field create_time of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("create_time"))
}

// Description returns a reference to field description of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("description"))
}

// Id returns a reference to field id of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("id"))
}

// Labels returns a reference to field labels of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nch.ref.Append("labels"))
}

// Name returns a reference to field name of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("name"))
}

// Project returns a reference to field project of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("project"))
}

// State returns a reference to field state of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("state"))
}

// UniqueId returns a reference to field unique_id of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) UniqueId() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("unique_id"))
}

// UpdateTime returns a reference to field update_time of google_network_connectivity_hub.
func (nch networkConnectivityHubAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(nch.ref.Append("update_time"))
}

func (nch networkConnectivityHubAttributes) RoutingVpcs() terra.ListValue[networkconnectivityhub.RoutingVpcsAttributes] {
	return terra.ReferenceAsList[networkconnectivityhub.RoutingVpcsAttributes](nch.ref.Append("routing_vpcs"))
}

func (nch networkConnectivityHubAttributes) Timeouts() networkconnectivityhub.TimeoutsAttributes {
	return terra.ReferenceAsSingle[networkconnectivityhub.TimeoutsAttributes](nch.ref.Append("timeouts"))
}

type networkConnectivityHubState struct {
	CreateTime  string                                    `json:"create_time"`
	Description string                                    `json:"description"`
	Id          string                                    `json:"id"`
	Labels      map[string]string                         `json:"labels"`
	Name        string                                    `json:"name"`
	Project     string                                    `json:"project"`
	State       string                                    `json:"state"`
	UniqueId    string                                    `json:"unique_id"`
	UpdateTime  string                                    `json:"update_time"`
	RoutingVpcs []networkconnectivityhub.RoutingVpcsState `json:"routing_vpcs"`
	Timeouts    *networkconnectivityhub.TimeoutsState     `json:"timeouts"`
}
