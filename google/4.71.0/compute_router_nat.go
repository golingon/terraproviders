// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computerouternat "github.com/golingon/terraproviders/google/4.71.0/computerouternat"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeRouterNat creates a new instance of [ComputeRouterNat].
func NewComputeRouterNat(name string, args ComputeRouterNatArgs) *ComputeRouterNat {
	return &ComputeRouterNat{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeRouterNat)(nil)

// ComputeRouterNat represents the Terraform resource google_compute_router_nat.
type ComputeRouterNat struct {
	Name      string
	Args      ComputeRouterNatArgs
	state     *computeRouterNatState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeRouterNat].
func (crn *ComputeRouterNat) Type() string {
	return "google_compute_router_nat"
}

// LocalName returns the local name for [ComputeRouterNat].
func (crn *ComputeRouterNat) LocalName() string {
	return crn.Name
}

// Configuration returns the configuration (args) for [ComputeRouterNat].
func (crn *ComputeRouterNat) Configuration() interface{} {
	return crn.Args
}

// DependOn is used for other resources to depend on [ComputeRouterNat].
func (crn *ComputeRouterNat) DependOn() terra.Reference {
	return terra.ReferenceResource(crn)
}

// Dependencies returns the list of resources [ComputeRouterNat] depends_on.
func (crn *ComputeRouterNat) Dependencies() terra.Dependencies {
	return crn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeRouterNat].
func (crn *ComputeRouterNat) LifecycleManagement() *terra.Lifecycle {
	return crn.Lifecycle
}

// Attributes returns the attributes for [ComputeRouterNat].
func (crn *ComputeRouterNat) Attributes() computeRouterNatAttributes {
	return computeRouterNatAttributes{ref: terra.ReferenceResource(crn)}
}

// ImportState imports the given attribute values into [ComputeRouterNat]'s state.
func (crn *ComputeRouterNat) ImportState(av io.Reader) error {
	crn.state = &computeRouterNatState{}
	if err := json.NewDecoder(av).Decode(crn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crn.Type(), crn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeRouterNat] has state.
func (crn *ComputeRouterNat) State() (*computeRouterNatState, bool) {
	return crn.state, crn.state != nil
}

// StateMust returns the state for [ComputeRouterNat]. Panics if the state is nil.
func (crn *ComputeRouterNat) StateMust() *computeRouterNatState {
	if crn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crn.Type(), crn.LocalName()))
	}
	return crn.state
}

