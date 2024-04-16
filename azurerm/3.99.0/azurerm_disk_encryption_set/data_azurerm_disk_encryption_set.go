// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurerm_disk_encryption_set

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurerm_disk_encryption_set.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (ades *DataSource) DataSource() string {
	return "azurerm_disk_encryption_set"
}

// LocalName returns the local name for [DataSource].
func (ades *DataSource) LocalName() string {
	return ades.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (ades *DataSource) Configuration() interface{} {
	return ades.Args
}

// Attributes returns the attributes for [DataSource].
func (ades *DataSource) Attributes() dataAzurermDiskEncryptionSetAttributes {
	return dataAzurermDiskEncryptionSetAttributes{ref: terra.ReferenceDataSource(ades)}
}

// DataArgs contains the configurations for azurerm_disk_encryption_set.
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

type dataAzurermDiskEncryptionSetAttributes struct {
	ref terra.Reference
}

// AutoKeyRotationEnabled returns a reference to field auto_key_rotation_enabled of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) AutoKeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(ades.ref.Append("auto_key_rotation_enabled"))
}

// Id returns a reference to field id of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("id"))
}

// KeyVaultKeyUrl returns a reference to field key_vault_key_url of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) KeyVaultKeyUrl() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("key_vault_key_url"))
}

// Location returns a reference to field location of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(ades.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_disk_encryption_set.
func (ades dataAzurermDiskEncryptionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](ades.ref.Append("tags"))
}

func (ades dataAzurermDiskEncryptionSetAttributes) Identity() terra.ListValue[DataIdentityAttributes] {
	return terra.ReferenceAsList[DataIdentityAttributes](ades.ref.Append("identity"))
}

func (ades dataAzurermDiskEncryptionSetAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](ades.ref.Append("timeouts"))
}
