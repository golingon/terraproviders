// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudwatcheventendpoint "github.com/golingon/terraproviders/aws/5.11.0/cloudwatcheventendpoint"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudwatchEventEndpoint creates a new instance of [CloudwatchEventEndpoint].
func NewCloudwatchEventEndpoint(name string, args CloudwatchEventEndpointArgs) *CloudwatchEventEndpoint {
	return &CloudwatchEventEndpoint{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudwatchEventEndpoint)(nil)

// CloudwatchEventEndpoint represents the Terraform resource aws_cloudwatch_event_endpoint.
type CloudwatchEventEndpoint struct {
	Name      string
	Args      CloudwatchEventEndpointArgs
	state     *cloudwatchEventEndpointState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) Type() string {
	return "aws_cloudwatch_event_endpoint"
}

// LocalName returns the local name for [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) LocalName() string {
	return cee.Name
}

// Configuration returns the configuration (args) for [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) Configuration() interface{} {
	return cee.Args
}

// DependOn is used for other resources to depend on [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) DependOn() terra.Reference {
	return terra.ReferenceResource(cee)
}

// Dependencies returns the list of resources [CloudwatchEventEndpoint] depends_on.
func (cee *CloudwatchEventEndpoint) Dependencies() terra.Dependencies {
	return cee.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) LifecycleManagement() *terra.Lifecycle {
	return cee.Lifecycle
}

// Attributes returns the attributes for [CloudwatchEventEndpoint].
func (cee *CloudwatchEventEndpoint) Attributes() cloudwatchEventEndpointAttributes {
	return cloudwatchEventEndpointAttributes{ref: terra.ReferenceResource(cee)}
}

// ImportState imports the given attribute values into [CloudwatchEventEndpoint]'s state.
func (cee *CloudwatchEventEndpoint) ImportState(av io.Reader) error {
	cee.state = &cloudwatchEventEndpointState{}
	if err := json.NewDecoder(av).Decode(cee.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cee.Type(), cee.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudwatchEventEndpoint] has state.
func (cee *CloudwatchEventEndpoint) State() (*cloudwatchEventEndpointState, bool) {
	return cee.state, cee.state != nil
}

// StateMust returns the state for [CloudwatchEventEndpoint]. Panics if the state is nil.
func (cee *CloudwatchEventEndpoint) StateMust() *cloudwatchEventEndpointState {
	if cee.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cee.Type(), cee.LocalName()))
	}
	return cee.state
}

// CloudwatchEventEndpointArgs contains the configurations for aws_cloudwatch_event_endpoint.
type CloudwatchEventEndpointArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RoleArn: string, optional
	RoleArn terra.StringValue `hcl:"role_arn,attr"`
	// EventBus: min=2,max=2
	EventBus []cloudwatcheventendpoint.EventBus `hcl:"event_bus,block" validate:"min=2,max=2"`
	// ReplicationConfig: optional
	ReplicationConfig *cloudwatcheventendpoint.ReplicationConfig `hcl:"replication_config,block"`
	// RoutingConfig: required
	RoutingConfig *cloudwatcheventendpoint.RoutingConfig `hcl:"routing_config,block" validate:"required"`
}
type cloudwatchEventEndpointAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("arn"))
}

// Description returns a reference to field description of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("description"))
}

// EndpointUrl returns a reference to field endpoint_url of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) EndpointUrl() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("endpoint_url"))
}

// Id returns a reference to field id of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_cloudwatch_event_endpoint.
func (cee cloudwatchEventEndpointAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(cee.ref.Append("role_arn"))
}

func (cee cloudwatchEventEndpointAttributes) EventBus() terra.ListValue[cloudwatcheventendpoint.EventBusAttributes] {
	return terra.ReferenceAsList[cloudwatcheventendpoint.EventBusAttributes](cee.ref.Append("event_bus"))
}

func (cee cloudwatchEventEndpointAttributes) ReplicationConfig() terra.ListValue[cloudwatcheventendpoint.ReplicationConfigAttributes] {
	return terra.ReferenceAsList[cloudwatcheventendpoint.ReplicationConfigAttributes](cee.ref.Append("replication_config"))
}

func (cee cloudwatchEventEndpointAttributes) RoutingConfig() terra.ListValue[cloudwatcheventendpoint.RoutingConfigAttributes] {
	return terra.ReferenceAsList[cloudwatcheventendpoint.RoutingConfigAttributes](cee.ref.Append("routing_config"))
}

type cloudwatchEventEndpointState struct {
	Arn               string                                           `json:"arn"`
	Description       string                                           `json:"description"`
	EndpointUrl       string                                           `json:"endpoint_url"`
	Id                string                                           `json:"id"`
	Name              string                                           `json:"name"`
	RoleArn           string                                           `json:"role_arn"`
	EventBus          []cloudwatcheventendpoint.EventBusState          `json:"event_bus"`
	ReplicationConfig []cloudwatcheventendpoint.ReplicationConfigState `json:"replication_config"`
	RoutingConfig     []cloudwatcheventendpoint.RoutingConfigState     `json:"routing_config"`
}
