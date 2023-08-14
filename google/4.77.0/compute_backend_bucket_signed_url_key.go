// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	computebackendbucketsignedurlkey "github.com/golingon/terraproviders/google/4.77.0/computebackendbucketsignedurlkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendBucketSignedUrlKey creates a new instance of [ComputeBackendBucketSignedUrlKey].
func NewComputeBackendBucketSignedUrlKey(name string, args ComputeBackendBucketSignedUrlKeyArgs) *ComputeBackendBucketSignedUrlKey {
	return &ComputeBackendBucketSignedUrlKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendBucketSignedUrlKey)(nil)

// ComputeBackendBucketSignedUrlKey represents the Terraform resource google_compute_backend_bucket_signed_url_key.
type ComputeBackendBucketSignedUrlKey struct {
	Name      string
	Args      ComputeBackendBucketSignedUrlKeyArgs
	state     *computeBackendBucketSignedUrlKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) Type() string {
	return "google_compute_backend_bucket_signed_url_key"
}

// LocalName returns the local name for [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) LocalName() string {
	return cbbsuk.Name
}

// Configuration returns the configuration (args) for [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) Configuration() interface{} {
	return cbbsuk.Args
}

// DependOn is used for other resources to depend on [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) DependOn() terra.Reference {
	return terra.ReferenceResource(cbbsuk)
}

// Dependencies returns the list of resources [ComputeBackendBucketSignedUrlKey] depends_on.
func (cbbsuk *ComputeBackendBucketSignedUrlKey) Dependencies() terra.Dependencies {
	return cbbsuk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) LifecycleManagement() *terra.Lifecycle {
	return cbbsuk.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendBucketSignedUrlKey].
func (cbbsuk *ComputeBackendBucketSignedUrlKey) Attributes() computeBackendBucketSignedUrlKeyAttributes {
	return computeBackendBucketSignedUrlKeyAttributes{ref: terra.ReferenceResource(cbbsuk)}
}

// ImportState imports the given attribute values into [ComputeBackendBucketSignedUrlKey]'s state.
func (cbbsuk *ComputeBackendBucketSignedUrlKey) ImportState(av io.Reader) error {
	cbbsuk.state = &computeBackendBucketSignedUrlKeyState{}
	if err := json.NewDecoder(av).Decode(cbbsuk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbbsuk.Type(), cbbsuk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendBucketSignedUrlKey] has state.
func (cbbsuk *ComputeBackendBucketSignedUrlKey) State() (*computeBackendBucketSignedUrlKeyState, bool) {
	return cbbsuk.state, cbbsuk.state != nil
}

// StateMust returns the state for [ComputeBackendBucketSignedUrlKey]. Panics if the state is nil.
func (cbbsuk *ComputeBackendBucketSignedUrlKey) StateMust() *computeBackendBucketSignedUrlKeyState {
	if cbbsuk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbbsuk.Type(), cbbsuk.LocalName()))
	}
	return cbbsuk.state
}

// ComputeBackendBucketSignedUrlKeyArgs contains the configurations for google_compute_backend_bucket_signed_url_key.
type ComputeBackendBucketSignedUrlKeyArgs struct {
	// BackendBucket: string, required
	BackendBucket terra.StringValue `hcl:"backend_bucket,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyValue: string, required
	KeyValue terra.StringValue `hcl:"key_value,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computebackendbucketsignedurlkey.Timeouts `hcl:"timeouts,block"`
}
type computeBackendBucketSignedUrlKeyAttributes struct {
	ref terra.Reference
}

// BackendBucket returns a reference to field backend_bucket of google_compute_backend_bucket_signed_url_key.
func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) BackendBucket() terra.StringValue {
	return terra.ReferenceAsString(cbbsuk.ref.Append("backend_bucket"))
}

// Id returns a reference to field id of google_compute_backend_bucket_signed_url_key.
func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbbsuk.ref.Append("id"))
}

// KeyValue returns a reference to field key_value of google_compute_backend_bucket_signed_url_key.
func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) KeyValue() terra.StringValue {
	return terra.ReferenceAsString(cbbsuk.ref.Append("key_value"))
}

// Name returns a reference to field name of google_compute_backend_bucket_signed_url_key.
func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbbsuk.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_bucket_signed_url_key.
func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbbsuk.ref.Append("project"))
}

func (cbbsuk computeBackendBucketSignedUrlKeyAttributes) Timeouts() computebackendbucketsignedurlkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computebackendbucketsignedurlkey.TimeoutsAttributes](cbbsuk.ref.Append("timeouts"))
}

type computeBackendBucketSignedUrlKeyState struct {
	BackendBucket string                                          `json:"backend_bucket"`
	Id            string                                          `json:"id"`
	KeyValue      string                                          `json:"key_value"`
	Name          string                                          `json:"name"`
	Project       string                                          `json:"project"`
	Timeouts      *computebackendbucketsignedurlkey.TimeoutsState `json:"timeouts"`
}
