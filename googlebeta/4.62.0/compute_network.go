// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computenetwork "github.com/golingon/terraproviders/googlebeta/4.62.0/computenetwork"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeNetwork creates a new instance of [ComputeNetwork].
func NewComputeNetwork(name string, args ComputeNetworkArgs) *ComputeNetwork {
	return &ComputeNetwork{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeNetwork)(nil)

// ComputeNetwork represents the Terraform resource google_compute_network.
type ComputeNetwork struct {
	Name      string
	Args      ComputeNetworkArgs
	state     *computeNetworkState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeNetwork].
func (cn *ComputeNetwork) Type() string {
	return "google_compute_network"
}

// LocalName returns the local name for [ComputeNetwork].
func (cn *ComputeNetwork) LocalName() string {
	return cn.Name
}

// Configuration returns the configuration (args) for [ComputeNetwork].
func (cn *ComputeNetwork) Configuration() interface{} {
	return cn.Args
}

// DependOn is used for other resources to depend on [ComputeNetwork].
func (cn *ComputeNetwork) DependOn() terra.Reference {
	return terra.ReferenceResource(cn)
}

// Dependencies returns the list of resources [ComputeNetwork] depends_on.
func (cn *ComputeNetwork) Dependencies() terra.Dependencies {
	return cn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeNetwork].
func (cn *ComputeNetwork) LifecycleManagement() *terra.Lifecycle {
	return cn.Lifecycle
}

// Attributes returns the attributes for [ComputeNetwork].
func (cn *ComputeNetwork) Attributes() computeNetworkAttributes {
	return computeNetworkAttributes{ref: terra.ReferenceResource(cn)}
}

// ImportState imports the given attribute values into [ComputeNetwork]'s state.
func (cn *ComputeNetwork) ImportState(av io.Reader) error {
	cn.state = &computeNetworkState{}
	if err := json.NewDecoder(av).Decode(cn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cn.Type(), cn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeNetwork] has state.
func (cn *ComputeNetwork) State() (*computeNetworkState, bool) {
	return cn.state, cn.state != nil
}

// StateMust returns the state for [ComputeNetwork]. Panics if the state is nil.
func (cn *ComputeNetwork) StateMust() *computeNetworkState {
	if cn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cn.Type(), cn.LocalName()))
	}
	return cn.state
}

// ComputeNetworkArgs contains the configurations for google_compute_network.
type ComputeNetworkArgs struct {
	// AutoCreateSubnetworks: bool, optional
	AutoCreateSubnetworks terra.BoolValue `hcl:"auto_create_subnetworks,attr"`
	// DeleteDefaultRoutesOnCreate: bool, optional
	DeleteDefaultRoutesOnCreate terra.BoolValue `hcl:"delete_default_routes_on_create,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnableUlaInternalIpv6: bool, optional
	EnableUlaInternalIpv6 terra.BoolValue `hcl:"enable_ula_internal_ipv6,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InternalIpv6Range: string, optional
	InternalIpv6Range terra.StringValue `hcl:"internal_ipv6_range,attr"`
	// Mtu: number, optional
	Mtu terra.NumberValue `hcl:"mtu,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NetworkFirewallPolicyEnforcementOrder: string, optional
	NetworkFirewallPolicyEnforcementOrder terra.StringValue `hcl:"network_firewall_policy_enforcement_order,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RoutingMode: string, optional
	RoutingMode terra.StringValue `hcl:"routing_mode,attr"`
	// Timeouts: optional
	Timeouts *computenetwork.Timeouts `hcl:"timeouts,block"`
}
type computeNetworkAttributes struct {
	ref terra.Reference
}

// AutoCreateSubnetworks returns a reference to field auto_create_subnetworks of google_compute_network.
func (cn computeNetworkAttributes) AutoCreateSubnetworks() terra.BoolValue {
	return terra.ReferenceAsBool(cn.ref.Append("auto_create_subnetworks"))
}

// DeleteDefaultRoutesOnCreate returns a reference to field delete_default_routes_on_create of google_compute_network.
func (cn computeNetworkAttributes) DeleteDefaultRoutesOnCreate() terra.BoolValue {
	return terra.ReferenceAsBool(cn.ref.Append("delete_default_routes_on_create"))
}

// Description returns a reference to field description of google_compute_network.
func (cn computeNetworkAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("description"))
}

// EnableUlaInternalIpv6 returns a reference to field enable_ula_internal_ipv6 of google_compute_network.
func (cn computeNetworkAttributes) EnableUlaInternalIpv6() terra.BoolValue {
	return terra.ReferenceAsBool(cn.ref.Append("enable_ula_internal_ipv6"))
}

// GatewayIpv4 returns a reference to field gateway_ipv4 of google_compute_network.
func (cn computeNetworkAttributes) GatewayIpv4() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("gateway_ipv4"))
}

// Id returns a reference to field id of google_compute_network.
func (cn computeNetworkAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("id"))
}

// InternalIpv6Range returns a reference to field internal_ipv6_range of google_compute_network.
func (cn computeNetworkAttributes) InternalIpv6Range() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("internal_ipv6_range"))
}

// Mtu returns a reference to field mtu of google_compute_network.
func (cn computeNetworkAttributes) Mtu() terra.NumberValue {
	return terra.ReferenceAsNumber(cn.ref.Append("mtu"))
}

// Name returns a reference to field name of google_compute_network.
func (cn computeNetworkAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("name"))
}

// NetworkFirewallPolicyEnforcementOrder returns a reference to field network_firewall_policy_enforcement_order of google_compute_network.
func (cn computeNetworkAttributes) NetworkFirewallPolicyEnforcementOrder() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("network_firewall_policy_enforcement_order"))
}

// Project returns a reference to field project of google_compute_network.
func (cn computeNetworkAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("project"))
}

// RoutingMode returns a reference to field routing_mode of google_compute_network.
func (cn computeNetworkAttributes) RoutingMode() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("routing_mode"))
}

// SelfLink returns a reference to field self_link of google_compute_network.
func (cn computeNetworkAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cn.ref.Append("self_link"))
}

func (cn computeNetworkAttributes) Timeouts() computenetwork.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computenetwork.TimeoutsAttributes](cn.ref.Append("timeouts"))
}

type computeNetworkState struct {
	AutoCreateSubnetworks                 bool                          `json:"auto_create_subnetworks"`
	DeleteDefaultRoutesOnCreate           bool                          `json:"delete_default_routes_on_create"`
	Description                           string                        `json:"description"`
	EnableUlaInternalIpv6                 bool                          `json:"enable_ula_internal_ipv6"`
	GatewayIpv4                           string                        `json:"gateway_ipv4"`
	Id                                    string                        `json:"id"`
	InternalIpv6Range                     string                        `json:"internal_ipv6_range"`
	Mtu                                   float64                       `json:"mtu"`
	Name                                  string                        `json:"name"`
	NetworkFirewallPolicyEnforcementOrder string                        `json:"network_firewall_policy_enforcement_order"`
	Project                               string                        `json:"project"`
	RoutingMode                           string                        `json:"routing_mode"`
	SelfLink                              string                        `json:"self_link"`
	Timeouts                              *computenetwork.TimeoutsState `json:"timeouts"`
}
