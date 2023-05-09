// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudactivedeployment "github.com/golingon/terraproviders/azurerm/3.55.0/springcloudactivedeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewSpringCloudActiveDeployment creates a new instance of [SpringCloudActiveDeployment].
func NewSpringCloudActiveDeployment(name string, args SpringCloudActiveDeploymentArgs) *SpringCloudActiveDeployment {
	return &SpringCloudActiveDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudActiveDeployment)(nil)

// SpringCloudActiveDeployment represents the Terraform resource azurerm_spring_cloud_active_deployment.
type SpringCloudActiveDeployment struct {
	Name      string
	Args      SpringCloudActiveDeploymentArgs
	state     *springCloudActiveDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) Type() string {
	return "azurerm_spring_cloud_active_deployment"
}

// LocalName returns the local name for [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) LocalName() string {
	return scad.Name
}

// Configuration returns the configuration (args) for [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) Configuration() interface{} {
	return scad.Args
}

// DependOn is used for other resources to depend on [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(scad)
}

// Dependencies returns the list of resources [SpringCloudActiveDeployment] depends_on.
func (scad *SpringCloudActiveDeployment) Dependencies() terra.Dependencies {
	return scad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) LifecycleManagement() *terra.Lifecycle {
	return scad.Lifecycle
}

// Attributes returns the attributes for [SpringCloudActiveDeployment].
func (scad *SpringCloudActiveDeployment) Attributes() springCloudActiveDeploymentAttributes {
	return springCloudActiveDeploymentAttributes{ref: terra.ReferenceResource(scad)}
}

// ImportState imports the given attribute values into [SpringCloudActiveDeployment]'s state.
func (scad *SpringCloudActiveDeployment) ImportState(av io.Reader) error {
	scad.state = &springCloudActiveDeploymentState{}
	if err := json.NewDecoder(av).Decode(scad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", scad.Type(), scad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [SpringCloudActiveDeployment] has state.
func (scad *SpringCloudActiveDeployment) State() (*springCloudActiveDeploymentState, bool) {
	return scad.state, scad.state != nil
}

// StateMust returns the state for [SpringCloudActiveDeployment]. Panics if the state is nil.
func (scad *SpringCloudActiveDeployment) StateMust() *springCloudActiveDeploymentState {
	if scad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", scad.Type(), scad.LocalName()))
	}
	return scad.state
}

// SpringCloudActiveDeploymentArgs contains the configurations for azurerm_spring_cloud_active_deployment.
type SpringCloudActiveDeploymentArgs struct {
	// DeploymentName: string, required
	DeploymentName terra.StringValue `hcl:"deployment_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// SpringCloudAppId: string, required
	SpringCloudAppId terra.StringValue `hcl:"spring_cloud_app_id,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *springcloudactivedeployment.Timeouts `hcl:"timeouts,block"`
}
type springCloudActiveDeploymentAttributes struct {
	ref terra.Reference
}

// DeploymentName returns a reference to field deployment_name of azurerm_spring_cloud_active_deployment.
func (scad springCloudActiveDeploymentAttributes) DeploymentName() terra.StringValue {
	return terra.ReferenceAsString(scad.ref.Append("deployment_name"))
}

// Id returns a reference to field id of azurerm_spring_cloud_active_deployment.
func (scad springCloudActiveDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scad.ref.Append("id"))
}

// SpringCloudAppId returns a reference to field spring_cloud_app_id of azurerm_spring_cloud_active_deployment.
func (scad springCloudActiveDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceAsString(scad.ref.Append("spring_cloud_app_id"))
}

func (scad springCloudActiveDeploymentAttributes) Timeouts() springcloudactivedeployment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[springcloudactivedeployment.TimeoutsAttributes](scad.ref.Append("timeouts"))
}

type springCloudActiveDeploymentState struct {
	DeploymentName   string                                     `json:"deployment_name"`
	Id               string                                     `json:"id"`
	SpringCloudAppId string                                     `json:"spring_cloud_app_id"`
	Timeouts         *springcloudactivedeployment.TimeoutsState `json:"timeouts"`
}
