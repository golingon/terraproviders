// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_media_streaming_policy

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CommonEncryptionCbcs struct {
	// CommonEncryptionCbcsClearKeyEncryption: optional
	ClearKeyEncryption *CommonEncryptionCbcsClearKeyEncryption `hcl:"clear_key_encryption,block"`
	// CommonEncryptionCbcsDefaultContentKey: optional
	DefaultContentKey *CommonEncryptionCbcsDefaultContentKey `hcl:"default_content_key,block"`
	// CommonEncryptionCbcsDrmFairplay: optional
	DrmFairplay *CommonEncryptionCbcsDrmFairplay `hcl:"drm_fairplay,block"`
	// CommonEncryptionCbcsEnabledProtocols: optional
	EnabledProtocols *CommonEncryptionCbcsEnabledProtocols `hcl:"enabled_protocols,block"`
}

type CommonEncryptionCbcsClearKeyEncryption struct {
	// CustomKeysAcquisitionUrlTemplate: string, required
	CustomKeysAcquisitionUrlTemplate terra.StringValue `hcl:"custom_keys_acquisition_url_template,attr" validate:"required"`
}

type CommonEncryptionCbcsDefaultContentKey struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// PolicyName: string, optional
	PolicyName terra.StringValue `hcl:"policy_name,attr"`
}

type CommonEncryptionCbcsDrmFairplay struct {
	// AllowPersistentLicense: bool, optional
	AllowPersistentLicense terra.BoolValue `hcl:"allow_persistent_license,attr"`
	// CustomLicenseAcquisitionUrlTemplate: string, optional
	CustomLicenseAcquisitionUrlTemplate terra.StringValue `hcl:"custom_license_acquisition_url_template,attr"`
}

type CommonEncryptionCbcsEnabledProtocols struct {
	// Dash: bool, optional
	Dash terra.BoolValue `hcl:"dash,attr"`
	// Download: bool, optional
	Download terra.BoolValue `hcl:"download,attr"`
	// Hls: bool, optional
	Hls terra.BoolValue `hcl:"hls,attr"`
	// SmoothStreaming: bool, optional
	SmoothStreaming terra.BoolValue `hcl:"smooth_streaming,attr"`
}

type CommonEncryptionCenc struct {
	// DrmWidevineCustomLicenseAcquisitionUrlTemplate: string, optional
	DrmWidevineCustomLicenseAcquisitionUrlTemplate terra.StringValue `hcl:"drm_widevine_custom_license_acquisition_url_template,attr"`
	// CommonEncryptionCencClearKeyEncryption: optional
	ClearKeyEncryption *CommonEncryptionCencClearKeyEncryption `hcl:"clear_key_encryption,block"`
	// CommonEncryptionCencClearTrack: min=0
	ClearTrack []CommonEncryptionCencClearTrack `hcl:"clear_track,block" validate:"min=0"`
	// CommonEncryptionCencContentKeyToTrackMapping: min=0
	ContentKeyToTrackMapping []CommonEncryptionCencContentKeyToTrackMapping `hcl:"content_key_to_track_mapping,block" validate:"min=0"`
	// CommonEncryptionCencDefaultContentKey: optional
	DefaultContentKey *CommonEncryptionCencDefaultContentKey `hcl:"default_content_key,block"`
	// CommonEncryptionCencDrmPlayready: optional
	DrmPlayready *CommonEncryptionCencDrmPlayready `hcl:"drm_playready,block"`
	// CommonEncryptionCencEnabledProtocols: optional
	EnabledProtocols *CommonEncryptionCencEnabledProtocols `hcl:"enabled_protocols,block"`
}

type CommonEncryptionCencClearKeyEncryption struct {
	// CustomKeysAcquisitionUrlTemplate: string, required
	CustomKeysAcquisitionUrlTemplate terra.StringValue `hcl:"custom_keys_acquisition_url_template,attr" validate:"required"`
}

type CommonEncryptionCencClearTrack struct {
	// CommonEncryptionCencClearTrackCondition: min=1
	Condition []CommonEncryptionCencClearTrackCondition `hcl:"condition,block" validate:"min=1"`
}

