// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_s3_bucket_object

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource aws_s3_bucket_object.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asbo *DataSource) DataSource() string {
	return "aws_s3_bucket_object"
}

// LocalName returns the local name for [DataSource].
func (asbo *DataSource) LocalName() string {
	return asbo.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asbo *DataSource) Configuration() interface{} {
	return asbo.Args
}

// Attributes returns the attributes for [DataSource].
func (asbo *DataSource) Attributes() dataAwsS3BucketObjectAttributes {
	return dataAwsS3BucketObjectAttributes{ref: terra.ReferenceDataSource(asbo)}
}

// DataArgs contains the configurations for aws_s3_bucket_object.
type DataArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Key: string, required
	Key terra.StringValue `hcl:"key,attr" validate:"required"`
	// Range: string, optional
	Range terra.StringValue `hcl:"range,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// VersionId: string, optional
	VersionId terra.StringValue `hcl:"version_id,attr"`
}

type dataAwsS3BucketObjectAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("arn"))
}

// Body returns a reference to field body of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("body"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("bucket"))
}

// BucketKeyEnabled returns a reference to field bucket_key_enabled of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(asbo.ref.Append("bucket_key_enabled"))
}

// CacheControl returns a reference to field cache_control of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("cache_control"))
}

// ContentDisposition returns a reference to field content_disposition of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("content_language"))
}

// ContentLength returns a reference to field content_length of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ContentLength() terra.NumberValue {
	return terra.ReferenceAsNumber(asbo.ref.Append("content_length"))
}

// ContentType returns a reference to field content_type of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("content_type"))
}

// Etag returns a reference to field etag of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("etag"))
}

// Expiration returns a reference to field expiration of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("expiration"))
}

// Expires returns a reference to field expires of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("expires"))
}

// Id returns a reference to field id of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("id"))
}

// Key returns a reference to field key of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("key"))
}

// LastModified returns a reference to field last_modified of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("last_modified"))
}

// Metadata returns a reference to field metadata of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asbo.ref.Append("metadata"))
}

// ObjectLockLegalHoldStatus returns a reference to field object_lock_legal_hold_status of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("object_lock_legal_hold_status"))
}

// ObjectLockMode returns a reference to field object_lock_mode of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("object_lock_mode"))
}

// ObjectLockRetainUntilDate returns a reference to field object_lock_retain_until_date of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("object_lock_retain_until_date"))
}

// Range returns a reference to field range of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Range() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("range"))
}

// ServerSideEncryption returns a reference to field server_side_encryption of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("server_side_encryption"))
}

// SseKmsKeyId returns a reference to field sse_kms_key_id of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) SseKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("sse_kms_key_id"))
}

// StorageClass returns a reference to field storage_class of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("storage_class"))
}

// Tags returns a reference to field tags of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asbo.ref.Append("tags"))
}

// VersionId returns a reference to field version_id of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("version_id"))
}

// WebsiteRedirectLocation returns a reference to field website_redirect_location of aws_s3_bucket_object.
func (asbo dataAwsS3BucketObjectAttributes) WebsiteRedirectLocation() terra.StringValue {
	return terra.ReferenceAsString(asbo.ref.Append("website_redirect_location"))
}
