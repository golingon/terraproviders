// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datacomputeglobalforwardingrule "github.com/golingon/terraproviders/google/4.59.0/datacomputeglobalforwardingrule"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataComputeGlobalForwardingRule(name string, args DataComputeGlobalForwardingRuleArgs) *DataComputeGlobalForwardingRule {
	return &DataComputeGlobalForwardingRule{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeGlobalForwardingRule)(nil)

type DataComputeGlobalForwardingRule struct {
	Name string
	Args DataComputeGlobalForwardingRuleArgs
}

func (cgfr *DataComputeGlobalForwardingRule) DataSource() string {
	return "google_compute_global_forwarding_rule"
}

func (cgfr *DataComputeGlobalForwardingRule) LocalName() string {
	return cgfr.Name
}

func (cgfr *DataComputeGlobalForwardingRule) Configuration() interface{} {
	return cgfr.Args
}

func (cgfr *DataComputeGlobalForwardingRule) Attributes() dataComputeGlobalForwardingRuleAttributes {
	return dataComputeGlobalForwardingRuleAttributes{ref: terra.ReferenceDataResource(cgfr)}
}

type DataComputeGlobalForwardingRuleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// MetadataFilters: min=0
	MetadataFilters []datacomputeglobalforwardingrule.MetadataFilters `hcl:"metadata_filters,block" validate:"min=0"`
}
type dataComputeGlobalForwardingRuleAttributes struct {
	ref terra.Reference
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("description"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("id"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) IpAddress() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("ip_address"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) IpProtocol() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("ip_protocol"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) IpVersion() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("ip_version"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) LabelFingerprint() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("label_fingerprint"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](cgfr.ref.Append("labels"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) LoadBalancingScheme() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("load_balancing_scheme"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("name"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Network() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("network"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) PortRange() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("port_range"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Project() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("project"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) PscConnectionId() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("psc_connection_id"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) PscConnectionStatus() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("psc_connection_status"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("self_link"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) Target() terra.StringValue {
	return terra.ReferenceString(cgfr.ref.Append("target"))
}

func (cgfr dataComputeGlobalForwardingRuleAttributes) MetadataFilters() terra.ListValue[datacomputeglobalforwardingrule.MetadataFiltersAttributes] {
	return terra.ReferenceList[datacomputeglobalforwardingrule.MetadataFiltersAttributes](cgfr.ref.Append("metadata_filters"))
}
