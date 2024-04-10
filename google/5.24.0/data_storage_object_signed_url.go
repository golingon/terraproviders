// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import "github.com/golingon/lingon/pkg/terra"

// NewDataStorageObjectSignedUrl creates a new instance of [DataStorageObjectSignedUrl].
func NewDataStorageObjectSignedUrl(name string, args DataStorageObjectSignedUrlArgs) *DataStorageObjectSignedUrl {
	return &DataStorageObjectSignedUrl{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStorageObjectSignedUrl)(nil)

// DataStorageObjectSignedUrl represents the Terraform data resource google_storage_object_signed_url.
type DataStorageObjectSignedUrl struct {
	Name string
	Args DataStorageObjectSignedUrlArgs
}

// DataSource returns the Terraform object type for [DataStorageObjectSignedUrl].
func (sosu *DataStorageObjectSignedUrl) DataSource() string {
	return "google_storage_object_signed_url"
}

// LocalName returns the local name for [DataStorageObjectSignedUrl].
func (sosu *DataStorageObjectSignedUrl) LocalName() string {
	return sosu.Name
}

// Configuration returns the configuration (args) for [DataStorageObjectSignedUrl].
func (sosu *DataStorageObjectSignedUrl) Configuration() interface{} {
	return sosu.Args
}

// Attributes returns the attributes for [DataStorageObjectSignedUrl].
func (sosu *DataStorageObjectSignedUrl) Attributes() dataStorageObjectSignedUrlAttributes {
	return dataStorageObjectSignedUrlAttributes{ref: terra.ReferenceDataResource(sosu)}
}

// DataStorageObjectSignedUrlArgs contains the configurations for google_storage_object_signed_url.
type DataStorageObjectSignedUrlArgs struct {
	// Bucket: string, required
	Bucket terra.StringValue `hcl:"bucket,attr" validate:"required"`
	// ContentMd5: string, optional
	ContentMd5 terra.StringValue `hcl:"content_md5,attr"`
	// ContentType: string, optional
	ContentType terra.StringValue `hcl:"content_type,attr"`
	// Credentials: string, optional
	Credentials terra.StringValue `hcl:"credentials,attr"`
	// Duration: string, optional
	Duration terra.StringValue `hcl:"duration,attr"`
	// ExtensionHeaders: map of string, optional
	ExtensionHeaders terra.MapValue[terra.StringValue] `hcl:"extension_headers,attr"`
	// HttpMethod: string, optional
	HttpMethod terra.StringValue `hcl:"http_method,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Path: string, required
	Path terra.StringValue `hcl:"path,attr" validate:"required"`
}
type dataStorageObjectSignedUrlAttributes struct {
	ref terra.Reference
}

// Bucket returns a reference to field bucket of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) Bucket() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("bucket"))
}

// ContentMd5 returns a reference to field content_md5 of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) ContentMd5() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("content_md5"))
}

// ContentType returns a reference to field content_type of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) ContentType() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("content_type"))
}

// Credentials returns a reference to field credentials of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) Credentials() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("credentials"))
}

// Duration returns a reference to field duration of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) Duration() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("duration"))
}

// ExtensionHeaders returns a reference to field extension_headers of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) ExtensionHeaders() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sosu.ref.Append("extension_headers"))
}

// HttpMethod returns a reference to field http_method of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("http_method"))
}

// Id returns a reference to field id of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("id"))
}

// Path returns a reference to field path of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) Path() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("path"))
}

// SignedUrl returns a reference to field signed_url of google_storage_object_signed_url.
func (sosu dataStorageObjectSignedUrlAttributes) SignedUrl() terra.StringValue {
	return terra.ReferenceAsString(sosu.ref.Append("signed_url"))
}
