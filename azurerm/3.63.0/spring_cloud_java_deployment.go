// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudjavadeployment "github.com/golingon/terraproviders/azurerm/3.63.0/springcloudjavadeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudJavaDeployment creates a new instance of [SpringCloudJavaDeployment].
func NewSpringCloudJavaDeployment(name string, args SpringCloudJavaDeploymentArgs) *SpringCloudJavaDeployment {
	return &SpringCloudJavaDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudJavaDeployment)(nil)

// SpringCloudJavaDeployment represents the Terraform resource azurerm_spring_cloud_java_deployment.
type SpringCloudJavaDeployment struct {
	Name      string
	Args      SpringCloudJavaDeploymentArgs
	state     *springCloudJavaDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) Type() string {
	return "azurerm_spring_cloud_java_deployment"
}

// LocalName returns the local name for [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) LocalName() string {
	return scjd.Name
}

// Configuration returns the configuration (args) for [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) Configuration() interface{} {
	return scjd.Args
}

// DependOn is used for other resources to depend on [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(scjd)
}

// Dependencies returns the list of resources [SpringCloudJavaDeployment] depends_on.
func (scjd *SpringCloudJavaDeployment) Dependencies() terra.Dependencies {
	return scjd.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) LifecycleManagement() *terra.Lifecycle {
	return scjd.Lifecycle
}

// Attributes returns the attributes for [SpringCloudJavaDeployment].
func (scjd *SpringCloudJavaDeployment) Attributes() springCloudJavaDeploymentAttributes {
	return springCloudJavaDeploymentAttributes{ref: terra.ReferenceResource(scjd)}
}

// ImportState imports the given attribute values into [SpringCloudJavaDeployment]'s state.
func (scjd *SpringCloudJavaDeployment) ImportState(av io.Reader) error {
	scjd.state = &springCloudJavaDeploymentState{}
	if err := json.NewDecoder(av).Decode(scjd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scjd.Type(), scjd.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudJavaDeployment] has state.
func (scjd *SpringCloudJavaDeployment) State() (*springCloudJavaDeploymentState, bool) {
	return scjd.state, scjd.state != nil
}

// StateMust returns the state for [SpringCloudJavaDeployment]. Panics if the state is nil.
func (scjd *SpringCloudJavaDeployment) StateMust() *springCloudJavaDeploymentState {
	if scjd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scjd.Type(), scjd.LocalName()))
	}
	return scjd.state
}

// SpringCloudJavaDeploymentArgs contains the configurations for azurerm_spring_cloud_java_deployment.
type SpringCloudJavaDeploymentArgs struct {
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
	Quota *springcloudjavadeployment.Quota `hcl:"quota,block"`
	// Timeouts: optional
	Timeouts *springcloudjavadeployment.Timeouts `hcl:"timeouts,block"`
}
type springCloudJavaDeploymentAttributes struct {
	ref terra.Reference
}

// EnvironmentVariables returns a reference to field environment_variables of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scjd.ref.Append("environment_variables"))
}

// Id returns a reference to field id of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scjd.ref.Append("id"))
}

// InstanceCount returns a reference to field instance_count of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceAsNumber(scjd.ref.Append("instance_count"))
}

// JvmOptions returns a reference to field jvm_options of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) JvmOptions() terra.StringValue {
	return terra.ReferenceAsString(scjd.ref.Append("jvm_options"))
}

// Name returns a reference to field name of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scjd.ref.Append("name"))
}

// RuntimeVersion returns a reference to field runtime_version of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) RuntimeVersion() terra.StringValue {
	return terra.ReferenceAsString(scjd.ref.Append("runtime_version"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_java_deployment.
func (scjd springCloudJavaDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(scjd.ref.Append("spring_cloud_app_id"))
}

func (scjd springCloudJavaDeploymentAttributes) Quota() terra.ListValue[springcloudjavadeployment.QuotaAttributes] {
	return terra.ReferenceAsList[springcloudjavadeployment.QuotaAttributes](scjd.ref.Append("quota"))
}

func (scjd springCloudJavaDeploymentAttributes) Timeouts() springcloudjavadeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudjavadeployment.TimeoutsAttributes](scjd.ref.Append("timeouts"))
}

type springCloudJavaDeploymentState struct {
	EnvironmentVariables map[string]string                        `json:"environment_variables"`
	Id                   string                                   `json:"id"`
	InstanceCount        float64                                  `json:"instance_count"`
	JvmOptions           string                                   `json:"jvm_options"`
	Name                 string                                   `json:"name"`
	RuntimeVersion       string                                   `json:"runtime_version"`
	SpringCloudAppId     string                                   `json:"spring_cloud_app_id"`
	Quota                []springcloudjavadeployment.QuotaState   `json:"quota"`
	Timeouts             *springcloudjavadeployment.TimeoutsState `json:"timeouts"`
}
