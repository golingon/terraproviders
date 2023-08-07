// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	computebackendservicesignedurlkey "github.com/golingon/terraproviders/googlebeta/4.76.0/computebackendservicesignedurlkey"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewComputeBackendServiceSignedUrlKey creates a new instance of [ComputeBackendServiceSignedUrlKey].
func NewComputeBackendServiceSignedUrlKey(name string, args ComputeBackendServiceSignedUrlKeyArgs) *ComputeBackendServiceSignedUrlKey {
	return &ComputeBackendServiceSignedUrlKey{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ComputeBackendServiceSignedUrlKey)(nil)

// ComputeBackendServiceSignedUrlKey represents the Terraform resource google_compute_backend_service_signed_url_key.
type ComputeBackendServiceSignedUrlKey struct {
	Name      string
	Args      ComputeBackendServiceSignedUrlKeyArgs
	state     *computeBackendServiceSignedUrlKeyState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) Type() string {
	return "google_compute_backend_service_signed_url_key"
}

// LocalName returns the local name for [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) LocalName() string {
	return cbssuk.Name
}

// Configuration returns the configuration (args) for [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) Configuration() interface{} {
	return cbssuk.Args
}

// DependOn is used for other resources to depend on [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) DependOn() terra.Reference {
	return terra.ReferenceResource(cbssuk)
}

// Dependencies returns the list of resources [ComputeBackendServiceSignedUrlKey] depends_on.
func (cbssuk *ComputeBackendServiceSignedUrlKey) Dependencies() terra.Dependencies {
	return cbssuk.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) LifecycleManagement() *terra.Lifecycle {
	return cbssuk.Lifecycle
}

// Attributes returns the attributes for [ComputeBackendServiceSignedUrlKey].
func (cbssuk *ComputeBackendServiceSignedUrlKey) Attributes() computeBackendServiceSignedUrlKeyAttributes {
	return computeBackendServiceSignedUrlKeyAttributes{ref: terra.ReferenceResource(cbssuk)}
}

// ImportState imports the given attribute values into [ComputeBackendServiceSignedUrlKey]'s state.
func (cbssuk *ComputeBackendServiceSignedUrlKey) ImportState(av io.Reader) error {
	cbssuk.state = &computeBackendServiceSignedUrlKeyState{}
	if err := json.NewDecoder(av).Decode(cbssuk.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cbssuk.Type(), cbssuk.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ComputeBackendServiceSignedUrlKey] has state.
func (cbssuk *ComputeBackendServiceSignedUrlKey) State() (*computeBackendServiceSignedUrlKeyState, bool) {
	return cbssuk.state, cbssuk.state != nil
}

// StateMust returns the state for [ComputeBackendServiceSignedUrlKey]. Panics if the state is nil.
func (cbssuk *ComputeBackendServiceSignedUrlKey) StateMust() *computeBackendServiceSignedUrlKeyState {
	if cbssuk.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cbssuk.Type(), cbssuk.LocalName()))
	}
	return cbssuk.state
}

// ComputeBackendServiceSignedUrlKeyArgs contains the configurations for google_compute_backend_service_signed_url_key.
type ComputeBackendServiceSignedUrlKeyArgs struct {
	// BackendService: string, required
	BackendService terra.StringValue `hcl:"backend_service,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// KeyValue: string, required
	KeyValue terra.StringValue `hcl:"key_value,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Timeouts: optional
	Timeouts *computebackendservicesignedurlkey.Timeouts `hcl:"timeouts,block"`
}
type computeBackendServiceSignedUrlKeyAttributes struct {
	ref terra.Reference
}

// BackendService returns a reference to field backend_service of google_compute_backend_service_signed_url_key.
func (cbssuk computeBackendServiceSignedUrlKeyAttributes) BackendService() terra.StringValue {
	return terra.ReferenceAsString(cbssuk.ref.Append("backend_service"))
}

// Id returns a reference to field id of google_compute_backend_service_signed_url_key.
func (cbssuk computeBackendServiceSignedUrlKeyAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cbssuk.ref.Append("id"))
}

// KeyValue returns a reference to field key_value of google_compute_backend_service_signed_url_key.
func (cbssuk computeBackendServiceSignedUrlKeyAttributes) KeyValue() terra.StringValue {
	return terra.ReferenceAsString(cbssuk.ref.Append("key_value"))
}

// Name returns a reference to field name of google_compute_backend_service_signed_url_key.
func (cbssuk computeBackendServiceSignedUrlKeyAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cbssuk.ref.Append("name"))
}

// Project returns a reference to field project of google_compute_backend_service_signed_url_key.
func (cbssuk computeBackendServiceSignedUrlKeyAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cbssuk.ref.Append("project"))
}

func (cbssuk computeBackendServiceSignedUrlKeyAttributes) Timeouts() computebackendservicesignedurlkey.TimeoutsAttributes {
	return terra.ReferenceAsSingle[computebackendservicesignedurlkey.TimeoutsAttributes](cbssuk.ref.Append("timeouts"))
}

type computeBackendServiceSignedUrlKeyState struct {
	BackendService string                                           `json:"backend_service"`
	Id             string                                           `json:"id"`
	KeyValue       string                                           `json:"key_value"`
	Name           string                                           `json:"name"`
	Project        string                                           `json:"project"`
	Timeouts       *computebackendservicesignedurlkey.TimeoutsState `json:"timeouts"`
}
