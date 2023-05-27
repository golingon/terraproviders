// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	managementgrouptemplatedeployment "github.com/golingon/terraproviders/azurerm/3.58.0/managementgrouptemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewManagementGroupTemplateDeployment creates a new instance of [ManagementGroupTemplateDeployment].
func NewManagementGroupTemplateDeployment(name string, args ManagementGroupTemplateDeploymentArgs) *ManagementGroupTemplateDeployment {
	return &ManagementGroupTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ManagementGroupTemplateDeployment)(nil)

// ManagementGroupTemplateDeployment represents the Terraform resource azurerm_management_group_template_deployment.
type ManagementGroupTemplateDeployment struct {
	Name      string
	Args      ManagementGroupTemplateDeploymentArgs
	state     *managementGroupTemplateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) Type() string {
	return "azurerm_management_group_template_deployment"
}

// LocalName returns the local name for [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) LocalName() string {
	return mgtd.Name
}

// Configuration returns the configuration (args) for [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) Configuration() interface{} {
	return mgtd.Args
}

// DependOn is used for other resources to depend on [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(mgtd)
}

// Dependencies returns the list of resources [ManagementGroupTemplateDeployment] depends_on.
func (mgtd *ManagementGroupTemplateDeployment) Dependencies() terra.Dependencies {
	return mgtd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) LifecycleManagement() *terra.Lifecycle {
	return mgtd.Lifecycle
}

// Attributes returns the attributes for [ManagementGroupTemplateDeployment].
func (mgtd *ManagementGroupTemplateDeployment) Attributes() managementGroupTemplateDeploymentAttributes {
	return managementGroupTemplateDeploymentAttributes{ref: terra.ReferenceResource(mgtd)}
}

// ImportState imports the given attribute values into [ManagementGroupTemplateDeployment]'s state.
func (mgtd *ManagementGroupTemplateDeployment) ImportState(av io.Reader) error {
	mgtd.state = &managementGroupTemplateDeploymentState{}
	if err := json.NewDecoder(av).Decode(mgtd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", mgtd.Type(), mgtd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [ManagementGroupTemplateDeployment] has state.
func (mgtd *ManagementGroupTemplateDeployment) State() (*managementGroupTemplateDeploymentState, bool) {
	return mgtd.state, mgtd.state != nil
}

// StateMust returns the state for [ManagementGroupTemplateDeployment]. Panics if the state is nil.
func (mgtd *ManagementGroupTemplateDeployment) StateMust() *managementGroupTemplateDeploymentState {
	if mgtd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", mgtd.Type(), mgtd.LocalName()))
	}
	return mgtd.state
}

// ManagementGroupTemplateDeploymentArgs contains the configurations for azurerm_management_group_template_deployment.
type ManagementGroupTemplateDeploymentArgs struct {
	// DebugLevel: string, optional
	DebugLevel terra.StringValue `hcl:"debug_level,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
	// ManagementGroupId: string, required
	ManagementGroupId terra.StringValue `hcl:"management_group_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ParametersContent: string, optional
	ParametersContent terra.StringValue `hcl:"parameters_content,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TemplateContent: string, optional
	TemplateContent terra.StringValue `hcl:"template_content,attr"`
	// TemplateSpecVersionId: string, optional
	TemplateSpecVersionId terra.StringValue `hcl:"template_spec_version_id,attr"`
	// Timeouts: optional
	Timeouts *managementgrouptemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type managementGroupTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// DebugLevel returns a reference to field debug_level of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) DebugLevel() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("debug_level"))
}

// Id returns a reference to field id of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("location"))
}

// ManagementGroupId returns a reference to field management_group_id of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) ManagementGroupId() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("management_group_id"))
}

// Name returns a reference to field name of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("output_content"))
}

// ParametersContent returns a reference to field parameters_content of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) ParametersContent() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("parameters_content"))
}

// Tags returns a reference to field tags of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](mgtd.ref.Append("tags"))
}

// TemplateContent returns a reference to field template_content of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) TemplateContent() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("template_content"))
}

// TemplateSpecVersionId returns a reference to field template_spec_version_id of azurerm_management_group_template_deployment.
func (mgtd managementGroupTemplateDeploymentAttributes) TemplateSpecVersionId() terra.StringValue {
	return terra.ReferenceAsString(mgtd.ref.Append("template_spec_version_id"))
}

func (mgtd managementGroupTemplateDeploymentAttributes) Timeouts() managementgrouptemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[managementgrouptemplatedeployment.TimeoutsAttributes](mgtd.ref.Append("timeouts"))
}

type managementGroupTemplateDeploymentState struct {
	DebugLevel            string                                           `json:"debug_level"`
	Id                    string                                           `json:"id"`
	Location              string                                           `json:"location"`
	ManagementGroupId     string                                           `json:"management_group_id"`
	Name                  string                                           `json:"name"`
	OutputContent         string                                           `json:"output_content"`
	ParametersContent     string                                           `json:"parameters_content"`
	Tags                  map[string]string                                `json:"tags"`
	TemplateContent       string                                           `json:"template_content"`
	TemplateSpecVersionId string                                           `json:"template_spec_version_id"`
	Timeouts              *managementgrouptemplatedeployment.TimeoutsState `json:"timeouts"`
}
