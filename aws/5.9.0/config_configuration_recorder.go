// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configconfigurationrecorder "github.com/golingon/terraproviders/aws/5.9.0/configconfigurationrecorder"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigConfigurationRecorder creates a new instance of [ConfigConfigurationRecorder].
func NewConfigConfigurationRecorder(name string, args ConfigConfigurationRecorderArgs) *ConfigConfigurationRecorder {
	return &ConfigConfigurationRecorder{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConfigurationRecorder)(nil)

// ConfigConfigurationRecorder represents the Terraform resource aws_config_configuration_recorder.
type ConfigConfigurationRecorder struct {
	Name      string
	Args      ConfigConfigurationRecorderArgs
	state     *configConfigurationRecorderState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) Type() string {
	return "aws_config_configuration_recorder"
}

// LocalName returns the local name for [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) LocalName() string {
	return ccr.Name
}

// Configuration returns the configuration (args) for [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) Configuration() interface{} {
	return ccr.Args
}

// DependOn is used for other resources to depend on [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) DependOn() terra.Reference {
	return terra.ReferenceResource(ccr)
}

// Dependencies returns the list of resources [ConfigConfigurationRecorder] depends_on.
func (ccr *ConfigConfigurationRecorder) Dependencies() terra.Dependencies {
	return ccr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) LifecycleManagement() *terra.Lifecycle {
	return ccr.Lifecycle
}

// Attributes returns the attributes for [ConfigConfigurationRecorder].
func (ccr *ConfigConfigurationRecorder) Attributes() configConfigurationRecorderAttributes {
	return configConfigurationRecorderAttributes{ref: terra.ReferenceResource(ccr)}
}

// ImportState imports the given attribute values into [ConfigConfigurationRecorder]'s state.
func (ccr *ConfigConfigurationRecorder) ImportState(av io.Reader) error {
	ccr.state = &configConfigurationRecorderState{}
	if err := json.NewDecoder(av).Decode(ccr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccr.Type(), ccr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigConfigurationRecorder] has state.
func (ccr *ConfigConfigurationRecorder) State() (*configConfigurationRecorderState, bool) {
	return ccr.state, ccr.state != nil
}

// StateMust returns the state for [ConfigConfigurationRecorder]. Panics if the state is nil.
func (ccr *ConfigConfigurationRecorder) StateMust() *configConfigurationRecorderState {
	if ccr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccr.Type(), ccr.LocalName()))
	}
	return ccr.state
}

// ConfigConfigurationRecorderArgs contains the configurations for aws_config_configuration_recorder.
type ConfigConfigurationRecorderArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, optional
	Name terra.StringValue `hcl:"name,attr"`
	// RoleArn: string, required
	RoleArn terra.StringValue `hcl:"role_arn,attr" validate:"required"`
	// RecordingGroup: optional
	RecordingGroup *configconfigurationrecorder.RecordingGroup `hcl:"recording_group,block"`
}
type configConfigurationRecorderAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_config_configuration_recorder.
func (ccr configConfigurationRecorderAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("id"))
}

// Name returns a reference to field name of aws_config_configuration_recorder.
func (ccr configConfigurationRecorderAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("name"))
}

// RoleArn returns a reference to field role_arn of aws_config_configuration_recorder.
func (ccr configConfigurationRecorderAttributes) RoleArn() terra.StringValue {
	return terra.ReferenceAsString(ccr.ref.Append("role_arn"))
}

func (ccr configConfigurationRecorderAttributes) RecordingGroup() terra.ListValue[configconfigurationrecorder.RecordingGroupAttributes] {
	return terra.ReferenceAsList[configconfigurationrecorder.RecordingGroupAttributes](ccr.ref.Append("recording_group"))
}

type configConfigurationRecorderState struct {
	Id             string                                            `json:"id"`
	Name           string                                            `json:"name"`
	RoleArn        string                                            `json:"role_arn"`
	RecordingGroup []configconfigurationrecorder.RecordingGroupState `json:"recording_group"`
}