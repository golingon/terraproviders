// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	globalacceleratorendpointgroup "github.com/golingon/terraproviders/aws/5.7.0/globalacceleratorendpointgroup"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewGlobalacceleratorEndpointGroup creates a new instance of [GlobalacceleratorEndpointGroup].
func NewGlobalacceleratorEndpointGroup(name string, args GlobalacceleratorEndpointGroupArgs) *GlobalacceleratorEndpointGroup {
	return &GlobalacceleratorEndpointGroup{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*GlobalacceleratorEndpointGroup)(nil)

// GlobalacceleratorEndpointGroup represents the Terraform resource aws_globalaccelerator_endpoint_group.
type GlobalacceleratorEndpointGroup struct {
	Name      string
	Args      GlobalacceleratorEndpointGroupArgs
	state     *globalacceleratorEndpointGroupState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) Type() string {
	return "aws_globalaccelerator_endpoint_group"
}

// LocalName returns the local name for [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) LocalName() string {
	return geg.Name
}

// Configuration returns the configuration (args) for [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) Configuration() interface{} {
	return geg.Args
}

// DependOn is used for other resources to depend on [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) DependOn() terra.Reference {
	return terra.ReferenceResource(geg)
}

// Dependencies returns the list of resources [GlobalacceleratorEndpointGroup] depends_on.
func (geg *GlobalacceleratorEndpointGroup) Dependencies() terra.Dependencies {
	return geg.DependsOn
}

// LifecycleManagement returns the lifecycle block for [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) LifecycleManagement() *terra.Lifecycle {
	return geg.Lifecycle
}

// Attributes returns the attributes for [GlobalacceleratorEndpointGroup].
func (geg *GlobalacceleratorEndpointGroup) Attributes() globalacceleratorEndpointGroupAttributes {
	return globalacceleratorEndpointGroupAttributes{ref: terra.ReferenceResource(geg)}
}

