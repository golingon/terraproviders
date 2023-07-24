// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	storagebucketobject "github.com/golingon/terraproviders/googlebeta/4.74.0/storagebucketobject"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStorageBucketObject creates a new instance of [StorageBucketObject].
func NewStorageBucketObject(name string, args StorageBucketObjectArgs) *StorageBucketObject {
	return &StorageBucketObject{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StorageBucketObject)(nil)

// StorageBucketObject represents the Terraform resource google_storage_bucket_object.
type StorageBucketObject struct {
	Name      string
	Args      StorageBucketObjectArgs
	state     *storageBucketObjectState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StorageBucketObject].
func (sbo *StorageBucketObject) Type() string {
	return "google_storage_bucket_object"
}

// LocalName returns the local name for [StorageBucketObject].
func (sbo *StorageBucketObject) LocalName() string {
	return sbo.Name
}

// Configuration returns the configuration (args) for [StorageBucketObject].
func (sbo *StorageBucketObject) Configuration() interface{} {
	return sbo.Args
}

// DependOn is used for other resources to depend on [StorageBucketObject].
func (sbo *StorageBucketObject) DependOn() terra.Reference {
	return terra.ReferenceResource(sbo)
}

// Dependencies returns the list of resources [StorageBucketObject] depends_on.
func (sbo *StorageBucketObject) Dependencies() terra.Dependencies {
	return sbo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StorageBucketObject].
func (sbo *StorageBucketObject) LifecycleManagement() *terra.Lifecycle {
	return sbo.Lifecycle
}

// Attributes returns the attributes for [StorageBucketObject].
func (sbo *StorageBucketObject) Attributes() storageBucketObjectAttributes {
	return storageBucketObjectAttributes{ref: terra.ReferenceResource(sbo)}
}

// ImportState imports the given attribute values into [StorageBucketObject]'s state.
func (sbo *StorageBucketObject) ImportState(av io.Reader) error {
	sbo.state = &storageBucketObjectState{}
	if err := json.NewDecoder(av).Decode(sbo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sbo.Type(), sbo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StorageBucketObject] has state.
func (sbo *StorageBucketObject) State() (*storageBucketObjectState, bool) {
	return sbo.state, sbo.state != nil
}

// StateMust returns the state for [StorageBucketObject]. Panics if the state is nil.
func (sbo *StorageBucketObject) StateMust() *storageBucketObjectState {
	if sbo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sbo.Type(), sbo.LocalName()))
	}
	return sbo.state
}

