// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataS3BucketObject creates a new instance of [DataS3BucketObject].
func NewDataS3BucketObject(name string, args DataS3BucketObjectArgs) *DataS3BucketObject {
	return &DataS3BucketObject{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3BucketObject)(nil)

// DataS3BucketObject represents the Terraform data resource aws_s3_bucket_object.
type DataS3BucketObject struct {
	Name string
	Args DataS3BucketObjectArgs
}

// DataSource returns the Terraform object type for [DataS3BucketObject].
func (sbo *DataS3BucketObject) DataSource() string {
	return "aws_s3_bucket_object"
}

// LocalName returns the local name for [DataS3BucketObject].
func (sbo *DataS3BucketObject) LocalName() string {
	return sbo.Name
}

// Configuration returns the configuration (args) for [DataS3BucketObject].
func (sbo *DataS3BucketObject) Configuration() interface{} {
	return sbo.Args
}

// Attributes returns the attributes for [DataS3BucketObject].
func (sbo *DataS3BucketObject) Attributes() dataS3BucketObjectAttributes {
	return dataS3BucketObjectAttributes{ref: terra.ReferenceDataResource(sbo)}
}

// DataS3BucketObjectArgs contains the configurations for aws_s3_bucket_object.
type DataS3BucketObjectArgs struct {
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
type dataS3BucketObjectAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("body"))
}

// Bucket returns a reference to field bucket of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("bucket"))
}

// BucketKeyEnabled returns a reference to field bucket_key_enabled of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(sbo.ref.Append("bucket_key_enabled"))
}

// CacheControl returns a reference to field cache_control of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("cache_control"))
}

// ContentDisposition returns a reference to field content_disposition of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_language"))
}

// ContentLength returns a reference to field content_length of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ContentLength() terra.NumberValue {
	return terra.ReferenceAsNumber(sbo.ref.Append("content_length"))
}

// ContentType returns a reference to field content_type of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_type"))
}

// Etag returns a reference to field etag of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("etag"))
}

// Expiration returns a reference to field expiration of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("expiration"))
}

// Expires returns a reference to field expires of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("expires"))
}

// Id returns a reference to field id of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("id"))
}

// Key returns a reference to field key of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("key"))
}

// LastModified returns a reference to field last_modified of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("last_modified"))
}

// Metadata returns a reference to field metadata of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sbo.ref.Append("metadata"))
}

// ObjectLockLegalHoldStatus returns a reference to field object_lock_legal_hold_status of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("object_lock_legal_hold_status"))
}

// ObjectLockMode returns a reference to field object_lock_mode of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("object_lock_mode"))
}

// ObjectLockRetainUntilDate returns a reference to field object_lock_retain_until_date of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("object_lock_retain_until_date"))
}

// Range returns a reference to field range of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Range() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("range"))
}

// ServerSideEncryption returns a reference to field server_side_encryption of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("server_side_encryption"))
}

// SseKmsKeyId returns a reference to field sse_kms_key_id of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) SseKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("sse_kms_key_id"))
}

// StorageClass returns a reference to field storage_class of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("storage_class"))
}

// Tags returns a reference to field tags of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sbo.ref.Append("tags"))
}

// VersionId returns a reference to field version_id of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("version_id"))
}

// WebsiteRedirectLocation returns a reference to field website_redirect_location of aws_s3_bucket_object.
func (sbo dataS3BucketObjectAttributes) WebsiteRedirectLocation() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("website_redirect_location"))
}
