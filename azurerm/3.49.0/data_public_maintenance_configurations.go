// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datapublicmaintenanceconfigurations "github.com/golingon/terraproviders/azurerm/3.49.0/datapublicmaintenanceconfigurations"
	"github.com/volvo-cars/lingon/pkg/terra"
)

func NewDataPublicMaintenanceConfigurations(name string, args DataPublicMaintenanceConfigurationsArgs) *DataPublicMaintenanceConfigurations {
	return &DataPublicMaintenanceConfigurations{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataPublicMaintenanceConfigurations)(nil)

type DataPublicMaintenanceConfigurations struct {
	Name string
	Args DataPublicMaintenanceConfigurationsArgs
}

func (pmc *DataPublicMaintenanceConfigurations) DataSource() string {
	return "azurerm_public_maintenance_configurations"
}

func (pmc *DataPublicMaintenanceConfigurations) LocalName() string {
	return pmc.Name
}

func (pmc *DataPublicMaintenanceConfigurations) Configuration() interface{} {
	return pmc.Args
}

func (pmc *DataPublicMaintenanceConfigurations) Attributes() dataPublicMaintenanceConfigurationsAttributes {
	return dataPublicMaintenanceConfigurationsAttributes{ref: terra.ReferenceDataResource(pmc)}
}

type DataPublicMaintenanceConfigurationsArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Location: string, optional
	Location terra.StringValue `hcl:"location,attr"`
	// RecurEvery: string, optional
	RecurEvery terra.StringValue `hcl:"recur_every,attr"`
	// Scope: string, optional
	Scope terra.StringValue `hcl:"scope,attr"`
	// Configs: min=0
	Configs []datapublicmaintenanceconfigurations.Configs `hcl:"configs,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datapublicmaintenanceconfigurations.Timeouts `hcl:"timeouts,block"`
}
type dataPublicMaintenanceConfigurationsAttributes struct {
	ref terra.Reference
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) Id() terra.StringValue {
	return terra.ReferenceString(pmc.ref.Append("id"))
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) Location() terra.StringValue {
	return terra.ReferenceString(pmc.ref.Append("location"))
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) RecurEvery() terra.StringValue {
	return terra.ReferenceString(pmc.ref.Append("recur_every"))
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) Scope() terra.StringValue {
	return terra.ReferenceString(pmc.ref.Append("scope"))
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) Configs() terra.ListValue[datapublicmaintenanceconfigurations.ConfigsAttributes] {
	return terra.ReferenceList[datapublicmaintenanceconfigurations.ConfigsAttributes](pmc.ref.Append("configs"))
}

func (pmc dataPublicMaintenanceConfigurationsAttributes) Timeouts() datapublicmaintenanceconfigurations.TimeoutsAttributes {
	return terra.ReferenceSingle[datapublicmaintenanceconfigurations.TimeoutsAttributes](pmc.ref.Append("timeouts"))
}
