// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigConfigurationRecorderStatus creates a new instance of [ConfigConfigurationRecorderStatus].
func NewConfigConfigurationRecorderStatus(name string, args ConfigConfigurationRecorderStatusArgs) *ConfigConfigurationRecorderStatus {
	return &ConfigConfigurationRecorderStatus{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigConfigurationRecorderStatus)(nil)

// ConfigConfigurationRecorderStatus represents the Terraform resource aws_config_configuration_recorder_status.
type ConfigConfigurationRecorderStatus struct {
	Name      string
	Args      ConfigConfigurationRecorderStatusArgs
	state     *configConfigurationRecorderStatusState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) Type() string {
	return "aws_config_configuration_recorder_status"
}

// LocalName returns the local name for [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) LocalName() string {
	return ccrs.Name
}

// Configuration returns the configuration (args) for [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) Configuration() interface{} {
	return ccrs.Args
}

// DependOn is used for other resources to depend on [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) DependOn() terra.Reference {
	return terra.ReferenceResource(ccrs)
}

// Dependencies returns the list of resources [ConfigConfigurationRecorderStatus] depends_on.
func (ccrs *ConfigConfigurationRecorderStatus) Dependencies() terra.Dependencies {
	return ccrs.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) LifecycleManagement() *terra.Lifecycle {
	return ccrs.Lifecycle
}

// Attributes returns the attributes for [ConfigConfigurationRecorderStatus].
func (ccrs *ConfigConfigurationRecorderStatus) Attributes() configConfigurationRecorderStatusAttributes {
	return configConfigurationRecorderStatusAttributes{ref: terra.ReferenceResource(ccrs)}
}

// ImportState imports the given attribute values into [ConfigConfigurationRecorderStatus]'s state.
func (ccrs *ConfigConfigurationRecorderStatus) ImportState(av io.Reader) error {
	ccrs.state = &configConfigurationRecorderStatusState{}
	if err := json.NewDecoder(av).Decode(ccrs.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ccrs.Type(), ccrs.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigConfigurationRecorderStatus] has state.
func (ccrs *ConfigConfigurationRecorderStatus) State() (*configConfigurationRecorderStatusState, bool) {
	return ccrs.state, ccrs.state != nil
}

// StateMust returns the state for [ConfigConfigurationRecorderStatus]. Panics if the state is nil.
func (ccrs *ConfigConfigurationRecorderStatus) StateMust() *configConfigurationRecorderStatusState {
	if ccrs.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ccrs.Type(), ccrs.LocalName()))
	}
	return ccrs.state
}

// ConfigConfigurationRecorderStatusArgs contains the configurations for aws_config_configuration_recorder_status.
type ConfigConfigurationRecorderStatusArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// IsEnabled: bool, required
	IsEnabled terra.BoolValue `hcl:"is_enabled,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
}
type configConfigurationRecorderStatusAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of aws_config_configuration_recorder_status.
func (ccrs configConfigurationRecorderStatusAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ccrs.ref.Append("id"))
}

// IsEnabled returns a reference to field is_enabled of aws_config_configuration_recorder_status.
func (ccrs configConfigurationRecorderStatusAttributes) IsEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ccrs.ref.Append("is_enabled"))
}

// Name returns a reference to field name of aws_config_configuration_recorder_status.
func (ccrs configConfigurationRecorderStatusAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ccrs.ref.Append("name"))
}

type configConfigurationRecorderStatusState struct {
	Id        string `json:"id"`
	IsEnabled bool   `json:"is_enabled"`
	Name      string `json:"name"`
}