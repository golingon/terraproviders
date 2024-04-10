// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	dialogflowcxsecuritysettings "github.com/golingon/terraproviders/google/5.24.0/dialogflowcxsecuritysettings"
	"io"
)

// NewDialogflowCxSecuritySettings creates a new instance of [DialogflowCxSecuritySettings].
func NewDialogflowCxSecuritySettings(name string, args DialogflowCxSecuritySettingsArgs) *DialogflowCxSecuritySettings {
	return &DialogflowCxSecuritySettings{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DialogflowCxSecuritySettings)(nil)

// DialogflowCxSecuritySettings represents the Terraform resource google_dialogflow_cx_security_settings.
type DialogflowCxSecuritySettings struct {
	Name      string
	Args      DialogflowCxSecuritySettingsArgs
	state     *dialogflowCxSecuritySettingsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) Type() string {
	return "google_dialogflow_cx_security_settings"
}

// LocalName returns the local name for [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) LocalName() string {
	return dcss.Name
}

// Configuration returns the configuration (args) for [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) Configuration() interface{} {
	return dcss.Args
}

// DependOn is used for other resources to depend on [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) DependOn() terra.Reference {
	return terra.ReferenceResource(dcss)
}

// Dependencies returns the list of resources [DialogflowCxSecuritySettings] depends_on.
func (dcss *DialogflowCxSecuritySettings) Dependencies() terra.Dependencies {
	return dcss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) LifecycleManagement() *terra.Lifecycle {
	return dcss.Lifecycle
}

// Attributes returns the attributes for [DialogflowCxSecuritySettings].
func (dcss *DialogflowCxSecuritySettings) Attributes() dialogflowCxSecuritySettingsAttributes {
	return dialogflowCxSecuritySettingsAttributes{ref: terra.ReferenceResource(dcss)}
}

// ImportState imports the given attribute values into [DialogflowCxSecuritySettings]'s state.
func (dcss *DialogflowCxSecuritySettings) ImportState(av io.Reader) error {
	dcss.state = &dialogflowCxSecuritySettingsState{}
	if err := json.NewDecoder(av).Decode(dcss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dcss.Type(), dcss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DialogflowCxSecuritySettings] has state.
func (dcss *DialogflowCxSecuritySettings) State() (*dialogflowCxSecuritySettingsState, bool) {
	return dcss.state, dcss.state != nil
}

// StateMust returns the state for [DialogflowCxSecuritySettings]. Panics if the state is nil.
func (dcss *DialogflowCxSecuritySettings) StateMust() *dialogflowCxSecuritySettingsState {
	if dcss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dcss.Type(), dcss.LocalName()))
	}
	return dcss.state
}

// DialogflowCxSecuritySettingsArgs contains the configurations for google_dialogflow_cx_security_settings.
type DialogflowCxSecuritySettingsArgs struct {
	// DeidentifyTemplate: string, optional
	DeidentifyTemplate terra.StringValue `hcl:"deidentify_template,attr"`
	// DisplayName: string, required
	DisplayName terra.StringValue `hcl:"display_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InspectTemplate: string, optional
	InspectTemplate terra.StringValue `hcl:"inspect_template,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// PurgeDataTypes: list of string, optional
	PurgeDataTypes terra.ListValue[terra.StringValue] `hcl:"purge_data_types,attr"`
	// RedactionScope: string, optional
	RedactionScope terra.StringValue `hcl:"redaction_scope,attr"`
	// RedactionStrategy: string, optional
	RedactionStrategy terra.StringValue `hcl:"redaction_strategy,attr"`
	// RetentionStrategy: string, optional
	RetentionStrategy terra.StringValue `hcl:"retention_strategy,attr"`
	// RetentionWindowDays: number, optional
	RetentionWindowDays terra.NumberValue `hcl:"retention_window_days,attr"`
	// AudioExportSettings: optional
	AudioExportSettings *dialogflowcxsecuritysettings.AudioExportSettings `hcl:"audio_export_settings,block"`
	// InsightsExportSettings: optional
	InsightsExportSettings *dialogflowcxsecuritysettings.InsightsExportSettings `hcl:"insights_export_settings,block"`
	// Timeouts: optional
	Timeouts *dialogflowcxsecuritysettings.Timeouts `hcl:"timeouts,block"`
}
type dialogflowCxSecuritySettingsAttributes struct {
	ref terra.Reference
}

