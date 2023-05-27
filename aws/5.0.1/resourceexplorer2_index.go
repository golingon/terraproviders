// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	resourceexplorer2index "github.com/golingon/terraproviders/aws/5.0.1/resourceexplorer2index"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceexplorer2Index creates a new instance of [Resourceexplorer2Index].
func NewResourceexplorer2Index(name string, args Resourceexplorer2IndexArgs) *Resourceexplorer2Index {
	return &Resourceexplorer2Index{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resourceexplorer2Index)(nil)

// Resourceexplorer2Index represents the Terraform resource aws_resourceexplorer2_index.
type Resourceexplorer2Index struct {
	Name      string
	Args      Resourceexplorer2IndexArgs
	state     *resourceexplorer2IndexState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) Type() string {
	return "aws_resourceexplorer2_index"
}

// LocalName returns the local name for [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) LocalName() string {
	return ri.Name
}

// Configuration returns the configuration (args) for [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) Configuration() interface{} {
	return ri.Args
}

// DependOn is used for other resources to depend on [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) DependOn() terra.Reference {
	return terra.ReferenceResource(ri)
}

// Dependencies returns the list of resources [Resourceexplorer2Index] depends_on.
func (ri *Resourceexplorer2Index) Dependencies() terra.Dependencies {
	return ri.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) LifecycleManagement() *terra.Lifecycle {
	return ri.Lifecycle
}

// Attributes returns the attributes for [Resourceexplorer2Index].
func (ri *Resourceexplorer2Index) Attributes() resourceexplorer2IndexAttributes {
	return resourceexplorer2IndexAttributes{ref: terra.ReferenceResource(ri)}
}

// ImportState imports the given attribute values into [Resourceexplorer2Index]'s state.
func (ri *Resourceexplorer2Index) ImportState(av io.Reader) error {
	ri.state = &resourceexplorer2IndexState{}
	if err := json.NewDecoder(av).Decode(ri.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ri.Type(), ri.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resourceexplorer2Index] has state.
func (ri *Resourceexplorer2Index) State() (*resourceexplorer2IndexState, bool) {
	return ri.state, ri.state != nil
}

// StateMust returns the state for [Resourceexplorer2Index]. Panics if the state is nil.
func (ri *Resourceexplorer2Index) StateMust() *resourceexplorer2IndexState {
	if ri.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ri.Type(), ri.LocalName()))
	}
	return ri.state
}

// Resourceexplorer2IndexArgs contains the configurations for aws_resourceexplorer2_index.
type Resourceexplorer2IndexArgs struct {
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *resourceexplorer2index.Timeouts `hcl:"timeouts,block"`
}
type resourceexplorer2IndexAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_resourceexplorer2_index.
func (ri resourceexplorer2IndexAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("arn"))
}

// Id returns a reference to field id of aws_resourceexplorer2_index.
func (ri resourceexplorer2IndexAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("id"))
}

// Tags returns a reference to field tags of aws_resourceexplorer2_index.
func (ri resourceexplorer2IndexAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_resourceexplorer2_index.
func (ri resourceexplorer2IndexAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ri.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_resourceexplorer2_index.
func (ri resourceexplorer2IndexAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ri.ref.Append("type"))
}

func (ri resourceexplorer2IndexAttributes) Timeouts() resourceexplorer2index.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourceexplorer2index.TimeoutsAttributes](ri.ref.Append("timeouts"))
}

type resourceexplorer2IndexState struct {
	Arn      string                                `json:"arn"`
	Id       string                                `json:"id"`
	Tags     map[string]string                     `json:"tags"`
	TagsAll  map[string]string                     `json:"tags_all"`
	Type     string                                `json:"type"`
	Timeouts *resourceexplorer2index.TimeoutsState `json:"timeouts"`
}
