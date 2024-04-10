// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package opsworkshaproxylayer

import (
	terra "github.com/golingon/lingon/pkg/terra"
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
)

type CloudwatchConfiguration struct {
	// Enabled: bool, optional
	Enabled terra.BoolValue `hcl:"enabled,attr"`
	// LogStreams: min=0
	LogStreams []LogStreams `hcl:"log_streams,block" validate:"min=0"`
}

type LogStreams struct {
	// BatchCount: number, optional
	BatchCount terra.NumberValue `hcl:"batch_count,attr"`
	// BatchSize: number, optional
	BatchSize terra.NumberValue `hcl:"batch_size,attr"`
	// BufferDuration: number, optional
	BufferDuration terra.NumberValue `hcl:"buffer_duration,attr"`
	// DatetimeFormat: string, optional
	DatetimeFormat terra.StringValue `hcl:"datetime_format,attr"`
	// Encoding: string, optional
	Encoding terra.StringValue `hcl:"encoding,attr"`
	// File: string, required
	File terra.StringValue `hcl:"file,attr" validate:"required"`
	// FileFingerprintLines: string, optional
	FileFingerprintLines terra.StringValue `hcl:"file_fingerprint_lines,attr"`
	// InitialPosition: string, optional
	InitialPosition terra.StringValue `hcl:"initial_position,attr"`
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
	// MultilineStartPattern: string, optional
	MultilineStartPattern terra.StringValue `hcl:"multiline_start_pattern,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
}

type EbsVolume struct {
	// Encrypted: bool, optional
	Encrypted terra.BoolValue `hcl:"encrypted,attr"`
	// Iops: number, optional
	Iops terra.NumberValue `hcl:"iops,attr"`
	// MountPoint: string, required
	MountPoint terra.StringValue `hcl:"mount_point,attr" validate:"required"`
	// NumberOfDisks: number, required
	NumberOfDisks terra.NumberValue `hcl:"number_of_disks,attr" validate:"required"`
	// RaidLevel: string, optional
	RaidLevel terra.StringValue `hcl:"raid_level,attr"`
	// Size: number, required
	Size terra.NumberValue `hcl:"size,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
}

type LoadBasedAutoScaling struct {
	// Enable: bool, optional
	Enable terra.BoolValue `hcl:"enable,attr"`
	// Downscaling: optional
	Downscaling *Downscaling `hcl:"downscaling,block"`
	// Upscaling: optional
	Upscaling *Upscaling `hcl:"upscaling,block"`
}

type Downscaling struct {
	// Alarms: list of string, optional
	Alarms terra.ListValue[terra.StringValue] `hcl:"alarms,attr"`
	// CpuThreshold: number, optional
	CpuThreshold terra.NumberValue `hcl:"cpu_threshold,attr"`
	// IgnoreMetricsTime: number, optional
	IgnoreMetricsTime terra.NumberValue `hcl:"ignore_metrics_time,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// LoadThreshold: number, optional
	LoadThreshold terra.NumberValue `hcl:"load_threshold,attr"`
	// MemoryThreshold: number, optional
	MemoryThreshold terra.NumberValue `hcl:"memory_threshold,attr"`
	// ThresholdsWaitTime: number, optional
	ThresholdsWaitTime terra.NumberValue `hcl:"thresholds_wait_time,attr"`
}

type Upscaling struct {
	// Alarms: list of string, optional
	Alarms terra.ListValue[terra.StringValue] `hcl:"alarms,attr"`
	// CpuThreshold: number, optional
	CpuThreshold terra.NumberValue `hcl:"cpu_threshold,attr"`
	// IgnoreMetricsTime: number, optional
	IgnoreMetricsTime terra.NumberValue `hcl:"ignore_metrics_time,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// LoadThreshold: number, optional
	LoadThreshold terra.NumberValue `hcl:"load_threshold,attr"`
	// MemoryThreshold: number, optional
	MemoryThreshold terra.NumberValue `hcl:"memory_threshold,attr"`
	// ThresholdsWaitTime: number, optional
	ThresholdsWaitTime terra.NumberValue `hcl:"thresholds_wait_time,attr"`
}

type CloudwatchConfigurationAttributes struct {
	ref terra.Reference
}

func (cc CloudwatchConfigurationAttributes) InternalRef() (terra.Reference, error) {
	return cc.ref, nil
}

func (cc CloudwatchConfigurationAttributes) InternalWithRef(ref terra.Reference) CloudwatchConfigurationAttributes {
	return CloudwatchConfigurationAttributes{ref: ref}
}