type CommonEncryptionCencClearTrackCondition struct {
	// Operation: string, required
	Operation terra.StringValue `hcl:"operation,attr" validate:"required"`
	// Property: string, required
	Property terra.StringValue `hcl:"property,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type CommonEncryptionCencContentKeyToTrackMapping struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// PolicyName: string, optional
	PolicyName terra.StringValue `hcl:"policy_name,attr"`
	// CommonEncryptionCencContentKeyToTrackMappingTrack: min=1
	Track []CommonEncryptionCencContentKeyToTrackMappingTrack `hcl:"track,block" validate:"min=1"`
}

type CommonEncryptionCencContentKeyToTrackMappingTrack struct {
	// CommonEncryptionCencContentKeyToTrackMappingTrackCondition: min=1
	Condition []CommonEncryptionCencContentKeyToTrackMappingTrackCondition `hcl:"condition,block" validate:"min=1"`
}

type CommonEncryptionCencContentKeyToTrackMappingTrackCondition struct {
	// Operation: string, required
	Operation terra.StringValue `hcl:"operation,attr" validate:"required"`
	// Property: string, required
	Property terra.StringValue `hcl:"property,attr" validate:"required"`
	// Value: string, required
	Value terra.StringValue `hcl:"value,attr" validate:"required"`
}

type CommonEncryptionCencDefaultContentKey struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// PolicyName: string, optional
	PolicyName terra.StringValue `hcl:"policy_name,attr"`
}

type CommonEncryptionCencDrmPlayready struct {
	// CustomAttributes: string, optional
	CustomAttributes terra.StringValue `hcl:"custom_attributes,attr"`
	// CustomLicenseAcquisitionUrlTemplate: string, optional
	CustomLicenseAcquisitionUrlTemplate terra.StringValue `hcl:"custom_license_acquisition_url_template,attr"`
}

type CommonEncryptionCencEnabledProtocols struct {
	// Dash: bool, optional
	Dash terra.BoolValue `hcl:"dash,attr"`
	// Download: bool, optional
	Download terra.BoolValue `hcl:"download,attr"`
	// Hls: bool, optional
	Hls terra.BoolValue `hcl:"hls,attr"`
	// SmoothStreaming: bool, optional
	SmoothStreaming terra.BoolValue `hcl:"smooth_streaming,attr"`
}

type EnvelopeEncryption struct {
	// CustomKeysAcquisitionUrlTemplate: string, optional
	CustomKeysAcquisitionUrlTemplate terra.StringValue `hcl:"custom_keys_acquisition_url_template,attr"`
	// EnvelopeEncryptionDefaultContentKey: optional
	DefaultContentKey *EnvelopeEncryptionDefaultContentKey `hcl:"default_content_key,block"`
	// EnvelopeEncryptionEnabledProtocols: optional
	EnabledProtocols *EnvelopeEncryptionEnabledProtocols `hcl:"enabled_protocols,block"`
}

type EnvelopeEncryptionDefaultContentKey struct {
	// Label: string, optional
	Label terra.StringValue `hcl:"label,attr"`
	// PolicyName: string, optional
	PolicyName terra.StringValue `hcl:"policy_name,attr"`
}

type EnvelopeEncryptionEnabledProtocols struct {
	// Dash: bool, optional
	Dash terra.BoolValue `hcl:"dash,attr"`
	// Download: bool, optional
	Download terra.BoolValue `hcl:"download,attr"`
	// Hls: bool, optional
	Hls terra.BoolValue `hcl:"hls,attr"`
	// SmoothStreaming: bool, optional
	SmoothStreaming terra.BoolValue `hcl:"smooth_streaming,attr"`
}

type NoEncryptionEnabledProtocols struct {
	// Dash: bool, optional
	Dash terra.BoolValue `hcl:"dash,attr"`
	// Download: bool, optional
	Download terra.BoolValue `hcl:"download,attr"`
	// Hls: bool, optional
	Hls terra.BoolValue `hcl:"hls,attr"`
	// SmoothStreaming: bool, optional
	SmoothStreaming terra.BoolValue `hcl:"smooth_streaming,attr"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Read: string, optional
	Read terra.StringValue `hcl:"read,attr"`
}

type CommonEncryptionCbcsAttributes struct {
	ref terra.Reference
}

func (cec CommonEncryptionCbcsAttributes) InternalRef() (terra.Reference, error) {
	return cec.ref, nil
}

func (cec CommonEncryptionCbcsAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCbcsAttributes {
	return CommonEncryptionCbcsAttributes{ref: ref}
}

func (cec CommonEncryptionCbcsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cec.ref.InternalTokens()
}

func (cec CommonEncryptionCbcsAttributes) ClearKeyEncryption() terra.ListValue[CommonEncryptionCbcsClearKeyEncryptionAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCbcsClearKeyEncryptionAttributes](cec.ref.Append("clear_key_encryption"))
}

func (cec CommonEncryptionCbcsAttributes) DefaultContentKey() terra.ListValue[CommonEncryptionCbcsDefaultContentKeyAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCbcsDefaultContentKeyAttributes](cec.ref.Append("default_content_key"))
}

func (cec CommonEncryptionCbcsAttributes) DrmFairplay() terra.ListValue[CommonEncryptionCbcsDrmFairplayAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCbcsDrmFairplayAttributes](cec.ref.Append("drm_fairplay"))
}

func (cec CommonEncryptionCbcsAttributes) EnabledProtocols() terra.ListValue[CommonEncryptionCbcsEnabledProtocolsAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCbcsEnabledProtocolsAttributes](cec.ref.Append("enabled_protocols"))
}

type CommonEncryptionCbcsClearKeyEncryptionAttributes struct {
	ref terra.Reference
}

func (cke CommonEncryptionCbcsClearKeyEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return cke.ref, nil
}

func (cke CommonEncryptionCbcsClearKeyEncryptionAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCbcsClearKeyEncryptionAttributes {
	return CommonEncryptionCbcsClearKeyEncryptionAttributes{ref: ref}
}

func (cke CommonEncryptionCbcsClearKeyEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cke.ref.InternalTokens()
}

func (cke CommonEncryptionCbcsClearKeyEncryptionAttributes) CustomKeysAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(cke.ref.Append("custom_keys_acquisition_url_template"))
}

type CommonEncryptionCbcsDefaultContentKeyAttributes struct {
	ref terra.Reference
}

func (dck CommonEncryptionCbcsDefaultContentKeyAttributes) InternalRef() (terra.Reference, error) {
	return dck.ref, nil
}

func (dck CommonEncryptionCbcsDefaultContentKeyAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCbcsDefaultContentKeyAttributes {
	return CommonEncryptionCbcsDefaultContentKeyAttributes{ref: ref}
}

func (dck CommonEncryptionCbcsDefaultContentKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dck.ref.InternalTokens()
}

func (dck CommonEncryptionCbcsDefaultContentKeyAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("label"))
}

func (dck CommonEncryptionCbcsDefaultContentKeyAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("policy_name"))
}

type CommonEncryptionCbcsDrmFairplayAttributes struct {
	ref terra.Reference
}

func (df CommonEncryptionCbcsDrmFairplayAttributes) InternalRef() (terra.Reference, error) {
	return df.ref, nil
}

func (df CommonEncryptionCbcsDrmFairplayAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCbcsDrmFairplayAttributes {
	return CommonEncryptionCbcsDrmFairplayAttributes{ref: ref}
}

func (df CommonEncryptionCbcsDrmFairplayAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return df.ref.InternalTokens()
}

func (df CommonEncryptionCbcsDrmFairplayAttributes) AllowPersistentLicense() terra.BoolValue {
	return terra.ReferenceAsBool(df.ref.Append("allow_persistent_license"))
}

func (df CommonEncryptionCbcsDrmFairplayAttributes) CustomLicenseAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(df.ref.Append("custom_license_acquisition_url_template"))
}

type CommonEncryptionCbcsEnabledProtocolsAttributes struct {
	ref terra.Reference
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCbcsEnabledProtocolsAttributes {
	return CommonEncryptionCbcsEnabledProtocolsAttributes{ref: ref}
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) Dash() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("dash"))
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) Download() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("download"))
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) Hls() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("hls"))
}

func (ep CommonEncryptionCbcsEnabledProtocolsAttributes) SmoothStreaming() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("smooth_streaming"))
}

type CommonEncryptionCencAttributes struct {
	ref terra.Reference
}

func (cec CommonEncryptionCencAttributes) InternalRef() (terra.Reference, error) {
	return cec.ref, nil
}

func (cec CommonEncryptionCencAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencAttributes {
	return CommonEncryptionCencAttributes{ref: ref}
}

func (cec CommonEncryptionCencAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cec.ref.InternalTokens()
}

func (cec CommonEncryptionCencAttributes) DrmWidevineCustomLicenseAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(cec.ref.Append("drm_widevine_custom_license_acquisition_url_template"))
}

