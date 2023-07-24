// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package azurerm

import (
	datasnapshot "github.com/golingon/terraproviders/azurerm/3.66.0/datasnapshot"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataSnapshot creates a new instance of [DataSnapshot].
func NewDataSnapshot(name string, args DataSnapshotArgs) *DataSnapshot {
	return &DataSnapshot{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataSnapshot)(nil)

// DataSnapshot represents the Terraform data resource azurerm_snapshot.
type DataSnapshot struct {
	Name string
	Args DataSnapshotArgs
}

// DataSource returns the Terraform object type for [DataSnapshot].
func (s *DataSnapshot) DataSource() string {
	return "azurerm_snapshot"
}

// LocalName returns the local name for [DataSnapshot].
func (s *DataSnapshot) LocalName() string {
	return s.Name
}

// Configuration returns the configuration (args) for [DataSnapshot].
func (s *DataSnapshot) Configuration() interface{} {
	return s.Args
}

// Attributes returns the attributes for [DataSnapshot].
func (s *DataSnapshot) Attributes() dataSnapshotAttributes {
	return dataSnapshotAttributes{ref: terra.ReferenceDataResource(s)}
}

// DataSnapshotArgs contains the configurations for azurerm_snapshot.
type DataSnapshotArgs struct {
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// Name: string, required
	Name terra.StringValue `hcl:"name,attr" validate:"required"`
	// ResourceGroupName: string, required
	ResourceGroupName terra.StringValue `hcl:"resource_group_name,attr" validate:"required"`
	// EncryptionSettings: min=0
	EncryptionSettings []datasnapshot.EncryptionSettings `hcl:"encryption_settings,block" validate:"min=0"`
	// Timeouts: optional
	Timeouts *datasnapshot.Timeouts `hcl:"timeouts,block"`
}
type dataSnapshotAttributes struct {
	ref terra.Reference
}

// CreationOption returns a reference to field creation_option of azurerm_snapshot.
func (s dataSnapshotAttributes) CreationOption() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("creation_option"))
}

// DiskSizeGb returns a reference to field disk_size_gb of azurerm_snapshot.
func (s dataSnapshotAttributes) DiskSizeGb() terra.NumberValue {
	return terra.ReferenceAsNumber(s.ref.Append("disk_size_gb"))
}

// Id returns a reference to field id of azurerm_snapshot.
func (s dataSnapshotAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("id"))
}

// Name returns a reference to field name of azurerm_snapshot.
func (s dataSnapshotAttributes) Name() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("name"))
}

// OsType returns a reference to field os_type of azurerm_snapshot.
func (s dataSnapshotAttributes) OsType() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("os_type"))
}

// ResourceGroupName returns a reference to field resource_group_name of azurerm_snapshot.
func (s dataSnapshotAttributes) ResourceGroupName() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("resource_group_name"))
}

// SourceResourceId returns a reference to field source_resource_id of azurerm_snapshot.
func (s dataSnapshotAttributes) SourceResourceId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("source_resource_id"))
}

// SourceUri returns a reference to field source_uri of azurerm_snapshot.
func (s dataSnapshotAttributes) SourceUri() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("source_uri"))
}

// StorageAccountId returns a reference to field storage_account_id of azurerm_snapshot.
func (s dataSnapshotAttributes) StorageAccountId() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("storage_account_id"))
}

// TimeCreated returns a reference to field time_created of azurerm_snapshot.
func (s dataSnapshotAttributes) TimeCreated() terra.StringValue {
	return terra.ReferenceAsString(s.ref.Append("time_created"))
}

// TrustedLaunchEnabled returns a reference to field trusted_launch_enabled of azurerm_snapshot.
func (s dataSnapshotAttributes) TrustedLaunchEnabled() terra.BoolValue {
	return terra.ReferenceAsBool(s.ref.Append("trusted_launch_enabled"))
}

func (s dataSnapshotAttributes) EncryptionSettings() terra.ListValue[datasnapshot.EncryptionSettingsAttributes] {
	return terra.ReferenceAsList[datasnapshot.EncryptionSettingsAttributes](s.ref.Append("encryption_settings"))
}

func (s dataSnapshotAttributes) Timeouts() datasnapshot.TimeoutsAttributes {
	return terra.ReferenceAsSingle[datasnapshot.TimeoutsAttributes](s.ref.Append("timeouts"))
}
