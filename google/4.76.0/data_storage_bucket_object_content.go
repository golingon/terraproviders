// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	datastoragebucketobjectcontent "github.com/golingon/terraproviders/google/4.76.0/datastoragebucketobjectcontent"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStorageBucketObjectContent creates a new instance of [DataStorageBucketObjectContent].
func NewDataStorageBucketObjectContent(name string, args DataStorageBucketObjectContentArgs) *DataStorageBucketObjectContent {
	return &DataStorageBucketObjectContent{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageBucketObjectContent)(nil)

// DataStorageBucketObjectContent represents the Terraform data resource google_storage_bucket_object_content.
type DataStorageBucketObjectContent struct {
	Name string
	Args DataStorageBucketObjectContentArgs
}

// DataSource returns the Terraform object type for [DataStorageBucketObjectContent].
func (sboc *DataStorageBucketObjectContent) DataSource() string {
	return "google_storage_bucket_object_content"
}

// LocalName returns the local name for [DataStorageBucketObjectContent].
func (sboc *DataStorageBucketObjectContent) LocalName() string {
	return sboc.Name
}

// Configuration returns the configuration (args) for [DataStorageBucketObjectContent].
func (sboc *DataStorageBucketObjectContent) Configuration() interface{} {
	return sboc.Args
}

// Attributes returns the attributes for [DataStorageBucketObjectContent].
func (sboc *DataStorageBucketObjectContent) Attributes() dataStorageBucketObjectContentAttributes {
	return dataStorageBucketObjectContentAttributes{ref: terra.ReferenceDataResource(sboc)}
}

// DataStorageBucketObjectContentArgs contains the configurations for google_storage_bucket_object_content.
type DataStorageBucketObjectContentArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// CustomerEncryption: min=0
	CustomerEncryption []datastoragebucketobjectcontent.CustomerEncryption `hcl:"customer_encryption,block" validate:"min=0"`
}
type dataStorageBucketObjectContentAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("bucket"))
}

// CacheControl returns a reference to field cache_control of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("cache_control"))
}

// Content returns a reference to field content of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("content"))
}

// ContentDisposition returns a reference to field content_disposition of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("content_type"))
}

// Crc32C returns a reference to field crc32c of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Crc32C() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("crc32c"))
}

// DetectMd5Hash returns a reference to field detect_md5hash of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("detect_md5hash"))
}

// EventBasedHold returns a reference to field event_based_hold of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) EventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(sboc.ref.Append("event_based_hold"))
}

// Id returns a reference to field id of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("kms_key_name"))
}

// Md5Hash returns a reference to field md5hash of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("md5hash"))
}

// MediaLink returns a reference to field media_link of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) MediaLink() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("media_link"))
}

// Metadata returns a reference to field metadata of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sboc.ref.Append("metadata"))
}

// Name returns a reference to field name of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("name"))
}

// OutputName returns a reference to field output_name of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) OutputName() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("output_name"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("self_link"))
}

// Source returns a reference to field source of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("source"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sboc.ref.Append("storage_class"))
}

// TemporaryHold returns a reference to field temporary_hold of google_storage_bucket_object_content.
func (sboc dataStorageBucketObjectContentAttributes) TemporaryHold() terra.BoolValue {
	return terra.ReferenceAsBool(sboc.ref.Append("temporary_hold"))
}

func (sboc dataStorageBucketObjectContentAttributes) CustomerEncryption() terra.ListValue[datastoragebucketobjectcontent.CustomerEncryptionAttributes] {
	return terra.ReferenceAsList[datastoragebucketobjectcontent.CustomerEncryptionAttributes](sboc.ref.Append("customer_encryption"))
}
