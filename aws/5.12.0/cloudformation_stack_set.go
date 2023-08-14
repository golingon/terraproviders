// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	cloudformationstackset "github.com/golingon/terraproviders/aws/5.12.0/cloudformationstackset"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCloudformationStackSet creates a new instance of [CloudformationStackSet].
func NewCloudformationStackSet(name string, args CloudformationStackSetArgs) *CloudformationStackSet {
	return &CloudformationStackSet{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CloudformationStackSet)(nil)

// CloudformationStackSet represents the Terraform resource aws_cloudformation_stack_set.
type CloudformationStackSet struct {
	Name      string
	Args      CloudformationStackSetArgs
	state     *cloudformationStackSetState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CloudformationStackSet].
func (css *CloudformationStackSet) Type() string {
	return "aws_cloudformation_stack_set"
}

// LocalName returns the local name for [CloudformationStackSet].
func (css *CloudformationStackSet) LocalName() string {
	return css.Name
}

// Configuration returns the configuration (args) for [CloudformationStackSet].
func (css *CloudformationStackSet) Configuration() interface{} {
	return css.Args
}

// DependOn is used for other resources to depend on [CloudformationStackSet].
func (css *CloudformationStackSet) DependOn() terra.Reference {
	return terra.ReferenceResource(css)
}

// Dependencies returns the list of resources [CloudformationStackSet] depends_on.
func (css *CloudformationStackSet) Dependencies() terra.Dependencies {
	return css.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CloudformationStackSet].
func (css *CloudformationStackSet) LifecycleManagement() *terra.Lifecycle {
	return css.Lifecycle
}

// Attributes returns the attributes for [CloudformationStackSet].
func (css *CloudformationStackSet) Attributes() cloudformationStackSetAttributes {
	return cloudformationStackSetAttributes{ref: terra.ReferenceResource(css)}
}

// ImportState imports the given attribute values into [CloudformationStackSet]'s state.
func (css *CloudformationStackSet) ImportState(av io.Reader) error {
	css.state = &cloudformationStackSetState{}
	if err := json.NewDecoder(av).Decode(css.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", css.Type(), css.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CloudformationStackSet] has state.
func (css *CloudformationStackSet) State() (*cloudformationStackSetState, bool) {
	return css.state, css.state != nil
}

// StateMust returns the state for [CloudformationStackSet]. Panics if the state is nil.
func (css *CloudformationStackSet) StateMust() *cloudformationStackSetState {
	if css.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", css.Type(), css.LocalName()))
	}
	return css.state
}

// CloudformationStackSetArgs contains the configurations for aws_cloudformation_stack_set.
type CloudformationStackSetArgs struct {
	// AdministrationRoleArn: string, optional
	AdministrationRoleArn terra.StringValue `hcl:"administration_role_arn,attr"`
	// CallAs: string, optional
	CallAs terra.StringValue `hcl:"call_as,attr"`
	// Capabilities: set of string, optional
	Capabilities terra.SetValue[terra.StringValue] `hcl:"capabilities,attr"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// ExecutionRoleName: string, optional
	ExecutionRoleName terra.StringValue `hcl:"execution_role_name,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// PermissionModel: string, optional
	PermissionModel terra.StringValue `hcl:"permission_model,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
	// TemplateBody: string, optional
	TemplateBody terra.StringValue `hcl:"template_body,attr"`
	// TemplateUrl: string, optional
	TemplateUrl terra.StringValue `hcl:"template_url,attr"`
	// AutoDeployment: optional
	AutoDeployment *cloudformationstackset.AutoDeployment `hcl:"auto_deployment,block"`
	// ManagedExecution: optional
	ManagedExecution *cloudformationstackset.ManagedExecution `hcl:"managed_execution,block"`
	// OperationPreferences: optional
	OperationPreferences *cloudformationstackset.OperationPreferences `hcl:"operation_preferences,block"`
	// Timeouts: optional
	Timeouts *cloudformationstackset.Timeouts `hcl:"timeouts,block"`
}
type cloudformationStackSetAttributes struct {
	ref terra.Reference
}

// AdministrationRoleArn returns a reference to field administration_role_arn of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) AdministrationRoleArn() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("administration_role_arn"))
}

// Arn returns a reference to field arn of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("arn"))
}

// CallAs returns a reference to field call_as of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) CallAs() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("call_as"))
}

// Capabilities returns a reference to field capabilities of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Capabilities() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](css.ref.Append("capabilities"))
}

// Description returns a reference to field description of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("description"))
}

// ExecutionRoleName returns a reference to field execution_role_name of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) ExecutionRoleName() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("execution_role_name"))
}

// Id returns a reference to field id of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("id"))
}

// Name returns a reference to field name of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("name"))
}

// Parameters returns a reference to field parameters of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](css.ref.Append("parameters"))
}

// PermissionModel returns a reference to field permission_model of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) PermissionModel() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("permission_model"))
}

// StackSetId returns a reference to field stack_set_id of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) StackSetId() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("stack_set_id"))
}

// Tags returns a reference to field tags of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](css.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](css.ref.Append("tags_all"))
}

// TemplateBody returns a reference to field template_body of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("template_body"))
}

// TemplateUrl returns a reference to field template_url of aws_cloudformation_stack_set.
func (css cloudformationStackSetAttributes) TemplateUrl() terra.StringValue {
	return terra.ReferenceAsString(css.ref.Append("template_url"))
}

func (css cloudformationStackSetAttributes) AutoDeployment() terra.ListValue[cloudformationstackset.AutoDeploymentAttributes] {
	return terra.ReferenceAsList[cloudformationstackset.AutoDeploymentAttributes](css.ref.Append("auto_deployment"))
}

func (css cloudformationStackSetAttributes) ManagedExecution() terra.ListValue[cloudformationstackset.ManagedExecutionAttributes] {
	return terra.ReferenceAsList[cloudformationstackset.ManagedExecutionAttributes](css.ref.Append("managed_execution"))
}

func (css cloudformationStackSetAttributes) OperationPreferences() terra.ListValue[cloudformationstackset.OperationPreferencesAttributes] {
	return terra.ReferenceAsList[cloudformationstackset.OperationPreferencesAttributes](css.ref.Append("operation_preferences"))
}

func (css cloudformationStackSetAttributes) Timeouts() cloudformationstackset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[cloudformationstackset.TimeoutsAttributes](css.ref.Append("timeouts"))
}

type cloudformationStackSetState struct {
	AdministrationRoleArn string                                             `json:"administration_role_arn"`
	Arn                   string                                             `json:"arn"`
	CallAs                string                                             `json:"call_as"`
	Capabilities          []string                                           `json:"capabilities"`
	Description           string                                             `json:"description"`
	ExecutionRoleName     string                                             `json:"execution_role_name"`
	Id                    string                                             `json:"id"`
	Name                  string                                             `json:"name"`
	Parameters            map[string]string                                  `json:"parameters"`
	PermissionModel       string                                             `json:"permission_model"`
	StackSetId            string                                             `json:"stack_set_id"`
	Tags                  map[string]string                                  `json:"tags"`
	TagsAll               map[string]string                                  `json:"tags_all"`
	TemplateBody          string                                             `json:"template_body"`
	TemplateUrl           string                                             `json:"template_url"`
	AutoDeployment        []cloudformationstackset.AutoDeploymentState       `json:"auto_deployment"`
	ManagedExecution      []cloudformationstackset.ManagedExecutionState     `json:"managed_execution"`
	OperationPreferences  []cloudformationstackset.OperationPreferencesState `json:"operation_preferences"`
	Timeouts              *cloudformationstackset.TimeoutsState              `json:"timeouts"`
}
