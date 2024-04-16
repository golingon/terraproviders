// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_spring_cloud_java_deployment

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

// Resource represents the Terraform resource azurerm_spring_cloud_java_deployment.
type Resource struct {
	Name      string
	Args      Args
	state     *azurermSpringCloudJavaDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [Resource].
func (ascjd *Resource) Type() string {
	return "azurerm_spring_cloud_java_deployment"
}

// LocalName returns the local name for [Resource].
func (ascjd *Resource) LocalName() string {
	return ascjd.Name
}

// Configuration returns the configuration (args) for [Resource].
func (ascjd *Resource) Configuration() interface{} {
	return ascjd.Args
}

// DependOn is used for other resources to depend on [Resource].
func (ascjd *Resource) DependOn() terra.Reference {
	return terra.ReferenceResource(ascjd)
}

// Dependencies returns the list of resources [Resource] depends_on.
func (ascjd *Resource) Dependencies() terra.Dependencies {
	return ascjd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [Resource].
func (ascjd *Resource) LifecycleManagement() *terra.Lifecycle {
	return ascjd.Lifecycle
}

// Attributes returns the attributes for [Resource].
func (ascjd *Resource) Attributes() azurermSpringCloudJavaDeploymentAttributes {
	return azurermSpringCloudJavaDeploymentAttributes{ref: terra.ReferenceResource(ascjd)}
}

// ImportState imports the given attribute values into [Resource]'s state.
func (ascjd *Resource) ImportState(state io.Reader) error {
	ascjd.state = &azurermSpringCloudJavaDeploymentState{}
	if err := json.NewDecoder(state).Decode(ascjd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ascjd.Type(), ascjd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [Resource] has state.
func (ascjd *Resource) State() (*azurermSpringCloudJavaDeploymentState, bool) {
	return ascjd.state, ascjd.state != nil
}

// StateMust returns the state for [Resource]. Panics if the state is nil.
func (ascjd *Resource) StateMust() *azurermSpringCloudJavaDeploymentState {
	if ascjd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ascjd.Type(), ascjd.LocalName()))
	}
	return ascjd.state
}

// Args contains the configurations for azurerm_spring_cloud_java_deployment.
type Args struct {
	// EnvironmentVariables: map of string, optional
	EnvironmentVariables terra.MapValue[terra.StringValue] `hcl:"environment_variables,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// InstanceCount: number, optional
	InstanceCount terra.NumberValue `hcl:"instance_count,attr"`
	// JvmOptions: string, optional
	JvmOptions terra.StringValue `hcl:"jvm_options,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RuntimeVersion: string, optional
	RuntimeVersion terra.StringValue `hcl:"runtime_version,attr"`
	// SpringCloudAppId: string, required
	SpringCloudAppId terra.StringValue `hcl:"spring_cloud_app_id,attr" validate:"required"`
	// Quota: optional
	Quota *Quota `hcl:"quota,block"`
	// Timeouts: optional
	Timeouts *Timeouts `hcl:"timeouts,block"`
}

type azurermSpringCloudJavaDeploymentAttributes struct {
	ref terra.Reference
}

// EnvironmentVariables returns a reference to field environment_variables of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ascjd.ref.Append("environment_variables"))
}

// Id returns a reference to field id of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ascjd.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(ascjd.ref.Append("instance_count"))
}

// JvmOptions returns a reference to field jvm_options of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) JvmOptions() terra.StringValue {
	return terra.ReferenceAsString(ascjd.ref.Append("jvm_options"))
}

// Name returns a reference to field name of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ascjd.ref.Append("name"))
}

// RuntimeVersion returns a reference to field runtime_version of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) RuntimeVersion() terra.StringValue {
	return terra.ReferenceAsString(ascjd.ref.Append("runtime_version"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_java_deployment.
func (ascjd azurermSpringCloudJavaDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(ascjd.ref.Append("spring_cloud_app_id"))
}

func (ascjd azurermSpringCloudJavaDeploymentAttributes) Quota() terra.ListValue[QuotaAttributes] {
	return terra.ReferenceAsList[QuotaAttributes](ascjd.ref.Append("quota"))
}

func (ascjd azurermSpringCloudJavaDeploymentAttributes) Timeouts() TimeoutsAttributes {
	return terra.ReferenceAsSingle[TimeoutsAttributes](ascjd.ref.Append("timeouts"))
}

type azurermSpringCloudJavaDeploymentState struct {
	EnvironmentVariables map[string]string `json:"environment_variables"`
	Id                   string            `json:"id"`
	InstanceCount        float64           `json:"instance_count"`
	JvmOptions           string            `json:"jvm_options"`
	Name                 string            `json:"name"`
	RuntimeVersion       string            `json:"runtime_version"`
	SpringCloudAppId     string            `json:"spring_cloud_app_id"`
	Quota                []QuotaState      `json:"quota"`
	Timeouts             *TimeoutsState    `json:"timeouts"`
}
