// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataspringcloudservice "github.com/golingon/terraproviders/azurerm/3.65.0/dataspringcloudservice"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSpringCloudService creates a new instance of [DataSpringCloudService].
func NewDataSpringCloudService(name string, args DataSpringCloudServiceArgs) *DataSpringCloudService {
	return &DataSpringCloudService{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSpringCloudService)(nil)

// DataSpringCloudService represents the Terraform data resource azurerm_spring_cloud_service.
type DataSpringCloudService struct {
	Name string
	Args DataSpringCloudServiceArgs
}

// DataSource returns the Terraform object type for [DataSpringCloudService].
func (scs *DataSpringCloudService) DataSource() string {
	return "azurerm_spring_cloud_service"
}

// LocalName returns the local name for [DataSpringCloudService].
func (scs *DataSpringCloudService) LocalName() string {
	return scs.Name
}

// Configuration returns the configuration (args) for [DataSpringCloudService].
func (scs *DataSpringCloudService) Configuration() interface{} {
	return scs.Args
}

// Attributes returns the attributes for [DataSpringCloudService].
func (scs *DataSpringCloudService) Attributes() dataSpringCloudServiceAttributes {
	return dataSpringCloudServiceAttributes{ref: terra.ReferenceDataResource(scs)}
}

// DataSpringCloudServiceArgs contains the configurations for azurerm_spring_cloud_service.
type DataSpringCloudServiceArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ConfigServerGitSetting: min=0
	ConfigServerGitSetting []dataspringcloudservice.ConfigServerGitSetting `hcl:"config_server_git_setting,block" validate:"min=0"`
	// RequiredNetworkTrafficRules: min=0
	RequiredNetworkTrafficRules []dataspringcloudservice.RequiredNetworkTrafficRules `hcl:"required_network_traffic_rules,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataspringcloudservice.Timeouts `hcl:"timeouts,block"`
}
type dataSpringCloudServiceAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("name"))
}

// OutboundPublicIpAddresses returns a reference to field outbound_public_ip_addresses of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) OutboundPublicIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](scs.ref.Append("outbound_public_ip_addresses"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(scs.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_spring_cloud_service.
func (scs dataSpringCloudServiceAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](scs.ref.Append("tags"))
}

func (scs dataSpringCloudServiceAttributes) ConfigServerGitSetting() terra.ListValue[dataspringcloudservice.ConfigServerGitSettingAttributes] {
	return terra.ReferenceAsList[dataspringcloudservice.ConfigServerGitSettingAttributes](scs.ref.Append("config_server_git_setting"))
}

func (scs dataSpringCloudServiceAttributes) RequiredNetworkTrafficRules() terra.ListValue[dataspringcloudservice.RequiredNetworkTrafficRulesAttributes] {
	return terra.ReferenceAsList[dataspringcloudservice.RequiredNetworkTrafficRulesAttributes](scs.ref.Append("required_network_traffic_rules"))
}

func (scs dataSpringCloudServiceAttributes) Timeouts() dataspringcloudservice.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataspringcloudservice.TimeoutsAttributes](scs.ref.Append("timeouts"))
}
