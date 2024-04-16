// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53recoveryreadiness_resource_set

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_route53recoveryreadiness_resource_set.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRoute53RecoveryreadinessResourceSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arrs *Resource) Type() string {
	return "aws_route53recoveryreadiness_resource_set"
}

// LocalName returns the local name for [Resource].
func (arrs *Resource) LocalName() string {
	return arrs.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arrs *Resource) Configuration() interface{} {
	return arrs.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arrs *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arrs)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arrs *Resource) Dependencies() terra.Dependencies {
	return arrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arrs *Resource) LifecycleManagement() *terra.Lifecycle {
	return arrs.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arrs *Resource) Attributes() awsRoute53RecoveryreadinessResourceSetAttributes {
	return awsRoute53RecoveryreadinessResourceSetAttributes{ref: terra.ReferenceResource(arrs)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arrs *Resource) ImportState(state io.Reader) error {
	arrs.state = &awsRoute53RecoveryreadinessResourceSetState{}
	if err := json.NewDecoder(state).Decode(arrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arrs.Type(), arrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arrs *Resource) State() (*awsRoute53RecoveryreadinessResourceSetState, bool) {
	return arrs.state, arrs.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arrs *Resource) StateMust() *awsRoute53RecoveryreadinessResourceSetState {
	if arrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arrs.Type(), arrs.LocalName()))
	}
	return arrs.state
}

// Args contains the configurations for aws_route53recoveryreadiness_resource_set.
type Args struct {
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
	Resources []Resources `hcl:"resources,block" validate:"min=1"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsRoute53RecoveryreadinessResourceSetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arrs.ref.Append("arn"))
}

// Id returns a reference to field id of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrs.ref.Append("id"))
}

// ResourceSetName returns a reference to field resource_set_name of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) ResourceSetName() terra.StringValue {
	return terra.ReferenceAsString(arrs.ref.Append("resource_set_name"))
}

// ResourceSetType returns a reference to field resource_set_type of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) ResourceSetType() terra.StringValue {
	return terra.ReferenceAsString(arrs.ref.Append("resource_set_type"))
}

// Tags returns a reference to field tags of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arrs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_route53recoveryreadiness_resource_set.
func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arrs.ref.Append("tags_all"))
}

func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) Resources() terra.ListValue[ResourcesAttributes] {
	return terra.ReferenceAsList[ResourcesAttributes](arrs.ref.Append("resources"))
}

func (arrs awsRoute53RecoveryreadinessResourceSetAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](arrs.ref.Append("timeouts"))
}

type awsRoute53RecoveryreadinessResourceSetState struct {
	Arn             string            `json:"arn"`
	Id              string            `json:"id"`
	ResourceSetName string            `json:"resource_set_name"`
	ResourceSetType string            `json:"resource_set_type"`
	Tags            map[string]string `json:"tags"`
	TagsAll         map[string]string `json:"tags_all"`
	Resources       []ResourcesState  `json:"resources"`
	Timeouts        *TimeoutsState    `json:"timeouts"`
}
