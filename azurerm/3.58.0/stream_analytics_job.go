// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsjob "github.com/golingon/terraproviders/azurerm/3.58.0/streamanalyticsjob"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsJob creates a new instance of [StreamAnalyticsJob].
func NewStreamAnalyticsJob(name string, args StreamAnalyticsJobArgs) *StreamAnalyticsJob {
	return &StreamAnalyticsJob{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsJob)(nil)

// StreamAnalyticsJob represents the Terraform resource azurerm_stream_analytics_job.
type StreamAnalyticsJob struct {
	Name      string
	Args      StreamAnalyticsJobArgs
	state     *streamAnalyticsJobState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) Type() string {
	return "azurerm_stream_analytics_job"
}

// LocalName returns the local name for [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) LocalName() string {
	return saj.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) Configuration() interface{} {
	return saj.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) DependOn() terra.Reference {
	return terra.ReferenceResource(saj)
}

// Dependencies returns the list of resources [StreamAnalyticsJob] depends_on.
func (saj *StreamAnalyticsJob) Dependencies() terra.Dependencies {
	return saj.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) LifecycleManagement() *terra.Lifecycle {
	return saj.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsJob].
func (saj *StreamAnalyticsJob) Attributes() streamAnalyticsJobAttributes {
	return streamAnalyticsJobAttributes{ref: terra.ReferenceResource(saj)}
}

// ImportState imports the given attribute values into [StreamAnalyticsJob]'s state.
func (saj *StreamAnalyticsJob) ImportState(av io.Reader) error {
	saj.state = &streamAnalyticsJobState{}
	if err := json.NewDecoder(av).Decode(saj.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", saj.Type(), saj.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsJob] has state.
func (saj *StreamAnalyticsJob) State() (*streamAnalyticsJobState, bool) {
	return saj.state, saj.state != nil
}

// StateMust returns the state for [StreamAnalyticsJob]. Panics if the state is nil.
func (saj *StreamAnalyticsJob) StateMust() *streamAnalyticsJobState {
	if saj.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", saj.Type(), saj.LocalName()))
	}
	return saj.state
}

// StreamAnalyticsJobArgs contains the configurations for azurerm_stream_analytics_job.
type StreamAnalyticsJobArgs struct {
	// CompatibilityLevel: string, optional
	CompatibilityLevel terra.StringValue `hcl:"compatibility_level,attr"`
	// ContentStoragePolicy: string, optional
	ContentStoragePolicy terra.StringValue `hcl:"content_storage_policy,attr"`
	// DataLocale: string, optional
	DataLocale terra.StringValue `hcl:"data_locale,attr"`
	// EventsLateArrivalMaxDelayInSeconds: number, optional
	EventsLateArrivalMaxDelayInSeconds terra.NumberValue `hcl:"events_late_arrival_max_delay_in_seconds,attr"`
	// EventsOutOfOrderMaxDelayInSeconds: number, optional
	EventsOutOfOrderMaxDelayInSeconds terra.NumberValue `hcl:"events_out_of_order_max_delay_in_seconds,attr"`
	// EventsOutOfOrderPolicy: string, optional
	EventsOutOfOrderPolicy terra.StringValue `hcl:"events_out_of_order_policy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// OutputErrorPolicy: string, optional
	OutputErrorPolicy terra.StringValue `hcl:"output_error_policy,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// StreamAnalyticsClusterId: string, optional
	StreamAnalyticsClusterId terra.StringValue `hcl:"stream_analytics_cluster_id,attr"`
	// StreamingUnits: number, optional
	StreamingUnits terra.NumberValue `hcl:"streaming_units,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TransformationQuery: string, required
	TransformationQuery terra.StringValue `hcl:"transformation_query,attr" validate:"required"`
	// Type: string, optional
	Type terra.StringValue `hcl:"type,attr"`
	// Identity: optional
	Identity *streamanalyticsjob.Identity `hcl:"identity,block"`
	// JobStorageAccount: min=0
	JobStorageAccount []streamanalyticsjob.JobStorageAccount `hcl:"job_storage_account,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *streamanalyticsjob.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsJobAttributes struct {
	ref terra.Reference
}

// CompatibilityLevel returns a reference to field compatibility_level of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) CompatibilityLevel() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("compatibility_level"))
}

