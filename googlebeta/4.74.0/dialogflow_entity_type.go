// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowentitytype "github.com/golingon/terraproviders/googlebeta/4.74.0/dialogflowentitytype"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowEntityType creates a new instance of [DialogflowEntityType].
func NewDialogflowEntityType(name string, args DialogflowEntityTypeArgs) *DialogflowEntityType {
	return &DialogflowEntityType{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowEntityType)(nil)

// DialogflowEntityType represents the Terraform resource google_dialogflow_entity_type.
type DialogflowEntityType struct {
	Name      string
	Args      DialogflowEntityTypeArgs
	state     *dialogflowEntityTypeState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowEntityType].
func (det *DialogflowEntityType) Type() string {
	return "google_dialogflow_entity_type"
}

// LocalName returns the local name for [DialogflowEntityType].
func (det *DialogflowEntityType) LocalName() string {
	return det.Name
}

// Configuration returns the configuration (args) for [DialogflowEntityType].
func (det *DialogflowEntityType) Configuration() interface{} {
	return det.Args
}

// DependOn is used for other resources to depend on [DialogflowEntityType].
func (det *DialogflowEntityType) DependOn() terra.Reference {
	return terra.ReferenceResource(det)
}

// Dependencies returns the list of resources [DialogflowEntityType] depends_on.
func (det *DialogflowEntityType) Dependencies() terra.Dependencies {
	return det.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowEntityType].
func (det *DialogflowEntityType) LifecycleManagement() *terra.Lifecycle {
	return det.Lifecycle
}

// Attributes returns the attributes for [DialogflowEntityType].
func (det *DialogflowEntityType) Attributes() dialogflowEntityTypeAttributes {
	return dialogflowEntityTypeAttributes{ref: terra.ReferenceResource(det)}
}

// ImportState imports the given attribute values into [DialogflowEntityType]'s state.
func (det *DialogflowEntityType) ImportState(av io.Reader) error {
	det.state = &dialogflowEntityTypeState{}
	if err := json.NewDecoder(av).Decode(det.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", det.Type(), det.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowEntityType] has state.
func (det *DialogflowEntityType) State() (*dialogflowEntityTypeState, bool) {
	return det.state, det.state != nil
}

// StateMust returns the state for [DialogflowEntityType]. Panics if the state is nil.
func (det *DialogflowEntityType) StateMust() *dialogflowEntityTypeState {
	if det.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", det.Type(), det.LocalName()))
	}
	return det.state
}

// DialogflowEntityTypeArgs contains the configurations for google_dialogflow_entity_type.
type DialogflowEntityTypeArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableFuzzyExtraction: bool, optional
	EnableFuzzyExtraction terra.BoolValue `hcl:"enable_fuzzy_extraction,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Kind: string, required
	Kind terra.StringValue `hcl:"kind,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Entities: min=0
	Entities []dialogflowentitytype.Entities `hcl:"entities,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dialogflowentitytype.Timeouts `hcl:"timeouts,block"`
}
type dialogflowEntityTypeAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(det.ref.Append("display_name"))
}

// EnableFuzzyExtraction returns a reference to field enable_fuzzy_extraction of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) EnableFuzzyExtraction() terra.BoolValue {
	return terra.ReferenceAsBool(det.ref.Append("enable_fuzzy_extraction"))
}

// Id returns a reference to field id of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(det.ref.Append("id"))
}

// Kind returns a reference to field kind of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) Kind() terra.StringValue {
	return terra.ReferenceAsString(det.ref.Append("kind"))
}

// Name returns a reference to field name of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(det.ref.Append("name"))
}

// Project returns a reference to field project of google_dialogflow_entity_type.
func (det dialogflowEntityTypeAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(det.ref.Append("project"))
}

func (det dialogflowEntityTypeAttributes) Entities() terra.ListValue[dialogflowentitytype.EntitiesAttributes] {
	return terra.ReferenceAsList[dialogflowentitytype.EntitiesAttributes](det.ref.Append("entities"))
}

func (det dialogflowEntityTypeAttributes) Timeouts() dialogflowentitytype.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowentitytype.TimeoutsAttributes](det.ref.Append("timeouts"))
}

type dialogflowEntityTypeState struct {
	DisplayName           string                               `json:"display_name"`
	EnableFuzzyExtraction bool                                 `json:"enable_fuzzy_extraction"`
	Id                    string                               `json:"id"`
	Kind                  string                               `json:"kind"`
	Name                  string                               `json:"name"`
	Project               string                               `json:"project"`
	Entities              []dialogflowentitytype.EntitiesState `json:"entities"`
	Timeouts              *dialogflowentitytype.TimeoutsState  `json:"timeouts"`
}
