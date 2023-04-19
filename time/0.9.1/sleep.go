// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package time

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSleep creates a new instance of [Sleep].
func NewSleep(name string, args SleepArgs) *Sleep {
	return &Sleep{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Sleep)(nil)

// Sleep represents the Terraform resource time_sleep.
type Sleep struct {
	Name      string
	Args      SleepArgs
	state     *sleepState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Sleep].
func (s *Sleep) Type() string {
	return "time_sleep"
}

// LocalName returns the local name for [Sleep].
func (s *Sleep) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [Sleep].
func (s *Sleep) Configuration() interface{} {
	return s.Args
}

// DependOn is used for other resources to depend on [Sleep].
func (s *Sleep) DependOn() terra.Reference {
	return terra.ReferenceResource(s)
}

// Dependencies returns the list of resources [Sleep] depends_on.
func (s *Sleep) Dependencies() terra.Dependencies {
	return s.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Sleep].
func (s *Sleep) LifecycleManagement() *terra.Lifecycle {
	return s.Lifecycle
}

// Attributes returns the attributes for [Sleep].
func (s *Sleep) Attributes() sleepAttributes {
	return sleepAttributes{ref: terra.ReferenceResource(s)}
}

// ImportState imports the given attribute values into [Sleep]'s state.
func (s *Sleep) ImportState(av io.Reader) error {
	s.state = &sleepState{}
	if err := json.NewDecoder(av).Decode(s.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", s.Type(), s.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Sleep] has state.
func (s *Sleep) State() (*sleepState, bool) {
	return s.state, s.state != nil
}

// StateMust returns the state for [Sleep]. Panics if the state is nil.
func (s *Sleep) StateMust() *sleepState {
	if s.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", s.Type(), s.LocalName()))
	}
	return s.state
}

// SleepArgs contains the configurations for time_sleep.
type SleepArgs struct {
	// CreateDuration: string, optional
	CreateDuration terra.StringValue `hcl:"create_duration,attr"`
	// DestroyDuration: string, optional
	DestroyDuration terra.StringValue `hcl:"destroy_duration,attr"`
	// Triggers: map of string, optional
	Triggers terra.MapValue[terra.StringValue] `hcl:"triggers,attr"`
}
type sleepAttributes struct {
	ref terra.Reference
}

// CreateDuration returns a reference to field create_duration of time_sleep.
func (s sleepAttributes) CreateDuration() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("create_duration"))
}

// DestroyDuration returns a reference to field destroy_duration of time_sleep.
func (s sleepAttributes) DestroyDuration() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("destroy_duration"))
}

// Id returns a reference to field id of time_sleep.
func (s sleepAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Triggers returns a reference to field triggers of time_sleep.
func (s sleepAttributes) Triggers() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](s.ref.Append("triggers"))
}

type sleepState struct {
	CreateDuration  string            `json:"create_duration"`
	DestroyDuration string            `json:"destroy_duration"`
	Id              string            `json:"id"`
	Triggers        map[string]string `json:"triggers"`
}
