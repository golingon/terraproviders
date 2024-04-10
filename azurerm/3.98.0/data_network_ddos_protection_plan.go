// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm

import (
	"github.com/golingon/lingon/pkg/terra"
	datanetworkddosprotectionplan "github.com/golingon/terraproviders/azurerm/3.98.0/datanetworkddosprotectionplan"
)

// NewDataNetworkDdosProtectionPlan creates a new instance of [DataNetworkDdosProtectionPlan].
func NewDataNetworkDdosProtectionPlan(name string, args DataNetworkDdosProtectionPlanArgs) *DataNetworkDdosProtectionPlan {
	return &DataNetworkDdosProtectionPlan{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetworkDdosProtectionPlan)(nil)

// DataNetworkDdosProtectionPlan represents the Terraform data resource azurerm_network_ddos_protection_plan.
type DataNetworkDdosProtectionPlan struct {
	Name string
	Args DataNetworkDdosProtectionPlanArgs
}

// DataSource returns the Terraform object type for [DataNetworkDdosProtectionPlan].
func (ndpp *DataNetworkDdosProtectionPlan) DataSource() string {
	return "azurerm_network_ddos_protection_plan"
}

// LocalName returns the local name for [DataNetworkDdosProtectionPlan].
func (ndpp *DataNetworkDdosProtectionPlan) LocalName() string {
	return ndpp.Name
}

// Configuration returns the configuration (args) for [DataNetworkDdosProtectionPlan].
func (ndpp *DataNetworkDdosProtectionPlan) Configuration() interface{} {
	return ndpp.Args
}

// Attributes returns the attributes for [DataNetworkDdosProtectionPlan].
func (ndpp *DataNetworkDdosProtectionPlan) Attributes() dataNetworkDdosProtectionPlanAttributes {
	return dataNetworkDdosProtectionPlanAttributes{ref: terra.ReferenceDataResource(ndpp)}
}

// DataNetworkDdosProtectionPlanArgs contains the configurations for azurerm_network_ddos_protection_plan.
type DataNetworkDdosProtectionPlanArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datanetworkddosprotectionplan.Timeouts `hcl:"timeouts,block"`
}
type dataNetworkDdosProtectionPlanAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ndpp.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ndpp.ref.Append("tags"))
}

// VirtualNetworkIds returns a reference to field virtual_network_ids of azurerm_network_ddos_protection_plan.
func (ndpp dataNetworkDdosProtectionPlanAttributes) VirtualNetworkIds() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ndpp.ref.Append("virtual_network_ids"))
}

func (ndpp dataNetworkDdosProtectionPlanAttributes) Timeouts() datanetworkddosprotectionplan.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetworkddosprotectionplan.TimeoutsAttributes](ndpp.ref.Append("timeouts"))
}
