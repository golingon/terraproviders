// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	qldbstream "github.com/golingon/terraproviders/aws/5.10.0/qldbstream"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewQldbStream creates a new instance of [QldbStream].
func NewQldbStream(name string, args QldbStreamArgs) *QldbStream {
	return &QldbStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*QldbStream)(nil)

// QldbStream represents the Terraform resource aws_qldb_stream.
type QldbStream struct {
	Name      string
	Args      QldbStreamArgs
	state     *qldbStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [QldbStream].
func (qs *QldbStream) Type() string {
	return "aws_qldb_stream"
}

// LocalName returns the local name for [QldbStream].
func (qs *QldbStream) LocalName() string {
	return qs.Name
}

// Configuration returns the configuration (args) for [QldbStream].
func (qs *QldbStream) Configuration() interface{} {
	return qs.Args
}

// DependOn is used for other resources to depend on [QldbStream].
func (qs *QldbStream) DependOn() terra.Reference {
	return terra.ReferenceResource(qs)
}

// Dependencies returns the list of resources [QldbStream] depends_on.
func (qs *QldbStream) Dependencies() terra.Dependencies {
	return qs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [QldbStream].
func (qs *QldbStream) LifecycleManagement() *terra.Lifecycle {
	return qs.Lifecycle
}

// Attributes returns the attributes for [QldbStream].
func (qs *QldbStream) Attributes() qldbStreamAttributes {
	return qldbStreamAttributes{ref: terra.ReferenceResource(qs)}
}

// ImportState imports the given attribute values into [QldbStream]'s state.
func (qs *QldbStream) ImportState(av io.Reader) error {
	qs.state = &qldbStreamState{}
	if err := json.NewDecoder(av).Decode(qs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", qs.Type(), qs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [QldbStream] has state.
func (qs *QldbStream) State() (*qldbStreamState, bool) {
	return qs.state, qs.state != nil
}

// StateMust returns the state for [QldbStream]. Panics if the state is nil.
func (qs *QldbStream) StateMust() *qldbStreamState {
	if qs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", qs.Type(), qs.LocalName()))
	}
	return qs.state
}

// QldbStreamArgs contains the configurations for aws_qldb_stream.
type QldbStreamArgs struct {
	// ExclusiveEndTime: string, optional
	ExclusiveEndTime terra.StringValue `hcl:"exclusive_end_time,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InclusiveStartTime: string, required
	InclusiveStartTime terra.StringValue `hcl:"inclusive_start_time,attr" validate:"required"`
	// LedgerName: string, required
	LedgerName terra.StringValue `hcl:"ledger_name,attr" validate:"required"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// StreamName: string, required
	StreamName terra.StringValue `hcl:"stream_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// KinesisConfiguration: required
	KinesisConfiguration *qldbstream.KinesisConfiguration `hcl:"kinesis_configuration,block" validate:"required"`
	// Timeouts: optional
	Timeouts *qldbstream.Timeouts `hcl:"timeouts,block"`
}
type qldbStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_qldb_stream.
func (qs qldbStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("arn"))
}

// ExclusiveEndTime returns a reference to field exclusive_end_time of aws_qldb_stream.
func (qs qldbStreamAttributes) ExclusiveEndTime() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("exclusive_end_time"))
}

// Id returns a reference to field id of aws_qldb_stream.
func (qs qldbStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("id"))
}

// InclusiveStartTime returns a reference to field inclusive_start_time of aws_qldb_stream.
func (qs qldbStreamAttributes) InclusiveStartTime() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("inclusive_start_time"))
}

// LedgerName returns a reference to field ledger_name of aws_qldb_stream.
func (qs qldbStreamAttributes) LedgerName() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("ledger_name"))
}

// RoleArn returns a reference to field role_arn of aws_qldb_stream.
func (qs qldbStreamAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("role_arn"))
}

// StreamName returns a reference to field stream_name of aws_qldb_stream.
func (qs qldbStreamAttributes) StreamName() terra.StringValue {
	return terra.ReferenceAsString(qs.ref.Append("stream_name"))
}

// Tags returns a reference to field tags of aws_qldb_stream.
func (qs qldbStreamAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qs.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_qldb_stream.
func (qs qldbStreamAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](qs.ref.Append("tags_all"))
}

func (qs qldbStreamAttributes) KinesisConfiguration() terra.ListValue[qldbstream.KinesisConfigurationAttributes] {
	return terra.ReferenceAsList[qldbstream.KinesisConfigurationAttributes](qs.ref.Append("kinesis_configuration"))
}

func (qs qldbStreamAttributes) Timeouts() qldbstream.TimeoutsAttributes {
	return terra.ReferenceAsSingle[qldbstream.TimeoutsAttributes](qs.ref.Append("timeouts"))
}

type qldbStreamState struct {
	Arn                  string                                 `json:"arn"`
	ExclusiveEndTime     string                                 `json:"exclusive_end_time"`
	Id                   string                                 `json:"id"`
	InclusiveStartTime   string                                 `json:"inclusive_start_time"`
	LedgerName           string                                 `json:"ledger_name"`
	RoleArn              string                                 `json:"role_arn"`
	StreamName           string                                 `json:"stream_name"`
	Tags                 map[string]string                      `json:"tags"`
	TagsAll              map[string]string                      `json:"tags_all"`
	KinesisConfiguration []qldbstream.KinesisConfigurationState `json:"kinesis_configuration"`
	Timeouts             *qldbstream.TimeoutsState              `json:"timeouts"`
}
