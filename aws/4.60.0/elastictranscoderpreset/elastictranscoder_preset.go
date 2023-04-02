// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package elastictranscoderpreset

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type Audio struct {
	// AudioPackingMode: string, optional
	AudioPackingMode terra.StringValue `hcl:"audio_packing_mode,attr"`
	// BitRate: string, optional
	BitRate terra.StringValue `hcl:"bit_rate,attr"`
	// Channels: string, optional
	Channels terra.StringValue `hcl:"channels,attr"`
	// Codec: string, optional
	Codec terra.StringValue `hcl:"codec,attr"`
	// SampleRate: string, optional
	SampleRate terra.StringValue `hcl:"sample_rate,attr"`
}

type AudioCodecOptions struct {
	// BitDepth: string, optional
	BitDepth terra.StringValue `hcl:"bit_depth,attr"`
	// BitOrder: string, optional
	BitOrder terra.StringValue `hcl:"bit_order,attr"`
	// Profile: string, optional
	Profile terra.StringValue `hcl:"profile,attr"`
	// Signed: string, optional
	Signed terra.StringValue `hcl:"signed,attr"`
}

type Thumbnails struct {
	// AspectRatio: string, optional
	AspectRatio terra.StringValue `hcl:"aspect_ratio,attr"`
	// Format: string, optional
	Format terra.StringValue `hcl:"format,attr"`
	// Interval: string, optional
	Interval terra.StringValue `hcl:"interval,attr"`
	// MaxHeight: string, optional
	MaxHeight terra.StringValue `hcl:"max_height,attr"`
	// MaxWidth: string, optional
	MaxWidth terra.StringValue `hcl:"max_width,attr"`
	// PaddingPolicy: string, optional
	PaddingPolicy terra.StringValue `hcl:"padding_policy,attr"`
	// Resolution: string, optional
	Resolution terra.StringValue `hcl:"resolution,attr"`
	// SizingPolicy: string, optional
	SizingPolicy terra.StringValue `hcl:"sizing_policy,attr"`
}

type Video struct {
	// AspectRatio: string, optional
	AspectRatio terra.StringValue `hcl:"aspect_ratio,attr"`
	// BitRate: string, optional
	BitRate terra.StringValue `hcl:"bit_rate,attr"`
	// Codec: string, optional
	Codec terra.StringValue `hcl:"codec,attr"`
	// DisplayAspectRatio: string, optional
	DisplayAspectRatio terra.StringValue `hcl:"display_aspect_ratio,attr"`
	// FixedGop: string, optional
	FixedGop terra.StringValue `hcl:"fixed_gop,attr"`
	// FrameRate: string, optional
	FrameRate terra.StringValue `hcl:"frame_rate,attr"`
	// KeyframesMaxDist: string, optional
	KeyframesMaxDist terra.StringValue `hcl:"keyframes_max_dist,attr"`
	// MaxFrameRate: string, optional
	MaxFrameRate terra.StringValue `hcl:"max_frame_rate,attr"`
	// MaxHeight: string, optional
	MaxHeight terra.StringValue `hcl:"max_height,attr"`
	// MaxWidth: string, optional
	MaxWidth terra.StringValue `hcl:"max_width,attr"`
	// PaddingPolicy: string, optional
	PaddingPolicy terra.StringValue `hcl:"padding_policy,attr"`
	// Resolution: string, optional
	Resolution terra.StringValue `hcl:"resolution,attr"`
	// SizingPolicy: string, optional
	SizingPolicy terra.StringValue `hcl:"sizing_policy,attr"`
}

