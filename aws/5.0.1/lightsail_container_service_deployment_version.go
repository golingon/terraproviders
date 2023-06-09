// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	lightsailcontainerservicedeploymentversion "github.com/golingon/terraproviders/aws/5.0.1/lightsailcontainerservicedeploymentversion"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewLightsailContainerServiceDeploymentVersion creates a new instance of [LightsailContainerServiceDeploymentVersion].
func NewLightsailContainerServiceDeploymentVersion(name string, args LightsailContainerServiceDeploymentVersionArgs) *LightsailContainerServiceDeploymentVersion {
	return &LightsailContainerServiceDeploymentVersion{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*LightsailContainerServiceDeploymentVersion)(nil)

// LightsailContainerServiceDeploymentVersion represents the Terraform resource aws_lightsail_container_service_deployment_version.
type LightsailContainerServiceDeploymentVersion struct {
	Name      string
	Args      LightsailContainerServiceDeploymentVersionArgs
	state     *lightsailContainerServiceDeploymentVersionState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) Type() string {
	return "aws_lightsail_container_service_deployment_version"
}

// LocalName returns the local name for [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) LocalName() string {
	return lcsdv.Name
}

// Configuration returns the configuration (args) for [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) Configuration() interface{} {
	return lcsdv.Args
}

// DependOn is used for other resources to depend on [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) DependOn() terra.Reference {
	return terra.ReferenceResource(lcsdv)
}

// Dependencies returns the list of resources [LightsailContainerServiceDeploymentVersion] depends_on.
func (lcsdv *LightsailContainerServiceDeploymentVersion) Dependencies() terra.Dependencies {
	return lcsdv.DependsOn
}

// LifecycleManagement returns the lifecycle block for [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) LifecycleManagement() *terra.Lifecycle {
	return lcsdv.Lifecycle
}

// Attributes returns the attributes for [LightsailContainerServiceDeploymentVersion].
func (lcsdv *LightsailContainerServiceDeploymentVersion) Attributes() lightsailContainerServiceDeploymentVersionAttributes {
	return lightsailContainerServiceDeploymentVersionAttributes{ref: terra.ReferenceResource(lcsdv)}
}

// ImportState imports the given attribute values into [LightsailContainerServiceDeploymentVersion]'s state.
func (lcsdv *LightsailContainerServiceDeploymentVersion) ImportState(av io.Reader) error {
	lcsdv.state = &lightsailContainerServiceDeploymentVersionState{}
	if err := json.NewDecoder(av).Decode(lcsdv.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", lcsdv.Type(), lcsdv.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [LightsailContainerServiceDeploymentVersion] has state.
func (lcsdv *LightsailContainerServiceDeploymentVersion) State() (*lightsailContainerServiceDeploymentVersionState, bool) {
	return lcsdv.state, lcsdv.state != nil
}

// StateMust returns the state for [LightsailContainerServiceDeploymentVersion]. Panics if the state is nil.
func (lcsdv *LightsailContainerServiceDeploymentVersion) StateMust() *lightsailContainerServiceDeploymentVersionState {
	if lcsdv.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", lcsdv.Type(), lcsdv.LocalName()))
	}
	return lcsdv.state
}

// LightsailContainerServiceDeploymentVersionArgs contains the configurations for aws_lightsail_container_service_deployment_version.
type LightsailContainerServiceDeploymentVersionArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// ServiceName: string, required
	ServiceName terra.StringValue `hcl:"service_name,attr" validate:"required"`
	// Container: min=1,max=53
	Container []lightsailcontainerservicedeploymentversion.Container `hcl:"container,block" validate:"min=1,max=53"`
	// PublicEndpoint: optional
	PublicEndpoint *lightsailcontainerservicedeploymentversion.PublicEndpoint `hcl:"public_endpoint,block"`
	// Timeouts: optional
	Timeouts *lightsailcontainerservicedeploymentversion.Timeouts `hcl:"timeouts,block"`
}
type lightsailContainerServiceDeploymentVersionAttributes struct {
	ref terra.Reference
}

// CreatedAt returns a reference to field created_at of aws_lightsail_container_service_deployment_version.
func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) CreatedAt() terra.StringValue {
	return terra.ReferenceAsString(lcsdv.ref.Append("created_at"))
}

// Id returns a reference to field id of aws_lightsail_container_service_deployment_version.
func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lcsdv.ref.Append("id"))
}

// ServiceName returns a reference to field service_name of aws_lightsail_container_service_deployment_version.
func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) ServiceName() terra.StringValue {
	return terra.ReferenceAsString(lcsdv.ref.Append("service_name"))
}

// State returns a reference to field state of aws_lightsail_container_service_deployment_version.
func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(lcsdv.ref.Append("state"))
}

// Version returns a reference to field version of aws_lightsail_container_service_deployment_version.
func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) Version() terra.NumberValue {
	return terra.ReferenceAsNumber(lcsdv.ref.Append("version"))
}

func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) Container() terra.SetValue[lightsailcontainerservicedeploymentversion.ContainerAttributes] {
	return terra.ReferenceAsSet[lightsailcontainerservicedeploymentversion.ContainerAttributes](lcsdv.ref.Append("container"))
}

func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) PublicEndpoint() terra.ListValue[lightsailcontainerservicedeploymentversion.PublicEndpointAttributes] {
	return terra.ReferenceAsList[lightsailcontainerservicedeploymentversion.PublicEndpointAttributes](lcsdv.ref.Append("public_endpoint"))
}

func (lcsdv lightsailContainerServiceDeploymentVersionAttributes) Timeouts() lightsailcontainerservicedeploymentversion.TimeoutsAttributes {
	return terra.ReferenceAsSingle[lightsailcontainerservicedeploymentversion.TimeoutsAttributes](lcsdv.ref.Append("timeouts"))
}

type lightsailContainerServiceDeploymentVersionState struct {
	CreatedAt      string                                                           `json:"created_at"`
	Id             string                                                           `json:"id"`
	ServiceName    string                                                           `json:"service_name"`
	State          string                                                           `json:"state"`
	Version        float64                                                          `json:"version"`
	Container      []lightsailcontainerservicedeploymentversion.ContainerState      `json:"container"`
	PublicEndpoint []lightsailcontainerservicedeploymentversion.PublicEndpointState `json:"public_endpoint"`
	Timeouts       *lightsailcontainerservicedeploymentversion.TimeoutsState        `json:"timeouts"`
}
