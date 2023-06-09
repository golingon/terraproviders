// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalbbackendaddresspool "github.com/golingon/terraproviders/azurerm/3.55.0/datalbbackendaddresspool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLbBackendAddressPool creates a new instance of [DataLbBackendAddressPool].
func NewDataLbBackendAddressPool(name string, args DataLbBackendAddressPoolArgs) *DataLbBackendAddressPool {
	return &DataLbBackendAddressPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbBackendAddressPool)(nil)

// DataLbBackendAddressPool represents the Terraform data resource azurerm_lb_backend_address_pool.
type DataLbBackendAddressPool struct {
	Name string
	Args DataLbBackendAddressPoolArgs
}

// DataSource returns the Terraform object type for [DataLbBackendAddressPool].
func (lbap *DataLbBackendAddressPool) DataSource() string {
	return "azurerm_lb_backend_address_pool"
}

// LocalName returns the local name for [DataLbBackendAddressPool].
func (lbap *DataLbBackendAddressPool) LocalName() string {
	return lbap.Name
}

// Configuration returns the configuration (args) for [DataLbBackendAddressPool].
func (lbap *DataLbBackendAddressPool) Configuration() interface{} {
	return lbap.Args
}

// Attributes returns the attributes for [DataLbBackendAddressPool].
func (lbap *DataLbBackendAddressPool) Attributes() dataLbBackendAddressPoolAttributes {
	return dataLbBackendAddressPoolAttributes{ref: terra.ReferenceDataResource(lbap)}
}

// DataLbBackendAddressPoolArgs contains the configurations for azurerm_lb_backend_address_pool.
type DataLbBackendAddressPoolArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// LoadbalancerId: string, required
	LoadbalancerId terra.StringValue `hcl:"loadbalancer_id,attr" validate:"required"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// BackendAddress: min=0
	BackendAddress []datalbbackendaddresspool.BackendAddress `hcl:"backend_address,block" validate:"min=0"`
	// BackendIpConfigurations: min=0
	BackendIpConfigurations []datalbbackendaddresspool.BackendIpConfigurations `hcl:"backend_ip_configurations,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datalbbackendaddresspool.Timeouts `hcl:"timeouts,block"`
}
type dataLbBackendAddressPoolAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("id"))
}

// InboundNatRules returns a reference to field inbound_nat_rules of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) InboundNatRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("inbound_nat_rules"))
}

// LoadBalancingRules returns a reference to field load_balancing_rules of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) LoadBalancingRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("load_balancing_rules"))
}

// LoadbalancerId returns a reference to field loadbalancer_id of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("loadbalancer_id"))
}

// Name returns a reference to field name of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(lbap.ref.Append("name"))
}

// OutboundRules returns a reference to field outbound_rules of azurerm_lb_backend_address_pool.
func (lbap dataLbBackendAddressPoolAttributes) OutboundRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](lbap.ref.Append("outbound_rules"))
}

func (lbap dataLbBackendAddressPoolAttributes) BackendAddress() terra.ListValue[datalbbackendaddresspool.BackendAddressAttributes] {
	return terra.ReferenceAsList[datalbbackendaddresspool.BackendAddressAttributes](lbap.ref.Append("backend_address"))
}

func (lbap dataLbBackendAddressPoolAttributes) BackendIpConfigurations() terra.ListValue[datalbbackendaddresspool.BackendIpConfigurationsAttributes] {
	return terra.ReferenceAsList[datalbbackendaddresspool.BackendIpConfigurationsAttributes](lbap.ref.Append("backend_ip_configurations"))
}

func (lbap dataLbBackendAddressPoolAttributes) Timeouts() datalbbackendaddresspool.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datalbbackendaddresspool.TimeoutsAttributes](lbap.ref.Append("timeouts"))
}
