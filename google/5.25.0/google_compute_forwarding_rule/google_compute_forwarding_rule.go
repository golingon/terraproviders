// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_compute_forwarding_rule

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

// Resource represents the Terraform resource google_compute_forwarding_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *googleComputeForwardingRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcfr *Resource) Type() string {
	return "google_compute_forwarding_rule"
}

// LocalName returns the local name for [Resource].
func (gcfr *Resource) LocalName() string {
	return gcfr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcfr *Resource) Configuration() interface{} {
	return gcfr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcfr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcfr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcfr *Resource) Dependencies() terra.Dependencies {
	return gcfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcfr *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcfr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcfr *Resource) Attributes() googleComputeForwardingRuleAttributes {
	return googleComputeForwardingRuleAttributes{ref: terra.ReferenceResource(gcfr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcfr *Resource) ImportState(state io.Reader) error {
	gcfr.state = &googleComputeForwardingRuleState{}
	if err := json.NewDecoder(state).Decode(gcfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcfr.Type(), gcfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcfr *Resource) State() (*googleComputeForwardingRuleState, bool) {
	return gcfr.state, gcfr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcfr *Resource) StateMust() *googleComputeForwardingRuleState {
	if gcfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcfr.Type(), gcfr.LocalName()))
	}
	return gcfr.state
}

// Args contains the configurations for google_compute_forwarding_rule.
type Args struct {
	// AllPorts: bool, optional
	AllPorts terra.BoolValue `hcl:"all_ports,attr"`
	// AllowGlobalAccess: bool, optional
	AllowGlobalAccess terra.BoolValue `hcl:"allow_global_access,attr"`
	// AllowPscGlobalAccess: bool, optional
	AllowPscGlobalAccess terra.BoolValue `hcl:"allow_psc_global_access,attr"`
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
	// IpVersion: string, optional
	IpVersion terra.StringValue `hcl:"ip_version,attr"`
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
	// NoAutomateDnsZone: bool, optional
	NoAutomateDnsZone terra.BoolValue `hcl:"no_automate_dns_zone,attr"`
	// PortRange: string, optional
	PortRange terra.StringValue `hcl:"port_range,attr"`
	// Ports: set of string, optional
	Ports terra.SetValue[terra.StringValue] `hcl:"ports,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// RecreateClosedPsc: bool, optional
	RecreateClosedPsc terra.BoolValue `hcl:"recreate_closed_psc,attr"`
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
	ServiceDirectoryRegistrations *ServiceDirectoryRegistrations `hcl:"service_directory_registrations,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleComputeForwardingRuleAttributes struct {
	ref terra.Reference
}

// AllPorts returns a reference to field all_ports of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) AllPorts() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("all_ports"))
}

// AllowGlobalAccess returns a reference to field allow_global_access of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) AllowGlobalAccess() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("allow_global_access"))
}

// AllowPscGlobalAccess returns a reference to field allow_psc_global_access of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) AllowPscGlobalAccess() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("allow_psc_global_access"))
}

// BackendService returns a reference to field backend_service of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("backend_service"))
}

// BaseForwardingRule returns a reference to field base_forwarding_rule of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) BaseForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("base_forwarding_rule"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("description"))
}

// EffectiveLabels returns a reference to field effective_labels of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) EffectiveLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcfr.ref.Append("effective_labels"))
}

// Id returns a reference to field id of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("ip_address"))
}

// IpProtocol returns a reference to field ip_protocol of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("ip_protocol"))
}

// IpVersion returns a reference to field ip_version of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("ip_version"))
}

// IsMirroringCollector returns a reference to field is_mirroring_collector of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) IsMirroringCollector() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("is_mirroring_collector"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcfr.ref.Append("labels"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("load_balancing_scheme"))
}

// Name returns a reference to field name of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("network"))
}

// NetworkTier returns a reference to field network_tier of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("network_tier"))
}

// NoAutomateDnsZone returns a reference to field no_automate_dns_zone of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) NoAutomateDnsZone() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("no_automate_dns_zone"))
}