// ImportState imports the given attribute values into [GlobalacceleratorEndpointGroup]'s state.
func (geg *GlobalacceleratorEndpointGroup) ImportState(av io.Reader) error {
	geg.state = &globalacceleratorEndpointGroupState{}
	if err := json.NewDecoder(av).Decode(geg.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", geg.Type(), geg.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [GlobalacceleratorEndpointGroup] has state.
func (geg *GlobalacceleratorEndpointGroup) State() (*globalacceleratorEndpointGroupState, bool) {
	return geg.state, geg.state != nil
}

// StateMust returns the state for [GlobalacceleratorEndpointGroup]. Panics if the state is nil.
func (geg *GlobalacceleratorEndpointGroup) StateMust() *globalacceleratorEndpointGroupState {
	if geg.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", geg.Type(), geg.LocalName()))
	}
	return geg.state
}

// GlobalacceleratorEndpointGroupArgs contains the configurations for aws_globalaccelerator_endpoint_group.
type GlobalacceleratorEndpointGroupArgs struct {
	// EndpointGroupRegion: string, optional
	EndpointGroupRegion terra.StringValue `hcl:"endpoint_group_region,attr"`
	// HealthCheckIntervalSeconds: number, optional
	HealthCheckIntervalSeconds terra.NumberValue `hcl:"health_check_interval_seconds,attr"`
	// HealthCheckPath: string, optional
	HealthCheckPath terra.StringValue `hcl:"health_check_path,attr"`
	// HealthCheckPort: number, optional
	HealthCheckPort terra.NumberValue `hcl:"health_check_port,attr"`
	// HealthCheckProtocol: string, optional
	HealthCheckProtocol terra.StringValue `hcl:"health_check_protocol,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ListenerArn: string, required
	ListenerArn terra.StringValue `hcl:"listener_arn,attr" validate:"required"`
	// ThresholdCount: number, optional
	ThresholdCount terra.NumberValue `hcl:"threshold_count,attr"`
	// TrafficDialPercentage: number, optional
	TrafficDialPercentage terra.NumberValue `hcl:"traffic_dial_percentage,attr"`
	// EndpointConfiguration: min=0
	EndpointConfiguration []globalacceleratorendpointgroup.EndpointConfiguration `hcl:"endpoint_configuration,block" validate:"min=0"`
	// PortOverride: min=0,max=10
	PortOverride []globalacceleratorendpointgroup.PortOverride `hcl:"port_override,block" validate:"min=0,max=10"`
	// Timeouts: optional
	Timeouts *globalacceleratorendpointgroup.Timeouts `hcl:"timeouts,block"`
}
type globalacceleratorEndpointGroupAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("arn"))
}

// EndpointGroupRegion returns a reference to field endpoint_group_region of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) EndpointGroupRegion() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("endpoint_group_region"))
}

// HealthCheckIntervalSeconds returns a reference to field health_check_interval_seconds of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) HealthCheckIntervalSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(geg.ref.Append("health_check_interval_seconds"))
}

// HealthCheckPath returns a reference to field health_check_path of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) HealthCheckPath() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("health_check_path"))
}

// HealthCheckPort returns a reference to field health_check_port of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) HealthCheckPort() terra.NumberValue {
	return terra.ReferenceAsNumber(geg.ref.Append("health_check_port"))
}

// HealthCheckProtocol returns a reference to field health_check_protocol of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) HealthCheckProtocol() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("health_check_protocol"))
}

// Id returns a reference to field id of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("id"))
}

// ListenerArn returns a reference to field listener_arn of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) ListenerArn() terra.StringValue {
	return terra.ReferenceAsString(geg.ref.Append("listener_arn"))
}

// ThresholdCount returns a reference to field threshold_count of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) ThresholdCount() terra.NumberValue {
	return terra.ReferenceAsNumber(geg.ref.Append("threshold_count"))
}

// TrafficDialPercentage returns a reference to field traffic_dial_percentage of aws_globalaccelerator_endpoint_group.
func (geg globalacceleratorEndpointGroupAttributes) TrafficDialPercentage() terra.NumberValue {
	return terra.ReferenceAsNumber(geg.ref.Append("traffic_dial_percentage"))
}

func (geg globalacceleratorEndpointGroupAttributes) EndpointConfiguration() terra.SetValue[globalacceleratorendpointgroup.EndpointConfigurationAttributes] {
	return terra.ReferenceAsSet[globalacceleratorendpointgroup.EndpointConfigurationAttributes](geg.ref.Append("endpoint_configuration"))
}

func (geg globalacceleratorEndpointGroupAttributes) PortOverride() terra.SetValue[globalacceleratorendpointgroup.PortOverrideAttributes] {
	return terra.ReferenceAsSet[globalacceleratorendpointgroup.PortOverrideAttributes](geg.ref.Append("port_override"))
}

func (geg globalacceleratorEndpointGroupAttributes) Timeouts() globalacceleratorendpointgroup.TimeoutsAttributes {
	return terra.ReferenceAsSingle[globalacceleratorendpointgroup.TimeoutsAttributes](geg.ref.Append("timeouts"))
}

type globalacceleratorEndpointGroupState struct {
	Arn                        string                                                      `json:"arn"`
	EndpointGroupRegion        string                                                      `json:"endpoint_group_region"`
	HealthCheckIntervalSeconds float64                                                     `json:"health_check_interval_seconds"`
	HealthCheckPath            string                                                      `json:"health_check_path"`
	HealthCheckPort            float64                                                     `json:"health_check_port"`
	HealthCheckProtocol        string                                                      `json:"health_check_protocol"`
	Id                         string                                                      `json:"id"`
	ListenerArn                string                                                      `json:"listener_arn"`
	ThresholdCount             float64                                                     `json:"threshold_count"`
	TrafficDialPercentage      float64                                                     `json:"traffic_dial_percentage"`
	EndpointConfiguration      []globalacceleratorendpointgroup.EndpointConfigurationState `json:"endpoint_configuration"`
	PortOverride               []globalacceleratorendpointgroup.PortOverrideState          `json:"port_override"`
	Timeouts                   *globalacceleratorendpointgroup.TimeoutsState               `json:"timeouts"`
}
