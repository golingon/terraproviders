// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computerouterinterface "github.com/golingon/terraproviders/google/4.63.1/computerouterinterface"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRouterInterface creates a new instance of [ComputeRouterInterface].
func NewComputeRouterInterface(name string, args ComputeRouterInterfaceArgs) *ComputeRouterInterface {
	return &ComputeRouterInterface{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRouterInterface)(nil)

// ComputeRouterInterface represents the Terraform resource google_compute_router_interface.
type ComputeRouterInterface struct {
	Name      string
	Args      ComputeRouterInterfaceArgs
	state     *computeRouterInterfaceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRouterInterface].
func (cri *ComputeRouterInterface) Type() string {
	return "google_compute_router_interface"
}

// LocalName returns the local name for [ComputeRouterInterface].
func (cri *ComputeRouterInterface) LocalName() string {
	return cri.Name
}

// Configuration returns the configuration (args) for [ComputeRouterInterface].
func (cri *ComputeRouterInterface) Configuration() interface{} {
	return cri.Args
}

// DependOn is used for other resources to depend on [ComputeRouterInterface].
func (cri *ComputeRouterInterface) DependOn() terra.Reference {
	return terra.ReferenceResource(cri)
}

// Dependencies returns the list of resources [ComputeRouterInterface] depends_on.
func (cri *ComputeRouterInterface) Dependencies() terra.Dependencies {
	return cri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRouterInterface].
func (cri *ComputeRouterInterface) LifecycleManagement() *terra.Lifecycle {
	return cri.Lifecycle
}

// Attributes returns the attributes for [ComputeRouterInterface].
func (cri *ComputeRouterInterface) Attributes() computeRouterInterfaceAttributes {
	return computeRouterInterfaceAttributes{ref: terra.ReferenceResource(cri)}
}

// ImportState imports the given attribute values into [ComputeRouterInterface]'s state.
func (cri *ComputeRouterInterface) ImportState(av io.Reader) error {
	cri.state = &computeRouterInterfaceState{}
	if err := json.NewDecoder(av).Decode(cri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cri.Type(), cri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRouterInterface] has state.
func (cri *ComputeRouterInterface) State() (*computeRouterInterfaceState, bool) {
	return cri.state, cri.state != nil
}

// StateMust returns the state for [ComputeRouterInterface]. Panics if the state is nil.
func (cri *ComputeRouterInterface) StateMust() *computeRouterInterfaceState {
	if cri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cri.Type(), cri.LocalName()))
	}
	return cri.state
}

// ComputeRouterInterfaceArgs contains the configurations for google_compute_router_interface.
type ComputeRouterInterfaceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InterconnectAttachment: string, optional
	InterconnectAttachment terra.StringValue `hcl:"interconnect_attachment,attr"`
	// IpRange: string, optional
	IpRange terra.StringValue `hcl:"ip_range,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateIpAddress: string, optional
	PrivateIpAddress terra.StringValue `hcl:"private_ip_address,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RedundantInterface: string, optional
	RedundantInterface terra.StringValue `hcl:"redundant_interface,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Router: string, required
	Router terra.StringValue `hcl:"router,attr" validate:"required"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// VpnTunnel: string, optional
	VpnTunnel terra.StringValue `hcl:"vpn_tunnel,attr"`
	// Timeouts: optional
	Timeouts *computerouterinterface.Timeouts `hcl:"timeouts,block"`
}
type computeRouterInterfaceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("id"))
}

// InterconnectAttachment returns a reference to field interconnect_attachment of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) InterconnectAttachment() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("interconnect_attachment"))
}

// IpRange returns a reference to field ip_range of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) IpRange() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("ip_range"))
}

// Name returns a reference to field name of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("name"))
}

// PrivateIpAddress returns a reference to field private_ip_address of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) PrivateIpAddress() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("private_ip_address"))
}

// Project returns a reference to field project of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("project"))
}

// RedundantInterface returns a reference to field redundant_interface of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) RedundantInterface() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("redundant_interface"))
}

// Region returns a reference to field region of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("router"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("subnetwork"))
}

// VpnTunnel returns a reference to field vpn_tunnel of google_compute_router_interface.
func (cri computeRouterInterfaceAttributes) VpnTunnel() terra.StringValue {
	return terra.ReferenceAsString(cri.ref.Append("vpn_tunnel"))
}

func (cri computeRouterInterfaceAttributes) Timeouts() computerouterinterface.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computerouterinterface.TimeoutsAttributes](cri.ref.Append("timeouts"))
}

type computeRouterInterfaceState struct {
	Id                     string                                `json:"id"`
	InterconnectAttachment string                                `json:"interconnect_attachment"`
	IpRange                string                                `json:"ip_range"`
	Name                   string                                `json:"name"`
	PrivateIpAddress       string                                `json:"private_ip_address"`
	Project                string                                `json:"project"`
	RedundantInterface     string                                `json:"redundant_interface"`
	Region                 string                                `json:"region"`
	Router                 string                                `json:"router"`
	Subnetwork             string                                `json:"subnetwork"`
	VpnTunnel              string                                `json:"vpn_tunnel"`
	Timeouts               *computerouterinterface.TimeoutsState `json:"timeouts"`
}