// DeidentifyTemplate returns a reference to field deidentify_template of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) DeidentifyTemplate() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("deidentify_template"))
}

// DisplayName returns a reference to field display_name of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("display_name"))
}

// Id returns a reference to field id of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("id"))
}

// InspectTemplate returns a reference to field inspect_template of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) InspectTemplate() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("inspect_template"))
}

// Location returns a reference to field location of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("location"))
}

// Name returns a reference to field name of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("name"))
}

// Project returns a reference to field project of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("project"))
}

// PurgeDataTypes returns a reference to field purge_data_types of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) PurgeDataTypes() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dcss.ref.Append("purge_data_types"))
}

// RedactionScope returns a reference to field redaction_scope of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) RedactionScope() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("redaction_scope"))
}

// RedactionStrategy returns a reference to field redaction_strategy of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) RedactionStrategy() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("redaction_strategy"))
}

// RetentionStrategy returns a reference to field retention_strategy of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) RetentionStrategy() terra.StringValue {
	return terra.ReferenceAsString(dcss.ref.Append("retention_strategy"))
}

// RetentionWindowDays returns a reference to field retention_window_days of google_dialogflow_cx_security_settings.
func (dcss dialogflowCxSecuritySettingsAttributes) RetentionWindowDays() terra.NumberValue {
	return terra.ReferenceAsNumber(dcss.ref.Append("retention_window_days"))
}

func (dcss dialogflowCxSecuritySettingsAttributes) AudioExportSettings() terra.ListValue[dialogflowcxsecuritysettings.AudioExportSettingsAttributes] {
	return terra.ReferenceAsList[dialogflowcxsecuritysettings.AudioExportSettingsAttributes](dcss.ref.Append("audio_export_settings"))
}

func (dcss dialogflowCxSecuritySettingsAttributes) InsightsExportSettings() terra.ListValue[dialogflowcxsecuritysettings.InsightsExportSettingsAttributes] {
	return terra.ReferenceAsList[dialogflowcxsecuritysettings.InsightsExportSettingsAttributes](dcss.ref.Append("insights_export_settings"))
}

func (dcss dialogflowCxSecuritySettingsAttributes) Timeouts() dialogflowcxsecuritysettings.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dialogflowcxsecuritysettings.TimeoutsAttributes](dcss.ref.Append("timeouts"))
}

type dialogflowCxSecuritySettingsState struct {
	DeidentifyTemplate     string                                                     `json:"deidentify_template"`
	DisplayName            string                                                     `json:"display_name"`
	Id                     string                                                     `json:"id"`
	InspectTemplate        string                                                     `json:"inspect_template"`
	Location               string                                                     `json:"location"`
	Name                   string                                                     `json:"name"`
	Project                string                                                     `json:"project"`
	PurgeDataTypes         []string                                                   `json:"purge_data_types"`
	RedactionScope         string                                                     `json:"redaction_scope"`
	RedactionStrategy      string                                                     `json:"redaction_strategy"`
	RetentionStrategy      string                                                     `json:"retention_strategy"`
	RetentionWindowDays    float64                                                    `json:"retention_window_days"`
	AudioExportSettings    []dialogflowcxsecuritysettings.AudioExportSettingsState    `json:"audio_export_settings"`
	InsightsExportSettings []dialogflowcxsecuritysettings.InsightsExportSettingsState `json:"insights_export_settings"`
	Timeouts               *dialogflowcxsecuritysettings.TimeoutsState                `json:"timeouts"`
}
