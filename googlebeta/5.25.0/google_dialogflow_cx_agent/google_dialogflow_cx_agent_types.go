// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_dialogflow_cx_agent

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type AdvancedSettings struct {
	// AdvancedSettingsAudioExportGcsDestination: optional
	AudioExportGcsDestination *AdvancedSettingsAudioExportGcsDestination `hcl:"audio_export_gcs_destination,block"`
	// AdvancedSettingsDtmfSettings: optional
	DtmfSettings *AdvancedSettingsDtmfSettings `hcl:"dtmf_settings,block"`
}

type AdvancedSettingsAudioExportGcsDestination struct {
	// Uri: string, optional
	Uri terra.StringValue `hcl:"uri,attr"`
}

type AdvancedSettingsDtmfSettings struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// FinishDigit: string, optional
	FinishDigit terra.StringValue `hcl:"finish_digit,attr"`
	// MaxDigits: number, optional
	MaxDigits terra.NumberValue `hcl:"max_digits,attr"`
}

type GitIntegrationSettings struct {
	// GitIntegrationSettingsGithubSettings: optional
	GithubSettings *GitIntegrationSettingsGithubSettings `hcl:"github_settings,block"`
}

type GitIntegrationSettingsGithubSettings struct {
	// AccessToken: string, optional
	AccessToken terra.StringValue `hcl:"access_token,attr"`
	// Branches: list of string, optional
	Branches terra.ListValue[terra.StringValue] `hcl:"branches,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// RepositoryUri: string, optional
	RepositoryUri terra.StringValue `hcl:"repository_uri,attr"`
	// TrackingBranch: string, optional
	TrackingBranch terra.StringValue `hcl:"tracking_branch,attr"`
}

type SpeechToTextSettings struct {
	// EnableSpeechAdaptation: bool, optional
	EnableSpeechAdaptation terra.BoolValue `hcl:"enable_speech_adaptation,attr"`
}

type TextToSpeechSettings struct {
	// SynthesizeSpeechConfigs: string, optional
	SynthesizeSpeechConfigs terra.StringValue `hcl:"synthesize_speech_configs,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type AdvancedSettingsAttributes struct {
	ref terra.Reference
}

func (as AdvancedSettingsAttributes) InternalRef() (terra.Reference, error) {
	return as.ref, nil
}

func (as AdvancedSettingsAttributes) InternalWithRef(ref terra.Reference) AdvancedSettingsAttributes {
	return AdvancedSettingsAttributes{ref: ref}
}

func (as AdvancedSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return as.ref.InternalTokens()
}

func (as AdvancedSettingsAttributes) AudioExportGcsDestination() terra.ListValue[AdvancedSettingsAudioExportGcsDestinationAttributes] {
	return terra.ReferenceAsList[AdvancedSettingsAudioExportGcsDestinationAttributes](as.ref.Append("audio_export_gcs_destination"))
}

func (as AdvancedSettingsAttributes) DtmfSettings() terra.ListValue[AdvancedSettingsDtmfSettingsAttributes] {
	return terra.ReferenceAsList[AdvancedSettingsDtmfSettingsAttributes](as.ref.Append("dtmf_settings"))
}

type AdvancedSettingsAudioExportGcsDestinationAttributes struct {
	ref terra.Reference
}

func (aegd AdvancedSettingsAudioExportGcsDestinationAttributes) InternalRef() (terra.Reference, error) {
	return aegd.ref, nil
}

func (aegd AdvancedSettingsAudioExportGcsDestinationAttributes) InternalWithRef(ref terra.Reference) AdvancedSettingsAudioExportGcsDestinationAttributes {
	return AdvancedSettingsAudioExportGcsDestinationAttributes{ref: ref}
}

func (aegd AdvancedSettingsAudioExportGcsDestinationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return aegd.ref.InternalTokens()
}

func (aegd AdvancedSettingsAudioExportGcsDestinationAttributes) Uri() terra.StringValue {
	return terra.ReferenceAsString(aegd.ref.Append("uri"))
}

type AdvancedSettingsDtmfSettingsAttributes struct {
	ref terra.Reference
}

func (ds AdvancedSettingsDtmfSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ds.ref, nil
}

func (ds AdvancedSettingsDtmfSettingsAttributes) InternalWithRef(ref terra.Reference) AdvancedSettingsDtmfSettingsAttributes {
	return AdvancedSettingsDtmfSettingsAttributes{ref: ref}
}

func (ds AdvancedSettingsDtmfSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ds.ref.InternalTokens()
}

func (ds AdvancedSettingsDtmfSettingsAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(ds.ref.Append("enabled"))
}

