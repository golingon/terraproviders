// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	macie2classificationjob "github.com/golingon/terraproviders/aws/5.9.0/macie2classificationjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewMacie2ClassificationJob creates a new instance of [Macie2ClassificationJob].
func NewMacie2ClassificationJob(name string, args Macie2ClassificationJobArgs) *Macie2ClassificationJob {
	return &Macie2ClassificationJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Macie2ClassificationJob)(nil)

// Macie2ClassificationJob represents the Terraform resource aws_macie2_classification_job.
type Macie2ClassificationJob struct {
	Name      string
	Args      Macie2ClassificationJobArgs
	state     *macie2ClassificationJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) Type() string {
	return "aws_macie2_classification_job"
}

// LocalName returns the local name for [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) LocalName() string {
	return mcj.Name
}

// Configuration returns the configuration (args) for [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) Configuration() interface{} {
	return mcj.Args
}

// DependOn is used for other resources to depend on [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) DependOn() terra.Reference {
	return terra.ReferenceResource(mcj)
}

// Dependencies returns the list of resources [Macie2ClassificationJob] depends_on.
func (mcj *Macie2ClassificationJob) Dependencies() terra.Dependencies {
	return mcj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) LifecycleManagement() *terra.Lifecycle {
	return mcj.Lifecycle
}

// Attributes returns the attributes for [Macie2ClassificationJob].
func (mcj *Macie2ClassificationJob) Attributes() macie2ClassificationJobAttributes {
	return macie2ClassificationJobAttributes{ref: terra.ReferenceResource(mcj)}
}

// ImportState imports the given attribute values into [Macie2ClassificationJob]'s state.
func (mcj *Macie2ClassificationJob) ImportState(av io.Reader) error {
	mcj.state = &macie2ClassificationJobState{}
	if err := json.NewDecoder(av).Decode(mcj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mcj.Type(), mcj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Macie2ClassificationJob] has state.
func (mcj *Macie2ClassificationJob) State() (*macie2ClassificationJobState, bool) {
	return mcj.state, mcj.state != nil
}

// StateMust returns the state for [Macie2ClassificationJob]. Panics if the state is nil.
func (mcj *Macie2ClassificationJob) StateMust() *macie2ClassificationJobState {
	if mcj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mcj.Type(), mcj.LocalName()))
	}
	return mcj.state
}

// Macie2ClassificationJobArgs contains the configurations for aws_macie2_classification_job.
type Macie2ClassificationJobArgs struct {
	// CustomDataIdentifierIds: list of string, optional
	CustomDataIdentifierIds terra.ListValue[terra.StringValue] `hcl:"custom_data_identifier_ids,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InitialRun: bool, optional
	InitialRun terra.BoolValue `hcl:"initial_run,attr"`
	// JobStatus: string, optional
	JobStatus terra.StringValue `hcl:"job_status,attr"`
	// JobType: string, required
	JobType terra.StringValue `hcl:"job_type,attr" validate:"required"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// SamplingPercentage: number, optional
	SamplingPercentage terra.NumberValue `hcl:"sampling_percentage,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// UserPausedDetails: min=0
	UserPausedDetails []macie2classificationjob.UserPausedDetails `hcl:"user_paused_details,block" validate:"min=0"`
	// S3JobDefinition: required
	S3JobDefinition *macie2classificationjob.S3JobDefinition `hcl:"s3_job_definition,block" validate:"required"`
	// ScheduleFrequency: optional
	ScheduleFrequency *macie2classificationjob.ScheduleFrequency `hcl:"schedule_frequency,block"`
}
type macie2ClassificationJobAttributes struct {
	ref terra.Reference
}

// CreatedAt returns a reference to field created_at of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("created_at"))
}

// CustomDataIdentifierIds returns a reference to field custom_data_identifier_ids of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) CustomDataIdentifierIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](mcj.ref.Append("custom_data_identifier_ids"))
}

// Description returns a reference to field description of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("description"))
}

// Id returns a reference to field id of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("id"))
}

// InitialRun returns a reference to field initial_run of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) InitialRun() terra.BoolValue {
	return terra.ReferenceAsBool(mcj.ref.Append("initial_run"))
}

// JobArn returns a reference to field job_arn of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) JobArn() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("job_arn"))
}

// JobId returns a reference to field job_id of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("job_id"))
}

// JobStatus returns a reference to field job_status of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) JobStatus() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("job_status"))
}

// JobType returns a reference to field job_type of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) JobType() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("job_type"))
}

// Name returns a reference to field name of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(mcj.ref.Append("name_prefix"))
}

// SamplingPercentage returns a reference to field sampling_percentage of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) SamplingPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(mcj.ref.Append("sampling_percentage"))
}

// Tags returns a reference to field tags of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcj.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_macie2_classification_job.
func (mcj macie2ClassificationJobAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mcj.ref.Append("tags_all"))
}

func (mcj macie2ClassificationJobAttributes) UserPausedDetails() terra.ListValue[macie2classificationjob.UserPausedDetailsAttributes] {
	return terra.ReferenceAsList[macie2classificationjob.UserPausedDetailsAttributes](mcj.ref.Append("user_paused_details"))
}

func (mcj macie2ClassificationJobAttributes) S3JobDefinition() terra.ListValue[macie2classificationjob.S3JobDefinitionAttributes] {
	return terra.ReferenceAsList[macie2classificationjob.S3JobDefinitionAttributes](mcj.ref.Append("s3_job_definition"))
}

func (mcj macie2ClassificationJobAttributes) ScheduleFrequency() terra.ListValue[macie2classificationjob.ScheduleFrequencyAttributes] {
	return terra.ReferenceAsList[macie2classificationjob.ScheduleFrequencyAttributes](mcj.ref.Append("schedule_frequency"))
}

type macie2ClassificationJobState struct {
	CreatedAt               string                                           `json:"created_at"`
	CustomDataIdentifierIds []string                                         `json:"custom_data_identifier_ids"`
	Description             string                                           `json:"description"`
	Id                      string                                           `json:"id"`
	InitialRun              bool                                             `json:"initial_run"`
	JobArn                  string                                           `json:"job_arn"`
	JobId                   string                                           `json:"job_id"`
	JobStatus               string                                           `json:"job_status"`
	JobType                 string                                           `json:"job_type"`
	Name                    string                                           `json:"name"`
	NamePrefix              string                                           `json:"name_prefix"`
	SamplingPercentage      float64                                          `json:"sampling_percentage"`
	Tags                    map[string]string                                `json:"tags"`
	TagsAll                 map[string]string                                `json:"tags_all"`
	UserPausedDetails       []macie2classificationjob.UserPausedDetailsState `json:"user_paused_details"`
	S3JobDefinition         []macie2classificationjob.S3JobDefinitionState   `json:"s3_job_definition"`
	ScheduleFrequency       []macie2classificationjob.ScheduleFrequencyState `json:"schedule_frequency"`
}
