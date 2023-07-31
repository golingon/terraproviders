// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datarecoveryservicesvault "github.com/golingon/terraproviders/azurerm/3.67.0/datarecoveryservicesvault"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataRecoveryServicesVault creates a new instance of [DataRecoveryServicesVault].
func NewDataRecoveryServicesVault(name string, args DataRecoveryServicesVaultArgs) *DataRecoveryServicesVault {
	return &DataRecoveryServicesVault{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataRecoveryServicesVault)(nil)

// DataRecoveryServicesVault represents the Terraform data resource azurerm_recovery_services_vault.
type DataRecoveryServicesVault struct {
	Name string
	Args DataRecoveryServicesVaultArgs
}

// DataSource returns the Terraform object type for [DataRecoveryServicesVault].
func (rsv *DataRecoveryServicesVault) DataSource() string {
	return "azurerm_recovery_services_vault"
}

// LocalName returns the local name for [DataRecoveryServicesVault].
func (rsv *DataRecoveryServicesVault) LocalName() string {
	return rsv.Name
}

// Configuration returns the configuration (args) for [DataRecoveryServicesVault].
func (rsv *DataRecoveryServicesVault) Configuration() interface{} {
	return rsv.Args
}

// Attributes returns the attributes for [DataRecoveryServicesVault].
func (rsv *DataRecoveryServicesVault) Attributes() dataRecoveryServicesVaultAttributes {
	return dataRecoveryServicesVaultAttributes{ref: terra.ReferenceDataResource(rsv)}
}

// DataRecoveryServicesVaultArgs contains the configurations for azurerm_recovery_services_vault.
type DataRecoveryServicesVaultArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datarecoveryservicesvault.Timeouts `hcl:"timeouts,block"`
}
type dataRecoveryServicesVaultAttributes struct {
	ref terra.Reference
}

// Id returns a reference to field id of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("resource_group_name"))
}

// Sku returns a reference to field sku of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) Sku() terra.StringValue {
	return terra.ReferenceAsString(rsv.ref.Append("sku"))
}

// Tags returns a reference to field tags of azurerm_recovery_services_vault.
func (rsv dataRecoveryServicesVaultAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](rsv.ref.Append("tags"))
}

func (rsv dataRecoveryServicesVaultAttributes) Timeouts() datarecoveryservicesvault.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datarecoveryservicesvault.TimeoutsAttributes](rsv.ref.Append("timeouts"))
}
