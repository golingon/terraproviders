// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_vpc_endpoint_connection_notification

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

// Resource represents the Terraform resource aws_vpc_endpoint_connection_notification.
type Resource struct {
	Name      string
	Args      Args
	state     *awsVpcEndpointConnectionNotificationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (avecn *Resource) Type() string {
	return "aws_vpc_endpoint_connection_notification"
}

// LocalName returns the local name for [Resource].
func (avecn *Resource) LocalName() string {
	return avecn.Name
}

// Configuration returns the configuration (args) for [Resource].
func (avecn *Resource) Configuration() interface{} {
	return avecn.Args
}

// DependOn is used for other resources to depend on [Resource].
func (avecn *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(avecn)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (avecn *Resource) Dependencies() terra.Dependencies {
	return avecn.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (avecn *Resource) LifecycleManagement() *terra.Lifecycle {
	return avecn.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (avecn *Resource) Attributes() awsVpcEndpointConnectionNotificationAttributes {
	return awsVpcEndpointConnectionNotificationAttributes{ref: terra.ReferenceResource(avecn)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (avecn *Resource) ImportState(state io.Reader) error {
	avecn.state = &awsVpcEndpointConnectionNotificationState{}
	if err := json.NewDecoder(state).Decode(avecn.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", avecn.Type(), avecn.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (avecn *Resource) State() (*awsVpcEndpointConnectionNotificationState, bool) {
	return avecn.state, avecn.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (avecn *Resource) StateMust() *awsVpcEndpointConnectionNotificationState {
	if avecn.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", avecn.Type(), avecn.LocalName()))
	}
	return avecn.state
}

// Args contains the configurations for aws_vpc_endpoint_connection_notification.
type Args struct {
	// ConnectionEvents: set of string, required
	ConnectionEvents terra.SetValue[terra.StringValue] `hcl:"connection_events,attr" validate:"required"`
	// ConnectionNotificationArn: string, required
	ConnectionNotificationArn terra.StringValue `hcl:"connection_notification_arn,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// VpcEndpointId: string, optional
	VpcEndpointId terra.StringValue `hcl:"vpc_endpoint_id,attr"`
	// VpcEndpointServiceId: string, optional
	VpcEndpointServiceId terra.StringValue `hcl:"vpc_endpoint_service_id,attr"`
}

type awsVpcEndpointConnectionNotificationAttributes struct {
	ref terra.Reference
}

// ConnectionEvents returns a reference to field connection_events of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) ConnectionEvents() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](avecn.ref.Append("connection_events"))
}

// ConnectionNotificationArn returns a reference to field connection_notification_arn of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) ConnectionNotificationArn() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("connection_notification_arn"))
}

// Id returns a reference to field id of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("id"))
}

// NotificationType returns a reference to field notification_type of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) NotificationType() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("notification_type"))
}

// State returns a reference to field state of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("state"))
}

// VpcEndpointId returns a reference to field vpc_endpoint_id of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) VpcEndpointId() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("vpc_endpoint_id"))
}

// VpcEndpointServiceId returns a reference to field vpc_endpoint_service_id of aws_vpc_endpoint_connection_notification.
func (avecn awsVpcEndpointConnectionNotificationAttributes) VpcEndpointServiceId() terra.StringValue {
	return terra.ReferenceAsString(avecn.ref.Append("vpc_endpoint_service_id"))
}

type awsVpcEndpointConnectionNotificationState struct {
	ConnectionEvents          []string `json:"connection_events"`
	ConnectionNotificationArn string   `json:"connection_notification_arn"`
	Id                        string   `json:"id"`
	NotificationType          string   `json:"notification_type"`
	State                     string   `json:"state"`
	VpcEndpointId             string   `json:"vpc_endpoint_id"`
	VpcEndpointServiceId      string   `json:"vpc_endpoint_service_id"`
}
