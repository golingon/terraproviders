// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudbuildworkerpool "github.com/golingon/terraproviders/googlebeta/4.62.0/cloudbuildworkerpool"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudbuildWorkerPool creates a new instance of [CloudbuildWorkerPool].
func NewCloudbuildWorkerPool(name string, args CloudbuildWorkerPoolArgs) *CloudbuildWorkerPool {
	return &CloudbuildWorkerPool{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudbuildWorkerPool)(nil)

// CloudbuildWorkerPool represents the Terraform resource google_cloudbuild_worker_pool.
type CloudbuildWorkerPool struct {
	Name      string
	Args      CloudbuildWorkerPoolArgs
	state     *cloudbuildWorkerPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) Type() string {
	return "google_cloudbuild_worker_pool"
}

// LocalName returns the local name for [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) LocalName() string {
	return cwp.Name
}

// Configuration returns the configuration (args) for [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) Configuration() interface{} {
	return cwp.Args
}

// DependOn is used for other resources to depend on [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) DependOn() terra.Reference {
	return terra.ReferenceResource(cwp)
}

// Dependencies returns the list of resources [CloudbuildWorkerPool] depends_on.
func (cwp *CloudbuildWorkerPool) Dependencies() terra.Dependencies {
	return cwp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) LifecycleManagement() *terra.Lifecycle {
	return cwp.Lifecycle
}

// Attributes returns the attributes for [CloudbuildWorkerPool].
func (cwp *CloudbuildWorkerPool) Attributes() cloudbuildWorkerPoolAttributes {
	return cloudbuildWorkerPoolAttributes{ref: terra.ReferenceResource(cwp)}
}

// ImportState imports the given attribute values into [CloudbuildWorkerPool]'s state.
func (cwp *CloudbuildWorkerPool) ImportState(av io.Reader) error {
	cwp.state = &cloudbuildWorkerPoolState{}
	if err := json.NewDecoder(av).Decode(cwp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cwp.Type(), cwp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudbuildWorkerPool] has state.
func (cwp *CloudbuildWorkerPool) State() (*cloudbuildWorkerPoolState, bool) {
	return cwp.state, cwp.state != nil
}

// StateMust returns the state for [CloudbuildWorkerPool]. Panics if the state is nil.
func (cwp *CloudbuildWorkerPool) StateMust() *cloudbuildWorkerPoolState {
	if cwp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cwp.Type(), cwp.LocalName()))
	}
	return cwp.state
}

// CloudbuildWorkerPoolArgs contains the configurations for google_cloudbuild_worker_pool.
type CloudbuildWorkerPoolArgs struct {
	// Annotations: map of string, optional
	Annotations terra.MapValue[terra.StringValue] `hcl:"annotations,attr"`
	// DisplayName: string, optional
	DisplayName terra.StringValue `hcl:"display_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// NetworkConfig: optional
	NetworkConfig *cloudbuildworkerpool.NetworkConfig `hcl:"network_config,block"`
	// Timeouts: optional
	Timeouts *cloudbuildworkerpool.Timeouts `hcl:"timeouts,block"`
	// WorkerConfig: optional
	WorkerConfig *cloudbuildworkerpool.WorkerConfig `hcl:"worker_config,block"`
}
type cloudbuildWorkerPoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](cwp.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("delete_time"))
}

// DisplayName returns a reference to field display_name of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("display_name"))
}

// Id returns a reference to field id of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("project"))
}

// State returns a reference to field state of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("state"))
}

// Uid returns a reference to field uid of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_cloudbuild_worker_pool.
func (cwp cloudbuildWorkerPoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(cwp.ref.Append("update_time"))
}

func (cwp cloudbuildWorkerPoolAttributes) NetworkConfig() terra.ListValue[cloudbuildworkerpool.NetworkConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildworkerpool.NetworkConfigAttributes](cwp.ref.Append("network_config"))
}

func (cwp cloudbuildWorkerPoolAttributes) Timeouts() cloudbuildworkerpool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudbuildworkerpool.TimeoutsAttributes](cwp.ref.Append("timeouts"))
}

func (cwp cloudbuildWorkerPoolAttributes) WorkerConfig() terra.ListValue[cloudbuildworkerpool.WorkerConfigAttributes] {
	return terra.ReferenceAsList[cloudbuildworkerpool.WorkerConfigAttributes](cwp.ref.Append("worker_config"))
}

type cloudbuildWorkerPoolState struct {
	Annotations   map[string]string                         `json:"annotations"`
	CreateTime    string                                    `json:"create_time"`
	DeleteTime    string                                    `json:"delete_time"`
	DisplayName   string                                    `json:"display_name"`
	Id            string                                    `json:"id"`
	Location      string                                    `json:"location"`
	Name          string                                    `json:"name"`
	Project       string                                    `json:"project"`
	State         string                                    `json:"state"`
	Uid           string                                    `json:"uid"`
	UpdateTime    string                                    `json:"update_time"`
	NetworkConfig []cloudbuildworkerpool.NetworkConfigState `json:"network_config"`
	Timeouts      *cloudbuildworkerpool.TimeoutsState       `json:"timeouts"`
	WorkerConfig  []cloudbuildworkerpool.WorkerConfigState  `json:"worker_config"`
}
