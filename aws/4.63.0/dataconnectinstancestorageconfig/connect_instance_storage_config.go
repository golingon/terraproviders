// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package dataconnectinstancestorageconfig

import (
	hclwrite "github.com/hashicorp/hcl/v2/hclwrite"
	terra "github.com/volvo-cars/lingon/pkg/terra"
)

type StorageConfig struct {
	// KinesisFirehoseConfig: min=0
	KinesisFirehoseConfig []KinesisFirehoseConfig `hcl:"kinesis_firehose_config,block" validate:"min=0"`
	// KinesisStreamConfig: min=0
	KinesisStreamConfig []KinesisStreamConfig `hcl:"kinesis_stream_config,block" validate:"min=0"`
	// KinesisVideoStreamConfig: min=0
	KinesisVideoStreamConfig []KinesisVideoStreamConfig `hcl:"kinesis_video_stream_config,block" validate:"min=0"`
	// S3Config: min=0
	S3Config []S3Config `hcl:"s3_config,block" validate:"min=0"`
}

type KinesisFirehoseConfig struct{}

type KinesisStreamConfig struct{}

type KinesisVideoStreamConfig struct {
	// KinesisVideoStreamConfigEncryptionConfig: min=0
	EncryptionConfig []KinesisVideoStreamConfigEncryptionConfig `hcl:"encryption_config,block" validate:"min=0"`
}

type KinesisVideoStreamConfigEncryptionConfig struct{}

type S3Config struct {
	// S3ConfigEncryptionConfig: min=0
	EncryptionConfig []S3ConfigEncryptionConfig `hcl:"encryption_config,block" validate:"min=0"`
}

type S3ConfigEncryptionConfig struct{}

type StorageConfigAttributes struct {
	ref terra.Reference
}

func (sc StorageConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc StorageConfigAttributes) InternalWithRef(ref terra.Reference) StorageConfigAttributes {
	return StorageConfigAttributes{ref: ref}
}

func (sc StorageConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc StorageConfigAttributes) StorageType() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("storage_type"))
}

func (sc StorageConfigAttributes) KinesisFirehoseConfig() terra.ListValue[KinesisFirehoseConfigAttributes] {
	return terra.ReferenceAsList[KinesisFirehoseConfigAttributes](sc.ref.Append("kinesis_firehose_config"))
}

func (sc StorageConfigAttributes) KinesisStreamConfig() terra.ListValue[KinesisStreamConfigAttributes] {
	return terra.ReferenceAsList[KinesisStreamConfigAttributes](sc.ref.Append("kinesis_stream_config"))
}

func (sc StorageConfigAttributes) KinesisVideoStreamConfig() terra.ListValue[KinesisVideoStreamConfigAttributes] {
	return terra.ReferenceAsList[KinesisVideoStreamConfigAttributes](sc.ref.Append("kinesis_video_stream_config"))
}

func (sc StorageConfigAttributes) S3Config() terra.ListValue[S3ConfigAttributes] {
	return terra.ReferenceAsList[S3ConfigAttributes](sc.ref.Append("s3_config"))
}

type KinesisFirehoseConfigAttributes struct {
	ref terra.Reference
}

func (kfc KinesisFirehoseConfigAttributes) InternalRef() (terra.Reference, error) {
	return kfc.ref, nil
}

func (kfc KinesisFirehoseConfigAttributes) InternalWithRef(ref terra.Reference) KinesisFirehoseConfigAttributes {
	return KinesisFirehoseConfigAttributes{ref: ref}
}

func (kfc KinesisFirehoseConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kfc.ref.InternalTokens()
}

func (kfc KinesisFirehoseConfigAttributes) FirehoseArn() terra.StringValue {
	return terra.ReferenceAsString(kfc.ref.Append("firehose_arn"))
}

type KinesisStreamConfigAttributes struct {
	ref terra.Reference
}

func (ksc KinesisStreamConfigAttributes) InternalRef() (terra.Reference, error) {
	return ksc.ref, nil
}

func (ksc KinesisStreamConfigAttributes) InternalWithRef(ref terra.Reference) KinesisStreamConfigAttributes {
	return KinesisStreamConfigAttributes{ref: ref}
}

func (ksc KinesisStreamConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ksc.ref.InternalTokens()
}

func (ksc KinesisStreamConfigAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(ksc.ref.Append("stream_arn"))
}

type KinesisVideoStreamConfigAttributes struct {
	ref terra.Reference
}

func (kvsc KinesisVideoStreamConfigAttributes) InternalRef() (terra.Reference, error) {
	return kvsc.ref, nil
}

func (kvsc KinesisVideoStreamConfigAttributes) InternalWithRef(ref terra.Reference) KinesisVideoStreamConfigAttributes {
	return KinesisVideoStreamConfigAttributes{ref: ref}
}

