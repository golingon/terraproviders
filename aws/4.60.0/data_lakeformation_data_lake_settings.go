// CODE GENERATED BY github.com/volvo-cars/lingon. DO NOT EDIT.

package aws

import (
	datalakeformationdatalakesettings "github.com/golingon/terraproviders/aws/4.60.0/datalakeformationdatalakesettings"
	"github.com/volvo-cars/lingon/pkg/terra"
)

// NewDataLakeformationDataLakeSettings creates a new instance of [DataLakeformationDataLakeSettings].
func NewDataLakeformationDataLakeSettings(name string, args DataLakeformationDataLakeSettingsArgs) *DataLakeformationDataLakeSettings {
	return &DataLakeformationDataLakeSettings{
		Args: args,
		Name: name,
	}
}

var _ terra.DataResource = (*DataLakeformationDataLakeSettings)(nil)

// DataLakeformationDataLakeSettings represents the Terraform data resource aws_lakeformation_data_lake_settings.
type DataLakeformationDataLakeSettings struct {
	Name string
	Args DataLakeformationDataLakeSettingsArgs
}

// DataSource returns the Terraform object type for [DataLakeformationDataLakeSettings].
func (ldls *DataLakeformationDataLakeSettings) DataSource() string {
	return "aws_lakeformation_data_lake_settings"
}

// LocalName returns the local name for [DataLakeformationDataLakeSettings].
func (ldls *DataLakeformationDataLakeSettings) LocalName() string {
	return ldls.Name
}

// Configuration returns the configuration (args) for [DataLakeformationDataLakeSettings].
func (ldls *DataLakeformationDataLakeSettings) Configuration() interface{} {
	return ldls.Args
}

// Attributes returns the attributes for [DataLakeformationDataLakeSettings].
func (ldls *DataLakeformationDataLakeSettings) Attributes() dataLakeformationDataLakeSettingsAttributes {
	return dataLakeformationDataLakeSettingsAttributes{ref: terra.ReferenceDataResource(ldls)}
}

// DataLakeformationDataLakeSettingsArgs contains the configurations for aws_lakeformation_data_lake_settings.
type DataLakeformationDataLakeSettingsArgs struct {
	// CatalogId: string, optional
	CatalogId terra.StringValue `hcl:"catalog_id,attr"`
	// Id: string, optional
	Id terra.StringValue `hcl:"id,attr"`
	// CreateDatabaseDefaultPermissions: min=0
	CreateDatabaseDefaultPermissions []datalakeformationdatalakesettings.CreateDatabaseDefaultPermissions `hcl:"create_database_default_permissions,block" validate:"min=0"`
	// CreateTableDefaultPermissions: min=0
	CreateTableDefaultPermissions []datalakeformationdatalakesettings.CreateTableDefaultPermissions `hcl:"create_table_default_permissions,block" validate:"min=0"`
}
type dataLakeformationDataLakeSettingsAttributes struct {
	ref terra.Reference
}

// Admins returns a reference to field admins of aws_lakeformation_data_lake_settings.
func (ldls dataLakeformationDataLakeSettingsAttributes) Admins() terra.SetValue[terra.StringValue] {
	return terra.ReferenceAsSet[terra.StringValue](ldls.ref.Append("admins"))
}

// CatalogId returns a reference to field catalog_id of aws_lakeformation_data_lake_settings.
func (ldls dataLakeformationDataLakeSettingsAttributes) CatalogId() terra.StringValue {
	return terra.ReferenceAsString(ldls.ref.Append("catalog_id"))
}

// Id returns a reference to field id of aws_lakeformation_data_lake_settings.
func (ldls dataLakeformationDataLakeSettingsAttributes) Id() terra.StringValue {
	return terra.ReferenceAsString(ldls.ref.Append("id"))
}

// TrustedResourceOwners returns a reference to field trusted_resource_owners of aws_lakeformation_data_lake_settings.
func (ldls dataLakeformationDataLakeSettingsAttributes) TrustedResourceOwners() terra.ListValue[terra.StringValue] {
	return terra.ReferenceAsList[terra.StringValue](ldls.ref.Append("trusted_resource_owners"))
}

func (ldls dataLakeformationDataLakeSettingsAttributes) CreateDatabaseDefaultPermissions() terra.ListValue[datalakeformationdatalakesettings.CreateDatabaseDefaultPermissionsAttributes] {
	return terra.ReferenceAsList[datalakeformationdatalakesettings.CreateDatabaseDefaultPermissionsAttributes](ldls.ref.Append("create_database_default_permissions"))
}

func (ldls dataLakeformationDataLakeSettingsAttributes) CreateTableDefaultPermissions() terra.ListValue[datalakeformationdatalakesettings.CreateTableDefaultPermissionsAttributes] {
	return terra.ReferenceAsList[datalakeformationdatalakesettings.CreateTableDefaultPermissionsAttributes](ldls.ref.Append("create_table_default_permissions"))
}
