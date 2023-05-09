// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	cloudrunv2job "github.com/golingon/terraproviders/google/4.64.0/cloudrunv2job"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudRunV2Job creates a new instance of [CloudRunV2Job].
func NewCloudRunV2Job(name string, args CloudRunV2JobArgs) *CloudRunV2Job {
	return &CloudRunV2Job{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudRunV2Job)(nil)

// CloudRunV2Job represents the Terraform resource google_cloud_run_v2_job.
type CloudRunV2Job struct {
	Name      string
	Args      CloudRunV2JobArgs
	state     *cloudRunV2JobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudRunV2Job].
func (crvj *CloudRunV2Job) Type() string {
	return "google_cloud_run_v2_job"
}

// LocalName returns the local name for [CloudRunV2Job].
func (crvj *CloudRunV2Job) LocalName() string {
	return crvj.Name
}

// Configuration returns the configuration (args) for [CloudRunV2Job].
func (crvj *CloudRunV2Job) Configuration() interface{} {
	return crvj.Args
}

// DependOn is used for other resources to depend on [CloudRunV2Job].
func (crvj *CloudRunV2Job) DependOn() terra.Reference {
	return terra.ReferenceResource(crvj)
}

// Dependencies returns the list of resources [CloudRunV2Job] depends_on.
func (crvj *CloudRunV2Job) Dependencies() terra.Dependencies {
	return crvj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudRunV2Job].
func (crvj *CloudRunV2Job) LifecycleManagement() *terra.Lifecycle {
	return crvj.Lifecycle
}

// Attributes returns the attributes for [CloudRunV2Job].
func (crvj *CloudRunV2Job) Attributes() cloudRunV2JobAttributes {
	return cloudRunV2JobAttributes{ref: terra.ReferenceResource(crvj)}
}

// ImportState imports the given attribute values into [CloudRunV2Job]'s state.
func (crvj *CloudRunV2Job) ImportState(av io.Reader) error {
	crvj.state = &cloudRunV2JobState{}
	if err := json.NewDecoder(av).Decode(crvj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crvj.Type(), crvj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudRunV2Job] has state.
func (crvj *CloudRunV2Job) State() (*cloudRunV2JobState, bool) {
	return crvj.state, crvj.state != nil
}

// StateMust returns the state for [CloudRunV2Job]. Panics if the state is nil.
func (crvj *CloudRunV2Job) StateMust() *cloudRunV2JobState {
	if crvj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crvj.Type(), crvj.LocalName()))
	}
	return crvj.state
}

// CloudRunV2JobArgs contains the configurations for google_cloud_run_v2_job.
type CloudRunV2JobArgs struct {
	// Client: string, optional
	Client terra.StringValue `hcl:"client,attr"`
	// ClientVersion: string, optional
	ClientVersion terra.StringValue `hcl:"client_version,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// LaunchStage: string, optional
	LaunchStage terra.StringValue `hcl:"launch_stage,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Conditions: min=0
	Conditions []cloudrunv2job.Conditions `hcl:"conditions,block" validate:"min=0"`
	// LatestCreatedExecution: min=0
	LatestCreatedExecution []cloudrunv2job.LatestCreatedExecution `hcl:"latest_created_execution,block" validate:"min=0"`
	// TerminalCondition: min=0
	TerminalCondition []cloudrunv2job.TerminalCondition `hcl:"terminal_condition,block" validate:"min=0"`
	// BinaryAuthorization: optional
	BinaryAuthorization *cloudrunv2job.BinaryAuthorization `hcl:"binary_authorization,block"`
	// Template: required
	Template *cloudrunv2job.Template `hcl:"template,block" validate:"required"`
	// Timeouts: optional
	Timeouts *cloudrunv2job.Timeouts `hcl:"timeouts,block"`
}
type cloudRunV2JobAttributes struct {
	ref terra.Reference
}

// Client returns a reference to field client of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Client() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("client"))
}

// ClientVersion returns a reference to field client_version of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) ClientVersion() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("client_version"))
}

// Etag returns a reference to field etag of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Etag() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("etag"))
}

