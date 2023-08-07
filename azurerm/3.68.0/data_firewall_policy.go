// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datafirewallpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/datafirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataFirewallPolicy creates a new instance of [DataFirewallPolicy].
func NewDataFirewallPolicy(name string, args DataFirewallPolicyArgs) *DataFirewallPolicy {
	return &DataFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataFirewallPolicy)(nil)

// DataFirewallPolicy represents the Terraform data resource azurerm_firewall_policy.
type DataFirewallPolicy struct {
	Name string
	Args DataFirewallPolicyArgs
}

// DataSource returns the Terraform object type for [DataFirewallPolicy].
func (fp *DataFirewallPolicy) DataSource() string {
	return "azurerm_firewall_policy"
}

// LocalName returns the local name for [DataFirewallPolicy].
func (fp *DataFirewallPolicy) LocalName() string {
	return fp.Name
}

// Configuration returns the configuration (args) for [DataFirewallPolicy].
func (fp *DataFirewallPolicy) Configuration() interface{} {
	return fp.Args
}

// Attributes returns the attributes for [DataFirewallPolicy].
func (fp *DataFirewallPolicy) Attributes() dataFirewallPolicyAttributes {
	return dataFirewallPolicyAttributes{ref: terra.ReferenceDataResource(fp)}
}

// DataFirewallPolicyArgs contains the configurations for azurerm_firewall_policy.
type DataFirewallPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Dns: min=0
	Dns []datafirewallpolicy.Dns `hcl:"dns,block" validate:"min=0"`
	// ThreatIntelligenceAllowlist: min=0
	ThreatIntelligenceAllowlist []datafirewallpolicy.ThreatIntelligenceAllowlist `hcl:"threat_intelligence_allowlist,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datafirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type dataFirewallPolicyAttributes struct {
	ref terra.Reference
}

// BasePolicyId returns a reference to field base_policy_id of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) BasePolicyId() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("base_policy_id"))
}

// ChildPolicies returns a reference to field child_policies of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) ChildPolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("child_policies"))
}

// Firewalls returns a reference to field firewalls of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) Firewalls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("firewalls"))
}

// Id returns a reference to field id of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("resource_group_name"))
}

// RuleCollectionGroups returns a reference to field rule_collection_groups of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) RuleCollectionGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("rule_collection_groups"))
}

// Tags returns a reference to field tags of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fp.ref.Append("tags"))
}

// ThreatIntelligenceMode returns a reference to field threat_intelligence_mode of azurerm_firewall_policy.
func (fp dataFirewallPolicyAttributes) ThreatIntelligenceMode() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("threat_intelligence_mode"))
}

func (fp dataFirewallPolicyAttributes) Dns() terra.ListValue[datafirewallpolicy.DnsAttributes] {
	return terra.ReferenceAsList[datafirewallpolicy.DnsAttributes](fp.ref.Append("dns"))
}

func (fp dataFirewallPolicyAttributes) ThreatIntelligenceAllowlist() terra.ListValue[datafirewallpolicy.ThreatIntelligenceAllowlistAttributes] {
	return terra.ReferenceAsList[datafirewallpolicy.ThreatIntelligenceAllowlistAttributes](fp.ref.Append("threat_intelligence_allowlist"))
}

func (fp dataFirewallPolicyAttributes) Timeouts() datafirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafirewallpolicy.TimeoutsAttributes](fp.ref.Append("timeouts"))
}
