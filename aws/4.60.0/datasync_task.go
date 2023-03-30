// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	datasynctask "github.com/golingon/terraproviders/aws/4.60.0/datasynctask"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewDatasyncTask(name string, args DatasyncTaskArgs) *DatasyncTask {
	return &DatasyncTask{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DatasyncTask)(nil)

type DatasyncTask struct {
	Name  string
	Args  DatasyncTaskArgs
	state *datasyncTaskState
}

func (dt *DatasyncTask) Type() string {
	return "aws_datasync_task"
}

func (dt *DatasyncTask) LocalName() string {
	return dt.Name
}

func (dt *DatasyncTask) Configuration() interface{} {
	return dt.Args
}

func (dt *DatasyncTask) Attributes() datasyncTaskAttributes {
	return datasyncTaskAttributes{ref: terra.ReferenceResource(dt)}
}

func (dt *DatasyncTask) ImportState(av io.Reader) error {
	dt.state = &datasyncTaskState{}
	if err := json.NewDecoder(av).Decode(dt.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dt.Type(), dt.LocalName(), err)
	}
	return nil
}

func (dt *DatasyncTask) State() (*datasyncTaskState, bool) {
	return dt.state, dt.state != nil
}

func (dt *DatasyncTask) StateMust() *datasyncTaskState {
	if dt.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dt.Type(), dt.LocalName()))
	}
	return dt.state
}

func (dt *DatasyncTask) DependOn() terra.Reference {
	return terra.ReferenceResource(dt)
}

type DatasyncTaskArgs struct {
	// CloudwatchLogGroupArn: string, optional
	CloudwatchLogGroupArn terra.StringValue `hcl:"cloudwatch_log_group_arn,attr"`
	// DestinationLocationArn: string, required
	DestinationLocationArn terra.StringValue `hcl:"destination_location_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// SourceLocationArn: string, required
	SourceLocationArn terra.StringValue `hcl:"source_location_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Excludes: optional
	Excludes *datasynctask.Excludes `hcl:"excludes,block"`
	// Includes: optional
	Includes *datasynctask.Includes `hcl:"includes,block"`
	// Options: optional
	Options *datasynctask.Options `hcl:"options,block"`
	// Schedule: optional
	Schedule *datasynctask.Schedule `hcl:"schedule,block"`
	// Timeouts: optional
	Timeouts *datasynctask.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that DatasyncTask depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type datasyncTaskAttributes struct {
	ref terra.Reference
}

func (dt datasyncTaskAttributes) Arn() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("arn"))
}

func (dt datasyncTaskAttributes) CloudwatchLogGroupArn() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("cloudwatch_log_group_arn"))
}

func (dt datasyncTaskAttributes) DestinationLocationArn() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("destination_location_arn"))
}

func (dt datasyncTaskAttributes) Id() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("id"))
}

func (dt datasyncTaskAttributes) Name() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("name"))
}

func (dt datasyncTaskAttributes) SourceLocationArn() terra.StringValue {
	return terra.ReferenceString(dt.ref.Append("source_location_arn"))
}

func (dt datasyncTaskAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dt.ref.Append("tags"))
}

func (dt datasyncTaskAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](dt.ref.Append("tags_all"))
}

func (dt datasyncTaskAttributes) Excludes() terra.ListValue[datasynctask.ExcludesAttributes] {
	return terra.ReferenceList[datasynctask.ExcludesAttributes](dt.ref.Append("excludes"))
}

func (dt datasyncTaskAttributes) Includes() terra.ListValue[datasynctask.IncludesAttributes] {
	return terra.ReferenceList[datasynctask.IncludesAttributes](dt.ref.Append("includes"))
}

func (dt datasyncTaskAttributes) Options() terra.ListValue[datasynctask.OptionsAttributes] {
	return terra.ReferenceList[datasynctask.OptionsAttributes](dt.ref.Append("options"))
}

func (dt datasyncTaskAttributes) Schedule() terra.ListValue[datasynctask.ScheduleAttributes] {
	return terra.ReferenceList[datasynctask.ScheduleAttributes](dt.ref.Append("schedule"))
}

func (dt datasyncTaskAttributes) Timeouts() datasynctask.TimeoutsAttributes {
	return terra.ReferenceSingle[datasynctask.TimeoutsAttributes](dt.ref.Append("timeouts"))
}

type datasyncTaskState struct {
	Arn                    string                       `json:"arn"`
	CloudwatchLogGroupArn  string                       `json:"cloudwatch_log_group_arn"`
	DestinationLocationArn string                       `json:"destination_location_arn"`
	Id                     string                       `json:"id"`
	Name                   string                       `json:"name"`
	SourceLocationArn      string                       `json:"source_location_arn"`
	Tags                   map[string]string            `json:"tags"`
	TagsAll                map[string]string            `json:"tags_all"`
	Excludes               []datasynctask.ExcludesState `json:"excludes"`
	Includes               []datasynctask.IncludesState `json:"includes"`
	Options                []datasynctask.OptionsState  `json:"options"`
	Schedule               []datasynctask.ScheduleState `json:"schedule"`
	Timeouts               *datasynctask.TimeoutsState  `json:"timeouts"`
}
