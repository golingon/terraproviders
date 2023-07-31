// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	route53recoveryreadinessresourceset "github.com/golingon/terraproviders/aws/5.10.0/route53recoveryreadinessresourceset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRoute53RecoveryreadinessResourceSet creates a new instance of [Route53RecoveryreadinessResourceSet].
func NewRoute53RecoveryreadinessResourceSet(name string, args Route53RecoveryreadinessResourceSetArgs) *Route53RecoveryreadinessResourceSet {
	return &Route53RecoveryreadinessResourceSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Route53RecoveryreadinessResourceSet)(nil)

// Route53RecoveryreadinessResourceSet represents the Terraform resource aws_route53recoveryreadiness_resource_set.
type Route53RecoveryreadinessResourceSet struct {
	Name      string
	Args      Route53RecoveryreadinessResourceSetArgs
	state     *route53RecoveryreadinessResourceSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) Type() string {
	return "aws_route53recoveryreadiness_resource_set"
}

// LocalName returns the local name for [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) LocalName() string {
	return rrs.Name
}

// Configuration returns the configuration (args) for [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) Configuration() interface{} {
	return rrs.Args
}

// DependOn is used for other resources to depend on [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) DependOn() terra.Reference {
	return terra.ReferenceResource(rrs)
}

// Dependencies returns the list of resources [Route53RecoveryreadinessResourceSet] depends_on.
func (rrs *Route53RecoveryreadinessResourceSet) Dependencies() terra.Dependencies {
	return rrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) LifecycleManagement() *terra.Lifecycle {
	return rrs.Lifecycle
}

// Attributes returns the attributes for [Route53RecoveryreadinessResourceSet].
func (rrs *Route53RecoveryreadinessResourceSet) Attributes() route53RecoveryreadinessResourceSetAttributes {
	return route53RecoveryreadinessResourceSetAttributes{ref: terra.ReferenceResource(rrs)}
}

// ImportState imports the given attribute values into [Route53RecoveryreadinessResourceSet]'s state.
func (rrs *Route53RecoveryreadinessResourceSet) ImportState(av io.Reader) error {
	rrs.state = &route53RecoveryreadinessResourceSetState{}
	if err := json.NewDecoder(av).Decode(rrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rrs.Type(), rrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Route53RecoveryreadinessResourceSet] has state.
func (rrs *Route53RecoveryreadinessResourceSet) State() (*route53RecoveryreadinessResourceSetState, bool) {
	return rrs.state, rrs.state != nil
}

// StateMust returns the state for [Route53RecoveryreadinessResourceSet]. Panics if the state is nil.
func (rrs *Route53RecoveryreadinessResourceSet) StateMust() *route53RecoveryreadinessResourceSetState {
	if rrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rrs.Type(), rrs.LocalName()))
	}
	return rrs.state
}

// Route53RecoveryreadinessResourceSetArgs contains the configurations for aws_route53recoveryreadiness_resource_set.
type Route53RecoveryreadinessResourceSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ResourceSetName: string, required
	ResourceSetName terra.StringValue `hcl:"resource_set_name,attr" validate:"required"`
	// ResourceSetType: string, required
	ResourceSetType terra.StringValue `hcl:"resource_set_type,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Resources: min=1
	Resources []route53recoveryreadinessresourceset.Resources `hcl:"resources,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *route53recoveryreadinessresourceset.Timeouts `hcl:"timeouts,block"`
}
type route53RecoveryreadinessResourceSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("id"))
}

// ResourceSetName returns a reference to field resource_set_name of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) ResourceSetName() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("resource_set_name"))
}

// ResourceSetType returns a reference to field resource_set_type of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) ResourceSetType() terra.StringValue {
	return terra.ReferenceAsString(rrs.ref.Append("resource_set_type"))
}

// Tags returns a reference to field tags of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53recoveryreadiness_resource_set.
func (rrs route53RecoveryreadinessResourceSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rrs.ref.Append("tags_all"))
}

func (rrs route53RecoveryreadinessResourceSetAttributes) Resources() terra.ListValue[route53recoveryreadinessresourceset.ResourcesAttributes] {
	return terra.ReferenceAsList[route53recoveryreadinessresourceset.ResourcesAttributes](rrs.ref.Append("resources"))
}

func (rrs route53RecoveryreadinessResourceSetAttributes) Timeouts() route53recoveryreadinessresourceset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[route53recoveryreadinessresourceset.TimeoutsAttributes](rrs.ref.Append("timeouts"))
}

type route53RecoveryreadinessResourceSetState struct {
	Arn             string                                               `json:"arn"`
	Id              string                                               `json:"id"`
	ResourceSetName string                                               `json:"resource_set_name"`
	ResourceSetType string                                               `json:"resource_set_type"`
	Tags            map[string]string                                    `json:"tags"`
	TagsAll         map[string]string                                    `json:"tags_all"`
	Resources       []route53recoveryreadinessresourceset.ResourcesState `json:"resources"`
	Timeouts        *route53recoveryreadinessresourceset.TimeoutsState   `json:"timeouts"`
}