func (cec CommonEncryptionCencAttributes) ClearKeyEncryption() terra.ListValue[CommonEncryptionCencClearKeyEncryptionAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCencClearKeyEncryptionAttributes](cec.ref.Append("clear_key_encryption"))
}

func (cec CommonEncryptionCencAttributes) ClearTrack() terra.SetValue[CommonEncryptionCencClearTrackAttributes] {
	return terra.ReferenceAsSet[CommonEncryptionCencClearTrackAttributes](cec.ref.Append("clear_track"))
}

func (cec CommonEncryptionCencAttributes) ContentKeyToTrackMapping() terra.SetValue[CommonEncryptionCencContentKeyToTrackMappingAttributes] {
	return terra.ReferenceAsSet[CommonEncryptionCencContentKeyToTrackMappingAttributes](cec.ref.Append("content_key_to_track_mapping"))
}

func (cec CommonEncryptionCencAttributes) DefaultContentKey() terra.ListValue[CommonEncryptionCencDefaultContentKeyAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCencDefaultContentKeyAttributes](cec.ref.Append("default_content_key"))
}

func (cec CommonEncryptionCencAttributes) DrmPlayready() terra.ListValue[CommonEncryptionCencDrmPlayreadyAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCencDrmPlayreadyAttributes](cec.ref.Append("drm_playready"))
}

func (cec CommonEncryptionCencAttributes) EnabledProtocols() terra.ListValue[CommonEncryptionCencEnabledProtocolsAttributes] {
	return terra.ReferenceAsList[CommonEncryptionCencEnabledProtocolsAttributes](cec.ref.Append("enabled_protocols"))
}

type CommonEncryptionCencClearKeyEncryptionAttributes struct {
	ref terra.Reference
}

func (cke CommonEncryptionCencClearKeyEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return cke.ref, nil
}

func (cke CommonEncryptionCencClearKeyEncryptionAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencClearKeyEncryptionAttributes {
	return CommonEncryptionCencClearKeyEncryptionAttributes{ref: ref}
}

func (cke CommonEncryptionCencClearKeyEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cke.ref.InternalTokens()
}

func (cke CommonEncryptionCencClearKeyEncryptionAttributes) CustomKeysAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(cke.ref.Append("custom_keys_acquisition_url_template"))
}

type CommonEncryptionCencClearTrackAttributes struct {
	ref terra.Reference
}

func (ct CommonEncryptionCencClearTrackAttributes) InternalRef() (terra.Reference, error) {
	return ct.ref, nil
}

func (ct CommonEncryptionCencClearTrackAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencClearTrackAttributes {
	return CommonEncryptionCencClearTrackAttributes{ref: ref}
}

func (ct CommonEncryptionCencClearTrackAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ct.ref.InternalTokens()
}

func (ct CommonEncryptionCencClearTrackAttributes) Condition() terra.SetValue[CommonEncryptionCencClearTrackConditionAttributes] {
	return terra.ReferenceAsSet[CommonEncryptionCencClearTrackConditionAttributes](ct.ref.Append("condition"))
}

type CommonEncryptionCencClearTrackConditionAttributes struct {
	ref terra.Reference
}

func (c CommonEncryptionCencClearTrackConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CommonEncryptionCencClearTrackConditionAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencClearTrackConditionAttributes {
	return CommonEncryptionCencClearTrackConditionAttributes{ref: ref}
}

func (c CommonEncryptionCencClearTrackConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CommonEncryptionCencClearTrackConditionAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("operation"))
}

func (c CommonEncryptionCencClearTrackConditionAttributes) Property() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("property"))
}

func (c CommonEncryptionCencClearTrackConditionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type CommonEncryptionCencContentKeyToTrackMappingAttributes struct {
	ref terra.Reference
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) InternalRef() (terra.Reference, error) {
	return ckttm.ref, nil
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencContentKeyToTrackMappingAttributes {
	return CommonEncryptionCencContentKeyToTrackMappingAttributes{ref: ref}
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ckttm.ref.InternalTokens()
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(ckttm.ref.Append("label"))
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceAsString(ckttm.ref.Append("policy_name"))
}

func (ckttm CommonEncryptionCencContentKeyToTrackMappingAttributes) Track() terra.SetValue[CommonEncryptionCencContentKeyToTrackMappingTrackAttributes] {
	return terra.ReferenceAsSet[CommonEncryptionCencContentKeyToTrackMappingTrackAttributes](ckttm.ref.Append("track"))
}

type CommonEncryptionCencContentKeyToTrackMappingTrackAttributes struct {
	ref terra.Reference
}

func (t CommonEncryptionCencContentKeyToTrackMappingTrackAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t CommonEncryptionCencContentKeyToTrackMappingTrackAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencContentKeyToTrackMappingTrackAttributes {
	return CommonEncryptionCencContentKeyToTrackMappingTrackAttributes{ref: ref}
}

func (t CommonEncryptionCencContentKeyToTrackMappingTrackAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return t.ref.InternalTokens()
}

func (t CommonEncryptionCencContentKeyToTrackMappingTrackAttributes) Condition() terra.SetValue[CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes] {
	return terra.ReferenceAsSet[CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes](t.ref.Append("condition"))
}

type CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes struct {
	ref terra.Reference
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) InternalRef() (terra.Reference, error) {
	return c.ref, nil
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes {
	return CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes{ref: ref}
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return c.ref.InternalTokens()
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) Operation() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("operation"))
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) Property() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("property"))
}

func (c CommonEncryptionCencContentKeyToTrackMappingTrackConditionAttributes) Value() terra.StringValue {
	return terra.ReferenceAsString(c.ref.Append("value"))
}

type CommonEncryptionCencDefaultContentKeyAttributes struct {
	ref terra.Reference
}

func (dck CommonEncryptionCencDefaultContentKeyAttributes) InternalRef() (terra.Reference, error) {
	return dck.ref, nil
}

func (dck CommonEncryptionCencDefaultContentKeyAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencDefaultContentKeyAttributes {
	return CommonEncryptionCencDefaultContentKeyAttributes{ref: ref}
}

func (dck CommonEncryptionCencDefaultContentKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dck.ref.InternalTokens()
}

func (dck CommonEncryptionCencDefaultContentKeyAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("label"))
}

func (dck CommonEncryptionCencDefaultContentKeyAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("policy_name"))
}

type CommonEncryptionCencDrmPlayreadyAttributes struct {
	ref terra.Reference
}

func (dp CommonEncryptionCencDrmPlayreadyAttributes) InternalRef() (terra.Reference, error) {
	return dp.ref, nil
}

func (dp CommonEncryptionCencDrmPlayreadyAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencDrmPlayreadyAttributes {
	return CommonEncryptionCencDrmPlayreadyAttributes{ref: ref}
}

func (dp CommonEncryptionCencDrmPlayreadyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dp.ref.InternalTokens()
}

func (dp CommonEncryptionCencDrmPlayreadyAttributes) CustomAttributes() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("custom_attributes"))
}

func (dp CommonEncryptionCencDrmPlayreadyAttributes) CustomLicenseAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(dp.ref.Append("custom_license_acquisition_url_template"))
}

type CommonEncryptionCencEnabledProtocolsAttributes struct {
	ref terra.Reference
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) InternalWithRef(ref terra.Reference) CommonEncryptionCencEnabledProtocolsAttributes {
	return CommonEncryptionCencEnabledProtocolsAttributes{ref: ref}
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) Dash() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("dash"))
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) Download() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("download"))
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) Hls() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("hls"))
}

func (ep CommonEncryptionCencEnabledProtocolsAttributes) SmoothStreaming() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("smooth_streaming"))
}

type EnvelopeEncryptionAttributes struct {
	ref terra.Reference
}

func (ee EnvelopeEncryptionAttributes) InternalRef() (terra.Reference, error) {
	return ee.ref, nil
}

func (ee EnvelopeEncryptionAttributes) InternalWithRef(ref terra.Reference) EnvelopeEncryptionAttributes {
	return EnvelopeEncryptionAttributes{ref: ref}
}

func (ee EnvelopeEncryptionAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ee.ref.InternalTokens()
}

func (ee EnvelopeEncryptionAttributes) CustomKeysAcquisitionUrlTemplate() terra.StringValue {
	return terra.ReferenceAsString(ee.ref.Append("custom_keys_acquisition_url_template"))
}

