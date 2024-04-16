// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spring_cloud_container_deployment

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

// Resource represents the Terraform resource azurerm_spring_cloud_container_deployment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpringCloudContainerDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (asccd *Resource) Type() string {
	return "azurerm_spring_cloud_container_deployment"
}

// LocalName returns the local name for [Resource].
func (asccd *Resource) LocalName() string {
	return asccd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (asccd *Resource) Configuration() interface{} {
	return asccd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (asccd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(asccd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (asccd *Resource) Dependencies() terra.Dependencies {
	return asccd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (asccd *Resource) LifecycleManagement() *terra.Lifecycle {
	return asccd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (asccd *Resource) Attributes() azurermSpringCloudContainerDeploymentAttributes {
	return azurermSpringCloudContainerDeploymentAttributes{ref: terra.ReferenceResource(asccd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (asccd *Resource) ImportState(state io.Reader) error {
	asccd.state = &azurermSpringCloudContainerDeploymentState{}
	if err := json.NewDecoder(state).Decode(asccd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", asccd.Type(), asccd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (asccd *Resource) State() (*azurermSpringCloudContainerDeploymentState, bool) {
	return asccd.state, asccd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (asccd *Resource) StateMust() *azurermSpringCloudContainerDeploymentState {
	if asccd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", asccd.Type(), asccd.LocalName()))
	}
	return asccd.state
}

// Args contains the configurations for azurerm_spring_cloud_container_deployment.
type Args struct {
	// AddonJson: string, optional
	AddonJson terra.StringValue `hcl:"addon_json,attr"`
	// ApplicationPerformanceMonitoringIds: list of string, optional
	ApplicationPerformanceMonitoringIds terra.ListValue[terra.StringValue] `hcl:"application_performance_monitoring_ids,attr"`
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
	Quota *Quota `hcl:"quota,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpringCloudContainerDeploymentAttributes struct {
	ref terra.Reference
}

// AddonJson returns a reference to field addon_json of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) AddonJson() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("addon_json"))
}

// ApplicationPerformanceMonitoringIds returns a reference to field application_performance_monitoring_ids of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) ApplicationPerformanceMonitoringIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asccd.ref.Append("application_performance_monitoring_ids"))
}

// Arguments returns a reference to field arguments of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Arguments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asccd.ref.Append("arguments"))
}

// Commands returns a reference to field commands of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Commands() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](asccd.ref.Append("commands"))
}

// EnvironmentVariables returns a reference to field environment_variables of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](asccd.ref.Append("environment_variables"))
}

// Id returns a reference to field id of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("id"))
}

// Image returns a reference to field image of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Image() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("image"))
}

// InstanceCount returns a reference to field instance_count of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(asccd.ref.Append("instance_count"))
}

// LanguageFramework returns a reference to field language_framework of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) LanguageFramework() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("language_framework"))
}

// Name returns a reference to field name of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("name"))
}

// Server returns a reference to field server of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) Server() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("server"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_container_deployment.
func (asccd azurermSpringCloudContainerDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(asccd.ref.Append("spring_cloud_app_id"))
}

func (asccd azurermSpringCloudContainerDeploymentAttributes) Quota() terra.ListValue[QuotaAttributes] {
	return terra.ReferenceAsList[QuotaAttributes](asccd.ref.Append("quota"))
}

func (asccd azurermSpringCloudContainerDeploymentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](asccd.ref.Append("timeouts"))
}

type azurermSpringCloudContainerDeploymentState struct {
	AddonJson                           string            `json:"addon_json"`
	ApplicationPerformanceMonitoringIds []string          `json:"application_performance_monitoring_ids"`
	Arguments                           []string          `json:"arguments"`
	Commands                            []string          `json:"commands"`
	EnvironmentVariables                map[string]string `json:"environment_variables"`
	Id                                  string            `json:"id"`
	Image                               string            `json:"image"`
	InstanceCount                       float64           `json:"instance_count"`
	LanguageFramework                   string            `json:"language_framework"`
	Name                                string            `json:"name"`
	Server                              string            `json:"server"`
	SpringCloudAppId                    string            `json:"spring_cloud_app_id"`
	Quota                               []QuotaState      `json:"quota"`
	Timeouts                            *TimeoutsState    `json:"timeouts"`
}
