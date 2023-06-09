// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataRoute53ResolverRules creates a new instance of [DataRoute53ResolverRules].
func NewDataRoute53ResolverRules(name string, args DataRoute53ResolverRulesArgs) *DataRoute53ResolverRules {
	return &DataRoute53ResolverRules{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRoute53ResolverRules)(nil)

// DataRoute53ResolverRules represents the Terraform data resource aws_route53_resolver_rules.
type DataRoute53ResolverRules struct {
	Name string
	Args DataRoute53ResolverRulesArgs
}

// DataSource returns the Terraform object type for [DataRoute53ResolverRules].
func (rrr *DataRoute53ResolverRules) DataSource() string {
	return "aws_route53_resolver_rules"
}

// LocalName returns the local name for [DataRoute53ResolverRules].
func (rrr *DataRoute53ResolverRules) LocalName() string {
	return rrr.Name
}

// Configuration returns the configuration (args) for [DataRoute53ResolverRules].
func (rrr *DataRoute53ResolverRules) Configuration() interface{} {
	return rrr.Args
}

// Attributes returns the attributes for [DataRoute53ResolverRules].
func (rrr *DataRoute53ResolverRules) Attributes() dataRoute53ResolverRulesAttributes {
	return dataRoute53ResolverRulesAttributes{ref: terra.ReferenceDataResource(rrr)}
}

// DataRoute53ResolverRulesArgs contains the configurations for aws_route53_resolver_rules.
type DataRoute53ResolverRulesArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// NameRegex: string, optional
	NameRegex terra.StringValue `hcl:"name_regex,attr"`
	// OwnerId: string, optional
	OwnerId terra.StringValue `hcl:"owner_id,attr"`
	// ResolverEndpointId: string, optional
	ResolverEndpointId terra.StringValue `hcl:"resolver_endpoint_id,attr"`
	// RuleType: string, optional
	RuleType terra.StringValue `hcl:"rule_type,attr"`
	// ShareStatus: string, optional
	ShareStatus terra.StringValue `hcl:"share_status,attr"`
}
type dataRoute53ResolverRulesAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("id"))
}

// NameRegex returns a reference to field name_regex of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) NameRegex() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("name_regex"))
}

// OwnerId returns a reference to field owner_id of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) OwnerId() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("owner_id"))
}

// ResolverEndpointId returns a reference to field resolver_endpoint_id of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) ResolverEndpointId() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("resolver_endpoint_id"))
}

// ResolverRuleIds returns a reference to field resolver_rule_ids of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) ResolverRuleIds() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rrr.ref.Append("resolver_rule_ids"))
}

// RuleType returns a reference to field rule_type of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) RuleType() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("rule_type"))
}

// ShareStatus returns a reference to field share_status of aws_route53_resolver_rules.
func (rrr dataRoute53ResolverRulesAttributes) ShareStatus() terra.StringValue {
	return terra.ReferenceAsString(rrr.ref.Append("share_status"))
}
