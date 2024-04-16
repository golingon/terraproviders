// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_route53recoverycontrolconfig_routing_control

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

// Resource represents the Terraform resource aws_route53recoverycontrolconfig_routing_control.
type Resource struct {
	Name      string
	Args      Args
	state     *awsRoute53RecoverycontrolconfigRoutingControlState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (arrc *Resource) Type() string {
	return "aws_route53recoverycontrolconfig_routing_control"
}

// LocalName returns the local name for [Resource].
func (arrc *Resource) LocalName() string {
	return arrc.Name
}

// Configuration returns the configuration (args) for [Resource].
func (arrc *Resource) Configuration() interface{} {
	return arrc.Args
}

// DependOn is used for other resources to depend on [Resource].
func (arrc *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(arrc)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (arrc *Resource) Dependencies() terra.Dependencies {
	return arrc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (arrc *Resource) LifecycleManagement() *terra.Lifecycle {
	return arrc.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (arrc *Resource) Attributes() awsRoute53RecoverycontrolconfigRoutingControlAttributes {
	return awsRoute53RecoverycontrolconfigRoutingControlAttributes{ref: terra.ReferenceResource(arrc)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (arrc *Resource) ImportState(state io.Reader) error {
	arrc.state = &awsRoute53RecoverycontrolconfigRoutingControlState{}
	if err := json.NewDecoder(state).Decode(arrc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", arrc.Type(), arrc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (arrc *Resource) State() (*awsRoute53RecoverycontrolconfigRoutingControlState, bool) {
	return arrc.state, arrc.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (arrc *Resource) StateMust() *awsRoute53RecoverycontrolconfigRoutingControlState {
	if arrc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", arrc.Type(), arrc.LocalName()))
	}
	return arrc.state
}

// Args contains the configurations for aws_route53recoverycontrolconfig_routing_control.
type Args struct {
	// ClusterArn: string, required
	ClusterArn terra.StringValue `hcl:"cluster_arn,attr" validate:"required"`
	// ControlPanelArn: string, optional
	ControlPanelArn terra.StringValue `hcl:"control_panel_arn,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}

type awsRoute53RecoverycontrolconfigRoutingControlAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("arn"))
}

// ClusterArn returns a reference to field cluster_arn of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) ClusterArn() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("cluster_arn"))
}

// ControlPanelArn returns a reference to field control_panel_arn of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) ControlPanelArn() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("control_panel_arn"))
}

// Id returns a reference to field id of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("id"))
}

// Name returns a reference to field name of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("name"))
}

// Status returns a reference to field status of aws_route53recoverycontrolconfig_routing_control.
func (arrc awsRoute53RecoverycontrolconfigRoutingControlAttributes) Status() terra.StringValue {
	return terra.ReferenceAsString(arrc.ref.Append("status"))
}

type awsRoute53RecoverycontrolconfigRoutingControlState struct {
	Arn             string `json:"arn"`
	ClusterArn      string `json:"cluster_arn"`
	ControlPanelArn string `json:"control_panel_arn"`
	Id              string `json:"id"`
	Name            string `json:"name"`
	Status          string `json:"status"`
}
