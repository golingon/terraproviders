// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeforwardingrule "github.com/golingon/terraproviders/google/4.72.1/datacomputeforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataComputeForwardingRule creates a new instance of [DataComputeForwardingRule].
func NewDataComputeForwardingRule(name string, args DataComputeForwardingRuleArgs) *DataComputeForwardingRule {
	return &DataComputeForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeForwardingRule)(nil)

// DataComputeForwardingRule represents the Terraform data resource google_compute_forwarding_rule.
type DataComputeForwardingRule struct {
	Name string
	Args DataComputeForwardingRuleArgs
}

// DataSource returns the Terraform object type for [DataComputeForwardingRule].
func (cfr *DataComputeForwardingRule) DataSource() string {
	return "google_compute_forwarding_rule"
}

// LocalName returns the local name for [DataComputeForwardingRule].
func (cfr *DataComputeForwardingRule) LocalName() string {
	return cfr.Name
}

// Configuration returns the configuration (args) for [DataComputeForwardingRule].
func (cfr *DataComputeForwardingRule) Configuration() interface{} {
	return cfr.Args
}

// Attributes returns the attributes for [DataComputeForwardingRule].
func (cfr *DataComputeForwardingRule) Attributes() dataComputeForwardingRuleAttributes {
	return dataComputeForwardingRuleAttributes{ref: terra.ReferenceDataResource(cfr)}
}

// DataComputeForwardingRuleArgs contains the configurations for google_compute_forwarding_rule.
type DataComputeForwardingRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// ServiceDirectoryRegistrations: min=0
	ServiceDirectoryRegistrations []datacomputeforwardingrule.ServiceDirectoryRegistrations `hcl:"service_directory_registrations,block" validate:"min=0"`
}
type dataComputeForwardingRuleAttributes struct {
	ref terra.Reference
}

// AllPorts returns a reference to field all_ports of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) AllPorts() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("all_ports"))
}

// AllowGlobalAccess returns a reference to field allow_global_access of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) AllowGlobalAccess() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("allow_global_access"))
}

// AllowPscGlobalAccess returns a reference to field allow_psc_global_access of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) AllowPscGlobalAccess() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("allow_psc_global_access"))
}

// BackendService returns a reference to field backend_service of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("backend_service"))
}

// BaseForwardingRule returns a reference to field base_forwarding_rule of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) BaseForwardingRule() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("base_forwarding_rule"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("creation_timestamp"))
}

// Description returns a reference to field description of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("description"))
}

// Id returns a reference to field id of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("id"))
}

// IpAddress returns a reference to field ip_address of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("ip_address"))
}

// IpProtocol returns a reference to field ip_protocol of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("ip_protocol"))
}

// IsMirroringCollector returns a reference to field is_mirroring_collector of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) IsMirroringCollector() terra.BoolValue {
	return terra.ReferenceAsBool(cfr.ref.Append("is_mirroring_collector"))
}

// LabelFingerprint returns a reference to field label_fingerprint of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("label_fingerprint"))
}

// Labels returns a reference to field labels of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cfr.ref.Append("labels"))
}

// LoadBalancingScheme returns a reference to field load_balancing_scheme of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("load_balancing_scheme"))
}

// Name returns a reference to field name of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("name"))
}

// Network returns a reference to field network of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Network() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("network"))
}

// NetworkTier returns a reference to field network_tier of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) NetworkTier() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("network_tier"))
}

// PortRange returns a reference to field port_range of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) PortRange() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("port_range"))
}

// Ports returns a reference to field ports of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Ports() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cfr.ref.Append("ports"))
}

// Project returns a reference to field project of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("project"))
}

// PscConnectionId returns a reference to field psc_connection_id of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("psc_connection_id"))
}

// PscConnectionStatus returns a reference to field psc_connection_status of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) PscConnectionStatus() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("psc_connection_status"))
}

// Region returns a reference to field region of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("region"))
}

// SelfLink returns a reference to field self_link of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("self_link"))
}

// ServiceLabel returns a reference to field service_label of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) ServiceLabel() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("service_label"))
}

// ServiceName returns a reference to field service_name of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("service_name"))
}

// SourceIpRanges returns a reference to field source_ip_ranges of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) SourceIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cfr.ref.Append("source_ip_ranges"))
}

// Subnetwork returns a reference to field subnetwork of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Subnetwork() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("subnetwork"))
}

// Target returns a reference to field target of google_compute_forwarding_rule.
func (cfr dataComputeForwardingRuleAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(cfr.ref.Append("target"))
}

func (cfr dataComputeForwardingRuleAttributes) ServiceDirectoryRegistrations() terra.ListValue[datacomputeforwardingrule.ServiceDirectoryRegistrationsAttributes] {
	return terra.ReferenceAsList[datacomputeforwardingrule.ServiceDirectoryRegistrationsAttributes](cfr.ref.Append("service_directory_registrations"))
}