func (kvsc KinesisVideoStreamConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return kvsc.ref.InternalTokens()
}

func (kvsc KinesisVideoStreamConfigAttributes) Prefix() terra.StringValue {
	return terra.ReferenceAsString(kvsc.ref.Append("prefix"))
}

func (kvsc KinesisVideoStreamConfigAttributes) RetentionPeriodHours() terra.NumberValue {
	return terra.ReferenceAsNumber(kvsc.ref.Append("retention_period_hours"))
}

func (kvsc KinesisVideoStreamConfigAttributes) EncryptionConfig() terra.ListValue[KinesisVideoStreamConfigEncryptionConfigAttributes] {
	return terra.ReferenceAsList[KinesisVideoStreamConfigEncryptionConfigAttributes](kvsc.ref.Append("encryption_config"))
}

type KinesisVideoStreamConfigEncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec KinesisVideoStreamConfigEncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec KinesisVideoStreamConfigEncryptionConfigAttributes) InternalWithRef(ref terra.Reference) KinesisVideoStreamConfigEncryptionConfigAttributes {
	return KinesisVideoStreamConfigEncryptionConfigAttributes{ref: ref}
}

func (ec KinesisVideoStreamConfigEncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec KinesisVideoStreamConfigEncryptionConfigAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_type"))
}

func (ec KinesisVideoStreamConfigEncryptionConfigAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("key_id"))
}

type S3ConfigAttributes struct {
	ref terra.Reference
}

func (sc S3ConfigAttributes) InternalRef() (terra.Reference, error) {
	return sc.ref, nil
}

func (sc S3ConfigAttributes) InternalWithRef(ref terra.Reference) S3ConfigAttributes {
	return S3ConfigAttributes{ref: ref}
}

func (sc S3ConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return sc.ref.InternalTokens()
}

func (sc S3ConfigAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("bucket_name"))
}

func (sc S3ConfigAttributes) BucketPrefix() terra.StringValue {
	return terra.ReferenceAsString(sc.ref.Append("bucket_prefix"))
}

func (sc S3ConfigAttributes) EncryptionConfig() terra.ListValue[S3ConfigEncryptionConfigAttributes] {
	return terra.ReferenceAsList[S3ConfigEncryptionConfigAttributes](sc.ref.Append("encryption_config"))
}

type S3ConfigEncryptionConfigAttributes struct {
	ref terra.Reference
}

func (ec S3ConfigEncryptionConfigAttributes) InternalRef() (terra.Reference, error) {
	return ec.ref, nil
}

func (ec S3ConfigEncryptionConfigAttributes) InternalWithRef(ref terra.Reference) S3ConfigEncryptionConfigAttributes {
	return S3ConfigEncryptionConfigAttributes{ref: ref}
}

func (ec S3ConfigEncryptionConfigAttributes) InternalTokens() (hclwrite.Tokens, error) {
	return ec.ref.InternalTokens()
}

func (ec S3ConfigEncryptionConfigAttributes) EncryptionType() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("encryption_type"))
}

func (ec S3ConfigEncryptionConfigAttributes) KeyId() terra.StringValue {
	return terra.ReferenceAsString(ec.ref.Append("key_id"))
}

type StorageConfigState struct {
	StorageType              string                          `json:"storage_type"`
	KinesisFirehoseConfig    []KinesisFirehoseConfigState    `json:"kinesis_firehose_config"`
	KinesisStreamConfig      []KinesisStreamConfigState      `json:"kinesis_stream_config"`
	KinesisVideoStreamConfig []KinesisVideoStreamConfigState `json:"kinesis_video_stream_config"`
	S3Config                 []S3ConfigState                 `json:"s3_config"`
}

type KinesisFirehoseConfigState struct {
	FirehoseArn string `json:"firehose_arn"`
}

type KinesisStreamConfigState struct {
	StreamArn string `json:"stream_arn"`
}

type KinesisVideoStreamConfigState struct {
	Prefix               string                                          `json:"prefix"`
	RetentionPeriodHours float64                                         `json:"retention_period_hours"`
	EncryptionConfig     []KinesisVideoStreamConfigEncryptionConfigState `json:"encryption_config"`
}

type KinesisVideoStreamConfigEncryptionConfigState struct {
	EncryptionType string `json:"encryption_type"`
	KeyId          string `json:"key_id"`
}

type S3ConfigState struct {
	BucketName       string                          `json:"bucket_name"`
	BucketPrefix     string                          `json:"bucket_prefix"`
	EncryptionConfig []S3ConfigEncryptionConfigState `json:"encryption_config"`
}

type S3ConfigEncryptionConfigState struct {
	EncryptionType string `json:"encryption_type"`
	KeyId          string `json:"key_id"`
}