// ExecutionCount returns a reference to field execution_count of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) ExecutionCount() terra.NumberValue {
	return terra.ReferenceAsNumber(crvj.ref.Append("execution_count"))
}

// Generation returns a reference to field generation of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Generation() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("generation"))
}

// Id returns a reference to field id of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("id"))
}

// Labels returns a reference to field labels of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](crvj.ref.Append("labels"))
}

// LaunchStage returns a reference to field launch_stage of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) LaunchStage() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("launch_stage"))
}

// Location returns a reference to field location of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("location"))
}

// Name returns a reference to field name of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("name"))
}

// ObservedGeneration returns a reference to field observed_generation of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) ObservedGeneration() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("observed_generation"))
}

// Project returns a reference to field project of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("project"))
}

// Reconciling returns a reference to field reconciling of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Reconciling() terra.BoolValue {
	return terra.ReferenceAsBool(crvj.ref.Append("reconciling"))
}

// Uid returns a reference to field uid of google_cloud_run_v2_job.
func (crvj cloudRunV2JobAttributes) Uid() terra.StringValue {
	return terra.ReferenceAsString(crvj.ref.Append("uid"))
}

func (crvj cloudRunV2JobAttributes) Conditions() terra.ListValue[cloudrunv2job.ConditionsAttributes] {
	return terra.ReferenceAsList[cloudrunv2job.ConditionsAttributes](crvj.ref.Append("conditions"))
}

func (crvj cloudRunV2JobAttributes) LatestCreatedExecution() terra.ListValue[cloudrunv2job.LatestCreatedExecutionAttributes] {
	return terra.ReferenceAsList[cloudrunv2job.LatestCreatedExecutionAttributes](crvj.ref.Append("latest_created_execution"))
}

func (crvj cloudRunV2JobAttributes) TerminalCondition() terra.ListValue[cloudrunv2job.TerminalConditionAttributes] {
	return terra.ReferenceAsList[cloudrunv2job.TerminalConditionAttributes](crvj.ref.Append("terminal_condition"))
}

func (crvj cloudRunV2JobAttributes) BinaryAuthorization() terra.ListValue[cloudrunv2job.BinaryAuthorizationAttributes] {
	return terra.ReferenceAsList[cloudrunv2job.BinaryAuthorizationAttributes](crvj.ref.Append("binary_authorization"))
}

func (crvj cloudRunV2JobAttributes) Template() terra.ListValue[cloudrunv2job.TemplateAttributes] {
	return terra.ReferenceAsList[cloudrunv2job.TemplateAttributes](crvj.ref.Append("template"))
}

func (crvj cloudRunV2JobAttributes) Timeouts() cloudrunv2job.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudrunv2job.TimeoutsAttributes](crvj.ref.Append("timeouts"))
}

type cloudRunV2JobState struct {
	Client                 string                                      `json:"client"`
	ClientVersion          string                                      `json:"client_version"`
	Etag                   string                                      `json:"etag"`
	ExecutionCount         float64                                     `json:"execution_count"`
	Generation             string                                      `json:"generation"`
	Id                     string                                      `json:"id"`
	Labels                 map[string]string                           `json:"labels"`
	LaunchStage            string                                      `json:"launch_stage"`
	Location               string                                      `json:"location"`
	Name                   string                                      `json:"name"`
	ObservedGeneration     string                                      `json:"observed_generation"`
	Project                string                                      `json:"project"`
	Reconciling            bool                                        `json:"reconciling"`
	Uid                    string                                      `json:"uid"`
	Conditions             []cloudrunv2job.ConditionsState             `json:"conditions"`
	LatestCreatedExecution []cloudrunv2job.LatestCreatedExecutionState `json:"latest_created_execution"`
	TerminalCondition      []cloudrunv2job.TerminalConditionState      `json:"terminal_condition"`
	BinaryAuthorization    []cloudrunv2job.BinaryAuthorizationState    `json:"binary_authorization"`
	Template               []cloudrunv2job.TemplateState               `json:"template"`
	Timeouts               *cloudrunv2job.TimeoutsState                `json:"timeouts"`
}
