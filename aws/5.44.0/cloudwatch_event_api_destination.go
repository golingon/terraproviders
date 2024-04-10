// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// NewCloudwatchEventApiDestination creates a new instance of [CloudwatchEventApiDestination].
func NewCloudwatchEventApiDestination(name string, args CloudwatchEventApiDestinationArgs) *CloudwatchEventApiDestination {
	return &CloudwatchEventApiDestination{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchEventApiDestination)(nil)

// CloudwatchEventApiDestination represents the Terraform resource aws_cloudwatch_event_api_destination.
type CloudwatchEventApiDestination struct {
	Name      string
	Args      CloudwatchEventApiDestinationArgs
	state     *cloudwatchEventApiDestinationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) Type() string {
	return "aws_cloudwatch_event_api_destination"
}

// LocalName returns the local name for [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) LocalName() string {
	return cead.Name
}

// Configuration returns the configuration (args) for [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) Configuration() interface{} {
	return cead.Args
}

// DependOn is used for other resources to depend on [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) DependOn() terra.Reference {
	return terra.ReferenceResource(cead)
}

// Dependencies returns the list of resources [CloudwatchEventApiDestination] depends_on.
func (cead *CloudwatchEventApiDestination) Dependencies() terra.Dependencies {
	return cead.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) LifecycleManagement() *terra.Lifecycle {
	return cead.Lifecycle
}

// Attributes returns the attributes for [CloudwatchEventApiDestination].
func (cead *CloudwatchEventApiDestination) Attributes() cloudwatchEventApiDestinationAttributes {
	return cloudwatchEventApiDestinationAttributes{ref: terra.ReferenceResource(cead)}
}

// ImportState imports the given attribute values into [CloudwatchEventApiDestination]'s state.
func (cead *CloudwatchEventApiDestination) ImportState(av io.Reader) error {
	cead.state = &cloudwatchEventApiDestinationState{}
	if err := json.NewDecoder(av).Decode(cead.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cead.Type(), cead.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchEventApiDestination] has state.
func (cead *CloudwatchEventApiDestination) State() (*cloudwatchEventApiDestinationState, bool) {
	return cead.state, cead.state != nil
}

// StateMust returns the state for [CloudwatchEventApiDestination]. Panics if the state is nil.
func (cead *CloudwatchEventApiDestination) StateMust() *cloudwatchEventApiDestinationState {
	if cead.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cead.Type(), cead.LocalName()))
	}
	return cead.state
}

// CloudwatchEventApiDestinationArgs contains the configurations for aws_cloudwatch_event_api_destination.
type CloudwatchEventApiDestinationArgs struct {
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
type cloudwatchEventApiDestinationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("arn"))
}

// ConnectionArn returns a reference to field connection_arn of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) ConnectionArn() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("connection_arn"))
}

// Description returns a reference to field description of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("description"))
}

// HttpMethod returns a reference to field http_method of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) HttpMethod() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("http_method"))
}

// Id returns a reference to field id of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("id"))
}

// InvocationEndpoint returns a reference to field invocation_endpoint of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) InvocationEndpoint() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("invocation_endpoint"))
}

// InvocationRateLimitPerSecond returns a reference to field invocation_rate_limit_per_second of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) InvocationRateLimitPerSecond() terra.NumberValue {
	return terra.ReferenceAsNumber(cead.ref.Append("invocation_rate_limit_per_second"))
}

// Name returns a reference to field name of aws_cloudwatch_event_api_destination.
func (cead cloudwatchEventApiDestinationAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cead.ref.Append("name"))
}

type cloudwatchEventApiDestinationState struct {
	Arn                          string  `json:"arn"`
	ConnectionArn                string  `json:"connection_arn"`
	Description                  string  `json:"description"`
	HttpMethod                   string  `json:"http_method"`
	Id                           string  `json:"id"`
	InvocationEndpoint           string  `json:"invocation_endpoint"`
	InvocationRateLimitPerSecond float64 `json:"invocation_rate_limit_per_second"`
	Name                         string  `json:"name"`
}
