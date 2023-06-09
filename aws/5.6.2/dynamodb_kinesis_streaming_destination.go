// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewDynamodbKinesisStreamingDestination creates a new instance of [DynamodbKinesisStreamingDestination].
func NewDynamodbKinesisStreamingDestination(name string, args DynamodbKinesisStreamingDestinationArgs) *DynamodbKinesisStreamingDestination {
	return &DynamodbKinesisStreamingDestination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*DynamodbKinesisStreamingDestination)(nil)

// DynamodbKinesisStreamingDestination represents the Terraform resource aws_dynamodb_kinesis_streaming_destination.
type DynamodbKinesisStreamingDestination struct {
	Name      string
	Args      DynamodbKinesisStreamingDestinationArgs
	state     *dynamodbKinesisStreamingDestinationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) Type() string {
	return "aws_dynamodb_kinesis_streaming_destination"
}

// LocalName returns the local name for [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) LocalName() string {
	return dksd.Name
}

// Configuration returns the configuration (args) for [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) Configuration() interface{} {
	return dksd.Args
}

// DependOn is used for other resources to depend on [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) DependOn() terra.Reference {
	return terra.ReferenceResource(dksd)
}

// Dependencies returns the list of resources [DynamodbKinesisStreamingDestination] depends_on.
func (dksd *DynamodbKinesisStreamingDestination) Dependencies() terra.Dependencies {
	return dksd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) LifecycleManagement() *terra.Lifecycle {
	return dksd.Lifecycle
}

// Attributes returns the attributes for [DynamodbKinesisStreamingDestination].
func (dksd *DynamodbKinesisStreamingDestination) Attributes() dynamodbKinesisStreamingDestinationAttributes {
	return dynamodbKinesisStreamingDestinationAttributes{ref: terra.ReferenceResource(dksd)}
}

// ImportState imports the given attribute values into [DynamodbKinesisStreamingDestination]'s state.
func (dksd *DynamodbKinesisStreamingDestination) ImportState(av io.Reader) error {
	dksd.state = &dynamodbKinesisStreamingDestinationState{}
	if err := json.NewDecoder(av).Decode(dksd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", dksd.Type(), dksd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [DynamodbKinesisStreamingDestination] has state.
func (dksd *DynamodbKinesisStreamingDestination) State() (*dynamodbKinesisStreamingDestinationState, bool) {
	return dksd.state, dksd.state != nil
}

// StateMust returns the state for [DynamodbKinesisStreamingDestination]. Panics if the state is nil.
func (dksd *DynamodbKinesisStreamingDestination) StateMust() *dynamodbKinesisStreamingDestinationState {
	if dksd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", dksd.Type(), dksd.LocalName()))
	}
	return dksd.state
}

// DynamodbKinesisStreamingDestinationArgs contains the configurations for aws_dynamodb_kinesis_streaming_destination.
type DynamodbKinesisStreamingDestinationArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// StreamArn: string, required
	StreamArn terra.StringValue `hcl:"stream_arn,attr" validate:"required"`
	// TableName: string, required
	TableName terra.StringValue `hcl:"table_name,attr" validate:"required"`
}
type dynamodbKinesisStreamingDestinationAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_dynamodb_kinesis_streaming_destination.
func (dksd dynamodbKinesisStreamingDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(dksd.ref.Append("id"))
}

// StreamArn returns a reference to field stream_arn of aws_dynamodb_kinesis_streaming_destination.
func (dksd dynamodbKinesisStreamingDestinationAttributes) StreamArn() terra.StringValue {
	return terra.ReferenceAsString(dksd.ref.Append("stream_arn"))
}

// TableName returns a reference to field table_name of aws_dynamodb_kinesis_streaming_destination.
func (dksd dynamodbKinesisStreamingDestinationAttributes) TableName() terra.StringValue {
	return terra.ReferenceAsString(dksd.ref.Append("table_name"))
}

type dynamodbKinesisStreamingDestinationState struct {
	Id        string `json:"id"`
	StreamArn string `json:"stream_arn"`
	TableName string `json:"table_name"`
}
