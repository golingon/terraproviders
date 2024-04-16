// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package aws_config_organization_custom_rule

import (
	"encoding/json"
	"fmt"
	"github.com/golingon/lingon/pkg/terra"
	"io"
)

// New creates a new instance of [Resource].
func New(name string, args Args) *Resource {
	return &Resource{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*Resource)(nil)

// Resource represents the Terraform resource aws_config_organization_custom_rule.
type Resource struct {
	Name      string
	Args      Args
	state     *awsConfigOrganizationCustomRuleState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (acocr *Resource) Type() string {
	return "aws_config_organization_custom_rule"
}

// LocalName returns the local name for [Resource].
func (acocr *Resource) LocalName() string {
	return acocr.Name
}

// Configuration returns the configuration (args) for [Resource].
func (acocr *Resource) Configuration() interface{} {
	return acocr.Args
}

// DependOn is used for other resources to depend on [Resource].
func (acocr *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(acocr)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (acocr *Resource) Dependencies() terra.Dependencies {
	return acocr.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (acocr *Resource) LifecycleManagement() *terra.Lifecycle {
	return acocr.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (acocr *Resource) Attributes() awsConfigOrganizationCustomRuleAttributes {
	return awsConfigOrganizationCustomRuleAttributes{ref: terra.ReferenceResource(acocr)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (acocr *Resource) ImportState(state io.Reader) error {
	acocr.state = &awsConfigOrganizationCustomRuleState{}
	if err := json.NewDecoder(state).Decode(acocr.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", acocr.Type(), acocr.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (acocr *Resource) State() (*awsConfigOrganizationCustomRuleState, bool) {
	return acocr.state, acocr.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (acocr *Resource) StateMust() *awsConfigOrganizationCustomRuleState {
	if acocr.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", acocr.Type(), acocr.LocalName()))
	}
	return acocr.state
}

// Args contains the configurations for aws_config_organization_custom_rule.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type awsConfigOrganizationCustomRuleAttributes struct {
	ref terra.Reference
}

// Arn returns a reference to field arn of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("arn"))
}

// Description returns a reference to field description of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("description"))
}

// ExcludedAccounts returns a reference to field excluded_accounts of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) ExcludedAccounts() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acocr.ref.Append("excluded_accounts"))
}

// Id returns a reference to field id of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("id"))
}

// InputParameters returns a reference to field input_parameters of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) InputParameters() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("input_parameters"))
}

// LambdaFunctionArn returns a reference to field lambda_function_arn of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) LambdaFunctionArn() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("lambda_function_arn"))
}

// MaximumExecutionFrequency returns a reference to field maximum_execution_frequency of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) MaximumExecutionFrequency() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("maximum_execution_frequency"))
}

// Name returns a reference to field name of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("name"))
}

// ResourceIdScope returns a reference to field resource_id_scope of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) ResourceIdScope() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("resource_id_scope"))
}

// ResourceTypesScope returns a reference to field resource_types_scope of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) ResourceTypesScope() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acocr.ref.Append("resource_types_scope"))
}

// TagKeyScope returns a reference to field tag_key_scope of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) TagKeyScope() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("tag_key_scope"))
}

// TagValueScope returns a reference to field tag_value_scope of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) TagValueScope() terra.StringValue {
	return terra.ReferenceAsString(acocr.ref.Append("tag_value_scope"))
}

// TriggerTypes returns a reference to field trigger_types of aws_config_organization_custom_rule.
func (acocr awsConfigOrganizationCustomRuleAttributes) TriggerTypes() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](acocr.ref.Append("trigger_types"))
}

func (acocr awsConfigOrganizationCustomRuleAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](acocr.ref.Append("timeouts"))
}

type awsConfigOrganizationCustomRuleState struct {
	Arn                       string         `json:"arn"`
	Description               string         `json:"description"`
	ExcludedAccounts          []string       `json:"excluded_accounts"`
	Id                        string         `json:"id"`
	InputParameters           string         `json:"input_parameters"`
	LambdaFunctionArn         string         `json:"lambda_function_arn"`
	MaximumExecutionFrequency string         `json:"maximum_execution_frequency"`
	Name                      string         `json:"name"`
	ResourceIdScope           string         `json:"resource_id_scope"`
	ResourceTypesScope        []string       `json:"resource_types_scope"`
	TagKeyScope               string         `json:"tag_key_scope"`
	TagValueScope             string         `json:"tag_value_scope"`
	TriggerTypes              []string       `json:"trigger_types"`
	Timeouts                  *TimeoutsState `json:"timeouts"`
}
