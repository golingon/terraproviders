// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	s3objectcopy "github.com/golingon/terraproviders/aws/5.9.0/s3objectcopy"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewS3ObjectCopy creates a new instance of [S3ObjectCopy].
func NewS3ObjectCopy(name string, args S3ObjectCopyArgs) *S3ObjectCopy {
	return &S3ObjectCopy{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*S3ObjectCopy)(nil)

// S3ObjectCopy represents the Terraform resource aws_s3_object_copy.
type S3ObjectCopy struct {
	Name      string
	Args      S3ObjectCopyArgs
	state     *s3ObjectCopyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [S3ObjectCopy].
func (soc *S3ObjectCopy) Type() string {
	return "aws_s3_object_copy"
}

// LocalName returns the local name for [S3ObjectCopy].
func (soc *S3ObjectCopy) LocalName() string {
	return soc.Name
}

// Configuration returns the configuration (args) for [S3ObjectCopy].
func (soc *S3ObjectCopy) Configuration() interface{} {
	return soc.Args
}

// DependOn is used for other resources to depend on [S3ObjectCopy].
func (soc *S3ObjectCopy) DependOn() terra.Reference {
	return terra.ReferenceResource(soc)
}

// Dependencies returns the list of resources [S3ObjectCopy] depends_on.
func (soc *S3ObjectCopy) Dependencies() terra.Dependencies {
	return soc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [S3ObjectCopy].
func (soc *S3ObjectCopy) LifecycleManagement() *terra.Lifecycle {
	return soc.Lifecycle
}

// Attributes returns the attributes for [S3ObjectCopy].
func (soc *S3ObjectCopy) Attributes() s3ObjectCopyAttributes {
	return s3ObjectCopyAttributes{ref: terra.ReferenceResource(soc)}
}

// ImportState imports the given attribute values into [S3ObjectCopy]'s state.
func (soc *S3ObjectCopy) ImportState(av io.Reader) error {
	soc.state = &s3ObjectCopyState{}
	if err := json.NewDecoder(av).Decode(soc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", soc.Type(), soc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [S3ObjectCopy] has state.
func (soc *S3ObjectCopy) State() (*s3ObjectCopyState, bool) {
	return soc.state, soc.state != nil
}

// StateMust returns the state for [S3ObjectCopy]. Panics if the state is nil.
func (soc *S3ObjectCopy) StateMust() *s3ObjectCopyState {
	if soc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", soc.Type(), soc.LocalName()))
	}
	return soc.state
}

// S3ObjectCopyArgs contains the configurations for aws_s3_object_copy.
type S3ObjectCopyArgs struct {
	// Acl: string, optional
	Acl terra.StringValue `hcl:"acl,attr"`
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// BucketKeyEnabled: bool, optional
	BucketKeyEnabled terra.BoolValue `hcl:"bucket_key_enabled,attr"`
	// CacheControl: string, optional
	CacheControl terra.StringValue `hcl:"cache_control,attr"`
	// ContentDisposition: string, optional
	ContentDisposition terra.StringValue `hcl:"content_disposition,attr"`
	// ContentEncoding: string, optional
	ContentEncoding terra.StringValue `hcl:"content_encoding,attr"`
	// ContentLanguage: string, optional
	ContentLanguage terra.StringValue `hcl:"content_language,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// CopyIfMatch: string, optional
	CopyIfMatch terra.StringValue `hcl:"copy_if_match,attr"`
	// CopyIfModifiedSince: string, optional
	CopyIfModifiedSince terra.StringValue `hcl:"copy_if_modified_since,attr"`
	// CopyIfNoneMatch: string, optional
	CopyIfNoneMatch terra.StringValue `hcl:"copy_if_none_match,attr"`
	// CopyIfUnmodifiedSince: string, optional
	CopyIfUnmodifiedSince terra.StringValue `hcl:"copy_if_unmodified_since,attr"`
	// CustomerAlgorithm: string, optional
	CustomerAlgorithm terra.StringValue `hcl:"customer_algorithm,attr"`
	// CustomerKey: string, optional
	CustomerKey terra.StringValue `hcl:"customer_key,attr"`
	// CustomerKeyMd5: string, optional
	CustomerKeyMd5 terra.StringValue `hcl:"customer_key_md5,attr"`
	// ExpectedBucketOwner: string, optional
	ExpectedBucketOwner terra.StringValue `hcl:"expected_bucket_owner,attr"`
	// ExpectedSourceBucketOwner: string, optional
	ExpectedSourceBucketOwner terra.StringValue `hcl:"expected_source_bucket_owner,attr"`
	// Expires: string, optional
	Expires terra.StringValue `hcl:"expires,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// KmsEncryptionContext: string, optional
	KmsEncryptionContext terra.StringValue `hcl:"kms_encryption_context,attr"`
	// KmsKeyId: string, optional
	KmsKeyId terra.StringValue `hcl:"kms_key_id,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// MetadataDirective: string, optional
	MetadataDirective terra.StringValue `hcl:"metadata_directive,attr"`
	// ObjectLockLegalHoldStatus: string, optional
	ObjectLockLegalHoldStatus terra.StringValue `hcl:"object_lock_legal_hold_status,attr"`
	// ObjectLockMode: string, optional
	ObjectLockMode terra.StringValue `hcl:"object_lock_mode,attr"`
	// ObjectLockRetainUntilDate: string, optional
	ObjectLockRetainUntilDate terra.StringValue `hcl:"object_lock_retain_until_date,attr"`
	// RequestPayer: string, optional
	RequestPayer terra.StringValue `hcl:"request_payer,attr"`
	// ServerSideEncryption: string, optional
	ServerSideEncryption terra.StringValue `hcl:"server_side_encryption,attr"`
	// Source: string, required
	Source terra.StringValue `hcl:"source,attr" validate:"required"`
	// SourceCustomerAlgorithm: string, optional
	SourceCustomerAlgorithm terra.StringValue `hcl:"source_customer_algorithm,attr"`
	// SourceCustomerKey: string, optional
	SourceCustomerKey terra.StringValue `hcl:"source_customer_key,attr"`
	// SourceCustomerKeyMd5: string, optional
	SourceCustomerKeyMd5 terra.StringValue `hcl:"source_customer_key_md5,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
	// TaggingDirective: string, optional
	TaggingDirective terra.StringValue `hcl:"tagging_directive,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// WebsiteRedirect: string, optional
	WebsiteRedirect terra.StringValue `hcl:"website_redirect,attr"`
	// Grant: min=0
	Grant []s3objectcopy.Grant `hcl:"grant,block" validate:"min=0"`
}
type s3ObjectCopyAttributes struct {
	ref terra.Reference
}

// Acl returns a reference to field acl of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Acl() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("acl"))
}

// Bucket returns a reference to field bucket of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("bucket"))
}

// BucketKeyEnabled returns a reference to field bucket_key_enabled of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(soc.ref.Append("bucket_key_enabled"))
}

// CacheControl returns a reference to field cache_control of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("cache_control"))
}

// ContentDisposition returns a reference to field content_disposition of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("content_type"))
}

// CopyIfMatch returns a reference to field copy_if_match of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CopyIfMatch() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("copy_if_match"))
}

// CopyIfModifiedSince returns a reference to field copy_if_modified_since of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CopyIfModifiedSince() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("copy_if_modified_since"))
}

// CopyIfNoneMatch returns a reference to field copy_if_none_match of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CopyIfNoneMatch() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("copy_if_none_match"))
}

// CopyIfUnmodifiedSince returns a reference to field copy_if_unmodified_since of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CopyIfUnmodifiedSince() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("copy_if_unmodified_since"))
}

// CustomerAlgorithm returns a reference to field customer_algorithm of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CustomerAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("customer_algorithm"))
}

// CustomerKey returns a reference to field customer_key of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CustomerKey() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("customer_key"))
}

// CustomerKeyMd5 returns a reference to field customer_key_md5 of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) CustomerKeyMd5() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("customer_key_md5"))
}

// Etag returns a reference to field etag of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("etag"))
}

// ExpectedBucketOwner returns a reference to field expected_bucket_owner of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ExpectedBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("expected_bucket_owner"))
}

// ExpectedSourceBucketOwner returns a reference to field expected_source_bucket_owner of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ExpectedSourceBucketOwner() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("expected_source_bucket_owner"))
}

// Expiration returns a reference to field expiration of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("expiration"))
}

// Expires returns a reference to field expires of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("expires"))
}

// ForceDestroy returns a reference to field force_destroy of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(soc.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("id"))
}

// Key returns a reference to field key of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("key"))
}

// KmsEncryptionContext returns a reference to field kms_encryption_context of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) KmsEncryptionContext() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("kms_encryption_context"))
}

// KmsKeyId returns a reference to field kms_key_id of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) KmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("kms_key_id"))
}

// LastModified returns a reference to field last_modified of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("last_modified"))
}

// Metadata returns a reference to field metadata of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](soc.ref.Append("metadata"))
}

// MetadataDirective returns a reference to field metadata_directive of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) MetadataDirective() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("metadata_directive"))
}

// ObjectLockLegalHoldStatus returns a reference to field object_lock_legal_hold_status of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("object_lock_legal_hold_status"))
}

// ObjectLockMode returns a reference to field object_lock_mode of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("object_lock_mode"))
}

// ObjectLockRetainUntilDate returns a reference to field object_lock_retain_until_date of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("object_lock_retain_until_date"))
}

// RequestCharged returns a reference to field request_charged of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) RequestCharged() terra.BoolValue {
	return terra.ReferenceAsBool(soc.ref.Append("request_charged"))
}

// RequestPayer returns a reference to field request_payer of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) RequestPayer() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("request_payer"))
}

// ServerSideEncryption returns a reference to field server_side_encryption of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("server_side_encryption"))
}

// Source returns a reference to field source of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("source"))
}

// SourceCustomerAlgorithm returns a reference to field source_customer_algorithm of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) SourceCustomerAlgorithm() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("source_customer_algorithm"))
}

// SourceCustomerKey returns a reference to field source_customer_key of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) SourceCustomerKey() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("source_customer_key"))
}

// SourceCustomerKeyMd5 returns a reference to field source_customer_key_md5 of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) SourceCustomerKeyMd5() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("source_customer_key_md5"))
}

// SourceVersionId returns a reference to field source_version_id of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) SourceVersionId() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("source_version_id"))
}

// StorageClass returns a reference to field storage_class of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("storage_class"))
}

// TaggingDirective returns a reference to field tagging_directive of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) TaggingDirective() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("tagging_directive"))
}

// Tags returns a reference to field tags of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](soc.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](soc.ref.Append("tags_all"))
}

// VersionId returns a reference to field version_id of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("version_id"))
}

// WebsiteRedirect returns a reference to field website_redirect of aws_s3_object_copy.
func (soc s3ObjectCopyAttributes) WebsiteRedirect() terra.StringValue {
	return terra.ReferenceAsString(soc.ref.Append("website_redirect"))
}

func (soc s3ObjectCopyAttributes) Grant() terra.SetValue[s3objectcopy.GrantAttributes] {
	return terra.ReferenceAsSet[s3objectcopy.GrantAttributes](soc.ref.Append("grant"))
}

type s3ObjectCopyState struct {
	Acl                       string                    `json:"acl"`
	Bucket                    string                    `json:"bucket"`
	BucketKeyEnabled          bool                      `json:"bucket_key_enabled"`
	CacheControl              string                    `json:"cache_control"`
	ContentDisposition        string                    `json:"content_disposition"`
	ContentEncoding           string                    `json:"content_encoding"`
	ContentLanguage           string                    `json:"content_language"`
	ContentType               string                    `json:"content_type"`
	CopyIfMatch               string                    `json:"copy_if_match"`
	CopyIfModifiedSince       string                    `json:"copy_if_modified_since"`
	CopyIfNoneMatch           string                    `json:"copy_if_none_match"`
	CopyIfUnmodifiedSince     string                    `json:"copy_if_unmodified_since"`
	CustomerAlgorithm         string                    `json:"customer_algorithm"`
	CustomerKey               string                    `json:"customer_key"`
	CustomerKeyMd5            string                    `json:"customer_key_md5"`
	Etag                      string                    `json:"etag"`
	ExpectedBucketOwner       string                    `json:"expected_bucket_owner"`
	ExpectedSourceBucketOwner string                    `json:"expected_source_bucket_owner"`
	Expiration                string                    `json:"expiration"`
	Expires                   string                    `json:"expires"`
	ForceDestroy              bool                      `json:"force_destroy"`
	Id                        string                    `json:"id"`
	Key                       string                    `json:"key"`
	KmsEncryptionContext      string                    `json:"kms_encryption_context"`
	KmsKeyId                  string                    `json:"kms_key_id"`
	LastModified              string                    `json:"last_modified"`
	Metadata                  map[string]string         `json:"metadata"`
	MetadataDirective         string                    `json:"metadata_directive"`
	ObjectLockLegalHoldStatus string                    `json:"object_lock_legal_hold_status"`
	ObjectLockMode            string                    `json:"object_lock_mode"`
	ObjectLockRetainUntilDate string                    `json:"object_lock_retain_until_date"`
	RequestCharged            bool                      `json:"request_charged"`
	RequestPayer              string                    `json:"request_payer"`
	ServerSideEncryption      string                    `json:"server_side_encryption"`
	Source                    string                    `json:"source"`
	SourceCustomerAlgorithm   string                    `json:"source_customer_algorithm"`
	SourceCustomerKey         string                    `json:"source_customer_key"`
	SourceCustomerKeyMd5      string                    `json:"source_customer_key_md5"`
	SourceVersionId           string                    `json:"source_version_id"`
	StorageClass              string                    `json:"storage_class"`
	TaggingDirective          string                    `json:"tagging_directive"`
	Tags                      map[string]string         `json:"tags"`
	TagsAll                   map[string]string         `json:"tags_all"`
	VersionId                 string                    `json:"version_id"`
	WebsiteRedirect           string                    `json:"website_redirect"`
	Grant                     []s3objectcopy.GrantState `json:"grant"`
}
