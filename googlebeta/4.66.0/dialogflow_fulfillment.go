// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	dialogflowfulfillment "github.com/golingon/terraproviders/googlebeta/4.66.0/dialogflowfulfillment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowFulfillment creates a new instance of [DialogflowFulfillment].
func NewDialogflowFulfillment(name string, args DialogflowFulfillmentArgs) *DialogflowFulfillment {
	return &DialogflowFulfillment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowFulfillment)(nil)

// DialogflowFulfillment represents the Terraform resource google_dialogflow_fulfillment.
type DialogflowFulfillment struct {
	Name      string
	Args      DialogflowFulfillmentArgs
	state     *dialogflowFulfillmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowFulfillment].
func (df *DialogflowFulfillment) Type() string {
	return "google_dialogflow_fulfillment"
}

// LocalName returns the local name for [DialogflowFulfillment].
func (df *DialogflowFulfillment) LocalName() string {
	return df.Name
}

// Configuration returns the configuration (args) for [DialogflowFulfillment].
func (df *DialogflowFulfillment) Configuration() interface{} {
	return df.Args
}

// DependOn is used for other resources to depend on [DialogflowFulfillment].
func (df *DialogflowFulfillment) DependOn() terra.Reference {
	return terra.ReferenceResource(df)
}

// Dependencies returns the list of resources [DialogflowFulfillment] depends_on.
func (df *DialogflowFulfillment) Dependencies() terra.Dependencies {
	return df.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowFulfillment].
func (df *DialogflowFulfillment) LifecycleManagement() *terra.Lifecycle {
	return df.Lifecycle
}

// Attributes returns the attributes for [DialogflowFulfillment].
func (df *DialogflowFulfillment) Attributes() dialogflowFulfillmentAttributes {
	return dialogflowFulfillmentAttributes{ref: terra.ReferenceResource(df)}
}

// ImportState imports the given attribute values into [DialogflowFulfillment]'s state.
func (df *DialogflowFulfillment) ImportState(av io.Reader) error {
	df.state = &dialogflowFulfillmentState{}
	if err := json.NewDecoder(av).Decode(df.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", df.Type(), df.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowFulfillment] has state.
func (df *DialogflowFulfillment) State() (*dialogflowFulfillmentState, bool) {
	return df.state, df.state != nil
}

// StateMust returns the state for [DialogflowFulfillment]. Panics if the state is nil.
func (df *DialogflowFulfillment) StateMust() *dialogflowFulfillmentState {
	if df.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", df.Type(), df.LocalName()))
	}
	return df.state
}

// DialogflowFulfillmentArgs contains the configurations for google_dialogflow_fulfillment.
type DialogflowFulfillmentArgs struct {
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Features: min=0
	Features []dialogflowfulfillment.Features `hcl:"features,block" validate:"min=0"`
	// GenericWebService: optional
	GenericWebService *dialogflowfulfillment.GenericWebService `hcl:"generic_web_service,block"`
	// Timeouts: optional
	Timeouts *dialogflowfulfillment.Timeouts `hcl:"timeouts,block"`
}
type dialogflowFulfillmentAttributes struct {
	ref terra.Reference
}

// DisplayName returns a reference to field display_name of google_dialogflow_fulfillment.
func (df dialogflowFulfillmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("display_name"))
}

// Enabled returns a reference to field enabled of google_dialogflow_fulfillment.
func (df dialogflowFulfillmentAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(df.ref.Append("enabled"))
}

// Id returns a reference to field id of google_dialogflow_fulfillment.
func (df dialogflowFulfillmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("id"))
}

// Name returns a reference to field name of google_dialogflow_fulfillment.
func (df dialogflowFulfillmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("name"))
}

// Project returns a reference to field project of google_dialogflow_fulfillment.
func (df dialogflowFulfillmentAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("project"))
}

func (df dialogflowFulfillmentAttributes) Features() terra.ListValue[dialogflowfulfillment.FeaturesAttributes] {
	return terra.ReferenceAsList[dialogflowfulfillment.FeaturesAttributes](df.ref.Append("features"))
}

func (df dialogflowFulfillmentAttributes) GenericWebService() terra.ListValue[dialogflowfulfillment.GenericWebServiceAttributes] {
	return terra.ReferenceAsList[dialogflowfulfillment.GenericWebServiceAttributes](df.ref.Append("generic_web_service"))
}

func (df dialogflowFulfillmentAttributes) Timeouts() dialogflowfulfillment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowfulfillment.TimeoutsAttributes](df.ref.Append("timeouts"))
}

type dialogflowFulfillmentState struct {
	DisplayName       string                                         `json:"display_name"`
	Enabled           bool                                           `json:"enabled"`
	Id                string                                         `json:"id"`
	Name              string                                         `json:"name"`
	Project           string                                         `json:"project"`
	Features          []dialogflowfulfillment.FeaturesState          `json:"features"`
	GenericWebService []dialogflowfulfillment.GenericWebServiceState `json:"generic_web_service"`
	Timeouts          *dialogflowfulfillment.TimeoutsState           `json:"timeouts"`
}
