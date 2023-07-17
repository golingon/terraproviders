// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkfirewallfirewallpolicy "github.com/golingon/terraproviders/aws/5.8.0/datanetworkfirewallfirewallpolicy"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkfirewallFirewallPolicy creates a new instance of [DataNetworkfirewallFirewallPolicy].
func NewDataNetworkfirewallFirewallPolicy(name string, args DataNetworkfirewallFirewallPolicyArgs) *DataNetworkfirewallFirewallPolicy {
	return &DataNetworkfirewallFirewallPolicy{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkfirewallFirewallPolicy)(nil)

// DataNetworkfirewallFirewallPolicy represents the Terraform data resource aws_networkfirewall_firewall_policy.
type DataNetworkfirewallFirewallPolicy struct {
	Name string
	Args DataNetworkfirewallFirewallPolicyArgs
}

// DataSource returns the Terraform object type for [DataNetworkfirewallFirewallPolicy].
func (nfp *DataNetworkfirewallFirewallPolicy) DataSource() string {
	return "aws_networkfirewall_firewall_policy"
}

// LocalName returns the local name for [DataNetworkfirewallFirewallPolicy].
func (nfp *DataNetworkfirewallFirewallPolicy) LocalName() string {
	return nfp.Name
}

// Configuration returns the configuration (args) for [DataNetworkfirewallFirewallPolicy].
func (nfp *DataNetworkfirewallFirewallPolicy) Configuration() interface{} {
	return nfp.Args
}

// Attributes returns the attributes for [DataNetworkfirewallFirewallPolicy].
func (nfp *DataNetworkfirewallFirewallPolicy) Attributes() dataNetworkfirewallFirewallPolicyAttributes {
	return dataNetworkfirewallFirewallPolicyAttributes{ref: terra.ReferenceDataResource(nfp)}
}

// DataNetworkfirewallFirewallPolicyArgs contains the configurations for aws_networkfirewall_firewall_policy.
type DataNetworkfirewallFirewallPolicyArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// FirewallPolicy: min=0
	FirewallPolicy []datanetworkfirewallfirewallpolicy.FirewallPolicy `hcl:"firewall_policy,block" validate:"min=0"`
}
type dataNetworkfirewallFirewallPolicyAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("arn"))
}

// Description returns a reference to field description of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("description"))
}

// Id returns a reference to field id of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("id"))
}

// Name returns a reference to field name of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("name"))
}

// Tags returns a reference to field tags of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nfp.ref.Append("tags"))
}

// UpdateToken returns a reference to field update_token of aws_networkfirewall_firewall_policy.
func (nfp dataNetworkfirewallFirewallPolicyAttributes) UpdateToken() terra.StringValue {
	return terra.ReferenceAsString(nfp.ref.Append("update_token"))
}

func (nfp dataNetworkfirewallFirewallPolicyAttributes) FirewallPolicy() terra.ListValue[datanetworkfirewallfirewallpolicy.FirewallPolicyAttributes] {
	return terra.ReferenceAsList[datanetworkfirewallfirewallpolicy.FirewallPolicyAttributes](nfp.ref.Append("firewall_policy"))
}
