// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	resourcegrouptemplatedeployment "github.com/golingon/terraproviders/azurerm/3.63.0/resourcegrouptemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewResourceGroupTemplateDeployment creates a new instance of [ResourceGroupTemplateDeployment].
func NewResourceGroupTemplateDeployment(name string, args ResourceGroupTemplateDeploymentArgs) *ResourceGroupTemplateDeployment {
	return &ResourceGroupTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ResourceGroupTemplateDeployment)(nil)

// ResourceGroupTemplateDeployment represents the Terraform resource azurerm_resource_group_template_deployment.
type ResourceGroupTemplateDeployment struct {
	Name      string
	Args      ResourceGroupTemplateDeploymentArgs
	state     *resourceGroupTemplateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) Type() string {
	return "azurerm_resource_group_template_deployment"
}

// LocalName returns the local name for [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) LocalName() string {
	return rgtd.Name
}

// Configuration returns the configuration (args) for [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) Configuration() interface{} {
	return rgtd.Args
}

// DependOn is used for other resources to depend on [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(rgtd)
}

// Dependencies returns the list of resources [ResourceGroupTemplateDeployment] depends_on.
func (rgtd *ResourceGroupTemplateDeployment) Dependencies() terra.Dependencies {
	return rgtd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) LifecycleManagement() *terra.Lifecycle {
	return rgtd.Lifecycle
}

// Attributes returns the attributes for [ResourceGroupTemplateDeployment].
func (rgtd *ResourceGroupTemplateDeployment) Attributes() resourceGroupTemplateDeploymentAttributes {
	return resourceGroupTemplateDeploymentAttributes{ref: terra.ReferenceResource(rgtd)}
}

// ImportState imports the given attribute values into [ResourceGroupTemplateDeployment]'s state.
func (rgtd *ResourceGroupTemplateDeployment) ImportState(av io.Reader) error {
	rgtd.state = &resourceGroupTemplateDeploymentState{}
	if err := json.NewDecoder(av).Decode(rgtd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", rgtd.Type(), rgtd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ResourceGroupTemplateDeployment] has state.
func (rgtd *ResourceGroupTemplateDeployment) State() (*resourceGroupTemplateDeploymentState, bool) {
	return rgtd.state, rgtd.state != nil
}

// StateMust returns the state for [ResourceGroupTemplateDeployment]. Panics if the state is nil.
func (rgtd *ResourceGroupTemplateDeployment) StateMust() *resourceGroupTemplateDeploymentState {
	if rgtd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", rgtd.Type(), rgtd.LocalName()))
	}
	return rgtd.state
}

// ResourceGroupTemplateDeploymentArgs contains the configurations for azurerm_resource_group_template_deployment.
type ResourceGroupTemplateDeploymentArgs struct {
	// DebugLevel: string, optional
	DebugLevel terra.StringValue `hcl:"debug_level,attr"`
	// DeploymentMode: string, required
	DeploymentMode terra.StringValue `hcl:"deployment_mode,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParametersContent: string, optional
	ParametersContent terra.StringValue `hcl:"parameters_content,attr"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TemplateContent: string, optional
	TemplateContent terra.StringValue `hcl:"template_content,attr"`
	// TemplateSpecVersionId: string, optional
	TemplateSpecVersionId terra.StringValue `hcl:"template_spec_version_id,attr"`
	// Timeouts: optional
	Timeouts *resourcegrouptemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type resourceGroupTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// DebugLevel returns a reference to field debug_level of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) DebugLevel() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("debug_level"))
}

// DeploymentMode returns a reference to field deployment_mode of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) DeploymentMode() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("deployment_mode"))
}

// Id returns a reference to field id of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("output_content"))
}

// ParametersContent returns a reference to field parameters_content of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) ParametersContent() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("parameters_content"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rgtd.ref.Append("tags"))
}

// TemplateContent returns a reference to field template_content of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) TemplateContent() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("template_content"))
}

// TemplateSpecVersionId returns a reference to field template_spec_version_id of azurerm_resource_group_template_deployment.
func (rgtd resourceGroupTemplateDeploymentAttributes) TemplateSpecVersionId() terra.StringValue {
	return terra.ReferenceAsString(rgtd.ref.Append("template_spec_version_id"))
}

func (rgtd resourceGroupTemplateDeploymentAttributes) Timeouts() resourcegrouptemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[resourcegrouptemplatedeployment.TimeoutsAttributes](rgtd.ref.Append("timeouts"))
}

type resourceGroupTemplateDeploymentState struct {
	DebugLevel            string                                         `json:"debug_level"`
	DeploymentMode        string                                         `json:"deployment_mode"`
	Id                    string                                         `json:"id"`
	Name                  string                                         `json:"name"`
	OutputContent         string                                         `json:"output_content"`
	ParametersContent     string                                         `json:"parameters_content"`
	ResourceGroupName     string                                         `json:"resource_group_name"`
	Tags                  map[string]string                              `json:"tags"`
	TemplateContent       string                                         `json:"template_content"`
	TemplateSpecVersionId string                                         `json:"template_spec_version_id"`
	Timeouts              *resourcegrouptemplatedeployment.TimeoutsState `json:"timeouts"`
}
