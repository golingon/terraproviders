// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	schedulerschedulegroup "github.com/golingon/terraproviders/aws/4.66.1/schedulerschedulegroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSchedulerScheduleGroup creates a new instance of [SchedulerScheduleGroup].
func NewSchedulerScheduleGroup(name string, args SchedulerScheduleGroupArgs) *SchedulerScheduleGroup {
	return &SchedulerScheduleGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SchedulerScheduleGroup)(nil)

// SchedulerScheduleGroup represents the Terraform resource aws_scheduler_schedule_group.
type SchedulerScheduleGroup struct {
	Name      string
	Args      SchedulerScheduleGroupArgs
	state     *schedulerScheduleGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) Type() string {
	return "aws_scheduler_schedule_group"
}

// LocalName returns the local name for [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) LocalName() string {
	return ssg.Name
}

// Configuration returns the configuration (args) for [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) Configuration() interface{} {
	return ssg.Args
}

// DependOn is used for other resources to depend on [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(ssg)
}

// Dependencies returns the list of resources [SchedulerScheduleGroup] depends_on.
func (ssg *SchedulerScheduleGroup) Dependencies() terra.Dependencies {
	return ssg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) LifecycleManagement() *terra.Lifecycle {
	return ssg.Lifecycle
}

// Attributes returns the attributes for [SchedulerScheduleGroup].
func (ssg *SchedulerScheduleGroup) Attributes() schedulerScheduleGroupAttributes {
	return schedulerScheduleGroupAttributes{ref: terra.ReferenceResource(ssg)}
}

// ImportState imports the given attribute values into [SchedulerScheduleGroup]'s state.
func (ssg *SchedulerScheduleGroup) ImportState(av io.Reader) error {
	ssg.state = &schedulerScheduleGroupState{}
	if err := json.NewDecoder(av).Decode(ssg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ssg.Type(), ssg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SchedulerScheduleGroup] has state.
func (ssg *SchedulerScheduleGroup) State() (*schedulerScheduleGroupState, bool) {
	return ssg.state, ssg.state != nil
}

// StateMust returns the state for [SchedulerScheduleGroup]. Panics if the state is nil.
func (ssg *SchedulerScheduleGroup) StateMust() *schedulerScheduleGroupState {
	if ssg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ssg.Type(), ssg.LocalName()))
	}
	return ssg.state
}

// SchedulerScheduleGroupArgs contains the configurations for aws_scheduler_schedule_group.
type SchedulerScheduleGroupArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// NamePrefix: string, optional
	NamePrefix terra.StringValue `hcl:"name_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// Timeouts: optional
	Timeouts *schedulerschedulegroup.Timeouts `hcl:"timeouts,block"`
}
type schedulerScheduleGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("arn"))
}

// CreationDate returns a reference to field creation_date of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) CreationDate() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("creation_date"))
}

// Id returns a reference to field id of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("id"))
}

// LastModificationDate returns a reference to field last_modification_date of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) LastModificationDate() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("last_modification_date"))
}

// Name returns a reference to field name of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("name"))
}

// NamePrefix returns a reference to field name_prefix of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) NamePrefix() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("name_prefix"))
}

// State returns a reference to field state of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ssg.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssg.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_scheduler_schedule_group.
func (ssg schedulerScheduleGroupAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ssg.ref.Append("tags_all"))
}

func (ssg schedulerScheduleGroupAttributes) Timeouts() schedulerschedulegroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[schedulerschedulegroup.TimeoutsAttributes](ssg.ref.Append("timeouts"))
}

type schedulerScheduleGroupState struct {
	Arn                  string                                `json:"arn"`
	CreationDate         string                                `json:"creation_date"`
	Id                   string                                `json:"id"`
	LastModificationDate string                                `json:"last_modification_date"`
	Name                 string                                `json:"name"`
	NamePrefix           string                                `json:"name_prefix"`
	State                string                                `json:"state"`
	Tags                 map[string]string                     `json:"tags"`
	TagsAll              map[string]string                     `json:"tags_all"`
	Timeouts             *schedulerschedulegroup.TimeoutsState `json:"timeouts"`
}
