// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_template_deployment

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

// Resource represents the Terraform resource azurerm_template_deployment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermTemplateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (atd *Resource) Type() string {
	return "azurerm_template_deployment"
}

// LocalName returns the local name for [Resource].
func (atd *Resource) LocalName() string {
	return atd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (atd *Resource) Configuration() interface{} {
	return atd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (atd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(atd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (atd *Resource) Dependencies() terra.Dependencies {
	return atd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (atd *Resource) LifecycleManagement() *terra.Lifecycle {
	return atd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (atd *Resource) Attributes() azurermTemplateDeploymentAttributes {
	return azurermTemplateDeploymentAttributes{ref: terra.ReferenceResource(atd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (atd *Resource) ImportState(state io.Reader) error {
	atd.state = &azurermTemplateDeploymentState{}
	if err := json.NewDecoder(state).Decode(atd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", atd.Type(), atd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (atd *Resource) State() (*azurermTemplateDeploymentState, bool) {
	return atd.state, atd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (atd *Resource) StateMust() *azurermTemplateDeploymentState {
	if atd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", atd.Type(), atd.LocalName()))
	}
	return atd.state
}

// Args contains the configurations for azurerm_template_deployment.
type Args struct {
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// DeploymentMode returns a reference to field deployment_mode of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) DeploymentMode() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("deployment_mode"))
}

// Id returns a reference to field id of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("name"))
}

// Outputs returns a reference to field outputs of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) Outputs() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atd.ref.Append("outputs"))
}

// Parameters returns a reference to field parameters of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) Parameters() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](atd.ref.Append("parameters"))
}

// ParametersBody returns a reference to field parameters_body of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) ParametersBody() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("parameters_body"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("resource_group_name"))
}

// TemplateBody returns a reference to field template_body of azurerm_template_deployment.
func (atd azurermTemplateDeploymentAttributes) TemplateBody() terra.StringValue {
	return terra.ReferenceAsString(atd.ref.Append("template_body"))
}

func (atd azurermTemplateDeploymentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](atd.ref.Append("timeouts"))
}

type azurermTemplateDeploymentState struct {
	DeploymentMode    string            `json:"deployment_mode"`
	Id                string            `json:"id"`
	Name              string            `json:"name"`
	Outputs           map[string]string `json:"outputs"`
	Parameters        map[string]string `json:"parameters"`
	ParametersBody    string            `json:"parameters_body"`
	ResourceGroupName string            `json:"resource_group_name"`
	TemplateBody      string            `json:"template_body"`
	Timeouts          *TimeoutsState    `json:"timeouts"`
}
