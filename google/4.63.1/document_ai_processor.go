// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	documentaiprocessor "github.com/golingon/terraproviders/google/4.63.1/documentaiprocessor"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDocumentAiProcessor creates a new instance of [DocumentAiProcessor].
func NewDocumentAiProcessor(name string, args DocumentAiProcessorArgs) *DocumentAiProcessor {
	return &DocumentAiProcessor{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DocumentAiProcessor)(nil)

// DocumentAiProcessor represents the Terraform resource google_document_ai_processor.
type DocumentAiProcessor struct {
	Name      string
	Args      DocumentAiProcessorArgs
	state     *documentAiProcessorState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DocumentAiProcessor].
func (dap *DocumentAiProcessor) Type() string {
	return "google_document_ai_processor"
}

// LocalName returns the local name for [DocumentAiProcessor].
func (dap *DocumentAiProcessor) LocalName() string {
	return dap.Name
}

// Configuration returns the configuration (args) for [DocumentAiProcessor].
func (dap *DocumentAiProcessor) Configuration() interface{} {
	return dap.Args
}

// DependOn is used for other resources to depend on [DocumentAiProcessor].
func (dap *DocumentAiProcessor) DependOn() terra.Reference {
	return terra.ReferenceResource(dap)
}

// Dependencies returns the list of resources [DocumentAiProcessor] depends_on.
func (dap *DocumentAiProcessor) Dependencies() terra.Dependencies {
	return dap.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DocumentAiProcessor].
func (dap *DocumentAiProcessor) LifecycleManagement() *terra.Lifecycle {
	return dap.Lifecycle
}

// Attributes returns the attributes for [DocumentAiProcessor].
func (dap *DocumentAiProcessor) Attributes() documentAiProcessorAttributes {
	return documentAiProcessorAttributes{ref: terra.ReferenceResource(dap)}
}

// ImportState imports the given attribute values into [DocumentAiProcessor]'s state.
func (dap *DocumentAiProcessor) ImportState(av io.Reader) error {
	dap.state = &documentAiProcessorState{}
	if err := json.NewDecoder(av).Decode(dap.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dap.Type(), dap.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DocumentAiProcessor] has state.
func (dap *DocumentAiProcessor) State() (*documentAiProcessorState, bool) {
	return dap.state, dap.state != nil
}

// StateMust returns the state for [DocumentAiProcessor]. Panics if the state is nil.
func (dap *DocumentAiProcessor) StateMust() *documentAiProcessorState {
	if dap.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dap.Type(), dap.LocalName()))
	}
	return dap.state
}

// DocumentAiProcessorArgs contains the configurations for google_document_ai_processor.
type DocumentAiProcessorArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Type: string, required
	Type terra.StringValue `hcl:"type,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *documentaiprocessor.Timeouts `hcl:"timeouts,block"`
}
type documentAiProcessorAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_document_ai_processor.
func (dap documentAiProcessorAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("display_name"))
}

// Id returns a reference to field id of google_document_ai_processor.
func (dap documentAiProcessorAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_document_ai_processor.
func (dap documentAiProcessorAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("kms_key_name"))
}

// Location returns a reference to field location of google_document_ai_processor.
func (dap documentAiProcessorAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("location"))
}

// Name returns a reference to field name of google_document_ai_processor.
func (dap documentAiProcessorAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("name"))
}

// Project returns a reference to field project of google_document_ai_processor.
func (dap documentAiProcessorAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("project"))
}

// Type returns a reference to field type of google_document_ai_processor.
func (dap documentAiProcessorAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(dap.ref.Append("type"))
}

func (dap documentAiProcessorAttributes) Timeouts() documentaiprocessor.TimeoutsAttributes {
	return terra.ReferenceAsSingle[documentaiprocessor.TimeoutsAttributes](dap.ref.Append("timeouts"))
}

type documentAiProcessorState struct {
	DisplayName string                             `json:"display_name"`
	Id          string                             `json:"id"`
	KmsKeyName  string                             `json:"kms_key_name"`
	Location    string                             `json:"location"`
	Name        string                             `json:"name"`
	Project     string                             `json:"project"`
	Type        string                             `json:"type"`
	Timeouts    *documentaiprocessor.TimeoutsState `json:"timeouts"`
}