func (ds AdvancedSettingsDtmfSettingsAttributes) FinishDigit() terra.StringValue {
	return terra.ReferenceAsString(ds.ref.Append("finish_digit"))
}

func (ds AdvancedSettingsDtmfSettingsAttributes) MaxDigits() terra.NumberValue {
	return terra.ReferenceAsNumber(ds.ref.Append("max_digits"))
}

type GitIntegrationSettingsAttributes struct {
	ref terra.Reference
}

func (gis GitIntegrationSettingsAttributes) InternalRef() (terra.Reference, error) {
	return gis.ref, nil
}

func (gis GitIntegrationSettingsAttributes) InternalWithRef(ref terra.Reference) GitIntegrationSettingsAttributes {
	return GitIntegrationSettingsAttributes{ref: ref}
}

func (gis GitIntegrationSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gis.ref.InternalTokens()
}

func (gis GitIntegrationSettingsAttributes) GithubSettings() terra.ListValue[GitIntegrationSettingsGithubSettingsAttributes] {
	return terra.ReferenceAsList[GitIntegrationSettingsGithubSettingsAttributes](gis.ref.Append("github_settings"))
}

type GitIntegrationSettingsGithubSettingsAttributes struct {
	ref terra.Reference
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) InternalRef() (terra.Reference, error) {
	return gs.ref, nil
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) InternalWithRef(ref terra.Reference) GitIntegrationSettingsGithubSettingsAttributes {
	return GitIntegrationSettingsGithubSettingsAttributes{ref: ref}
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return gs.ref.InternalTokens()
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) AccessToken() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("access_token"))
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) Branches() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](gs.ref.Append("branches"))
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("display_name"))
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) RepositoryUri() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("repository_uri"))
}

func (gs GitIntegrationSettingsGithubSettingsAttributes) TrackingBranch() terra.StringValue {
	return terra.ReferenceAsString(gs.ref.Append("tracking_branch"))
}

type SpeechToTextSettingsAttributes struct {
	ref terra.Reference
}

func (stts SpeechToTextSettingsAttributes) InternalRef() (terra.Reference, error) {
	return stts.ref, nil
}

func (stts SpeechToTextSettingsAttributes) InternalWithRef(ref terra.Reference) SpeechToTextSettingsAttributes {
	return SpeechToTextSettingsAttributes{ref: ref}
}

func (stts SpeechToTextSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return stts.ref.InternalTokens()
}

func (stts SpeechToTextSettingsAttributes) EnableSpeechAdaptation() terra.BoolValue {
	return terra.ReferenceAsBool(stts.ref.Append("enable_speech_adaptation"))
}

type TextToSpeechSettingsAttributes struct {
	ref terra.Reference
}

func (ttss TextToSpeechSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ttss.ref, nil
}

func (ttss TextToSpeechSettingsAttributes) InternalWithRef(ref terra.Reference) TextToSpeechSettingsAttributes {
	return TextToSpeechSettingsAttributes{ref: ref}
}

func (ttss TextToSpeechSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ttss.ref.InternalTokens()
}

func (ttss TextToSpeechSettingsAttributes) SynthesizeSpeechConfigs() terra.StringValue {
	return terra.ReferenceAsString(ttss.ref.Append("synthesize_speech_configs"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("update"))
}

type AdvancedSettingsState struct {
	AudioExportGcsDestination []AdvancedSettingsAudioExportGcsDestinationState `json:"audio_export_gcs_destination"`
	DtmfSettings              []AdvancedSettingsDtmfSettingsState              `json:"dtmf_settings"`
}

type AdvancedSettingsAudioExportGcsDestinationState struct {
	Uri string `json:"uri"`
}

type AdvancedSettingsDtmfSettingsState struct {
	Enabled     bool    `json:"enabled"`
	FinishDigit string  `json:"finish_digit"`
	MaxDigits   float64 `json:"max_digits"`
}

type GitIntegrationSettingsState struct {
	GithubSettings []GitIntegrationSettingsGithubSettingsState `json:"github_settings"`
}

type GitIntegrationSettingsGithubSettingsState struct {
	AccessToken    string   `json:"access_token"`
	Branches       []string `json:"branches"`
	DisplayName    string   `json:"display_name"`
	RepositoryUri  string   `json:"repository_uri"`
	TrackingBranch string   `json:"tracking_branch"`
}

type SpeechToTextSettingsState struct {
	EnableSpeechAdaptation bool `json:"enable_speech_adaptation"`
}

type TextToSpeechSettingsState struct {
	SynthesizeSpeechConfigs string `json:"synthesize_speech_configs"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
