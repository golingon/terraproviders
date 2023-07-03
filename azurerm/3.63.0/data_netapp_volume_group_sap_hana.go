// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datanetappvolumegroupsaphana "github.com/golingon/terraproviders/azurerm/3.63.0/datanetappvolumegroupsaphana"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataNetappVolumeGroupSapHana creates a new instance of [DataNetappVolumeGroupSapHana].
func NewDataNetappVolumeGroupSapHana(name string, args DataNetappVolumeGroupSapHanaArgs) *DataNetappVolumeGroupSapHana {
	return &DataNetappVolumeGroupSapHana{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataNetappVolumeGroupSapHana)(nil)

// DataNetappVolumeGroupSapHana represents the Terraform data resource azurerm_netapp_volume_group_sap_hana.
type DataNetappVolumeGroupSapHana struct {
	Name string
	Args DataNetappVolumeGroupSapHanaArgs
}

// DataSource returns the Terraform object type for [DataNetappVolumeGroupSapHana].
func (nvgsh *DataNetappVolumeGroupSapHana) DataSource() string {
	return "azurerm_netapp_volume_group_sap_hana"
}

// LocalName returns the local name for [DataNetappVolumeGroupSapHana].
func (nvgsh *DataNetappVolumeGroupSapHana) LocalName() string {
	return nvgsh.Name
}

// Configuration returns the configuration (args) for [DataNetappVolumeGroupSapHana].
func (nvgsh *DataNetappVolumeGroupSapHana) Configuration() interface{} {
	return nvgsh.Args
}

// Attributes returns the attributes for [DataNetappVolumeGroupSapHana].
func (nvgsh *DataNetappVolumeGroupSapHana) Attributes() dataNetappVolumeGroupSapHanaAttributes {
	return dataNetappVolumeGroupSapHanaAttributes{ref: terra.ReferenceDataResource(nvgsh)}
}

// DataNetappVolumeGroupSapHanaArgs contains the configurations for azurerm_netapp_volume_group_sap_hana.
type DataNetappVolumeGroupSapHanaArgs struct {
	// AccountName: string, required
	AccountName terra.StringValue `hcl:"account_name,attr" validate:"required"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Volume: min=0
	Volume []datanetappvolumegroupsaphana.Volume `hcl:"volume,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datanetappvolumegroupsaphana.Timeouts `hcl:"timeouts,block"`
}
type dataNetappVolumeGroupSapHanaAttributes struct {
	ref terra.Reference
}

// AccountName returns a reference to field account_name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) AccountName() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("account_name"))
}

// ApplicationIdentifier returns a reference to field application_identifier of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) ApplicationIdentifier() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("application_identifier"))
}

// GroupDescription returns a reference to field group_description of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) GroupDescription() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("group_description"))
}

// Id returns a reference to field id of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_netapp_volume_group_sap_hana.
func (nvgsh dataNetappVolumeGroupSapHanaAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(nvgsh.ref.Append("resource_group_name"))
}

func (nvgsh dataNetappVolumeGroupSapHanaAttributes) Volume() terra.ListValue[datanetappvolumegroupsaphana.VolumeAttributes] {
	return terra.ReferenceAsList[datanetappvolumegroupsaphana.VolumeAttributes](nvgsh.ref.Append("volume"))
}

func (nvgsh dataNetappVolumeGroupSapHanaAttributes) Timeouts() datanetappvolumegroupsaphana.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datanetappvolumegroupsaphana.TimeoutsAttributes](nvgsh.ref.Append("timeouts"))
}
