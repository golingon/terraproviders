// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package googlebeta

import (
	"encoding/json"
	"fmt"
	cloudschedulerjob "github.com/golingon/terraproviders/googlebeta/4.62.0/cloudschedulerjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudSchedulerJob creates a new instance of [CloudSchedulerJob].
func NewCloudSchedulerJob(name string, args CloudSchedulerJobArgs) *CloudSchedulerJob {
	return &CloudSchedulerJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudSchedulerJob)(nil)

// CloudSchedulerJob represents the Terraform resource google_cloud_scheduler_job.
type CloudSchedulerJob struct {
	Name      string
	Args      CloudSchedulerJobArgs
	state     *cloudSchedulerJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudSchedulerJob].
func (csj *CloudSchedulerJob) Type() string {
	return "google_cloud_scheduler_job"
}

// LocalName returns the local name for [CloudSchedulerJob].
func (csj *CloudSchedulerJob) LocalName() string {
	return csj.Name
}

// Configuration returns the configuration (args) for [CloudSchedulerJob].
func (csj *CloudSchedulerJob) Configuration() interface{} {
	return csj.Args
}

// DependOn is used for other resources to depend on [CloudSchedulerJob].
func (csj *CloudSchedulerJob) DependOn() terra.Reference {
	return terra.ReferenceResource(csj)
}

// Dependencies returns the list of resources [CloudSchedulerJob] depends_on.
func (csj *CloudSchedulerJob) Dependencies() terra.Dependencies {
	return csj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudSchedulerJob].
func (csj *CloudSchedulerJob) LifecycleManagement() *terra.Lifecycle {
	return csj.Lifecycle
}

// Attributes returns the attributes for [CloudSchedulerJob].
func (csj *CloudSchedulerJob) Attributes() cloudSchedulerJobAttributes {
	return cloudSchedulerJobAttributes{ref: terra.ReferenceResource(csj)}
}

// ImportState imports the given attribute values into [CloudSchedulerJob]'s state.
func (csj *CloudSchedulerJob) ImportState(av io.Reader) error {
	csj.state = &cloudSchedulerJobState{}
	if err := json.NewDecoder(av).Decode(csj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", csj.Type(), csj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudSchedulerJob] has state.
func (csj *CloudSchedulerJob) State() (*cloudSchedulerJobState, bool) {
	return csj.state, csj.state != nil
}

// StateMust returns the state for [CloudSchedulerJob]. Panics if the state is nil.
func (csj *CloudSchedulerJob) StateMust() *cloudSchedulerJobState {
	if csj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", csj.Type(), csj.LocalName()))
	}
	return csj.state
}

// CloudSchedulerJobArgs contains the configurations for google_cloud_scheduler_job.
type CloudSchedulerJobArgs struct {
	// AttemptDeadline: string, optional
	AttemptDeadline terra.StringValue `hcl:"attempt_deadline,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Paused: bool, optional
	Paused terra.BoolValue `hcl:"paused,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Region: string, optional
	Region terra.StringValue `hcl:"region,attr"`
	// Schedule: string, optional
	Schedule terra.StringValue `hcl:"schedule,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// AppEngineHttpTarget: optional
	AppEngineHttpTarget *cloudschedulerjob.AppEngineHttpTarget `hcl:"app_engine_http_target,block"`
	// HttpTarget: optional
	HttpTarget *cloudschedulerjob.HttpTarget `hcl:"http_target,block"`
	// PubsubTarget: optional
	PubsubTarget *cloudschedulerjob.PubsubTarget `hcl:"pubsub_target,block"`
	// RetryConfig: optional
	RetryConfig *cloudschedulerjob.RetryConfig `hcl:"retry_config,block"`
	// Timeouts: optional
	Timeouts *cloudschedulerjob.Timeouts `hcl:"timeouts,block"`
}
type cloudSchedulerJobAttributes struct {
	ref terra.Reference
}

// AttemptDeadline returns a reference to field attempt_deadline of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) AttemptDeadline() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("attempt_deadline"))
}

// Description returns a reference to field description of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("description"))
}

// Id returns a reference to field id of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("id"))
}

// Name returns a reference to field name of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("name"))
}

// Paused returns a reference to field paused of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Paused() terra.BoolValue {
	return terra.ReferenceAsBool(csj.ref.Append("paused"))
}

// Project returns a reference to field project of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("project"))
}

// Region returns a reference to field region of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Region() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("region"))
}

// Schedule returns a reference to field schedule of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) Schedule() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("schedule"))
}

// State returns a reference to field state of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("state"))
}

// TimeZone returns a reference to field time_zone of google_cloud_scheduler_job.
func (csj cloudSchedulerJobAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(csj.ref.Append("time_zone"))
}

func (csj cloudSchedulerJobAttributes) AppEngineHttpTarget() terra.ListValue[cloudschedulerjob.AppEngineHttpTargetAttributes] {
	return terra.ReferenceAsList[cloudschedulerjob.AppEngineHttpTargetAttributes](csj.ref.Append("app_engine_http_target"))
}

func (csj cloudSchedulerJobAttributes) HttpTarget() terra.ListValue[cloudschedulerjob.HttpTargetAttributes] {
	return terra.ReferenceAsList[cloudschedulerjob.HttpTargetAttributes](csj.ref.Append("http_target"))
}

func (csj cloudSchedulerJobAttributes) PubsubTarget() terra.ListValue[cloudschedulerjob.PubsubTargetAttributes] {
	return terra.ReferenceAsList[cloudschedulerjob.PubsubTargetAttributes](csj.ref.Append("pubsub_target"))
}

func (csj cloudSchedulerJobAttributes) RetryConfig() terra.ListValue[cloudschedulerjob.RetryConfigAttributes] {
	return terra.ReferenceAsList[cloudschedulerjob.RetryConfigAttributes](csj.ref.Append("retry_config"))
}

func (csj cloudSchedulerJobAttributes) Timeouts() cloudschedulerjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudschedulerjob.TimeoutsAttributes](csj.ref.Append("timeouts"))
}

type cloudSchedulerJobState struct {
	AttemptDeadline     string                                       `json:"attempt_deadline"`
	Description         string                                       `json:"description"`
	Id                  string                                       `json:"id"`
	Name                string                                       `json:"name"`
	Paused              bool                                         `json:"paused"`
	Project             string                                       `json:"project"`
	Region              string                                       `json:"region"`
	Schedule            string                                       `json:"schedule"`
	State               string                                       `json:"state"`
	TimeZone            string                                       `json:"time_zone"`
	AppEngineHttpTarget []cloudschedulerjob.AppEngineHttpTargetState `json:"app_engine_http_target"`
	HttpTarget          []cloudschedulerjob.HttpTargetState          `json:"http_target"`
	PubsubTarget        []cloudschedulerjob.PubsubTargetState        `json:"pubsub_target"`
	RetryConfig         []cloudschedulerjob.RetryConfigState         `json:"retry_config"`
	Timeouts            *cloudschedulerjob.TimeoutsState             `json:"timeouts"`
}
