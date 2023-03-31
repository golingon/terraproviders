// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	springcloudcontainerdeployment "github.com/golingon/terraproviders/azurerm/3.49.0/springcloudcontainerdeployment"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewSpringCloudContainerDeployment(name string, args SpringCloudContainerDeploymentArgs) *SpringCloudContainerDeployment {
	return &SpringCloudContainerDeployment{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*SpringCloudContainerDeployment)(nil)

type SpringCloudContainerDeployment struct {
	Name  string
	Args  SpringCloudContainerDeploymentArgs
	state *springCloudContainerDeploymentState
}

func (sccd *SpringCloudContainerDeployment) Type() string {
	return "azurerm_spring_cloud_container_deployment"
}

func (sccd *SpringCloudContainerDeployment) LocalName() string {
	return sccd.Name
}

func (sccd *SpringCloudContainerDeployment) Configuration() interface{} {
	return sccd.Args
}

func (sccd *SpringCloudContainerDeployment) Attributes() springCloudContainerDeploymentAttributes {
	return springCloudContainerDeploymentAttributes{ref: terra.ReferenceResource(sccd)}
}

func (sccd *SpringCloudContainerDeployment) ImportState(av io.Reader) error {
	sccd.state = &springCloudContainerDeploymentState{}
	if err := json.NewDecoder(av).Decode(sccd.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", sccd.Type(), sccd.LocalName(), err)
	}
	return nil
}

func (sccd *SpringCloudContainerDeployment) State() (*springCloudContainerDeploymentState, bool) {
	return sccd.state, sccd.state != nil
}

func (sccd *SpringCloudContainerDeployment) StateMust() *springCloudContainerDeploymentState {
	if sccd.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", sccd.Type(), sccd.LocalName()))
	}
	return sccd.state
}

func (sccd *SpringCloudContainerDeployment) DependOn() terra.Reference {
	return terra.ReferenceResource(sccd)
}

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
	// DependsOn contains resources that SpringCloudContainerDeployment depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type springCloudContainerDeploymentAttributes struct {
	ref terra.Reference
}

func (sccd springCloudContainerDeploymentAttributes) AddonJson() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("addon_json"))
}

func (sccd springCloudContainerDeploymentAttributes) Arguments() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sccd.ref.Append("arguments"))
}

func (sccd springCloudContainerDeploymentAttributes) Commands() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](sccd.ref.Append("commands"))
}

func (sccd springCloudContainerDeploymentAttributes) EnvironmentVariables() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](sccd.ref.Append("environment_variables"))
}

func (sccd springCloudContainerDeploymentAttributes) Id() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("id"))
}

func (sccd springCloudContainerDeploymentAttributes) Image() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("image"))
}

func (sccd springCloudContainerDeploymentAttributes) InstanceCount() terra.NumberValue {
	return terra.ReferenceNumber(sccd.ref.Append("instance_count"))
}

func (sccd springCloudContainerDeploymentAttributes) LanguageFramework() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("language_framework"))
}

func (sccd springCloudContainerDeploymentAttributes) Name() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("name"))
}

func (sccd springCloudContainerDeploymentAttributes) Server() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("server"))
}

func (sccd springCloudContainerDeploymentAttributes) SpringCloudAppId() terra.StringValue {
	return terra.ReferenceString(sccd.ref.Append("spring_cloud_app_id"))
}

func (sccd springCloudContainerDeploymentAttributes) Quota() terra.ListValue[springcloudcontainerdeployment.QuotaAttributes] {
	return terra.ReferenceList[springcloudcontainerdeployment.QuotaAttributes](sccd.ref.Append("quota"))
}

func (sccd springCloudContainerDeploymentAttributes) Timeouts() springcloudcontainerdeployment.TimeoutsAttributes {
	return terra.ReferenceSingle[springcloudcontainerdeployment.TimeoutsAttributes](sccd.ref.Append("timeouts"))
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