// ComputeRouterNatArgs contains the configurations for google_compute_router_nat.
type ComputeRouterNatArgs struct {
	// DrainNatIps: set of string, optional
	DrainNatIps terra.SetValue[terra.StringValue] `hcl:"drain_nat_ips,attr"`
	// EnableDynamicPortAllocation: bool, optional
	EnableDynamicPortAllocation terra.BoolValue `hcl:"enable_dynamic_port_allocation,attr"`
	// EnableEndpointIndependentMapping: bool, optional
	EnableEndpointIndependentMapping terra.BoolValue `hcl:"enable_endpoint_independent_mapping,attr"`
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
	// NatIpAllocateOption: string, required
	NatIpAllocateOption terra.StringValue `hcl:"nat_ip_allocate_option,attr" validate:"required"`
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
	// UdpIdleTimeoutSec: number, optional
	UdpIdleTimeoutSec terra.NumberValue `hcl:"udp_idle_timeout_sec,attr"`
	// LogConfig: optional
	LogConfig *computerouternat.LogConfig `hcl:"log_config,block"`
	// Rules: min=0
	Rules []computerouternat.Rules `hcl:"rules,block" validate:"min=0"`
	// Subnetwork: min=0
	Subnetwork []computerouternat.Subnetwork `hcl:"subnetwork,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computerouternat.Timeouts `hcl:"timeouts,block"`
}
type computeRouterNatAttributes struct {
	ref terra.Reference
}

// DrainNatIps returns a reference to field drain_nat_ips of google_compute_router_nat.
func (crn computeRouterNatAttributes) DrainNatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crn.ref.Append("drain_nat_ips"))
}

// EnableDynamicPortAllocation returns a reference to field enable_dynamic_port_allocation of google_compute_router_nat.
func (crn computeRouterNatAttributes) EnableDynamicPortAllocation() terra.BoolValue {
	return terra.ReferenceAsBool(crn.ref.Append("enable_dynamic_port_allocation"))
}

// EnableEndpointIndependentMapping returns a reference to field enable_endpoint_independent_mapping of google_compute_router_nat.
func (crn computeRouterNatAttributes) EnableEndpointIndependentMapping() terra.BoolValue {
	return terra.ReferenceAsBool(crn.ref.Append("enable_endpoint_independent_mapping"))
}

// IcmpIdleTimeoutSec returns a reference to field icmp_idle_timeout_sec of google_compute_router_nat.
func (crn computeRouterNatAttributes) IcmpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("icmp_idle_timeout_sec"))
}

// Id returns a reference to field id of google_compute_router_nat.
func (crn computeRouterNatAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("id"))
}

// MaxPortsPerVm returns a reference to field max_ports_per_vm of google_compute_router_nat.
func (crn computeRouterNatAttributes) MaxPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("max_ports_per_vm"))
}

// MinPortsPerVm returns a reference to field min_ports_per_vm of google_compute_router_nat.
func (crn computeRouterNatAttributes) MinPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("min_ports_per_vm"))
}

// Name returns a reference to field name of google_compute_router_nat.
func (crn computeRouterNatAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("name"))
}

// NatIpAllocateOption returns a reference to field nat_ip_allocate_option of google_compute_router_nat.
func (crn computeRouterNatAttributes) NatIpAllocateOption() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("nat_ip_allocate_option"))
}

// NatIps returns a reference to field nat_ips of google_compute_router_nat.
func (crn computeRouterNatAttributes) NatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crn.ref.Append("nat_ips"))
}

// Project returns a reference to field project of google_compute_router_nat.
func (crn computeRouterNatAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router_nat.
func (crn computeRouterNatAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_router_nat.
func (crn computeRouterNatAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("router"))
}

// SourceSubnetworkIpRangesToNat returns a reference to field source_subnetwork_ip_ranges_to_nat of google_compute_router_nat.
func (crn computeRouterNatAttributes) SourceSubnetworkIpRangesToNat() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("source_subnetwork_ip_ranges_to_nat"))
}

// TcpEstablishedIdleTimeoutSec returns a reference to field tcp_established_idle_timeout_sec of google_compute_router_nat.
func (crn computeRouterNatAttributes) TcpEstablishedIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_established_idle_timeout_sec"))
}

// TcpTimeWaitTimeoutSec returns a reference to field tcp_time_wait_timeout_sec of google_compute_router_nat.
func (crn computeRouterNatAttributes) TcpTimeWaitTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_time_wait_timeout_sec"))
}

// TcpTransitoryIdleTimeoutSec returns a reference to field tcp_transitory_idle_timeout_sec of google_compute_router_nat.
func (crn computeRouterNatAttributes) TcpTransitoryIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_transitory_idle_timeout_sec"))
}

// UdpIdleTimeoutSec returns a reference to field udp_idle_timeout_sec of google_compute_router_nat.
func (crn computeRouterNatAttributes) UdpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("udp_idle_timeout_sec"))
}

func (crn computeRouterNatAttributes) LogConfig() terra.ListValue[computerouternat.LogConfigAttributes] {
	return terra.ReferenceAsList[computerouternat.LogConfigAttributes](crn.ref.Append("log_config"))
}

func (crn computeRouterNatAttributes) Rules() terra.SetValue[computerouternat.RulesAttributes] {
	return terra.ReferenceAsSet[computerouternat.RulesAttributes](crn.ref.Append("rules"))
}

func (crn computeRouterNatAttributes) Subnetwork() terra.SetValue[computerouternat.SubnetworkAttributes] {
	return terra.ReferenceAsSet[computerouternat.SubnetworkAttributes](crn.ref.Append("subnetwork"))
}

func (crn computeRouterNatAttributes) Timeouts() computerouternat.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computerouternat.TimeoutsAttributes](crn.ref.Append("timeouts"))
}

type computeRouterNatState struct {
	DrainNatIps                      []string                           `json:"drain_nat_ips"`
	EnableDynamicPortAllocation      bool                               `json:"enable_dynamic_port_allocation"`
	EnableEndpointIndependentMapping bool                               `json:"enable_endpoint_independent_mapping"`
	IcmpIdleTimeoutSec               float64                            `json:"icmp_idle_timeout_sec"`
	Id                               string                             `json:"id"`
	MaxPortsPerVm                    float64                            `json:"max_ports_per_vm"`
	MinPortsPerVm                    float64                            `json:"min_ports_per_vm"`
	Name                             string                             `json:"name"`
	NatIpAllocateOption              string                             `json:"nat_ip_allocate_option"`
	NatIps                           []string                           `json:"nat_ips"`
	Project                          string                             `json:"project"`
	Region                           string                             `json:"region"`
	Router                           string                             `json:"router"`
	SourceSubnetworkIpRangesToNat    string                             `json:"source_subnetwork_ip_ranges_to_nat"`
	TcpEstablishedIdleTimeoutSec     float64                            `json:"tcp_established_idle_timeout_sec"`
	TcpTimeWaitTimeoutSec            float64                            `json:"tcp_time_wait_timeout_sec"`
	TcpTransitoryIdleTimeoutSec      float64                            `json:"tcp_transitory_idle_timeout_sec"`
	UdpIdleTimeoutSec                float64                            `json:"udp_idle_timeout_sec"`
	LogConfig                        []computerouternat.LogConfigState  `json:"log_config"`
	Rules                            []computerouternat.RulesState      `json:"rules"`
	Subnetwork                       []computerouternat.SubnetworkState `json:"subnetwork"`
	Timeouts                         *computerouternat.TimeoutsState    `json:"timeouts"`
}
