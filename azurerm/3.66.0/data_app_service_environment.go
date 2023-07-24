// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataappserviceenvironment "github.com/golingon/terraproviders/azurerm/3.66.0/dataappserviceenvironment"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataAppServiceEnvironment creates a new instance of [DataAppServiceEnvironment].
func NewDataAppServiceEnvironment(name string, args DataAppServiceEnvironmentArgs) *DataAppServiceEnvironment {
	return &DataAppServiceEnvironment{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataAppServiceEnvironment)(nil)

// DataAppServiceEnvironment represents the Terraform data resource azurerm_app_service_environment.
type DataAppServiceEnvironment struct {
	Name string
	Args DataAppServiceEnvironmentArgs
}

// DataSource returns the Terraform object type for [DataAppServiceEnvironment].
func (ase *DataAppServiceEnvironment) DataSource() string {
	return "azurerm_app_service_environment"
}

// LocalName returns the local name for [DataAppServiceEnvironment].
func (ase *DataAppServiceEnvironment) LocalName() string {
	return ase.Name
}

// Configuration returns the configuration (args) for [DataAppServiceEnvironment].
func (ase *DataAppServiceEnvironment) Configuration() interface{} {
	return ase.Args
}

// Attributes returns the attributes for [DataAppServiceEnvironment].
func (ase *DataAppServiceEnvironment) Attributes() dataAppServiceEnvironmentAttributes {
	return dataAppServiceEnvironmentAttributes{ref: terra.ReferenceDataResource(ase)}
}

// DataAppServiceEnvironmentArgs contains the configurations for azurerm_app_service_environment.
type DataAppServiceEnvironmentArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// ClusterSetting: min=0
	ClusterSetting []dataappserviceenvironment.ClusterSetting `hcl:"cluster_setting,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataappserviceenvironment.Timeouts `hcl:"timeouts,block"`
}
type dataAppServiceEnvironmentAttributes struct {
	ref terra.Reference
}

// FrontEndScaleFactor returns a reference to field front_end_scale_factor of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) FrontEndScaleFactor() terra.NumberValue {
	return terra.ReferenceAsNumber(ase.ref.Append("front_end_scale_factor"))
}

// Id returns a reference to field id of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("id"))
}

// InternalIpAddress returns a reference to field internal_ip_address of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) InternalIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("internal_ip_address"))
}

// Location returns a reference to field location of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("name"))
}

// OutboundIpAddresses returns a reference to field outbound_ip_addresses of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) OutboundIpAddresses() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ase.ref.Append("outbound_ip_addresses"))
}

// PricingTier returns a reference to field pricing_tier of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) PricingTier() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("pricing_tier"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("resource_group_name"))
}

// ServiceIpAddress returns a reference to field service_ip_address of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) ServiceIpAddress() terra.StringValue {
	return terra.ReferenceAsString(ase.ref.Append("service_ip_address"))
}

// Tags returns a reference to field tags of azurerm_app_service_environment.
func (ase dataAppServiceEnvironmentAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ase.ref.Append("tags"))
}

func (ase dataAppServiceEnvironmentAttributes) ClusterSetting() terra.ListValue[dataappserviceenvironment.ClusterSettingAttributes] {
	return terra.ReferenceAsList[dataappserviceenvironment.ClusterSettingAttributes](ase.ref.Append("cluster_setting"))
}

func (ase dataAppServiceEnvironmentAttributes) Timeouts() dataappserviceenvironment.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataappserviceenvironment.TimeoutsAttributes](ase.ref.Append("timeouts"))
}
