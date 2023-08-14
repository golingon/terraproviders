// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computeroute "github.com/golingon/terraproviders/googlebeta/4.75.1/computeroute"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRoute creates a new instance of [ComputeRoute].
func NewComputeRoute(name string, args ComputeRouteArgs) *ComputeRoute {
	return &ComputeRoute{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRoute)(nil)

// ComputeRoute represents the Terraform resource google_compute_route.
type ComputeRoute struct {
	Name      string
	Args      ComputeRouteArgs
	state     *computeRouteState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRoute].
func (cr *ComputeRoute) Type() string {
	return "google_compute_route"
}

// LocalName returns the local name for [ComputeRoute].
func (cr *ComputeRoute) LocalName() string {
	return cr.Name
}

// Configuration returns the configuration (args) for [ComputeRoute].
func (cr *ComputeRoute) Configuration() interface{} {
	return cr.Args
}

// DependOn is used for other resources to depend on [ComputeRoute].
func (cr *ComputeRoute) DependOn() terra.Reference {
	return terra.ReferenceResource(cr)
}

// Dependencies returns the list of resources [ComputeRoute] depends_on.
func (cr *ComputeRoute) Dependencies() terra.Dependencies {
	return cr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRoute].
func (cr *ComputeRoute) LifecycleManagement() *terra.Lifecycle {
	return cr.Lifecycle
}

// Attributes returns the attributes for [ComputeRoute].
func (cr *ComputeRoute) Attributes() computeRouteAttributes {
	return computeRouteAttributes{ref: terra.ReferenceResource(cr)}
}

// ImportState imports the given attribute values into [ComputeRoute]'s state.
func (cr *ComputeRoute) ImportState(av io.Reader) error {
	cr.state = &computeRouteState{}
	if err := json.NewDecoder(av).Decode(cr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cr.Type(), cr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRoute] has state.
func (cr *ComputeRoute) State() (*computeRouteState, bool) {
	return cr.state, cr.state != nil
}

// StateMust returns the state for [ComputeRoute]. Panics if the state is nil.
func (cr *ComputeRoute) StateMust() *computeRouteState {
	if cr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cr.Type(), cr.LocalName()))
	}
	return cr.state
}

// ComputeRouteArgs contains the configurations for google_compute_route.
type ComputeRouteArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DestRange: string, required
	DestRange terra.StringValue `hcl:"dest_range,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, required
	Network terra.StringValue `hcl:"network,attr" validate:"required"`
	// NextHopGateway: string, optional
	NextHopGateway terra.StringValue `hcl:"next_hop_gateway,attr"`
	// NextHopIlb: string, optional
	NextHopIlb terra.StringValue `hcl:"next_hop_ilb,attr"`
	// NextHopInstance: string, optional
	NextHopInstance terra.StringValue `hcl:"next_hop_instance,attr"`
	// NextHopInstanceZone: string, optional
	NextHopInstanceZone terra.StringValue `hcl:"next_hop_instance_zone,attr"`
	// NextHopIp: string, optional
	NextHopIp terra.StringValue `hcl:"next_hop_ip,attr"`
	// NextHopVpnTunnel: string, optional
	NextHopVpnTunnel terra.StringValue `hcl:"next_hop_vpn_tunnel,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Tags: set of string, optional
	Tags terra.SetValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *computeroute.Timeouts `hcl:"timeouts,block"`
}
type computeRouteAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_compute_route.
func (cr computeRouteAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("description"))
}

// DestRange returns a reference to field dest_range of google_compute_route.
func (cr computeRouteAttributes) DestRange() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("dest_range"))
}

// Id returns a reference to field id of google_compute_route.
func (cr computeRouteAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_route.
func (cr computeRouteAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_route.
func (cr computeRouteAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("network"))
}

// NextHopGateway returns a reference to field next_hop_gateway of google_compute_route.
func (cr computeRouteAttributes) NextHopGateway() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_gateway"))
}

// NextHopIlb returns a reference to field next_hop_ilb of google_compute_route.
func (cr computeRouteAttributes) NextHopIlb() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_ilb"))
}

// NextHopInstance returns a reference to field next_hop_instance of google_compute_route.
func (cr computeRouteAttributes) NextHopInstance() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_instance"))
}

// NextHopInstanceZone returns a reference to field next_hop_instance_zone of google_compute_route.
func (cr computeRouteAttributes) NextHopInstanceZone() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_instance_zone"))
}

// NextHopIp returns a reference to field next_hop_ip of google_compute_route.
func (cr computeRouteAttributes) NextHopIp() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_ip"))
}

// NextHopNetwork returns a reference to field next_hop_network of google_compute_route.
func (cr computeRouteAttributes) NextHopNetwork() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_network"))
}

// NextHopVpnTunnel returns a reference to field next_hop_vpn_tunnel of google_compute_route.
func (cr computeRouteAttributes) NextHopVpnTunnel() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("next_hop_vpn_tunnel"))
}

// Priority returns a reference to field priority of google_compute_route.
func (cr computeRouteAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(cr.ref.Append("priority"))
}

// Project returns a reference to field project of google_compute_route.
func (cr computeRouteAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_route.
func (cr computeRouteAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cr.ref.Append("self_link"))
}

// Tags returns a reference to field tags of google_compute_route.
func (cr computeRouteAttributes) Tags() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cr.ref.Append("tags"))
}

func (cr computeRouteAttributes) Timeouts() computeroute.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeroute.TimeoutsAttributes](cr.ref.Append("timeouts"))
}

type computeRouteState struct {
	Description         string                      `json:"description"`
	DestRange           string                      `json:"dest_range"`
	Id                  string                      `json:"id"`
	Name                string                      `json:"name"`
	Network             string                      `json:"network"`
	NextHopGateway      string                      `json:"next_hop_gateway"`
	NextHopIlb          string                      `json:"next_hop_ilb"`
	NextHopInstance     string                      `json:"next_hop_instance"`
	NextHopInstanceZone string                      `json:"next_hop_instance_zone"`
	NextHopIp           string                      `json:"next_hop_ip"`
	NextHopNetwork      string                      `json:"next_hop_network"`
	NextHopVpnTunnel    string                      `json:"next_hop_vpn_tunnel"`
	Priority            float64                     `json:"priority"`
	Project             string                      `json:"project"`
	SelfLink            string                      `json:"self_link"`
	Tags                []string                    `json:"tags"`
	Timeouts            *computeroute.TimeoutsState `json:"timeouts"`
}