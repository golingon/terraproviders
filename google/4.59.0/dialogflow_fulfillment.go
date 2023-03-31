// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowfulfillment "github.com/golingon/terraproviders/google/4.59.0/dialogflowfulfillment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDialogflowFulfillment(name string, args DialogflowFulfillmentArgs) *DialogflowFulfillment {
	return &DialogflowFulfillment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowFulfillment)(nil)

type DialogflowFulfillment struct {
	Name  string
	Args  DialogflowFulfillmentArgs
	state *dialogflowFulfillmentState
}

func (df *DialogflowFulfillment) Type() string {
	return "google_dialogflow_fulfillment"
}

func (df *DialogflowFulfillment) LocalName() string {
	return df.Name
}

func (df *DialogflowFulfillment) Configuration() interface{} {
	return df.Args
}

func (df *DialogflowFulfillment) Attributes() dialogflowFulfillmentAttributes {
	return dialogflowFulfillmentAttributes{ref: terra.ReferenceResource(df)}
}

func (df *DialogflowFulfillment) ImportState(av io.Reader) error {
	df.state = &dialogflowFulfillmentState{}
	if err := json.NewDecoder(av).Decode(df.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", df.Type(), df.LocalName(), err)
	}
	return nil
}

func (df *DialogflowFulfillment) State() (*dialogflowFulfillmentState, bool) {
	return df.state, df.state != nil
}

func (df *DialogflowFulfillment) StateMust() *dialogflowFulfillmentState {
	if df.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", df.Type(), df.LocalName()))
	}
	return df.state
}

func (df *DialogflowFulfillment) DependOn() terra.Reference {
	return terra.ReferenceResource(df)
}

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
	// DependsOn contains resources that DialogflowFulfillment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dialogflowFulfillmentAttributes struct {
	ref terra.Reference
}

func (df dialogflowFulfillmentAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceString(df.ref.Append("display_name"))
}

func (df dialogflowFulfillmentAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceBool(df.ref.Append("enabled"))
}

func (df dialogflowFulfillmentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(df.ref.Append("id"))
}

func (df dialogflowFulfillmentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(df.ref.Append("name"))
}

func (df dialogflowFulfillmentAttributes) Project() terra.StringValue {
	return terra.ReferenceString(df.ref.Append("project"))
}

func (df dialogflowFulfillmentAttributes) Features() terra.ListValue[dialogflowfulfillment.FeaturesAttributes] {
	return terra.ReferenceList[dialogflowfulfillment.FeaturesAttributes](df.ref.Append("features"))
}

func (df dialogflowFulfillmentAttributes) GenericWebService() terra.ListValue[dialogflowfulfillment.GenericWebServiceAttributes] {
	return terra.ReferenceList[dialogflowfulfillment.GenericWebServiceAttributes](df.ref.Append("generic_web_service"))
}

func (df dialogflowFulfillmentAttributes) Timeouts() dialogflowfulfillment.TimeoutsAttributes {
	return terra.ReferenceSingle[dialogflowfulfillment.TimeoutsAttributes](df.ref.Append("timeouts"))
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
