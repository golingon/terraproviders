// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowcxentitytype "github.com/golingon/terraproviders/googlebeta/4.71.0/dialogflowcxentitytype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxEntityType creates a new instance of [DialogflowCxEntityType].
func NewDialogflowCxEntityType(name string, args DialogflowCxEntityTypeArgs) *DialogflowCxEntityType {
	return &DialogflowCxEntityType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxEntityType)(nil)

// DialogflowCxEntityType represents the Terraform resource google_dialogflow_cx_entity_type.
type DialogflowCxEntityType struct {
	Name      string
	Args      DialogflowCxEntityTypeArgs
	state     *dialogflowCxEntityTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) Type() string {
	return "google_dialogflow_cx_entity_type"
}

// LocalName returns the local name for [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) LocalName() string {
	return dcet.Name
}

// Configuration returns the configuration (args) for [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) Configuration() interface{} {
	return dcet.Args
}

// DependOn is used for other resources to depend on [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) DependOn() terra.Reference {
	return terra.ReferenceResource(dcet)
}

// Dependencies returns the list of resources [DialogflowCxEntityType] depends_on.
func (dcet *DialogflowCxEntityType) Dependencies() terra.Dependencies {
	return dcet.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) LifecycleManagement() *terra.Lifecycle {
	return dcet.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxEntityType].
func (dcet *DialogflowCxEntityType) Attributes() dialogflowCxEntityTypeAttributes {
	return dialogflowCxEntityTypeAttributes{ref: terra.ReferenceResource(dcet)}
}

// ImportState imports the given attribute values into [DialogflowCxEntityType]'s state.
func (dcet *DialogflowCxEntityType) ImportState(av io.Reader) error {
	dcet.state = &dialogflowCxEntityTypeState{}
	if err := json.NewDecoder(av).Decode(dcet.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcet.Type(), dcet.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxEntityType] has state.
func (dcet *DialogflowCxEntityType) State() (*dialogflowCxEntityTypeState, bool) {
	return dcet.state, dcet.state != nil
}

// StateMust returns the state for [DialogflowCxEntityType]. Panics if the state is nil.
func (dcet *DialogflowCxEntityType) StateMust() *dialogflowCxEntityTypeState {
	if dcet.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcet.Type(), dcet.LocalName()))
	}
	return dcet.state
}

// DialogflowCxEntityTypeArgs contains the configurations for google_dialogflow_cx_entity_type.
type DialogflowCxEntityTypeArgs struct {
	// AutoExpansionMode: string, optional
	AutoExpansionMode terra.StringValue `hcl:"auto_expansion_mode,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableFuzzyExtraction: bool, optional
	EnableFuzzyExtraction terra.BoolValue `hcl:"enable_fuzzy_extraction,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// LanguageCode: string, optional
	LanguageCode terra.StringValue `hcl:"language_code,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// Redact: bool, optional
	Redact terra.BoolValue `hcl:"redact,attr"`
	// Entities: min=1
	Entities []dialogflowcxentitytype.Entities `hcl:"entities,block" validate:"min=1"`
	// ExcludedPhrases: min=0
	ExcludedPhrases []dialogflowcxentitytype.ExcludedPhrases `hcl:"excluded_phrases,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dialogflowcxentitytype.Timeouts `hcl:"timeouts,block"`
}
type dialogflowCxEntityTypeAttributes struct {
	ref terra.Reference
}

// AutoExpansionMode returns a reference to field auto_expansion_mode of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) AutoExpansionMode() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("auto_expansion_mode"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("display_name"))
}

// EnableFuzzyExtraction returns a reference to field enable_fuzzy_extraction of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) EnableFuzzyExtraction() terra.BoolValue {
	return terra.ReferenceAsBool(dcet.ref.Append("enable_fuzzy_extraction"))
}

// Id returns a reference to field id of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("id"))
}

// Kind returns a reference to field kind of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("kind"))
}

// LanguageCode returns a reference to field language_code of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) LanguageCode() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("language_code"))
}

// Name returns a reference to field name of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dcet.ref.Append("parent"))
}

// Redact returns a reference to field redact of google_dialogflow_cx_entity_type.
func (dcet dialogflowCxEntityTypeAttributes) Redact() terra.BoolValue {
	return terra.ReferenceAsBool(dcet.ref.Append("redact"))
}

func (dcet dialogflowCxEntityTypeAttributes) Entities() terra.ListValue[dialogflowcxentitytype.EntitiesAttributes] {
	return terra.ReferenceAsList[dialogflowcxentitytype.EntitiesAttributes](dcet.ref.Append("entities"))
}

func (dcet dialogflowCxEntityTypeAttributes) ExcludedPhrases() terra.ListValue[dialogflowcxentitytype.ExcludedPhrasesAttributes] {
	return terra.ReferenceAsList[dialogflowcxentitytype.ExcludedPhrasesAttributes](dcet.ref.Append("excluded_phrases"))
}

func (dcet dialogflowCxEntityTypeAttributes) Timeouts() dialogflowcxentitytype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxentitytype.TimeoutsAttributes](dcet.ref.Append("timeouts"))
}

type dialogflowCxEntityTypeState struct {
	AutoExpansionMode     string                                        `json:"auto_expansion_mode"`
	DisplayName           string                                        `json:"display_name"`
	EnableFuzzyExtraction bool                                          `json:"enable_fuzzy_extraction"`
	Id                    string                                        `json:"id"`
	Kind                  string                                        `json:"kind"`
	LanguageCode          string                                        `json:"language_code"`
	Name                  string                                        `json:"name"`
	Parent                string                                        `json:"parent"`
	Redact                bool                                          `json:"redact"`
	Entities              []dialogflowcxentitytype.EntitiesState        `json:"entities"`
	ExcludedPhrases       []dialogflowcxentitytype.ExcludedPhrasesState `json:"excluded_phrases"`
	Timeouts              *dialogflowcxentitytype.TimeoutsState         `json:"timeouts"`
}
