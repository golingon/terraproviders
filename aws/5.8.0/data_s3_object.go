// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

// NewDataS3Object creates a new instance of [DataS3Object].
func NewDataS3Object(name string, args DataS3ObjectArgs) *DataS3Object {
	return &DataS3Object{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3Object)(nil)

// DataS3Object represents the Terraform data resource aws_s3_object.
type DataS3Object struct {
	Name string
	Args DataS3ObjectArgs
}

// DataSource returns the Terraform object type for [DataS3Object].
func (so *DataS3Object) DataSource() string {
	return "aws_s3_object"
}

// LocalName returns the local name for [DataS3Object].
func (so *DataS3Object) LocalName() string {
	return so.Name
}

// Configuration returns the configuration (args) for [DataS3Object].
func (so *DataS3Object) Configuration() interface{} {
	return so.Args
}

// Attributes returns the attributes for [DataS3Object].
func (so *DataS3Object) Attributes() dataS3ObjectAttributes {
	return dataS3ObjectAttributes{ref: terra.ReferenceDataResource(so)}
}

// DataS3ObjectArgs contains the configurations for aws_s3_object.
type DataS3ObjectArgs struct {
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
type dataS3ObjectAttributes struct {
	ref terra.Reference
}

// Body returns a reference to field body of aws_s3_object.
func (so dataS3ObjectAttributes) Body() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("body"))
}

// Bucket returns a reference to field bucket of aws_s3_object.
func (so dataS3ObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("bucket"))
}

// BucketKeyEnabled returns a reference to field bucket_key_enabled of aws_s3_object.
func (so dataS3ObjectAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(so.ref.Append("bucket_key_enabled"))
}

// CacheControl returns a reference to field cache_control of aws_s3_object.
func (so dataS3ObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("cache_control"))
}

// ContentDisposition returns a reference to field content_disposition of aws_s3_object.
func (so dataS3ObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of aws_s3_object.
func (so dataS3ObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of aws_s3_object.
func (so dataS3ObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_language"))
}

// ContentLength returns a reference to field content_length of aws_s3_object.
func (so dataS3ObjectAttributes) ContentLength() terra.NumberValue {
	return terra.ReferenceAsNumber(so.ref.Append("content_length"))
}

// ContentType returns a reference to field content_type of aws_s3_object.
func (so dataS3ObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("content_type"))
}

// Etag returns a reference to field etag of aws_s3_object.
func (so dataS3ObjectAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("etag"))
}

// Expiration returns a reference to field expiration of aws_s3_object.
func (so dataS3ObjectAttributes) Expiration() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("expiration"))
}

// Expires returns a reference to field expires of aws_s3_object.
func (so dataS3ObjectAttributes) Expires() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("expires"))
}

// Id returns a reference to field id of aws_s3_object.
func (so dataS3ObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("id"))
}

// Key returns a reference to field key of aws_s3_object.
func (so dataS3ObjectAttributes) Key() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("key"))
}

// LastModified returns a reference to field last_modified of aws_s3_object.
func (so dataS3ObjectAttributes) LastModified() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("last_modified"))
}

// Metadata returns a reference to field metadata of aws_s3_object.
func (so dataS3ObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](so.ref.Append("metadata"))
}

// ObjectLockLegalHoldStatus returns a reference to field object_lock_legal_hold_status of aws_s3_object.
func (so dataS3ObjectAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_legal_hold_status"))
}

// ObjectLockMode returns a reference to field object_lock_mode of aws_s3_object.
func (so dataS3ObjectAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_mode"))
}

// ObjectLockRetainUntilDate returns a reference to field object_lock_retain_until_date of aws_s3_object.
func (so dataS3ObjectAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("object_lock_retain_until_date"))
}

// Range returns a reference to field range of aws_s3_object.
func (so dataS3ObjectAttributes) Range() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("range"))
}

// ServerSideEncryption returns a reference to field server_side_encryption of aws_s3_object.
func (so dataS3ObjectAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("server_side_encryption"))
}

// SseKmsKeyId returns a reference to field sse_kms_key_id of aws_s3_object.
func (so dataS3ObjectAttributes) SseKmsKeyId() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("sse_kms_key_id"))
}

// StorageClass returns a reference to field storage_class of aws_s3_object.
func (so dataS3ObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("storage_class"))
}

// Tags returns a reference to field tags of aws_s3_object.
func (so dataS3ObjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](so.ref.Append("tags"))
}

// VersionId returns a reference to field version_id of aws_s3_object.
func (so dataS3ObjectAttributes) VersionId() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("version_id"))
}

// WebsiteRedirectLocation returns a reference to field website_redirect_location of aws_s3_object.
func (so dataS3ObjectAttributes) WebsiteRedirectLocation() terra.StringValue {
	return terra.ReferenceAsString(so.ref.Append("website_redirect_location"))
}