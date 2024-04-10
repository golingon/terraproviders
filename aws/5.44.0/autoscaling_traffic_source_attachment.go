// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	autoscalingtrafficsourceattachment "github.com/golingon/terraproviders/aws/5.44.0/autoscalingtrafficsourceattachment"
	"io"
)

// NewAutoscalingTrafficSourceAttachment creates a new instance of [AutoscalingTrafficSourceAttachment].
func NewAutoscalingTrafficSourceAttachment(name string, args AutoscalingTrafficSourceAttachmentArgs) *AutoscalingTrafficSourceAttachment {
	return &AutoscalingTrafficSourceAttachment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AutoscalingTrafficSourceAttachment)(nil)

// AutoscalingTrafficSourceAttachment represents the Terraform resource aws_autoscaling_traffic_source_attachment.
type AutoscalingTrafficSourceAttachment struct {
	Name      string
	Args      AutoscalingTrafficSourceAttachmentArgs
	state     *autoscalingTrafficSourceAttachmentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) Type() string {
	return "aws_autoscaling_traffic_source_attachment"
}

// LocalName returns the local name for [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) LocalName() string {
	return atsa.Name
}

// Configuration returns the configuration (args) for [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) Configuration() interface{} {
	return atsa.Args
}

// DependOn is used for other resources to depend on [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) DependOn() terra.Reference {
	return terra.ReferenceResource(atsa)
}

// Dependencies returns the list of resources [AutoscalingTrafficSourceAttachment] depends_on.
func (atsa *AutoscalingTrafficSourceAttachment) Dependencies() terra.Dependencies {
	return atsa.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) LifecycleManagement() *terra.Lifecycle {
	return atsa.Lifecycle
}

// Attributes returns the attributes for [AutoscalingTrafficSourceAttachment].
func (atsa *AutoscalingTrafficSourceAttachment) Attributes() autoscalingTrafficSourceAttachmentAttributes {
	return autoscalingTrafficSourceAttachmentAttributes{ref: terra.ReferenceResource(atsa)}
}

// ImportState imports the given attribute values into [AutoscalingTrafficSourceAttachment]'s state.
func (atsa *AutoscalingTrafficSourceAttachment) ImportState(av io.Reader) error {
	atsa.state = &autoscalingTrafficSourceAttachmentState{}
	if err := json.NewDecoder(av).Decode(atsa.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atsa.Type(), atsa.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AutoscalingTrafficSourceAttachment] has state.
func (atsa *AutoscalingTrafficSourceAttachment) State() (*autoscalingTrafficSourceAttachmentState, bool) {
	return atsa.state, atsa.state != nil
}

// StateMust returns the state for [AutoscalingTrafficSourceAttachment]. Panics if the state is nil.
func (atsa *AutoscalingTrafficSourceAttachment) StateMust() *autoscalingTrafficSourceAttachmentState {
	if atsa.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atsa.Type(), atsa.LocalName()))
	}
	return atsa.state
}

// AutoscalingTrafficSourceAttachmentArgs contains the configurations for aws_autoscaling_traffic_source_attachment.
type AutoscalingTrafficSourceAttachmentArgs struct {
	// AutoscalingGroupName: string, required
	AutoscalingGroupName terra.StringValue `hcl:"autoscaling_group_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Timeouts: optional
	Timeouts *autoscalingtrafficsourceattachment.Timeouts `hcl:"timeouts,block"`
	// TrafficSource: optional
	TrafficSource *autoscalingtrafficsourceattachment.TrafficSource `hcl:"traffic_source,block"`
}
type autoscalingTrafficSourceAttachmentAttributes struct {
	ref terra.Reference
}

// AutoscalingGroupName returns a reference to field autoscaling_group_name of aws_autoscaling_traffic_source_attachment.
func (atsa autoscalingTrafficSourceAttachmentAttributes) AutoscalingGroupName() terra.StringValue {
	return terra.ReferenceAsString(atsa.ref.Append("autoscaling_group_name"))
}

// Id returns a reference to field id of aws_autoscaling_traffic_source_attachment.
func (atsa autoscalingTrafficSourceAttachmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atsa.ref.Append("id"))
}

func (atsa autoscalingTrafficSourceAttachmentAttributes) Timeouts() autoscalingtrafficsourceattachment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[autoscalingtrafficsourceattachment.TimeoutsAttributes](atsa.ref.Append("timeouts"))
}

func (atsa autoscalingTrafficSourceAttachmentAttributes) TrafficSource() terra.ListValue[autoscalingtrafficsourceattachment.TrafficSourceAttributes] {
	return terra.ReferenceAsList[autoscalingtrafficsourceattachment.TrafficSourceAttributes](atsa.ref.Append("traffic_source"))
}

type autoscalingTrafficSourceAttachmentState struct {
	AutoscalingGroupName string                                                  `json:"autoscaling_group_name"`
	Id                   string                                                  `json:"id"`
	Timeouts             *autoscalingtrafficsourceattachment.TimeoutsState       `json:"timeouts"`
	TrafficSource        []autoscalingtrafficsourceattachment.TrafficSourceState `json:"traffic_source"`
}
