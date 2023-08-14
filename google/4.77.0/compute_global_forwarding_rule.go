// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computeglobalforwardingrule "github.com/golingon/terraproviders/google/4.77.0/computeglobalforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeGlobalForwardingRule creates a new instance of [ComputeGlobalForwardingRule].
func NewComputeGlobalForwardingRule(name string, args ComputeGlobalForwardingRuleArgs) *ComputeGlobalForwardingRule {
	return &ComputeGlobalForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeGlobalForwardingRule)(nil)

// ComputeGlobalForwardingRule represents the Terraform resource google_compute_global_forwarding_rule.
type ComputeGlobalForwardingRule struct {
	Name      string
	Args      ComputeGlobalForwardingRuleArgs
	state     *computeGlobalForwardingRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) Type() string {
	return "google_compute_global_forwarding_rule"
}

// LocalName returns the local name for [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) LocalName() string {
	return cgfr.Name
}

// Configuration returns the configuration (args) for [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) Configuration() interface{} {
	return cgfr.Args
}

// DependOn is used for other resources to depend on [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cgfr)
}

// Dependencies returns the list of resources [ComputeGlobalForwardingRule] depends_on.
func (cgfr *ComputeGlobalForwardingRule) Dependencies() terra.Dependencies {
	return cgfr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) LifecycleManagement() *terra.Lifecycle {
	return cgfr.Lifecycle
}

// Attributes returns the attributes for [ComputeGlobalForwardingRule].
func (cgfr *ComputeGlobalForwardingRule) Attributes() computeGlobalForwardingRuleAttributes {
	return computeGlobalForwardingRuleAttributes{ref: terra.ReferenceResource(cgfr)}
}

// ImportState imports the given attribute values into [ComputeGlobalForwardingRule]'s state.
func (cgfr *ComputeGlobalForwardingRule) ImportState(av io.Reader) error {
	cgfr.state = &computeGlobalForwardingRuleState{}
	if err := json.NewDecoder(av).Decode(cgfr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cgfr.Type(), cgfr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeGlobalForwardingRule] has state.
func (cgfr *ComputeGlobalForwardingRule) State() (*computeGlobalForwardingRuleState, bool) {
	return cgfr.state, cgfr.state != nil
}

// StateMust returns the state for [ComputeGlobalForwardingRule]. Panics if the state is nil.
func (cgfr *ComputeGlobalForwardingRule) StateMust() *computeGlobalForwardingRuleState {
	if cgfr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cgfr.Type(), cgfr.LocalName()))
	}
	return cgfr.state
}

// ComputeGlobalForwardingRuleArgs contains the configurations for google_compute_global_forwarding_rule.
type ComputeGlobalForwardingRuleArgs struct {
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
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LoadBalancingScheme: string, optional
	LoadBalancingScheme terra.StringValue `hcl:"load_balancing_scheme,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Network: string, optional
	Network terra.StringValue `hcl:"network,attr"`
	// NoAutomateDnsZone: bool, optional
	NoAutomateDnsZone terra.BoolValue `hcl:"no_automate_dns_zone,attr"`
	// PortRange: string, optional
	PortRange terra.StringValue `hcl:"port_range,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// SourceIpRanges: list of string, optional
	SourceIpRanges terra.ListValue[terra.StringValue] `hcl:"source_ip_ranges,attr"`
	// Target: string, required
	Target terra.StringValue `hcl:"target,attr" validate:"required"`
	// MetadataFilters: min=0
	MetadataFilters []computeglobalforwardingrule.MetadataFilters `hcl:"metadata_filters,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *computeglobalforwardingrule.Timeouts `hcl:"timeouts,block"`
}
type computeGlobalForwardingRuleAttributes struct {
	ref terra.Reference
}

// BaseForwardingRule returns a reference to field base_forwarding_rule of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) BaseForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("base_forwarding_rule"))
}

// Description returns a reference to field description of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("ip_address"))
}

// IpProtocol returns a reference to field ip_protocol of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("ip_protocol"))
}

// IpVersion returns a reference to field ip_version of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("ip_version"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cgfr.ref.Append("labels"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("load_balancing_scheme"))
}

// Name returns a reference to field name of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("network"))
}

// NoAutomateDnsZone returns a reference to field no_automate_dns_zone of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) NoAutomateDnsZone() terra.BoolValue {
	return terra.ReferenceAsBool(cgfr.ref.Append("no_automate_dns_zone"))
}

// PortRange returns a reference to field port_range of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) PortRange() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("port_range"))
}

// Project returns a reference to field project of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("project"))
}

// PscConnectionId returns a reference to field psc_connection_id of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("psc_connection_id"))
}

// PscConnectionStatus returns a reference to field psc_connection_status of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) PscConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("psc_connection_status"))
}

// SelfLink returns a reference to field self_link of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("self_link"))
}

// SourceIpRanges returns a reference to field source_ip_ranges of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) SourceIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cgfr.ref.Append("source_ip_ranges"))
}

// Target returns a reference to field target of google_compute_global_forwarding_rule.
func (cgfr computeGlobalForwardingRuleAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(cgfr.ref.Append("target"))
}

func (cgfr computeGlobalForwardingRuleAttributes) MetadataFilters() terra.ListValue[computeglobalforwardingrule.MetadataFiltersAttributes] {
	return terra.ReferenceAsList[computeglobalforwardingrule.MetadataFiltersAttributes](cgfr.ref.Append("metadata_filters"))
}

func (cgfr computeGlobalForwardingRuleAttributes) Timeouts() computeglobalforwardingrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computeglobalforwardingrule.TimeoutsAttributes](cgfr.ref.Append("timeouts"))
}

type computeGlobalForwardingRuleState struct {
	BaseForwardingRule  string                                             `json:"base_forwarding_rule"`
	Description         string                                             `json:"description"`
	Id                  string                                             `json:"id"`
	IpAddress           string                                             `json:"ip_address"`
	IpProtocol          string                                             `json:"ip_protocol"`
	IpVersion           string                                             `json:"ip_version"`
	LabelFingerprint    string                                             `json:"label_fingerprint"`
	Labels              map[string]string                                  `json:"labels"`
	LoadBalancingScheme string                                             `json:"load_balancing_scheme"`
	Name                string                                             `json:"name"`
	Network             string                                             `json:"network"`
	NoAutomateDnsZone   bool                                               `json:"no_automate_dns_zone"`
	PortRange           string                                             `json:"port_range"`
	Project             string                                             `json:"project"`
	PscConnectionId     string                                             `json:"psc_connection_id"`
	PscConnectionStatus string                                             `json:"psc_connection_status"`
	SelfLink            string                                             `json:"self_link"`
	SourceIpRanges      []string                                           `json:"source_ip_ranges"`
	Target              string                                             `json:"target"`
	MetadataFilters     []computeglobalforwardingrule.MetadataFiltersState `json:"metadata_filters"`
	Timeouts            *computeglobalforwardingrule.TimeoutsState         `json:"timeouts"`
}
