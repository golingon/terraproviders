// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_tenant_template_deployment

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

// Resource represents the Terraform resource azurerm_tenant_template_deployment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermTenantTemplateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (attd *Resource) Type() string {
	return "azurerm_tenant_template_deployment"
}

// LocalName returns the local name for [Resource].
func (attd *Resource) LocalName() string {
	return attd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (attd *Resource) Configuration() interface{} {
	return attd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (attd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(attd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (attd *Resource) Dependencies() terra.Dependencies {
	return attd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (attd *Resource) LifecycleManagement() *terra.Lifecycle {
	return attd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (attd *Resource) Attributes() azurermTenantTemplateDeploymentAttributes {
	return azurermTenantTemplateDeploymentAttributes{ref: terra.ReferenceResource(attd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (attd *Resource) ImportState(state io.Reader) error {
	attd.state = &azurermTenantTemplateDeploymentState{}
	if err := json.NewDecoder(state).Decode(attd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", attd.Type(), attd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (attd *Resource) State() (*azurermTenantTemplateDeploymentState, bool) {
	return attd.state, attd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (attd *Resource) StateMust() *azurermTenantTemplateDeploymentState {
	if attd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", attd.Type(), attd.LocalName()))
	}
	return attd.state
}

// Args contains the configurations for azurerm_tenant_template_deployment.
type Args struct {
	// DebugLevel: string, optional
	DebugLevel terra.StringValue `hcl:"debug_level,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, required
	Location terra.StringValue `hcl:"location,attr" validate:"required"`
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
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermTenantTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// DebugLevel returns a reference to field debug_level of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) DebugLevel() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("debug_level"))
}

// Id returns a reference to field id of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("output_content"))
}

// ParametersContent returns a reference to field parameters_content of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) ParametersContent() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("parameters_content"))
}

// Tags returns a reference to field tags of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](attd.ref.Append("tags"))
}

// TemplateContent returns a reference to field template_content of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) TemplateContent() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("template_content"))
}

// TemplateSpecVersionId returns a reference to field template_spec_version_id of azurerm_tenant_template_deployment.
func (attd azurermTenantTemplateDeploymentAttributes) TemplateSpecVersionId() terra.StringValue {
	return terra.ReferenceAsString(attd.ref.Append("template_spec_version_id"))
}

func (attd azurermTenantTemplateDeploymentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](attd.ref.Append("timeouts"))
}

type azurermTenantTemplateDeploymentState struct {
	DebugLevel            string            `json:"debug_level"`
	Id                    string            `json:"id"`
	Location              string            `json:"location"`
	Name                  string            `json:"name"`
	OutputContent         string            `json:"output_content"`
	ParametersContent     string            `json:"parameters_content"`
	Tags                  map[string]string `json:"tags"`
	TemplateContent       string            `json:"template_content"`
	TemplateSpecVersionId string            `json:"template_spec_version_id"`
	Timeouts              *TimeoutsState    `json:"timeouts"`
}
