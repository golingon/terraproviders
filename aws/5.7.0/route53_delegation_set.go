// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53DelegationSet creates a new instance of [Route53DelegationSet].
func NewRoute53DelegationSet(name string, args Route53DelegationSetArgs) *Route53DelegationSet {
	return &Route53DelegationSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53DelegationSet)(nil)

// Route53DelegationSet represents the Terraform resource aws_route53_delegation_set.
type Route53DelegationSet struct {
	Name      string
	Args      Route53DelegationSetArgs
	state     *route53DelegationSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53DelegationSet].
func (rds *Route53DelegationSet) Type() string {
	return "aws_route53_delegation_set"
}

// LocalName returns the local name for [Route53DelegationSet].
func (rds *Route53DelegationSet) LocalName() string {
	return rds.Name
}

// Configuration returns the configuration (args) for [Route53DelegationSet].
func (rds *Route53DelegationSet) Configuration() interface{} {
	return rds.Args
}

// DependOn is used for other resources to depend on [Route53DelegationSet].
func (rds *Route53DelegationSet) DependOn() terra.Reference {
	return terra.ReferenceResource(rds)
}

// Dependencies returns the list of resources [Route53DelegationSet] depends_on.
func (rds *Route53DelegationSet) Dependencies() terra.Dependencies {
	return rds.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53DelegationSet].
func (rds *Route53DelegationSet) LifecycleManagement() *terra.Lifecycle {
	return rds.Lifecycle
}

// Attributes returns the attributes for [Route53DelegationSet].
func (rds *Route53DelegationSet) Attributes() route53DelegationSetAttributes {
	return route53DelegationSetAttributes{ref: terra.ReferenceResource(rds)}
}

// ImportState imports the given attribute values into [Route53DelegationSet]'s state.
func (rds *Route53DelegationSet) ImportState(av io.Reader) error {
	rds.state = &route53DelegationSetState{}
	if err := json.NewDecoder(av).Decode(rds.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rds.Type(), rds.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53DelegationSet] has state.
func (rds *Route53DelegationSet) State() (*route53DelegationSetState, bool) {
	return rds.state, rds.state != nil
}

// StateMust returns the state for [Route53DelegationSet]. Panics if the state is nil.
func (rds *Route53DelegationSet) StateMust() *route53DelegationSetState {
	if rds.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rds.Type(), rds.LocalName()))
	}
	return rds.state
}

// Route53DelegationSetArgs contains the configurations for aws_route53_delegation_set.
type Route53DelegationSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ReferenceName: string, optional
	ReferenceName terra.StringValue `hcl:"reference_name,attr"`
}
type route53DelegationSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_delegation_set.
func (rds route53DelegationSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rds.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53_delegation_set.
func (rds route53DelegationSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rds.ref.Append("id"))
}

// NameServers returns a reference to field name_servers of aws_route53_delegation_set.
func (rds route53DelegationSetAttributes) NameServers() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](rds.ref.Append("name_servers"))
}

// ReferenceName returns a reference to field reference_name of aws_route53_delegation_set.
func (rds route53DelegationSetAttributes) ReferenceName() terra.StringValue {
	return terra.ReferenceAsString(rds.ref.Append("reference_name"))
}

type route53DelegationSetState struct {
	Arn           string   `json:"arn"`
	Id            string   `json:"id"`
	NameServers   []string `json:"name_servers"`
	ReferenceName string   `json:"reference_name"`
}
