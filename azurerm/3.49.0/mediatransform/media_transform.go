// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package mediatransform

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Output struct {
	// OnErrorAction: string, optional
	OnErrorAction terra.StringValue `hcl:"on_error_action,attr"`
	// RelativePriority: string, optional
	RelativePriority terra.StringValue `hcl:"relative_priority,attr"`
	// AudioAnalyzerPreset: optional
	AudioAnalyzerPreset *AudioAnalyzerPreset `hcl:"audio_analyzer_preset,block"`
	// BuiltinPreset: optional
	BuiltinPreset *BuiltinPreset `hcl:"builtin_preset,block"`
	// FaceDetectorPreset: optional
	FaceDetectorPreset *FaceDetectorPreset `hcl:"face_detector_preset,block"`
	// VideoAnalyzerPreset: optional
	VideoAnalyzerPreset *VideoAnalyzerPreset `hcl:"video_analyzer_preset,block"`
}

type AudioAnalyzerPreset struct {
	// AudioAnalysisMode: string, optional
	AudioAnalysisMode terra.StringValue `hcl:"audio_analysis_mode,attr"`
	// AudioLanguage: string, optional
	AudioLanguage terra.StringValue `hcl:"audio_language,attr"`
}

type BuiltinPreset struct {
	// PresetName: string, required
	PresetName terra.StringValue `hcl:"preset_name,attr" validate:"required"`
}

type FaceDetectorPreset struct {
	// AnalysisResolution: string, optional
	AnalysisResolution terra.StringValue `hcl:"analysis_resolution,attr"`
}

