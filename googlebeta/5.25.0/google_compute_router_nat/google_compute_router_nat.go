// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_router_nat

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

// Resource represents the Terraform resource google_compute_router_nat.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeRouterNatState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcrn *Resource) Type() string {
	return "google_compute_router_nat"
}

// LocalName returns the local name for [Resource].
func (gcrn *Resource) LocalName() string {
	return gcrn.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcrn *Resource) Configuration() interface{} {
	return gcrn.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcrn *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcrn)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcrn *Resource) Dependencies() terra.Dependencies {
	return gcrn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcrn *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcrn.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcrn *Resource) Attributes() googleComputeRouterNatAttributes {
	return googleComputeRouterNatAttributes{ref: terra.ReferenceResource(gcrn)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcrn *Resource) ImportState(state io.Reader) error {
	gcrn.state = &googleComputeRouterNatState{}
	if err := json.NewDecoder(state).Decode(gcrn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcrn.Type(), gcrn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcrn *Resource) State() (*googleComputeRouterNatState, bool) {
	return gcrn.state, gcrn.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcrn *Resource) StateMust() *googleComputeRouterNatState {
	if gcrn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcrn.Type(), gcrn.LocalName()))
	}
	return gcrn.state
}

// Args contains the configurations for google_compute_router_nat.
type Args struct {
	// DrainNatIps: set of string, optional
	DrainNatIps terra.SetValue[terra.StringValue] `hcl:"drain_nat_ips,attr"`
	// EnableDynamicPortAllocation: bool, optional
	EnableDynamicPortAllocation terra.BoolValue `hcl:"enable_dynamic_port_allocation,attr"`
	// EnableEndpointIndependentMapping: bool, optional
	EnableEndpointIndependentMapping terra.BoolValue `hcl:"enable_endpoint_independent_mapping,attr"`
	// EndpointTypes: list of string, optional
	EndpointTypes terra.ListValue[terra.StringValue] `hcl:"endpoint_types,attr"`
	// IcmpIdleTimeoutSec: number, optional
	IcmpIdleTimeoutSec terra.NumberValue `hcl:"icmp_idle_timeout_sec,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxPortsPerVm: number, optional
	MaxPortsPerVm terra.NumberValue `hcl:"max_ports_per_vm,attr"`
	// MinPortsPerVm: number, optional
	MinPortsPerVm terra.NumberValue `hcl:"min_ports_per_vm,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// NatIpAllocateOption: string, optional
	NatIpAllocateOption terra.StringValue `hcl:"nat_ip_allocate_option,attr"`
	// NatIps: set of string, optional
	NatIps terra.SetValue[terra.StringValue] `hcl:"nat_ips,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Router: string, required
	Router terra.StringValue `hcl:"router,attr" validate:"required"`
	// SourceSubnetworkIpRangesToNat: string, required
	SourceSubnetworkIpRangesToNat terra.StringValue `hcl:"source_subnetwork_ip_ranges_to_nat,attr" validate:"required"`
	// TcpEstablishedIdleTimeoutSec: number, optional
	TcpEstablishedIdleTimeoutSec terra.NumberValue `hcl:"tcp_established_idle_timeout_sec,attr"`
	// TcpTimeWaitTimeoutSec: number, optional
	TcpTimeWaitTimeoutSec terra.NumberValue `hcl:"tcp_time_wait_timeout_sec,attr"`
	// TcpTransitoryIdleTimeoutSec: number, optional
	TcpTransitoryIdleTimeoutSec terra.NumberValue `hcl:"tcp_transitory_idle_timeout_sec,attr"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// UdpIdleTimeoutSec: number, optional
	UdpIdleTimeoutSec terra.NumberValue `hcl:"udp_idle_timeout_sec,attr"`
	// LogConfig: optional
	LogConfig *LogConfig `hcl:"log_config,block"`
	// Rules: min=0
	Rules []Rules `hcl:"rules,block" validate:"min=0"`
	// Subnetwork: min=0
	Subnetwork []Subnetwork `hcl:"subnetwork,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeRouterNatAttributes struct {
	ref terra.Reference
}

// DrainNatIps returns a reference to field drain_nat_ips of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) DrainNatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrn.ref.Append("drain_nat_ips"))
}

// EnableDynamicPortAllocation returns a reference to field enable_dynamic_port_allocation of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) EnableDynamicPortAllocation() terra.BoolValue {
	return terra.ReferenceAsBool(gcrn.ref.Append("enable_dynamic_port_allocation"))
}

// EnableEndpointIndependentMapping returns a reference to field enable_endpoint_independent_mapping of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) EnableEndpointIndependentMapping() terra.BoolValue {
	return terra.ReferenceAsBool(gcrn.ref.Append("enable_endpoint_independent_mapping"))
}

// EndpointTypes returns a reference to field endpoint_types of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) EndpointTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcrn.ref.Append("endpoint_types"))
}

// IcmpIdleTimeoutSec returns a reference to field icmp_idle_timeout_sec of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) IcmpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("icmp_idle_timeout_sec"))
}

// Id returns a reference to field id of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("id"))
}

// MaxPortsPerVm returns a reference to field max_ports_per_vm of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) MaxPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("max_ports_per_vm"))
}

// MinPortsPerVm returns a reference to field min_ports_per_vm of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) MinPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("min_ports_per_vm"))
}

// Name returns a reference to field name of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("name"))
}

// NatIpAllocateOption returns a reference to field nat_ip_allocate_option of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) NatIpAllocateOption() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("nat_ip_allocate_option"))
}

// NatIps returns a reference to field nat_ips of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) NatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcrn.ref.Append("nat_ips"))
}

// Project returns a reference to field project of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("router"))
}

// SourceSubnetworkIpRangesToNat returns a reference to field source_subnetwork_ip_ranges_to_nat of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) SourceSubnetworkIpRangesToNat() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("source_subnetwork_ip_ranges_to_nat"))
}

// TcpEstablishedIdleTimeoutSec returns a reference to field tcp_established_idle_timeout_sec of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) TcpEstablishedIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("tcp_established_idle_timeout_sec"))
}

// TcpTimeWaitTimeoutSec returns a reference to field tcp_time_wait_timeout_sec of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) TcpTimeWaitTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("tcp_time_wait_timeout_sec"))
}

// TcpTransitoryIdleTimeoutSec returns a reference to field tcp_transitory_idle_timeout_sec of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) TcpTransitoryIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("tcp_transitory_idle_timeout_sec"))
}

// Type returns a reference to field type of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(gcrn.ref.Append("type"))
}

// UdpIdleTimeoutSec returns a reference to field udp_idle_timeout_sec of google_compute_router_nat.
func (gcrn googleComputeRouterNatAttributes) UdpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(gcrn.ref.Append("udp_idle_timeout_sec"))
}

func (gcrn googleComputeRouterNatAttributes) LogConfig() terra.ListValue[LogConfigAttributes] {
	return terra.ReferenceAsList[LogConfigAttributes](gcrn.ref.Append("log_config"))
}

func (gcrn googleComputeRouterNatAttributes) Rules() terra.SetValue[RulesAttributes] {
	return terra.ReferenceAsSet[RulesAttributes](gcrn.ref.Append("rules"))
}

func (gcrn googleComputeRouterNatAttributes) Subnetwork() terra.SetValue[SubnetworkAttributes] {
	return terra.ReferenceAsSet[SubnetworkAttributes](gcrn.ref.Append("subnetwork"))
}

func (gcrn googleComputeRouterNatAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcrn.ref.Append("timeouts"))
}

type googleComputeRouterNatState struct {
	DrainNatIps                      []string          `json:"drain_nat_ips"`
	EnableDynamicPortAllocation      bool              `json:"enable_dynamic_port_allocation"`
	EnableEndpointIndependentMapping bool              `json:"enable_endpoint_independent_mapping"`
	EndpointTypes                    []string          `json:"endpoint_types"`
	IcmpIdleTimeoutSec               float64           `json:"icmp_idle_timeout_sec"`
	Id                               string            `json:"id"`
	MaxPortsPerVm                    float64           `json:"max_ports_per_vm"`
	MinPortsPerVm                    float64           `json:"min_ports_per_vm"`
	Name                             string            `json:"name"`
	NatIpAllocateOption              string            `json:"nat_ip_allocate_option"`
	NatIps                           []string          `json:"nat_ips"`
	Project                          string            `json:"project"`
	Region                           string            `json:"region"`
	Router                           string            `json:"router"`
	SourceSubnetworkIpRangesToNat    string            `json:"source_subnetwork_ip_ranges_to_nat"`
	TcpEstablishedIdleTimeoutSec     float64           `json:"tcp_established_idle_timeout_sec"`
	TcpTimeWaitTimeoutSec            float64           `json:"tcp_time_wait_timeout_sec"`
	TcpTransitoryIdleTimeoutSec      float64           `json:"tcp_transitory_idle_timeout_sec"`
	Type                             string            `json:"type"`
	UdpIdleTimeoutSec                float64           `json:"udp_idle_timeout_sec"`
	LogConfig                        []LogConfigState  `json:"log_config"`
	Rules                            []RulesState      `json:"rules"`
	Subnetwork                       []SubnetworkState `json:"subnetwork"`
	Timeouts                         *TimeoutsState    `json:"timeouts"`
}
