// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_cloudwatch_event_api_destination

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

// Resource represents the Terraform resource aws_cloudwatch_event_api_destination.
type Resource struct {
	Name      string
	Args      Args
	state     *awsCloudwatchEventApiDestinationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acead *Resource) Type() string {
	return "aws_cloudwatch_event_api_destination"
}

// LocalName returns the local name for [Resource].
func (acead *Resource) LocalName() string {
	return acead.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acead *Resource) Configuration() interface{} {
	return acead.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acead *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acead)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acead *Resource) Dependencies() terra.Dependencies {
	return acead.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acead *Resource) LifecycleManagement() *terra.Lifecycle {
	return acead.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acead *Resource) Attributes() awsCloudwatchEventApiDestinationAttributes {
	return awsCloudwatchEventApiDestinationAttributes{ref: terra.ReferenceResource(acead)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acead *Resource) ImportState(state io.Reader) error {
	acead.state = &awsCloudwatchEventApiDestinationState{}
	if err := json.NewDecoder(state).Decode(acead.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acead.Type(), acead.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acead *Resource) State() (*awsCloudwatchEventApiDestinationState, bool) {
	return acead.state, acead.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acead *Resource) StateMust() *awsCloudwatchEventApiDestinationState {
	if acead.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acead.Type(), acead.LocalName()))
	}
	return acead.state
}

// Args contains the configurations for aws_cloudwatch_event_api_destination.
type Args struct {
	// ConnectionArn: string, required
	ConnectionArn terra.StringValue `hcl:"connection_arn,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// HttpMethod: string, required
	HttpMethod terra.StringValue `hcl:"http_method,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InvocationEndpoint: string, required
	InvocationEndpoint terra.StringValue `hcl:"invocation_endpoint,attr" validate:"required"`
	// InvocationRateLimitPerSecond: number, optional
	InvocationRateLimitPerSecond terra.NumberValue `hcl:"invocation_rate_limit_per_second,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type awsCloudwatchEventApiDestinationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("arn"))
}

// ConnectionArn returns a reference to field connection_arn of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("connection_arn"))
}

// Description returns a reference to field description of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("description"))
}

// HttpMethod returns a reference to field http_method of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("http_method"))
}

// Id returns a reference to field id of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("id"))
}

// InvocationEndpoint returns a reference to field invocation_endpoint of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) InvocationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("invocation_endpoint"))
}

// InvocationRateLimitPerSecond returns a reference to field invocation_rate_limit_per_second of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) InvocationRateLimitPerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(acead.ref.Append("invocation_rate_limit_per_second"))
}

// Name returns a reference to field name of aws_cloudwatch_event_api_destination.
func (acead awsCloudwatchEventApiDestinationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acead.ref.Append("name"))
}

type awsCloudwatchEventApiDestinationState struct {
	Arn                          string  `json:"arn"`
	ConnectionArn                string  `json:"connection_arn"`
	Description                  string  `json:"description"`
	HttpMethod                   string  `json:"http_method"`
	Id                           string  `json:"id"`
	InvocationEndpoint           string  `json:"invocation_endpoint"`
	InvocationRateLimitPerSecond float64 `json:"invocation_rate_limit_per_second"`
	Name                         string  `json:"name"`
}
