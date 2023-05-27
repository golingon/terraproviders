// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	firewall "github.com/golingon/terraproviders/azurerm/3.58.0/firewall"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewFirewall creates a new instance of [Firewall].
func NewFirewall(name string, args FirewallArgs) *Firewall {
	return &Firewall{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Firewall)(nil)

// Firewall represents the Terraform resource azurerm_firewall.
type Firewall struct {
	Name      string
	Args      FirewallArgs
	state     *firewallState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Firewall].
func (f *Firewall) Type() string {
	return "azurerm_firewall"
}

// LocalName returns the local name for [Firewall].
func (f *Firewall) LocalName() string {
	return f.Name
}

// Configuration returns the configuration (args) for [Firewall].
func (f *Firewall) Configuration() interface{} {
	return f.Args
}

// DependOn is used for other resources to depend on [Firewall].
func (f *Firewall) DependOn() terra.Reference {
	return terra.ReferenceResource(f)
}

// Dependencies returns the list of resources [Firewall] depends_on.
func (f *Firewall) Dependencies() terra.Dependencies {
	return f.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Firewall].
func (f *Firewall) LifecycleManagement() *terra.Lifecycle {
	return f.Lifecycle
}

// Attributes returns the attributes for [Firewall].
func (f *Firewall) Attributes() firewallAttributes {
	return firewallAttributes{ref: terra.ReferenceResource(f)}
}

// ImportState imports the given attribute values into [Firewall]'s state.
func (f *Firewall) ImportState(av io.Reader) error {
	f.state = &firewallState{}
	if err := json.NewDecoder(av).Decode(f.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", f.Type(), f.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Firewall] has state.
func (f *Firewall) State() (*firewallState, bool) {
	return f.state, f.state != nil
}

// StateMust returns the state for [Firewall]. Panics if the state is nil.
func (f *Firewall) StateMust() *firewallState {
	if f.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", f.Type(), f.LocalName()))
	}
	return f.state
}

// FirewallArgs contains the configurations for azurerm_firewall.
type FirewallArgs struct {
	// DnsServers: list of string, optional
	DnsServers terra.ListValue[terra.StringValue] `hcl:"dns_servers,attr"`
	// FirewallPolicyId: string, optional
	FirewallPolicyId terra.StringValue `hcl:"firewall_policy_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PrivateIpRanges: set of string, optional
	PrivateIpRanges terra.SetValue[terra.StringValue] `hcl:"private_ip_ranges,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// SkuName: string, required
	SkuName terra.StringValue `hcl:"sku_name,attr" validate:"required"`
	// SkuTier: string, required
	SkuTier terra.StringValue `hcl:"sku_tier,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// ThreatIntelMode: string, optional
	ThreatIntelMode terra.StringValue `hcl:"threat_intel_mode,attr"`
	// Zones: set of string, optional
	Zones terra.SetValue[terra.StringValue] `hcl:"zones,attr"`
	// IpConfiguration: min=0
	IpConfiguration []firewall.IpConfiguration `hcl:"ip_configuration,block" validate:"min=0"`
	// ManagementIpConfiguration: optional
	ManagementIpConfiguration *firewall.ManagementIpConfiguration `hcl:"management_ip_configuration,block"`
	// Timeouts: optional
	Timeouts *firewall.Timeouts `hcl:"timeouts,block"`
	// VirtualHub: optional
	VirtualHub *firewall.VirtualHub `hcl:"virtual_hub,block"`
}
type firewallAttributes struct {
	ref terra.Reference
}

// DnsServers returns a reference to field dns_servers of azurerm_firewall.
func (f firewallAttributes) DnsServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](f.ref.Append("dns_servers"))
}

// FirewallPolicyId returns a reference to field firewall_policy_id of azurerm_firewall.
func (f firewallAttributes) FirewallPolicyId() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("firewall_policy_id"))
}

// Id returns a reference to field id of azurerm_firewall.
func (f firewallAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_firewall.
func (f firewallAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_firewall.
func (f firewallAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("name"))
}

// PrivateIpRanges returns a reference to field private_ip_ranges of azurerm_firewall.
func (f firewallAttributes) PrivateIpRanges() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("private_ip_ranges"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_firewall.
func (f firewallAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_firewall.
func (f firewallAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("sku_name"))
}

// SkuTier returns a reference to field sku_tier of azurerm_firewall.
func (f firewallAttributes) SkuTier() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("sku_tier"))
}

// Tags returns a reference to field tags of azurerm_firewall.
func (f firewallAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](f.ref.Append("tags"))
}

// ThreatIntelMode returns a reference to field threat_intel_mode of azurerm_firewall.
func (f firewallAttributes) ThreatIntelMode() terra.StringValue {
	return terra.ReferenceAsString(f.ref.Append("threat_intel_mode"))
}

// Zones returns a reference to field zones of azurerm_firewall.
func (f firewallAttributes) Zones() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](f.ref.Append("zones"))
}

func (f firewallAttributes) IpConfiguration() terra.ListValue[firewall.IpConfigurationAttributes] {
	return terra.ReferenceAsList[firewall.IpConfigurationAttributes](f.ref.Append("ip_configuration"))
}

func (f firewallAttributes) ManagementIpConfiguration() terra.ListValue[firewall.ManagementIpConfigurationAttributes] {
	return terra.ReferenceAsList[firewall.ManagementIpConfigurationAttributes](f.ref.Append("management_ip_configuration"))
}

func (f firewallAttributes) Timeouts() firewall.TimeoutsAttributes {
	return terra.ReferenceAsSingle[firewall.TimeoutsAttributes](f.ref.Append("timeouts"))
}

func (f firewallAttributes) VirtualHub() terra.ListValue[firewall.VirtualHubAttributes] {
	return terra.ReferenceAsList[firewall.VirtualHubAttributes](f.ref.Append("virtual_hub"))
}

type firewallState struct {
	DnsServers                []string                                  `json:"dns_servers"`
	FirewallPolicyId          string                                    `json:"firewall_policy_id"`
	Id                        string                                    `json:"id"`
	Location                  string                                    `json:"location"`
	Name                      string                                    `json:"name"`
	PrivateIpRanges           []string                                  `json:"private_ip_ranges"`
	ResourceGroupName         string                                    `json:"resource_group_name"`
	SkuName                   string                                    `json:"sku_name"`
	SkuTier                   string                                    `json:"sku_tier"`
	Tags                      map[string]string                         `json:"tags"`
	ThreatIntelMode           string                                    `json:"threat_intel_mode"`
	Zones                     []string                                  `json:"zones"`
	IpConfiguration           []firewall.IpConfigurationState           `json:"ip_configuration"`
	ManagementIpConfiguration []firewall.ManagementIpConfigurationState `json:"management_ip_configuration"`
	Timeouts                  *firewall.TimeoutsState                   `json:"timeouts"`
	VirtualHub                []firewall.VirtualHubState                `json:"virtual_hub"`
}
