// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	firewallpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/firewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirewallPolicy creates a new instance of [FirewallPolicy].
func NewFirewallPolicy(name string, args FirewallPolicyArgs) *FirewallPolicy {
	return &FirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*FirewallPolicy)(nil)

// FirewallPolicy represents the Terraform resource azurerm_firewall_policy.
type FirewallPolicy struct {
	Name      string
	Args      FirewallPolicyArgs
	state     *firewallPolicyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [FirewallPolicy].
func (fp *FirewallPolicy) Type() string {
	return "azurerm_firewall_policy"
}

// LocalName returns the local name for [FirewallPolicy].
func (fp *FirewallPolicy) LocalName() string {
	return fp.Name
}

// Configuration returns the configuration (args) for [FirewallPolicy].
func (fp *FirewallPolicy) Configuration() interface{} {
	return fp.Args
}

// DependOn is used for other resources to depend on [FirewallPolicy].
func (fp *FirewallPolicy) DependOn() terra.Reference {
	return terra.ReferenceResource(fp)
}

// Dependencies returns the list of resources [FirewallPolicy] depends_on.
func (fp *FirewallPolicy) Dependencies() terra.Dependencies {
	return fp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [FirewallPolicy].
func (fp *FirewallPolicy) LifecycleManagement() *terra.Lifecycle {
	return fp.Lifecycle
}

// Attributes returns the attributes for [FirewallPolicy].
func (fp *FirewallPolicy) Attributes() firewallPolicyAttributes {
	return firewallPolicyAttributes{ref: terra.ReferenceResource(fp)}
}

// ImportState imports the given attribute values into [FirewallPolicy]'s state.
func (fp *FirewallPolicy) ImportState(av io.Reader) error {
	fp.state = &firewallPolicyState{}
	if err := json.NewDecoder(av).Decode(fp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", fp.Type(), fp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [FirewallPolicy] has state.
func (fp *FirewallPolicy) State() (*firewallPolicyState, bool) {
	return fp.state, fp.state != nil
}

// StateMust returns the state for [FirewallPolicy]. Panics if the state is nil.
func (fp *FirewallPolicy) StateMust() *firewallPolicyState {
	if fp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", fp.Type(), fp.LocalName()))
	}
	return fp.state
}

// FirewallPolicyArgs contains the configurations for azurerm_firewall_policy.
type FirewallPolicyArgs struct {
	// AutoLearnPrivateRangesEnabled: bool, optional
	AutoLearnPrivateRangesEnabled terra.BoolValue `hcl:"auto_learn_private_ranges_enabled,attr"`
	// BasePolicyId: string, optional
	BasePolicyId terra.StringValue `hcl:"base_policy_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateIpRanges: list of string, optional
	PrivateIpRanges terra.ListValue[terra.StringValue] `hcl:"private_ip_ranges,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Sku: string, optional
	Sku terra.StringValue `hcl:"sku,attr"`
	// SqlRedirectAllowed: bool, optional
	SqlRedirectAllowed terra.BoolValue `hcl:"sql_redirect_allowed,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ThreatIntelligenceMode: string, optional
	ThreatIntelligenceMode terra.StringValue `hcl:"threat_intelligence_mode,attr"`
	// Dns: optional
	Dns *firewallpolicy.Dns `hcl:"dns,block"`
	// ExplicitProxy: optional
	ExplicitProxy *firewallpolicy.ExplicitProxy `hcl:"explicit_proxy,block"`
	// Identity: optional
	Identity *firewallpolicy.Identity `hcl:"identity,block"`
	// Insights: optional
	Insights *firewallpolicy.Insights `hcl:"insights,block"`
	// IntrusionDetection: optional
	IntrusionDetection *firewallpolicy.IntrusionDetection `hcl:"intrusion_detection,block"`
	// ThreatIntelligenceAllowlist: optional
	ThreatIntelligenceAllowlist *firewallpolicy.ThreatIntelligenceAllowlist `hcl:"threat_intelligence_allowlist,block"`
	// Timeouts: optional
	Timeouts *firewallpolicy.Timeouts `hcl:"timeouts,block"`
	// TlsCertificate: optional
	TlsCertificate *firewallpolicy.TlsCertificate `hcl:"tls_certificate,block"`
}
type firewallPolicyAttributes struct {
	ref terra.Reference
}

// AutoLearnPrivateRangesEnabled returns a reference to field auto_learn_private_ranges_enabled of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) AutoLearnPrivateRangesEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(fp.ref.Append("auto_learn_private_ranges_enabled"))
}

// BasePolicyId returns a reference to field base_policy_id of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) BasePolicyId() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("base_policy_id"))
}

// ChildPolicies returns a reference to field child_policies of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) ChildPolicies() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("child_policies"))
}

// Firewalls returns a reference to field firewalls of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Firewalls() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("firewalls"))
}

// Id returns a reference to field id of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("name"))
}

// PrivateIpRanges returns a reference to field private_ip_ranges of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) PrivateIpRanges() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("private_ip_ranges"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("resource_group_name"))
}

// RuleCollectionGroups returns a reference to field rule_collection_groups of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) RuleCollectionGroups() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](fp.ref.Append("rule_collection_groups"))
}

// Sku returns a reference to field sku of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("sku"))
}

// SqlRedirectAllowed returns a reference to field sql_redirect_allowed of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) SqlRedirectAllowed() terra.BoolValue {
	return terra.ReferenceAsBool(fp.ref.Append("sql_redirect_allowed"))
}

// Tags returns a reference to field tags of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](fp.ref.Append("tags"))
}

// ThreatIntelligenceMode returns a reference to field threat_intelligence_mode of azurerm_firewall_policy.
func (fp firewallPolicyAttributes) ThreatIntelligenceMode() terra.StringValue {
	return terra.ReferenceAsString(fp.ref.Append("threat_intelligence_mode"))
}

func (fp firewallPolicyAttributes) Dns() terra.ListValue[firewallpolicy.DnsAttributes] {
	return terra.ReferenceAsList[firewallpolicy.DnsAttributes](fp.ref.Append("dns"))
}

func (fp firewallPolicyAttributes) ExplicitProxy() terra.ListValue[firewallpolicy.ExplicitProxyAttributes] {
	return terra.ReferenceAsList[firewallpolicy.ExplicitProxyAttributes](fp.ref.Append("explicit_proxy"))
}

func (fp firewallPolicyAttributes) Identity() terra.ListValue[firewallpolicy.IdentityAttributes] {
	return terra.ReferenceAsList[firewallpolicy.IdentityAttributes](fp.ref.Append("identity"))
}

func (fp firewallPolicyAttributes) Insights() terra.ListValue[firewallpolicy.InsightsAttributes] {
	return terra.ReferenceAsList[firewallpolicy.InsightsAttributes](fp.ref.Append("insights"))
}

func (fp firewallPolicyAttributes) IntrusionDetection() terra.ListValue[firewallpolicy.IntrusionDetectionAttributes] {
	return terra.ReferenceAsList[firewallpolicy.IntrusionDetectionAttributes](fp.ref.Append("intrusion_detection"))
}

func (fp firewallPolicyAttributes) ThreatIntelligenceAllowlist() terra.ListValue[firewallpolicy.ThreatIntelligenceAllowlistAttributes] {
	return terra.ReferenceAsList[firewallpolicy.ThreatIntelligenceAllowlistAttributes](fp.ref.Append("threat_intelligence_allowlist"))
}

func (fp firewallPolicyAttributes) Timeouts() firewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firewallpolicy.TimeoutsAttributes](fp.ref.Append("timeouts"))
}

func (fp firewallPolicyAttributes) TlsCertificate() terra.ListValue[firewallpolicy.TlsCertificateAttributes] {
	return terra.ReferenceAsList[firewallpolicy.TlsCertificateAttributes](fp.ref.Append("tls_certificate"))
}

type firewallPolicyState struct {
	AutoLearnPrivateRangesEnabled bool                                              `json:"auto_learn_private_ranges_enabled"`
	BasePolicyId                  string                                            `json:"base_policy_id"`
	ChildPolicies                 []string                                          `json:"child_policies"`
	Firewalls                     []string                                          `json:"firewalls"`
	Id                            string                                            `json:"id"`
	Location                      string                                            `json:"location"`
	Name                          string                                            `json:"name"`
	PrivateIpRanges               []string                                          `json:"private_ip_ranges"`
	ResourceGroupName             string                                            `json:"resource_group_name"`
	RuleCollectionGroups          []string                                          `json:"rule_collection_groups"`
	Sku                           string                                            `json:"sku"`
	SqlRedirectAllowed            bool                                              `json:"sql_redirect_allowed"`
	Tags                          map[string]string                                 `json:"tags"`
	ThreatIntelligenceMode        string                                            `json:"threat_intelligence_mode"`
	Dns                           []firewallpolicy.DnsState                         `json:"dns"`
	ExplicitProxy                 []firewallpolicy.ExplicitProxyState               `json:"explicit_proxy"`
	Identity                      []firewallpolicy.IdentityState                    `json:"identity"`
	Insights                      []firewallpolicy.InsightsState                    `json:"insights"`
	IntrusionDetection            []firewallpolicy.IntrusionDetectionState          `json:"intrusion_detection"`
	ThreatIntelligenceAllowlist   []firewallpolicy.ThreatIntelligenceAllowlistState `json:"threat_intelligence_allowlist"`
	Timeouts                      *firewallpolicy.TimeoutsState                     `json:"timeouts"`
	TlsCertificate                []firewallpolicy.TlsCertificateState              `json:"tls_certificate"`
}