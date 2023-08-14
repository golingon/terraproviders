// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewIotLoggingOptions creates a new instance of [IotLoggingOptions].
func NewIotLoggingOptions(name string, args IotLoggingOptionsArgs) *IotLoggingOptions {
	return &IotLoggingOptions{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*IotLoggingOptions)(nil)

// IotLoggingOptions represents the Terraform resource aws_iot_logging_options.
type IotLoggingOptions struct {
	Name      string
	Args      IotLoggingOptionsArgs
	state     *iotLoggingOptionsState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [IotLoggingOptions].
func (ilo *IotLoggingOptions) Type() string {
	return "aws_iot_logging_options"
}

// LocalName returns the local name for [IotLoggingOptions].
func (ilo *IotLoggingOptions) LocalName() string {
	return ilo.Name
}

// Configuration returns the configuration (args) for [IotLoggingOptions].
func (ilo *IotLoggingOptions) Configuration() interface{} {
	return ilo.Args
}

// DependOn is used for other resources to depend on [IotLoggingOptions].
func (ilo *IotLoggingOptions) DependOn() terra.Reference {
	return terra.ReferenceResource(ilo)
}

// Dependencies returns the list of resources [IotLoggingOptions] depends_on.
func (ilo *IotLoggingOptions) Dependencies() terra.Dependencies {
	return ilo.DependsOn
}

// LifecycleManagement returns the lifecycle block for [IotLoggingOptions].
func (ilo *IotLoggingOptions) LifecycleManagement() *terra.Lifecycle {
	return ilo.Lifecycle
}

// Attributes returns the attributes for [IotLoggingOptions].
func (ilo *IotLoggingOptions) Attributes() iotLoggingOptionsAttributes {
	return iotLoggingOptionsAttributes{ref: terra.ReferenceResource(ilo)}
}

// ImportState imports the given attribute values into [IotLoggingOptions]'s state.
func (ilo *IotLoggingOptions) ImportState(av io.Reader) error {
	ilo.state = &iotLoggingOptionsState{}
	if err := json.NewDecoder(av).Decode(ilo.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ilo.Type(), ilo.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [IotLoggingOptions] has state.
func (ilo *IotLoggingOptions) State() (*iotLoggingOptionsState, bool) {
	return ilo.state, ilo.state != nil
}

// StateMust returns the state for [IotLoggingOptions]. Panics if the state is nil.
func (ilo *IotLoggingOptions) StateMust() *iotLoggingOptionsState {
	if ilo.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ilo.Type(), ilo.LocalName()))
	}
	return ilo.state
}

// IotLoggingOptionsArgs contains the configurations for aws_iot_logging_options.
type IotLoggingOptionsArgs struct {
	// DefaultLogLevel: string, required
	DefaultLogLevel terra.StringValue `hcl:"default_log_level,attr" validate:"required"`
	// DisableAllLogs: bool, optional
	DisableAllLogs terra.BoolValue `hcl:"disable_all_logs,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
}
type iotLoggingOptionsAttributes struct {
	ref terra.Reference
}

// DefaultLogLevel returns a reference to field default_log_level of aws_iot_logging_options.
func (ilo iotLoggingOptionsAttributes) DefaultLogLevel() terra.StringValue {
	return terra.ReferenceAsString(ilo.ref.Append("default_log_level"))
}

// DisableAllLogs returns a reference to field disable_all_logs of aws_iot_logging_options.
func (ilo iotLoggingOptionsAttributes) DisableAllLogs() terra.BoolValue {
	return terra.ReferenceAsBool(ilo.ref.Append("disable_all_logs"))
}

// Id returns a reference to field id of aws_iot_logging_options.
func (ilo iotLoggingOptionsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ilo.ref.Append("id"))
}

// RoleArn returns a reference to field role_arn of aws_iot_logging_options.
func (ilo iotLoggingOptionsAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ilo.ref.Append("role_arn"))
}

type iotLoggingOptionsState struct {
	DefaultLogLevel string `json:"default_log_level"`
	DisableAllLogs  bool   `json:"disable_all_logs"`
	Id              string `json:"id"`
	RoleArn         string `json:"role_arn"`
}