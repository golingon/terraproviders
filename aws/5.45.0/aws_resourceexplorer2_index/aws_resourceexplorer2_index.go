// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_resourceexplorer2_index

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

// Resource represents the Terraform resource aws_resourceexplorer2_index.
type Resource struct {
	Name      string
	Args      Args
	state     *awsResourceexplorer2IndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ari *Resource) Type() string {
	return "aws_resourceexplorer2_index"
}

// LocalName returns the local name for [Resource].
func (ari *Resource) LocalName() string {
	return ari.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ari *Resource) Configuration() interface{} {
	return ari.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ari *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ari)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ari *Resource) Dependencies() terra.Dependencies {
	return ari.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ari *Resource) LifecycleManagement() *terra.Lifecycle {
	return ari.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ari *Resource) Attributes() awsResourceexplorer2IndexAttributes {
	return awsResourceexplorer2IndexAttributes{ref: terra.ReferenceResource(ari)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ari *Resource) ImportState(state io.Reader) error {
	ari.state = &awsResourceexplorer2IndexState{}
	if err := json.NewDecoder(state).Decode(ari.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ari.Type(), ari.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ari *Resource) State() (*awsResourceexplorer2IndexState, bool) {
	return ari.state, ari.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ari *Resource) StateMust() *awsResourceexplorer2IndexState {
	if ari.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ari.Type(), ari.LocalName()))
	}
	return ari.state
}

// Args contains the configurations for aws_resourceexplorer2_index.
type Args struct {
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsResourceexplorer2IndexAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_resourceexplorer2_index.
func (ari awsResourceexplorer2IndexAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ari.ref.Append("arn"))
}

// Id returns a reference to field id of aws_resourceexplorer2_index.
func (ari awsResourceexplorer2IndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ari.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_resourceexplorer2_index.
func (ari awsResourceexplorer2IndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ari.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_resourceexplorer2_index.
func (ari awsResourceexplorer2IndexAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ari.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_resourceexplorer2_index.
func (ari awsResourceexplorer2IndexAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ari.ref.Append("type"))
}

func (ari awsResourceexplorer2IndexAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ari.ref.Append("timeouts"))
}

type awsResourceexplorer2IndexState struct {
	Arn      string            `json:"arn"`
	Id       string            `json:"id"`
	Tags     map[string]string `json:"tags"`
	TagsAll  map[string]string `json:"tags_all"`
	Type     string            `json:"type"`
	Timeouts *TimeoutsState    `json:"timeouts"`
}
