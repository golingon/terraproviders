// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_redshift_snapshot_schedule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_redshift_snapshot_schedule.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRedshiftSnapshotScheduleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arss *Resource) Type() string {
	return "aws_redshift_snapshot_schedule"
}

// LocalName returns the local name for [Resource].
func (arss *Resource) LocalName() string {
	return arss.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arss *Resource) Configuration() interface{} {
	return arss.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arss *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arss)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arss *Resource) Dependencies() terra.Dependencies {
	return arss.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arss *Resource) LifecycleManagement() *terra.Lifecycle {
	return arss.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arss *Resource) Attributes() awsRedshiftSnapshotScheduleAttributes {
	return awsRedshiftSnapshotScheduleAttributes{ref: terra.ReferenceResource(arss)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arss *Resource) ImportState(state io.Reader) error {
	arss.state = &awsRedshiftSnapshotScheduleState{}
	if err := json.NewDecoder(state).Decode(arss.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arss.Type(), arss.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arss *Resource) State() (*awsRedshiftSnapshotScheduleState, bool) {
	return arss.state, arss.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arss *Resource) StateMust() *awsRedshiftSnapshotScheduleState {
	if arss.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arss.Type(), arss.LocalName()))
	}
	return arss.state
}

// Args contains the configurations for aws_redshift_snapshot_schedule.
type Args struct {
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

type awsRedshiftSnapshotScheduleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arss.ref.Append("arn"))
}

// Definitions returns a reference to field definitions of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Definitions() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](arss.ref.Append("definitions"))
}

// Description returns a reference to field description of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(arss.ref.Append("description"))
}

// ForceDestroy returns a reference to field force_destroy of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) ForceDestroy() terra.BoolValue {
	return terra.ReferenceAsBool(arss.ref.Append("force_destroy"))
}

// Id returns a reference to field id of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arss.ref.Append("id"))
}

// Identifier returns a reference to field identifier of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Identifier() terra.StringValue {
	return terra.ReferenceAsString(arss.ref.Append("identifier"))
}

// IdentifierPrefix returns a reference to field identifier_prefix of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) IdentifierPrefix() terra.StringValue {
	return terra.ReferenceAsString(arss.ref.Append("identifier_prefix"))
}

// Tags returns a reference to field tags of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arss.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_redshift_snapshot_schedule.
func (arss awsRedshiftSnapshotScheduleAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](arss.ref.Append("tags_all"))
}

type awsRedshiftSnapshotScheduleState struct {
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
