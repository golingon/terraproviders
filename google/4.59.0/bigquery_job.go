// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package google

import (
	"encoding/json"
	"fmt"
	bigqueryjob "github.com/golingon/terraproviders/google/4.59.0/bigqueryjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewBigqueryJob creates a new instance of [BigqueryJob].
func NewBigqueryJob(name string, args BigqueryJobArgs) *BigqueryJob {
	return &BigqueryJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*BigqueryJob)(nil)

// BigqueryJob represents the Terraform resource google_bigquery_job.
type BigqueryJob struct {
	Name      string
	Args      BigqueryJobArgs
	state     *bigqueryJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [BigqueryJob].
func (bj *BigqueryJob) Type() string {
	return "google_bigquery_job"
}

// LocalName returns the local name for [BigqueryJob].
func (bj *BigqueryJob) LocalName() string {
	return bj.Name
}

// Configuration returns the configuration (args) for [BigqueryJob].
func (bj *BigqueryJob) Configuration() interface{} {
	return bj.Args
}

// DependOn is used for other resources to depend on [BigqueryJob].
func (bj *BigqueryJob) DependOn() terra.Reference {
	return terra.ReferenceResource(bj)
}

// Dependencies returns the list of resources [BigqueryJob] depends_on.
func (bj *BigqueryJob) Dependencies() terra.Dependencies {
	return bj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [BigqueryJob].
func (bj *BigqueryJob) LifecycleManagement() *terra.Lifecycle {
	return bj.Lifecycle
}

// Attributes returns the attributes for [BigqueryJob].
func (bj *BigqueryJob) Attributes() bigqueryJobAttributes {
	return bigqueryJobAttributes{ref: terra.ReferenceResource(bj)}
}

// ImportState imports the given attribute values into [BigqueryJob]'s state.
func (bj *BigqueryJob) ImportState(av io.Reader) error {
	bj.state = &bigqueryJobState{}
	if err := json.NewDecoder(av).Decode(bj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", bj.Type(), bj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [BigqueryJob] has state.
func (bj *BigqueryJob) State() (*bigqueryJobState, bool) {
	return bj.state, bj.state != nil
}

// StateMust returns the state for [BigqueryJob]. Panics if the state is nil.
func (bj *BigqueryJob) StateMust() *bigqueryJobState {
	if bj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", bj.Type(), bj.LocalName()))
	}
	return bj.state
}

// BigqueryJobArgs contains the configurations for google_bigquery_job.
type BigqueryJobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// JobId: string, required
	JobId terra.StringValue `hcl:"job_id,attr" validate:"required"`
	// JobTimeoutMs: string, optional
	JobTimeoutMs terra.StringValue `hcl:"job_timeout_ms,attr"`
	// Labels: map of string, optional
	Labels terra.MapValue[terra.StringValue] `hcl:"labels,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// Project: string, optional
	Project terra.StringValue `hcl:"project,attr"`
	// Status: min=0
	Status []bigqueryjob.Status `hcl:"status,block" validate:"min=0"`
	// Copy: optional
	Copy *bigqueryjob.Copy `hcl:"copy,block"`
	// Extract: optional
	Extract *bigqueryjob.Extract `hcl:"extract,block"`
	// Load: optional
	Load *bigqueryjob.Load `hcl:"load,block"`
	// Query: optional
	Query *bigqueryjob.Query `hcl:"query,block"`
	// Timeouts: optional
	Timeouts *bigqueryjob.Timeouts `hcl:"timeouts,block"`
}
type bigqueryJobAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of google_bigquery_job.
func (bj bigqueryJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("id"))
}

// JobId returns a reference to field job_id of google_bigquery_job.
func (bj bigqueryJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("job_id"))
}

// JobTimeoutMs returns a reference to field job_timeout_ms of google_bigquery_job.
func (bj bigqueryJobAttributes) JobTimeoutMs() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("job_timeout_ms"))
}

// JobType returns a reference to field job_type of google_bigquery_job.
func (bj bigqueryJobAttributes) JobType() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("job_type"))
}

// Labels returns a reference to field labels of google_bigquery_job.
func (bj bigqueryJobAttributes) Labels() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](bj.ref.Append("labels"))
}

// Location returns a reference to field location of google_bigquery_job.
func (bj bigqueryJobAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("location"))
}

// Project returns a reference to field project of google_bigquery_job.
func (bj bigqueryJobAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("project"))
}

// UserEmail returns a reference to field user_email of google_bigquery_job.
func (bj bigqueryJobAttributes) UserEmail() terra.StringValue {
	return terra.ReferenceAsString(bj.ref.Append("user_email"))
}

func (bj bigqueryJobAttributes) Status() terra.ListValue[bigqueryjob.StatusAttributes] {
	return terra.ReferenceAsList[bigqueryjob.StatusAttributes](bj.ref.Append("status"))
}

func (bj bigqueryJobAttributes) Copy() terra.ListValue[bigqueryjob.CopyAttributes] {
	return terra.ReferenceAsList[bigqueryjob.CopyAttributes](bj.ref.Append("copy"))
}

func (bj bigqueryJobAttributes) Extract() terra.ListValue[bigqueryjob.ExtractAttributes] {
	return terra.ReferenceAsList[bigqueryjob.ExtractAttributes](bj.ref.Append("extract"))
}

func (bj bigqueryJobAttributes) Load() terra.ListValue[bigqueryjob.LoadAttributes] {
	return terra.ReferenceAsList[bigqueryjob.LoadAttributes](bj.ref.Append("load"))
}

func (bj bigqueryJobAttributes) Query() terra.ListValue[bigqueryjob.QueryAttributes] {
	return terra.ReferenceAsList[bigqueryjob.QueryAttributes](bj.ref.Append("query"))
}

func (bj bigqueryJobAttributes) Timeouts() bigqueryjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[bigqueryjob.TimeoutsAttributes](bj.ref.Append("timeouts"))
}

type bigqueryJobState struct {
	Id           string                     `json:"id"`
	JobId        string                     `json:"job_id"`
	JobTimeoutMs string                     `json:"job_timeout_ms"`
	JobType      string                     `json:"job_type"`
	Labels       map[string]string          `json:"labels"`
	Location     string                     `json:"location"`
	Project      string                     `json:"project"`
	UserEmail    string                     `json:"user_email"`
	Status       []bigqueryjob.StatusState  `json:"status"`
	Copy         []bigqueryjob.CopyState    `json:"copy"`
	Extract      []bigqueryjob.ExtractState `json:"extract"`
	Load         []bigqueryjob.LoadState    `json:"load"`
	Query        []bigqueryjob.QueryState   `json:"query"`
	Timeouts     *bigqueryjob.TimeoutsState `json:"timeouts"`
}
