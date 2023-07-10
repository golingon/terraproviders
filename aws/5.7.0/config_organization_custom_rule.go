// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configorganizationcustomrule "github.com/golingon/terraproviders/aws/5.7.0/configorganizationcustomrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigOrganizationCustomRule creates a new instance of [ConfigOrganizationCustomRule].
func NewConfigOrganizationCustomRule(name string, args ConfigOrganizationCustomRuleArgs) *ConfigOrganizationCustomRule {
	return &ConfigOrganizationCustomRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigOrganizationCustomRule)(nil)

// ConfigOrganizationCustomRule represents the Terraform resource aws_config_organization_custom_rule.
type ConfigOrganizationCustomRule struct {
	Name      string
	Args      ConfigOrganizationCustomRuleArgs
	state     *configOrganizationCustomRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) Type() string {
	return "aws_config_organization_custom_rule"
}

// LocalName returns the local name for [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) LocalName() string {
	return cocr.Name
}

// Configuration returns the configuration (args) for [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) Configuration() interface{} {
	return cocr.Args
}

// DependOn is used for other resources to depend on [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cocr)
}

// Dependencies returns the list of resources [ConfigOrganizationCustomRule] depends_on.
func (cocr *ConfigOrganizationCustomRule) Dependencies() terra.Dependencies {
	return cocr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) LifecycleManagement() *terra.Lifecycle {
	return cocr.Lifecycle
}

// Attributes returns the attributes for [ConfigOrganizationCustomRule].
func (cocr *ConfigOrganizationCustomRule) Attributes() configOrganizationCustomRuleAttributes {
	return configOrganizationCustomRuleAttributes{ref: terra.ReferenceResource(cocr)}
}

// ImportState imports the given attribute values into [ConfigOrganizationCustomRule]'s state.
func (cocr *ConfigOrganizationCustomRule) ImportState(av io.Reader) error {
	cocr.state = &configOrganizationCustomRuleState{}
	if err := json.NewDecoder(av).Decode(cocr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cocr.Type(), cocr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigOrganizationCustomRule] has state.
func (cocr *ConfigOrganizationCustomRule) State() (*configOrganizationCustomRuleState, bool) {
	return cocr.state, cocr.state != nil
}

// StateMust returns the state for [ConfigOrganizationCustomRule]. Panics if the state is nil.
func (cocr *ConfigOrganizationCustomRule) StateMust() *configOrganizationCustomRuleState {
	if cocr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cocr.Type(), cocr.LocalName()))
	}
	return cocr.state
}

// ConfigOrganizationCustomRuleArgs contains the configurations for aws_config_organization_custom_rule.
type ConfigOrganizationCustomRuleArgs struct {
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExcludedAccounts: set of string, optional
	ExcludedAccounts terra.SetValue[terra.StringValue] `hcl:"excluded_accounts,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputParameters: string, optional
	InputParameters terra.StringValue `hcl:"input_parameters,attr"`
	// LambdaFunctionArn: string, required
	LambdaFunctionArn terra.StringValue `hcl:"lambda_function_arn,attr" validate:"required"`
	// MaximumExecutionFrequency: string, optional
	MaximumExecutionFrequency terra.StringValue `hcl:"maximum_execution_frequency,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceIdScope: string, optional
	ResourceIdScope terra.StringValue `hcl:"resource_id_scope,attr"`
	// ResourceTypesScope: set of string, optional
	ResourceTypesScope terra.SetValue[terra.StringValue] `hcl:"resource_types_scope,attr"`
	// TagKeyScope: string, optional
	TagKeyScope terra.StringValue `hcl:"tag_key_scope,attr"`
	// TagValueScope: string, optional
	TagValueScope terra.StringValue `hcl:"tag_value_scope,attr"`
	// TriggerTypes: set of string, required
	TriggerTypes terra.SetValue[terra.StringValue] `hcl:"trigger_types,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *configorganizationcustomrule.Timeouts `hcl:"timeouts,block"`
}
type configOrganizationCustomRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("description"))
}

// ExcludedAccounts returns a reference to field excluded_accounts of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) ExcludedAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocr.ref.Append("excluded_accounts"))
}

// Id returns a reference to field id of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("id"))
}

// InputParameters returns a reference to field input_parameters of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) InputParameters() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("input_parameters"))
}

// LambdaFunctionArn returns a reference to field lambda_function_arn of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) LambdaFunctionArn() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("lambda_function_arn"))
}

// MaximumExecutionFrequency returns a reference to field maximum_execution_frequency of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) MaximumExecutionFrequency() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("maximum_execution_frequency"))
}

// Name returns a reference to field name of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("name"))
}

// ResourceIdScope returns a reference to field resource_id_scope of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) ResourceIdScope() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("resource_id_scope"))
}

// ResourceTypesScope returns a reference to field resource_types_scope of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) ResourceTypesScope() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocr.ref.Append("resource_types_scope"))
}

// TagKeyScope returns a reference to field tag_key_scope of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) TagKeyScope() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("tag_key_scope"))
}

// TagValueScope returns a reference to field tag_value_scope of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) TagValueScope() terra.StringValue {
	return terra.ReferenceAsString(cocr.ref.Append("tag_value_scope"))
}

// TriggerTypes returns a reference to field trigger_types of aws_config_organization_custom_rule.
func (cocr configOrganizationCustomRuleAttributes) TriggerTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocr.ref.Append("trigger_types"))
}

func (cocr configOrganizationCustomRuleAttributes) Timeouts() configorganizationcustomrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[configorganizationcustomrule.TimeoutsAttributes](cocr.ref.Append("timeouts"))
}

type configOrganizationCustomRuleState struct {
	Arn                       string                                      `json:"arn"`
	Description               string                                      `json:"description"`
	ExcludedAccounts          []string                                    `json:"excluded_accounts"`
	Id                        string                                      `json:"id"`
	InputParameters           string                                      `json:"input_parameters"`
	LambdaFunctionArn         string                                      `json:"lambda_function_arn"`
	MaximumExecutionFrequency string                                      `json:"maximum_execution_frequency"`
	Name                      string                                      `json:"name"`
	ResourceIdScope           string                                      `json:"resource_id_scope"`
	ResourceTypesScope        []string                                    `json:"resource_types_scope"`
	TagKeyScope               string                                      `json:"tag_key_scope"`
	TagValueScope             string                                      `json:"tag_value_scope"`
	TriggerTypes              []string                                    `json:"trigger_types"`
	Timeouts                  *configorganizationcustomrule.TimeoutsState `json:"timeouts"`
}
