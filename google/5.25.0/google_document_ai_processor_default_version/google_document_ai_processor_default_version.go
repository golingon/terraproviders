// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_document_ai_processor_default_version

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

// Resource represents the Terraform resource google_document_ai_processor_default_version.
type Resource struct {
	Name      string
	Args      Args
	state     *googleDocumentAiProcessorDefaultVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gdapdv *Resource) Type() string {
	return "google_document_ai_processor_default_version"
}

// LocalName returns the local name for [Resource].
func (gdapdv *Resource) LocalName() string {
	return gdapdv.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gdapdv *Resource) Configuration() interface{} {
	return gdapdv.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gdapdv *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gdapdv)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gdapdv *Resource) Dependencies() terra.Dependencies {
	return gdapdv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gdapdv *Resource) LifecycleManagement() *terra.Lifecycle {
	return gdapdv.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gdapdv *Resource) Attributes() googleDocumentAiProcessorDefaultVersionAttributes {
	return googleDocumentAiProcessorDefaultVersionAttributes{ref: terra.ReferenceResource(gdapdv)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gdapdv *Resource) ImportState(state io.Reader) error {
	gdapdv.state = &googleDocumentAiProcessorDefaultVersionState{}
	if err := json.NewDecoder(state).Decode(gdapdv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gdapdv.Type(), gdapdv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gdapdv *Resource) State() (*googleDocumentAiProcessorDefaultVersionState, bool) {
	return gdapdv.state, gdapdv.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gdapdv *Resource) StateMust() *googleDocumentAiProcessorDefaultVersionState {
	if gdapdv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gdapdv.Type(), gdapdv.LocalName()))
	}
	return gdapdv.state
}

// Args contains the configurations for google_document_ai_processor_default_version.
type Args struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Processor: string, required
	Processor terra.StringValue `hcl:"processor,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type googleDocumentAiProcessorDefaultVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_document_ai_processor_default_version.
func (gdapdv googleDocumentAiProcessorDefaultVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gdapdv.ref.Append("id"))
}

// Processor returns a reference to field processor of google_document_ai_processor_default_version.
func (gdapdv googleDocumentAiProcessorDefaultVersionAttributes) Processor() terra.StringValue {
	return terra.ReferenceAsString(gdapdv.ref.Append("processor"))
}

// Version returns a reference to field version of google_document_ai_processor_default_version.
func (gdapdv googleDocumentAiProcessorDefaultVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(gdapdv.ref.Append("version"))
}

func (gdapdv googleDocumentAiProcessorDefaultVersionAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gdapdv.ref.Append("timeouts"))
}

type googleDocumentAiProcessorDefaultVersionState struct {
	Id        string         `json:"id"`
	Processor string         `json:"processor"`
	Version   string         `json:"version"`
	Timeouts  *TimeoutsState `json:"timeouts"`
}
