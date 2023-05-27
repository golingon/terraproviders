// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudfunctions2function "github.com/golingon/terraproviders/google/4.66.0/cloudfunctions2function"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudfunctions2Function creates a new instance of [Cloudfunctions2Function].
func NewCloudfunctions2Function(name string, args Cloudfunctions2FunctionArgs) *Cloudfunctions2Function {
	return &Cloudfunctions2Function{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Cloudfunctions2Function)(nil)

// Cloudfunctions2Function represents the Terraform resource google_cloudfunctions2_function.
type Cloudfunctions2Function struct {
	Name      string
	Args      Cloudfunctions2FunctionArgs
	state     *cloudfunctions2FunctionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) Type() string {
	return "google_cloudfunctions2_function"
}

// LocalName returns the local name for [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) LocalName() string {
	return cf.Name
}

// Configuration returns the configuration (args) for [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) Configuration() interface{} {
	return cf.Args
}

// DependOn is used for other resources to depend on [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) DependOn() terra.Reference {
	return terra.ReferenceResource(cf)
}

// Dependencies returns the list of resources [Cloudfunctions2Function] depends_on.
func (cf *Cloudfunctions2Function) Dependencies() terra.Dependencies {
	return cf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) LifecycleManagement() *terra.Lifecycle {
	return cf.Lifecycle
}

// Attributes returns the attributes for [Cloudfunctions2Function].
func (cf *Cloudfunctions2Function) Attributes() cloudfunctions2FunctionAttributes {
	return cloudfunctions2FunctionAttributes{ref: terra.ReferenceResource(cf)}
}

// ImportState imports the given attribute values into [Cloudfunctions2Function]'s state.
func (cf *Cloudfunctions2Function) ImportState(av io.Reader) error {
	cf.state = &cloudfunctions2FunctionState{}
	if err := json.NewDecoder(av).Decode(cf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cf.Type(), cf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Cloudfunctions2Function] has state.
func (cf *Cloudfunctions2Function) State() (*cloudfunctions2FunctionState, bool) {
	return cf.state, cf.state != nil
}

// StateMust returns the state for [Cloudfunctions2Function]. Panics if the state is nil.
func (cf *Cloudfunctions2Function) StateMust() *cloudfunctions2FunctionState {
	if cf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cf.Type(), cf.LocalName()))
	}
	return cf.state
}

// Cloudfunctions2FunctionArgs contains the configurations for google_cloudfunctions2_function.
type Cloudfunctions2FunctionArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// BuildConfig: optional
	BuildConfig *cloudfunctions2function.BuildConfig `hcl:"build_config,block"`
	// EventTrigger: optional
	EventTrigger *cloudfunctions2function.EventTrigger `hcl:"event_trigger,block"`
	// ServiceConfig: optional
	ServiceConfig *cloudfunctions2function.ServiceConfig `hcl:"service_config,block"`
	// Timeouts: optional
	Timeouts *cloudfunctions2function.Timeouts `hcl:"timeouts,block"`
}
type cloudfunctions2FunctionAttributes struct {
	ref terra.Reference
}

// Description returns a reference to field description of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("description"))
}

// Environment returns a reference to field environment of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Environment() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("environment"))
}

// Id returns a reference to field id of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("id"))
}

// Labels returns a reference to field labels of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cf.ref.Append("labels"))
}

// Location returns a reference to field location of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("project"))
}

// State returns a reference to field state of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("state"))
}

// UpdateTime returns a reference to field update_time of google_cloudfunctions2_function.
func (cf cloudfunctions2FunctionAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cf.ref.Append("update_time"))
}

func (cf cloudfunctions2FunctionAttributes) BuildConfig() terra.ListValue[cloudfunctions2function.BuildConfigAttributes] {
	return terra.ReferenceAsList[cloudfunctions2function.BuildConfigAttributes](cf.ref.Append("build_config"))
}

func (cf cloudfunctions2FunctionAttributes) EventTrigger() terra.ListValue[cloudfunctions2function.EventTriggerAttributes] {
	return terra.ReferenceAsList[cloudfunctions2function.EventTriggerAttributes](cf.ref.Append("event_trigger"))
}

func (cf cloudfunctions2FunctionAttributes) ServiceConfig() terra.ListValue[cloudfunctions2function.ServiceConfigAttributes] {
	return terra.ReferenceAsList[cloudfunctions2function.ServiceConfigAttributes](cf.ref.Append("service_config"))
}

func (cf cloudfunctions2FunctionAttributes) Timeouts() cloudfunctions2function.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudfunctions2function.TimeoutsAttributes](cf.ref.Append("timeouts"))
}

type cloudfunctions2FunctionState struct {
	Description   string                                       `json:"description"`
	Environment   string                                       `json:"environment"`
	Id            string                                       `json:"id"`
	Labels        map[string]string                            `json:"labels"`
	Location      string                                       `json:"location"`
	Name          string                                       `json:"name"`
	Project       string                                       `json:"project"`
	State         string                                       `json:"state"`
	UpdateTime    string                                       `json:"update_time"`
	BuildConfig   []cloudfunctions2function.BuildConfigState   `json:"build_config"`
	EventTrigger  []cloudfunctions2function.EventTriggerState  `json:"event_trigger"`
	ServiceConfig []cloudfunctions2function.ServiceConfigState `json:"service_config"`
	Timeouts      *cloudfunctions2function.TimeoutsState       `json:"timeouts"`
}