func (ee EnvelopeEncryptionAttributes) DefaultContentKey() terra.ListValue[EnvelopeEncryptionDefaultContentKeyAttributes] {
	return terra.ReferenceAsList[EnvelopeEncryptionDefaultContentKeyAttributes](ee.ref.Append("default_content_key"))
}

func (ee EnvelopeEncryptionAttributes) EnabledProtocols() terra.ListValue[EnvelopeEncryptionEnabledProtocolsAttributes] {
	return terra.ReferenceAsList[EnvelopeEncryptionEnabledProtocolsAttributes](ee.ref.Append("enabled_protocols"))
}

type EnvelopeEncryptionDefaultContentKeyAttributes struct {
	ref terra.Reference
}

func (dck EnvelopeEncryptionDefaultContentKeyAttributes) InternalRef() (terra.Reference, error) {
	return dck.ref, nil
}

func (dck EnvelopeEncryptionDefaultContentKeyAttributes) InternalWithRef(ref terra.Reference) EnvelopeEncryptionDefaultContentKeyAttributes {
	return EnvelopeEncryptionDefaultContentKeyAttributes{ref: ref}
}

func (dck EnvelopeEncryptionDefaultContentKeyAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return dck.ref.InternalTokens()
}

func (dck EnvelopeEncryptionDefaultContentKeyAttributes) Label() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("label"))
}

func (dck EnvelopeEncryptionDefaultContentKeyAttributes) PolicyName() terra.StringValue {
	return terra.ReferenceAsString(dck.ref.Append("policy_name"))
}

type EnvelopeEncryptionEnabledProtocolsAttributes struct {
	ref terra.Reference
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) InternalRef() (terra.Reference, error) {
	return ep.ref, nil
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) InternalWithRef(ref terra.Reference) EnvelopeEncryptionEnabledProtocolsAttributes {
	return EnvelopeEncryptionEnabledProtocolsAttributes{ref: ref}
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ep.ref.InternalTokens()
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) Dash() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("dash"))
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) Download() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("download"))
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) Hls() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("hls"))
}

func (ep EnvelopeEncryptionEnabledProtocolsAttributes) SmoothStreaming() terra.BoolValue {
	return terra.ReferenceAsBool(ep.ref.Append("smooth_streaming"))
}

type NoEncryptionEnabledProtocolsAttributes struct {
	ref terra.Reference
}

func (neep NoEncryptionEnabledProtocolsAttributes) InternalRef() (terra.Reference, error) {
	return neep.ref, nil
}

func (neep NoEncryptionEnabledProtocolsAttributes) InternalWithRef(ref terra.Reference) NoEncryptionEnabledProtocolsAttributes {
	return NoEncryptionEnabledProtocolsAttributes{ref: ref}
}

func (neep NoEncryptionEnabledProtocolsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return neep.ref.InternalTokens()
}

func (neep NoEncryptionEnabledProtocolsAttributes) Dash() terra.BoolValue {
	return terra.ReferenceAsBool(neep.ref.Append("dash"))
}

func (neep NoEncryptionEnabledProtocolsAttributes) Download() terra.BoolValue {
	return terra.ReferenceAsBool(neep.ref.Append("download"))
}

func (neep NoEncryptionEnabledProtocolsAttributes) Hls() terra.BoolValue {
	return terra.ReferenceAsBool(neep.ref.Append("hls"))
}

func (neep NoEncryptionEnabledProtocolsAttributes) SmoothStreaming() terra.BoolValue {
	return terra.ReferenceAsBool(neep.ref.Append("smooth_streaming"))
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

func (t TimeoutsAttributes) Read() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("read"))
}

type CommonEncryptionCbcsState struct {
	ClearKeyEncryption []CommonEncryptionCbcsClearKeyEncryptionState `json:"clear_key_encryption"`
	DefaultContentKey  []CommonEncryptionCbcsDefaultContentKeyState  `json:"default_content_key"`
	DrmFairplay        []CommonEncryptionCbcsDrmFairplayState        `json:"drm_fairplay"`
	EnabledProtocols   []CommonEncryptionCbcsEnabledProtocolsState   `json:"enabled_protocols"`
}

type CommonEncryptionCbcsClearKeyEncryptionState struct {
	CustomKeysAcquisitionUrlTemplate string `json:"custom_keys_acquisition_url_template"`
}

