// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	datafactorytriggerschedule "github.com/golingon/terraproviders/azurerm/3.49.0/datafactorytriggerschedule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDataFactoryTriggerSchedule(name string, args DataFactoryTriggerScheduleArgs) *DataFactoryTriggerSchedule {
	return &DataFactoryTriggerSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DataFactoryTriggerSchedule)(nil)

type DataFactoryTriggerSchedule struct {
	Name  string
	Args  DataFactoryTriggerScheduleArgs
	state *dataFactoryTriggerScheduleState
}

func (dfts *DataFactoryTriggerSchedule) Type() string {
	return "azurerm_data_factory_trigger_schedule"
}

func (dfts *DataFactoryTriggerSchedule) LocalName() string {
	return dfts.Name
}

func (dfts *DataFactoryTriggerSchedule) Configuration() interface{} {
	return dfts.Args
}

func (dfts *DataFactoryTriggerSchedule) Attributes() dataFactoryTriggerScheduleAttributes {
	return dataFactoryTriggerScheduleAttributes{ref: terra.ReferenceResource(dfts)}
}

func (dfts *DataFactoryTriggerSchedule) ImportState(av io.Reader) error {
	dfts.state = &dataFactoryTriggerScheduleState{}
	if err := json.NewDecoder(av).Decode(dfts.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dfts.Type(), dfts.LocalName(), err)
	}
	return nil
}

func (dfts *DataFactoryTriggerSchedule) State() (*dataFactoryTriggerScheduleState, bool) {
	return dfts.state, dfts.state != nil
}

func (dfts *DataFactoryTriggerSchedule) StateMust() *dataFactoryTriggerScheduleState {
	if dfts.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dfts.Type(), dfts.LocalName()))
	}
	return dfts.state
}

func (dfts *DataFactoryTriggerSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(dfts)
}

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
	// DependsOn contains resources that DataFactoryTriggerSchedule depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type dataFactoryTriggerScheduleAttributes struct {
	ref terra.Reference
}

func (dfts dataFactoryTriggerScheduleAttributes) Activated() terra.BoolValue {
	return terra.ReferenceBool(dfts.ref.Append("activated"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Annotations() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](dfts.ref.Append("annotations"))
}

func (dfts dataFactoryTriggerScheduleAttributes) DataFactoryId() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("data_factory_id"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("description"))
}

func (dfts dataFactoryTriggerScheduleAttributes) EndTime() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("end_time"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Frequency() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("frequency"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("id"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Interval() terra.NumberValue {
	return terra.ReferenceNumber(dfts.ref.Append("interval"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("name"))
}

func (dfts dataFactoryTriggerScheduleAttributes) PipelineName() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("pipeline_name"))
}

func (dfts dataFactoryTriggerScheduleAttributes) PipelineParameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dfts.ref.Append("pipeline_parameters"))
}

func (dfts dataFactoryTriggerScheduleAttributes) StartTime() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("start_time"))
}

func (dfts dataFactoryTriggerScheduleAttributes) TimeZone() terra.StringValue {
	return terra.ReferenceString(dfts.ref.Append("time_zone"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Pipeline() terra.ListValue[datafactorytriggerschedule.PipelineAttributes] {
	return terra.ReferenceList[datafactorytriggerschedule.PipelineAttributes](dfts.ref.Append("pipeline"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Schedule() terra.ListValue[datafactorytriggerschedule.ScheduleAttributes] {
	return terra.ReferenceList[datafactorytriggerschedule.ScheduleAttributes](dfts.ref.Append("schedule"))
}

func (dfts dataFactoryTriggerScheduleAttributes) Timeouts() datafactorytriggerschedule.TimeoutsAttributes {
	return terra.ReferenceSingle[datafactorytriggerschedule.TimeoutsAttributes](dfts.ref.Append("timeouts"))
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
