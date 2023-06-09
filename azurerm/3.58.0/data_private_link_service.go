// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatelinkservice "github.com/golingon/terraproviders/azurerm/3.58.0/dataprivatelinkservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataPrivateLinkService creates a new instance of [DataPrivateLinkService].
func NewDataPrivateLinkService(name string, args DataPrivateLinkServiceArgs) *DataPrivateLinkService {
	return &DataPrivateLinkService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateLinkService)(nil)

// DataPrivateLinkService represents the Terraform data resource azurerm_private_link_service.
type DataPrivateLinkService struct {
	Name string
	Args DataPrivateLinkServiceArgs
}

// DataSource returns the Terraform object type for [DataPrivateLinkService].
func (pls *DataPrivateLinkService) DataSource() string {
	return "azurerm_private_link_service"
}

// LocalName returns the local name for [DataPrivateLinkService].
func (pls *DataPrivateLinkService) LocalName() string {
	return pls.Name
}

// Configuration returns the configuration (args) for [DataPrivateLinkService].
func (pls *DataPrivateLinkService) Configuration() interface{} {
	return pls.Args
}

// Attributes returns the attributes for [DataPrivateLinkService].
func (pls *DataPrivateLinkService) Attributes() dataPrivateLinkServiceAttributes {
	return dataPrivateLinkServiceAttributes{ref: terra.ReferenceDataResource(pls)}
}

// DataPrivateLinkServiceArgs contains the configurations for azurerm_private_link_service.
type DataPrivateLinkServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// NatIpConfiguration: min=0
	NatIpConfiguration []dataprivatelinkservice.NatIpConfiguration `hcl:"nat_ip_configuration,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataprivatelinkservice.Timeouts `hcl:"timeouts,block"`
}
type dataPrivateLinkServiceAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("alias"))
}

// AutoApprovalSubscriptionIds returns a reference to field auto_approval_subscription_ids of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) AutoApprovalSubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pls.ref.Append("auto_approval_subscription_ids"))
}

// EnableProxyProtocol returns a reference to field enable_proxy_protocol of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) EnableProxyProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(pls.ref.Append("enable_proxy_protocol"))
}

// Id returns a reference to field id of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("id"))
}

// LoadBalancerFrontendIpConfigurationIds returns a reference to field load_balancer_frontend_ip_configuration_ids of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) LoadBalancerFrontendIpConfigurationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pls.ref.Append("load_balancer_frontend_ip_configuration_ids"))
}

// Location returns a reference to field location of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(pls.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](pls.ref.Append("tags"))
}

// VisibilitySubscriptionIds returns a reference to field visibility_subscription_ids of azurerm_private_link_service.
func (pls dataPrivateLinkServiceAttributes) VisibilitySubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](pls.ref.Append("visibility_subscription_ids"))
}

func (pls dataPrivateLinkServiceAttributes) NatIpConfiguration() terra.ListValue[dataprivatelinkservice.NatIpConfigurationAttributes] {
	return terra.ReferenceAsList[dataprivatelinkservice.NatIpConfigurationAttributes](pls.ref.Append("nat_ip_configuration"))
}

func (pls dataPrivateLinkServiceAttributes) Timeouts() dataprivatelinkservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataprivatelinkservice.TimeoutsAttributes](pls.ref.Append("timeouts"))
}
