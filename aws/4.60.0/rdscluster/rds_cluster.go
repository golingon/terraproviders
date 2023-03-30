// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package rdscluster

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type RestoreToPointInTime struct {
	// RestoreToTime: string, optional
	RestoreToTime terra.StringValue `hcl:"restore_to_time,attr"`
	// RestoreType: string, optional
	RestoreType terra.StringValue `hcl:"restore_type,attr"`
	// SourceClusterIdentifier: string, required
	SourceClusterIdentifier terra.StringValue `hcl:"source_cluster_identifier,attr" validate:"required"`
	// UseLatestRestorableTime: bool, optional
	UseLatestRestorableTime terra.BoolValue `hcl:"use_latest_restorable_time,attr"`
}

type S3Import struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// BucketPrefix: string, optional
	BucketPrefix terra.StringValue `hcl:"bucket_prefix,attr"`
	// IngestionRole: string, required
	IngestionRole terra.StringValue `hcl:"ingestion_role,attr" validate:"required"`
	// SourceEngine: string, required
	SourceEngine terra.StringValue `hcl:"source_engine,attr" validate:"required"`
	// SourceEngineVersion: string, required
	SourceEngineVersion terra.StringValue `hcl:"source_engine_version,attr" validate:"required"`
}

type ScalingConfiguration struct {
	// AutoPause: bool, optional
	AutoPause terra.BoolValue `hcl:"auto_pause,attr"`
	// MaxCapacity: number, optional
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr"`
	// MinCapacity: number, optional
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr"`
	// SecondsUntilAutoPause: number, optional
	SecondsUntilAutoPause terra.NumberValue `hcl:"seconds_until_auto_pause,attr"`
	// TimeoutAction: string, optional
	TimeoutAction terra.StringValue `hcl:"timeout_action,attr"`
}

type Serverlessv2ScalingConfiguration struct {
	// MaxCapacity: number, required
	MaxCapacity terra.NumberValue `hcl:"max_capacity,attr" validate:"required"`
	// MinCapacity: number, required
	MinCapacity terra.NumberValue `hcl:"min_capacity,attr" validate:"required"`
}

type Timeouts struct {
	// Create: string, optional
	Create terra.StringValue `hcl:"create,attr"`
	// Delete: string, optional
	Delete terra.StringValue `hcl:"delete,attr"`
	// Update: string, optional
	Update terra.StringValue `hcl:"update,attr"`
}

type RestoreToPointInTimeAttributes struct {
	ref terra.Reference
}

func (rtpit RestoreToPointInTimeAttributes) InternalRef() terra.Reference {
	return rtpit.ref
}

func (rtpit RestoreToPointInTimeAttributes) InternalWithRef(ref terra.Reference) RestoreToPointInTimeAttributes {
	return RestoreToPointInTimeAttributes{ref: ref}
}

func (rtpit RestoreToPointInTimeAttributes) InternalTokens() hclwrite.Tokens {
	return rtpit.ref.InternalTokens()
}

func (rtpit RestoreToPointInTimeAttributes) RestoreToTime() terra.StringValue {
	return terra.ReferenceString(rtpit.ref.Append("restore_to_time"))
}

func (rtpit RestoreToPointInTimeAttributes) RestoreType() terra.StringValue {
	return terra.ReferenceString(rtpit.ref.Append("restore_type"))
}

func (rtpit RestoreToPointInTimeAttributes) SourceClusterIdentifier() terra.StringValue {
	return terra.ReferenceString(rtpit.ref.Append("source_cluster_identifier"))
}

func (rtpit RestoreToPointInTimeAttributes) UseLatestRestorableTime() terra.BoolValue {
	return terra.ReferenceBool(rtpit.ref.Append("use_latest_restorable_time"))
}

type S3ImportAttributes struct {
	ref terra.Reference
}

func (si S3ImportAttributes) InternalRef() terra.Reference {
	return si.ref
}

func (si S3ImportAttributes) InternalWithRef(ref terra.Reference) S3ImportAttributes {
	return S3ImportAttributes{ref: ref}
}

