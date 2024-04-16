// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package google_cloudbuild_worker_pool

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource google_cloudbuild_worker_pool.
type Resource struct {
	Name      string
	Args      Args
	state     *googleCloudbuildWorkerPoolState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (gcwp *Resource) Type() string {
	return "google_cloudbuild_worker_pool"
}

// LocalName returns the local name for [Resource].
func (gcwp *Resource) LocalName() string {
	return gcwp.Name
}

// Configuration returns the configuration (args) for [Resource].
func (gcwp *Resource) Configuration() interface{} {
	return gcwp.Args
}

// DependOn is used for other resources to depend on [Resource].
func (gcwp *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(gcwp)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (gcwp *Resource) Dependencies() terra.Dependencies {
	return gcwp.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (gcwp *Resource) LifecycleManagement() *terra.Lifecycle {
	return gcwp.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (gcwp *Resource) Attributes() googleCloudbuildWorkerPoolAttributes {
	return googleCloudbuildWorkerPoolAttributes{ref: terra.ReferenceResource(gcwp)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (gcwp *Resource) ImportState(state io.Reader) error {
	gcwp.state = &googleCloudbuildWorkerPoolState{}
	if err := json.NewDecoder(state).Decode(gcwp.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", gcwp.Type(), gcwp.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (gcwp *Resource) State() (*googleCloudbuildWorkerPoolState, bool) {
	return gcwp.state, gcwp.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (gcwp *Resource) StateMust() *googleCloudbuildWorkerPoolState {
	if gcwp.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", gcwp.Type(), gcwp.LocalName()))
	}
	return gcwp.state
}

// Args contains the configurations for google_cloudbuild_worker_pool.
type Args struct {
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
	NetworkConfig *NetworkConfig `hcl:"network_config,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
	// WorkerConfig: optional
	WorkerConfig *WorkerConfig `hcl:"worker_config,block"`
}

type googleCloudbuildWorkerPoolAttributes struct {
	ref terra.Reference
}

// Annotations returns a reference to field annotations of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Annotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcwp.ref.Append("annotations"))
}

// CreateTime returns a reference to field create_time of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) CreateTime() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("create_time"))
}

// DeleteTime returns a reference to field delete_time of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) DeleteTime() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("delete_time"))
}

// DisplayName returns a reference to field display_name of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) DisplayName() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("display_name"))
}

// EffectiveAnnotations returns a reference to field effective_annotations of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) EffectiveAnnotations() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](gcwp.ref.Append("effective_annotations"))
}

// Id returns a reference to field id of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("id"))
}

// Location returns a reference to field location of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("location"))
}

// Name returns a reference to field name of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("name"))
}

// Project returns a reference to field project of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("project"))
}

// State returns a reference to field state of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("state"))
}

// Uid returns a reference to field uid of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("uid"))
}

// UpdateTime returns a reference to field update_time of google_cloudbuild_worker_pool.
func (gcwp googleCloudbuildWorkerPoolAttributes) UpdateTime() terra.StringValue {
	return terra.ReferenceAsString(gcwp.ref.Append("update_time"))
}

func (gcwp googleCloudbuildWorkerPoolAttributes) NetworkConfig() terra.ListValue[NetworkConfigAttributes] {
	return terra.ReferenceAsList[NetworkConfigAttributes](gcwp.ref.Append("network_config"))
}

func (gcwp googleCloudbuildWorkerPoolAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](gcwp.ref.Append("timeouts"))
}

func (gcwp googleCloudbuildWorkerPoolAttributes) WorkerConfig() terra.ListValue[WorkerConfigAttributes] {
	return terra.ReferenceAsList[WorkerConfigAttributes](gcwp.ref.Append("worker_config"))
}

type googleCloudbuildWorkerPoolState struct {
	Annotations          map[string]string    `json:"annotations"`
	CreateTime           string               `json:"create_time"`
	DeleteTime           string               `json:"delete_time"`
	DisplayName          string               `json:"display_name"`
	EffectiveAnnotations map[string]string    `json:"effective_annotations"`
	Id                   string               `json:"id"`
	Location             string               `json:"location"`
	Name                 string               `json:"name"`
	Project              string               `json:"project"`
	State                string               `json:"state"`
	Uid                  string               `json:"uid"`
	UpdateTime           string               `json:"update_time"`
	NetworkConfig        []NetworkConfigState `json:"network_config"`
	Timeouts             *TimeoutsState       `json:"timeouts"`
	WorkerConfig         []WorkerConfigState  `json:"worker_config"`
}
