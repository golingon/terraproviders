// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datacontainerapp "github.com/golingon/terraproviders/azurerm/3.69.0/datacontainerapp"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataContainerApp creates a new instance of [DataContainerApp].
func NewDataContainerApp(name string, args DataContainerAppArgs) *DataContainerApp {
	return &DataContainerApp{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataContainerApp)(nil)

// DataContainerApp represents the Terraform data resource azurerm_container_app.
type DataContainerApp struct {
	Name string
	Args DataContainerAppArgs
}

// DataSource returns the Terraform object type for [DataContainerApp].
func (ca *DataContainerApp) DataSource() string {
	return "azurerm_container_app"
}

// LocalName returns the local name for [DataContainerApp].
func (ca *DataContainerApp) LocalName() string {
	return ca.Name
}

// Configuration returns the configuration (args) for [DataContainerApp].
func (ca *DataContainerApp) Configuration() interface{} {
	return ca.Args
}

// Attributes returns the attributes for [DataContainerApp].
func (ca *DataContainerApp) Attributes() dataContainerAppAttributes {
	return dataContainerAppAttributes{ref: terra.ReferenceDataResource(ca)}
}

// DataContainerAppArgs contains the configurations for azurerm_container_app.
type DataContainerAppArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Dapr: min=0
	Dapr []datacontainerapp.Dapr `hcl:"dapr,block" validate:"min=0"`
	// Identity: min=0
	Identity []datacontainerapp.Identity `hcl:"identity,block" validate:"min=0"`
	// Ingress: min=0
	Ingress []datacontainerapp.Ingress `hcl:"ingress,block" validate:"min=0"`
	// Registry: min=0
	Registry []datacontainerapp.Registry `hcl:"registry,block" validate:"min=0"`
	// Secret: min=0
	Secret []datacontainerapp.Secret `hcl:"secret,block" validate:"min=0"`
	// Template: min=0
	Template []datacontainerapp.Template `hcl:"template,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datacontainerapp.Timeouts `hcl:"timeouts,block"`
}
type dataContainerAppAttributes struct {
	ref terra.Reference
}

// ContainerAppEnvironmentId returns a reference to field container_app_environment_id of azurerm_container_app.
func (ca dataContainerAppAttributes) ContainerAppEnvironmentId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("container_app_environment_id"))
}

// CustomDomainVerificationId returns a reference to field custom_domain_verification_id of azurerm_container_app.
func (ca dataContainerAppAttributes) CustomDomainVerificationId() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("custom_domain_verification_id"))
}

// Id returns a reference to field id of azurerm_container_app.
func (ca dataContainerAppAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("id"))
}

// LatestRevisionFqdn returns a reference to field latest_revision_fqdn of azurerm_container_app.
func (ca dataContainerAppAttributes) LatestRevisionFqdn() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("latest_revision_fqdn"))
}

// LatestRevisionName returns a reference to field latest_revision_name of azurerm_container_app.
func (ca dataContainerAppAttributes) LatestRevisionName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("latest_revision_name"))
}

// Location returns a reference to field location of azurerm_container_app.
func (ca dataContainerAppAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_container_app.
func (ca dataContainerAppAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_container_app.
func (ca dataContainerAppAttributes) OutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ca.ref.Append("outbound_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_container_app.
func (ca dataContainerAppAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("resource_group_name"))
}

// RevisionMode returns a reference to field revision_mode of azurerm_container_app.
func (ca dataContainerAppAttributes) RevisionMode() terra.StringValue {
	return terra.ReferenceAsString(ca.ref.Append("revision_mode"))
}

// Tags returns a reference to field tags of azurerm_container_app.
func (ca dataContainerAppAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ca.ref.Append("tags"))
}

func (ca dataContainerAppAttributes) Dapr() terra.ListValue[datacontainerapp.DaprAttributes] {
	return terra.ReferenceAsList[datacontainerapp.DaprAttributes](ca.ref.Append("dapr"))
}

func (ca dataContainerAppAttributes) Identity() terra.ListValue[datacontainerapp.IdentityAttributes] {
	return terra.ReferenceAsList[datacontainerapp.IdentityAttributes](ca.ref.Append("identity"))
}

func (ca dataContainerAppAttributes) Ingress() terra.ListValue[datacontainerapp.IngressAttributes] {
	return terra.ReferenceAsList[datacontainerapp.IngressAttributes](ca.ref.Append("ingress"))
}

func (ca dataContainerAppAttributes) Registry() terra.ListValue[datacontainerapp.RegistryAttributes] {
	return terra.ReferenceAsList[datacontainerapp.RegistryAttributes](ca.ref.Append("registry"))
}

func (ca dataContainerAppAttributes) Secret() terra.ListValue[datacontainerapp.SecretAttributes] {
	return terra.ReferenceAsList[datacontainerapp.SecretAttributes](ca.ref.Append("secret"))
}

func (ca dataContainerAppAttributes) Template() terra.ListValue[datacontainerapp.TemplateAttributes] {
	return terra.ReferenceAsList[datacontainerapp.TemplateAttributes](ca.ref.Append("template"))
}

func (ca dataContainerAppAttributes) Timeouts() datacontainerapp.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datacontainerapp.TimeoutsAttributes](ca.ref.Append("timeouts"))
}
