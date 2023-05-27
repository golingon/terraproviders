// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computebackendbucket "github.com/golingon/terraproviders/google/4.66.0/computebackendbucket"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendBucket creates a new instance of [ComputeBackendBucket].
func NewComputeBackendBucket(name string, args ComputeBackendBucketArgs) *ComputeBackendBucket {
	return &ComputeBackendBucket{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendBucket)(nil)

// ComputeBackendBucket represents the Terraform resource google_compute_backend_bucket.
type ComputeBackendBucket struct {
	Name      string
	Args      ComputeBackendBucketArgs
	state     *computeBackendBucketState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) Type() string {
	return "google_compute_backend_bucket"
}

// LocalName returns the local name for [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) LocalName() string {
	return cbb.Name
}

// Configuration returns the configuration (args) for [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) Configuration() interface{} {
	return cbb.Args
}

// DependOn is used for other resources to depend on [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) DependOn() terra.Reference {
	return terra.ReferenceResource(cbb)
}

// Dependencies returns the list of resources [ComputeBackendBucket] depends_on.
func (cbb *ComputeBackendBucket) Dependencies() terra.Dependencies {
	return cbb.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) LifecycleManagement() *terra.Lifecycle {
	return cbb.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendBucket].
func (cbb *ComputeBackendBucket) Attributes() computeBackendBucketAttributes {
	return computeBackendBucketAttributes{ref: terra.ReferenceResource(cbb)}
}

// ImportState imports the given attribute values into [ComputeBackendBucket]'s state.
func (cbb *ComputeBackendBucket) ImportState(av io.Reader) error {
	cbb.state = &computeBackendBucketState{}
	if err := json.NewDecoder(av).Decode(cbb.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbb.Type(), cbb.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendBucket] has state.
func (cbb *ComputeBackendBucket) State() (*computeBackendBucketState, bool) {
	return cbb.state, cbb.state != nil
}

// StateMust returns the state for [ComputeBackendBucket]. Panics if the state is nil.
func (cbb *ComputeBackendBucket) StateMust() *computeBackendBucketState {
	if cbb.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbb.Type(), cbb.LocalName()))
	}
	return cbb.state
}

// ComputeBackendBucketArgs contains the configurations for google_compute_backend_bucket.
type ComputeBackendBucketArgs struct {
	// BucketName: string, required
	BucketName terra.StringValue `hcl:"bucket_name,attr" validate:"required"`
	// CompressionMode: string, optional
	CompressionMode terra.StringValue `hcl:"compression_mode,attr"`
	// CustomResponseHeaders: list of string, optional
	CustomResponseHeaders terra.ListValue[terra.StringValue] `hcl:"custom_response_headers,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EdgeSecurityPolicy: string, optional
	EdgeSecurityPolicy terra.StringValue `hcl:"edge_security_policy,attr"`
	// EnableCdn: bool, optional
	EnableCdn terra.BoolValue `hcl:"enable_cdn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// CdnPolicy: optional
	CdnPolicy *computebackendbucket.CdnPolicy `hcl:"cdn_policy,block"`
	// Timeouts: optional
	Timeouts *computebackendbucket.Timeouts `hcl:"timeouts,block"`
}
type computeBackendBucketAttributes struct {
	ref terra.Reference
}

// BucketName returns a reference to field bucket_name of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) BucketName() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("bucket_name"))
}

// CompressionMode returns a reference to field compression_mode of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) CompressionMode() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("compression_mode"))
}

// CreationTimestamp returns a reference to field creation_timestamp of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) CreationTimestamp() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("creation_timestamp"))
}

// CustomResponseHeaders returns a reference to field custom_response_headers of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) CustomResponseHeaders() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](cbb.ref.Append("custom_response_headers"))
}

// Description returns a reference to field description of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("description"))
}

// EdgeSecurityPolicy returns a reference to field edge_security_policy of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) EdgeSecurityPolicy() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("edge_security_policy"))
}

// EnableCdn returns a reference to field enable_cdn of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) EnableCdn() terra.BoolValue {
	return terra.ReferenceAsBool(cbb.ref.Append("enable_cdn"))
}

// Id returns a reference to field id of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("id"))
}

// Name returns a reference to field name of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("project"))
}

// SelfLink returns a reference to field self_link of google_compute_backend_bucket.
func (cbb computeBackendBucketAttributes) SelfLink() terra.StringValue {
	return terra.ReferenceAsString(cbb.ref.Append("self_link"))
}

func (cbb computeBackendBucketAttributes) CdnPolicy() terra.ListValue[computebackendbucket.CdnPolicyAttributes] {
	return terra.ReferenceAsList[computebackendbucket.CdnPolicyAttributes](cbb.ref.Append("cdn_policy"))
}

func (cbb computeBackendBucketAttributes) Timeouts() computebackendbucket.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computebackendbucket.TimeoutsAttributes](cbb.ref.Append("timeouts"))
}

type computeBackendBucketState struct {
	BucketName            string                                `json:"bucket_name"`
	CompressionMode       string                                `json:"compression_mode"`
	CreationTimestamp     string                                `json:"creation_timestamp"`
	CustomResponseHeaders []string                              `json:"custom_response_headers"`
	Description           string                                `json:"description"`
	EdgeSecurityPolicy    string                                `json:"edge_security_policy"`
	EnableCdn             bool                                  `json:"enable_cdn"`
	Id                    string                                `json:"id"`
	Name                  string                                `json:"name"`
	Project               string                                `json:"project"`
	SelfLink              string                                `json:"self_link"`
	CdnPolicy             []computebackendbucket.CdnPolicyState `json:"cdn_policy"`
	Timeouts              *computebackendbucket.TimeoutsState   `json:"timeouts"`
}
