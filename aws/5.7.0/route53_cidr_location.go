// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53CidrLocation creates a new instance of [Route53CidrLocation].
func NewRoute53CidrLocation(name string, args Route53CidrLocationArgs) *Route53CidrLocation {
	return &Route53CidrLocation{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53CidrLocation)(nil)

// Route53CidrLocation represents the Terraform resource aws_route53_cidr_location.
type Route53CidrLocation struct {
	Name      string
	Args      Route53CidrLocationArgs
	state     *route53CidrLocationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53CidrLocation].
func (rcl *Route53CidrLocation) Type() string {
	return "aws_route53_cidr_location"
}

// LocalName returns the local name for [Route53CidrLocation].
func (rcl *Route53CidrLocation) LocalName() string {
	return rcl.Name
}

// Configuration returns the configuration (args) for [Route53CidrLocation].
func (rcl *Route53CidrLocation) Configuration() interface{} {
	return rcl.Args
}

// DependOn is used for other resources to depend on [Route53CidrLocation].
func (rcl *Route53CidrLocation) DependOn() terra.Reference {
	return terra.ReferenceResource(rcl)
}

// Dependencies returns the list of resources [Route53CidrLocation] depends_on.
func (rcl *Route53CidrLocation) Dependencies() terra.Dependencies {
	return rcl.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53CidrLocation].
func (rcl *Route53CidrLocation) LifecycleManagement() *terra.Lifecycle {
	return rcl.Lifecycle
}

// Attributes returns the attributes for [Route53CidrLocation].
func (rcl *Route53CidrLocation) Attributes() route53CidrLocationAttributes {
	return route53CidrLocationAttributes{ref: terra.ReferenceResource(rcl)}
}

// ImportState imports the given attribute values into [Route53CidrLocation]'s state.
func (rcl *Route53CidrLocation) ImportState(av io.Reader) error {
	rcl.state = &route53CidrLocationState{}
	if err := json.NewDecoder(av).Decode(rcl.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rcl.Type(), rcl.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53CidrLocation] has state.
func (rcl *Route53CidrLocation) State() (*route53CidrLocationState, bool) {
	return rcl.state, rcl.state != nil
}

// StateMust returns the state for [Route53CidrLocation]. Panics if the state is nil.
func (rcl *Route53CidrLocation) StateMust() *route53CidrLocationState {
	if rcl.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rcl.Type(), rcl.LocalName()))
	}
	return rcl.state
}

// Route53CidrLocationArgs contains the configurations for aws_route53_cidr_location.
type Route53CidrLocationArgs struct {
	// CidrBlocks: set of string, required
	CidrBlocks terra.SetValue[terra.StringValue] `hcl:"cidr_blocks,attr" validate:"required"`
	// CidrCollectionId: string, required
	CidrCollectionId terra.StringValue `hcl:"cidr_collection_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type route53CidrLocationAttributes struct {
	ref terra.Reference
}

// CidrBlocks returns a reference to field cidr_blocks of aws_route53_cidr_location.
func (rcl route53CidrLocationAttributes) CidrBlocks() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rcl.ref.Append("cidr_blocks"))
}

// CidrCollectionId returns a reference to field cidr_collection_id of aws_route53_cidr_location.
func (rcl route53CidrLocationAttributes) CidrCollectionId() terra.StringValue {
	return terra.ReferenceAsString(rcl.ref.Append("cidr_collection_id"))
}

// Id returns a reference to field id of aws_route53_cidr_location.
func (rcl route53CidrLocationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rcl.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53_cidr_location.
func (rcl route53CidrLocationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rcl.ref.Append("name"))
}

type route53CidrLocationState struct {
	CidrBlocks       []string `json:"cidr_blocks"`
	CidrCollectionId string   `json:"cidr_collection_id"`
	Id               string   `json:"id"`
	Name             string   `json:"name"`
}