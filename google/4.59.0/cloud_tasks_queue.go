// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudtasksqueue "github.com/golingon/terraproviders/google/4.59.0/cloudtasksqueue"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudTasksQueue creates a new instance of [CloudTasksQueue].
func NewCloudTasksQueue(name string, args CloudTasksQueueArgs) *CloudTasksQueue {
	return &CloudTasksQueue{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudTasksQueue)(nil)

// CloudTasksQueue represents the Terraform resource google_cloud_tasks_queue.
type CloudTasksQueue struct {
	Name      string
	Args      CloudTasksQueueArgs
	state     *cloudTasksQueueState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudTasksQueue].
func (ctq *CloudTasksQueue) Type() string {
	return "google_cloud_tasks_queue"
}

// LocalName returns the local name for [CloudTasksQueue].
func (ctq *CloudTasksQueue) LocalName() string {
	return ctq.Name
}

// Configuration returns the configuration (args) for [CloudTasksQueue].
func (ctq *CloudTasksQueue) Configuration() interface{} {
	return ctq.Args
}

// DependOn is used for other resources to depend on [CloudTasksQueue].
func (ctq *CloudTasksQueue) DependOn() terra.Reference {
	return terra.ReferenceResource(ctq)
}

// Dependencies returns the list of resources [CloudTasksQueue] depends_on.
func (ctq *CloudTasksQueue) Dependencies() terra.Dependencies {
	return ctq.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudTasksQueue].
func (ctq *CloudTasksQueue) LifecycleManagement() *terra.Lifecycle {
	return ctq.Lifecycle
}

// Attributes returns the attributes for [CloudTasksQueue].
func (ctq *CloudTasksQueue) Attributes() cloudTasksQueueAttributes {
	return cloudTasksQueueAttributes{ref: terra.ReferenceResource(ctq)}
}

// ImportState imports the given attribute values into [CloudTasksQueue]'s state.
func (ctq *CloudTasksQueue) ImportState(av io.Reader) error {
	ctq.state = &cloudTasksQueueState{}
	if err := json.NewDecoder(av).Decode(ctq.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ctq.Type(), ctq.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudTasksQueue] has state.
func (ctq *CloudTasksQueue) State() (*cloudTasksQueueState, bool) {
	return ctq.state, ctq.state != nil
}

// StateMust returns the state for [CloudTasksQueue]. Panics if the state is nil.
func (ctq *CloudTasksQueue) StateMust() *cloudTasksQueueState {
	if ctq.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ctq.Type(), ctq.LocalName()))
	}
	return ctq.state
}

// CloudTasksQueueArgs contains the configurations for google_cloud_tasks_queue.
type CloudTasksQueueArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// AppEngineRoutingOverride: optional
	AppEngineRoutingOverride *cloudtasksqueue.AppEngineRoutingOverride `hcl:"app_engine_routing_override,block"`
	// RateLimits: optional
	RateLimits *cloudtasksqueue.RateLimits `hcl:"rate_limits,block"`
	// RetryConfig: optional
	RetryConfig *cloudtasksqueue.RetryConfig `hcl:"retry_config,block"`
	// StackdriverLoggingConfig: optional
	StackdriverLoggingConfig *cloudtasksqueue.StackdriverLoggingConfig `hcl:"stackdriver_logging_config,block"`
	// Timeouts: optional
	Timeouts *cloudtasksqueue.Timeouts `hcl:"timeouts,block"`
}
type cloudTasksQueueAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_cloud_tasks_queue.
func (ctq cloudTasksQueueAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ctq.ref.Append("id"))
}

// Location returns a reference to field location of google_cloud_tasks_queue.
func (ctq cloudTasksQueueAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ctq.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_tasks_queue.
func (ctq cloudTasksQueueAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ctq.ref.Append("name"))
}

// Project returns a reference to field project of google_cloud_tasks_queue.
func (ctq cloudTasksQueueAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(ctq.ref.Append("project"))
}

func (ctq cloudTasksQueueAttributes) AppEngineRoutingOverride() terra.ListValue[cloudtasksqueue.AppEngineRoutingOverrideAttributes] {
	return terra.ReferenceAsList[cloudtasksqueue.AppEngineRoutingOverrideAttributes](ctq.ref.Append("app_engine_routing_override"))
}

func (ctq cloudTasksQueueAttributes) RateLimits() terra.ListValue[cloudtasksqueue.RateLimitsAttributes] {
	return terra.ReferenceAsList[cloudtasksqueue.RateLimitsAttributes](ctq.ref.Append("rate_limits"))
}

func (ctq cloudTasksQueueAttributes) RetryConfig() terra.ListValue[cloudtasksqueue.RetryConfigAttributes] {
	return terra.ReferenceAsList[cloudtasksqueue.RetryConfigAttributes](ctq.ref.Append("retry_config"))
}

func (ctq cloudTasksQueueAttributes) StackdriverLoggingConfig() terra.ListValue[cloudtasksqueue.StackdriverLoggingConfigAttributes] {
	return terra.ReferenceAsList[cloudtasksqueue.StackdriverLoggingConfigAttributes](ctq.ref.Append("stackdriver_logging_config"))
}

func (ctq cloudTasksQueueAttributes) Timeouts() cloudtasksqueue.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudtasksqueue.TimeoutsAttributes](ctq.ref.Append("timeouts"))
}

type cloudTasksQueueState struct {
	Id                       string                                          `json:"id"`
	Location                 string                                          `json:"location"`
	Name                     string                                          `json:"name"`
	Project                  string                                          `json:"project"`
	AppEngineRoutingOverride []cloudtasksqueue.AppEngineRoutingOverrideState `json:"app_engine_routing_override"`
	RateLimits               []cloudtasksqueue.RateLimitsState               `json:"rate_limits"`
	RetryConfig              []cloudtasksqueue.RetryConfigState              `json:"retry_config"`
	StackdriverLoggingConfig []cloudtasksqueue.StackdriverLoggingConfigState `json:"stackdriver_logging_config"`
	Timeouts                 *cloudtasksqueue.TimeoutsState                  `json:"timeouts"`
}
