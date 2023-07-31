// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configorganizationcustompolicyrule "github.com/golingon/terraproviders/aws/5.10.0/configorganizationcustompolicyrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigOrganizationCustomPolicyRule creates a new instance of [ConfigOrganizationCustomPolicyRule].
func NewConfigOrganizationCustomPolicyRule(name string, args ConfigOrganizationCustomPolicyRuleArgs) *ConfigOrganizationCustomPolicyRule {
	return &ConfigOrganizationCustomPolicyRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigOrganizationCustomPolicyRule)(nil)

// ConfigOrganizationCustomPolicyRule represents the Terraform resource aws_config_organization_custom_policy_rule.
type ConfigOrganizationCustomPolicyRule struct {
	Name      string
	Args      ConfigOrganizationCustomPolicyRuleArgs
	state     *configOrganizationCustomPolicyRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) Type() string {
	return "aws_config_organization_custom_policy_rule"
}

// LocalName returns the local name for [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) LocalName() string {
	return cocpr.Name
}

// Configuration returns the configuration (args) for [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) Configuration() interface{} {
	return cocpr.Args
}

// DependOn is used for other resources to depend on [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) DependOn() terra.Reference {
	return terra.ReferenceResource(cocpr)
}

// Dependencies returns the list of resources [ConfigOrganizationCustomPolicyRule] depends_on.
func (cocpr *ConfigOrganizationCustomPolicyRule) Dependencies() terra.Dependencies {
	return cocpr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) LifecycleManagement() *terra.Lifecycle {
	return cocpr.Lifecycle
}

// Attributes returns the attributes for [ConfigOrganizationCustomPolicyRule].
func (cocpr *ConfigOrganizationCustomPolicyRule) Attributes() configOrganizationCustomPolicyRuleAttributes {
	return configOrganizationCustomPolicyRuleAttributes{ref: terra.ReferenceResource(cocpr)}
}

// ImportState imports the given attribute values into [ConfigOrganizationCustomPolicyRule]'s state.
func (cocpr *ConfigOrganizationCustomPolicyRule) ImportState(av io.Reader) error {
	cocpr.state = &configOrganizationCustomPolicyRuleState{}
	if err := json.NewDecoder(av).Decode(cocpr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cocpr.Type(), cocpr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigOrganizationCustomPolicyRule] has state.
func (cocpr *ConfigOrganizationCustomPolicyRule) State() (*configOrganizationCustomPolicyRuleState, bool) {
	return cocpr.state, cocpr.state != nil
}

// StateMust returns the state for [ConfigOrganizationCustomPolicyRule]. Panics if the state is nil.
func (cocpr *ConfigOrganizationCustomPolicyRule) StateMust() *configOrganizationCustomPolicyRuleState {
	if cocpr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cocpr.Type(), cocpr.LocalName()))
	}
	return cocpr.state
}

// ConfigOrganizationCustomPolicyRuleArgs contains the configurations for aws_config_organization_custom_policy_rule.
type ConfigOrganizationCustomPolicyRuleArgs struct {
	// DebugLogDeliveryAccounts: set of string, optional
	DebugLogDeliveryAccounts terra.SetValue[terra.StringValue] `hcl:"debug_log_delivery_accounts,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExcludedAccounts: set of string, optional
	ExcludedAccounts terra.SetValue[terra.StringValue] `hcl:"excluded_accounts,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InputParameters: string, optional
	InputParameters terra.StringValue `hcl:"input_parameters,attr"`
	// MaximumExecutionFrequency: string, optional
	MaximumExecutionFrequency terra.StringValue `hcl:"maximum_execution_frequency,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// PolicyRuntime: string, required
	PolicyRuntime terra.StringValue `hcl:"policy_runtime,attr" validate:"required"`
	// PolicyText: string, required
	PolicyText terra.StringValue `hcl:"policy_text,attr" validate:"required"`
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
	Timeouts *configorganizationcustompolicyrule.Timeouts `hcl:"timeouts,block"`
}
type configOrganizationCustomPolicyRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("arn"))
}

// DebugLogDeliveryAccounts returns a reference to field debug_log_delivery_accounts of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) DebugLogDeliveryAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocpr.ref.Append("debug_log_delivery_accounts"))
}

// Description returns a reference to field description of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("description"))
}

// ExcludedAccounts returns a reference to field excluded_accounts of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) ExcludedAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocpr.ref.Append("excluded_accounts"))
}

// Id returns a reference to field id of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("id"))
}

// InputParameters returns a reference to field input_parameters of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) InputParameters() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("input_parameters"))
}

// MaximumExecutionFrequency returns a reference to field maximum_execution_frequency of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) MaximumExecutionFrequency() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("maximum_execution_frequency"))
}

// Name returns a reference to field name of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("name"))
}

// PolicyRuntime returns a reference to field policy_runtime of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) PolicyRuntime() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("policy_runtime"))
}

// PolicyText returns a reference to field policy_text of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) PolicyText() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("policy_text"))
}

// ResourceIdScope returns a reference to field resource_id_scope of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) ResourceIdScope() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("resource_id_scope"))
}

// ResourceTypesScope returns a reference to field resource_types_scope of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) ResourceTypesScope() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocpr.ref.Append("resource_types_scope"))
}

// TagKeyScope returns a reference to field tag_key_scope of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) TagKeyScope() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("tag_key_scope"))
}

// TagValueScope returns a reference to field tag_value_scope of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) TagValueScope() terra.StringValue {
	return terra.ReferenceAsString(cocpr.ref.Append("tag_value_scope"))
}

// TriggerTypes returns a reference to field trigger_types of aws_config_organization_custom_policy_rule.
func (cocpr configOrganizationCustomPolicyRuleAttributes) TriggerTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](cocpr.ref.Append("trigger_types"))
}

func (cocpr configOrganizationCustomPolicyRuleAttributes) Timeouts() configorganizationcustompolicyrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[configorganizationcustompolicyrule.TimeoutsAttributes](cocpr.ref.Append("timeouts"))
}

type configOrganizationCustomPolicyRuleState struct {
	Arn                       string                                            `json:"arn"`
	DebugLogDeliveryAccounts  []string                                          `json:"debug_log_delivery_accounts"`
	Description               string                                            `json:"description"`
	ExcludedAccounts          []string                                          `json:"excluded_accounts"`
	Id                        string                                            `json:"id"`
	InputParameters           string                                            `json:"input_parameters"`
	MaximumExecutionFrequency string                                            `json:"maximum_execution_frequency"`
	Name                      string                                            `json:"name"`
	PolicyRuntime             string                                            `json:"policy_runtime"`
	PolicyText                string                                            `json:"policy_text"`
	ResourceIdScope           string                                            `json:"resource_id_scope"`
	ResourceTypesScope        []string                                          `json:"resource_types_scope"`
	TagKeyScope               string                                            `json:"tag_key_scope"`
	TagValueScope             string                                            `json:"tag_value_scope"`
	TriggerTypes              []string                                          `json:"trigger_types"`
	Timeouts                  *configorganizationcustompolicyrule.TimeoutsState `json:"timeouts"`
}