// StorageBucketObjectArgs contains the configurations for google_storage_bucket_object.
type StorageBucketObjectArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// CacheControl: string, optional
	CacheControl terra.StringValue `hcl:"cache_control,attr"`
	// Content: string, optional
	Content terra.StringValue `hcl:"content,attr"`
	// ContentDisposition: string, optional
	ContentDisposition terra.StringValue `hcl:"content_disposition,attr"`
	// ContentEncoding: string, optional
	ContentEncoding terra.StringValue `hcl:"content_encoding,attr"`
	// ContentLanguage: string, optional
	ContentLanguage terra.StringValue `hcl:"content_language,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// DetectMd5Hash: string, optional
	DetectMd5Hash terra.StringValue `hcl:"detect_md5hash,attr"`
	// EventBasedHold: bool, optional
	EventBasedHold terra.BoolValue `hcl:"event_based_hold,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KmsKeyName: string, optional
	KmsKeyName terra.StringValue `hcl:"kms_key_name,attr"`
	// Metadata: map of string, optional
	Metadata terra.MapValue[terra.StringValue] `hcl:"metadata,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Source: string, optional
	Source terra.StringValue `hcl:"source,attr"`
	// StorageClass: string, optional
	StorageClass terra.StringValue `hcl:"storage_class,attr"`
	// TemporaryHold: bool, optional
	TemporaryHold terra.BoolValue `hcl:"temporary_hold,attr"`
	// CustomerEncryption: optional
	CustomerEncryption *storagebucketobject.CustomerEncryption `hcl:"customer_encryption,block"`
	// Timeouts: optional
	Timeouts *storagebucketobject.Timeouts `hcl:"timeouts,block"`
}
type storageBucketObjectAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("bucket"))
}

// CacheControl returns a reference to field cache_control of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) CacheControl() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("cache_control"))
}

// Content returns a reference to field content of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Content() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content"))
}

// ContentDisposition returns a reference to field content_disposition of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) ContentDisposition() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_disposition"))
}

// ContentEncoding returns a reference to field content_encoding of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) ContentEncoding() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_encoding"))
}

// ContentLanguage returns a reference to field content_language of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) ContentLanguage() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_language"))
}

// ContentType returns a reference to field content_type of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("content_type"))
}

// Crc32C returns a reference to field crc32c of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Crc32C() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("crc32c"))
}

// DetectMd5Hash returns a reference to field detect_md5hash of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) DetectMd5Hash() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("detect_md5hash"))
}

// EventBasedHold returns a reference to field event_based_hold of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) EventBasedHold() terra.BoolValue {
	return terra.ReferenceAsBool(sbo.ref.Append("event_based_hold"))
}

// Id returns a reference to field id of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("id"))
}

// KmsKeyName returns a reference to field kms_key_name of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) KmsKeyName() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("kms_key_name"))
}

// Md5Hash returns a reference to field md5hash of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Md5Hash() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("md5hash"))
}

// MediaLink returns a reference to field media_link of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) MediaLink() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("media_link"))
}

// Metadata returns a reference to field metadata of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Metadata() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sbo.ref.Append("metadata"))
}

// Name returns a reference to field name of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("name"))
}

// OutputName returns a reference to field output_name of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) OutputName() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("output_name"))
}

// SelfLink returns a reference to field self_link of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("self_link"))
}

// Source returns a reference to field source of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) Source() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("source"))
}

// StorageClass returns a reference to field storage_class of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) StorageClass() terra.StringValue {
	return terra.ReferenceAsString(sbo.ref.Append("storage_class"))
}

// TemporaryHold returns a reference to field temporary_hold of google_storage_bucket_object.
func (sbo storageBucketObjectAttributes) TemporaryHold() terra.BoolValue {
	return terra.ReferenceAsBool(sbo.ref.Append("temporary_hold"))
}

func (sbo storageBucketObjectAttributes) CustomerEncryption() terra.ListValue[storagebucketobject.CustomerEncryptionAttributes] {
	return terra.ReferenceAsList[storagebucketobject.CustomerEncryptionAttributes](sbo.ref.Append("customer_encryption"))
}

func (sbo storageBucketObjectAttributes) Timeouts() storagebucketobject.TimeoutsAttributes {
	return terra.ReferenceAsSingle[storagebucketobject.TimeoutsAttributes](sbo.ref.Append("timeouts"))
}

type storageBucketObjectState struct {
	Bucket             string                                        `json:"bucket"`
	CacheControl       string                                        `json:"cache_control"`
	Content            string                                        `json:"content"`
	ContentDisposition string                                        `json:"content_disposition"`
	ContentEncoding    string                                        `json:"content_encoding"`
	ContentLanguage    string                                        `json:"content_language"`
	ContentType        string                                        `json:"content_type"`
	Crc32C             string                                        `json:"crc32c"`
	DetectMd5Hash      string                                        `json:"detect_md5hash"`
	EventBasedHold     bool                                          `json:"event_based_hold"`
	Id                 string                                        `json:"id"`
	KmsKeyName         string                                        `json:"kms_key_name"`
	Md5Hash            string                                        `json:"md5hash"`
	MediaLink          string                                        `json:"media_link"`
	Metadata           map[string]string                             `json:"metadata"`
	Name               string                                        `json:"name"`
	OutputName         string                                        `json:"output_name"`
	SelfLink           string                                        `json:"self_link"`
	Source             string                                        `json:"source"`
	StorageClass       string                                        `json:"storage_class"`
	TemporaryHold      bool                                          `json:"temporary_hold"`
	CustomerEncryption []storagebucketobject.CustomerEncryptionState `json:"customer_encryption"`
	Timeouts           *storagebucketobject.TimeoutsState            `json:"timeouts"`
}