func (cc CloudwatchConfigurationAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return cc.ref.InternalTokens()
}

func (cc CloudwatchConfigurationAttributes) Enabled() terra.BoolValue {
	return terra.ReferenceAsBool(cc.ref.Append("enabled"))
}

func (cc CloudwatchConfigurationAttributes) LogStreams() terra.ListValue[LogStreamsAttributes] {
	return terra.ReferenceAsList[LogStreamsAttributes](cc.ref.Append("log_streams"))
}

type LogStreamsAttributes struct {
	ref terra.Reference
}

func (ls LogStreamsAttributes) InternalRef() (terra.Reference, error) {
	return ls.ref, nil
}

func (ls LogStreamsAttributes) InternalWithRef(ref terra.Reference) LogStreamsAttributes {
	return LogStreamsAttributes{ref: ref}
}

func (ls LogStreamsAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ls.ref.InternalTokens()
}

func (ls LogStreamsAttributes) BatchCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ls.ref.Append("batch_count"))
}

func (ls LogStreamsAttributes) BatchSize() terra.NumberValue {
	return terra.ReferenceAsNumber(ls.ref.Append("batch_size"))
}

func (ls LogStreamsAttributes) BufferDuration() terra.NumberValue {
	return terra.ReferenceAsNumber(ls.ref.Append("buffer_duration"))
}

func (ls LogStreamsAttributes) DatetimeFormat() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("datetime_format"))
}

func (ls LogStreamsAttributes) Encoding() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("encoding"))
}

func (ls LogStreamsAttributes) File() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("file"))
}

func (ls LogStreamsAttributes) FileFingerprintLines() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("file_fingerprint_lines"))
}

func (ls LogStreamsAttributes) InitialPosition() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("initial_position"))
}

func (ls LogStreamsAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("log_group_name"))
}

func (ls LogStreamsAttributes) MultilineStartPattern() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("multiline_start_pattern"))
}

func (ls LogStreamsAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(ls.ref.Append("time_zone"))
}

type EbsVolumeAttributes struct {
	ref terra.Reference
}

func (ev EbsVolumeAttributes) InternalRef() (terra.Reference, error) {
	return ev.ref, nil
}

func (ev EbsVolumeAttributes) InternalWithRef(ref terra.Reference) EbsVolumeAttributes {
	return EbsVolumeAttributes{ref: ref}
}

func (ev EbsVolumeAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ev.ref.InternalTokens()
}

func (ev EbsVolumeAttributes) Encrypted() terra.BoolValue {
	return terra.ReferenceAsBool(ev.ref.Append("encrypted"))
}

func (ev EbsVolumeAttributes) Iops() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("iops"))
}

func (ev EbsVolumeAttributes) MountPoint() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("mount_point"))
}

func (ev EbsVolumeAttributes) NumberOfDisks() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("number_of_disks"))
}

func (ev EbsVolumeAttributes) RaidLevel() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("raid_level"))
}

func (ev EbsVolumeAttributes) Size() terra.NumberValue {
	return terra.ReferenceAsNumber(ev.ref.Append("size"))
}

func (ev EbsVolumeAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(ev.ref.Append("type"))
}

type LoadBasedAutoScalingAttributes struct {
	ref terra.Reference
}

func (lbas LoadBasedAutoScalingAttributes) InternalRef() (terra.Reference, error) {
	return lbas.ref, nil
}

func (lbas LoadBasedAutoScalingAttributes) InternalWithRef(ref terra.Reference) LoadBasedAutoScalingAttributes {
	return LoadBasedAutoScalingAttributes{ref: ref}
}

func (lbas LoadBasedAutoScalingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return lbas.ref.InternalTokens()
}

func (lbas LoadBasedAutoScalingAttributes) Enable() terra.BoolValue {
	return terra.ReferenceAsBool(lbas.ref.Append("enable"))
}

func (lbas LoadBasedAutoScalingAttributes) Downscaling() terra.ListValue[DownscalingAttributes] {
	return terra.ReferenceAsList[DownscalingAttributes](lbas.ref.Append("downscaling"))
}

func (lbas LoadBasedAutoScalingAttributes) Upscaling() terra.ListValue[UpscalingAttributes] {
	return terra.ReferenceAsList[UpscalingAttributes](lbas.ref.Append("upscaling"))
}

type DownscalingAttributes struct {
	ref terra.Reference
}