func (si S3ImportAttributes) InternalTokens() hclwrite.Tokens {
	return si.ref.InternalTokens()
}

func (si S3ImportAttributes) BucketName() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("bucket_name"))
}

func (si S3ImportAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("bucket_prefix"))
}

func (si S3ImportAttributes) IngestionRole() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("ingestion_role"))
}

func (si S3ImportAttributes) SourceEngine() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("source_engine"))
}

func (si S3ImportAttributes) SourceEngineVersion() terra.StringValue {
	return terra.ReferenceString(si.ref.Append("source_engine_version"))
}

type ScalingConfigurationAttributes struct {
	ref terra.Reference
}

func (sc ScalingConfigurationAttributes) InternalRef() terra.Reference {
	return sc.ref
}

func (sc ScalingConfigurationAttributes) InternalWithRef(ref terra.Reference) ScalingConfigurationAttributes {
	return ScalingConfigurationAttributes{ref: ref}
}

func (sc ScalingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return sc.ref.InternalTokens()
}

func (sc ScalingConfigurationAttributes) AutoPause() terra.BoolValue {
	return terra.ReferenceBool(sc.ref.Append("auto_pause"))
}

func (sc ScalingConfigurationAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("max_capacity"))
}

func (sc ScalingConfigurationAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("min_capacity"))
}

func (sc ScalingConfigurationAttributes) SecondsUntilAutoPause() terra.NumberValue {
	return terra.ReferenceNumber(sc.ref.Append("seconds_until_auto_pause"))
}

func (sc ScalingConfigurationAttributes) TimeoutAction() terra.StringValue {
	return terra.ReferenceString(sc.ref.Append("timeout_action"))
}

type Serverlessv2ScalingConfigurationAttributes struct {
	ref terra.Reference
}

func (ssc Serverlessv2ScalingConfigurationAttributes) InternalRef() terra.Reference {
	return ssc.ref
}

func (ssc Serverlessv2ScalingConfigurationAttributes) InternalWithRef(ref terra.Reference) Serverlessv2ScalingConfigurationAttributes {
	return Serverlessv2ScalingConfigurationAttributes{ref: ref}
}

func (ssc Serverlessv2ScalingConfigurationAttributes) InternalTokens() hclwrite.Tokens {
	return ssc.ref.InternalTokens()
}

func (ssc Serverlessv2ScalingConfigurationAttributes) MaxCapacity() terra.NumberValue {
	return terra.ReferenceNumber(ssc.ref.Append("max_capacity"))
}

func (ssc Serverlessv2ScalingConfigurationAttributes) MinCapacity() terra.NumberValue {
	return terra.ReferenceNumber(ssc.ref.Append("min_capacity"))
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

func (t TimeoutsAttributes) Update() terra.StringValue {
	return terra.ReferenceString(t.ref.Append("update"))
}

type RestoreToPointInTimeState struct {
	RestoreToTime           string `json:"restore_to_time"`
	RestoreType             string `json:"restore_type"`
	SourceClusterIdentifier string `json:"source_cluster_identifier"`
	UseLatestRestorableTime bool   `json:"use_latest_restorable_time"`
}

type S3ImportState struct {
	BucketName          string `json:"bucket_name"`
	BucketPrefix        string `json:"bucket_prefix"`
	IngestionRole       string `json:"ingestion_role"`
	SourceEngine        string `json:"source_engine"`
	SourceEngineVersion string `json:"source_engine_version"`
}

type ScalingConfigurationState struct {
	AutoPause             bool    `json:"auto_pause"`
	MaxCapacity           float64 `json:"max_capacity"`
	MinCapacity           float64 `json:"min_capacity"`
	SecondsUntilAutoPause float64 `json:"seconds_until_auto_pause"`
	TimeoutAction         string  `json:"timeout_action"`
}

type Serverlessv2ScalingConfigurationState struct {
	MaxCapacity float64 `json:"max_capacity"`
	MinCapacity float64 `json:"min_capacity"`
}

type TimeoutsState struct {
	Create string `json:"create"`
	Delete string `json:"delete"`
	Update string `json:"update"`
}
