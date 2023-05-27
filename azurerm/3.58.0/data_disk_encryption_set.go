// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datadiskencryptionset "github.com/golingon/terraproviders/azurerm/3.58.0/datadiskencryptionset"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataDiskEncryptionSet creates a new instance of [DataDiskEncryptionSet].
func NewDataDiskEncryptionSet(name string, args DataDiskEncryptionSetArgs) *DataDiskEncryptionSet {
	return &DataDiskEncryptionSet{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataDiskEncryptionSet)(nil)

// DataDiskEncryptionSet represents the Terraform data resource azurerm_disk_encryption_set.
type DataDiskEncryptionSet struct {
	Name string
	Args DataDiskEncryptionSetArgs
}

// DataSource returns the Terraform object type for [DataDiskEncryptionSet].
func (des *DataDiskEncryptionSet) DataSource() string {
	return "azurerm_disk_encryption_set"
}

// LocalName returns the local name for [DataDiskEncryptionSet].
func (des *DataDiskEncryptionSet) LocalName() string {
	return des.Name
}

// Configuration returns the configuration (args) for [DataDiskEncryptionSet].
func (des *DataDiskEncryptionSet) Configuration() interface{} {
	return des.Args
}

// Attributes returns the attributes for [DataDiskEncryptionSet].
func (des *DataDiskEncryptionSet) Attributes() dataDiskEncryptionSetAttributes {
	return dataDiskEncryptionSetAttributes{ref: terra.ReferenceDataResource(des)}
}

// DataDiskEncryptionSetArgs contains the configurations for azurerm_disk_encryption_set.
type DataDiskEncryptionSetArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Timeouts: optional
	Timeouts *datadiskencryptionset.Timeouts `hcl:"timeouts,block"`
}
type dataDiskEncryptionSetAttributes struct {
	ref terra.Reference
}

// AutoKeyRotationEnabled returns a reference to field auto_key_rotation_enabled of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) AutoKeyRotationEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(des.ref.Append("auto_key_rotation_enabled"))
}

// Id returns a reference to field id of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("id"))
}

// Location returns a reference to field location of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) Location() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("location"))
}

// Name returns a reference to field name of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("name"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(des.ref.Append("resource_group_name"))
}

// Tags returns a reference to field tags of azurerm_disk_encryption_set.
func (des dataDiskEncryptionSetAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](des.ref.Append("tags"))
}

func (des dataDiskEncryptionSetAttributes) Timeouts() datadiskencryptionset.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datadiskencryptionset.TimeoutsAttributes](des.ref.Append("timeouts"))
}