// PortRange returns a reference to field port_range of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) PortRange() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("port_range"))
}

// Ports returns a reference to field ports of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Ports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](gcfr.ref.Append("ports"))
}

// Project returns a reference to field project of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("project"))
}

// PscConnectionId returns a reference to field psc_connection_id of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("psc_connection_id"))
}

// PscConnectionStatus returns a reference to field psc_connection_status of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) PscConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("psc_connection_status"))
}

// RecreateClosedPsc returns a reference to field recreate_closed_psc of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) RecreateClosedPsc() terra.BoolValue {
	return terra.ReferenceAsBool(gcfr.ref.Append("recreate_closed_psc"))
}

// Region returns a reference to field region of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("self_link"))
}

// ServiceLabel returns a reference to field service_label of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) ServiceLabel() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("service_label"))
}

// ServiceName returns a reference to field service_name of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("service_name"))
}

// SourceIpRanges returns a reference to field source_ip_ranges of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) SourceIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gcfr.ref.Append("source_ip_ranges"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("subnetwork"))
}

// Target returns a reference to field target of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(gcfr.ref.Append("target"))
}

// TerraformLabels returns a reference to field terraform_labels of google_compute_forwarding_rule.
func (gcfr googleComputeForwardingRuleAttributes) TerraformLabels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcfr.ref.Append("terraform_labels"))
}

func (gcfr googleComputeForwardingRuleAttributes) ServiceDirectoryRegistrations() terra.ListValue[ServiceDirectoryRegistrationsAttributes] {
	return terra.ReferenceAsList[ServiceDirectoryRegistrationsAttributes](gcfr.ref.Append("service_directory_registrations"))
}

func (gcfr googleComputeForwardingRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcfr.ref.Append("timeouts"))
}

type googleComputeForwardingRuleState struct {
	AllPorts                      bool                                 `json:"all_ports"`
	AllowGlobalAccess             bool                                 `json:"allow_global_access"`
	AllowPscGlobalAccess          bool                                 `json:"allow_psc_global_access"`
	BackendService                string                               `json:"backend_service"`
	BaseForwardingRule            string                               `json:"base_forwarding_rule"`
	CreationTimestamp             string                               `json:"creation_timestamp"`
	Description                   string                               `json:"description"`
	EffectiveLabels               map[string]string                    `json:"effective_labels"`
	Id                            string                               `json:"id"`
	IpAddress                     string                               `json:"ip_address"`
	IpProtocol                    string                               `json:"ip_protocol"`
	IpVersion                     string                               `json:"ip_version"`
	IsMirroringCollector          bool                                 `json:"is_mirroring_collector"`
	LabelFingerprint              string                               `json:"label_fingerprint"`
	Labels                        map[string]string                    `json:"labels"`
	LoadBalancingScheme           string                               `json:"load_balancing_scheme"`
	Name                          string                               `json:"name"`
	Network                       string                               `json:"network"`
	NetworkTier                   string                               `json:"network_tier"`
	NoAutomateDnsZone             bool                                 `json:"no_automate_dns_zone"`
	PortRange                     string                               `json:"port_range"`
	Ports                         []string                             `json:"ports"`
	Project                       string                               `json:"project"`
	PscConnectionId               string                               `json:"psc_connection_id"`
	PscConnectionStatus           string                               `json:"psc_connection_status"`
	RecreateClosedPsc             bool                                 `json:"recreate_closed_psc"`
	Region                        string                               `json:"region"`
	SelfLink                      string                               `json:"self_link"`
	ServiceLabel                  string                               `json:"service_label"`
	ServiceName                   string                               `json:"service_name"`
	SourceIpRanges                []string                             `json:"source_ip_ranges"`
	Subnetwork                    string                               `json:"subnetwork"`
	Target                        string                               `json:"target"`
	TerraformLabels               map[string]string                    `json:"terraform_labels"`
	ServiceDirectoryRegistrations []ServiceDirectoryRegistrationsState `json:"service_directory_registrations"`
	Timeouts                      *TimeoutsState                       `json:"timeouts"`
}
