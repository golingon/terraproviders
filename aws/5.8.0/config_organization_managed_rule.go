// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	configorganizationmanagedrule "github.com/golingon/terraproviders/aws/5.8.0/configorganizationmanagedrule"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewConfigOrganizationManagedRule creates a new instance of [ConfigOrganizationManagedRule].
func NewConfigOrganizationManagedRule(name string, args ConfigOrganizationManagedRuleArgs) *ConfigOrganizationManagedRule {
	return &ConfigOrganizationManagedRule{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ConfigOrganizationManagedRule)(nil)

// ConfigOrganizationManagedRule represents the Terraform resource aws_config_organization_managed_rule.
type ConfigOrganizationManagedRule struct {
	Name      string
	Args      ConfigOrganizationManagedRuleArgs
	state     *configOrganizationManagedRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) Type() string {
	return "aws_config_organization_managed_rule"
}

// LocalName returns the local name for [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) LocalName() string {
	return comr.Name
}

// Configuration returns the configuration (args) for [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) Configuration() interface{} {
	return comr.Args
}

// DependOn is used for other resources to depend on [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) DependOn() terra.Reference {
	return terra.ReferenceResource(comr)
}

// Dependencies returns the list of resources [ConfigOrganizationManagedRule] depends_on.
func (comr *ConfigOrganizationManagedRule) Dependencies() terra.Dependencies {
	return comr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) LifecycleManagement() *terra.Lifecycle {
	return comr.Lifecycle
}

// Attributes returns the attributes for [ConfigOrganizationManagedRule].
func (comr *ConfigOrganizationManagedRule) Attributes() configOrganizationManagedRuleAttributes {
	return configOrganizationManagedRuleAttributes{ref: terra.ReferenceResource(comr)}
}

// ImportState imports the given attribute values into [ConfigOrganizationManagedRule]'s state.
func (comr *ConfigOrganizationManagedRule) ImportState(av io.Reader) error {
	comr.state = &configOrganizationManagedRuleState{}
	if err := json.NewDecoder(av).Decode(comr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", comr.Type(), comr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ConfigOrganizationManagedRule] has state.
func (comr *ConfigOrganizationManagedRule) State() (*configOrganizationManagedRuleState, bool) {
	return comr.state, comr.state != nil
}

// StateMust returns the state for [ConfigOrganizationManagedRule]. Panics if the state is nil.
func (comr *ConfigOrganizationManagedRule) StateMust() *configOrganizationManagedRuleState {
	if comr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", comr.Type(), comr.LocalName()))
	}
	return comr.state
}

// ConfigOrganizationManagedRuleArgs contains the configurations for aws_config_organization_managed_rule.
type ConfigOrganizationManagedRuleArgs struct {
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
	// ResourceIdScope: string, optional
	ResourceIdScope terra.StringValue `hcl:"resource_id_scope,attr"`
	// ResourceTypesScope: set of string, optional
	ResourceTypesScope terra.SetValue[terra.StringValue] `hcl:"resource_types_scope,attr"`
	// RuleIdentifier: string, required
	RuleIdentifier terra.StringValue `hcl:"rule_identifier,attr" validate:"required"`
	// TagKeyScope: string, optional
	TagKeyScope terra.StringValue `hcl:"tag_key_scope,attr"`
	// TagValueScope: string, optional
	TagValueScope terra.StringValue `hcl:"tag_value_scope,attr"`
	// Timeouts: optional
	Timeouts *configorganizationmanagedrule.Timeouts `hcl:"timeouts,block"`
}
type configOrganizationManagedRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("description"))
}

// ExcludedAccounts returns a reference to field excluded_accounts of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) ExcludedAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](comr.ref.Append("excluded_accounts"))
}

// Id returns a reference to field id of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("id"))
}

// InputParameters returns a reference to field input_parameters of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) InputParameters() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("input_parameters"))
}

// MaximumExecutionFrequency returns a reference to field maximum_execution_frequency of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) MaximumExecutionFrequency() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("maximum_execution_frequency"))
}

// Name returns a reference to field name of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("name"))
}

// ResourceIdScope returns a reference to field resource_id_scope of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) ResourceIdScope() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("resource_id_scope"))
}

// ResourceTypesScope returns a reference to field resource_types_scope of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) ResourceTypesScope() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](comr.ref.Append("resource_types_scope"))
}

// RuleIdentifier returns a reference to field rule_identifier of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) RuleIdentifier() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("rule_identifier"))
}

// TagKeyScope returns a reference to field tag_key_scope of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) TagKeyScope() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("tag_key_scope"))
}

// TagValueScope returns a reference to field tag_value_scope of aws_config_organization_managed_rule.
func (comr configOrganizationManagedRuleAttributes) TagValueScope() terra.StringValue {
	return terra.ReferenceAsString(comr.ref.Append("tag_value_scope"))
}

func (comr configOrganizationManagedRuleAttributes) Timeouts() configorganizationmanagedrule.TimeoutsAttributes {
	return terra.ReferenceAsSingle[configorganizationmanagedrule.TimeoutsAttributes](comr.ref.Append("timeouts"))
}

type configOrganizationManagedRuleState struct {
	Arn                       string                                       `json:"arn"`
	Description               string                                       `json:"description"`
	ExcludedAccounts          []string                                     `json:"excluded_accounts"`
	Id                        string                                       `json:"id"`
	InputParameters           string                                       `json:"input_parameters"`
	MaximumExecutionFrequency string                                       `json:"maximum_execution_frequency"`
	Name                      string                                       `json:"name"`
	ResourceIdScope           string                                       `json:"resource_id_scope"`
	ResourceTypesScope        []string                                     `json:"resource_types_scope"`
	RuleIdentifier            string                                       `json:"rule_identifier"`
	TagKeyScope               string                                       `json:"tag_key_scope"`
	TagValueScope             string                                       `json:"tag_value_scope"`
	Timeouts                  *configorganizationmanagedrule.TimeoutsState `json:"timeouts"`
}
