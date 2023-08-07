// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacdnfrontdoorfirewallpolicy "github.com/golingon/terraproviders/azurerm/3.68.0/datacdnfrontdoorfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataCdnFrontdoorFirewallPolicy creates a new instance of [DataCdnFrontdoorFirewallPolicy].
func NewDataCdnFrontdoorFirewallPolicy(name string, args DataCdnFrontdoorFirewallPolicyArgs) *DataCdnFrontdoorFirewallPolicy {
	return &DataCdnFrontdoorFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataCdnFrontdoorFirewallPolicy)(nil)

// DataCdnFrontdoorFirewallPolicy represents the Terraform data resource azurerm_cdn_frontdoor_firewall_policy.
type DataCdnFrontdoorFirewallPolicy struct {
	Name string
	Args DataCdnFrontdoorFirewallPolicyArgs
}

// DataSource returns the Terraform object type for [DataCdnFrontdoorFirewallPolicy].
func (cffp *DataCdnFrontdoorFirewallPolicy) DataSource() string {
	return "azurerm_cdn_frontdoor_firewall_policy"
}

// LocalName returns the local name for [DataCdnFrontdoorFirewallPolicy].
func (cffp *DataCdnFrontdoorFirewallPolicy) LocalName() string {
	return cffp.Name
}

// Configuration returns the configuration (args) for [DataCdnFrontdoorFirewallPolicy].
func (cffp *DataCdnFrontdoorFirewallPolicy) Configuration() interface{} {
	return cffp.Args
}

// Attributes returns the attributes for [DataCdnFrontdoorFirewallPolicy].
func (cffp *DataCdnFrontdoorFirewallPolicy) Attributes() dataCdnFrontdoorFirewallPolicyAttributes {
	return dataCdnFrontdoorFirewallPolicyAttributes{ref: terra.ReferenceDataResource(cffp)}
}

// DataCdnFrontdoorFirewallPolicyArgs contains the configurations for azurerm_cdn_frontdoor_firewall_policy.
type DataCdnFrontdoorFirewallPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datacdnfrontdoorfirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type dataCdnFrontdoorFirewallPolicyAttributes struct {
	ref terra.Reference
}

// Enabled returns a reference to field enabled of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cffp.ref.Append("enabled"))
}

// FrontendEndpointIds returns a reference to field frontend_endpoint_ids of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) FrontendEndpointIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cffp.ref.Append("frontend_endpoint_ids"))
}

// Id returns a reference to field id of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("id"))
}

// Mode returns a reference to field mode of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) Mode() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("mode"))
}

// Name returns a reference to field name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("name"))
}

// RedirectUrl returns a reference to field redirect_url of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) RedirectUrl() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("redirect_url"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("resource_group_name"))
}

// SkuName returns a reference to field sku_name of azurerm_cdn_frontdoor_firewall_policy.
func (cffp dataCdnFrontdoorFirewallPolicyAttributes) SkuName() terra.StringValue {
	return terra.ReferenceAsString(cffp.ref.Append("sku_name"))
}

func (cffp dataCdnFrontdoorFirewallPolicyAttributes) Timeouts() datacdnfrontdoorfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacdnfrontdoorfirewallpolicy.TimeoutsAttributes](cffp.ref.Append("timeouts"))
}
