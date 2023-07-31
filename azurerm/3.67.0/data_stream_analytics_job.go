// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datastreamanalyticsjob "github.com/golingon/terraproviders/azurerm/3.67.0/datastreamanalyticsjob"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataStreamAnalyticsJob creates a new instance of [DataStreamAnalyticsJob].
func NewDataStreamAnalyticsJob(name string, args DataStreamAnalyticsJobArgs) *DataStreamAnalyticsJob {
	return &DataStreamAnalyticsJob{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataStreamAnalyticsJob)(nil)

// DataStreamAnalyticsJob represents the Terraform data resource azurerm_stream_analytics_job.
type DataStreamAnalyticsJob struct {
	Name string
	Args DataStreamAnalyticsJobArgs
}

// DataSource returns the Terraform object type for [DataStreamAnalyticsJob].
func (saj *DataStreamAnalyticsJob) DataSource() string {
	return "azurerm_stream_analytics_job"
}

// LocalName returns the local name for [DataStreamAnalyticsJob].
func (saj *DataStreamAnalyticsJob) LocalName() string {
	return saj.Name
}

// Configuration returns the configuration (args) for [DataStreamAnalyticsJob].
func (saj *DataStreamAnalyticsJob) Configuration() interface{} {
	return saj.Args
}

// Attributes returns the attributes for [DataStreamAnalyticsJob].
func (saj *DataStreamAnalyticsJob) Attributes() dataStreamAnalyticsJobAttributes {
	return dataStreamAnalyticsJobAttributes{ref: terra.ReferenceDataResource(saj)}
}

// DataStreamAnalyticsJobArgs contains the configurations for azurerm_stream_analytics_job.
type DataStreamAnalyticsJobArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Identity: min=0
	Identity []datastreamanalyticsjob.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datastreamanalyticsjob.Timeouts `hcl:"timeouts,block"`
}
type dataStreamAnalyticsJobAttributes struct {
	ref terra.Reference
}

// CompatibilityLevel returns a reference to field compatibility_level of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) CompatibilityLevel() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("compatibility_level"))
}

// DataLocale returns a reference to field data_locale of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) DataLocale() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("data_locale"))
}

// EventsLateArrivalMaxDelayInSeconds returns a reference to field events_late_arrival_max_delay_in_seconds of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) EventsLateArrivalMaxDelayInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("events_late_arrival_max_delay_in_seconds"))
}

// EventsOutOfOrderMaxDelayInSeconds returns a reference to field events_out_of_order_max_delay_in_seconds of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) EventsOutOfOrderMaxDelayInSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("events_out_of_order_max_delay_in_seconds"))
}

// EventsOutOfOrderPolicy returns a reference to field events_out_of_order_policy of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) EventsOutOfOrderPolicy() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("events_out_of_order_policy"))
}

// Id returns a reference to field id of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("id"))
}

// JobId returns a reference to field job_id of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) JobId() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("job_id"))
}

// LastOutputTime returns a reference to field last_output_time of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) LastOutputTime() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("last_output_time"))
}

// Location returns a reference to field location of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("name"))
}

// OutputErrorPolicy returns a reference to field output_error_policy of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) OutputErrorPolicy() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("output_error_policy"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("resource_group_name"))
}

// StartMode returns a reference to field start_mode of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) StartMode() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("start_mode"))
}

// StartTime returns a reference to field start_time of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("start_time"))
}

// StreamingUnits returns a reference to field streaming_units of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) StreamingUnits() terra.NumberValue {
	return terra.ReferenceAsNumber(saj.ref.Append("streaming_units"))
}

// TransformationQuery returns a reference to field transformation_query of azurerm_stream_analytics_job.
func (saj dataStreamAnalyticsJobAttributes) TransformationQuery() terra.StringValue {
	return terra.ReferenceAsString(saj.ref.Append("transformation_query"))
}

func (saj dataStreamAnalyticsJobAttributes) Identity() terra.ListValue[datastreamanalyticsjob.IdentityAttributes] {
	return terra.ReferenceAsList[datastreamanalyticsjob.IdentityAttributes](saj.ref.Append("identity"))
}

func (saj dataStreamAnalyticsJobAttributes) Timeouts() datastreamanalyticsjob.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datastreamanalyticsjob.TimeoutsAttributes](saj.ref.Append("timeouts"))
}
