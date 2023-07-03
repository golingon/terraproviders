// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	templatedeployment "github.com/golingon/terraproviders/azurerm/3.63.0/templatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewTemplateDeployment creates a new instance of [TemplateDeployment].
func NewTemplateDeployment(name string, args TemplateDeploymentArgs) *TemplateDeployment {
	return &TemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*TemplateDeployment)(nil)

// TemplateDeployment represents the Terraform resource azurerm_template_deployment.
type TemplateDeployment struct {
	Name      string
	Args      TemplateDeploymentArgs
	state     *templateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [TemplateDeployment].
func (td *TemplateDeployment) Type() string {
	return "azurerm_template_deployment"
}

// LocalName returns the local name for [TemplateDeployment].
func (td *TemplateDeployment) LocalName() string {
	return td.Name
}

// Configuration returns the configuration (args) for [TemplateDeployment].
func (td *TemplateDeployment) Configuration() interface{} {
	return td.Args
}

// DependOn is used for other resources to depend on [TemplateDeployment].
func (td *TemplateDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(td)
}

// Dependencies returns the list of resources [TemplateDeployment] depends_on.
func (td *TemplateDeployment) Dependencies() terra.Dependencies {
	return td.DependsOn
}

// LifecycleManagement returns the lifecycle block for [TemplateDeployment].
func (td *TemplateDeployment) LifecycleManagement() *terra.Lifecycle {
	return td.Lifecycle
}

// Attributes returns the attributes for [TemplateDeployment].
func (td *TemplateDeployment) Attributes() templateDeploymentAttributes {
	return templateDeploymentAttributes{ref: terra.ReferenceResource(td)}
}

// ImportState imports the given attribute values into [TemplateDeployment]'s state.
func (td *TemplateDeployment) ImportState(av io.Reader) error {
	td.state = &templateDeploymentState{}
	if err := json.NewDecoder(av).Decode(td.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", td.Type(), td.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [TemplateDeployment] has state.
func (td *TemplateDeployment) State() (*templateDeploymentState, bool) {
	return td.state, td.state != nil
}

// StateMust returns the state for [TemplateDeployment]. Panics if the state is nil.
func (td *TemplateDeployment) StateMust() *templateDeploymentState {
	if td.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", td.Type(), td.LocalName()))
	}
	return td.state
}

// TemplateDeploymentArgs contains the configurations for azurerm_template_deployment.
type TemplateDeploymentArgs struct {
	// DeploymentMode: string, required
	DeploymentMode terra.StringValue `hcl:"deployment_mode,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Parameters: map of string, optional
	Parameters terra.MapValue[terra.StringValue] `hcl:"parameters,attr"`
	// ParametersBody: string, optional
	ParametersBody terra.StringValue `hcl:"parameters_body,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// TemplateBody: string, optional
	TemplateBody terra.StringValue `hcl:"template_body,attr"`
	// Timeouts: optional
	Timeouts *templatedeployment.Timeouts `hcl:"timeouts,block"`
}
type templateDeploymentAttributes struct {
	ref terra.Reference
}

// DeploymentMode returns a reference to field deployment_mode of azurerm_template_deployment.
func (td templateDeploymentAttributes) DeploymentMode() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("deployment_mode"))
}

// Id returns a reference to field id of azurerm_template_deployment.
func (td templateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_template_deployment.
func (td templateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("name"))
}

// Outputs returns a reference to field outputs of azurerm_template_deployment.
func (td templateDeploymentAttributes) Outputs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](td.ref.Append("outputs"))
}

// Parameters returns a reference to field parameters of azurerm_template_deployment.
func (td templateDeploymentAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](td.ref.Append("parameters"))
}

// ParametersBody returns a reference to field parameters_body of azurerm_template_deployment.
func (td templateDeploymentAttributes) ParametersBody() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("parameters_body"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_template_deployment.
func (td templateDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("resource_group_name"))
}

// TemplateBody returns a reference to field template_body of azurerm_template_deployment.
func (td templateDeploymentAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(td.ref.Append("template_body"))
}

func (td templateDeploymentAttributes) Timeouts() templatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[templatedeployment.TimeoutsAttributes](td.ref.Append("timeouts"))
}

type templateDeploymentState struct {
	DeploymentMode    string                            `json:"deployment_mode"`
	Id                string                            `json:"id"`
	Name              string                            `json:"name"`
	Outputs           map[string]string                 `json:"outputs"`
	Parameters        map[string]string                 `json:"parameters"`
	ParametersBody    string                            `json:"parameters_body"`
	ResourceGroupName string                            `json:"resource_group_name"`
	TemplateBody      string                            `json:"template_body"`
	Timeouts          *templatedeployment.TimeoutsState `json:"timeouts"`
}
