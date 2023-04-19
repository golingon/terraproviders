// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputerouternat "github.com/golingon/terraproviders/google/4.62.0/datacomputerouternat"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeRouterNat creates a new instance of [DataComputeRouterNat].
func NewDataComputeRouterNat(name string, args DataComputeRouterNatArgs) *DataComputeRouterNat {
	return &DataComputeRouterNat{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeRouterNat)(nil)

// DataComputeRouterNat represents the Terraform data resource google_compute_router_nat.
type DataComputeRouterNat struct {
	Name string
	Args DataComputeRouterNatArgs
}

// DataSource returns the Terraform object type for [DataComputeRouterNat].
func (crn *DataComputeRouterNat) DataSource() string {
	return "google_compute_router_nat"
}

// LocalName returns the local name for [DataComputeRouterNat].
func (crn *DataComputeRouterNat) LocalName() string {
	return crn.Name
}

// Configuration returns the configuration (args) for [DataComputeRouterNat].
func (crn *DataComputeRouterNat) Configuration() interface{} {
	return crn.Args
}

// Attributes returns the attributes for [DataComputeRouterNat].
func (crn *DataComputeRouterNat) Attributes() dataComputeRouterNatAttributes {
	return dataComputeRouterNatAttributes{ref: terra.ReferenceDataResource(crn)}
}

// DataComputeRouterNatArgs contains the configurations for google_compute_router_nat.
type DataComputeRouterNatArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Router: string, required
	Router terra.StringValue `hcl:"router,attr" validate:"required"`
	// LogConfig: min=0
	LogConfig []datacomputerouternat.LogConfig `hcl:"log_config,block" validate:"min=0"`
	// Rules: min=0
	Rules []datacomputerouternat.Rules `hcl:"rules,block" validate:"min=0"`
	// Subnetwork: min=0
	Subnetwork []datacomputerouternat.Subnetwork `hcl:"subnetwork,block" validate:"min=0"`
}
type dataComputeRouterNatAttributes struct {
	ref terra.Reference
}

// DrainNatIps returns a reference to field drain_nat_ips of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) DrainNatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crn.ref.Append("drain_nat_ips"))
}

// EnableDynamicPortAllocation returns a reference to field enable_dynamic_port_allocation of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) EnableDynamicPortAllocation() terra.BoolValue {
	return terra.ReferenceAsBool(crn.ref.Append("enable_dynamic_port_allocation"))
}

// EnableEndpointIndependentMapping returns a reference to field enable_endpoint_independent_mapping of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) EnableEndpointIndependentMapping() terra.BoolValue {
	return terra.ReferenceAsBool(crn.ref.Append("enable_endpoint_independent_mapping"))
}

// IcmpIdleTimeoutSec returns a reference to field icmp_idle_timeout_sec of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) IcmpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("icmp_idle_timeout_sec"))
}

// Id returns a reference to field id of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("id"))
}

// MaxPortsPerVm returns a reference to field max_ports_per_vm of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) MaxPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("max_ports_per_vm"))
}

// MinPortsPerVm returns a reference to field min_ports_per_vm of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) MinPortsPerVm() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("min_ports_per_vm"))
}

// Name returns a reference to field name of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("name"))
}

// NatIpAllocateOption returns a reference to field nat_ip_allocate_option of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) NatIpAllocateOption() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("nat_ip_allocate_option"))
}

// NatIps returns a reference to field nat_ips of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) NatIps() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](crn.ref.Append("nat_ips"))
}

// Project returns a reference to field project of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("project"))
}

// Region returns a reference to field region of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("region"))
}

// Router returns a reference to field router of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) Router() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("router"))
}

// SourceSubnetworkIpRangesToNat returns a reference to field source_subnetwork_ip_ranges_to_nat of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) SourceSubnetworkIpRangesToNat() terra.StringValue {
	return terra.ReferenceAsString(crn.ref.Append("source_subnetwork_ip_ranges_to_nat"))
}

// TcpEstablishedIdleTimeoutSec returns a reference to field tcp_established_idle_timeout_sec of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) TcpEstablishedIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_established_idle_timeout_sec"))
}

// TcpTimeWaitTimeoutSec returns a reference to field tcp_time_wait_timeout_sec of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) TcpTimeWaitTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_time_wait_timeout_sec"))
}

// TcpTransitoryIdleTimeoutSec returns a reference to field tcp_transitory_idle_timeout_sec of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) TcpTransitoryIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("tcp_transitory_idle_timeout_sec"))
}

// UdpIdleTimeoutSec returns a reference to field udp_idle_timeout_sec of google_compute_router_nat.
func (crn dataComputeRouterNatAttributes) UdpIdleTimeoutSec() terra.NumberValue {
	return terra.ReferenceAsNumber(crn.ref.Append("udp_idle_timeout_sec"))
}

func (crn dataComputeRouterNatAttributes) LogConfig() terra.ListValue[datacomputerouternat.LogConfigAttributes] {
	return terra.ReferenceAsList[datacomputerouternat.LogConfigAttributes](crn.ref.Append("log_config"))
}

func (crn dataComputeRouterNatAttributes) Rules() terra.SetValue[datacomputerouternat.RulesAttributes] {
	return terra.ReferenceAsSet[datacomputerouternat.RulesAttributes](crn.ref.Append("rules"))
}

func (crn dataComputeRouterNatAttributes) Subnetwork() terra.SetValue[datacomputerouternat.SubnetworkAttributes] {
	return terra.ReferenceAsSet[datacomputerouternat.SubnetworkAttributes](crn.ref.Append("subnetwork"))
}
