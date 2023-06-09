// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudcontainerdeployment "github.com/golingon/terraproviders/azurerm/3.52.0/springcloudcontainerdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudContainerDeployment creates a new instance of [SpringCloudContainerDeployment].
func NewSpringCloudContainerDeployment(name string, args SpringCloudContainerDeploymentArgs) *SpringCloudContainerDeployment {
	return &SpringCloudContainerDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudContainerDeployment)(nil)

// SpringCloudContainerDeployment represents the Terraform resource azurerm_spring_cloud_container_deployment.
type SpringCloudContainerDeployment struct {
	Name      string
	Args      SpringCloudContainerDeploymentArgs
	state     *springCloudContainerDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) Type() string {
	return "azurerm_spring_cloud_container_deployment"
}

// LocalName returns the local name for [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) LocalName() string {
	return sccd.Name
}

// Configuration returns the configuration (args) for [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) Configuration() interface{} {
	return sccd.Args
}

// DependOn is used for other resources to depend on [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(sccd)
}

// Dependencies returns the list of resources [SpringCloudContainerDeployment] depends_on.
func (sccd *SpringCloudContainerDeployment) Dependencies() terra.Dependencies {
	return sccd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) LifecycleManagement() *terra.Lifecycle {
	return sccd.Lifecycle
}

// Attributes returns the attributes for [SpringCloudContainerDeployment].
func (sccd *SpringCloudContainerDeployment) Attributes() springCloudContainerDeploymentAttributes {
	return springCloudContainerDeploymentAttributes{ref: terra.ReferenceResource(sccd)}
}

// ImportState imports the given attribute values into [SpringCloudContainerDeployment]'s state.
func (sccd *SpringCloudContainerDeployment) ImportState(av io.Reader) error {
	sccd.state = &springCloudContainerDeploymentState{}
	if err := json.NewDecoder(av).Decode(sccd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sccd.Type(), sccd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudContainerDeployment] has state.
func (sccd *SpringCloudContainerDeployment) State() (*springCloudContainerDeploymentState, bool) {
	return sccd.state, sccd.state != nil
}

// StateMust returns the state for [SpringCloudContainerDeployment]. Panics if the state is nil.
func (sccd *SpringCloudContainerDeployment) StateMust() *springCloudContainerDeploymentState {
	if sccd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sccd.Type(), sccd.LocalName()))
	}
	return sccd.state
}

// SpringCloudContainerDeploymentArgs contains the configurations for azurerm_spring_cloud_container_deployment.
type SpringCloudContainerDeploymentArgs struct {
	// AddonJson: string, optional
	AddonJson terra.StringValue `hcl:"addon_json,attr"`
	// Arguments: list of string, optional
	Arguments terra.ListValue[terra.StringValue] `hcl:"arguments,attr"`
	// Commands: list of string, optional
	Commands terra.ListValue[terra.StringValue] `hcl:"commands,attr"`
	// EnvironmentVariables: map of string, optional
	EnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"environment_variables,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Image: string, required
	Image terra.StringValue `hcl:"image,attr" validate:"required"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// LanguageFramework: string, optional
	LanguageFramework terra.StringValue `hcl:"language_framework,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// Server: string, required
	Server terra.StringValue `hcl:"server,attr" validate:"required"`
	// SpringCloudAppId: string, required
	SpringCloudAppId terra.StringValue `hcl:"spring_cloud_app_id,attr" validate:"required"`
	// Quota: optional
	Quota *springcloudcontainerdeployment.Quota `hcl:"quota,block"`
	// Timeouts: optional
	Timeouts *springcloudcontainerdeployment.Timeouts `hcl:"timeouts,block"`
}
type springCloudContainerDeploymentAttributes struct {
	ref terra.Reference
}

// AddonJson returns a reference to field addon_json of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) AddonJson() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("addon_json"))
}

// Arguments returns a reference to field arguments of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Arguments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sccd.ref.Append("arguments"))
}

// Commands returns a reference to field commands of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Commands() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](sccd.ref.Append("commands"))
}

// EnvironmentVariables returns a reference to field environment_variables of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](sccd.ref.Append("environment_variables"))
}

// Id returns a reference to field id of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("id"))
}

// Image returns a reference to field image of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("image"))
}

// InstanceCount returns a reference to field instance_count of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(sccd.ref.Append("instance_count"))
}

// LanguageFramework returns a reference to field language_framework of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) LanguageFramework() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("language_framework"))
}

// Name returns a reference to field name of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("name"))
}

// Server returns a reference to field server of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("server"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_container_deployment.
func (sccd springCloudContainerDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(sccd.ref.Append("spring_cloud_app_id"))
}

func (sccd springCloudContainerDeploymentAttributes) Quota() terra.ListValue[springcloudcontainerdeployment.QuotaAttributes] {
	return terra.ReferenceAsList[springcloudcontainerdeployment.QuotaAttributes](sccd.ref.Append("quota"))
}

func (sccd springCloudContainerDeploymentAttributes) Timeouts() springcloudcontainerdeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudcontainerdeployment.TimeoutsAttributes](sccd.ref.Append("timeouts"))
}

type springCloudContainerDeploymentState struct {
	AddonJson            string                                        `json:"addon_json"`
	Arguments            []string                                      `json:"arguments"`
	Commands             []string                                      `json:"commands"`
	EnvironmentVariables map[string]string                             `json:"environment_variables"`
	Id                   string                                        `json:"id"`
	Image                string                                        `json:"image"`
	InstanceCount        float64                                       `json:"instance_count"`
	LanguageFramework    string                                        `json:"language_framework"`
	Name                 string                                        `json:"name"`
	Server               string                                        `json:"server"`
	SpringCloudAppId     string                                        `json:"spring_cloud_app_id"`
	Quota                []springcloudcontainerdeployment.QuotaState   `json:"quota"`
	Timeouts             *springcloudcontainerdeployment.TimeoutsState `json:"timeouts"`
}
