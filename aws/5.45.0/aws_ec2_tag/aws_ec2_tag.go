// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_ec2_tag

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

// Resource represents the Terraform resource aws_ec2_tag.
type Resource struct {
	Name      string
	Args      Args
	state     *awsEc2TagState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (aet *Resource) Type() string {
	return "aws_ec2_tag"
}

// LocalName returns the local name for [Resource].
func (aet *Resource) LocalName() string {
	return aet.Name
}

// Configuration returns the configuration (args) for [Resource].
func (aet *Resource) Configuration() interface{} {
	return aet.Args
}

// DependOn is used for other resources to depend on [Resource].
func (aet *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(aet)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (aet *Resource) Dependencies() terra.Dependencies {
	return aet.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (aet *Resource) LifecycleManagement() *terra.Lifecycle {
	return aet.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (aet *Resource) Attributes() awsEc2TagAttributes {
	return awsEc2TagAttributes{ref: terra.ReferenceResource(aet)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (aet *Resource) ImportState(state io.Reader) error {
	aet.state = &awsEc2TagState{}
	if err := json.NewDecoder(state).Decode(aet.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", aet.Type(), aet.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (aet *Resource) State() (*awsEc2TagState, bool) {
	return aet.state, aet.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (aet *Resource) StateMust() *awsEc2TagState {
	if aet.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", aet.Type(), aet.LocalName()))
	}
	return aet.state
}

// Args contains the configurations for aws_ec2_tag.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// ResourceId: string, required
	ResourceId terra.StringValue `hcl:"resource_id,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type awsEc2TagAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_ec2_tag.
func (aet awsEc2TagAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aet.ref.Append("id"))
}

// Key returns a reference to field key of aws_ec2_tag.
func (aet awsEc2TagAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(aet.ref.Append("key"))
}

// ResourceId returns a reference to field resource_id of aws_ec2_tag.
func (aet awsEc2TagAttributes) ResourceId() terra.StringValue {
	return terra.ReferenceAsString(aet.ref.Append("resource_id"))
}

// Value returns a reference to field value of aws_ec2_tag.
func (aet awsEc2TagAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(aet.ref.Append("value"))
}

type awsEc2TagState struct {
	Id         string `json:"id"`
	Key        string `json:"key"`
	ResourceId string `json:"resource_id"`
	Value      string `json:"value"`
}
