// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewRoute53CidrCollection creates a new instance of [Route53CidrCollection].
func NewRoute53CidrCollection(name string, args Route53CidrCollectionArgs) *Route53CidrCollection {
	return &Route53CidrCollection{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53CidrCollection)(nil)

// Route53CidrCollection represents the Terraform resource aws_route53_cidr_collection.
type Route53CidrCollection struct {
	Name      string
	Args      Route53CidrCollectionArgs
	state     *route53CidrCollectionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53CidrCollection].
func (rcc *Route53CidrCollection) Type() string {
	return "aws_route53_cidr_collection"
}

// LocalName returns the local name for [Route53CidrCollection].
func (rcc *Route53CidrCollection) LocalName() string {
	return rcc.Name
}

// Configuration returns the configuration (args) for [Route53CidrCollection].
func (rcc *Route53CidrCollection) Configuration() interface{} {
	return rcc.Args
}

// DependOn is used for other resources to depend on [Route53CidrCollection].
func (rcc *Route53CidrCollection) DependOn() terra.Reference {
	return terra.ReferenceResource(rcc)
}

// Dependencies returns the list of resources [Route53CidrCollection] depends_on.
func (rcc *Route53CidrCollection) Dependencies() terra.Dependencies {
	return rcc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53CidrCollection].
func (rcc *Route53CidrCollection) LifecycleManagement() *terra.Lifecycle {
	return rcc.Lifecycle
}

// Attributes returns the attributes for [Route53CidrCollection].
func (rcc *Route53CidrCollection) Attributes() route53CidrCollectionAttributes {
	return route53CidrCollectionAttributes{ref: terra.ReferenceResource(rcc)}
}

// ImportState imports the given attribute values into [Route53CidrCollection]'s state.
func (rcc *Route53CidrCollection) ImportState(av io.Reader) error {
	rcc.state = &route53CidrCollectionState{}
	if err := json.NewDecoder(av).Decode(rcc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcc.Type(), rcc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53CidrCollection] has state.
func (rcc *Route53CidrCollection) State() (*route53CidrCollectionState, bool) {
	return rcc.state, rcc.state != nil
}

// StateMust returns the state for [Route53CidrCollection]. Panics if the state is nil.
func (rcc *Route53CidrCollection) StateMust() *route53CidrCollectionState {
	if rcc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcc.Type(), rcc.LocalName()))
	}
	return rcc.state
}

// Route53CidrCollectionArgs contains the configurations for aws_route53_cidr_collection.
type Route53CidrCollectionArgs struct {
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type route53CidrCollectionAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53_cidr_collection.
func (rcc route53CidrCollectionAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rcc.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53_cidr_collection.
func (rcc route53CidrCollectionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcc.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_cidr_collection.
func (rcc route53CidrCollectionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rcc.ref.Append("name"))
}

// Version returns a reference to field version of aws_route53_cidr_collection.
func (rcc route53CidrCollectionAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(rcc.ref.Append("version"))
}

type route53CidrCollectionState struct {
	Arn     string  `json:"arn"`
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Version float64 `json:"version"`
}