func (d DownscalingAttributes) InternalRef() (terra.Reference, error) {
	return d.ref, nil
}

func (d DownscalingAttributes) InternalWithRef(ref terra.Reference) DownscalingAttributes {
	return DownscalingAttributes{ref: ref}
}

func (d DownscalingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return d.ref.InternalTokens()
}

func (d DownscalingAttributes) Alarms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](d.ref.Append("alarms"))
}

func (d DownscalingAttributes) CpuThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("cpu_threshold"))
}

func (d DownscalingAttributes) IgnoreMetricsTime() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("ignore_metrics_time"))
}

func (d DownscalingAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("instance_count"))
}

func (d DownscalingAttributes) LoadThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("load_threshold"))
}

func (d DownscalingAttributes) MemoryThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("memory_threshold"))
}

func (d DownscalingAttributes) ThresholdsWaitTime() terra.NumberValue {
	return terra.ReferenceAsNumber(d.ref.Append("thresholds_wait_time"))
}

type UpscalingAttributes struct {
	ref terra.Reference
}

func (u UpscalingAttributes) InternalRef() (terra.Reference, error) {
	return u.ref, nil
}

func (u UpscalingAttributes) InternalWithRef(ref terra.Reference) UpscalingAttributes {
	return UpscalingAttributes{ref: ref}
}

func (u UpscalingAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return u.ref.InternalTokens()
}

func (u UpscalingAttributes) Alarms() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](u.ref.Append("alarms"))
}

func (u UpscalingAttributes) CpuThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("cpu_threshold"))
}

func (u UpscalingAttributes) IgnoreMetricsTime() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("ignore_metrics_time"))
}

func (u UpscalingAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("instance_count"))
}

func (u UpscalingAttributes) LoadThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("load_threshold"))
}

func (u UpscalingAttributes) MemoryThreshold() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("memory_threshold"))
}

func (u UpscalingAttributes) ThresholdsWaitTime() terra.NumberValue {
	return terra.ReferenceAsNumber(u.ref.Append("thresholds_wait_time"))
}

type CloudwatchConfigurationState struct {
	Enabled    bool              `json:"enabled"`
	LogStreams []LogStreamsState `json:"log_streams"`
}

type LogStreamsState struct {
	BatchCount            float64 `json:"batch_count"`
	BatchSize             float64 `json:"batch_size"`
	BufferDuration        float64 `json:"buffer_duration"`
	DatetimeFormat        string  `json:"datetime_format"`
	Encoding              string  `json:"encoding"`
	File                  string  `json:"file"`
	FileFingerprintLines  string  `json:"file_fingerprint_lines"`
	InitialPosition       string  `json:"initial_position"`
	LogGroupName          string  `json:"log_group_name"`
	MultilineStartPattern string  `json:"multiline_start_pattern"`
	TimeZone              string  `json:"time_zone"`
}

type EbsVolumeState struct {
	Encrypted     bool    `json:"encrypted"`
	Iops          float64 `json:"iops"`
	MountPoint    string  `json:"mount_point"`
	NumberOfDisks float64 `json:"number_of_disks"`
	RaidLevel     string  `json:"raid_level"`
	Size          float64 `json:"size"`
	Type          string  `json:"type"`
}

type LoadBasedAutoScalingState struct {
	Enable      bool               `json:"enable"`
	Downscaling []DownscalingState `json:"downscaling"`
	Upscaling   []UpscalingState   `json:"upscaling"`
}

type DownscalingState struct {
	Alarms             []string `json:"alarms"`
	CpuThreshold       float64  `json:"cpu_threshold"`
	IgnoreMetricsTime  float64  `json:"ignore_metrics_time"`
	InstanceCount      float64  `json:"instance_count"`
	LoadThreshold      float64  `json:"load_threshold"`
	MemoryThreshold    float64  `json:"memory_threshold"`
	ThresholdsWaitTime float64  `json:"thresholds_wait_time"`
}

type UpscalingState struct {
	Alarms             []string `json:"alarms"`
	CpuThreshold       float64  `json:"cpu_threshold"`
	IgnoreMetricsTime  float64  `json:"ignore_metrics_time"`
	InstanceCount      float64  `json:"instance_count"`
	LoadThreshold      float64  `json:"load_threshold"`
	MemoryThreshold    float64  `json:"memory_threshold"`
	ThresholdsWaitTime float64  `json:"thresholds_wait_time"`
}
