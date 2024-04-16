// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package null_resource

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

// Resource represents the Terraform resource null_resource.
type Resource struct {
	Name      string
	Args      Args
	state     *nullResourceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (nr *Resource) Type() string {
	return "null_resource"
}

// LocalName returns the local name for [Resource].
func (nr *Resource) LocalName() string {
	return nr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (nr *Resource) Configuration() interface{} {
	return nr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (nr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(nr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (nr *Resource) Dependencies() terra.Dependencies {
	return nr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (nr *Resource) LifecycleManagement() *terra.Lifecycle {
	return nr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (nr *Resource) Attributes() nullResourceAttributes {
	return nullResourceAttributes{ref: terra.ReferenceResource(nr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (nr *Resource) ImportState(state io.Reader) error {
	nr.state = &nullResourceState{}
	if err := json.NewDecoder(state).Decode(nr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", nr.Type(), nr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (nr *Resource) State() (*nullResourceState, bool) {
	return nr.state, nr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (nr *Resource) StateMust() *nullResourceState {
	if nr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", nr.Type(), nr.LocalName()))
	}
	return nr.state
}

// Args contains the configurations for null_resource.
type Args struct {
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
}

type nullResourceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of null_resource.
func (nr nullResourceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nr.ref.Append("id"))
}

// Triggers returns a reference to field triggers of null_resource.
func (nr nullResourceAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](nr.ref.Append("triggers"))
}

type nullResourceState struct {
	Id       string            `json:"id"`
	Triggers map[string]string `json:"triggers"`
}
