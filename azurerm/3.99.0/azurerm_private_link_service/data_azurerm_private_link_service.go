// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_private_link_service

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_private_link_service.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (apls *DataSource) DataSource() string {
	return "azurerm_private_link_service"
}

// LocalName returns the local name for [DataSource].
func (apls *DataSource) LocalName() string {
	return apls.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (apls *DataSource) Configuration() interface{} {
	return apls.Args
}

// Attributes returns the attributes for [DataSource].
func (apls *DataSource) Attributes() dataAzurermPrivateLinkServiceAttributes {
	return dataAzurermPrivateLinkServiceAttributes{ref: terra.ReferenceDataSource(apls)}
}

// DataArgs contains the configurations for azurerm_private_link_service.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermPrivateLinkServiceAttributes struct {
	ref terra.Reference
}

// Alias returns a reference to field alias of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) Alias() terra.StringValue {
	return terra.ReferenceAsString(apls.ref.Append("alias"))
}

// AutoApprovalSubscriptionIds returns a reference to field auto_approval_subscription_ids of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) AutoApprovalSubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](apls.ref.Append("auto_approval_subscription_ids"))
}

// EnableProxyProtocol returns a reference to field enable_proxy_protocol of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) EnableProxyProtocol() terra.BoolValue {
	return terra.ReferenceAsBool(apls.ref.Append("enable_proxy_protocol"))
}

// Id returns a reference to field id of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(apls.ref.Append("id"))
}

// LoadBalancerFrontendIpConfigurationIds returns a reference to field load_balancer_frontend_ip_configuration_ids of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) LoadBalancerFrontendIpConfigurationIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](apls.ref.Append("load_balancer_frontend_ip_configuration_ids"))
}

// Location returns a reference to field location of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(apls.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(apls.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(apls.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](apls.ref.Append("tags"))
}

// VisibilitySubscriptionIds returns a reference to field visibility_subscription_ids of azurerm_private_link_service.
func (apls dataAzurermPrivateLinkServiceAttributes) VisibilitySubscriptionIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](apls.ref.Append("visibility_subscription_ids"))
}

func (apls dataAzurermPrivateLinkServiceAttributes) NatIpConfiguration() terra.ListValue[DataNatIpConfigurationAttributes] {
	return terra.ReferenceAsList[DataNatIpConfigurationAttributes](apls.ref.Append("nat_ip_configuration"))
}

func (apls dataAzurermPrivateLinkServiceAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](apls.ref.Append("timeouts"))
}
