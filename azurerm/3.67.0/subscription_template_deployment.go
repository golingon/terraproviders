// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	subscriptiontemplatedeployment "github.com/golingon/terraproviders/azurerm/3.67.0/subscriptiontemplatedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSubscriptionTemplateDeployment creates a new instance of [SubscriptionTemplateDeployment].
func NewSubscriptionTemplateDeployment(name string, args SubscriptionTemplateDeploymentArgs) *SubscriptionTemplateDeployment {
	return &SubscriptionTemplateDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SubscriptionTemplateDeployment)(nil)

// SubscriptionTemplateDeployment represents the Terraform resource azurerm_subscription_template_deployment.
type SubscriptionTemplateDeployment struct {
	Name      string
	Args      SubscriptionTemplateDeploymentArgs
	state     *subscriptionTemplateDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) Type() string {
	return "azurerm_subscription_template_deployment"
}

// LocalName returns the local name for [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) LocalName() string {
	return std.Name
}

// Configuration returns the configuration (args) for [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) Configuration() interface{} {
	return std.Args
}

// DependOn is used for other resources to depend on [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(std)
}

// Dependencies returns the list of resources [SubscriptionTemplateDeployment] depends_on.
func (std *SubscriptionTemplateDeployment) Dependencies() terra.Dependencies {
	return std.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) LifecycleManagement() *terra.Lifecycle {
	return std.Lifecycle
}

// Attributes returns the attributes for [SubscriptionTemplateDeployment].
func (std *SubscriptionTemplateDeployment) Attributes() subscriptionTemplateDeploymentAttributes {
	return subscriptionTemplateDeploymentAttributes{ref: terra.ReferenceResource(std)}
}

// ImportState imports the given attribute values into [SubscriptionTemplateDeployment]'s state.
func (std *SubscriptionTemplateDeployment) ImportState(av io.Reader) error {
	std.state = &subscriptionTemplateDeploymentState{}
	if err := json.NewDecoder(av).Decode(std.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", std.Type(), std.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SubscriptionTemplateDeployment] has state.
func (std *SubscriptionTemplateDeployment) State() (*subscriptionTemplateDeploymentState, bool) {
	return std.state, std.state != nil
}

// StateMust returns the state for [SubscriptionTemplateDeployment]. Panics if the state is nil.
func (std *SubscriptionTemplateDeployment) StateMust() *subscriptionTemplateDeploymentState {
	if std.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", std.Type(), std.LocalName()))
	}
	return std.state
}

// SubscriptionTemplateDeploymentArgs contains the configurations for azurerm_subscription_template_deployment.
type SubscriptionTemplateDeploymentArgs struct {
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
	Timeouts *subscriptiontemplatedeployment.Timeouts `hcl:"timeouts,block"`
}
type subscriptionTemplateDeploymentAttributes struct {
	ref terra.Reference
}

// DebugLevel returns a reference to field debug_level of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) DebugLevel() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("debug_level"))
}

// Id returns a reference to field id of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("name"))
}

// OutputContent returns a reference to field output_content of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) OutputContent() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("output_content"))
}

// ParametersContent returns a reference to field parameters_content of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) ParametersContent() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("parameters_content"))
}

// Tags returns a reference to field tags of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](std.ref.Append("tags"))
}

// TemplateContent returns a reference to field template_content of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) TemplateContent() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("template_content"))
}

// TemplateSpecVersionId returns a reference to field template_spec_version_id of azurerm_subscription_template_deployment.
func (std subscriptionTemplateDeploymentAttributes) TemplateSpecVersionId() terra.StringValue {
	return terra.ReferenceAsString(std.ref.Append("template_spec_version_id"))
}

func (std subscriptionTemplateDeploymentAttributes) Timeouts() subscriptiontemplatedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[subscriptiontemplatedeployment.TimeoutsAttributes](std.ref.Append("timeouts"))
}

type subscriptionTemplateDeploymentState struct {
	DebugLevel            string                                        `json:"debug_level"`
	Id                    string                                        `json:"id"`
	Location              string                                        `json:"location"`
	Name                  string                                        `json:"name"`
	OutputContent         string                                        `json:"output_content"`
	ParametersContent     string                                        `json:"parameters_content"`
	Tags                  map[string]string                             `json:"tags"`
	TemplateContent       string                                        `json:"template_content"`
	TemplateSpecVersionId string                                        `json:"template_spec_version_id"`
	Timeouts              *subscriptiontemplatedeployment.TimeoutsState `json:"timeouts"`
}