type VideoWatermarks struct {
	// HorizontalAlign: string, optional
	HorizontalAlign terra.StringValue `hcl:"horizontal_align,attr"`
	// HorizontalOffset: string, optional
	HorizontalOffset terra.StringValue `hcl:"horizontal_offset,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaxHeight: string, optional
	MaxHeight terra.StringValue `hcl:"max_height,attr"`
	// MaxWidth: string, optional
	MaxWidth terra.StringValue `hcl:"max_width,attr"`
	// Opacity: string, optional
	Opacity terra.StringValue `hcl:"opacity,attr"`
	// SizingPolicy: string, optional
	SizingPolicy terra.StringValue `hcl:"sizing_policy,attr"`
	// Target: string, optional
	Target terra.StringValue `hcl:"target,attr"`
	// VerticalAlign: string, optional
	VerticalAlign terra.StringValue `hcl:"vertical_align,attr"`
	// VerticalOffset: string, optional
	VerticalOffset terra.StringValue `hcl:"vertical_offset,attr"`
}

type AudioAttributes struct {
	ref terra.Reference
}

func (a AudioAttributes) InternalRef() (terra.Reference, error) {
	return a.ref, nil
}

func (a AudioAttributes) InternalWithRef(ref terra.Reference) AudioAttributes {
	return AudioAttributes{ref: ref}
}

func (a AudioAttributes) InternalTokens() hclwrite.Tokens {
	return a.ref.InternalTokens()
}

func (a AudioAttributes) AudioPackingMode() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("audio_packing_mode"))
}

func (a AudioAttributes) BitRate() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("bit_rate"))
}

func (a AudioAttributes) Channels() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("channels"))
}

func (a AudioAttributes) Codec() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("codec"))
}

func (a AudioAttributes) SampleRate() terra.StringValue {
	return terra.ReferenceAsString(a.ref.Append("sample_rate"))
}

type AudioCodecOptionsAttributes struct {
	ref terra.Reference
}

func (aco AudioCodecOptionsAttributes) InternalRef() (terra.Reference, error) {
	return aco.ref, nil
}

func (aco AudioCodecOptionsAttributes) InternalWithRef(ref terra.Reference) AudioCodecOptionsAttributes {
	return AudioCodecOptionsAttributes{ref: ref}
}

func (aco AudioCodecOptionsAttributes) InternalTokens() hclwrite.Tokens {
	return aco.ref.InternalTokens()
}

func (aco AudioCodecOptionsAttributes) BitDepth() terra.StringValue {
	return terra.ReferenceAsString(aco.ref.Append("bit_depth"))
}

func (aco AudioCodecOptionsAttributes) BitOrder() terra.StringValue {
	return terra.ReferenceAsString(aco.ref.Append("bit_order"))
}

func (aco AudioCodecOptionsAttributes) Profile() terra.StringValue {
	return terra.ReferenceAsString(aco.ref.Append("profile"))
}

func (aco AudioCodecOptionsAttributes) Signed() terra.StringValue {
	return terra.ReferenceAsString(aco.ref.Append("signed"))
}

type ThumbnailsAttributes struct {
	ref terra.Reference
}

func (t ThumbnailsAttributes) InternalRef() (terra.Reference, error) {
	return t.ref, nil
}

func (t ThumbnailsAttributes) InternalWithRef(ref terra.Reference) ThumbnailsAttributes {
	return ThumbnailsAttributes{ref: ref}
}

func (t ThumbnailsAttributes) InternalTokens() hclwrite.Tokens {
	return t.ref.InternalTokens()
}

func (t ThumbnailsAttributes) AspectRatio() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("aspect_ratio"))
}

func (t ThumbnailsAttributes) Format() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("format"))
}

func (t ThumbnailsAttributes) Interval() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("interval"))
}

func (t ThumbnailsAttributes) MaxHeight() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("max_height"))
}

func (t ThumbnailsAttributes) MaxWidth() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("max_width"))
}

func (t ThumbnailsAttributes) PaddingPolicy() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("padding_policy"))
}

func (t ThumbnailsAttributes) Resolution() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("resolution"))
}

func (t ThumbnailsAttributes) SizingPolicy() terra.StringValue {
	return terra.ReferenceAsString(t.ref.Append("sizing_policy"))
}

type VideoAttributes struct {
	ref terra.Reference
}

func (v VideoAttributes) InternalRef() (terra.Reference, error) {
	return v.ref, nil
}

func (v VideoAttributes) InternalWithRef(ref terra.Reference) VideoAttributes {
	return VideoAttributes{ref: ref}
}

func (v VideoAttributes) InternalTokens() hclwrite.Tokens {
	return v.ref.InternalTokens()
}

func (v VideoAttributes) AspectRatio() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("aspect_ratio"))
}

func (v VideoAttributes) BitRate() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("bit_rate"))
}

func (v VideoAttributes) Codec() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("codec"))
}

func (v VideoAttributes) DisplayAspectRatio() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("display_aspect_ratio"))
}

func (v VideoAttributes) FixedGop() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("fixed_gop"))
}

func (v VideoAttributes) FrameRate() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("frame_rate"))
}

func (v VideoAttributes) KeyframesMaxDist() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("keyframes_max_dist"))
}

func (v VideoAttributes) MaxFrameRate() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("max_frame_rate"))
}

func (v VideoAttributes) MaxHeight() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("max_height"))
}

func (v VideoAttributes) MaxWidth() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("max_width"))
}

func (v VideoAttributes) PaddingPolicy() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("padding_policy"))
}

func (v VideoAttributes) Resolution() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("resolution"))
}

func (v VideoAttributes) SizingPolicy() terra.StringValue {
	return terra.ReferenceAsString(v.ref.Append("sizing_policy"))
}

type VideoWatermarksAttributes struct {
	ref terra.Reference
}

func (vw VideoWatermarksAttributes) InternalRef() (terra.Reference, error) {
	return vw.ref, nil
}

func (vw VideoWatermarksAttributes) InternalWithRef(ref terra.Reference) VideoWatermarksAttributes {
	return VideoWatermarksAttributes{ref: ref}
}

func (vw VideoWatermarksAttributes) InternalTokens() hclwrite.Tokens {
	return vw.ref.InternalTokens()
}

func (vw VideoWatermarksAttributes) HorizontalAlign() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("horizontal_align"))
}

func (vw VideoWatermarksAttributes) HorizontalOffset() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("horizontal_offset"))
}

func (vw VideoWatermarksAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("id"))
}

func (vw VideoWatermarksAttributes) MaxHeight() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("max_height"))
}

func (vw VideoWatermarksAttributes) MaxWidth() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("max_width"))
}

func (vw VideoWatermarksAttributes) Opacity() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("opacity"))
}

func (vw VideoWatermarksAttributes) SizingPolicy() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("sizing_policy"))
}

func (vw VideoWatermarksAttributes) Target() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("target"))
}

func (vw VideoWatermarksAttributes) VerticalAlign() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("vertical_align"))
}

func (vw VideoWatermarksAttributes) VerticalOffset() terra.StringValue {
	return terra.ReferenceAsString(vw.ref.Append("vertical_offset"))
}

type AudioState struct {
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
	SampleRate       string `json:"sample_rate"`
}

type AudioCodecOptionsState struct {
	BitDepth string `json:"bit_depth"`
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
}

type ThumbnailsState struct {
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
}

type VideoState struct {
	AspectRatio        string `json:"aspect_ratio"`
	BitRate            string `json:"bit_rate"`
	Codec              string `json:"codec"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FixedGop           string `json:"fixed_gop"`
	FrameRate          string `json:"frame_rate"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxFrameRate       string `json:"max_frame_rate"`
	MaxHeight          string `json:"max_height"`
	MaxWidth           string `json:"max_width"`
	PaddingPolicy      string `json:"padding_policy"`
	Resolution         string `json:"resolution"`
	SizingPolicy       string `json:"sizing_policy"`
}

type VideoWatermarksState struct {
	HorizontalAlign  string `json:"horizontal_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	Id               string `json:"id"`
	MaxHeight        string `json:"max_height"`
	MaxWidth         string `json:"max_width"`
	Opacity          string `json:"opacity"`
	SizingPolicy     string `json:"sizing_policy"`
	Target           string `json:"target"`
	VerticalAlign    string `json:"vertical_align"`
	VerticalOffset   string `json:"vertical_offset"`
}
