// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	sagemakerdevice "github.com/golingon/terraproviders/aws/5.44.0/sagemakerdevice"
	"io"
)

// NewSagemakerDevice creates a new instance of [SagemakerDevice].
func NewSagemakerDevice(name string, args SagemakerDeviceArgs) *SagemakerDevice {
	return &SagemakerDevice{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerDevice)(nil)

// SagemakerDevice represents the Terraform resource aws_sagemaker_device.
type SagemakerDevice struct {
	Name      string
	Args      SagemakerDeviceArgs
	state     *sagemakerDeviceState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerDevice].
func (sd *SagemakerDevice) Type() string {
	return "aws_sagemaker_device"
}

// LocalName returns the local name for [SagemakerDevice].
func (sd *SagemakerDevice) LocalName() string {
	return sd.Name
}

// Configuration returns the configuration (args) for [SagemakerDevice].
func (sd *SagemakerDevice) Configuration() interface{} {
	return sd.Args
}

// DependOn is used for other resources to depend on [SagemakerDevice].
func (sd *SagemakerDevice) DependOn() terra.Reference {
	return terra.ReferenceResource(sd)
}

// Dependencies returns the list of resources [SagemakerDevice] depends_on.
func (sd *SagemakerDevice) Dependencies() terra.Dependencies {
	return sd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerDevice].
func (sd *SagemakerDevice) LifecycleManagement() *terra.Lifecycle {
	return sd.Lifecycle
}

// Attributes returns the attributes for [SagemakerDevice].
func (sd *SagemakerDevice) Attributes() sagemakerDeviceAttributes {
	return sagemakerDeviceAttributes{ref: terra.ReferenceResource(sd)}
}

// ImportState imports the given attribute values into [SagemakerDevice]'s state.
func (sd *SagemakerDevice) ImportState(av io.Reader) error {
	sd.state = &sagemakerDeviceState{}
	if err := json.NewDecoder(av).Decode(sd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sd.Type(), sd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerDevice] has state.
func (sd *SagemakerDevice) State() (*sagemakerDeviceState, bool) {
	return sd.state, sd.state != nil
}

// StateMust returns the state for [SagemakerDevice]. Panics if the state is nil.
func (sd *SagemakerDevice) StateMust() *sagemakerDeviceState {
	if sd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sd.Type(), sd.LocalName()))
	}
	return sd.state
}

// SagemakerDeviceArgs contains the configurations for aws_sagemaker_device.
type SagemakerDeviceArgs struct {
	// DeviceFleetName: string, required
	DeviceFleetName terra.StringValue `hcl:"device_fleet_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Device: required
	Device *sagemakerdevice.Device `hcl:"device,block" validate:"required"`
}
type sagemakerDeviceAttributes struct {
	ref terra.Reference
}

// AgentVersion returns a reference to field agent_version of aws_sagemaker_device.
func (sd sagemakerDeviceAttributes) AgentVersion() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("agent_version"))
}

// Arn returns a reference to field arn of aws_sagemaker_device.
func (sd sagemakerDeviceAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("arn"))
}

// DeviceFleetName returns a reference to field device_fleet_name of aws_sagemaker_device.
func (sd sagemakerDeviceAttributes) DeviceFleetName() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("device_fleet_name"))
}

// Id returns a reference to field id of aws_sagemaker_device.
func (sd sagemakerDeviceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sd.ref.Append("id"))
}

func (sd sagemakerDeviceAttributes) Device() terra.ListValue[sagemakerdevice.DeviceAttributes] {
	return terra.ReferenceAsList[sagemakerdevice.DeviceAttributes](sd.ref.Append("device"))
}

type sagemakerDeviceState struct {
	AgentVersion    string                        `json:"agent_version"`
	Arn             string                        `json:"arn"`
	DeviceFleetName string                        `json:"device_fleet_name"`
	Id              string                        `json:"id"`
	Device          []sagemakerdevice.DeviceState `json:"device"`
}
