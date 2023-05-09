// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	dialogflowcxwebhook "github.com/golingon/terraproviders/google/4.63.1/dialogflowcxwebhook"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDialogflowCxWebhook creates a new instance of [DialogflowCxWebhook].
func NewDialogflowCxWebhook(name string, args DialogflowCxWebhookArgs) *DialogflowCxWebhook {
	return &DialogflowCxWebhook{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxWebhook)(nil)

// DialogflowCxWebhook represents the Terraform resource google_dialogflow_cx_webhook.
type DialogflowCxWebhook struct {
	Name      string
	Args      DialogflowCxWebhookArgs
	state     *dialogflowCxWebhookState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) Type() string {
	return "google_dialogflow_cx_webhook"
}

// LocalName returns the local name for [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) LocalName() string {
	return dcw.Name
}

// Configuration returns the configuration (args) for [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) Configuration() interface{} {
	return dcw.Args
}

// DependOn is used for other resources to depend on [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) DependOn() terra.Reference {
	return terra.ReferenceResource(dcw)
}

// Dependencies returns the list of resources [DialogflowCxWebhook] depends_on.
func (dcw *DialogflowCxWebhook) Dependencies() terra.Dependencies {
	return dcw.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) LifecycleManagement() *terra.Lifecycle {
	return dcw.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxWebhook].
func (dcw *DialogflowCxWebhook) Attributes() dialogflowCxWebhookAttributes {
	return dialogflowCxWebhookAttributes{ref: terra.ReferenceResource(dcw)}
}

// ImportState imports the given attribute values into [DialogflowCxWebhook]'s state.
func (dcw *DialogflowCxWebhook) ImportState(av io.Reader) error {
	dcw.state = &dialogflowCxWebhookState{}
	if err := json.NewDecoder(av).Decode(dcw.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcw.Type(), dcw.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxWebhook] has state.
func (dcw *DialogflowCxWebhook) State() (*dialogflowCxWebhookState, bool) {
	return dcw.state, dcw.state != nil
}

// StateMust returns the state for [DialogflowCxWebhook]. Panics if the state is nil.
func (dcw *DialogflowCxWebhook) StateMust() *dialogflowCxWebhookState {
	if dcw.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcw.Type(), dcw.LocalName()))
	}
	return dcw.state
}

// DialogflowCxWebhookArgs contains the configurations for google_dialogflow_cx_webhook.
type DialogflowCxWebhookArgs struct {
	// Disabled: bool, optional
	Disabled terra.BoolValue `hcl:"disabled,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// EnableSpellCorrection: bool, optional
	EnableSpellCorrection terra.BoolValue `hcl:"enable_spell_correction,attr"`
	// EnableStackdriverLogging: bool, optional
	EnableStackdriverLogging terra.BoolValue `hcl:"enable_stackdriver_logging,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Parent: string, optional
	Parent terra.StringValue `hcl:"parent,attr"`
	// SecuritySettings: string, optional
	SecuritySettings terra.StringValue `hcl:"security_settings,attr"`
	// Timeout: string, optional
	Timeout terra.StringValue `hcl:"timeout,attr"`
	// GenericWebService: optional
	GenericWebService *dialogflowcxwebhook.GenericWebService `hcl:"generic_web_service,block"`
	// ServiceDirectory: optional
	ServiceDirectory *dialogflowcxwebhook.ServiceDirectory `hcl:"service_directory,block"`
	// Timeouts: optional
	Timeouts *dialogflowcxwebhook.Timeouts `hcl:"timeouts,block"`
}
type dialogflowCxWebhookAttributes struct {
	ref terra.Reference
}

// Disabled returns a reference to field disabled of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) Disabled() terra.BoolValue {
	return terra.ReferenceAsBool(dcw.ref.Append("disabled"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("display_name"))
}

// EnableSpellCorrection returns a reference to field enable_spell_correction of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) EnableSpellCorrection() terra.BoolValue {
	return terra.ReferenceAsBool(dcw.ref.Append("enable_spell_correction"))
}

// EnableStackdriverLogging returns a reference to field enable_stackdriver_logging of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) EnableStackdriverLogging() terra.BoolValue {
	return terra.ReferenceAsBool(dcw.ref.Append("enable_stackdriver_logging"))
}

// Id returns a reference to field id of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("id"))
}

// Name returns a reference to field name of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("name"))
}

// Parent returns a reference to field parent of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) Parent() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("parent"))
}

// SecuritySettings returns a reference to field security_settings of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) SecuritySettings() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("security_settings"))
}

// StartFlow returns a reference to field start_flow of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) StartFlow() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("start_flow"))
}

// Timeout returns a reference to field timeout of google_dialogflow_cx_webhook.
func (dcw dialogflowCxWebhookAttributes) Timeout() terra.StringValue {
	return terra.ReferenceAsString(dcw.ref.Append("timeout"))
}

func (dcw dialogflowCxWebhookAttributes) GenericWebService() terra.ListValue[dialogflowcxwebhook.GenericWebServiceAttributes] {
	return terra.ReferenceAsList[dialogflowcxwebhook.GenericWebServiceAttributes](dcw.ref.Append("generic_web_service"))
}

func (dcw dialogflowCxWebhookAttributes) ServiceDirectory() terra.ListValue[dialogflowcxwebhook.ServiceDirectoryAttributes] {
	return terra.ReferenceAsList[dialogflowcxwebhook.ServiceDirectoryAttributes](dcw.ref.Append("service_directory"))
}

func (dcw dialogflowCxWebhookAttributes) Timeouts() dialogflowcxwebhook.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxwebhook.TimeoutsAttributes](dcw.ref.Append("timeouts"))
}

type dialogflowCxWebhookState struct {
	Disabled                 bool                                         `json:"disabled"`
	DisplayName              string                                       `json:"display_name"`
	EnableSpellCorrection    bool                                         `json:"enable_spell_correction"`
	EnableStackdriverLogging bool                                         `json:"enable_stackdriver_logging"`
	Id                       string                                       `json:"id"`
	Name                     string                                       `json:"name"`
	Parent                   string                                       `json:"parent"`
	SecuritySettings         string                                       `json:"security_settings"`
	StartFlow                string                                       `json:"start_flow"`
	Timeout                  string                                       `json:"timeout"`
	GenericWebService        []dialogflowcxwebhook.GenericWebServiceState `json:"generic_web_service"`
	ServiceDirectory         []dialogflowcxwebhook.ServiceDirectoryState  `json:"service_directory"`
	Timeouts                 *dialogflowcxwebhook.TimeoutsState           `json:"timeouts"`
}
