// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	streamanalyticsjobschedule "github.com/golingon/terraproviders/azurerm/3.64.0/streamanalyticsjobschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewStreamAnalyticsJobSchedule creates a new instance of [StreamAnalyticsJobSchedule].
func NewStreamAnalyticsJobSchedule(name string, args StreamAnalyticsJobScheduleArgs) *StreamAnalyticsJobSchedule {
	return &StreamAnalyticsJobSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*StreamAnalyticsJobSchedule)(nil)

// StreamAnalyticsJobSchedule represents the Terraform resource azurerm_stream_analytics_job_schedule.
type StreamAnalyticsJobSchedule struct {
	Name      string
	Args      StreamAnalyticsJobScheduleArgs
	state     *streamAnalyticsJobScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) Type() string {
	return "azurerm_stream_analytics_job_schedule"
}

// LocalName returns the local name for [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) LocalName() string {
	return sajs.Name
}

// Configuration returns the configuration (args) for [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) Configuration() interface{} {
	return sajs.Args
}

// DependOn is used for other resources to depend on [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(sajs)
}

// Dependencies returns the list of resources [StreamAnalyticsJobSchedule] depends_on.
func (sajs *StreamAnalyticsJobSchedule) Dependencies() terra.Dependencies {
	return sajs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) LifecycleManagement() *terra.Lifecycle {
	return sajs.Lifecycle
}

// Attributes returns the attributes for [StreamAnalyticsJobSchedule].
func (sajs *StreamAnalyticsJobSchedule) Attributes() streamAnalyticsJobScheduleAttributes {
	return streamAnalyticsJobScheduleAttributes{ref: terra.ReferenceResource(sajs)}
}

// ImportState imports the given attribute values into [StreamAnalyticsJobSchedule]'s state.
func (sajs *StreamAnalyticsJobSchedule) ImportState(av io.Reader) error {
	sajs.state = &streamAnalyticsJobScheduleState{}
	if err := json.NewDecoder(av).Decode(sajs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sajs.Type(), sajs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [StreamAnalyticsJobSchedule] has state.
func (sajs *StreamAnalyticsJobSchedule) State() (*streamAnalyticsJobScheduleState, bool) {
	return sajs.state, sajs.state != nil
}

// StateMust returns the state for [StreamAnalyticsJobSchedule]. Panics if the state is nil.
func (sajs *StreamAnalyticsJobSchedule) StateMust() *streamAnalyticsJobScheduleState {
	if sajs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sajs.Type(), sajs.LocalName()))
	}
	return sajs.state
}

// StreamAnalyticsJobScheduleArgs contains the configurations for azurerm_stream_analytics_job_schedule.
type StreamAnalyticsJobScheduleArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StartMode: string, required
	StartMode terra.StringValue `hcl:"start_mode,attr" validate:"required"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// StreamAnalyticsJobId: string, required
	StreamAnalyticsJobId terra.StringValue `hcl:"stream_analytics_job_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *streamanalyticsjobschedule.Timeouts `hcl:"timeouts,block"`
}
type streamAnalyticsJobScheduleAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_stream_analytics_job_schedule.
func (sajs streamAnalyticsJobScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sajs.ref.Append("id"))
}

// LastOutputTime returns a reference to field last_output_time of azurerm_stream_analytics_job_schedule.
func (sajs streamAnalyticsJobScheduleAttributes) LastOutputTime() terra.StringValue {
	return terra.ReferenceAsString(sajs.ref.Append("last_output_time"))
}

// StartMode returns a reference to field start_mode of azurerm_stream_analytics_job_schedule.
func (sajs streamAnalyticsJobScheduleAttributes) StartMode() terra.StringValue {
	return terra.ReferenceAsString(sajs.ref.Append("start_mode"))
}

// StartTime returns a reference to field start_time of azurerm_stream_analytics_job_schedule.
func (sajs streamAnalyticsJobScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(sajs.ref.Append("start_time"))
}

// StreamAnalyticsJobId returns a reference to field stream_analytics_job_id of azurerm_stream_analytics_job_schedule.
func (sajs streamAnalyticsJobScheduleAttributes) StreamAnalyticsJobId() terra.StringValue {
	return terra.ReferenceAsString(sajs.ref.Append("stream_analytics_job_id"))
}

func (sajs streamAnalyticsJobScheduleAttributes) Timeouts() streamanalyticsjobschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[streamanalyticsjobschedule.TimeoutsAttributes](sajs.ref.Append("timeouts"))
}

type streamAnalyticsJobScheduleState struct {
	Id                   string                                    `json:"id"`
	LastOutputTime       string                                    `json:"last_output_time"`
	StartMode            string                                    `json:"start_mode"`
	StartTime            string                                    `json:"start_time"`
	StreamAnalyticsJobId string                                    `json:"stream_analytics_job_id"`
	Timeouts             *streamanalyticsjobschedule.TimeoutsState `json:"timeouts"`
}
