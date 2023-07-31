// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	evidentlylaunch "github.com/golingon/terraproviders/aws/5.10.0/evidentlylaunch"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewEvidentlyLaunch creates a new instance of [EvidentlyLaunch].
func NewEvidentlyLaunch(name string, args EvidentlyLaunchArgs) *EvidentlyLaunch {
	return &EvidentlyLaunch{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*EvidentlyLaunch)(nil)

// EvidentlyLaunch represents the Terraform resource aws_evidently_launch.
type EvidentlyLaunch struct {
	Name      string
	Args      EvidentlyLaunchArgs
	state     *evidentlyLaunchState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [EvidentlyLaunch].
func (el *EvidentlyLaunch) Type() string {
	return "aws_evidently_launch"
}

// LocalName returns the local name for [EvidentlyLaunch].
func (el *EvidentlyLaunch) LocalName() string {
	return el.Name
}

// Configuration returns the configuration (args) for [EvidentlyLaunch].
func (el *EvidentlyLaunch) Configuration() interface{} {
	return el.Args
}

// DependOn is used for other resources to depend on [EvidentlyLaunch].
func (el *EvidentlyLaunch) DependOn() terra.Reference {
	return terra.ReferenceResource(el)
}

// Dependencies returns the list of resources [EvidentlyLaunch] depends_on.
func (el *EvidentlyLaunch) Dependencies() terra.Dependencies {
	return el.DependsOn
}

// LifecycleManagement returns the lifecycle block for [EvidentlyLaunch].
func (el *EvidentlyLaunch) LifecycleManagement() *terra.Lifecycle {
	return el.Lifecycle
}

// Attributes returns the attributes for [EvidentlyLaunch].
func (el *EvidentlyLaunch) Attributes() evidentlyLaunchAttributes {
	return evidentlyLaunchAttributes{ref: terra.ReferenceResource(el)}
}

// ImportState imports the given attribute values into [EvidentlyLaunch]'s state.
func (el *EvidentlyLaunch) ImportState(av io.Reader) error {
	el.state = &evidentlyLaunchState{}
	if err := json.NewDecoder(av).Decode(el.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", el.Type(), el.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [EvidentlyLaunch] has state.
func (el *EvidentlyLaunch) State() (*evidentlyLaunchState, bool) {
	return el.state, el.state != nil
}

// StateMust returns the state for [EvidentlyLaunch]. Panics if the state is nil.
func (el *EvidentlyLaunch) StateMust() *evidentlyLaunchState {
	if el.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", el.Type(), el.LocalName()))
	}
	return el.state
}

// EvidentlyLaunchArgs contains the configurations for aws_evidently_launch.
type EvidentlyLaunchArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Project: string, required
	Project terra.StringValue `hcl:"project,attr" validate:"required"`
	// RandomizationSalt: string, optional
	RandomizationSalt terra.StringValue `hcl:"randomization_salt,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Execution: min=0
	Execution []evidentlylaunch.Execution `hcl:"execution,block" validate:"min=0"`
	// Groups: min=1,max=5
	Groups []evidentlylaunch.Groups `hcl:"groups,block" validate:"min=1,max=5"`
	// MetricMonitors: min=0,max=3
	MetricMonitors []evidentlylaunch.MetricMonitors `hcl:"metric_monitors,block" validate:"min=0,max=3"`
	// ScheduledSplitsConfig: optional
	ScheduledSplitsConfig *evidentlylaunch.ScheduledSplitsConfig `hcl:"scheduled_splits_config,block"`
	// Timeouts: optional
	Timeouts *evidentlylaunch.Timeouts `hcl:"timeouts,block"`
}
type evidentlyLaunchAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("arn"))
}

// CreatedTime returns a reference to field created_time of aws_evidently_launch.
func (el evidentlyLaunchAttributes) CreatedTime() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("created_time"))
}

// Description returns a reference to field description of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("description"))
}

// Id returns a reference to field id of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("id"))
}

// LastUpdatedTime returns a reference to field last_updated_time of aws_evidently_launch.
func (el evidentlyLaunchAttributes) LastUpdatedTime() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("last_updated_time"))
}

// Name returns a reference to field name of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("name"))
}

// Project returns a reference to field project of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Project() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("project"))
}

// RandomizationSalt returns a reference to field randomization_salt of aws_evidently_launch.
func (el evidentlyLaunchAttributes) RandomizationSalt() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("randomization_salt"))
}

// Status returns a reference to field status of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("status"))
}

// StatusReason returns a reference to field status_reason of aws_evidently_launch.
func (el evidentlyLaunchAttributes) StatusReason() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("status_reason"))
}

// Tags returns a reference to field tags of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](el.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_evidently_launch.
func (el evidentlyLaunchAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](el.ref.Append("tags_all"))
}

// Type returns a reference to field type of aws_evidently_launch.
func (el evidentlyLaunchAttributes) Type() terra.StringValue {
	return terra.ReferenceAsString(el.ref.Append("type"))
}

func (el evidentlyLaunchAttributes) Execution() terra.ListValue[evidentlylaunch.ExecutionAttributes] {
	return terra.ReferenceAsList[evidentlylaunch.ExecutionAttributes](el.ref.Append("execution"))
}

func (el evidentlyLaunchAttributes) Groups() terra.ListValue[evidentlylaunch.GroupsAttributes] {
	return terra.ReferenceAsList[evidentlylaunch.GroupsAttributes](el.ref.Append("groups"))
}

func (el evidentlyLaunchAttributes) MetricMonitors() terra.ListValue[evidentlylaunch.MetricMonitorsAttributes] {
	return terra.ReferenceAsList[evidentlylaunch.MetricMonitorsAttributes](el.ref.Append("metric_monitors"))
}

func (el evidentlyLaunchAttributes) ScheduledSplitsConfig() terra.ListValue[evidentlylaunch.ScheduledSplitsConfigAttributes] {
	return terra.ReferenceAsList[evidentlylaunch.ScheduledSplitsConfigAttributes](el.ref.Append("scheduled_splits_config"))
}

func (el evidentlyLaunchAttributes) Timeouts() evidentlylaunch.TimeoutsAttributes {
	return terra.ReferenceAsSingle[evidentlylaunch.TimeoutsAttributes](el.ref.Append("timeouts"))
}

type evidentlyLaunchState struct {
	Arn                   string                                       `json:"arn"`
	CreatedTime           string                                       `json:"created_time"`
	Description           string                                       `json:"description"`
	Id                    string                                       `json:"id"`
	LastUpdatedTime       string                                       `json:"last_updated_time"`
	Name                  string                                       `json:"name"`
	Project               string                                       `json:"project"`
	RandomizationSalt     string                                       `json:"randomization_salt"`
	Status                string                                       `json:"status"`
	StatusReason          string                                       `json:"status_reason"`
	Tags                  map[string]string                            `json:"tags"`
	TagsAll               map[string]string                            `json:"tags_all"`
	Type                  string                                       `json:"type"`
	Execution             []evidentlylaunch.ExecutionState             `json:"execution"`
	Groups                []evidentlylaunch.GroupsState                `json:"groups"`
	MetricMonitors        []evidentlylaunch.MetricMonitorsState        `json:"metric_monitors"`
	ScheduledSplitsConfig []evidentlylaunch.ScheduledSplitsConfigState `json:"scheduled_splits_config"`
	Timeouts              *evidentlylaunch.TimeoutsState               `json:"timeouts"`
}
