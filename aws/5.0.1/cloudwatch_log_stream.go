// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchLogStream creates a new instance of [CloudwatchLogStream].
func NewCloudwatchLogStream(name string, args CloudwatchLogStreamArgs) *CloudwatchLogStream {
	return &CloudwatchLogStream{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchLogStream)(nil)

// CloudwatchLogStream represents the Terraform resource aws_cloudwatch_log_stream.
type CloudwatchLogStream struct {
	Name      string
	Args      CloudwatchLogStreamArgs
	state     *cloudwatchLogStreamState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchLogStream].
func (cls *CloudwatchLogStream) Type() string {
	return "aws_cloudwatch_log_stream"
}

// LocalName returns the local name for [CloudwatchLogStream].
func (cls *CloudwatchLogStream) LocalName() string {
	return cls.Name
}

// Configuration returns the configuration (args) for [CloudwatchLogStream].
func (cls *CloudwatchLogStream) Configuration() interface{} {
	return cls.Args
}

// DependOn is used for other resources to depend on [CloudwatchLogStream].
func (cls *CloudwatchLogStream) DependOn() terra.Reference {
	return terra.ReferenceResource(cls)
}

// Dependencies returns the list of resources [CloudwatchLogStream] depends_on.
func (cls *CloudwatchLogStream) Dependencies() terra.Dependencies {
	return cls.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchLogStream].
func (cls *CloudwatchLogStream) LifecycleManagement() *terra.Lifecycle {
	return cls.Lifecycle
}

// Attributes returns the attributes for [CloudwatchLogStream].
func (cls *CloudwatchLogStream) Attributes() cloudwatchLogStreamAttributes {
	return cloudwatchLogStreamAttributes{ref: terra.ReferenceResource(cls)}
}

// ImportState imports the given attribute values into [CloudwatchLogStream]'s state.
func (cls *CloudwatchLogStream) ImportState(av io.Reader) error {
	cls.state = &cloudwatchLogStreamState{}
	if err := json.NewDecoder(av).Decode(cls.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cls.Type(), cls.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchLogStream] has state.
func (cls *CloudwatchLogStream) State() (*cloudwatchLogStreamState, bool) {
	return cls.state, cls.state != nil
}

// StateMust returns the state for [CloudwatchLogStream]. Panics if the state is nil.
func (cls *CloudwatchLogStream) StateMust() *cloudwatchLogStreamState {
	if cls.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cls.Type(), cls.LocalName()))
	}
	return cls.state
}

// CloudwatchLogStreamArgs contains the configurations for aws_cloudwatch_log_stream.
type CloudwatchLogStreamArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LogGroupName: string, required
	LogGroupName terra.StringValue `hcl:"log_group_name,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type cloudwatchLogStreamAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_log_stream.
func (cls cloudwatchLogStreamAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cls.ref.Append("arn"))
}

// Id returns a reference to field id of aws_cloudwatch_log_stream.
func (cls cloudwatchLogStreamAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cls.ref.Append("id"))
}

// LogGroupName returns a reference to field log_group_name of aws_cloudwatch_log_stream.
func (cls cloudwatchLogStreamAttributes) LogGroupName() terra.StringValue {
	return terra.ReferenceAsString(cls.ref.Append("log_group_name"))
}

// Name returns a reference to field name of aws_cloudwatch_log_stream.
func (cls cloudwatchLogStreamAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cls.ref.Append("name"))
}

type cloudwatchLogStreamState struct {
	Arn          string `json:"arn"`
	Id           string `json:"id"`
	LogGroupName string `json:"log_group_name"`
	Name         string `json:"name"`
}
