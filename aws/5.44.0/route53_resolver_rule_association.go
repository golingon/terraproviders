// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	route53resolverruleassociation "github.com/golingon/terraproviders/aws/5.44.0/route53resolverruleassociation"
	"io"
)

// NewRoute53ResolverRuleAssociation creates a new instance of [Route53ResolverRuleAssociation].
func NewRoute53ResolverRuleAssociation(name string, args Route53ResolverRuleAssociationArgs) *Route53ResolverRuleAssociation {
	return &Route53ResolverRuleAssociation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53ResolverRuleAssociation)(nil)

// Route53ResolverRuleAssociation represents the Terraform resource aws_route53_resolver_rule_association.
type Route53ResolverRuleAssociation struct {
	Name      string
	Args      Route53ResolverRuleAssociationArgs
	state     *route53ResolverRuleAssociationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) Type() string {
	return "aws_route53_resolver_rule_association"
}

// LocalName returns the local name for [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) LocalName() string {
	return rrra.Name
}

// Configuration returns the configuration (args) for [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) Configuration() interface{} {
	return rrra.Args
}

// DependOn is used for other resources to depend on [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) DependOn() terra.Reference {
	return terra.ReferenceResource(rrra)
}

// Dependencies returns the list of resources [Route53ResolverRuleAssociation] depends_on.
func (rrra *Route53ResolverRuleAssociation) Dependencies() terra.Dependencies {
	return rrra.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) LifecycleManagement() *terra.Lifecycle {
	return rrra.Lifecycle
}

// Attributes returns the attributes for [Route53ResolverRuleAssociation].
func (rrra *Route53ResolverRuleAssociation) Attributes() route53ResolverRuleAssociationAttributes {
	return route53ResolverRuleAssociationAttributes{ref: terra.ReferenceResource(rrra)}
}

// ImportState imports the given attribute values into [Route53ResolverRuleAssociation]'s state.
func (rrra *Route53ResolverRuleAssociation) ImportState(av io.Reader) error {
	rrra.state = &route53ResolverRuleAssociationState{}
	if err := json.NewDecoder(av).Decode(rrra.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrra.Type(), rrra.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53ResolverRuleAssociation] has state.
func (rrra *Route53ResolverRuleAssociation) State() (*route53ResolverRuleAssociationState, bool) {
	return rrra.state, rrra.state != nil
}

// StateMust returns the state for [Route53ResolverRuleAssociation]. Panics if the state is nil.
func (rrra *Route53ResolverRuleAssociation) StateMust() *route53ResolverRuleAssociationState {
	if rrra.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrra.Type(), rrra.LocalName()))
	}
	return rrra.state
}

// Route53ResolverRuleAssociationArgs contains the configurations for aws_route53_resolver_rule_association.
type Route53ResolverRuleAssociationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// ResolverRuleId: string, required
	ResolverRuleId terra.StringValue `hcl:"resolver_rule_id,attr" validate:"required"`
	// VpcId: string, required
	VpcId terra.StringValue `hcl:"vpc_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *route53resolverruleassociation.Timeouts `hcl:"timeouts,block"`
}
type route53ResolverRuleAssociationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_route53_resolver_rule_association.
func (rrra route53ResolverRuleAssociationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrra.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_resolver_rule_association.
func (rrra route53ResolverRuleAssociationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rrra.ref.Append("name"))
}

// ResolverRuleId returns a reference to field resolver_rule_id of aws_route53_resolver_rule_association.
func (rrra route53ResolverRuleAssociationAttributes) ResolverRuleId() terra.StringValue {
	return terra.ReferenceAsString(rrra.ref.Append("resolver_rule_id"))
}

// VpcId returns a reference to field vpc_id of aws_route53_resolver_rule_association.
func (rrra route53ResolverRuleAssociationAttributes) VpcId() terra.StringValue {
	return terra.ReferenceAsString(rrra.ref.Append("vpc_id"))
}

func (rrra route53ResolverRuleAssociationAttributes) Timeouts() route53resolverruleassociation.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53resolverruleassociation.TimeoutsAttributes](rrra.ref.Append("timeouts"))
}

type route53ResolverRuleAssociationState struct {
	Id             string                                        `json:"id"`
	Name           string                                        `json:"name"`
	ResolverRuleId string                                        `json:"resolver_rule_id"`
	VpcId          string                                        `json:"vpc_id"`
	Timeouts       *route53resolverruleassociation.TimeoutsState `json:"timeouts"`
}
