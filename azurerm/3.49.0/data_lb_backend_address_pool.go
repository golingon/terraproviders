// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datalbbackendaddresspool "github.com/golingon/terraproviders/azurerm/3.49.0/datalbbackendaddresspool"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataLbBackendAddressPool(name string, args DataLbBackendAddressPoolArgs) *DataLbBackendAddressPool {
	return &DataLbBackendAddressPool{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLbBackendAddressPool)(nil)

type DataLbBackendAddressPool struct {
	Name string
	Args DataLbBackendAddressPoolArgs
}

func (lbap *DataLbBackendAddressPool) DataSource() string {
	return "azurerm_lb_backend_address_pool"
}

func (lbap *DataLbBackendAddressPool) LocalName() string {
	return lbap.Name
}

func (lbap *DataLbBackendAddressPool) Configuration() interface{} {
	return lbap.Args
}

func (lbap *DataLbBackendAddressPool) Attributes() dataLbBackendAddressPoolAttributes {
	return dataLbBackendAddressPoolAttributes{ref: terra.ReferenceDataResource(lbap)}
}

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

func (lbap dataLbBackendAddressPoolAttributes) Id() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("id"))
}

func (lbap dataLbBackendAddressPoolAttributes) InboundNatRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("inbound_nat_rules"))
}

func (lbap dataLbBackendAddressPoolAttributes) LoadBalancingRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("load_balancing_rules"))
}

func (lbap dataLbBackendAddressPoolAttributes) LoadbalancerId() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("loadbalancer_id"))
}

func (lbap dataLbBackendAddressPoolAttributes) Name() terra.StringValue {
	return terra.ReferenceString(lbap.ref.Append("name"))
}

func (lbap dataLbBackendAddressPoolAttributes) OutboundRules() terra.ListValue[terra.StringValue] {
	return terra.ReferenceList[terra.StringValue](lbap.ref.Append("outbound_rules"))
}

func (lbap dataLbBackendAddressPoolAttributes) BackendAddress() terra.ListValue[datalbbackendaddresspool.BackendAddressAttributes] {
	return terra.ReferenceList[datalbbackendaddresspool.BackendAddressAttributes](lbap.ref.Append("backend_address"))
}

func (lbap dataLbBackendAddressPoolAttributes) BackendIpConfigurations() terra.ListValue[datalbbackendaddresspool.BackendIpConfigurationsAttributes] {
	return terra.ReferenceList[datalbbackendaddresspool.BackendIpConfigurationsAttributes](lbap.ref.Append("backend_ip_configurations"))
}

func (lbap dataLbBackendAddressPoolAttributes) Timeouts() datalbbackendaddresspool.TimeoutsAttributes {
	return terra.ReferenceSingle[datalbbackendaddresspool.TimeoutsAttributes](lbap.ref.Append("timeouts"))
}
