// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	composerenvironment "github.com/golingon/terraproviders/googlebeta/4.64.0/composerenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComposerEnvironment creates a new instance of [ComposerEnvironment].
func NewComposerEnvironment(name string, args ComposerEnvironmentArgs) *ComposerEnvironment {
	return &ComposerEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComposerEnvironment)(nil)

// ComposerEnvironment represents the Terraform resource google_composer_environment.
type ComposerEnvironment struct {
	Name      string
	Args      ComposerEnvironmentArgs
	state     *composerEnvironmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComposerEnvironment].
func (ce *ComposerEnvironment) Type() string {
	return "google_composer_environment"
}

// LocalName returns the local name for [ComposerEnvironment].
func (ce *ComposerEnvironment) LocalName() string {
	return ce.Name
}

// Configuration returns the configuration (args) for [ComposerEnvironment].
func (ce *ComposerEnvironment) Configuration() interface{} {
	return ce.Args
}

// DependOn is used for other resources to depend on [ComposerEnvironment].
func (ce *ComposerEnvironment) DependOn() terra.Reference {
	return terra.ReferenceResource(ce)
}

// Dependencies returns the list of resources [ComposerEnvironment] depends_on.
func (ce *ComposerEnvironment) Dependencies() terra.Dependencies {
	return ce.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComposerEnvironment].
func (ce *ComposerEnvironment) LifecycleManagement() *terra.Lifecycle {
	return ce.Lifecycle
}

// Attributes returns the attributes for [ComposerEnvironment].
func (ce *ComposerEnvironment) Attributes() composerEnvironmentAttributes {
	return composerEnvironmentAttributes{ref: terra.ReferenceResource(ce)}
}

// ImportState imports the given attribute values into [ComposerEnvironment]'s state.
func (ce *ComposerEnvironment) ImportState(av io.Reader) error {
	ce.state = &composerEnvironmentState{}
	if err := json.NewDecoder(av).Decode(ce.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ce.Type(), ce.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComposerEnvironment] has state.
func (ce *ComposerEnvironment) State() (*composerEnvironmentState, bool) {
	return ce.state, ce.state != nil
}

// StateMust returns the state for [ComposerEnvironment]. Panics if the state is nil.
func (ce *ComposerEnvironment) StateMust() *composerEnvironmentState {
	if ce.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ce.Type(), ce.LocalName()))
	}
	return ce.state
}

// ComposerEnvironmentArgs contains the configurations for google_composer_environment.
type ComposerEnvironmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Config: optional
	Config *composerenvironment.Config `hcl:"config,block"`
	// Timeouts: optional
	Timeouts *composerenvironment.Timeouts `hcl:"timeouts,block"`
}
type composerEnvironmentAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_composer_environment.
func (ce composerEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("id"))
}

// Labels returns a reference to field labels of google_composer_environment.
func (ce composerEnvironmentAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ce.ref.Append("labels"))
}

// Name returns a reference to field name of google_composer_environment.
func (ce composerEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("name"))
}

// Project returns a reference to field project of google_composer_environment.
func (ce composerEnvironmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("project"))
}

// Region returns a reference to field region of google_composer_environment.
func (ce composerEnvironmentAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(ce.ref.Append("region"))
}

func (ce composerEnvironmentAttributes) Config() terra.ListValue[composerenvironment.ConfigAttributes] {
	return terra.ReferenceAsList[composerenvironment.ConfigAttributes](ce.ref.Append("config"))
}

func (ce composerEnvironmentAttributes) Timeouts() composerenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[composerenvironment.TimeoutsAttributes](ce.ref.Append("timeouts"))
}

type composerEnvironmentState struct {
	Id       string                             `json:"id"`
	Labels   map[string]string                  `json:"labels"`
	Name     string                             `json:"name"`
	Project  string                             `json:"project"`
	Region   string                             `json:"region"`
	Config   []composerenvironment.ConfigState  `json:"config"`
	Timeouts *composerenvironment.TimeoutsState `json:"timeouts"`
}
