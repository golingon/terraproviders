// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datastoragebucketobject "github.com/golingon/terraproviders/google/5.24.0/datastoragebucketobject"
)

// NewDataStorageBucketObject creates a new instance of [DataStorageBucketObject].
func NewDataStorageBucketObject(name string, args DataStorageBucketObjectArgs) *DataStorageBucketObject {
	return &DataStorageBucketObject{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageBucketObject)(nil)

// DataStorageBucketObject represents the Terraform data resource google_storage_bucket_object.
type DataStorageBucketObject struct {
	Name string
	Args DataStorageBucketObjectArgs
}

// DataSource returns the Terraform object type for [DataStorageBucketObject].
func (sbo *DataStorageBucketObject) DataSource() string {
	return "google_storage_bucket_object"
}

// LocalName returns the local name for [DataStorageBucketObject].
func (sbo *DataStorageBucketObject) LocalName() string {
	return sbo.Name
}

// Configuration returns the configuration (args) for [DataStorageBucketObject].
func (sbo *DataStorageBucketObject) Configuration() interface{} {
	return sbo.Args
}

// Attributes returns the attributes for [DataStorageBucketObject].
func (sbo *DataStorageBucketObject) Attributes() dataStorageBucketObjectAttributes {
	return dataStorageBucketObjectAttributes{ref: terra.ReferenceDataResource(sbo)}
}

// DataStorageBucketObjectArgs contains the configurations for google_storage_bucket_object.
type DataStorageBucketObjectArgs struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// CustomerEncryption: min=0
	CustomerEncryption []datastoragebucketobject.CustomerEncryption `hcl:"customer_encryption,block" validate:"min=0"`
	// Retention: min=0
	Retention []datastoragebucketobject.Retention `hcl:"retention,block" validate:"min=0"`
}
type dataStorageBucketObjectAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("bucket"))
}

// CacheControl returns a reference to field cache_control of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("cache_control"))
}

// Content returns a reference to field content of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content"))
}

// ContentDisposition returns a reference to field content_disposition of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_type"))
}

// Crc32C returns a reference to field crc32c of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Crc32C() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("crc32c"))
}

// DetectMd5Hash returns a reference to field detect_md5hash of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("detect_md5hash"))
}

// EventBasedHold returns a reference to field event_based_hold of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) EventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(sbo.ref.Append("event_based_hold"))
}

// Id returns a reference to field id of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("kms_key_name"))
}

// Md5Hash returns a reference to field md5hash of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("md5hash"))
}

// MediaLink returns a reference to field media_link of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) MediaLink() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("media_link"))
}

// Metadata returns a reference to field metadata of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sbo.ref.Append("metadata"))
}

// Name returns a reference to field name of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("name"))
}

// OutputName returns a reference to field output_name of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) OutputName() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("output_name"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("self_link"))
}

// Source returns a reference to field source of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("source"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("storage_class"))
}

// TemporaryHold returns a reference to field temporary_hold of google_storage_bucket_object.
func (sbo dataStorageBucketObjectAttributes) TemporaryHold() terra.BoolValue {
	return terra.ReferenceAsBool(sbo.ref.Append("temporary_hold"))
}

func (sbo dataStorageBucketObjectAttributes) CustomerEncryption() terra.ListValue[datastoragebucketobject.CustomerEncryptionAttributes] {
	return terra.ReferenceAsList[datastoragebucketobject.CustomerEncryptionAttributes](sbo.ref.Append("customer_encryption"))
}

func (sbo dataStorageBucketObjectAttributes) Retention() terra.ListValue[datastoragebucketobject.RetentionAttributes] {
	return terra.ReferenceAsList[datastoragebucketobject.RetentionAttributes](sbo.ref.Append("retention"))
}
