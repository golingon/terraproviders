// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configremediationconfiguration "github.com/golingon/terraproviders/aws/5.9.0/configremediationconfiguration"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigRemediationConfiguration creates a new instance of [ConfigRemediationConfiguration].
func NewConfigRemediationConfiguration(name string, args ConfigRemediationConfigurationArgs) *ConfigRemediationConfiguration {
	return &ConfigRemediationConfiguration{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigRemediationConfiguration)(nil)

// ConfigRemediationConfiguration represents the Terraform resource aws_config_remediation_configuration.
type ConfigRemediationConfiguration struct {
	Name      string
	Args      ConfigRemediationConfigurationArgs
	state     *configRemediationConfigurationState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) Type() string {
	return "aws_config_remediation_configuration"
}

// LocalName returns the local name for [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) LocalName() string {
	return crc.Name
}

// Configuration returns the configuration (args) for [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) Configuration() interface{} {
	return crc.Args
}

// DependOn is used for other resources to depend on [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) DependOn() terra.Reference {
	return terra.ReferenceResource(crc)
}

// Dependencies returns the list of resources [ConfigRemediationConfiguration] depends_on.
func (crc *ConfigRemediationConfiguration) Dependencies() terra.Dependencies {
	return crc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) LifecycleManagement() *terra.Lifecycle {
	return crc.Lifecycle
}

// Attributes returns the attributes for [ConfigRemediationConfiguration].
func (crc *ConfigRemediationConfiguration) Attributes() configRemediationConfigurationAttributes {
	return configRemediationConfigurationAttributes{ref: terra.ReferenceResource(crc)}
}

// ImportState imports the given attribute values into [ConfigRemediationConfiguration]'s state.
func (crc *ConfigRemediationConfiguration) ImportState(av io.Reader) error {
	crc.state = &configRemediationConfigurationState{}
	if err := json.NewDecoder(av).Decode(crc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", crc.Type(), crc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigRemediationConfiguration] has state.
func (crc *ConfigRemediationConfiguration) State() (*configRemediationConfigurationState, bool) {
	return crc.state, crc.state != nil
}

// StateMust returns the state for [ConfigRemediationConfiguration]. Panics if the state is nil.
func (crc *ConfigRemediationConfiguration) StateMust() *configRemediationConfigurationState {
	if crc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", crc.Type(), crc.LocalName()))
	}
	return crc.state
}

// ConfigRemediationConfigurationArgs contains the configurations for aws_config_remediation_configuration.
type ConfigRemediationConfigurationArgs struct {
	// Automatic: bool, optional
	Automatic terra.BoolValue `hcl:"automatic,attr"`
	// ConfigRuleName: string, required
	ConfigRuleName terra.StringValue `hcl:"config_rule_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MaximumAutomaticAttempts: number, optional
	MaximumAutomaticAttempts terra.NumberValue `hcl:"maximum_automatic_attempts,attr"`
	// ResourceType: string, optional
	ResourceType terra.StringValue `hcl:"resource_type,attr"`
	// RetryAttemptSeconds: number, optional
	RetryAttemptSeconds terra.NumberValue `hcl:"retry_attempt_seconds,attr"`
	// TargetId: string, required
	TargetId terra.StringValue `hcl:"target_id,attr" validate:"required"`
	// TargetType: string, required
	TargetType terra.StringValue `hcl:"target_type,attr" validate:"required"`
	// TargetVersion: string, optional
	TargetVersion terra.StringValue `hcl:"target_version,attr"`
	// ExecutionControls: optional
	ExecutionControls *configremediationconfiguration.ExecutionControls `hcl:"execution_controls,block"`
	// Parameter: min=0,max=25
	Parameter []configremediationconfiguration.Parameter `hcl:"parameter,block" validate:"min=0,max=25"`
}
type configRemediationConfigurationAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("arn"))
}

// Automatic returns a reference to field automatic of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) Automatic() terra.BoolValue {
	return terra.ReferenceAsBool(crc.ref.Append("automatic"))
}

// ConfigRuleName returns a reference to field config_rule_name of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) ConfigRuleName() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("config_rule_name"))
}

// Id returns a reference to field id of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("id"))
}

// MaximumAutomaticAttempts returns a reference to field maximum_automatic_attempts of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) MaximumAutomaticAttempts() terra.NumberValue {
	return terra.ReferenceAsNumber(crc.ref.Append("maximum_automatic_attempts"))
}

// ResourceType returns a reference to field resource_type of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) ResourceType() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("resource_type"))
}

// RetryAttemptSeconds returns a reference to field retry_attempt_seconds of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) RetryAttemptSeconds() terra.NumberValue {
	return terra.ReferenceAsNumber(crc.ref.Append("retry_attempt_seconds"))
}

// TargetId returns a reference to field target_id of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) TargetId() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("target_id"))
}

// TargetType returns a reference to field target_type of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) TargetType() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("target_type"))
}

// TargetVersion returns a reference to field target_version of aws_config_remediation_configuration.
func (crc configRemediationConfigurationAttributes) TargetVersion() terra.StringValue {
	return terra.ReferenceAsString(crc.ref.Append("target_version"))
}

func (crc configRemediationConfigurationAttributes) ExecutionControls() terra.ListValue[configremediationconfiguration.ExecutionControlsAttributes] {
	return terra.ReferenceAsList[configremediationconfiguration.ExecutionControlsAttributes](crc.ref.Append("execution_controls"))
}

func (crc configRemediationConfigurationAttributes) Parameter() terra.ListValue[configremediationconfiguration.ParameterAttributes] {
	return terra.ReferenceAsList[configremediationconfiguration.ParameterAttributes](crc.ref.Append("parameter"))
}

type configRemediationConfigurationState struct {
	Arn                      string                                                  `json:"arn"`
	Automatic                bool                                                    `json:"automatic"`
	ConfigRuleName           string                                                  `json:"config_rule_name"`
	Id                       string                                                  `json:"id"`
	MaximumAutomaticAttempts float64                                                 `json:"maximum_automatic_attempts"`
	ResourceType             string                                                  `json:"resource_type"`
	RetryAttemptSeconds      float64                                                 `json:"retry_attempt_seconds"`
	TargetId                 string                                                  `json:"target_id"`
	TargetType               string                                                  `json:"target_type"`
	TargetVersion            string                                                  `json:"target_version"`
	ExecutionControls        []configremediationconfiguration.ExecutionControlsState `json:"execution_controls"`
	Parameter                []configremediationconfiguration.ParameterState         `json:"parameter"`
}