type VideoAnalyzerPreset struct {
	// AudioAnalysisMode: string, optional
	AudioAnalysisMode terra.StringValue `hcl:"audio_analysis_mode,attr"`
	// AudioLanguage: string, optional
	AudioLanguage terra.StringValue `hcl:"audio_language,attr"`
	// InsightsType: string, optional
	InsightsType terra.StringValue `hcl:"insights_type,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type OutputAttributes struct {
	ref terra.Reference
}

func (o OutputAttributes) InternalRef() terra.Reference {
	return o.ref
}

func (o OutputAttributes) InternalWithRef(ref terra.Reference) OutputAttributes {
	return OutputAttributes{ref: ref}
}

func (o OutputAttributes) InternalTokens() hclwrite.Tokens {
	return o.ref.InternalTokens()
}

func (o OutputAttributes) OnErrorAction() terra.StringValue {
	return terra.ReferenceString(o.ref.Append("on_error_action"))
}

func (o OutputAttributes) RelativePriority() terra.StringValue {
	return terra.ReferenceString(o.ref.Append("relative_priority"))
}

func (o OutputAttributes) AudioAnalyzerPreset() terra.ListValue[AudioAnalyzerPresetAttributes] {
	return terra.ReferenceList[AudioAnalyzerPresetAttributes](o.ref.Append("audio_analyzer_preset"))
}

func (o OutputAttributes) BuiltinPreset() terra.ListValue[BuiltinPresetAttributes] {
	return terra.ReferenceList[BuiltinPresetAttributes](o.ref.Append("builtin_preset"))
}

func (o OutputAttributes) FaceDetectorPreset() terra.ListValue[FaceDetectorPresetAttributes] {
	return terra.ReferenceList[FaceDetectorPresetAttributes](o.ref.Append("face_detector_preset"))
}

func (o OutputAttributes) VideoAnalyzerPreset() terra.ListValue[VideoAnalyzerPresetAttributes] {
	return terra.ReferenceList[VideoAnalyzerPresetAttributes](o.ref.Append("video_analyzer_preset"))
}

type AudioAnalyzerPresetAttributes struct {
	ref terra.Reference
}

func (aap AudioAnalyzerPresetAttributes) InternalRef() terra.Reference {
	return aap.ref
}

func (aap AudioAnalyzerPresetAttributes) InternalWithRef(ref terra.Reference) AudioAnalyzerPresetAttributes {
	return AudioAnalyzerPresetAttributes{ref: ref}
}

func (aap AudioAnalyzerPresetAttributes) InternalTokens() hclwrite.Tokens {
	return aap.ref.InternalTokens()
}

func (aap AudioAnalyzerPresetAttributes) AudioAnalysisMode() terra.StringValue {
	return terra.ReferenceString(aap.ref.Append("audio_analysis_mode"))
}

func (aap AudioAnalyzerPresetAttributes) AudioLanguage() terra.StringValue {
	return terra.ReferenceString(aap.ref.Append("audio_language"))
}

type BuiltinPresetAttributes struct {
	ref terra.Reference
}

func (bp BuiltinPresetAttributes) InternalRef() terra.Reference {
	return bp.ref
}

func (bp BuiltinPresetAttributes) InternalWithRef(ref terra.Reference) BuiltinPresetAttributes {
	return BuiltinPresetAttributes{ref: ref}
}

func (bp BuiltinPresetAttributes) InternalTokens() hclwrite.Tokens {
	return bp.ref.InternalTokens()
}

func (bp BuiltinPresetAttributes) PresetName() terra.StringValue {
	return terra.ReferenceString(bp.ref.Append("preset_name"))
}

type FaceDetectorPresetAttributes struct {
	ref terra.Reference
}

func (fdp FaceDetectorPresetAttributes) InternalRef() terra.Reference {
	return fdp.ref
}

func (fdp FaceDetectorPresetAttributes) InternalWithRef(ref terra.Reference) FaceDetectorPresetAttributes {
	return FaceDetectorPresetAttributes{ref: ref}
}

func (fdp FaceDetectorPresetAttributes) InternalTokens() hclwrite.Tokens {
	return fdp.ref.InternalTokens()
}

func (fdp FaceDetectorPresetAttributes) AnalysisResolution() terra.StringValue {
	return terra.ReferenceString(fdp.ref.Append("analysis_resolution"))
}

type VideoAnalyzerPresetAttributes struct {
	ref terra.Reference
}

func (vap VideoAnalyzerPresetAttributes) InternalRef() terra.Reference {
	return vap.ref
}

func (vap VideoAnalyzerPresetAttributes) InternalWithRef(ref terra.Reference) VideoAnalyzerPresetAttributes {
	return VideoAnalyzerPresetAttributes{ref: ref}
}

func (vap VideoAnalyzerPresetAttributes) InternalTokens() hclwrite.Tokens {
	return vap.ref.InternalTokens()
}

func (vap VideoAnalyzerPresetAttributes) AudioAnalysisMode() terra.StringValue {
	return terra.ReferenceString(vap.ref.Append("audio_analysis_mode"))
}

func (vap VideoAnalyzerPresetAttributes) AudioLanguage() terra.StringValue {
	return terra.ReferenceString(vap.ref.Append("audio_language"))
}

func (vap VideoAnalyzerPresetAttributes) InsightsType() terra.StringValue {
	return terra.ReferenceString(vap.ref.Append("insights_type"))
}

type TimeoutsAttributes struct {
	ref terra.Reference
}

func (t TimeoutsAttributes) InternalRef() terra.Reference {
	return t.ref
}

func (t TimeoutsAttributes) InternalWithRef(ref terra.Reference) TimeoutsAttributes {
	return TimeoutsAttributes{ref: ref}
}

func (t TimeoutsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t TimeoutsAttributes) Create() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("create"))
}

func (t TimeoutsAttributes) Delete() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("delete"))
}

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("read"))
}

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type OutputState struct {
	OnErrorAction       string                     `json:"on_error_action"`
	RelativePriority    string                     `json:"relative_priority"`
	AudioAnalyzerPreset []AudioAnalyzerPresetState `json:"audio_analyzer_preset"`
	BuiltinPreset       []BuiltinPresetState       `json:"builtin_preset"`
	FaceDetectorPreset  []FaceDetectorPresetState  `json:"face_detector_preset"`
	VideoAnalyzerPreset []VideoAnalyzerPresetState `json:"video_analyzer_preset"`
}

type AudioAnalyzerPresetState struct {
	AudioAnalysisMode string `json:"audio_analysis_mode"`
	AudioLanguage     string `json:"audio_language"`
}

type BuiltinPresetState struct {
	PresetName string `json:"preset_name"`
}

type FaceDetectorPresetState struct {
	AnalysisResolution string `json:"analysis_resolution"`
}

type VideoAnalyzerPresetState struct {
	AudioAnalysisMode string `json:"audio_analysis_mode"`
	AudioLanguage     string `json:"audio_language"`
	InsightsType      string `json:"insights_type"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
	Update string `json:"update"`
}
