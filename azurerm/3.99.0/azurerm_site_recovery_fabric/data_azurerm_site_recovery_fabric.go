// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_site_recovery_fabric

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_site_recovery_fabric.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (asrf *DataSource) DataSource() string {
	return "azurerm_site_recovery_fabric"
}

// LocalName returns the local name for [DataSource].
func (asrf *DataSource) LocalName() string {
	return asrf.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (asrf *DataSource) Configuration() interface{} {
	return asrf.Args
}

// Attributes returns the attributes for [DataSource].
func (asrf *DataSource) Attributes() dataAzurermSiteRecoveryFabricAttributes {
	return dataAzurermSiteRecoveryFabricAttributes{ref: terra.ReferenceDataSource(asrf)}
}

// DataArgs contains the configurations for azurerm_site_recovery_fabric.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurermSiteRecoveryFabricAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_fabric.
func (asrf dataAzurermSiteRecoveryFabricAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(asrf.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_site_recovery_fabric.
func (asrf dataAzurermSiteRecoveryFabricAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(asrf.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_site_recovery_fabric.
func (asrf dataAzurermSiteRecoveryFabricAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(asrf.ref.Append("name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_fabric.
func (asrf dataAzurermSiteRecoveryFabricAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(asrf.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_fabric.
func (asrf dataAzurermSiteRecoveryFabricAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(asrf.ref.Append("resource_group_name"))
}

func (asrf dataAzurermSiteRecoveryFabricAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](asrf.ref.Append("timeouts"))
}
