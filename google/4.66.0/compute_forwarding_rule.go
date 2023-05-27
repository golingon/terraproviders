// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeforwardingrule "github.com/golingon/terraproviders/google/4.66.0/computeforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeForwardingRule creates a new instance of [ComputeForwardingRule].
func NewComputeForwardingRule(name string, args ComputeForwardingRuleArgs) *ComputeForwardingRule {
	return &ComputeForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeForwardingRule)(nil)

// ComputeForwardingRule represents the Terraform resource google_compute_forwarding_rule.
type ComputeForwardingRule struct {
	Name      string
	Args      ComputeForwardingRuleArgs
	state     *computeForwardingRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) Type() string {
	return "google_compute_forwarding_rule"
}

// LocalName returns the local name for [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) LocalName() string {
	return cfr.Name
}

// Configuration returns the configuration (args) for [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) Configuration() interface{} {
	return cfr.Args
}

// DependOn is used for other resources to depend on [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cfr)
}

// Dependencies returns the list of resources [ComputeForwardingRule] depends_on.
func (cfr *ComputeForwardingRule) Dependencies() terra.Dependencies {
	return cfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) LifecycleManagement() *terra.Lifecycle {
	return cfr.Lifecycle
}

// Attributes returns the attributes for [ComputeForwardingRule].
func (cfr *ComputeForwardingRule) Attributes() computeForwardingRuleAttributes {
	return computeForwardingRuleAttributes{ref: terra.ReferenceResource(cfr)}
}

// ImportState imports the given attribute values into [ComputeForwardingRule]'s state.
func (cfr *ComputeForwardingRule) ImportState(av io.Reader) error {
	cfr.state = &computeForwardingRuleState{}
	if err := json.NewDecoder(av).Decode(cfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cfr.Type(), cfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeForwardingRule] has state.
func (cfr *ComputeForwardingRule) State() (*computeForwardingRuleState, bool) {
	return cfr.state, cfr.state != nil
}

// StateMust returns the state for [ComputeForwardingRule]. Panics if the state is nil.
func (cfr *ComputeForwardingRule) StateMust() *computeForwardingRuleState {
	if cfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cfr.Type(), cfr.LocalName()))
	}
	return cfr.state
}

// ComputeForwardingRuleArgs contains the configurations for google_compute_forwarding_rule.
type ComputeForwardingRuleArgs struct {
	// AllPorts: bool, optional
	AllPorts terra.BoolValue `hcl:"all_ports,attr"`
	// AllowGlobalAccess: bool, optional
	AllowGlobalAccess terra.BoolValue `hcl:"allow_global_access,attr"`
	// BackendService: string, optional
	BackendService terra.StringValue `hcl:"backend_service,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IpAddress: string, optional
	IpAddress terra.StringValue `hcl:"ip_address,attr"`
	// IpProtocol: string, optional
	IpProtocol terra.StringValue `hcl:"ip_protocol,attr"`
	// IsMirroringCollector: bool, optional
	IsMirroringCollector terra.BoolValue `hcl:"is_mirroring_collector,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LoadBalancingScheme: string, optional
	LoadBalancingScheme terra.StringValue `hcl:"load_balancing_scheme,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NetworkTier: string, optional
	NetworkTier terra.StringValue `hcl:"network_tier,attr"`
	// PortRange: string, optional
	PortRange terra.StringValue `hcl:"port_range,attr"`
	// Ports: set of string, optional
	Ports terra.SetValue[terra.StringValue] `hcl:"ports,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ServiceLabel: string, optional
	ServiceLabel terra.StringValue `hcl:"service_label,attr"`
	// SourceIpRanges: list of string, optional
	SourceIpRanges terra.ListValue[terra.StringValue] `hcl:"source_ip_ranges,attr"`
	// Subnetwork: string, optional
	Subnetwork terra.StringValue `hcl:"subnetwork,attr"`
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// ServiceDirectoryRegistrations: optional
	ServiceDirectoryRegistrations *computeforwardingrule.ServiceDirectoryRegistrations `hcl:"service_directory_registrations,block"`
	// Timeouts: optional
	Timeouts *computeforwardingrule.Timeouts `hcl:"timeouts,block"`
}
type computeForwardingRuleAttributes struct {
	ref terra.Reference
}

