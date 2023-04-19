// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	dataapplicationgateway "github.com/golingon/terraproviders/azurerm/3.52.0/dataapplicationgateway"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataApplicationGateway creates a new instance of [DataApplicationGateway].
func NewDataApplicationGateway(name string, args DataApplicationGatewayArgs) *DataApplicationGateway {
	return &DataApplicationGateway{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataApplicationGateway)(nil)

// DataApplicationGateway represents the Terraform data resource azurerm_application_gateway.
type DataApplicationGateway struct {
	Name string
	Args DataApplicationGatewayArgs
}

// DataSource returns the Terraform object type for [DataApplicationGateway].
func (ag *DataApplicationGateway) DataSource() string {
	return "azurerm_application_gateway"
}

// LocalName returns the local name for [DataApplicationGateway].
func (ag *DataApplicationGateway) LocalName() string {
	return ag.Name
}

// Configuration returns the configuration (args) for [DataApplicationGateway].
func (ag *DataApplicationGateway) Configuration() interface{} {
	return ag.Args
}

// Attributes returns the attributes for [DataApplicationGateway].
func (ag *DataApplicationGateway) Attributes() dataApplicationGatewayAttributes {
	return dataApplicationGatewayAttributes{ref: terra.ReferenceDataResource(ag)}
}

// DataApplicationGatewayArgs contains the configurations for azurerm_application_gateway.
type DataApplicationGatewayArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// BackendAddressPool: min=0
	BackendAddressPool []dataapplicationgateway.BackendAddressPool `hcl:"backend_address_pool,block" validate:"min=0"`
	// Identity: min=0
	Identity []dataapplicationgateway.Identity `hcl:"identity,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *dataapplicationgateway.Timeouts `hcl:"timeouts,block"`
}
type dataApplicationGatewayAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ag.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_application_gateway.
func (ag dataApplicationGatewayAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ag.ref.Append("tags"))
}

func (ag dataApplicationGatewayAttributes) BackendAddressPool() terra.ListValue[dataapplicationgateway.BackendAddressPoolAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.BackendAddressPoolAttributes](ag.ref.Append("backend_address_pool"))
}

func (ag dataApplicationGatewayAttributes) Identity() terra.ListValue[dataapplicationgateway.IdentityAttributes] {
	return terra.ReferenceAsList[dataapplicationgateway.IdentityAttributes](ag.ref.Append("identity"))
}

func (ag dataApplicationGatewayAttributes) Timeouts() dataapplicationgateway.TimeoutsAttributes {
	return terra.ReferenceAsSingle[dataapplicationgateway.TimeoutsAttributes](ag.ref.Append("timeouts"))
}
