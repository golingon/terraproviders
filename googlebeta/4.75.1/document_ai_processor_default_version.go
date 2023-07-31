// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	documentaiprocessordefaultversion "github.com/golingon/terraproviders/googlebeta/4.75.1/documentaiprocessordefaultversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocumentAiProcessorDefaultVersion creates a new instance of [DocumentAiProcessorDefaultVersion].
func NewDocumentAiProcessorDefaultVersion(name string, args DocumentAiProcessorDefaultVersionArgs) *DocumentAiProcessorDefaultVersion {
	return &DocumentAiProcessorDefaultVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocumentAiProcessorDefaultVersion)(nil)

// DocumentAiProcessorDefaultVersion represents the Terraform resource google_document_ai_processor_default_version.
type DocumentAiProcessorDefaultVersion struct {
	Name      string
	Args      DocumentAiProcessorDefaultVersionArgs
	state     *documentAiProcessorDefaultVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) Type() string {
	return "google_document_ai_processor_default_version"
}

// LocalName returns the local name for [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) LocalName() string {
	return dapdv.Name
}

// Configuration returns the configuration (args) for [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) Configuration() interface{} {
	return dapdv.Args
}

// DependOn is used for other resources to depend on [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(dapdv)
}

// Dependencies returns the list of resources [DocumentAiProcessorDefaultVersion] depends_on.
func (dapdv *DocumentAiProcessorDefaultVersion) Dependencies() terra.Dependencies {
	return dapdv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) LifecycleManagement() *terra.Lifecycle {
	return dapdv.Lifecycle
}

// Attributes returns the attributes for [DocumentAiProcessorDefaultVersion].
func (dapdv *DocumentAiProcessorDefaultVersion) Attributes() documentAiProcessorDefaultVersionAttributes {
	return documentAiProcessorDefaultVersionAttributes{ref: terra.ReferenceResource(dapdv)}
}

// ImportState imports the given attribute values into [DocumentAiProcessorDefaultVersion]'s state.
func (dapdv *DocumentAiProcessorDefaultVersion) ImportState(av io.Reader) error {
	dapdv.state = &documentAiProcessorDefaultVersionState{}
	if err := json.NewDecoder(av).Decode(dapdv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dapdv.Type(), dapdv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocumentAiProcessorDefaultVersion] has state.
func (dapdv *DocumentAiProcessorDefaultVersion) State() (*documentAiProcessorDefaultVersionState, bool) {
	return dapdv.state, dapdv.state != nil
}

// StateMust returns the state for [DocumentAiProcessorDefaultVersion]. Panics if the state is nil.
func (dapdv *DocumentAiProcessorDefaultVersion) StateMust() *documentAiProcessorDefaultVersionState {
	if dapdv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dapdv.Type(), dapdv.LocalName()))
	}
	return dapdv.state
}

// DocumentAiProcessorDefaultVersionArgs contains the configurations for google_document_ai_processor_default_version.
type DocumentAiProcessorDefaultVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Processor: string, required
	Processor terra.StringValue `hcl:"processor,attr" validate:"required"`
	// Version: string, required
	Version terra.StringValue `hcl:"version,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *documentaiprocessordefaultversion.Timeouts `hcl:"timeouts,block"`
}
type documentAiProcessorDefaultVersionAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_document_ai_processor_default_version.
func (dapdv documentAiProcessorDefaultVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dapdv.ref.Append("id"))
}

// Processor returns a reference to field processor of google_document_ai_processor_default_version.
func (dapdv documentAiProcessorDefaultVersionAttributes) Processor() terra.StringValue {
	return terra.ReferenceAsString(dapdv.ref.Append("processor"))
}

// Version returns a reference to field version of google_document_ai_processor_default_version.
func (dapdv documentAiProcessorDefaultVersionAttributes) Version() terra.StringValue {
	return terra.ReferenceAsString(dapdv.ref.Append("version"))
}

func (dapdv documentAiProcessorDefaultVersionAttributes) Timeouts() documentaiprocessordefaultversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[documentaiprocessordefaultversion.TimeoutsAttributes](dapdv.ref.Append("timeouts"))
}

type documentAiProcessorDefaultVersionState struct {
	Id        string                                           `json:"id"`
	Processor string                                           `json:"processor"`
	Version   string                                           `json:"version"`
	Timeouts  *documentaiprocessordefaultversion.TimeoutsState `json:"timeouts"`
}