// AllPorts returns a reference to field all_ports of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) AllPorts() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("all_ports"))
}

// AllowGlobalAccess returns a reference to field allow_global_access of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) AllowGlobalAccess() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("allow_global_access"))
}

// BackendService returns a reference to field backend_service of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("backend_service"))
}

// BaseForwardingRule returns a reference to field base_forwarding_rule of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) BaseForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("base_forwarding_rule"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("ip_address"))
}

// IpProtocol returns a reference to field ip_protocol of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("ip_protocol"))
}

// IsMirroringCollector returns a reference to field is_mirroring_collector of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) IsMirroringCollector() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("is_mirroring_collector"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cfr.ref.Append("labels"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("load_balancing_scheme"))
}

// Name returns a reference to field name of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("network"))
}

// NetworkTier returns a reference to field network_tier of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("network_tier"))
}

// PortRange returns a reference to field port_range of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) PortRange() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("port_range"))
}

// Ports returns a reference to field ports of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Ports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfr.ref.Append("ports"))
}

// Project returns a reference to field project of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("project"))
}

// PscConnectionId returns a reference to field psc_connection_id of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("psc_connection_id"))
}

// PscConnectionStatus returns a reference to field psc_connection_status of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) PscConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("psc_connection_status"))
}

// Region returns a reference to field region of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("self_link"))
}

// ServiceLabel returns a reference to field service_label of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) ServiceLabel() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("service_label"))
}

// ServiceName returns a reference to field service_name of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("service_name"))
}

// SourceIpRanges returns a reference to field source_ip_ranges of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) SourceIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfr.ref.Append("source_ip_ranges"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("subnetwork"))
}

// Target returns a reference to field target of google_compute_forwarding_rule.
func (cfr computeForwardingRuleAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("target"))
}

func (cfr computeForwardingRuleAttributes) ServiceDirectoryRegistrations() terra.ListValue[computeforwardingrule.ServiceDirectoryRegistrationsAttributes] {
	return terra.ReferenceAsList[computeforwardingrule.ServiceDirectoryRegistrationsAttributes](cfr.ref.Append("service_directory_registrations"))
}

func (cfr computeForwardingRuleAttributes) Timeouts() computeforwardingrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeforwardingrule.TimeoutsAttributes](cfr.ref.Append("timeouts"))
}

type computeForwardingRuleState struct {
	AllPorts                      bool                                                       `json:"all_ports"`
	AllowGlobalAccess             bool                                                       `json:"allow_global_access"`
	BackendService                string                                                     `json:"backend_service"`
	BaseForwardingRule            string                                                     `json:"base_forwarding_rule"`
	CreationTimestamp             string                                                     `json:"creation_timestamp"`
	Description                   string                                                     `json:"description"`
	Id                            string                                                     `json:"id"`
	IpAddress                     string                                                     `json:"ip_address"`
	IpProtocol                    string                                                     `json:"ip_protocol"`
	IsMirroringCollector          bool                                                       `json:"is_mirroring_collector"`
	LabelFingerprint              string                                                     `json:"label_fingerprint"`
	Labels                        map[string]string                                          `json:"labels"`
	LoadBalancingScheme           string                                                     `json:"load_balancing_scheme"`
	Name                          string                                                     `json:"name"`
	Network                       string                                                     `json:"network"`
	NetworkTier                   string                                                     `json:"network_tier"`
	PortRange                     string                                                     `json:"port_range"`
	Ports                         []string                                                   `json:"ports"`
	Project                       string                                                     `json:"project"`
	PscConnectionId               string                                                     `json:"psc_connection_id"`
	PscConnectionStatus           string                                                     `json:"psc_connection_status"`
	Region                        string                                                     `json:"region"`
	SelfLink                      string                                                     `json:"self_link"`
	ServiceLabel                  string                                                     `json:"service_label"`
	ServiceName                   string                                                     `json:"service_name"`
	SourceIpRanges                []string                                                   `json:"source_ip_ranges"`
	Subnetwork                    string                                                     `json:"subnetwork"`
	Target                        string                                                     `json:"target"`
	ServiceDirectoryRegistrations []computeforwardingrule.ServiceDirectoryRegistrationsState `json:"service_directory_registrations"`
	Timeouts                      *computeforwardingrule.TimeoutsState                       `json:"timeouts"`
}
