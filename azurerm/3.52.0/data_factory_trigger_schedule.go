// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorytriggerschedule "github.com/golingon/terraproviders/azurerm/3.52.0/datafactorytriggerschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDataFactoryTriggerSchedule creates a new instance of [DataFactoryTriggerSchedule].
func NewDataFactoryTriggerSchedule(name string, args DataFactoryTriggerScheduleArgs) *DataFactoryTriggerSchedule {
	return &DataFactoryTriggerSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryTriggerSchedule)(nil)

// DataFactoryTriggerSchedule represents the Terraform resource azurerm_data_factory_trigger_schedule.
type DataFactoryTriggerSchedule struct {
	Name      string
	Args      DataFactoryTriggerScheduleArgs
	state     *dataFactoryTriggerScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) Type() string {
	return "azurerm_data_factory_trigger_schedule"
}

// LocalName returns the local name for [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) LocalName() string {
	return dfts.Name
}

// Configuration returns the configuration (args) for [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) Configuration() interface{} {
	return dfts.Args
}

// DependOn is used for other resources to depend on [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(dfts)
}

// Dependencies returns the list of resources [DataFactoryTriggerSchedule] depends_on.
func (dfts *DataFactoryTriggerSchedule) Dependencies() terra.Dependencies {
	return dfts.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) LifecycleManagement() *terra.Lifecycle {
	return dfts.Lifecycle
}

// Attributes returns the attributes for [DataFactoryTriggerSchedule].
func (dfts *DataFactoryTriggerSchedule) Attributes() dataFactoryTriggerScheduleAttributes {
	return dataFactoryTriggerScheduleAttributes{ref: terra.ReferenceResource(dfts)}
}

// ImportState imports the given attribute values into [DataFactoryTriggerSchedule]'s state.
func (dfts *DataFactoryTriggerSchedule) ImportState(av io.Reader) error {
	dfts.state = &dataFactoryTriggerScheduleState{}
	if err := json.NewDecoder(av).Decode(dfts.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfts.Type(), dfts.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DataFactoryTriggerSchedule] has state.
func (dfts *DataFactoryTriggerSchedule) State() (*dataFactoryTriggerScheduleState, bool) {
	return dfts.state, dfts.state != nil
}

// StateMust returns the state for [DataFactoryTriggerSchedule]. Panics if the state is nil.
func (dfts *DataFactoryTriggerSchedule) StateMust() *dataFactoryTriggerScheduleState {
	if dfts.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfts.Type(), dfts.LocalName()))
	}
	return dfts.state
}

// DataFactoryTriggerScheduleArgs contains the configurations for azurerm_data_factory_trigger_schedule.
type DataFactoryTriggerScheduleArgs struct {
	// Activated: bool, optional
	Activated terra.BoolValue `hcl:"activated,attr"`
	// Annotations: list of string, optional
	Annotations terra.ListValue[terra.StringValue] `hcl:"annotations,attr"`
	// DataFactoryId: string, required
	DataFactoryId terra.StringValue `hcl:"data_factory_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EndTime: string, optional
	EndTime terra.StringValue `hcl:"end_time,attr"`
	// Frequency: string, optional
	Frequency terra.StringValue `hcl:"frequency,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Interval: number, optional
	Interval terra.NumberValue `hcl:"interval,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PipelineName: string, optional
	PipelineName terra.StringValue `hcl:"pipeline_name,attr"`
	// PipelineParameters: map of string, optional
	PipelineParameters terra.MapValue[terra.StringValue] `hcl:"pipeline_parameters,attr"`
	// StartTime: string, optional
	StartTime terra.StringValue `hcl:"start_time,attr"`
	// TimeZone: string, optional
	TimeZone terra.StringValue `hcl:"time_zone,attr"`
	// Pipeline: min=0
	Pipeline []datafactorytriggerschedule.Pipeline `hcl:"pipeline,block" validate:"min=0"`
	// Schedule: optional
	Schedule *datafactorytriggerschedule.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *datafactorytriggerschedule.Timeouts `hcl:"timeouts,block"`
}
type dataFactoryTriggerScheduleAttributes struct {
	ref terra.Reference
}

// Activated returns a reference to field activated of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Activated() terra.BoolValue {
	return terra.ReferenceAsBool(dfts.ref.Append("activated"))
}

// Annotations returns a reference to field annotations of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](dfts.ref.Append("annotations"))
}

// DataFactoryId returns a reference to field data_factory_id of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("data_factory_id"))
}

// Description returns a reference to field description of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("description"))
}

// EndTime returns a reference to field end_time of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) EndTime() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("end_time"))
}

// Frequency returns a reference to field frequency of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Frequency() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("frequency"))
}

// Id returns a reference to field id of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("id"))
}

// Interval returns a reference to field interval of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Interval() terra.NumberValue {
	return terra.ReferenceAsNumber(dfts.ref.Append("interval"))
}

// Name returns a reference to field name of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("name"))
}

// PipelineName returns a reference to field pipeline_name of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) PipelineName() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("pipeline_name"))
}

// PipelineParameters returns a reference to field pipeline_parameters of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) PipelineParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](dfts.ref.Append("pipeline_parameters"))
}

// StartTime returns a reference to field start_time of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("start_time"))
}

// TimeZone returns a reference to field time_zone of azurerm_data_factory_trigger_schedule.
func (dfts dataFactoryTriggerScheduleAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceAsString(dfts.ref.Append("time_zone"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Pipeline() terra.ListValue[datafactorytriggerschedule.PipelineAttributes] {
	return terra.ReferenceAsList[datafactorytriggerschedule.PipelineAttributes](dfts.ref.Append("pipeline"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Schedule() terra.ListValue[datafactorytriggerschedule.ScheduleAttributes] {
	return terra.ReferenceAsList[datafactorytriggerschedule.ScheduleAttributes](dfts.ref.Append("schedule"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Timeouts() datafactorytriggerschedule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datafactorytriggerschedule.TimeoutsAttributes](dfts.ref.Append("timeouts"))
}

type dataFactoryTriggerScheduleState struct {
	Activated          bool                                       `json:"activated"`
	Annotations        []string                                   `json:"annotations"`
	DataFactoryId      string                                     `json:"data_factory_id"`
	Description        string                                     `json:"description"`
	EndTime            string                                     `json:"end_time"`
	Frequency          string                                     `json:"frequency"`
	Id                 string                                     `json:"id"`
	Interval           float64                                    `json:"interval"`
	Name               string                                     `json:"name"`
	PipelineName       string                                     `json:"pipeline_name"`
	PipelineParameters map[string]string                          `json:"pipeline_parameters"`
	StartTime          string                                     `json:"start_time"`
	TimeZone           string                                     `json:"time_zone"`
	Pipeline           []datafactorytriggerschedule.PipelineState `json:"pipeline"`
	Schedule           []datafactorytriggerschedule.ScheduleState `json:"schedule"`
	Timeouts           *datafactorytriggerschedule.TimeoutsState  `json:"timeouts"`
}
