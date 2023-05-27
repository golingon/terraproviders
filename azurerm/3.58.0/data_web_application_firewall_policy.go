// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datawebapplicationfirewallpolicy "github.com/golingon/terraproviders/azurerm/3.58.0/datawebapplicationfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataWebApplicationFirewallPolicy creates a new instance of [DataWebApplicationFirewallPolicy].
func NewDataWebApplicationFirewallPolicy(name string, args DataWebApplicationFirewallPolicyArgs) *DataWebApplicationFirewallPolicy {
	return &DataWebApplicationFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataWebApplicationFirewallPolicy)(nil)

// DataWebApplicationFirewallPolicy represents the Terraform data resource azurerm_web_application_firewall_policy.
type DataWebApplicationFirewallPolicy struct {
	Name string
	Args DataWebApplicationFirewallPolicyArgs
}

// DataSource returns the Terraform object type for [DataWebApplicationFirewallPolicy].
func (wafp *DataWebApplicationFirewallPolicy) DataSource() string {
	return "azurerm_web_application_firewall_policy"
}

// LocalName returns the local name for [DataWebApplicationFirewallPolicy].
func (wafp *DataWebApplicationFirewallPolicy) LocalName() string {
	return wafp.Name
}

// Configuration returns the configuration (args) for [DataWebApplicationFirewallPolicy].
func (wafp *DataWebApplicationFirewallPolicy) Configuration() interface{} {
	return wafp.Args
}

// Attributes returns the attributes for [DataWebApplicationFirewallPolicy].
func (wafp *DataWebApplicationFirewallPolicy) Attributes() dataWebApplicationFirewallPolicyAttributes {
	return dataWebApplicationFirewallPolicyAttributes{ref: terra.ReferenceDataResource(wafp)}
}

// DataWebApplicationFirewallPolicyArgs contains the configurations for azurerm_web_application_firewall_policy.
type DataWebApplicationFirewallPolicyArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *datawebapplicationfirewallpolicy.Timeouts `hcl:"timeouts,block"`
}
type dataWebApplicationFirewallPolicyAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_web_application_firewall_policy.
func (wafp dataWebApplicationFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(wafp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_web_application_firewall_policy.
func (wafp dataWebApplicationFirewallPolicyAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(wafp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_web_application_firewall_policy.
func (wafp dataWebApplicationFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(wafp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_web_application_firewall_policy.
func (wafp dataWebApplicationFirewallPolicyAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(wafp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_web_application_firewall_policy.
func (wafp dataWebApplicationFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](wafp.ref.Append("tags"))
}

func (wafp dataWebApplicationFirewallPolicyAttributes) Timeouts() datawebapplicationfirewallpolicy.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datawebapplicationfirewallpolicy.TimeoutsAttributes](wafp.ref.Append("timeouts"))
}
