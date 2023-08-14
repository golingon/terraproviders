// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3Object creates a new instance of [S3Object].
func NewS3Object(name string, args S3ObjectArgs) *S3Object {
	return &S3Object{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3Object)(nil)

// S3Object represents the Terraform resource aws_s3_object.
type S3Object struct {
	Name      string
	Args      S3ObjectArgs
	state     *s3ObjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3Object].
func (so *S3Object) Type() string {
	return "aws_s3_object"
}

// LocalName returns the local name for [S3Object].
func (so *S3Object) LocalName() string {
	return so.Name
}

// Configuration returns the configuration (args) for [S3Object].
func (so *S3Object) Configuration() interface{} {
	return so.Args
}

// DependOn is used for other resources to depend on [S3Object].
func (so *S3Object) DependOn() terra.Reference {
	return terra.ReferenceResource(so)
}

// Dependencies returns the list of resources [S3Object] depends_on.
func (so *S3Object) Dependencies() terra.Dependencies {
	return so.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3Object].
func (so *S3Object) LifecycleManagement() *terra.Lifecycle {
	return so.Lifecycle
}

// Attributes returns the attributes for [S3Object].
func (so *S3Object) Attributes() s3ObjectAttributes {
	return s3ObjectAttributes{ref: terra.ReferenceResource(so)}
}

// ImportState imports the given attribute values into [S3Object]'s state.
func (so *S3Object) ImportState(av io.Reader) error {
	so.state = &s3ObjectState{}
	if err := json.NewDecoder(av).Decode(so.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", so.Type(), so.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3Object] has state.
func (so *S3Object) State() (*s3ObjectState, bool) {
	return so.state, so.state != nil
}

// StateMust returns the state for [S3Object]. Panics if the state is nil.
func (so *S3Object) StateMust() *s3ObjectState {
	if so.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", so.Type(), so.LocalName()))
	}
	return so.state
}

// S3ObjectArgs contains the configurations for aws_s3_object.
type S3ObjectArgs struct {
	// Acl: string, optional
	Acl terra.StringValue `hcl:"acl,attr"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// BucketKeyEnabled: bool, optional
	BucketKeyEnabled terra.BoolValue `hcl:"bucket_key_enabled,attr"`
	// CacheControl: string, optional
	CacheControl terra.StringValue `hcl:"cache_control,attr"`
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// ContentBase64: string, optional
	ContentBase64 terra.StringValue `hcl:"content_base64,attr"`
	// ContentDisposition: string, optional
	ContentDisposition terra.StringValue `hcl:"content_disposition,attr"`
	// ContentEncoding: string, optional
	ContentEncoding terra.StringValue `hcl:"content_encoding,attr"`
	// ContentLanguage: string, optional
	ContentLanguage terra.StringValue `hcl:"content_language,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Etag: string, optional
	Etag terra.StringValue `hcl:"etag,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// ObjectLockLegalHoldStatus: string, optional
	ObjectLockLegalHoldStatus terra.StringValue `hcl:"object_lock_legal_hold_status,attr"`
	// ObjectLockMode: string, optional
	ObjectLockMode terra.StringValue `hcl:"object_lock_mode,attr"`
	// ObjectLockRetainUntilDate: string, optional
	ObjectLockRetainUntilDate terra.StringValue `hcl:"object_lock_retain_until_date,attr"`
	// ServerSideEncryption: string, optional
	ServerSideEncryption terra.StringValue `hcl:"server_side_encryption,attr"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
	// SourceHash: string, optional
	SourceHash terra.StringValue `hcl:"source_hash,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WebsiteRedirect: string, optional
	WebsiteRedirect terra.StringValue `hcl:"website_redirect,attr"`
}
type s3ObjectAttributes struct {
	ref terra.Reference
}

// Acl returns a reference to field acl of aws_s3_object.
func (so s3ObjectAttributes) Acl() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("acl"))
}

// Bucket returns a reference to field bucket of aws_s3_object.
func (so s3ObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("bucket"))
}

// BucketKeyEnabled returns a reference to field bucket_key_enabled of aws_s3_object.
func (so s3ObjectAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(so.ref.Append("bucket_key_enabled"))
}

// CacheControl returns a reference to field cache_control of aws_s3_object.
func (so s3ObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("cache_control"))
}

// Content returns a reference to field content of aws_s3_object.
func (so s3ObjectAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content"))
}

// ContentBase64 returns a reference to field content_base64 of aws_s3_object.
func (so s3ObjectAttributes) ContentBase64() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_base64"))
}

// ContentDisposition returns a reference to field content_disposition of aws_s3_object.
func (so s3ObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of aws_s3_object.
func (so s3ObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of aws_s3_object.
func (so s3ObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of aws_s3_object.
func (so s3ObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_type"))
}

// Etag returns a reference to field etag of aws_s3_object.
func (so s3ObjectAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("etag"))
}

// ForceDestroy returns a reference to field force_destroy of aws_s3_object.
func (so s3ObjectAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(so.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_s3_object.
func (so s3ObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("id"))
}

// Key returns a reference to field key of aws_s3_object.
func (so s3ObjectAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("key"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_s3_object.
func (so s3ObjectAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("kms_key_id"))
}

// Metadata returns a reference to field metadata of aws_s3_object.
func (so s3ObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](so.ref.Append("metadata"))
}

// ObjectLockLegalHoldStatus returns a reference to field object_lock_legal_hold_status of aws_s3_object.
func (so s3ObjectAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_legal_hold_status"))
}

// ObjectLockMode returns a reference to field object_lock_mode of aws_s3_object.
func (so s3ObjectAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_mode"))
}

// ObjectLockRetainUntilDate returns a reference to field object_lock_retain_until_date of aws_s3_object.
func (so s3ObjectAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_retain_until_date"))
}

// ServerSideEncryption returns a reference to field server_side_encryption of aws_s3_object.
func (so s3ObjectAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("server_side_encryption"))
}

// Source returns a reference to field source of aws_s3_object.
func (so s3ObjectAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("source"))
}

// SourceHash returns a reference to field source_hash of aws_s3_object.
func (so s3ObjectAttributes) SourceHash() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("source_hash"))
}

// StorageClass returns a reference to field storage_class of aws_s3_object.
func (so s3ObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("storage_class"))
}

// Tags returns a reference to field tags of aws_s3_object.
func (so s3ObjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](so.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_s3_object.
func (so s3ObjectAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](so.ref.Append("tags_all"))
}

// VersionId returns a reference to field version_id of aws_s3_object.
func (so s3ObjectAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("version_id"))
}

// WebsiteRedirect returns a reference to field website_redirect of aws_s3_object.
func (so s3ObjectAttributes) WebsiteRedirect() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("website_redirect"))
}

type s3ObjectState struct {
	Acl                       string            `json:"acl"`
	Bucket                    string            `json:"bucket"`
	BucketKeyEnabled          bool              `json:"bucket_key_enabled"`
	CacheControl              string            `json:"cache_control"`
	Content                   string            `json:"content"`
	ContentBase64             string            `json:"content_base64"`
	ContentDisposition        string            `json:"content_disposition"`
	ContentEncoding           string            `json:"content_encoding"`
	ContentLanguage           string            `json:"content_language"`
	ContentType               string            `json:"content_type"`
	Etag                      string            `json:"etag"`
	ForceDestroy              bool              `json:"force_destroy"`
	Id                        string            `json:"id"`
	Key                       string            `json:"key"`
	KmsKeyId                  string            `json:"kms_key_id"`
	Metadata                  map[string]string `json:"metadata"`
	ObjectLockLegalHoldStatus string            `json:"object_lock_legal_hold_status"`
	ObjectLockMode            string            `json:"object_lock_mode"`
	ObjectLockRetainUntilDate string            `json:"object_lock_retain_until_date"`
	ServerSideEncryption      string            `json:"server_side_encryption"`
	Source                    string            `json:"source"`
	SourceHash                string            `json:"source_hash"`
	StorageClass              string            `json:"storage_class"`
	Tags                      map[string]string `json:"tags"`
	TagsAll                   map[string]string `json:"tags_all"`
	VersionId                 string            `json:"version_id"`
	WebsiteRedirect           string            `json:"website_redirect"`
}