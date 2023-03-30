// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import "github.com/volvo-cars/lingon/pkg/terra"

func NewDataS3BucketObject(name string, args DataS3BucketObjectArgs) *DataS3BucketObject {
	return &DataS3BucketObject{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataS3BucketObject)(nil)

type DataS3BucketObject struct {
	Name string
	Args DataS3BucketObjectArgs
}

func (sbo *DataS3BucketObject) DataSource() string {
	return "aws_s3_bucket_object"
}

func (sbo *DataS3BucketObject) LocalName() string {
	return sbo.Name
}

func (sbo *DataS3BucketObject) Configuration() interface{} {
	return sbo.Args
}

func (sbo *DataS3BucketObject) Attributes() dataS3BucketObjectAttributes {
	return dataS3BucketObjectAttributes{ref: terra.ReferenceDataResource(sbo)}
}

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

func (sbo dataS3BucketObjectAttributes) Body() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("body"))
}

func (sbo dataS3BucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("bucket"))
}

func (sbo dataS3BucketObjectAttributes) BucketKeyEnabled() terra.BoolValue {
	return terra.ReferenceBool(sbo.ref.Append("bucket_key_enabled"))
}

func (sbo dataS3BucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("cache_control"))
}

func (sbo dataS3BucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("content_disposition"))
}

func (sbo dataS3BucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("content_encoding"))
}

func (sbo dataS3BucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("content_language"))
}

func (sbo dataS3BucketObjectAttributes) ContentLength() terra.NumberValue {
	return terra.ReferenceNumber(sbo.ref.Append("content_length"))
}

func (sbo dataS3BucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("content_type"))
}

func (sbo dataS3BucketObjectAttributes) Etag() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("etag"))
}

func (sbo dataS3BucketObjectAttributes) Expiration() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("expiration"))
}

func (sbo dataS3BucketObjectAttributes) Expires() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("expires"))
}

func (sbo dataS3BucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("id"))
}

func (sbo dataS3BucketObjectAttributes) Key() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("key"))
}

func (sbo dataS3BucketObjectAttributes) LastModified() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("last_modified"))
}

func (sbo dataS3BucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sbo.ref.Append("metadata"))
}

func (sbo dataS3BucketObjectAttributes) ObjectLockLegalHoldStatus() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("object_lock_legal_hold_status"))
}

func (sbo dataS3BucketObjectAttributes) ObjectLockMode() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("object_lock_mode"))
}

func (sbo dataS3BucketObjectAttributes) ObjectLockRetainUntilDate() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("object_lock_retain_until_date"))
}

func (sbo dataS3BucketObjectAttributes) Range() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("range"))
}

func (sbo dataS3BucketObjectAttributes) ServerSideEncryption() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("server_side_encryption"))
}

func (sbo dataS3BucketObjectAttributes) SseKmsKeyId() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("sse_kms_key_id"))
}

func (sbo dataS3BucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("storage_class"))
}

func (sbo dataS3BucketObjectAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sbo.ref.Append("tags"))
}

func (sbo dataS3BucketObjectAttributes) VersionId() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("version_id"))
}

func (sbo dataS3BucketObjectAttributes) WebsiteRedirectLocation() terra.StringValue {
	return terra.ReferenceString(sbo.ref.Append("website_redirect_location"))
}
