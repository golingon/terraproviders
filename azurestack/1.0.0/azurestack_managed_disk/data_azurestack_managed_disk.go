// CODE GENERATED BY github.com/golingon/lingon. DO NOT EDIT.

package azurestack_managed_disk

import "github.com/golingon/lingon/pkg/terra"

// Data creates a new instance of [DataSource].
func Data(name string, args DataArgs) *DataSource {
	return &DataSource{
		Args: args,
		Name: name,
	}
}

var _ terra.DataSource = (*DataSource)(nil)

// DataSource represents the Terraform data resource azurestack_managed_disk.
type DataSource struct {
	Name string
	Args DataArgs
}

// DataSource returns the Terraform object type for [DataSource].
func (amd *DataSource) DataSource() string {
	return "azurestack_managed_disk"
}

// LocalName returns the local name for [DataSource].
func (amd *DataSource) LocalName() string {
	return amd.Name
}

// Configuration returns the configuration (args) for [DataSource].
func (amd *DataSource) Configuration() interface{} {
	return amd.Args
}

// Attributes returns the attributes for [DataSource].
func (amd *DataSource) Attributes() dataAzurestackManagedDiskAttributes {
	return dataAzurestackManagedDiskAttributes{ref: terra.ReferenceDataSource(amd)}
}

// DataArgs contains the configurations for azurestack_managed_disk.
type DataArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// Tags: map of string, optional
	Tags terra.MapValue[terra.StringValue] `hcl:"tags,attr"`
	// Timeouts: optional
	Timeouts *DataTimeouts `hcl:"timeouts,block"`
}

type dataAzurestackManagedDiskAttributes struct {
	ref terra.Reference
}

// CreateOption returns a reference to field create_option of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) CreateOption() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("create_option"))
}

// DiskSizeGb returns a reference to field disk_size_gb of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(amd.ref.Append("disk_size_gb"))
}

// Id returns a reference to field id of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("id"))
}

// ImageReferenceId returns a reference to field image_reference_id of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) ImageReferenceId() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("image_reference_id"))
}

// Name returns a reference to field name of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("os_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("resource_group_name"))
}

// SourceResourceId returns a reference to field source_resource_id of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) SourceResourceId() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("source_resource_id"))
}

// SourceUri returns a reference to field source_uri of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) SourceUri() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("source_uri"))
}

// StorageAccountId returns a reference to field storage_account_id of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("storage_account_id"))
}

// StorageAccountType returns a reference to field storage_account_type of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) StorageAccountType() terra.StringValue {
	return terra.ReferenceAsString(amd.ref.Append("storage_account_type"))
}

// Tags returns a reference to field tags of azurestack_managed_disk.
func (amd dataAzurestackManagedDiskAttributes) Tags() terra.MapValue[terra.StringValue] {
	return terra.ReferenceAsMap[terra.StringValue](amd.ref.Append("tags"))
}

func (amd dataAzurestackManagedDiskAttributes) Timeouts() DataTimeoutsAttributes {
	return terra.ReferenceAsSingle[DataTimeoutsAttributes](amd.ref.Append("timeouts"))
}
