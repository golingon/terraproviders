// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53_resolver_firewall_domain_list

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_route53_resolver_firewall_domain_list.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (arrfdl *DataSource) DataSource() string {
	return "aws_route53_resolver_firewall_domain_list"
}

// LocalName returns the local name for [DataSource].
func (arrfdl *DataSource) LocalName() string {
	return arrfdl.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (arrfdl *DataSource) Configuration() interface{} {
	return arrfdl.Args
}

// Attributes returns the attributes for [DataSource].
func (arrfdl *DataSource) Attributes() dataAwsRoute53ResolverFirewallDomainListAttributes {
	return dataAwsRoute53ResolverFirewallDomainListAttributes{ref: terra.ReferenceDataSource(arrfdl)}
}

// DataArgs contains the configurations for aws_route53_resolver_firewall_domain_list.
type DataArgs struct {
	// FirewallDomainListId: string, required
	FirewallDomainListId terra.StringValue `hcl:"firewall_domain_list_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
}

type dataAwsRoute53ResolverFirewallDomainListAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("arn"))
}

// CreationTime returns a reference to field creation_time of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) CreationTime() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("creation_time"))
}

// CreatorRequestId returns a reference to field creator_request_id of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) CreatorRequestId() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("creator_request_id"))
}

// DomainCount returns a reference to field domain_count of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) DomainCount() terra.NumberValue {
	return terra.ReferenceAsNumber(arrfdl.ref.Append("domain_count"))
}

// FirewallDomainListId returns a reference to field firewall_domain_list_id of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) FirewallDomainListId() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("firewall_domain_list_id"))
}

// Id returns a reference to field id of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("id"))
}

// ManagedOwnerName returns a reference to field managed_owner_name of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) ManagedOwnerName() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("managed_owner_name"))
}

// ModificationTime returns a reference to field modification_time of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) ModificationTime() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("modification_time"))
}

// Name returns a reference to field name of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("name"))
}

// Status returns a reference to field status of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("status"))
}

// StatusMessage returns a reference to field status_message of aws_route53_resolver_firewall_domain_list.
func (arrfdl dataAwsRoute53ResolverFirewallDomainListAttributes) StatusMessage() terra.StringValue {
	return terra.ReferenceAsString(arrfdl.ref.Append("status_message"))
}
