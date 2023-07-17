// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package medialivemultiplexprogram

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type MultiplexProgramSettings struct {
	// PreferredChannelPipeline: string, required
	PreferredChannelPipeline terra.StringValue `hcl:"preferred_channel_pipeline,attr" validate:"required"`
	// ProgramNumber: number, required
	ProgramNumber terra.NumberValue `hcl:"program_number,attr" validate:"required"`
	// ServiceDescriptor: min=0
	ServiceDescriptor []ServiceDescriptor `hcl:"service_descriptor,block" validate:"min=0"`
	// VideoSettings: min=0
	VideoSettings []VideoSettings `hcl:"video_settings,block" validate:"min=0"`
}

type ServiceDescriptor struct {
	// ProviderName: string, required
	ProviderName terra.StringValue `hcl:"provider_name,attr" validate:"required"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
}

type VideoSettings struct {
	// ConstantBitrate: number, optional
	ConstantBitrate terra.NumberValue `hcl:"constant_bitrate,attr"`
	// StatmuxSettings: min=0
	StatmuxSettings []StatmuxSettings `hcl:"statmux_settings,block" validate:"min=0"`
}

type StatmuxSettings struct {
	// MaximumBitrate: number, optional
	MaximumBitrate terra.NumberValue `hcl:"maximum_bitrate,attr"`
	// MinimumBitrate: number, optional
	MinimumBitrate terra.NumberValue `hcl:"minimum_bitrate,attr"`
	// Priority: number, optional
	Priority terra.NumberValue `hcl:"priority,attr"`
}

type MultiplexProgramSettingsAttributes struct {
	ref terra.Reference
}

func (mps MultiplexProgramSettingsAttributes) InternalRef() (terra.Reference, error) {
	return mps.ref, nil
}

func (mps MultiplexProgramSettingsAttributes) InternalWithRef(ref terra.Reference) MultiplexProgramSettingsAttributes {
	return MultiplexProgramSettingsAttributes{ref: ref}
}

func (mps MultiplexProgramSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return mps.ref.InternalTokens()
}

func (mps MultiplexProgramSettingsAttributes) PreferredChannelPipeline() terra.StringValue {
	return terra.ReferenceAsString(mps.ref.Append("preferred_channel_pipeline"))
}

func (mps MultiplexProgramSettingsAttributes) ProgramNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(mps.ref.Append("program_number"))
}

func (mps MultiplexProgramSettingsAttributes) ServiceDescriptor() terra.ListValue[ServiceDescriptorAttributes] {
	return terra.ReferenceAsList[ServiceDescriptorAttributes](mps.ref.Append("service_descriptor"))
}

func (mps MultiplexProgramSettingsAttributes) VideoSettings() terra.ListValue[VideoSettingsAttributes] {
	return terra.ReferenceAsList[VideoSettingsAttributes](mps.ref.Append("video_settings"))
}

type ServiceDescriptorAttributes struct {
	ref terra.Reference
}

func (sd ServiceDescriptorAttributes) InternalRef() (terra.Reference, error) {
	return sd.ref, nil
}

func (sd ServiceDescriptorAttributes) InternalWithRef(ref terra.Reference) ServiceDescriptorAttributes {
	return ServiceDescriptorAttributes{ref: ref}
}

func (sd ServiceDescriptorAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sd.ref.InternalTokens()
}

func (sd ServiceDescriptorAttributes) ProviderName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("provider_name"))
}

func (sd ServiceDescriptorAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("service_name"))
}

type VideoSettingsAttributes struct {
	ref terra.Reference
}

func (vs VideoSettingsAttributes) InternalRef() (terra.Reference, error) {
	return vs.ref, nil
}

func (vs VideoSettingsAttributes) InternalWithRef(ref terra.Reference) VideoSettingsAttributes {
	return VideoSettingsAttributes{ref: ref}
}

func (vs VideoSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return vs.ref.InternalTokens()
}

func (vs VideoSettingsAttributes) ConstantBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(vs.ref.Append("constant_bitrate"))
}

func (vs VideoSettingsAttributes) StatmuxSettings() terra.ListValue[StatmuxSettingsAttributes] {
	return terra.ReferenceAsList[StatmuxSettingsAttributes](vs.ref.Append("statmux_settings"))
}

type StatmuxSettingsAttributes struct {
	ref terra.Reference
}

func (ss StatmuxSettingsAttributes) InternalRef() (terra.Reference, error) {
	return ss.ref, nil
}

func (ss StatmuxSettingsAttributes) InternalWithRef(ref terra.Reference) StatmuxSettingsAttributes {
	return StatmuxSettingsAttributes{ref: ref}
}

func (ss StatmuxSettingsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ss.ref.InternalTokens()
}

func (ss StatmuxSettingsAttributes) MaximumBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("maximum_bitrate"))
}

func (ss StatmuxSettingsAttributes) MinimumBitrate() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("minimum_bitrate"))
}

func (ss StatmuxSettingsAttributes) Priority() terra.NumberValue {
	return terra.ReferenceAsNumber(ss.ref.Append("priority"))
}

type MultiplexProgramSettingsState struct {
	PreferredChannelPipeline string                   `json:"preferred_channel_pipeline"`
	ProgramNumber            float64                  `json:"program_number"`
	ServiceDescriptor        []ServiceDescriptorState `json:"service_descriptor"`
	VideoSettings            []VideoSettingsState     `json:"video_settings"`
}

type ServiceDescriptorState struct {
	ProviderName string `json:"provider_name"`
	ServiceName  string `json:"service_name"`
}

type VideoSettingsState struct {
	ConstantBitrate float64                `json:"constant_bitrate"`
	StatmuxSettings []StatmuxSettingsState `json:"statmux_settings"`
}

type StatmuxSettingsState struct {
	MaximumBitrate float64 `json:"maximum_bitrate"`
	MinimumBitrate float64 `json:"minimum_bitrate"`
	Priority       float64 `json:"priority"`
}
