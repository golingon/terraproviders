// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	"encoding/json"
	"fmt"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

// NewAppconfigDeployment creates a new instance of [AppconfigDeployment].
func NewAppconfigDeployment(name string, args AppconfigDeploymentArgs) *AppconfigDeployment {
	return &AppconfigDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*AppconfigDeployment)(nil)

// AppconfigDeployment represents the Terraform resource aws_appconfig_deployment.
type AppconfigDeployment struct {
	Name      string
	Args      AppconfigDeploymentArgs
	state     *appconfigDeploymentState
	DependsOn terra.Dependencies
	Lifecycle *terra.Lifecycle
}

// Type returns the Terraform object type for [AppconfigDeployment].
func (ad *AppconfigDeployment) Type() string {
	return "aws_appconfig_deployment"
}

// LocalName returns the local name for [AppconfigDeployment].
func (ad *AppconfigDeployment) LocalName() string {
	return ad.Name
}

// Configuration returns the configuration (args) for [AppconfigDeployment].
func (ad *AppconfigDeployment) Configuration() interface{} {
	return ad.Args
}

// DependOn is used for other resources to depend on [AppconfigDeployment].
func (ad *AppconfigDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(ad)
}

// Dependencies returns the list of resources [AppconfigDeployment] depends_on.
func (ad *AppconfigDeployment) Dependencies() terra.Dependencies {
	return ad.DependsOn
}

// LifecycleManagement returns the lifecycle block for [AppconfigDeployment].
func (ad *AppconfigDeployment) LifecycleManagement() *terra.Lifecycle {
	return ad.Lifecycle
}

// Attributes returns the attributes for [AppconfigDeployment].
func (ad *AppconfigDeployment) Attributes() appconfigDeploymentAttributes {
	return appconfigDeploymentAttributes{ref: terra.ReferenceResource(ad)}
}

// ImportState imports the given attribute values into [AppconfigDeployment]'s state.
func (ad *AppconfigDeployment) ImportState(av io.Reader) error {
	ad.state = &appconfigDeploymentState{}
	if err := json.NewDecoder(av).Decode(ad.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ad.Type(), ad.LocalName(), err)
	}
	return nil
}

// State returns the state and a bool indicating if [AppconfigDeployment] has state.
func (ad *AppconfigDeployment) State() (*appconfigDeploymentState, bool) {
	return ad.state, ad.state != nil
}

// StateMust returns the state for [AppconfigDeployment]. Panics if the state is nil.
func (ad *AppconfigDeployment) StateMust() *appconfigDeploymentState {
	if ad.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ad.Type(), ad.LocalName()))
	}
	return ad.state
}

// AppconfigDeploymentArgs contains the configurations for aws_appconfig_deployment.
type AppconfigDeploymentArgs struct {
	// ApplicationId: string, required
	ApplicationId terra.StringValue `hcl:"application_id,attr" validate:"required"`
	// ConfigurationProfileId: string, required
	ConfigurationProfileId terra.StringValue `hcl:"configuration_profile_id,attr" validate:"required"`
	// ConfigurationVersion: string, required
	ConfigurationVersion terra.StringValue `hcl:"configuration_version,attr" validate:"required"`
	// DeploymentStrategyId: string, required
	DeploymentStrategyId terra.StringValue `hcl:"deployment_strategy_id,attr" validate:"required"`
	// Description: string, optional
	Description terra.StringValue `hcl:"description,attr"`
	// EnvironmentId: string, required
	EnvironmentId terra.StringValue `hcl:"environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// TagsAll: map of string, optional
	TagsAll terra.MapValue[terra.StringValue] `hcl:"tags_all,attr"`
}
type appconfigDeploymentAttributes struct {
	ref terra.Reference
}

// ApplicationId returns a reference to field application_id of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) ApplicationId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("application_id"))
}

// Arn returns a reference to field arn of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) Arn() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("arn"))
}

// ConfigurationProfileId returns a reference to field configuration_profile_id of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) ConfigurationProfileId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("configuration_profile_id"))
}

// ConfigurationVersion returns a reference to field configuration_version of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) ConfigurationVersion() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("configuration_version"))
}

// DeploymentNumber returns a reference to field deployment_number of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) DeploymentNumber() terra.NumberValue {
	return terra.ReferenceAsNumber(ad.ref.Append("deployment_number"))
}

// DeploymentStrategyId returns a reference to field deployment_strategy_id of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) DeploymentStrategyId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("deployment_strategy_id"))
}

// Description returns a reference to field description of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("description"))
}

// EnvironmentId returns a reference to field environment_id of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) EnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("environment_id"))
}

// Id returns a reference to field id of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("id"))
}

// State returns a reference to field state of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) State() terra.StringValue {
	return terra.ReferenceAsString(ad.ref.Append("state"))
}

// Tags returns a reference to field tags of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ad.ref.Append("tags"))
}

// TagsAll returns a reference to field tags_all of aws_appconfig_deployment.
func (ad appconfigDeploymentAttributes) TagsAll() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ad.ref.Append("tags_all"))
}

type appconfigDeploymentState struct {
	ApplicationId          string            `json:"application_id"`
	Arn                    string            `json:"arn"`
	ConfigurationProfileId string            `json:"configuration_profile_id"`
	ConfigurationVersion   string            `json:"configuration_version"`
	DeploymentNumber       float64           `json:"deployment_number"`
	DeploymentStrategyId   string            `json:"deployment_strategy_id"`
	Description            string            `json:"description"`
	EnvironmentId          string            `json:"environment_id"`
	Id                     string            `json:"id"`
	State                  string            `json:"state"`
	Tags                   map[string]string `json:"tags"`
	TagsAll                map[string]string `json:"tags_all"`
}