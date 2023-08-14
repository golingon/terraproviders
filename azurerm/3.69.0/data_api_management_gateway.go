// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapimanagementgateway "github.com/golingon/terraproviders/azurerm/3.69.0/dataapimanagementgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApiManagementGateway creates a new instance of [DataApiManagementGateway].
func NewDataApiManagementGateway(name string, args DataApiManagementGatewayArgs) *DataApiManagementGateway {
	return &DataApiManagementGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApiManagementGateway)(nil)

// DataApiManagementGateway represents the Terraform data resource azurerm_api_management_gateway.
type DataApiManagementGateway struct {
	Name string
	Args DataApiManagementGatewayArgs
}

// DataSource returns the Terraform object type for [DataApiManagementGateway].
func (amg *DataApiManagementGateway) DataSource() string {
	return "azurerm_api_management_gateway"
}

// LocalName returns the local name for [DataApiManagementGateway].
func (amg *DataApiManagementGateway) LocalName() string {
	return amg.Name
}

// Configuration returns the configuration (args) for [DataApiManagementGateway].
func (amg *DataApiManagementGateway) Configuration() interface{} {
	return amg.Args
}

// Attributes returns the attributes for [DataApiManagementGateway].
func (amg *DataApiManagementGateway) Attributes() dataApiManagementGatewayAttributes {
	return dataApiManagementGatewayAttributes{ref: terra.ReferenceDataResource(amg)}
}

// DataApiManagementGatewayArgs contains the configurations for azurerm_api_management_gateway.
type DataApiManagementGatewayArgs struct {
	// ApiManagementId: string, required
	ApiManagementId terra.StringValue `hcl:"api_management_id,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// LocationData: min=0
	LocationData []dataapimanagementgateway.LocationData `hcl:"location_data,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataapimanagementgateway.Timeouts `hcl:"timeouts,block"`
}
type dataApiManagementGatewayAttributes struct {
	ref terra.Reference
}

// ApiManagementId returns a reference to field api_management_id of azurerm_api_management_gateway.
func (amg dataApiManagementGatewayAttributes) ApiManagementId() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("api_management_id"))
}

// Description returns a reference to field description of azurerm_api_management_gateway.
func (amg dataApiManagementGatewayAttributes) Description() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("description"))
}

// Id returns a reference to field id of azurerm_api_management_gateway.
func (amg dataApiManagementGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_api_management_gateway.
func (amg dataApiManagementGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amg.ref.Append("name"))
}

func (amg dataApiManagementGatewayAttributes) LocationData() terra.ListValue[dataapimanagementgateway.LocationDataAttributes] {
	return terra.ReferenceAsList[dataapimanagementgateway.LocationDataAttributes](amg.ref.Append("location_data"))
}

func (amg dataApiManagementGatewayAttributes) Timeouts() dataapimanagementgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapimanagementgateway.TimeoutsAttributes](amg.ref.Append("timeouts"))
}
