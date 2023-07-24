// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	codedeploydeploymentconfig "github.com/golingon/terraproviders/aws/5.9.0/codedeploydeploymentconfig"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewCodedeployDeploymentConfig creates a new instance of [CodedeployDeploymentConfig].
func NewCodedeployDeploymentConfig(name string, args CodedeployDeploymentConfigArgs) *CodedeployDeploymentConfig {
	return &CodedeployDeploymentConfig{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*CodedeployDeploymentConfig)(nil)

// CodedeployDeploymentConfig represents the Terraform resource aws_codedeploy_deployment_config.
type CodedeployDeploymentConfig struct {
	Name      string
	Args      CodedeployDeploymentConfigArgs
	state     *codedeployDeploymentConfigState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) Type() string {
	return "aws_codedeploy_deployment_config"
}

// LocalName returns the local name for [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) LocalName() string {
	return cdc.Name
}

// Configuration returns the configuration (args) for [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) Configuration() interface{} {
	return cdc.Args
}

// DependOn is used for other resources to depend on [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) DependOn() terra.Reference {
	return terra.ReferenceResource(cdc)
}

// Dependencies returns the list of resources [CodedeployDeploymentConfig] depends_on.
func (cdc *CodedeployDeploymentConfig) Dependencies() terra.Dependencies {
	return cdc.DependsOn
}

// LifecycleManagement returns the lifecycle block for [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) LifecycleManagement() *terra.Lifecycle {
	return cdc.Lifecycle
}

// Attributes returns the attributes for [CodedeployDeploymentConfig].
func (cdc *CodedeployDeploymentConfig) Attributes() codedeployDeploymentConfigAttributes {
	return codedeployDeploymentConfigAttributes{ref: terra.ReferenceResource(cdc)}
}

// ImportState imports the given attribute values into [CodedeployDeploymentConfig]'s state.
func (cdc *CodedeployDeploymentConfig) ImportState(av io.Reader) error {
	cdc.state = &codedeployDeploymentConfigState{}
	if err := json.NewDecoder(av).Decode(cdc.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", cdc.Type(), cdc.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [CodedeployDeploymentConfig] has state.
func (cdc *CodedeployDeploymentConfig) State() (*codedeployDeploymentConfigState, bool) {
	return cdc.state, cdc.state != nil
}

// StateMust returns the state for [CodedeployDeploymentConfig]. Panics if the state is nil.
func (cdc *CodedeployDeploymentConfig) StateMust() *codedeployDeploymentConfigState {
	if cdc.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", cdc.Type(), cdc.LocalName()))
	}
	return cdc.state
}

// CodedeployDeploymentConfigArgs contains the configurations for aws_codedeploy_deployment_config.
type CodedeployDeploymentConfigArgs struct {
	// ComputePlatform: string, optional
	ComputePlatform terra.StringValue `hcl:"compute_platform,attr"`
	// DeploymentConfigName: string, required
	DeploymentConfigName terra.StringValue `hcl:"deployment_config_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// MinimumHealthyHosts: optional
	MinimumHealthyHosts *codedeploydeploymentconfig.MinimumHealthyHosts `hcl:"minimum_healthy_hosts,block"`
	// TrafficRoutingConfig: optional
	TrafficRoutingConfig *codedeploydeploymentconfig.TrafficRoutingConfig `hcl:"traffic_routing_config,block"`
}
type codedeployDeploymentConfigAttributes struct {
	ref terra.Reference
}

// ComputePlatform returns a reference to field compute_platform of aws_codedeploy_deployment_config.
func (cdc codedeployDeploymentConfigAttributes) ComputePlatform() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("compute_platform"))
}

// DeploymentConfigId returns a reference to field deployment_config_id of aws_codedeploy_deployment_config.
func (cdc codedeployDeploymentConfigAttributes) DeploymentConfigId() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("deployment_config_id"))
}

// DeploymentConfigName returns a reference to field deployment_config_name of aws_codedeploy_deployment_config.
func (cdc codedeployDeploymentConfigAttributes) DeploymentConfigName() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("deployment_config_name"))
}

// Id returns a reference to field id of aws_codedeploy_deployment_config.
func (cdc codedeployDeploymentConfigAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(cdc.ref.Append("id"))
}

func (cdc codedeployDeploymentConfigAttributes) MinimumHealthyHosts() terra.ListValue[codedeploydeploymentconfig.MinimumHealthyHostsAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentconfig.MinimumHealthyHostsAttributes](cdc.ref.Append("minimum_healthy_hosts"))
}

func (cdc codedeployDeploymentConfigAttributes) TrafficRoutingConfig() terra.ListValue[codedeploydeploymentconfig.TrafficRoutingConfigAttributes] {
	return terra.ReferenceAsList[codedeploydeploymentconfig.TrafficRoutingConfigAttributes](cdc.ref.Append("traffic_routing_config"))
}

type codedeployDeploymentConfigState struct {
	ComputePlatform      string                                                 `json:"compute_platform"`
	DeploymentConfigId   string                                                 `json:"deployment_config_id"`
	DeploymentConfigName string                                                 `json:"deployment_config_name"`
	Id                   string                                                 `json:"id"`
	MinimumHealthyHosts  []codedeploydeploymentconfig.MinimumHealthyHostsState  `json:"minimum_healthy_hosts"`
	TrafficRoutingConfig []codedeploydeploymentconfig.TrafficRoutingConfigState `json:"traffic_routing_config"`
}
