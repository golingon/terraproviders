// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google

import (
	"github.com/golingon/lingon/pkg/terra"
	datacomputebackendbucket "github.com/golingon/terraproviders/google/5.24.0/datacomputebackendbucket"
)

// NewDataComputeBackendBucket creates a new instance of [DataComputeBackendBucket].
func NewDataComputeBackendBucket(name string, args DataComputeBackendBucketArgs) *DataComputeBackendBucket {
	return &DataComputeBackendBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataComputeBackendBucket)(nil)

// DataComputeBackendBucket represents the Terraform data resource google_compute_backend_bucket.
type DataComputeBackendBucket struct {
	Name string
	Args DataComputeBackendBucketArgs
}

// DataSource returns the Terraform object type for [DataComputeBackendBucket].
func (cbb *DataComputeBackendBucket) DataSource() string {
	return "google_compute_backend_bucket"
}

// LocalName returns the local name for [DataComputeBackendBucket].
func (cbb *DataComputeBackendBucket) LocalName() string {
	return cbb.Name
}

// Configuration returns the configuration (args) for [DataComputeBackendBucket].
func (cbb *DataComputeBackendBucket) Configuration() interface{} {
	return cbb.Args
}

// Attributes returns the attributes for [DataComputeBackendBucket].
func (cbb *DataComputeBackendBucket) Attributes() dataComputeBackendBucketAttributes {
	return dataComputeBackendBucketAttributes{ref: terra.ReferenceDataResource(cbb)}
}

// DataComputeBackendBucketArgs contains the configurations for google_compute_backend_bucket.
type DataComputeBackendBucketArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// CdnPolicy: min=0
	CdnPolicy []datacomputebackendbucket.CdnPolicy `hcl:"cdn_policy,block" validate:"min=0"`
}
type dataComputeBackendBucketAttributes struct {
	ref terra.Reference
}

// BucketName returns a reference to field bucket_name of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("bucket_name"))
}

// CompressionMode returns a reference to field compression_mode of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) CompressionMode() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("compression_mode"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("creation_timestamp"))
}

// CustomResponseHeaders returns a reference to field custom_response_headers of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) CustomResponseHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cbb.ref.Append("custom_response_headers"))
}

// Description returns a reference to field description of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("description"))
}

// EdgeSecurityPolicy returns a reference to field edge_security_policy of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("edge_security_policy"))
}

// EnableCdn returns a reference to field enable_cdn of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceAsBool(cbb.ref.Append("enable_cdn"))
}

// Id returns a reference to field id of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_backend_bucket.
func (cbb dataComputeBackendBucketAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("self_link"))
}

func (cbb dataComputeBackendBucketAttributes) CdnPolicy() terra.ListValue[datacomputebackendbucket.CdnPolicyAttributes] {
	return terra.ReferenceAsList[datacomputebackendbucket.CdnPolicyAttributes](cbb.ref.Append("cdn_policy"))
}
