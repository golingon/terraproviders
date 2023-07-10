// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasiterecoveryprotectioncontainer "github.com/golingon/terraproviders/azurerm/3.64.0/datasiterecoveryprotectioncontainer"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSiteRecoveryProtectionContainer creates a new instance of [DataSiteRecoveryProtectionContainer].
func NewDataSiteRecoveryProtectionContainer(name string, args DataSiteRecoveryProtectionContainerArgs) *DataSiteRecoveryProtectionContainer {
	return &DataSiteRecoveryProtectionContainer{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSiteRecoveryProtectionContainer)(nil)

// DataSiteRecoveryProtectionContainer represents the Terraform data resource azurerm_site_recovery_protection_container.
type DataSiteRecoveryProtectionContainer struct {
	Name string
	Args DataSiteRecoveryProtectionContainerArgs
}

// DataSource returns the Terraform object type for [DataSiteRecoveryProtectionContainer].
func (srpc *DataSiteRecoveryProtectionContainer) DataSource() string {
	return "azurerm_site_recovery_protection_container"
}

// LocalName returns the local name for [DataSiteRecoveryProtectionContainer].
func (srpc *DataSiteRecoveryProtectionContainer) LocalName() string {
	return srpc.Name
}

// Configuration returns the configuration (args) for [DataSiteRecoveryProtectionContainer].
func (srpc *DataSiteRecoveryProtectionContainer) Configuration() interface{} {
	return srpc.Args
}

// Attributes returns the attributes for [DataSiteRecoveryProtectionContainer].
func (srpc *DataSiteRecoveryProtectionContainer) Attributes() dataSiteRecoveryProtectionContainerAttributes {
	return dataSiteRecoveryProtectionContainerAttributes{ref: terra.ReferenceDataResource(srpc)}
}

// DataSiteRecoveryProtectionContainerArgs contains the configurations for azurerm_site_recovery_protection_container.
type DataSiteRecoveryProtectionContainerArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// RecoveryFabricName: string, required
	RecoveryFabricName terra.StringValue `hcl:"recovery_fabric_name,attr" validate:"required"`
	// RecoveryVaultName: string, required
	RecoveryVaultName terra.StringValue `hcl:"recovery_vault_name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datasiterecoveryprotectioncontainer.Timeouts `hcl:"timeouts,block"`
}
type dataSiteRecoveryProtectionContainerAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_site_recovery_protection_container.
func (srpc dataSiteRecoveryProtectionContainerAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_site_recovery_protection_container.
func (srpc dataSiteRecoveryProtectionContainerAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("name"))
}

// RecoveryFabricName returns a reference to field recovery_fabric_name of azurerm_site_recovery_protection_container.
func (srpc dataSiteRecoveryProtectionContainerAttributes) RecoveryFabricName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("recovery_fabric_name"))
}

// RecoveryVaultName returns a reference to field recovery_vault_name of azurerm_site_recovery_protection_container.
func (srpc dataSiteRecoveryProtectionContainerAttributes) RecoveryVaultName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("recovery_vault_name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_site_recovery_protection_container.
func (srpc dataSiteRecoveryProtectionContainerAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(srpc.ref.Append("resource_group_name"))
}

func (srpc dataSiteRecoveryProtectionContainerAttributes) Timeouts() datasiterecoveryprotectioncontainer.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasiterecoveryprotectioncontainer.TimeoutsAttributes](srpc.ref.Append("timeouts"))
}
