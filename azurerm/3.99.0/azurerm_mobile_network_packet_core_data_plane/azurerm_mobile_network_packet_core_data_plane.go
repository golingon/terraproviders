// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_mobile_network_packet_core_data_plane

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

// Resource represents the Terraform resource azurerm_mobile_network_packet_core_data_plane.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermMobileNetworkPacketCoreDataPlaneState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (amnpcdp *Resource) Type() string {
	return "azurerm_mobile_network_packet_core_data_plane"
}

// LocalName returns the local name for [Resource].
func (amnpcdp *Resource) LocalName() string {
	return amnpcdp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (amnpcdp *Resource) Configuration() interface{} {
	return amnpcdp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (amnpcdp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(amnpcdp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (amnpcdp *Resource) Dependencies() terra.Dependencies {
	return amnpcdp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (amnpcdp *Resource) LifecycleManagement() *terra.Lifecycle {
	return amnpcdp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (amnpcdp *Resource) Attributes() azurermMobileNetworkPacketCoreDataPlaneAttributes {
	return azurermMobileNetworkPacketCoreDataPlaneAttributes{ref: terra.ReferenceResource(amnpcdp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (amnpcdp *Resource) ImportState(state io.Reader) error {
	amnpcdp.state = &azurermMobileNetworkPacketCoreDataPlaneState{}
	if err := json.NewDecoder(state).Decode(amnpcdp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", amnpcdp.Type(), amnpcdp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (amnpcdp *Resource) State() (*azurermMobileNetworkPacketCoreDataPlaneState, bool) {
	return amnpcdp.state, amnpcdp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (amnpcdp *Resource) StateMust() *azurermMobileNetworkPacketCoreDataPlaneState {
	if amnpcdp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", amnpcdp.Type(), amnpcdp.LocalName()))
	}
	return amnpcdp.state
}

// Args contains the configurations for azurerm_mobile_network_packet_core_data_plane.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// MobileNetworkPacketCoreControlPlaneId: string, required
	MobileNetworkPacketCoreControlPlaneId terra.StringValue `hcl:"mobile_network_packet_core_control_plane_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// UserPlaneAccessIpv4Address: string, optional
	UserPlaneAccessIpv4Address terra.StringValue `hcl:"user_plane_access_ipv4_address,attr"`
	// UserPlaneAccessIpv4Gateway: string, optional
	UserPlaneAccessIpv4Gateway terra.StringValue `hcl:"user_plane_access_ipv4_gateway,attr"`
	// UserPlaneAccessIpv4Subnet: string, optional
	UserPlaneAccessIpv4Subnet terra.StringValue `hcl:"user_plane_access_ipv4_subnet,attr"`
	// UserPlaneAccessName: string, optional
	UserPlaneAccessName terra.StringValue `hcl:"user_plane_access_name,attr"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermMobileNetworkPacketCoreDataPlaneAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("location"))
}

// MobileNetworkPacketCoreControlPlaneId returns a reference to field mobile_network_packet_core_control_plane_id of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) MobileNetworkPacketCoreControlPlaneId() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("mobile_network_packet_core_control_plane_id"))
}

// Name returns a reference to field name of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("name"))
}

// Tags returns a reference to field tags of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amnpcdp.ref.Append("tags"))
}

// UserPlaneAccessIpv4Address returns a reference to field user_plane_access_ipv4_address of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) UserPlaneAccessIpv4Address() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("user_plane_access_ipv4_address"))
}

// UserPlaneAccessIpv4Gateway returns a reference to field user_plane_access_ipv4_gateway of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) UserPlaneAccessIpv4Gateway() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("user_plane_access_ipv4_gateway"))
}

// UserPlaneAccessIpv4Subnet returns a reference to field user_plane_access_ipv4_subnet of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) UserPlaneAccessIpv4Subnet() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("user_plane_access_ipv4_subnet"))
}

// UserPlaneAccessName returns a reference to field user_plane_access_name of azurerm_mobile_network_packet_core_data_plane.
func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) UserPlaneAccessName() terra.StringValue {
	return terra.ReferenceAsString(amnpcdp.ref.Append("user_plane_access_name"))
}

func (amnpcdp azurermMobileNetworkPacketCoreDataPlaneAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](amnpcdp.ref.Append("timeouts"))
}

type azurermMobileNetworkPacketCoreDataPlaneState struct {
	Id                                    string            `json:"id"`
	Location                              string            `json:"location"`
	MobileNetworkPacketCoreControlPlaneId string            `json:"mobile_network_packet_core_control_plane_id"`
	Name                                  string            `json:"name"`
	Tags                                  map[string]string `json:"tags"`
	UserPlaneAccessIpv4Address            string            `json:"user_plane_access_ipv4_address"`
	UserPlaneAccessIpv4Gateway            string            `json:"user_plane_access_ipv4_gateway"`
	UserPlaneAccessIpv4Subnet             string            `json:"user_plane_access_ipv4_subnet"`
	UserPlaneAccessName                   string            `json:"user_plane_access_name"`
	Timeouts                              *TimeoutsState    `json:"timeouts"`
}
