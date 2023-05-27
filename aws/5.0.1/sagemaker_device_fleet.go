// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	sagemakerdevicefleet "github.com/golingon/terraproviders/aws/5.0.1/sagemakerdevicefleet"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSagemakerDeviceFleet creates a new instance of [SagemakerDeviceFleet].
func NewSagemakerDeviceFleet(name string, args SagemakerDeviceFleetArgs) *SagemakerDeviceFleet {
	return &SagemakerDeviceFleet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SagemakerDeviceFleet)(nil)

// SagemakerDeviceFleet represents the Terraform resource aws_sagemaker_device_fleet.
type SagemakerDeviceFleet struct {
	Name      string
	Args      SagemakerDeviceFleetArgs
	state     *sagemakerDeviceFleetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) Type() string {
	return "aws_sagemaker_device_fleet"
}

// LocalName returns the local name for [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) LocalName() string {
	return sdf.Name
}

// Configuration returns the configuration (args) for [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) Configuration() interface{} {
	return sdf.Args
}

// DependOn is used for other resources to depend on [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) DependOn() terra.Reference {
	return terra.ReferenceResource(sdf)
}

// Dependencies returns the list of resources [SagemakerDeviceFleet] depends_on.
func (sdf *SagemakerDeviceFleet) Dependencies() terra.Dependencies {
	return sdf.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) LifecycleManagement() *terra.Lifecycle {
	return sdf.Lifecycle
}

// Attributes returns the attributes for [SagemakerDeviceFleet].
func (sdf *SagemakerDeviceFleet) Attributes() sagemakerDeviceFleetAttributes {
	return sagemakerDeviceFleetAttributes{ref: terra.ReferenceResource(sdf)}
}

// ImportState imports the given attribute values into [SagemakerDeviceFleet]'s state.
func (sdf *SagemakerDeviceFleet) ImportState(av io.Reader) error {
	sdf.state = &sagemakerDeviceFleetState{}
	if err := json.NewDecoder(av).Decode(sdf.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sdf.Type(), sdf.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SagemakerDeviceFleet] has state.
func (sdf *SagemakerDeviceFleet) State() (*sagemakerDeviceFleetState, bool) {
	return sdf.state, sdf.state != nil
}

// StateMust returns the state for [SagemakerDeviceFleet]. Panics if the state is nil.
func (sdf *SagemakerDeviceFleet) StateMust() *sagemakerDeviceFleetState {
	if sdf.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sdf.Type(), sdf.LocalName()))
	}
	return sdf.state
}

// SagemakerDeviceFleetArgs contains the configurations for aws_sagemaker_device_fleet.
type SagemakerDeviceFleetArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// DeviceFleetName: string, required
	DeviceFleetName terra.StringValue `hcl:"device_fleet_name,attr" validate:"required"`
	// EnableIotRoleAlias: bool, optional
	EnableIotRoleAlias terra.BoolValue `hcl:"enable_iot_role_alias,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// OutputConfig: required
	OutputConfig *sagemakerdevicefleet.OutputConfig `hcl:"output_config,block" validate:"required"`
}
type sagemakerDeviceFleetAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("arn"))
}

// Description returns a reference to field description of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("description"))
}

// DeviceFleetName returns a reference to field device_fleet_name of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) DeviceFleetName() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("device_fleet_name"))
}

// EnableIotRoleAlias returns a reference to field enable_iot_role_alias of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) EnableIotRoleAlias() terra.BoolValue {
	return terra.ReferenceAsBool(sdf.ref.Append("enable_iot_role_alias"))
}

// Id returns a reference to field id of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("id"))
}

// IotRoleAlias returns a reference to field iot_role_alias of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) IotRoleAlias() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("iot_role_alias"))
}

// RoleArn returns a reference to field role_arn of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(sdf.ref.Append("role_arn"))
}

// Tags returns a reference to field tags of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdf.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_sagemaker_device_fleet.
func (sdf sagemakerDeviceFleetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sdf.ref.Append("tags_all"))
}

func (sdf sagemakerDeviceFleetAttributes) OutputConfig() terra.ListValue[sagemakerdevicefleet.OutputConfigAttributes] {
	return terra.ReferenceAsList[sagemakerdevicefleet.OutputConfigAttributes](sdf.ref.Append("output_config"))
}

type sagemakerDeviceFleetState struct {
	Arn                string                                   `json:"arn"`
	Description        string                                   `json:"description"`
	DeviceFleetName    string                                   `json:"device_fleet_name"`
	EnableIotRoleAlias bool                                     `json:"enable_iot_role_alias"`
	Id                 string                                   `json:"id"`
	IotRoleAlias       string                                   `json:"iot_role_alias"`
	RoleArn            string                                   `json:"role_arn"`
	Tags               map[string]string                        `json:"tags"`
	TagsAll            map[string]string                        `json:"tags_all"`
	OutputConfig       []sagemakerdevicefleet.OutputConfigState `json:"output_config"`
}
