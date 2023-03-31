// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	"encoding/json"
	"fmt"
	containerapp "github.com/golingon/terraproviders/azurerm/3.49.0/containerapp"
	"github.com/volvo-cars/lingon/pkg/terra"
	"io"
)

func NewContainerApp(name string, args ContainerAppArgs) *ContainerApp {
	return &ContainerApp{
		Args: args,
		Name: name,
	}
}

var _ terra.Resource = (*ContainerApp)(nil)

type ContainerApp struct {
	Name  string
	Args  ContainerAppArgs
	state *containerAppState
}

func (ca *ContainerApp) Type() string {
	return "azurerm_container_app"
}

func (ca *ContainerApp) LocalName() string {
	return ca.Name
}

func (ca *ContainerApp) Configuration() interface{} {
	return ca.Args
}

func (ca *ContainerApp) Attributes() containerAppAttributes {
	return containerAppAttributes{ref: terra.ReferenceResource(ca)}
}

func (ca *ContainerApp) ImportState(av io.Reader) error {
	ca.state = &containerAppState{}
	if err := json.NewDecoder(av).Decode(ca.state); err != nil {
		return fmt.Errorf("decoding state into resource %s.%s: %w", ca.Type(), ca.LocalName(), err)
	}
	return nil
}

func (ca *ContainerApp) State() (*containerAppState, bool) {
	return ca.state, ca.state != nil
}

func (ca *ContainerApp) StateMust() *containerAppState {
	if ca.state == nil {
		panic(fmt.Sprintf("state is nil for resource %s.%s", ca.Type(), ca.LocalName()))
	}
	return ca.state
}

func (ca *ContainerApp) DependOn() terra.Reference {
	return terra.ReferenceResource(ca)
}

type ContainerAppArgs struct {
	// ContainerAppEnvironmentId: string, required
	ContainerAppEnvironmentId terra.StringValue `hcl:"container_app_environment_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// RevisionMode: string, required
	RevisionMode terra.StringValue `hcl:"revision_mode,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Dapr: optional
	Dapr *containerapp.Dapr `hcl:"dapr,block"`
	// Identity: optional
	Identity *containerapp.Identity `hcl:"identity,block"`
	// Ingress: optional
	Ingress *containerapp.Ingress `hcl:"ingress,block"`
	// Registry: min=0
	Registry []containerapp.Registry `hcl:"registry,block" validate:"min=0"`
	// Secret: min=0
	Secret []containerapp.Secret `hcl:"secret,block" validate:"min=0"`
	// Template: required
	Template *containerapp.Template `hcl:"template,block" validate:"required"`
	// Timeouts: optional
	Timeouts *containerapp.Timeouts `hcl:"timeouts,block"`
	// DependsOn contains resources that ContainerApp depends on
	DependsOn terra.Dependencies `hcl:"depends_on,attr"`
}
type containerAppAttributes struct {
	ref terra.Reference
}

func (ca containerAppAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("container_app_environment_id"))
}

func (ca containerAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("custom_domain_verification_id"))
}

func (ca containerAppAttributes) Id() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("id"))
}

func (ca containerAppAttributes) LatestRevisionFqdn() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("latest_revision_fqdn"))
}

func (ca containerAppAttributes) LatestRevisionName() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("latest_revision_name"))
}

func (ca containerAppAttributes) Location() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("location"))
}

func (ca containerAppAttributes) Name() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("name"))
}

func (ca containerAppAttributes) OutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](ca.ref.Append("outbound_ip_addresses"))
}

func (ca containerAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("resource_group_name"))
}

func (ca containerAppAttributes) RevisionMode() terra.StringValue {
	return terra.ReferenceString(ca.ref.Append("revision_mode"))
}

func (ca containerAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](ca.ref.Append("tags"))
}

func (ca containerAppAttributes) Dapr() terra.ListValue[containerapp.DaprAttributes] {
	return terra.ReferenceList[containerapp.DaprAttributes](ca.ref.Append("dapr"))
}

func (ca containerAppAttributes) Identity() terra.ListValue[containerapp.IdentityAttributes] {
	return terra.ReferenceList[containerapp.IdentityAttributes](ca.ref.Append("identity"))
}

func (ca containerAppAttributes) Ingress() terra.ListValue[containerapp.IngressAttributes] {
	return terra.ReferenceList[containerapp.IngressAttributes](ca.ref.Append("ingress"))
}

func (ca containerAppAttributes) Registry() terra.ListValue[containerapp.RegistryAttributes] {
	return terra.ReferenceList[containerapp.RegistryAttributes](ca.ref.Append("registry"))
}

func (ca containerAppAttributes) Secret() terra.SetValue[containerapp.SecretAttributes] {
	return terra.ReferenceSet[containerapp.SecretAttributes](ca.ref.Append("secret"))
}

func (ca containerAppAttributes) Template() terra.ListValue[containerapp.TemplateAttributes] {
	return terra.ReferenceList[containerapp.TemplateAttributes](ca.ref.Append("template"))
}

func (ca containerAppAttributes) Timeouts() containerapp.TimeoutsAttributes {
	return terra.ReferenceSingle[containerapp.TimeoutsAttributes](ca.ref.Append("timeouts"))
}

type containerAppState struct {
	ContainerAppEnvironmentId  string                       `json:"container_app_environment_id"`
	CustomDomainVerificationId string                       `json:"custom_domain_verification_id"`
	Id                         string                       `json:"id"`
	LatestRevisionFqdn         string                       `json:"latest_revision_fqdn"`
	LatestRevisionName         string                       `json:"latest_revision_name"`
	Location                   string                       `json:"location"`
	Name                       string                       `json:"name"`
	OutboundIpAddresses        []string                     `json:"outbound_ip_addresses"`
	ResourceGroupName          string                       `json:"resource_group_name"`
	RevisionMode               string                       `json:"revision_mode"`
	Tags                       map[string]string            `json:"tags"`
	Dapr                       []containerapp.DaprState     `json:"dapr"`
	Identity                   []containerapp.IdentityState `json:"identity"`
	Ingress                    []containerapp.IngressState  `json:"ingress"`
	Registry                   []containerapp.RegistryState `json:"registry"`
	Secret                     []containerapp.SecretState   `json:"secret"`
	Template                   []containerapp.TemplateState `json:"template"`
	Timeouts                   *containerapp.TimeoutsState  `json:"timeouts"`
}
