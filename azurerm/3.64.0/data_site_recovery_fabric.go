// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasiterecoveryfabric "github.com/golingon/terraproviders/azurerm/3.64.0/datasiterecoveryfabric"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSiteRecoveryFabric creates a new instance of [DataSiteRecoveryFabric].
func NewDataSiteRecoveryFabric(name string, args DataSiteRecoveryFabricArgs) *DataSiteRecoveryFabric {
	return &DataSiteRecoveryFabric{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSiteRecoveryFabric)(nil)

// DataSiteRecoveryFabric represents the Terraform data resource azurerm_site_recovery_fabric.
type DataSiteRecoveryFabric struct {
	Name string
	Args DataSiteRecoveryFabricArgs
}

// DataSource returns the Terraform object type for [DataSiteRecoveryFabric].
func (srf *DataSiteRecoveryFabric) DataSource() string {
	return "azurerm_site_recovery_fabric"
}

// LocalName returns the local name for [DataSiteRecoveryFabric].
func (srf *DataSiteRecoveryFabric) LocalName() string {
	return srf.Name
}

// Configuration returns the configuration (args) for [DataSiteRecoveryFabric].
func (srf *DataSiteRecoveryFabric) Configuration() interface{} {
	return srf.Args
}

// Attributes returns the attributes for [DataSiteRecoveryFabric].
func (srf *DataSiteRecoveryFabric) Attributes() dataSiteRecoveryFabricAttributes {
	return dataSiteRecoveryFabricAttributes{ref: terra.ReferenceDataResource(srf)}
}

// DataSiteRecoveryFabricArgs contains the configurations for azurerm_site_recovery_fabric.
type DataSiteRecoveryFabricArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasiterecoveryfabric.Timeouts `hcl:"timeouts,block"`
}
type dataSiteRecoveryFabricAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_fabric.
func (srf dataSiteRecoveryFabricAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_site_recovery_fabric.
func (srf dataSiteRecoveryFabricAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_site_recovery_fabric.
func (srf dataSiteRecoveryFabricAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_fabric.
func (srf dataSiteRecoveryFabricAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_fabric.
func (srf dataSiteRecoveryFabricAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srf.ref.Append("resource_group_name"))
}

func (srf dataSiteRecoveryFabricAttributes) Timeouts() datasiterecoveryfabric.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasiterecoveryfabric.TimeoutsAttributes](srf.ref.Append("timeouts"))
}
