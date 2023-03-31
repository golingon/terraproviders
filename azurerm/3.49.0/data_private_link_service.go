// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataprivatelinkservice "github.com/golingon/terraproviders/azurerm/3.49.0/dataprivatelinkservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataPrivateLinkService(name string, args DataPrivateLinkServiceArgs) *DataPrivateLinkService {
	return &DataPrivateLinkService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPrivateLinkService)(nil)

type DataPrivateLinkService struct {
	Name string
	Args DataPrivateLinkServiceArgs
}

func (pls *DataPrivateLinkService) DataSource() string {
	return "azurerm_private_link_service"
}

func (pls *DataPrivateLinkService) LocalName() string {
	return pls.Name
}

func (pls *DataPrivateLinkService) Configuration() interface{} {
	return pls.Args
}

func (pls *DataPrivateLinkService) Attributes() dataPrivateLinkServiceAttributes {
	return dataPrivateLinkServiceAttributes{ref: terra.ReferenceDataResource(pls)}
}

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

func (pls dataPrivateLinkServiceAttributes) Alias() terra.StringValue {
	return terra.ReferenceString(pls.ref.Append("alias"))
}

func (pls dataPrivateLinkServiceAttributes) AutoApprovalSubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](pls.ref.Append("auto_approval_subscription_ids"))
}

func (pls dataPrivateLinkServiceAttributes) EnableProxyProtocol() terra.BoolValue {
	return terra.ReferenceBool(pls.ref.Append("enable_proxy_protocol"))
}

func (pls dataPrivateLinkServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pls.ref.Append("id"))
}

func (pls dataPrivateLinkServiceAttributes) LoadBalancerFrontendIpConfigurationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](pls.ref.Append("load_balancer_frontend_ip_configuration_ids"))
}

func (pls dataPrivateLinkServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceString(pls.ref.Append("location"))
}

func (pls dataPrivateLinkServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceString(pls.ref.Append("name"))
}

func (pls dataPrivateLinkServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceString(pls.ref.Append("resource_group_name"))
}

func (pls dataPrivateLinkServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceMap[terra.StringValue](pls.ref.Append("tags"))
}

func (pls dataPrivateLinkServiceAttributes) VisibilitySubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](pls.ref.Append("visibility_subscription_ids"))
}

func (pls dataPrivateLinkServiceAttributes) NatIpConfiguration() terra.ListValue[dataprivatelinkservice.NatIpConfigurationAttributes] {
	return terra.ReferenceList[dataprivatelinkservice.NatIpConfigurationAttributes](pls.ref.Append("nat_ip_configuration"))
}

func (pls dataPrivateLinkServiceAttributes) Timeouts() dataprivatelinkservice.TimeoutsAttributes {
	return terra.ReferenceSingle[dataprivatelinkservice.TimeoutsAttributes](pls.ref.Append("timeouts"))
}