type CommonEncryptionCbcsDefaultContentKeyState struct {
	Label      string `json:"label"`
	PolicyName string `json:"policy_name"`
}

type CommonEncryptionCbcsDrmFairplayState struct {
	AllowPersistentLicense              bool   `json:"allow_persistent_license"`
	CustomLicenseAcquisitionUrlTemplate string `json:"custom_license_acquisition_url_template"`
}

type CommonEncryptionCbcsEnabledProtocolsState struct {
	Dash            bool `json:"dash"`
	Download        bool `json:"download"`
	Hls             bool `json:"hls"`
	SmoothStreaming bool `json:"smooth_streaming"`
}

type CommonEncryptionCencState struct {
	DrmWidevineCustomLicenseAcquisitionUrlTemplate string                                              `json:"drm_widevine_custom_license_acquisition_url_template"`
	ClearKeyEncryption                             []CommonEncryptionCencClearKeyEncryptionState       `json:"clear_key_encryption"`
	ClearTrack                                     []CommonEncryptionCencClearTrackState               `json:"clear_track"`
	ContentKeyToTrackMapping                       []CommonEncryptionCencContentKeyToTrackMappingState `json:"content_key_to_track_mapping"`
	DefaultContentKey                              []CommonEncryptionCencDefaultContentKeyState        `json:"default_content_key"`
	DrmPlayready                                   []CommonEncryptionCencDrmPlayreadyState             `json:"drm_playready"`
	EnabledProtocols                               []CommonEncryptionCencEnabledProtocolsState         `json:"enabled_protocols"`
}

type CommonEncryptionCencClearKeyEncryptionState struct {
	CustomKeysAcquisitionUrlTemplate string `json:"custom_keys_acquisition_url_template"`
}

type CommonEncryptionCencClearTrackState struct {
	Condition []CommonEncryptionCencClearTrackConditionState `json:"condition"`
}

type CommonEncryptionCencClearTrackConditionState struct {
	Operation string `json:"operation"`
	Property  string `json:"property"`
	Value     string `json:"value"`
}

type CommonEncryptionCencContentKeyToTrackMappingState struct {
	Label      string                                                   `json:"label"`
	PolicyName string                                                   `json:"policy_name"`
	Track      []CommonEncryptionCencContentKeyToTrackMappingTrackState `json:"track"`
}

type CommonEncryptionCencContentKeyToTrackMappingTrackState struct {
	Condition []CommonEncryptionCencContentKeyToTrackMappingTrackConditionState `json:"condition"`
}

type CommonEncryptionCencContentKeyToTrackMappingTrackConditionState struct {
	Operation string `json:"operation"`
	Property  string `json:"property"`
	Value     string `json:"value"`
}

type CommonEncryptionCencDefaultContentKeyState struct {
	Label      string `json:"label"`
	PolicyName string `json:"policy_name"`
}

type CommonEncryptionCencDrmPlayreadyState struct {
	CustomAttributes                    string `json:"custom_attributes"`
	CustomLicenseAcquisitionUrlTemplate string `json:"custom_license_acquisition_url_template"`
}

type CommonEncryptionCencEnabledProtocolsState struct {
	Dash            bool `json:"dash"`
	Download        bool `json:"download"`
	Hls             bool `json:"hls"`
	SmoothStreaming bool `json:"smooth_streaming"`
}

type EnvelopeEncryptionState struct {
	CustomKeysAcquisitionUrlTemplate string                                     `json:"custom_keys_acquisition_url_template"`
	DefaultContentKey                []EnvelopeEncryptionDefaultContentKeyState `json:"default_content_key"`
	EnabledProtocols                 []EnvelopeEncryptionEnabledProtocolsState  `json:"enabled_protocols"`
}

type EnvelopeEncryptionDefaultContentKeyState struct {
	Label      string `json:"label"`
	PolicyName string `json:"policy_name"`
}

type EnvelopeEncryptionEnabledProtocolsState struct {
	Dash            bool `json:"dash"`
	Download        bool `json:"download"`
	Hls             bool `json:"hls"`
	SmoothStreaming bool `json:"smooth_streaming"`
}

type NoEncryptionEnabledProtocolsState struct {
	Dash            bool `json:"dash"`
	Download        bool `json:"download"`
	Hls             bool `json:"hls"`
	SmoothStreaming bool `json:"smooth_streaming"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Read   string `json:"read"`
}
