// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computerouterpeer "github.com/golingon/terraproviders/google/4.66.0/computerouterpeer"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRouterPeer creates a new instance of [ComputeRouterPeer].
func NewComputeRouterPeer(name string, args ComputeRouterPeerArgs) *ComputeRouterPeer {
	return &ComputeRouterPeer{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRouterPeer)(nil)

// ComputeRouterPeer represents the Terraform resource google_compute_router_peer.
type ComputeRouterPeer struct {
	Name      string
	Args      ComputeRouterPeerArgs
	state     *computeRouterPeerState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRouterPeer].
func (crp *ComputeRouterPeer) Type() string {
	return "google_compute_router_peer"
}

// LocalName returns the local name for [ComputeRouterPeer].
func (crp *ComputeRouterPeer) LocalName() string {
	return crp.Name
}

// Configuration returns the configuration (args) for [ComputeRouterPeer].
func (crp *ComputeRouterPeer) Configuration() interface{} {
	return crp.Args
}

// DependOn is used for other resources to depend on [ComputeRouterPeer].
func (crp *ComputeRouterPeer) DependOn() terra.Reference {
	return terra.ReferenceResource(crp)
}

// Dependencies returns the list of resources [ComputeRouterPeer] depends_on.
func (crp *ComputeRouterPeer) Dependencies() terra.Dependencies {
	return crp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRouterPeer].
func (crp *ComputeRouterPeer) LifecycleManagement() *terra.Lifecycle {
	return crp.Lifecycle
}

// Attributes returns the attributes for [ComputeRouterPeer].
func (crp *ComputeRouterPeer) Attributes() computeRouterPeerAttributes {
	return computeRouterPeerAttributes{ref: terra.ReferenceResource(crp)}
}

// ImportState imports the given attribute values into [ComputeRouterPeer]'s state.
func (crp *ComputeRouterPeer) ImportState(av io.Reader) error {
	crp.state = &computeRouterPeerState{}
	if err := json.NewDecoder(av).Decode(crp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crp.Type(), crp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRouterPeer] has state.
func (crp *ComputeRouterPeer) State() (*computeRouterPeerState, bool) {
	return crp.state, crp.state != nil
}

// StateMust returns the state for [ComputeRouterPeer]. Panics if the state is nil.
func (crp *ComputeRouterPeer) StateMust() *computeRouterPeerState {
	if crp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crp.Type(), crp.LocalName()))
	}
	return crp.state
}

// ComputeRouterPeerArgs contains the configurations for google_compute_router_peer.
type ComputeRouterPeerArgs struct {
	// AdvertiseMode: string, optional
	AdvertiseMode terra.StringValue `hcl:"advertise_mode,attr"`
	// AdvertisedGroups: list of string, optional
	AdvertisedGroups terra.ListValue[terra.StringValue] `hcl:"advertised_groups,attr"`
	// AdvertisedRoutePriority: number, optional
	AdvertisedRoutePriority terra.NumberValue `hcl:"advertised_route_priority,attr"`
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
	// EnableIpv6: bool, optional
	EnableIpv6 terra.BoolValue `hcl:"enable_ipv6,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interface: string, required
	Interface terra.StringValue `hcl:"interface,attr" validate:"required"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// Ipv6NexthopAddress: string, optional
	Ipv6NexthopAddress terra.StringValue `hcl:"ipv6_nexthop_address,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PeerAsn: number, required
	PeerAsn terra.NumberValue `hcl:"peer_asn,attr" validate:"required"`
	// PeerIpAddress: string, required
	PeerIpAddress terra.StringValue `hcl:"peer_ip_address,attr" validate:"required"`
	// PeerIpv6NexthopAddress: string, optional
	PeerIpv6NexthopAddress terra.StringValue `hcl:"peer_ipv6_nexthop_address,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Router: string, required
	Router terra.StringValue `hcl:"router,attr" validate:"required"`
	// RouterApplianceInstance: string, optional
	RouterApplianceInstance terra.StringValue `hcl:"router_appliance_instance,attr"`
	// AdvertisedIpRanges: min=0
	AdvertisedIpRanges []computerouterpeer.AdvertisedIpRanges `hcl:"advertised_ip_ranges,block" validate:"min=0"`
	// Bfd: optional
	Bfd *computerouterpeer.Bfd `hcl:"bfd,block"`
	// Timeouts: optional
	Timeouts *computerouterpeer.Timeouts `hcl:"timeouts,block"`
}
type computeRouterPeerAttributes struct {
	ref terra.Reference
}

// AdvertiseMode returns a reference to field advertise_mode of google_compute_router_peer.
func (crp computeRouterPeerAttributes) AdvertiseMode() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("advertise_mode"))
}

