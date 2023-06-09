// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewRedshiftSnapshotSchedule creates a new instance of [RedshiftSnapshotSchedule].
func NewRedshiftSnapshotSchedule(name string, args RedshiftSnapshotScheduleArgs) *RedshiftSnapshotSchedule {
	return &RedshiftSnapshotSchedule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*RedshiftSnapshotSchedule)(nil)

// RedshiftSnapshotSchedule represents the Terraform resource aws_redshift_snapshot_schedule.
type RedshiftSnapshotSchedule struct {
	Name      string
	Args      RedshiftSnapshotScheduleArgs
	state     *redshiftSnapshotScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) Type() string {
	return "aws_redshift_snapshot_schedule"
}

// LocalName returns the local name for [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) LocalName() string {
	return rss.Name
}

// Configuration returns the configuration (args) for [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) Configuration() interface{} {
	return rss.Args
}

// DependOn is used for other resources to depend on [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) DependOn() terra.Reference {
	return terra.ReferenceResource(rss)
}

// Dependencies returns the list of resources [RedshiftSnapshotSchedule] depends_on.
func (rss *RedshiftSnapshotSchedule) Dependencies() terra.Dependencies {
	return rss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) LifecycleManagement() *terra.Lifecycle {
	return rss.Lifecycle
}

// Attributes returns the attributes for [RedshiftSnapshotSchedule].
func (rss *RedshiftSnapshotSchedule) Attributes() redshiftSnapshotScheduleAttributes {
	return redshiftSnapshotScheduleAttributes{ref: terra.ReferenceResource(rss)}
}

// ImportState imports the given attribute values into [RedshiftSnapshotSchedule]'s state.
func (rss *RedshiftSnapshotSchedule) ImportState(av io.Reader) error {
	rss.state = &redshiftSnapshotScheduleState{}
	if err := json.NewDecoder(av).Decode(rss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rss.Type(), rss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [RedshiftSnapshotSchedule] has state.
func (rss *RedshiftSnapshotSchedule) State() (*redshiftSnapshotScheduleState, bool) {
	return rss.state, rss.state != nil
}

// StateMust returns the state for [RedshiftSnapshotSchedule]. Panics if the state is nil.
func (rss *RedshiftSnapshotSchedule) StateMust() *redshiftSnapshotScheduleState {
	if rss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rss.Type(), rss.LocalName()))
	}
	return rss.state
}

// RedshiftSnapshotScheduleArgs contains the configurations for aws_redshift_snapshot_schedule.
type RedshiftSnapshotScheduleArgs struct {
	// Definitions: set of string, required
	Definitions terra.SetValue[terra.StringValue] `hcl:"definitions,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ForceDestroy: bool, optional
	ForceDestroy terra.BoolValue `hcl:"force_destroy,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Identifier: string, optional
	Identifier terra.StringValue `hcl:"identifier,attr"`
	// IdentifierPrefix: string, optional
	IdentifierPrefix terra.StringValue `hcl:"identifier_prefix,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type redshiftSnapshotScheduleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(rss.ref.Append("arn"))
}

// Definitions returns a reference to field definitions of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Definitions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](rss.ref.Append("definitions"))
}

// Description returns a reference to field description of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(rss.ref.Append("description"))
}

// ForceDestroy returns a reference to field force_destroy of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(rss.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rss.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(rss.ref.Append("identifier"))
}

// IdentifierPrefix returns a reference to field identifier_prefix of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) IdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(rss.ref.Append("identifier_prefix"))
}

// Tags returns a reference to field tags of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rss.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_snapshot_schedule.
func (rss redshiftSnapshotScheduleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rss.ref.Append("tags_all"))
}

type redshiftSnapshotScheduleState struct {
	Arn              string            `json:"arn"`
	Definitions      []string          `json:"definitions"`
	Description      string            `json:"description"`
	ForceDestroy     bool              `json:"force_destroy"`
	Id               string            `json:"id"`
	Identifier       string            `json:"identifier"`
	IdentifierPrefix string            `json:"identifier_prefix"`
	Tags             map[string]string `json:"tags"`
	TagsAll          map[string]string `json:"tags_all"`
}
