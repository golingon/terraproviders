// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_express_route_circuit

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_express_route_circuit.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (aerc *DataSource) DataSource() string {
	return "azurerm_express_route_circuit"
}

// LocalName returns the local name for [DataSource].
func (aerc *DataSource) LocalName() string {
	return aerc.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (aerc *DataSource) Configuration() interface{} {
	return aerc.Args
}

// Attributes returns the attributes for [DataSource].
func (aerc *DataSource) Attributes() dataAzurermExpressRouteCircuitAttributes {
	return dataAzurermExpressRouteCircuitAttributes{ref: terra.ReferenceDataSource(aerc)}
}

// DataArgs contains the configurations for azurerm_express_route_circuit.
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

type dataAzurermExpressRouteCircuitAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("resource_group_name"))
}

// ServiceKey returns a reference to field service_key of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) ServiceKey() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("service_key"))
}

// ServiceProviderProvisioningState returns a reference to field service_provider_provisioning_state of azurerm_express_route_circuit.
func (aerc dataAzurermExpressRouteCircuitAttributes) ServiceProviderProvisioningState() terra.StringValue {
	return terra.ReferenceAsString(aerc.ref.Append("service_provider_provisioning_state"))
}

func (aerc dataAzurermExpressRouteCircuitAttributes) Peerings() terra.ListValue[DataPeeringsAttributes] {
	return terra.ReferenceAsList[DataPeeringsAttributes](aerc.ref.Append("peerings"))
}

func (aerc dataAzurermExpressRouteCircuitAttributes) ServiceProviderProperties() terra.ListValue[DataServiceProviderPropertiesAttributes] {
	return terra.ReferenceAsList[DataServiceProviderPropertiesAttributes](aerc.ref.Append("service_provider_properties"))
}

func (aerc dataAzurermExpressRouteCircuitAttributes) Sku() terra.ListValue[DataSkuAttributes] {
	return terra.ReferenceAsList[DataSkuAttributes](aerc.ref.Append("sku"))
}

func (aerc dataAzurermExpressRouteCircuitAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](aerc.ref.Append("timeouts"))
}
