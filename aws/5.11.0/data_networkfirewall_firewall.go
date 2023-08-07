// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datanetworkfirewallfirewall "github.com/golingon/terraproviders/aws/5.11.0/datanetworkfirewallfirewall"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetworkfirewallFirewall creates a new instance of [DataNetworkfirewallFirewall].
func NewDataNetworkfirewallFirewall(name string, args DataNetworkfirewallFirewallArgs) *DataNetworkfirewallFirewall {
	return &DataNetworkfirewallFirewall{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkfirewallFirewall)(nil)

// DataNetworkfirewallFirewall represents the Terraform data resource aws_networkfirewall_firewall.
type DataNetworkfirewallFirewall struct {
	Name string
	Args DataNetworkfirewallFirewallArgs
}

// DataSource returns the Terraform object type for [DataNetworkfirewallFirewall].
func (nf *DataNetworkfirewallFirewall) DataSource() string {
	return "aws_networkfirewall_firewall"
}

// LocalName returns the local name for [DataNetworkfirewallFirewall].
func (nf *DataNetworkfirewallFirewall) LocalName() string {
	return nf.Name
}

// Configuration returns the configuration (args) for [DataNetworkfirewallFirewall].
func (nf *DataNetworkfirewallFirewall) Configuration() interface{} {
	return nf.Args
}

// Attributes returns the attributes for [DataNetworkfirewallFirewall].
func (nf *DataNetworkfirewallFirewall) Attributes() dataNetworkfirewallFirewallAttributes {
	return dataNetworkfirewallFirewallAttributes{ref: terra.ReferenceDataResource(nf)}
}

// DataNetworkfirewallFirewallArgs contains the configurations for aws_networkfirewall_firewall.
type DataNetworkfirewallFirewallArgs struct {
	// Arn: string, optional
	Arn terra.StringValue `hcl:"arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// EncryptionConfiguration: min=0
	EncryptionConfiguration []datanetworkfirewallfirewall.EncryptionConfiguration `hcl:"encryption_configuration,block" validate:"min=0"`
	// FirewallStatus: min=0
	FirewallStatus []datanetworkfirewallfirewall.FirewallStatus `hcl:"firewall_status,block" validate:"min=0"`
	// SubnetMapping: min=0
	SubnetMapping []datanetworkfirewallfirewall.SubnetMapping `hcl:"subnet_mapping,block" validate:"min=0"`
}
type dataNetworkfirewallFirewallAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("arn"))
}

// DeleteProtection returns a reference to field delete_protection of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) DeleteProtection() terra.BoolValue {
	return terra.ReferenceAsBool(nf.ref.Append("delete_protection"))
}

// Description returns a reference to field description of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("description"))
}

// FirewallPolicyArn returns a reference to field firewall_policy_arn of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) FirewallPolicyArn() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("firewall_policy_arn"))
}

// FirewallPolicyChangeProtection returns a reference to field firewall_policy_change_protection of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) FirewallPolicyChangeProtection() terra.BoolValue {
	return terra.ReferenceAsBool(nf.ref.Append("firewall_policy_change_protection"))
}

// Id returns a reference to field id of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("id"))
}

// Name returns a reference to field name of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("name"))
}

// SubnetChangeProtection returns a reference to field subnet_change_protection of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) SubnetChangeProtection() terra.BoolValue {
	return terra.ReferenceAsBool(nf.ref.Append("subnet_change_protection"))
}

// Tags returns a reference to field tags of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nf.ref.Append("tags"))
}

// UpdateToken returns a reference to field update_token of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) UpdateToken() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("update_token"))
}

// VpcId returns a reference to field vpc_id of aws_networkfirewall_firewall.
func (nf dataNetworkfirewallFirewallAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(nf.ref.Append("vpc_id"))
}

func (nf dataNetworkfirewallFirewallAttributes) EncryptionConfiguration() terra.SetValue[datanetworkfirewallfirewall.EncryptionConfigurationAttributes] {
	return terra.ReferenceAsSet[datanetworkfirewallfirewall.EncryptionConfigurationAttributes](nf.ref.Append("encryption_configuration"))
}

func (nf dataNetworkfirewallFirewallAttributes) FirewallStatus() terra.ListValue[datanetworkfirewallfirewall.FirewallStatusAttributes] {
	return terra.ReferenceAsList[datanetworkfirewallfirewall.FirewallStatusAttributes](nf.ref.Append("firewall_status"))
}

func (nf dataNetworkfirewallFirewallAttributes) SubnetMapping() terra.SetValue[datanetworkfirewallfirewall.SubnetMappingAttributes] {
	return terra.ReferenceAsSet[datanetworkfirewallfirewall.SubnetMappingAttributes](nf.ref.Append("subnet_mapping"))
}