// ContentStoragePolicy returns a reference to field content_storage_policy of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) ContentStoragePolicy() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("content_storage_policy"))
}

// DataLocale returns a reference to field data_locale of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) DataLocale() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("data_locale"))
}

// EventsLateArrivalMaxDelayInSeconds returns a reference to field events_late_arrival_max_delay_in_seconds of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) EventsLateArrivalMaxDelayInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("events_late_arrival_max_delay_in_seconds"))
}

// EventsOutOfOrderMaxDelayInSeconds returns a reference to field events_out_of_order_max_delay_in_seconds of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) EventsOutOfOrderMaxDelayInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("events_out_of_order_max_delay_in_seconds"))
}

// EventsOutOfOrderPolicy returns a reference to field events_out_of_order_policy of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) EventsOutOfOrderPolicy() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("events_out_of_order_policy"))
}

// Id returns a reference to field id of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("id"))
}

// JobId returns a reference to field job_id of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("job_id"))
}

// Location returns a reference to field location of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("name"))
}

// OutputErrorPolicy returns a reference to field output_error_policy of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) OutputErrorPolicy() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("output_error_policy"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("resource_group_name"))
}

// StreamAnalyticsClusterId returns a reference to field stream_analytics_cluster_id of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) StreamAnalyticsClusterId() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("stream_analytics_cluster_id"))
}

// StreamingUnits returns a reference to field streaming_units of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) StreamingUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("streaming_units"))
}

// Tags returns a reference to field tags of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](saj.ref.Append("tags"))
}

// TransformationQuery returns a reference to field transformation_query of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) TransformationQuery() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("transformation_query"))
}

// Type returns a reference to field type of azurerm_stream_analytics_job.
func (saj streamAnalyticsJobAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("type"))
}

func (saj streamAnalyticsJobAttributes) Identity() terra.ListValue[streamanalyticsjob.IdentityAttributes] {
	return terra.ReferenceAsList[streamanalyticsjob.IdentityAttributes](saj.ref.Append("identity"))
}

func (saj streamAnalyticsJobAttributes) JobStorageAccount() terra.ListValue[streamanalyticsjob.JobStorageAccountAttributes] {
	return terra.ReferenceAsList[streamanalyticsjob.JobStorageAccountAttributes](saj.ref.Append("job_storage_account"))
}

func (saj streamAnalyticsJobAttributes) Timeouts() streamanalyticsjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsjob.TimeoutsAttributes](saj.ref.Append("timeouts"))
}

type streamAnalyticsJobState struct {
	CompatibilityLevel                 string                                      `json:"compatibility_level"`
	ContentStoragePolicy               string                                      `json:"content_storage_policy"`
	DataLocale                         string                                      `json:"data_locale"`
	EventsLateArrivalMaxDelayInSeconds float64                                     `json:"events_late_arrival_max_delay_in_seconds"`
	EventsOutOfOrderMaxDelayInSeconds  float64                                     `json:"events_out_of_order_max_delay_in_seconds"`
	EventsOutOfOrderPolicy             string                                      `json:"events_out_of_order_policy"`
	Id                                 string                                      `json:"id"`
	JobId                              string                                      `json:"job_id"`
	Location                           string                                      `json:"location"`
	Name                               string                                      `json:"name"`
	OutputErrorPolicy                  string                                      `json:"output_error_policy"`
	ResourceGroupName                  string                                      `json:"resource_group_name"`
	StreamAnalyticsClusterId           string                                      `json:"stream_analytics_cluster_id"`
	StreamingUnits                     float64                                     `json:"streaming_units"`
	Tags                               map[string]string                           `json:"tags"`
	TransformationQuery                string                                      `json:"transformation_query"`
	Type                               string                                      `json:"type"`
	Identity                           []streamanalyticsjob.IdentityState          `json:"identity"`
	JobStorageAccount                  []streamanalyticsjob.JobStorageAccountState `json:"job_storage_account"`
	Timeouts                           *streamanalyticsjob.TimeoutsState           `json:"timeouts"`
}
