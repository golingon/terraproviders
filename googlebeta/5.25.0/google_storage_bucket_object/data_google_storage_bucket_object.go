// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_storage_bucket_object

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource google_storage_bucket_object.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (gsbo *DataSource) DataSource() string {
	return "google_storage_bucket_object"
}

// LocalName returns the local name for [DataSource].
func (gsbo *DataSource) LocalName() string {
	return gsbo.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (gsbo *DataSource) Configuration() interface{} {
	return gsbo.Args
}

// Attributes returns the attributes for [DataSource].
func (gsbo *DataSource) Attributes() dataGoogleStorageBucketObjectAttributes {
	return dataGoogleStorageBucketObjectAttributes{ref: terra.ReferenceDataSource(gsbo)}
}

// DataArgs contains the configurations for google_storage_bucket_object.
type DataArgs struct {
	// Bucket: string, optional
	Bucket terra.StringValue `hcl:"bucket,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
}

type dataGoogleStorageBucketObjectAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("bucket"))
}

// CacheControl returns a reference to field cache_control of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("cache_control"))
}

// Content returns a reference to field content of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("content"))
}

// ContentDisposition returns a reference to field content_disposition of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("content_type"))
}

// Crc32C returns a reference to field crc32c of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Crc32C() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("crc32c"))
}

// DetectMd5Hash returns a reference to field detect_md5hash of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("detect_md5hash"))
}

// EventBasedHold returns a reference to field event_based_hold of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) EventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(gsbo.ref.Append("event_based_hold"))
}

// Id returns a reference to field id of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("kms_key_name"))
}

// Md5Hash returns a reference to field md5hash of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("md5hash"))
}

// MediaLink returns a reference to field media_link of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) MediaLink() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("media_link"))
}

// Metadata returns a reference to field metadata of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gsbo.ref.Append("metadata"))
}

// Name returns a reference to field name of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("name"))
}

// OutputName returns a reference to field output_name of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) OutputName() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("output_name"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("self_link"))
}

// Source returns a reference to field source of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("source"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(gsbo.ref.Append("storage_class"))
}

// TemporaryHold returns a reference to field temporary_hold of google_storage_bucket_object.
func (gsbo dataGoogleStorageBucketObjectAttributes) TemporaryHold() terra.BoolValue {
	return terra.ReferenceAsBool(gsbo.ref.Append("temporary_hold"))
}

func (gsbo dataGoogleStorageBucketObjectAttributes) CustomerEncryption() terra.ListValue[DataCustomerEncryptionAttributes] {
	return terra.ReferenceAsList[DataCustomerEncryptionAttributes](gsbo.ref.Append("customer_encryption"))
}

func (gsbo dataGoogleStorageBucketObjectAttributes) Retention() terra.ListValue[DataRetentionAttributes] {
	return terra.ReferenceAsList[DataRetentionAttributes](gsbo.ref.Append("retention"))
}