// AdvertisedGroups returns a reference to field advertised_groups of google_compute_router_peer.
func (crp computeRouterPeerAttributes) AdvertisedGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](crp.ref.Append("advertised_groups"))
}

// AdvertisedRoutePriority returns a reference to field advertised_route_priority of google_compute_router_peer.
func (crp computeRouterPeerAttributes) AdvertisedRoutePriority() terra.NumberValue {
	return terra.ReferenceAsNumber(crp.ref.Append("advertised_route_priority"))
}

// Enable returns a reference to field enable of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(crp.ref.Append("enable"))
}

// EnableIpv6 returns a reference to field enable_ipv6 of google_compute_router_peer.
func (crp computeRouterPeerAttributes) EnableIpv6() terra.BoolValue {
	return terra.ReferenceAsBool(crp.ref.Append("enable_ipv6"))
}

// Id returns a reference to field id of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("id"))
}

// Interface returns a reference to field interface of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Interface() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("interface"))
}

// IpAddress returns a reference to field ip_address of google_compute_router_peer.
func (crp computeRouterPeerAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("ip_address"))
}

// Ipv6NexthopAddress returns a reference to field ipv6_nexthop_address of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Ipv6NexthopAddress() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("ipv6_nexthop_address"))
}

// ManagementType returns a reference to field management_type of google_compute_router_peer.
func (crp computeRouterPeerAttributes) ManagementType() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("management_type"))
}

// Name returns a reference to field name of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("name"))
}

// PeerAsn returns a reference to field peer_asn of google_compute_router_peer.
func (crp computeRouterPeerAttributes) PeerAsn() terra.NumberValue {
	return terra.ReferenceAsNumber(crp.ref.Append("peer_asn"))
}

// PeerIpAddress returns a reference to field peer_ip_address of google_compute_router_peer.
func (crp computeRouterPeerAttributes) PeerIpAddress() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("peer_ip_address"))
}

// PeerIpv6NexthopAddress returns a reference to field peer_ipv6_nexthop_address of google_compute_router_peer.
func (crp computeRouterPeerAttributes) PeerIpv6NexthopAddress() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("peer_ipv6_nexthop_address"))
}

// Project returns a reference to field project of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_router_peer.
func (crp computeRouterPeerAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("router"))
}

// RouterApplianceInstance returns a reference to field router_appliance_instance of google_compute_router_peer.
func (crp computeRouterPeerAttributes) RouterApplianceInstance() terra.StringValue {
	return terra.ReferenceAsString(crp.ref.Append("router_appliance_instance"))
}

func (crp computeRouterPeerAttributes) AdvertisedIpRanges() terra.ListValue[computerouterpeer.AdvertisedIpRangesAttributes] {
	return terra.ReferenceAsList[computerouterpeer.AdvertisedIpRangesAttributes](crp.ref.Append("advertised_ip_ranges"))
}

func (crp computeRouterPeerAttributes) Bfd() terra.ListValue[computerouterpeer.BfdAttributes] {
	return terra.ReferenceAsList[computerouterpeer.BfdAttributes](crp.ref.Append("bfd"))
}

func (crp computeRouterPeerAttributes) Timeouts() computerouterpeer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computerouterpeer.TimeoutsAttributes](crp.ref.Append("timeouts"))
}

type computeRouterPeerState struct {
	AdvertiseMode           string                                      `json:"advertise_mode"`
	AdvertisedGroups        []string                                    `json:"advertised_groups"`
	AdvertisedRoutePriority float64                                     `json:"advertised_route_priority"`
	Enable                  bool                                        `json:"enable"`
	EnableIpv6              bool                                        `json:"enable_ipv6"`
	Id                      string                                      `json:"id"`
	Interface               string                                      `json:"interface"`
	IpAddress               string                                      `json:"ip_address"`
	Ipv6NexthopAddress      string                                      `json:"ipv6_nexthop_address"`
	ManagementType          string                                      `json:"management_type"`
	Name                    string                                      `json:"name"`
	PeerAsn                 float64                                     `json:"peer_asn"`
	PeerIpAddress           string                                      `json:"peer_ip_address"`
	PeerIpv6NexthopAddress  string                                      `json:"peer_ipv6_nexthop_address"`
	Project                 string                                      `json:"project"`
	Region                  string                                      `json:"region"`
	Router                  string                                      `json:"router"`
	RouterApplianceInstance string                                      `json:"router_appliance_instance"`
	AdvertisedIpRanges      []computerouterpeer.AdvertisedIpRangesState `json:"advertised_ip_ranges"`
	Bfd                     []computerouterpeer.BfdState                `json:"bfd"`
	Timeouts                *computerouterpeer.TimeoutsState            `json:"timeouts"`
}
