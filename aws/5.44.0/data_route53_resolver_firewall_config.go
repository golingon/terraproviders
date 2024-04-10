// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import "github.com/golingon/lingon/pkg/terra"

// NewDataRoute53ResolverFirewallConfig creates a new instance of [DataRoute53ResolverFirewallConfig].
func NewDataRoute53ResolverFirewallConfig(name string, args DataRoute53ResolverFirewallConfigArgs) *DataRoute53ResolverFirewallConfig {
	return &DataRoute53ResolverFirewallConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53ResolverFirewallConfig)(nil)

// DataRoute53ResolverFirewallConfig represents the Terraform data resource aws_route53_resolver_firewall_config.
type DataRoute53ResolverFirewallConfig struct {
	Name string
	Args DataRoute53ResolverFirewallConfigArgs
}

// DataSource returns the Terraform object type for [DataRoute53ResolverFirewallConfig].
func (rrfc *DataRoute53ResolverFirewallConfig) DataSource() string {
	return "aws_route53_resolver_firewall_config"
}

// LocalName returns the local name for [DataRoute53ResolverFirewallConfig].
func (rrfc *DataRoute53ResolverFirewallConfig) LocalName() string {
	return rrfc.Name
}

// Configuration returns the configuration (args) for [DataRoute53ResolverFirewallConfig].
func (rrfc *DataRoute53ResolverFirewallConfig) Configuration() interface{} {
	return rrfc.Args
}

// Attributes returns the attributes for [DataRoute53ResolverFirewallConfig].
func (rrfc *DataRoute53ResolverFirewallConfig) Attributes() dataRoute53ResolverFirewallConfigAttributes {
	return dataRoute53ResolverFirewallConfigAttributes{ref: terra.ReferenceDataResource(rrfc)}
}

// DataRoute53ResolverFirewallConfigArgs contains the configurations for aws_route53_resolver_firewall_config.
type DataRoute53ResolverFirewallConfigArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
}
type dataRoute53ResolverFirewallConfigAttributes struct {
	ref terra.Reference
}

// FirewallFailOpen returns a reference to field firewall_fail_open of aws_route53_resolver_firewall_config.
func (rrfc dataRoute53ResolverFirewallConfigAttributes) FirewallFailOpen() terra.StringValue {
	return terra.ReferenceAsString(rrfc.ref.Append("firewall_fail_open"))
}

// Id returns a reference to field id of aws_route53_resolver_firewall_config.
func (rrfc dataRoute53ResolverFirewallConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrfc.ref.Append("id"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_firewall_config.
func (rrfc dataRoute53ResolverFirewallConfigAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rrfc.ref.Append("owner_id"))
}

// ResourceId returns a reference to field resource_id of aws_route53_resolver_firewall_config.
func (rrfc dataRoute53ResolverFirewallConfigAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(rrfc.ref.Append("resource_id"))
}
